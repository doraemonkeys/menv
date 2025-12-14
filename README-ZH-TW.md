# menv

[English](README.md) | [简体中文](README-ZH.md) | [繁體中文](README-ZH-TW.md) | [한국어](README-KO.md) | [日本語](README-JA.md)

Windows 環境變數管理工具，讓你在命令列輕鬆管理環境變數和 PATH。

## 安裝

```bash
go install github.com/doraemonkeys/menv@latest
```

或者克隆後本地安裝：

```bash
git clone https://github.com/doraemonkeys/menv.git
cd menv
go install
```

## 快速上手

### 查看環境變數

```bash
menv -list                  # 列出所有使用者環境變數
menv -get JAVA_HOME         # 獲取指定變數的值
menv -search java           # 搜尋包含 java 的變數
```

### 設定/刪除環境變數

```bash
menv GOPATH C:\Go           # 設定使用者環境變數
menv -d GOPATH              # 刪除使用者環境變數
```

### PATH 管理

```bash
menv -path                  # 查看 PATH（每行顯示一個路徑）
menv -add "C:\bin"          # 添加路徑到 PATH
menv -rm "C:\bin"           # 從 PATH 移除路徑
menv -clean                 # 清理 PATH（去重 + 刪除無效路徑）
menv -search java -path     # 在 PATH 中搜尋包含 java 的路徑
```

### 檢查無效路徑

```bash
menv -check                 # 檢查 PATH 中的無效目錄
menv -check -fix            # 自動移除無效路徑
menv -check -fix -i         # 移除前逐個確認
```

### 備份與恢復

```bash
menv -backup backup.json    # 備份環境變數
menv -restore backup.json   # 恢復環境變數
menv -export env.sh         # 導出為 shell 腳本
```

### 作業系統環境變數

以上命令預設操作**使用者**環境變數，加 `-sys` 可操作**系統**環境變數（需管理員權限）：

```bash
menv -list -sys             # 列出系統環境變數
menv -path -sys             # 查看系統 PATH
menv -add "C:\bin" -sys     # 添加到系統 PATH
menv -clean -sys            # 清理系統 PATH
```

## 常用場景

| 場景 | 命令 |
|------|------|
| 查看某個變數 | `menv -get JAVA_HOME` |
| 添加目錄到 PATH | `menv -add "C:\tools\bin"` |
| 清理重複和無效路徑 | `menv -clean` |
| 備份當前環境變數 | `menv -backup my-env.json` |
| 查找 Java 相關路徑 | `menv -search java -path` |
