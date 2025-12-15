<p align="center">
  <img src="./assets/logo2.png" width="120" alt="menv logo">
</p>

<h1 align="center">✨ menv ✨</h1>

<p align="center">
  <strong>🪟 A cute & powerful Windows environment variable manager</strong>
</p>

<p align="center">
  <a href="https://github.com/doraemonkeys/menv/actions"><img src="https://github.com/doraemonkeys/menv/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://goreportcard.com/report/github.com/doraemonkeys/menv"><img src="https://goreportcard.com/badge/github.com/doraemonkeys/menv" alt="Go Report Card"></a>
  <a href="https://github.com/doraemonkeys/menv/blob/main/LICENSE"><img src="https://img.shields.io/github/license/doraemonkeys/menv" alt="License"></a>
  <a href="https://github.com/doraemonkeys/menv/releases"><img src="https://img.shields.io/github/v/release/doraemonkeys/menv?include_prereleases" alt="Release"></a>
</p>

<p align="center">
  <a href="README.md">English</a> •
  <a href="README-ZH.md">简体中文</a> •
  <a href="README-ZH-TW.md">繁體中文</a> •
  <a href="README-KO.md">한국어</a> •
  <a href="README-JA.md">日本語</a>
</p>

---

## 🚀 Features

- 📋 **List & Search** - View and search environment variables with ease
- ✏️ **Set & Delete** - Manage environment variables from the command line
- 📁 **PATH Management** - Add, remove, and clean PATH entries
- 🔍 **Health Check** - Find and fix invalid paths automatically
- 💾 **Backup & Restore** - Never lose your environment settings
- 🛡️ **System & User** - Support for both user and system variables

## 📦 Installation

Download the latest binary from [GitHub Releases](https://github.com/doraemonkeys/menv/releases) and add it to your PATH.

Or install via Go:

```bash
go install github.com/doraemonkeys/menv@latest
```

Or clone and install locally:

```bash
git clone https://github.com/doraemonkeys/menv.git
cd menv
go install
```

## 🎯 Quick Start

### 👀 View Environment Variables

```bash
menv -list                  # List all user environment variables
menv -get JAVA_HOME         # Get the value of a specific variable
menv -search java           # Search for variables containing "java"
```

### ✏️ Set/Delete Environment Variables

```bash
menv GOPATH C:\Go           # Set a user environment variable
menv -d GOPATH              # Delete a user environment variable
```

### 📁 PATH Management

```bash
menv -path                  # View PATH (one path per line)
menv -add "C:\bin"          # Add a path to PATH
menv -rm "C:\bin"           # Remove a path from PATH
menv -clean                 # Clean PATH (deduplicate + remove invalid paths)
menv -search java -path     # Search for paths containing "java" in PATH
```

### 🔍 Check Invalid Paths

```bash
menv -check                 # Check for invalid directories in PATH
menv -check -fix            # Automatically remove invalid paths
menv -check -fix -i         # Confirm before removing each path
```

### 💾 Backup and Restore

```bash
menv -backup backup.json    # Backup environment variables
menv -restore backup.json   # Restore environment variables
menv -export env.sh         # Export as shell script
```

### 🛡️ System Environment Variables

The above commands operate on **user** environment variables by default. Add `-sys` to operate on **system** environment variables (requires administrator privileges):

```bash
menv -list -sys             # List system environment variables
menv -path -sys             # View system PATH
menv -add "C:\bin" -sys     # Add to system PATH
menv -clean -sys            # Clean system PATH
```

## 📖 Common Scenarios

| Scenario | Command |
|:---------|:--------|
| 👁️ View a specific variable | `menv -get JAVA_HOME` |
| ➕ Add a directory to PATH | `menv -add "C:\tools\bin"` |
| 🧹 Clean duplicate and invalid paths | `menv -clean` |
| 💾 Backup current environment variables | `menv -backup my-env.json` |
| 🔎 Find Java-related paths | `menv -search java -path` |

## 📄 License

[LICENSE](LICENSE)
