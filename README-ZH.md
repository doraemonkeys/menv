# menv

[English](README.md) | [简体中文](README-ZH.md) | [繁體中文](README-ZH-TW.md) | [한국어](README-KO.md) | [日本語](README-JA.md)

Windows 环境变量管理工具，让你在命令行轻松管理环境变量和 PATH。

## 安装

```bash
go install github.com/doraemonkeys/menv@latest
```

或者克隆后本地安装：

```bash
git clone https://github.com/doraemonkeys/menv.git
cd menv
go install
```

## 快速上手

### 查看环境变量

```bash
menv -list                  # 列出所有用户环境变量
menv -get JAVA_HOME         # 获取指定变量的值
menv -search java           # 搜索包含 java 的变量
```

### 设置/删除环境变量

```bash
menv GOPATH C:\Go           # 设置用户环境变量
menv -d GOPATH              # 删除用户环境变量
```

### PATH 管理

```bash
menv -path                  # 查看 PATH（每行显示一个路径）
menv -add "C:\bin"          # 添加路径到 PATH
menv -rm "C:\bin"           # 从 PATH 移除路径
menv -clean                 # 清理 PATH（去重 + 删除无效路径）
menv -search java -path     # 在 PATH 中搜索包含 java 的路径
```

### 检查无效路径

```bash
menv -check                 # 检查 PATH 中的无效目录
menv -check -fix            # 自动移除无效路径
menv -check -fix -i         # 移除前逐个确认
```

### 备份与恢复

```bash
menv -backup backup.json    # 备份环境变量
menv -restore backup.json   # 恢复环境变量
menv -export env.sh         # 导出为 shell 脚本
```

### 操作系统环境变量

以上命令默认操作**用户**环境变量，加 `-sys` 可操作**系统**环境变量（需管理员权限）：

```bash
menv -list -sys             # 列出系统环境变量
menv -path -sys             # 查看系统 PATH
menv -add "C:\bin" -sys     # 添加到系统 PATH
menv -clean -sys            # 清理系统 PATH
```

## 常用场景

| 场景 | 命令 |
|------|------|
| 查看某个变量 | `menv -get JAVA_HOME` |
| 添加目录到 PATH | `menv -add "C:\tools\bin"` |
| 清理重复和无效路径 | `menv -clean` |
| 备份当前环境变量 | `menv -backup my-env.json` |
| 查找 Java 相关路径 | `menv -search java -path` |

