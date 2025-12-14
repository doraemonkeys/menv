package cmd

import "flag"

var (
	EnvFilePath = flag.String("file", "", "env file path")
	DelEnv      = flag.Bool("d", false, "delete env")
	SetSystem   = flag.Bool("sys", false, "set system env")
	StartWith   = flag.String("startWith", "", "line start with")
	AddPath     = flag.String("add", "", "add path")
)
