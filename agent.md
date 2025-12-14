# menv - Windows 环境变量管理工具

## 项目结构

```
menv/
├── main.go              # 入口，命令行解析与分发
├── Makefile             # 构建与测试命令
├── cmd/
│   └── flags.go         # 命令行标志定义
├── env/
│   ├── user.go          # 用户环境变量 Set/Unset/SetPS
│   ├── system.go        # 系统环境变量 SetSystem/UnsetSystem
│   ├── parser.go        # 环境文件解析 ParseEnvFile
│   ├── query.go         # 环境变量查询 (List/Get)
│   └── export.go        # 环境变量导出 (ExportToFile)
├── path/
│   ├── query.go         # 注册表查询 PATH (QueryUserPath/QuerySystemPath)
│   └── modify.go        # PATH 操作 (Add/Remove/Clean)
├── color/
│   └── color.go         # ANSI 彩色输出 (Success/Error/Warning/Info)
├── scripts/
│   └── check-file-count.sh  # CI 文件数量检查脚本
└── .github/
    └── workflows/
        └── ci.yml       # GitHub Actions CI 配置
```

这是一个Windows环境变量管理工具。

## 核心功能

PS E:\Doraemon\IT\Repository\menv> .\menv.exe
menv - Windows Environment Variable Manager

Usage:
  menv [options] [key] [value]

Options:
  -list             List all env vars
  -get <key>        Get env var value
  -path             Display PATH (one per line)
  -add <path>       Add path to PATH variable
  -rm <path>        Remove path from PATH variable
  -clean            Clean PATH (dedupe + remove invalid)
  -check            Check PATH for invalid directories
  -fix              Auto-remove invalid paths (use with -check)
  -i                Interactive confirmation
  -d                Delete environment variable
  -sys              Target system env (default: user)
  -file <path>      Read env vars from file
  -startWith <str>  Filter lines starting with string
  -export <path>    Export env vars to file (sh/bat/json)
  -backup <path>    Backup env vars to JSON file
  -restore <path>   Restore env vars from backup file
  -search <keyword> Search env vars by keyword
                    Use with -path to search in PATH

Examples:
  menv -list                         # List user env vars
  menv -list -sys                    # List system env vars
  menv -get JAVA_HOME                # Get JAVA_HOME value
  menv -path                         # Display user PATH
  menv -path -sys                    # Display system PATH
  menv GOPATH C:\Go                  # Set user env var
  menv -sys GOPATH C:\Go             # Set system env var
  menv -d GOPATH                     # Delete user env var
  menv -d -sys GOPATH                # Delete system env var
  menv -add "C:\bin"                 # Add to user PATH
  menv -add "C:\bin" -sys            # Add to system PATH
  menv -rm "C:\bin"                  # Remove from user PATH
  menv -rm "C:\bin" -sys             # Remove from system PATH
  menv -clean                        # Clean user PATH
  menv -clean -sys                   # Clean system PATH
  menv -clean -i                     # Clean with confirmation
  menv -file env.sh -startWith export
  menv -export env.sh                # Export user env as shell
  menv -export env.bat               # Export user env as batch
  menv -export env.json              # Export user env as JSON
  menv -export env.json -sys         # Export system env as JSON
  menv -backup backup.json           # Backup user env vars
  menv -backup backup.json -sys      # Backup system env vars
  menv -restore backup.json          # Restore user env vars
  menv -restore backup.json -sys     # Restore system env vars
  menv -search java                  # Search env vars for 'java'
  menv -search java -path            # Search PATH for 'java'
  menv -check                        # Check user PATH for invalid dirs
  menv -check -sys                   # Check system PATH for invalid dirs
  menv -check -fix                   # Check and remove invalid paths
  menv -check -fix -i                # Check and remove with confirmation

## CI

```bash
make test check-coverage lint check-file-count
```



