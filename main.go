package main

import (
	"flag"
	"fmt"
	"os"

	"setenv/cmd"
	"setenv/env"
	"setenv/path"
)

// Usage examples:
// setenv -file xxx.sh -startWith export  # Set env vars from file
// setenv key value                        # Set user env var
// setenv -sys key value                   # Set system env var
// setenv -d key                           # Delete user env var
// setenv -d -sys key                      # Delete system env var
// setenv -d -file xxx.sh                  # Delete env vars from file
// setenv -add "C:\path\to\dir"            # Add to user PATH
// setenv -add "C:\path\to\dir" -sys       # Add to system PATH

func main() {
	flag.Parse()
	args := flag.Args()

	if err := run(args); err != nil {
		panic(err)
	}
}

func run(args []string) error {
	// Validate arguments
	if *cmd.DelEnv && *cmd.EnvFilePath == "" && len(args) != 1 || len(os.Args) <= 1 {
		flag.Usage()
		return nil
	}

	// Handle -add flag: add path to PATH
	if *cmd.AddPath != "" {
		if len(args) != 0 {
			fmt.Printf("invalid args: %v\n", args)
			return nil
		}
		return path.Add(*cmd.AddPath, *cmd.SetSystem)
	}

	// Handle -file flag: process env file
	if *cmd.EnvFilePath != "" {
		return processEnvFile()
	}

	// Handle -d flag: delete env var
	if *cmd.DelEnv {
		if *cmd.SetSystem {
			return env.UnsetSystem(args[0])
		}
		return env.Unset(args[0])
	}

	// Default: set env var
	if *cmd.SetSystem {
		return env.SetSystem(args[0], args[1])
	}
	return env.Set(args[0], args[1])
}

func processEnvFile() error {
	content, err := os.ReadFile(*cmd.EnvFilePath)
	if err != nil {
		return err
	}

	envMap, err := env.ParseEnvFile(content, *cmd.StartWith)
	if err != nil {
		return err
	}

	for _, v := range envMap {
		if *cmd.DelEnv {
			if *cmd.SetSystem {
				env.UnsetSystem(v.First)
			} else {
				env.Unset(v.First)
			}
		} else {
			if *cmd.SetSystem {
				env.SetSystem(v.First, v.Second)
			} else {
				env.Set(v.First, v.Second)
			}
		}
	}

	return nil
}
