<p align="center">
  <img src="https://raw.githubusercontent.com/egonelbre/gophers/master/vector/superhero/zorro.svg" width="120" alt="menv logo">
</p>

<h1 align="center">✨ menv ✨</h1>

<p align="center">
  <strong>🪟 かわいくて強力な Windows 環境変数管理ツール</strong>
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

## 🚀 特徴

- 📋 **一覧と検索** - 環境変数を簡単に表示・検索
- ✏️ **設定と削除** - コマンドラインから環境変数を管理
- 📁 **PATH 管理** - PATH エントリの追加、削除、クリーンアップ
- 🔍 **ヘルスチェック** - 無効なパスを自動的に検出・修正
- 💾 **バックアップと復元** - 環境設定を失うことはありません
- 🛡️ **システムとユーザー** - ユーザーとシステム変数の両方をサポート

## 📦 インストール

[GitHub Releases](https://github.com/doraemonkeys/menv/releases) から最新のバイナリをダウンロードし、PATH に追加してください。

または Go でインストール：

```bash
go install github.com/doraemonkeys/menv@latest
```

またはクローンしてローカルでインストール：

```bash
git clone https://github.com/doraemonkeys/menv.git
cd menv
go install
```

## 🎯 クイックスタート

### 👀 環境変数の表示

```bash
menv -list                  # すべてのユーザー環境変数を一覧表示
menv -get JAVA_HOME         # 特定の変数の値を取得
menv -search java           # "java" を含む変数を検索
```

### ✏️ 環境変数の設定/削除

```bash
menv GOPATH C:\Go           # ユーザー環境変数を設定
menv -d GOPATH              # ユーザー環境変数を削除
```

### 📁 PATH 管理

```bash
menv -path                  # PATH を表示（1行に1パス）
menv -add "C:\bin"          # PATH にパスを追加
menv -rm "C:\bin"           # PATH からパスを削除
menv -clean                 # PATH をクリーンアップ（重複削除 + 無効パス削除）
menv -search java -path     # PATH 内で "java" を含むパスを検索
```

### 🔍 無効なパスのチェック

```bash
menv -check                 # PATH 内の無効なディレクトリをチェック
menv -check -fix            # 無効なパスを自動削除
menv -check -fix -i         # 各パスを削除する前に確認
```

### 💾 バックアップと復元

```bash
menv -backup backup.json    # 環境変数をバックアップ
menv -restore backup.json   # 環境変数を復元
menv -export env.sh         # シェルスクリプトとしてエクスポート
```

### 🛡️ システム環境変数

上記のコマンドはデフォルトで**ユーザー**環境変数を操作します。**システム**環境変数を操作するには `-sys` を追加してください（管理者権限が必要）：

```bash
menv -list -sys             # システム環境変数を一覧表示
menv -path -sys             # システム PATH を表示
menv -add "C:\bin" -sys     # システム PATH に追加
menv -clean -sys            # システム PATH をクリーンアップ
```

## 📖 よくあるシナリオ

| シナリオ | コマンド |
|:---------|:---------|
| 👁️ 特定の変数を表示 | `menv -get JAVA_HOME` |
| ➕ PATH にディレクトリを追加 | `menv -add "C:\tools\bin"` |
| 🧹 重複と無効なパスをクリーンアップ | `menv -clean` |
| 💾 現在の環境変数をバックアップ | `menv -backup my-env.json` |
| 🔎 Java 関連のパスを検索 | `menv -search java -path` |

## 📄 ライセンス

[LICENSE](LICENSE)
