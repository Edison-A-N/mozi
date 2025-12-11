# Mozi

![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)
![Go Version](https://img.shields.io/github/go-mod/go-version/Edison-A-N/mozi?style=flat-square)

> A CLI tool for managing and sharing Cursor IDE Rules and Commands

Mozi helps you organize and install Cursor IDE rules and commands by role. Define capabilities for specific roles (backend engineers, frontend engineers, etc.) and automatically bind corresponding rules to these roles. It also provides workflow commands for requirements analysis, task breakdown, and implementation execution.

## Features

- 📝 **Rules Management**: Manage rule files in `.mdc` format
- 🎯 **Commands Management**: Define role capabilities through commands
- 🔗 **Role Binding**: Automatically bind rules to roles
- 📦 **CLI Tool**: One-click installation to Cursor configuration
- 🔄 **Version Control**: Manage changes through Git
- 🛠️ **Workflow Commands**: Pre-built commands for requirements-to-implementation workflow
- 🔍 **Smart Conflict Handling**: Interactive conflict resolution (skip, overwrite, diff, merge)

## Installation

### Build from Source

```bash
git clone https://github.com/Edison-A-N/mozi.git
cd mozi
go build -o mozi main.go
sudo mv mozi /usr/local/bin/  # Optional: install to system path
```

### Using Go Install

```bash
go install github.com/Edison-A-N/mozi@latest
```

Make sure `$GOPATH/bin` or `$GOBIN` is in your PATH.

## Usage

### Install Commands and Rules

```bash
# Install to current directory (default: ./.cursor)
mozi install

# Install to home directory (~/.cursor)
mozi install --home

# Install to specified directory
mozi install --dir /path/to/config
```

The tool will:
- Copy `.md` files from `commands/` to `<target>/.cursor/commands/mozi/`
- Copy `.mdc` files from `rules/` to `<target>/.cursor/rules/mozi/<role>/` (if rules exist)
- Automatically bind rules to roles based on "Related Rules" section in command files (if rules exist)
- Handle file conflicts with options: skip, overwrite, diff, or merge

### Installation Structure

```
<target>/.cursor/
├── commands/
│   └── mozi/
│       ├── ask.md
│       ├── git-help/
│       │   └── git-commit.md
│       └── requirements-to-implementation/
│           ├── 01-requirements-analysis.md
│           ├── 02-clarify-questions.md
│           ├── 03-breakdown-tasks.md
│           └── 04-execute-todo.md
```

## File Format

### Rules (`.mdc`)

```markdown
---
description: Rule description
globs: ["**/*.ts", "**/*.tsx"]
alwaysApply: false
---

# Rule Title

Rule content...
```

### Commands (`.md`)

```markdown
---
description: Command description
---

# Command Title

Command content...

## Related Rules
This role will apply the following rules:
- rule-name-1
- rule-name-2
```

## Included Commands

### Workflow Commands

#### Requirements-to-Implementation Workflow
A complete workflow for transforming requirements into working code:

1. **01-requirements-analysis**: Break down requirements and design solutions
   - Analyze background information
   - Identify core requirements
   - Design technical solutions
   - Define goals and success criteria

2. **02-clarify-questions**: Identify information that needs clarification
   - Goal-oriented question identification
   - Priority and impact assessment
   - Blocking vs non-blocking questions

3. **03-breakdown-tasks**: Break down work into executable TODO items
   - Implementation-oriented task breakdown
   - Dependency management
   - Clear acceptance criteria

4. **04-execute-todo**: Execute TODO items step by step
   - One task at a time execution
   - Progress tracking
   - Completion verification

#### Utility Commands

- **ask**: Ask questions only - no code changes (web search enabled)
- **git-help/git-commit**: Help with Git commit message generation

## Project Status

This project is actively maintained and ready for public use. We welcome contributions and feedback!

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-rule`)
3. Commit your changes (`git commit -m 'Add some amazing rule'`)
4. Push to the branch (`git push origin feature/amazing-rule`)
5. Open a Pull Request

### Adding New Commands

To add a new command:
1. Create a `.md` file in the `commands/` directory
2. Add frontmatter with `description` field
3. Optionally add a "Related Rules" section to bind rules to the command

### Adding New Rules

To add a new rule:
1. Create a `.mdc` file in the `rules/` directory
2. Add frontmatter with `description`, `globs`, and `alwaysApply` fields
3. Reference the rule in command files' "Related Rules" section

## Related Links

- [Cursor Official Documentation](https://docs.cursor.com)
- [Cursor Rules Documentation](https://docs.cursor.com/zh/context/rules)

## License

MIT
