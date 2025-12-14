package cmd

import "flag"

var (
	EnvFilePath = flag.String("file", "", "env file path")
	DelEnv      = flag.Bool("d", false, "delete env")
	SetSystem   = flag.Bool("sys", false, "set system env")
	StartWith   = flag.String("startWith", "", "line start with")
	AddPath     = flag.String("add", "", "add path")
	ListEnv     = flag.Bool("list", false, "list all env vars")
	GetEnv      = flag.String("get", "", "get env var value")
	ShowPath    = flag.Bool("path", false, "display PATH")
)
