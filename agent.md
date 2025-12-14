# menv - Windows 环境变量管理工具

## 项目结构

```
menv/
├── main.go          # 入口，命令行解析与分发
├── cmd/flags.go     # 命令行标志定义
├── env/
│   ├── user.go      # 用户环境变量 Set/Unset/SetPS
│   ├── system.go    # 系统环境变量 SetSystem/UnsetSystem
│   └── parser.go    # 环境文件解析 ParseEnvFile
├── path/
│   ├── query.go     # 注册表查询 PATH (QueryUserPath/QuerySystemPath)
│   └── modify.go    # PATH 操作 (Add/CleanUser)
└── color/color.go   # ANSI 彩色输出 (Success/Error/Warning/Info)
```

## 核心功能

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
  -d                Delete environment variable
  -sys              Target system env (default: user)
  -file <path>      Read env vars from file
  -startWith <str>  Filter lines starting with string

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
  menv -file env.sh -startWith export



现在我已经了解了完整的代码结构。这是一个Windows环境变量管理工具，当前支持：

**已有功能：**

- 设置/删除用户环境变量
- 设置/删除系统环境变量
- 从文件批量导入环境变量
- 添加路径到 PATH
- PATH 去重清理（`CleanUser` 已实现但未暴露）

---

## 建议增加的实用功能

### 1. **查看环境变量 (`-list`, `-get`) 已实现**
```bash
menv -list              # 列出所有用户环境变量
menv -list -sys         # 列出所有系统环境变量
menv -get JAVA_HOME     # 查看指定变量的值
menv -path              # 美化显示 PATH（每行一个路径）
```

### 2. **PATH 管理增强 已实现**
```bash
menv -rm "C:\path"      # 从 PATH 中移除指定路径
menv -clean             # 清理 PATH（去重 + 移除不存在的目录）
menv -clean -sys        # 清理系统 PATH
```

### 3. **导出环境变量 (`-export`)**
```bash
menv -export env.sh             # 导出为 shell 格式
menv -export env.bat            # 导出为 bat 格式
menv -export env.json           # 导出为 JSON 格式
```

### 4. **备份与恢复**
```bash
menv -backup backup.json        # 备份当前环境变量
menv -restore backup.json       # 从备份恢复
```

### 5. **搜索功能**
```bash
menv -search java       # 搜索包含 "java" 的变量名或值
menv -search java -path # 在 PATH 中搜索包含 "java" 的路径
```

### 6. **PATH 检查**
```bash
menv -check             # 检查 PATH 中不存在的目录
menv -check -fix        # 自动移除不存在的目录
```

### 7. **交互式确认 (`-i`)**
```bash
menv -d JAVA_HOME -i    # 删除前确认
```

---

