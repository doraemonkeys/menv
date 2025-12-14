package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/doraemonkeys/menv/cmd"
	"github.com/doraemonkeys/menv/color"
	"github.com/doraemonkeys/menv/env"
	"github.com/doraemonkeys/menv/path"
)

func init() {
	flag.Usage = func() {
		fmt.Println(color.Sprintf(color.BoldCyan, "menv") + " - Windows Environment Variable Manager")
		fmt.Println()
		color.Info("Usage:")
		fmt.Println("  menv [options] [key] [value]")
		fmt.Println()
		color.Info("Options:")
		fmt.Println("  -add <path>       Add path to PATH variable")
		fmt.Println("  -d                Delete environment variable")
		fmt.Println("  -sys              Target system env (default: user)")
		fmt.Println("  -file <path>      Read env vars from file")
		fmt.Println("  -startWith <str>  Filter lines starting with string")
		fmt.Println()
		color.Info("Examples:")
		fmt.Println("  menv GOPATH C:\\Go                  # Set user env var")
		fmt.Println("  menv -sys GOPATH C:\\Go             # Set system env var")
		fmt.Println("  menv -d GOPATH                     # Delete user env var")
		fmt.Println("  menv -d -sys GOPATH                # Delete system env var")
		fmt.Println("  menv -add \"C:\\bin\"                 # Add to user PATH")
		fmt.Println("  menv -add \"C:\\bin\" -sys            # Add to system PATH")
		fmt.Println("  menv -file env.sh -startWith export")
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	if err := run(args); err != nil {
		color.Error("Error: %v", err)
		os.Exit(1)
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
			return fmt.Errorf("unexpected arguments: %v", args)
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
	if len(args) < 2 {
		return fmt.Errorf("missing value for key '%s'", args[0])
	}
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
