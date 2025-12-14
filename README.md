# menv

[English](README.md) | [简体中文](README-ZH.md) | [繁體中文](README-ZH-TW.md) | [한국어](README-KO.md) | [日本語](README-JA.md)

A Windows environment variable management tool that makes it easy to manage environment variables and PATH from the command line.

## Installation

```bash
go install github.com/doraemonkeys/menv@latest
```

Or clone and install locally:

```bash
git clone https://github.com/doraemonkeys/menv.git
cd menv
go install
```

## Quick Start

### View Environment Variables

```bash
menv -list                  # List all user environment variables
menv -get JAVA_HOME         # Get the value of a specific variable
menv -search java           # Search for variables containing "java"
```

### Set/Delete Environment Variables

```bash
menv GOPATH C:\Go           # Set a user environment variable
menv -d GOPATH              # Delete a user environment variable
```

### PATH Management

```bash
menv -path                  # View PATH (one path per line)
menv -add "C:\bin"          # Add a path to PATH
menv -rm "C:\bin"           # Remove a path from PATH
menv -clean                 # Clean PATH (deduplicate + remove invalid paths)
menv -search java -path     # Search for paths containing "java" in PATH
```

### Check Invalid Paths

```bash
menv -check                 # Check for invalid directories in PATH
menv -check -fix            # Automatically remove invalid paths
menv -check -fix -i         # Confirm before removing each path
```

### Backup and Restore

```bash
menv -backup backup.json    # Backup environment variables
menv -restore backup.json   # Restore environment variables
menv -export env.sh         # Export as shell script
```

### System Environment Variables

The above commands operate on **user** environment variables by default. Add `-sys` to operate on **system** environment variables (requires administrator privileges):

```bash
menv -list -sys             # List system environment variables
menv -path -sys             # View system PATH
menv -add "C:\bin" -sys     # Add to system PATH
menv -clean -sys            # Clean system PATH
```

## Common Scenarios

| Scenario | Command |
|----------|---------|
| View a specific variable | `menv -get JAVA_HOME` |
| Add a directory to PATH | `menv -add "C:\tools\bin"` |
| Clean duplicate and invalid paths | `menv -clean` |
| Backup current environment variables | `menv -backup my-env.json` |
| Find Java-related paths | `menv -search java -path` |
