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

| 功能 | 命令示例 | 实现位置 |
|------|----------|----------|
| 设置用户变量 | `menv KEY VALUE` | `env.Set()` |
| 设置系统变量 | `menv -sys KEY VALUE` | `env.SetSystem()` |
| 删除用户变量 | `menv -d KEY` | `env.Unset()` |
| 删除系统变量 | `menv -d -sys KEY` | `env.UnsetSystem()` |
| 添加到 PATH | `menv -add "C:\bin"` | `path.Add()` |
| 从文件导入 | `menv -file env.sh -startWith export` | `main.processEnvFile()` |

## 命令行标志 (cmd/flags.go)

- `-file <path>` - 环境文件路径
- `-d` - 删除模式
- `-sys` - 操作系统级变量
- `-startWith <str>` - 文件行前缀过滤
- `-add <path>` - 添加到 PATH

## 技术要点

1. **setx vs PowerShell**: PATH 使用 PowerShell 设置（绕过 setx 1024 字符限制）
2. **注册表路径**:
   - 用户: `HKEY_CURRENT_USER\Environment`
   - 系统: `HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Environment`
3. **系统变量操作需要管理员权限**

- 
