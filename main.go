package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

//go:embed commands rules
var embeddedFS embed.FS

const (
	commandsSourceDir = "commands"
	rulesSourceDir    = "rules"
	cursorHomeDir     = ".cursor"
	commandsTargetDir = "commands"
	rulesTargetBase   = "rules/mozi"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case "install":
		runInstall(os.Args[2:])
	case "-h", "--help", "help":
		printHelp()
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command '%s'\n\n", command)
		printHelp()
		os.Exit(1)
	}
}

// runInstall runs the install subcommand
func runInstall(args []string) {
	// Parse install command arguments
	targetBaseDir, err := parseInstallArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		printInstallHelp()
		os.Exit(1)
	}

	cursorDir := filepath.Join(targetBaseDir, cursorHomeDir)
	commandsTarget := filepath.Join(cursorDir, commandsTargetDir)

	// Create target directories
	if err := os.MkdirAll(commandsTarget, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating %s: %v\n", commandsTarget, err)
		os.Exit(1)
	}

	fmt.Printf("Installing commands and rules to %s\n", cursorDir)
	fmt.Println()

	// Install commands and collect role-rules mapping
	roleRulesMap, commandsCount, err := installCommands(commandsSourceDir, commandsTarget)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error installing commands: %v\n", err)
		os.Exit(1)
	}

	// Install rules for each role
	totalRulesCount := 0
	for role, rules := range roleRulesMap {
		rulesTarget := filepath.Join(cursorDir, rulesTargetBase, role)
		rulesCount, err := installRulesForRole(rulesSourceDir, rulesTarget, role, rules)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error installing rules for role %s: %v\n", role, err)
			os.Exit(1)
		}
		totalRulesCount += rulesCount
	}

	fmt.Println()
	fmt.Printf("✓ Successfully installed %d commands and %d rules to %s\n", commandsCount, totalRulesCount, cursorDir)
}

// installCommands installs command files and returns a map of role -> rules list
func installCommands(sourceDir, targetDir string) (map[string][]string, int, error) {
	roleRulesMap := make(map[string][]string)
	count := 0

	// Walk embedded filesystem
	err := fs.WalkDir(embeddedFS, sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Only process .md files
		ext := filepath.Ext(path)
		if ext != ".md" {
			return nil
		}

		// Extract role name from filename (without extension)
		roleName := strings.TrimSuffix(filepath.Base(path), ext)

		// Read command file to extract related rules
		rules, err := extractRulesFromCommand(path)
		if err != nil {
			return fmt.Errorf("error reading command file %s: %v", path, err)
		}

		// Store role-rules mapping
		roleRulesMap[roleName] = rules

		// Calculate relative path
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		// Target file path
		targetPath := filepath.Join(targetDir, relPath)

		// Create target file directory
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return fmt.Errorf("error creating target directory: %v", err)
		}

		// Copy file with conflict handling
		err = copyFileWithConflictHandling(path, targetPath)
		if err == ErrSkipped {
			// File was skipped (identical or user chose to skip)
			return nil
		}
		if err != nil {
			return fmt.Errorf("error copying file %s: %v", path, err)
		}

		fmt.Printf("  ✓ Installed command: %s\n", relPath)
		count++

		return nil
	})

	return roleRulesMap, count, err
}

// extractRulesFromCommand extracts rule names from a command file
func extractRulesFromCommand(filePath string) ([]string, error) {
	file, err := embeddedFS.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rules []string
	inRelatedRules := false
	// Match lines like "- rule-name" or "- rule-name.mdc" (with optional .mdc extension)
	rulePattern := regexp.MustCompile(`^-\s+([a-z0-9-]+)(?:\.mdc)?\s*$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Check if we're in the "Related Rules" section
		if strings.Contains(strings.ToLower(line), "related rules") {
			inRelatedRules = true
			continue
		}

		// Stop if we hit another section (##) after entering Related Rules section
		if inRelatedRules && strings.HasPrefix(line, "##") {
			break
		}

		// Extract rule names from list items
		if inRelatedRules {
			matches := rulePattern.FindStringSubmatch(line)
			if len(matches) > 1 {
				rules = append(rules, matches[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rules, nil
}

// installRulesForRole installs rules for a specific role
func installRulesForRole(sourceDir, targetDir, role string, ruleNames []string) (int, error) {
	// Create target directory
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return 0, fmt.Errorf("error creating target directory: %v", err)
	}

	count := 0

	// Build a map of rule names for quick lookup
	ruleMap := make(map[string]bool)
	for _, ruleName := range ruleNames {
		ruleMap[ruleName] = true
	}

	// Walk through embedded rules directory
	err := fs.WalkDir(embeddedFS, sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Only process .mdc files
		ext := filepath.Ext(path)
		if ext != ".mdc" {
			return nil
		}

		// Extract rule name from filename (without extension)
		ruleName := strings.TrimSuffix(filepath.Base(path), ext)

		// Check if this rule is in the role's rule list
		if !ruleMap[ruleName] {
			return nil
		}

		// Target file path
		targetPath := filepath.Join(targetDir, filepath.Base(path))

		// Copy file with conflict handling
		err = copyFileWithConflictHandling(path, targetPath)
		if err == ErrSkipped {
			// File was skipped (identical or user chose to skip)
			return nil
		}
		if err != nil {
			return fmt.Errorf("error copying file %s: %v", path, err)
		}

		fmt.Printf("  ✓ Installed rule for %s: %s\n", role, ruleName)
		count++

		return nil
	})

	return count, err
}

// copyFile copies a file from embedded filesystem to destination path
func copyFile(src, dst string) error {
	sourceFile, err := embeddedFS.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Set default file permissions (0644 for regular files)
	return os.Chmod(dst, 0644)
}

// parseInstallArgs parses install command arguments and returns the target base directory
// Priority: --dir > --home > default (current directory)
func parseInstallArgs(args []string) (string, error) {
	var targetDir string
	var useHome bool

	// Check for help flag
	for _, arg := range args {
		if arg == "-h" || arg == "--help" || arg == "help" {
			printInstallHelp()
			os.Exit(0)
		}
	}

	// Parse arguments
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "--home" || arg == "-H" {
			useHome = true
		} else if arg == "--dir" || arg == "-d" {
			if i+1 >= len(args) {
				return "", fmt.Errorf("--dir requires a directory path")
			}
			targetDir = args[i+1]
			i++ // Skip the next argument as it's the directory path
		}
	}

	// Priority: --dir > --home > default
	if targetDir != "" {
		// Validate directory exists or can be created
		absPath, err := filepath.Abs(targetDir)
		if err != nil {
			return "", fmt.Errorf("invalid directory path: %v", err)
		}
		return absPath, nil
	}

	if useHome {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting home directory: %v", err)
		}
		return homeDir, nil
	}

	// Default: current directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current directory: %v", err)
	}
	return cwd, nil
}

// ErrSkipped is returned when a file copy is skipped
var ErrSkipped = fmt.Errorf("file skipped")

// copyFileWithConflictHandling copies a file with conflict detection and handling
// Returns ErrSkipped if the file was skipped, nil if copied successfully, or an error
func copyFileWithConflictHandling(src, dst string) error {
	// Check if destination file exists
	if _, err := os.Stat(dst); err == nil {
		// File exists, check if it's different
		if filesAreDifferent(src, dst) {
			return handleFileConflict(src, dst)
		}
		// Files are identical, skip silently
		return ErrSkipped
	}

	// File doesn't exist, proceed with normal copy
	return copyFile(src, dst)
}

// filesAreDifferent checks if two files have different content
// file1 is from embedded filesystem, file2 is from filesystem
func filesAreDifferent(file1, file2 string) bool {
	hash1, err := fileHashEmbedded(file1)
	if err != nil {
		return true // Assume different if we can't read
	}

	hash2, err := fileHashFilesystem(file2)
	if err != nil {
		return true // Assume different if we can't read
	}

	return !bytes.Equal(hash1, hash2)
}

// fileHashEmbedded calculates SHA256 hash of a file from embedded filesystem
func fileHashEmbedded(filePath string) ([]byte, error) {
	file, err := embeddedFS.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

// fileHashFilesystem calculates SHA256 hash of a file from filesystem
func fileHashFilesystem(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

// handleFileConflict handles file conflicts by prompting user for action
func handleFileConflict(src, dst string) error {
	fmt.Printf("\n⚠️  File conflict detected: %s\n", dst)
	fmt.Println("  Source:", src)
	fmt.Println("  Target:", dst)
	fmt.Println()
	fmt.Println("Choose an action:")
	fmt.Println("  [s]kip - Skip this file (keep existing)")
	fmt.Println("  [o]verwrite - Overwrite with source file")
	fmt.Println("  [d]iff - View differences")
	fmt.Println("  [m]erge - Merge files (manual edit required)")
	fmt.Print("Your choice (s/o/d/m): ")

	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	choice = strings.TrimSpace(strings.ToLower(choice))

	switch choice {
	case "s", "skip":
		fmt.Println("  → Skipped")
		return ErrSkipped

	case "o", "overwrite":
		if err := copyFile(src, dst); err != nil {
			return err
		}
		fmt.Println("  → Overwritten")
		return nil

	case "d", "diff":
		if err := showDiff(src, dst); err != nil {
			fmt.Printf("  ⚠️  Error showing diff: %v\n", err)
			// Fall through to ask again
			return handleFileConflict(src, dst)
		}
		// After showing diff, ask again
		return handleFileConflict(src, dst)

	case "m", "merge":
		if err := openForMerge(src, dst); err != nil {
			fmt.Printf("  ⚠️  Error opening for merge: %v\n", err)
			return err
		}
		fmt.Println("  → Opened for manual merge")
		return nil

	default:
		fmt.Println("  ⚠️  Invalid choice, please try again")
		return handleFileConflict(src, dst)
	}
}

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
	colorBold   = "\033[1m"
)

// showDiff shows the difference between two files with enhanced formatting
// file1 is from embedded filesystem, file2 is from filesystem
func showDiff(file1, file2 string) error {
	// Create a temporary file for the embedded file
	tmpFile, err := os.CreateTemp("", "mozi-embed-*.tmp")
	if err != nil {
		return fmt.Errorf("error creating temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Copy embedded file to temp file
	embeddedFile, err := embeddedFS.Open(file1)
	if err != nil {
		return fmt.Errorf("error opening embedded file: %v", err)
	}
	defer embeddedFile.Close()

	if _, err := io.Copy(tmpFile, embeddedFile); err != nil {
		return fmt.Errorf("error copying embedded file: %v", err)
	}
	tmpFile.Close()

	// Try git diff first (better color support), fallback to diff
	var cmd *exec.Cmd
	var useGitDiff bool
	if _, err := exec.LookPath("git"); err == nil {
		cmd = exec.Command("git", "diff", "--no-index", "--color=always", file2, tmpFile.Name())
		useGitDiff = true
	} else {
		cmd = exec.Command("diff", "-u", file2, tmpFile.Name())
		useGitDiff = false
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		// diff/git-diff returns non-zero exit code if files differ, which is expected
		if cmd.ProcessState.ExitCode() == 1 {
			fmt.Println()
			fmt.Printf("%s%s═══════════════════════════════════════════════════════════════%s\n", colorBold, colorCyan, colorReset)
			fmt.Printf("%s%s  File Differences%s\n", colorBold, colorCyan, colorReset)
			fmt.Printf("%s%s═══════════════════════════════════════════════════════════════%s\n", colorBold, colorCyan, colorReset)
			fmt.Printf("%sExisting file:%s %s\n", colorYellow, colorReset, file2)
			fmt.Printf("%sNew file:     %s %s\n", colorGreen, colorReset, file1)
			fmt.Printf("%s%s───────────────────────────────────────────────────────────────%s\n", colorGray, colorReset, colorReset)
			fmt.Println()

			if useGitDiff {
				// git diff already has colors, just print it
				fmt.Print(string(output))
			} else {
				// Format unified diff with colors
				printColoredDiff(string(output))
			}

			fmt.Printf("%s%s───────────────────────────────────────────────────────────────%s\n", colorGray, colorReset, colorReset)
			fmt.Println()
			return nil
		}
		return err
	}

	// Files are identical
	fmt.Printf("%s%sFiles are identical%s\n", colorGreen, colorBold, colorReset)
	return nil
}

// printColoredDiff prints a unified diff with color highlighting
func printColoredDiff(diffOutput string) {
	lines := strings.Split(diffOutput, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "---") {
			fmt.Printf("%s%s%s\n", colorRed, line, colorReset)
		} else if strings.HasPrefix(line, "+++") {
			fmt.Printf("%s%s%s\n", colorGreen, line, colorReset)
		} else if strings.HasPrefix(line, "@@") {
			fmt.Printf("%s%s%s\n", colorCyan, line, colorReset)
		} else if strings.HasPrefix(line, "-") {
			fmt.Printf("%s%s%s\n", colorRed, line, colorReset)
		} else if strings.HasPrefix(line, "+") {
			fmt.Printf("%s%s%s\n", colorGreen, line, colorReset)
		} else {
			fmt.Println(line)
		}
	}
}

// openForMerge opens both files for manual merging
func openForMerge(src, dst string) error {
	// Try to open files with default editor
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim" // Default to vim if EDITOR is not set
	}

	fmt.Printf("  Opening files for merge with %s...\n", editor)
	fmt.Printf("  Source: %s\n", src)
	fmt.Printf("  Target: %s\n", dst)

	// Open target file (user should manually merge)
	cmd := exec.Command(editor, dst)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error opening editor: %v", err)
	}

	return nil
}

// printHelp prints the main help message
func printHelp() {
	fmt.Println("Mozi - Cursor Rules & Commands Installer")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  mozi <COMMAND> [OPTIONS]")
	fmt.Println()
	fmt.Println("COMMANDS:")
	fmt.Println("  install    Install commands and rules to Cursor IDE configuration")
	fmt.Println("  help       Show this help message")
	fmt.Println()
	fmt.Println("OPTIONS:")
	fmt.Println("  -h, --help    Show help message")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  mozi install              # Install to ./.cursor (default)")
	fmt.Println("  mozi install --home        # Install to ~/.cursor")
	fmt.Println("  mozi install --dir <path>  # Install to specified directory")
	fmt.Println("  mozi help                 # Show help message")
	fmt.Println()
	fmt.Println("For more information about a command, run:")
	fmt.Println("  mozi <COMMAND> --help")
	fmt.Println()
}

// printInstallHelp prints the install command help message
func printInstallHelp() {
	fmt.Println("mozi install - Install commands and rules to Cursor IDE configuration")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  mozi install [OPTIONS]")
	fmt.Println()
	fmt.Println("DESCRIPTION:")
	fmt.Println("  Installs commands and rules from the current directory to Cursor IDE configuration.")
	fmt.Println("  The tool will:")
	fmt.Println("    - Copy all .md files from commands/ to <target>/.cursor/commands/")
	fmt.Println("    - Copy .mdc files from rules/ to <target>/.cursor/rules/mozi/<role>/")
	fmt.Println("    - Automatically bind rules to roles based on 'Related Rules' section in command files")
	fmt.Println("    - Handle file conflicts with options: skip, overwrite, diff, or merge")
	fmt.Println()
	fmt.Println("TARGET DIRECTORY (priority order):")
	fmt.Println("  1. --dir, -d <path>    Use specified directory (highest priority)")
	fmt.Println("  2. --home, -H          Use home directory (~/.cursor)")
	fmt.Println("  3. (default)          Use current directory (./.cursor)")
	fmt.Println()
	fmt.Println("OPTIONS:")
	fmt.Println("  -h, --help              Show this help message")
	fmt.Println("  -d, --dir <path>        Install to specified directory")
	fmt.Println("  -H, --home              Install to home directory (~/.cursor)")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  mozi install                    # Install to ./.cursor (default)")
	fmt.Println("  mozi install --home               # Install to ~/.cursor")
	fmt.Println("  mozi install --dir /path/to/config  # Install to /path/to/config/.cursor")
	fmt.Println("  mozi install -d ./my-config      # Install to ./my-config/.cursor")
	fmt.Println()
}
