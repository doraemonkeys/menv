package cmd

import "flag"

var (
	EnvFilePath = flag.String("file", "", "env file path")
	DelEnv      = flag.Bool("d", false, "delete env")
	SetSystem   = flag.Bool("sys", false, "set system env")
	StartWith   = flag.String("startWith", "", "line start with")
	AddPath     = flag.String("add", "", "add path")
	RemovePath  = flag.String("rm", "", "remove path from PATH")
	CleanPath   = flag.Bool("clean", false, "clean PATH (dedupe + remove invalid)")
	CheckPath   = flag.Bool("check", false, "check PATH for invalid directories")
	FixPath     = flag.Bool("fix", false, "auto-remove invalid paths (use with -check)")
	Interactive = flag.Bool("i", false, "interactive confirmation")
	ListEnv     = flag.Bool("list", false, "list all env vars")
	GetEnv      = flag.String("get", "", "get env var value")
	ShowPath    = flag.Bool("path", false, "display PATH")
	ExportPath  = flag.String("export", "", "export env vars to file (sh/bat/json)")
	BackupPath  = flag.String("backup", "", "backup env vars to JSON file")
	RestorePath = flag.String("restore", "", "restore env vars from backup file")
	Search      = flag.String("search", "", "search env vars by keyword")
)
