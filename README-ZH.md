<p align="center">
  <img src="https://raw.githubusercontent.com/egonelbre/gophers/master/vector/superhero/zorro.svg" width="120" alt="menv logo">
</p>

<h1 align="center">✨ menv ✨</h1>

<p align="center">
  <strong>🪟 可爱又强大的 Windows 环境变量管理工具</strong>
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

## 🚀 特性

- 📋 **列出与搜索** - 轻松查看和搜索环境变量
- ✏️ **设置与删除** - 从命令行管理环境变量
- 📁 **PATH 管理** - 添加、删除和清理 PATH 条目
- 🔍 **健康检查** - 自动查找并修复无效路径
- 💾 **备份与恢复** - 永不丢失你的环境设置
- 🛡️ **系统与用户** - 支持用户和系统环境变量

## 📦 安装

```bash
go install github.com/doraemonkeys/menv@latest
```

或者克隆并本地安装：

```bash
git clone https://github.com/doraemonkeys/menv.git
cd menv
go install
```

## 🎯 快速开始

### 👀 查看环境变量

```bash
menv -list                  # 列出所有用户环境变量
menv -get JAVA_HOME         # 获取特定变量的值
menv -search java           # 搜索包含 "java" 的变量
```

### ✏️ 设置/删除环境变量

```bash
menv GOPATH C:\Go           # 设置用户环境变量
menv -d GOPATH              # 删除用户环境变量
```

### 📁 PATH 管理

```bash
menv -path                  # 查看 PATH（每行一个路径）
menv -add "C:\bin"          # 向 PATH 添加路径
menv -rm "C:\bin"           # 从 PATH 移除路径
menv -clean                 # 清理 PATH（去重 + 移除无效路径）
menv -search java -path     # 在 PATH 中搜索包含 "java" 的路径
```

### 🔍 检查无效路径

```bash
menv -check                 # 检查 PATH 中的无效目录
menv -check -fix            # 自动移除无效路径
menv -check -fix -i         # 移除每个路径前确认
```

### 💾 备份与恢复

```bash
menv -backup backup.json    # 备份环境变量
menv -restore backup.json   # 恢复环境变量
menv -export env.sh         # 导出为 shell 脚本
```

### 🛡️ 系统环境变量

以上命令默认操作**用户**环境变量。添加 `-sys` 参数可操作**系统**环境变量（需要管理员权限）：

```bash
menv -list -sys             # 列出系统环境变量
menv -path -sys             # 查看系统 PATH
menv -add "C:\bin" -sys     # 添加到系统 PATH
menv -clean -sys            # 清理系统 PATH
```

## 📖 常用场景

| 场景 | 命令 |
|:-----|:-----|
| 👁️ 查看特定变量 | `menv -get JAVA_HOME` |
| ➕ 向 PATH 添加目录 | `menv -add "C:\tools\bin"` |
| 🧹 清理重复和无效路径 | `menv -clean` |
| 💾 备份当前环境变量 | `menv -backup my-env.json` |
| 🔎 查找 Java 相关路径 | `menv -search java -path` |

## 📄 许可证

[LICENSE](LICENSE)
