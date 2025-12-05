# Mozi

![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)

> A CLI tool for managing and sharing Cursor IDE Rules and Commands

Mozi helps you organize and install Cursor IDE rules and commands by role. Define capabilities for specific roles (backend engineers, frontend engineers, etc.) and automatically bind corresponding rules to these roles.

## Features

- 📝 **Rules Management**: Manage rule files in `.mdc` format
- 🎯 **Commands Management**: Define role capabilities through commands
- 🔗 **Role Binding**: Automatically bind rules to roles
- 📦 **CLI Tool**: One-click installation to Cursor configuration
- 🔄 **Version Control**: Manage changes through Git

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
- Copy `.mdc` files from `rules/` to `<target>/.cursor/rules/mozi/<role>/`
- Automatically bind rules to roles based on "Related Rules" section in command files
- Handle file conflicts with options: skip, overwrite, diff, or merge

### Installation Structure

```
<target>/.cursor/
├── commands/
│   └── mozi/
│       └── backend-engineer.md
└── rules/
    └── mozi/
        └── backend-engineer/
            ├── backend-coding-standards.mdc
            ├── api-design-guidelines.mdc
            ├── database-best-practices.mdc
            ├── testing-requirements.mdc
            └── security-guidelines.mdc
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

## Supported Roles

### Backend Engineer

Includes rules for:
- Backend coding standards
- API design guidelines
- Database best practices
- Testing requirements
- Security guidelines

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-rule`)
3. Commit your changes (`git commit -m 'Add some amazing rule'`)
4. Push to the branch (`git push origin feature/amazing-rule`)
5. Open a Pull Request

## Related Links

- [Cursor Official Documentation](https://docs.cursor.com)
- [Cursor Rules Documentation](https://docs.cursor.com/zh/context/rules)

## License

MIT
