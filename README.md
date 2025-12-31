# Mozi

![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)
![Version](https://img.shields.io/badge/version-alpha-orange?style=flat-square)
![Go Version](https://img.shields.io/github/go-mod/go-version/Edison-A-N/mozi?style=flat-square)

> A CLI tool for managing and sharing Cursor IDE Rules and Commands

Mozi helps you organize and install Cursor IDE rules and commands by role. Define capabilities for specific roles (backend engineers, frontend engineers, etc.) and automatically bind corresponding rules to these roles. It also provides workflow commands for requirements analysis, task breakdown, and implementation execution.

## Features

- 📝 Manage Cursor IDE rules and commands
- 🎯 Organize commands by role
- 🔗 Automatically bind rules to roles
- 🔍 Interactive conflict resolution

## Installation

### Quick Install

```bash
curl -fsSL https://raw.githubusercontent.com/Edison-A-N/mozi/main/install.sh | bash
```

### Build from Source

```bash
git clone https://github.com/Edison-A-N/mozi.git
cd mozi
go build -o mozi main.go
sudo mv mozi /usr/local/bin/
```

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

The tool will copy commands and rules to your Cursor configuration directory and handle file conflicts interactively.

## Included Commands

- **requirements-to-implementation**: Complete workflow for transforming requirements into working code
- **ask**: Ask questions only - no code changes (web search enabled)
- **git-help/git-commit**: Help with Git commit message generation

## License

MIT
