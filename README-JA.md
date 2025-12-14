# menv

[English](README.md) | [简体中文](README-ZH.md) | [繁體中文](README-ZH-TW.md) | [한국어](README-KO.md) | [日本語](README-JA.md)

コマンドラインで環境変数と PATH を簡単に管理できる Windows 環境変数管理ツールです。

## インストール

```bash
go install github.com/doraemonkeys/menv@latest
```

またはクローンしてローカルインストール：

```bash
git clone https://github.com/doraemonkeys/menv.git
cd menv
go install
```

## クイックスタート

### 環境変数の表示

```bash
menv -list                  # すべてのユーザー環境変数を一覧表示
menv -get JAVA_HOME         # 特定の変数の値を取得
menv -search java           # java を含む変数を検索
```

### 環境変数の設定/削除

```bash
menv GOPATH C:\Go           # ユーザー環境変数を設定
menv -d GOPATH              # ユーザー環境変数を削除
```

### PATH 管理

```bash
menv -path                  # PATH を表示（1行に1つのパス）
menv -add "C:\bin"          # PATH にパスを追加
menv -rm "C:\bin"           # PATH からパスを削除
menv -clean                 # PATH をクリーンアップ（重複削除 + 無効なパスを削除）
menv -search java -path     # PATH 内で java を含むパスを検索
```

### 無効なパスの確認

```bash
menv -check                 # PATH 内の無効なディレクトリを確認
menv -check -fix            # 無効なパスを自動的に削除
menv -check -fix -i         # 各パスを削除する前に確認
```

### バックアップと復元

```bash
menv -backup backup.json    # 環境変数をバックアップ
menv -restore backup.json   # 環境変数を復元
menv -export env.sh         # シェルスクリプトとしてエクスポート
```

### システム環境変数

上記のコマンドはデフォルトで**ユーザー**環境変数を操作します。**システム**環境変数を操作するには `-sys` を追加してください（管理者権限が必要）：

```bash
menv -list -sys             # システム環境変数を一覧表示
menv -path -sys             # システム PATH を表示
menv -add "C:\bin" -sys     # システム PATH に追加
menv -clean -sys            # システム PATH をクリーンアップ
```

## よくあるシナリオ

| シナリオ | コマンド |
|----------|----------|
| 特定の変数を表示 | `menv -get JAVA_HOME` |
| PATH にディレクトリを追加 | `menv -add "C:\tools\bin"` |
| 重複と無効なパスをクリーンアップ | `menv -clean` |
| 現在の環境変数をバックアップ | `menv -backup my-env.json` |
| Java 関連のパスを検索 | `menv -search java -path` |
