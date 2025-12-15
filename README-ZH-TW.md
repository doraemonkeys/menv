<p align="center">
  <img src="./assets/logo2.png" width="120" alt="menv logo">
</p>

<h1 align="center">✨ menv ✨</h1>

<p align="center">
  <strong>🪟 可愛又強大的 Windows 環境變數管理工具</strong>
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

- 📋 **列出與搜尋** - 輕鬆查看和搜尋環境變數
- ✏️ **設定與刪除** - 從命令列管理環境變數
- 📁 **PATH 管理** - 新增、刪除和清理 PATH 條目
- 🔍 **健康檢查** - 自動查找並修復無效路徑
- 💾 **備份與還原** - 永不遺失你的環境設定
- 🛡️ **系統與使用者** - 支援使用者和系統環境變數

## 📦 安裝

從 [GitHub Releases](https://github.com/doraemonkeys/menv/releases) 下載最新版本，並將其新增到 PATH。

或者透過 Go 安裝：

```bash
go install github.com/doraemonkeys/menv@latest
```

或者克隆並本地安裝：

```bash
git clone https://github.com/doraemonkeys/menv.git
cd menv
go install
```

## 🎯 快速開始

### 👀 查看環境變數

```bash
menv -list                  # 列出所有使用者環境變數
menv -get JAVA_HOME         # 取得特定變數的值
menv -search java           # 搜尋包含 "java" 的變數
```

### ✏️ 設定/刪除環境變數

```bash
menv GOPATH C:\Go           # 設定使用者環境變數
menv -d GOPATH              # 刪除使用者環境變數
```

### 📁 PATH 管理

```bash
menv -path                  # 查看 PATH（每行一個路徑）
menv -add "C:\bin"          # 向 PATH 新增路徑
menv -rm "C:\bin"           # 從 PATH 移除路徑
menv -clean                 # 清理 PATH（去重 + 移除無效路徑）
menv -search java -path     # 在 PATH 中搜尋包含 "java" 的路徑
```

### 🔍 檢查無效路徑

```bash
menv -check                 # 檢查 PATH 中的無效目錄
menv -check -fix            # 自動移除無效路徑
menv -check -fix -y         # 無需確認直接移除
```

### 💾 備份與還原

```bash
menv -backup backup.json    # 備份環境變數
menv -restore backup.json   # 還原環境變數
menv -export env.sh         # 匯出為 shell 腳本
```

### 🛡️ 系統環境變數

以上命令預設操作**使用者**環境變數。新增 `-sys` 參數可操作**系統**環境變數（需要管理員權限）：

```bash
menv -list -sys             # 列出系統環境變數
menv -path -sys             # 查看系統 PATH
menv -add "C:\bin" -sys     # 新增到系統 PATH
menv -clean -sys            # 清理系統 PATH
```

## 📖 常用場景

| 場景 | 命令 |
|:-----|:-----|
| 👁️ 查看特定變數 | `menv -get JAVA_HOME` |
| ➕ 向 PATH 新增目錄 | `menv -add "C:\tools\bin"` |
| 🧹 清理重複和無效路徑 | `menv -clean` |
| 💾 備份當前環境變數 | `menv -backup my-env.json` |
| 🔎 查找 Java 相關路徑 | `menv -search java -path` |

## 📄 授權條款

[LICENSE](LICENSE)
