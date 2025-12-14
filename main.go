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
		fmt.Println("  -list             List all env vars")
		fmt.Println("  -get <key>        Get env var value")
		fmt.Println("  -path             Display PATH (one per line)")
		fmt.Println("  -add <path>       Add path to PATH variable")
		fmt.Println("  -rm <path>        Remove path from PATH variable")
		fmt.Println("  -clean            Clean PATH (dedupe + remove invalid)")
		fmt.Println("  -d                Delete environment variable")
		fmt.Println("  -sys              Target system env (default: user)")
		fmt.Println("  -file <path>      Read env vars from file")
		fmt.Println("  -startWith <str>  Filter lines starting with string")
		fmt.Println("  -export <path>    Export env vars to file (sh/bat/json)")
		fmt.Println()
		color.Info("Examples:")
		fmt.Println("  menv -list                         # List user env vars")
		fmt.Println("  menv -list -sys                    # List system env vars")
		fmt.Println("  menv -get JAVA_HOME                # Get JAVA_HOME value")
		fmt.Println("  menv -path                         # Display user PATH")
		fmt.Println("  menv -path -sys                    # Display system PATH")
		fmt.Println("  menv GOPATH C:\\Go                  # Set user env var")
		fmt.Println("  menv -sys GOPATH C:\\Go             # Set system env var")
		fmt.Println("  menv -d GOPATH                     # Delete user env var")
		fmt.Println("  menv -d -sys GOPATH                # Delete system env var")
		fmt.Println("  menv -add \"C:\\bin\"                 # Add to user PATH")
		fmt.Println("  menv -add \"C:\\bin\" -sys            # Add to system PATH")
		fmt.Println("  menv -rm \"C:\\bin\"                  # Remove from user PATH")
		fmt.Println("  menv -rm \"C:\\bin\" -sys             # Remove from system PATH")
		fmt.Println("  menv -clean                        # Clean user PATH")
		fmt.Println("  menv -clean -sys                   # Clean system PATH")
		fmt.Println("  menv -file env.sh -startWith export")
		fmt.Println("  menv -export env.sh                # Export user env as shell")
		fmt.Println("  menv -export env.bat               # Export user env as batch")
		fmt.Println("  menv -export env.json              # Export user env as JSON")
		fmt.Println("  menv -export env.json -sys         # Export system env as JSON")
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
	// Handle -list flag: list all env vars
	if *cmd.ListEnv {
		return listEnvVars()
	}

	// Handle -get flag: get specific env var
	if *cmd.GetEnv != "" {
		return getEnvVar(*cmd.GetEnv)
	}

	// Handle -path flag: display PATH
	if *cmd.ShowPath {
		return showPath()
	}

	// Handle -export flag: export env vars to file
	if *cmd.ExportPath != "" {
		return exportEnvVars(*cmd.ExportPath)
	}

	// Validate arguments
	noArgsRequired := *cmd.RemovePath != "" || *cmd.CleanPath
	if !noArgsRequired && (*cmd.DelEnv && *cmd.EnvFilePath == "" && len(args) != 1 || len(os.Args) <= 1) {
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

	// Handle -rm flag: remove path from PATH
	if *cmd.RemovePath != "" {
		if len(args) != 0 {
			return fmt.Errorf("unexpected arguments: %v", args)
		}
		return path.Remove(*cmd.RemovePath, *cmd.SetSystem)
	}

	// Handle -clean flag: clean PATH
	if *cmd.CleanPath {
		if len(args) != 0 {
			return fmt.Errorf("unexpected arguments: %v", args)
		}
		return path.Clean(*cmd.SetSystem)
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
		if err := applyEnvVar(v.First, v.Second); err != nil {
			return err
		}
	}

	return nil
}

func applyEnvVar(key, value string) error {
	if *cmd.DelEnv {
		if *cmd.SetSystem {
			return env.UnsetSystem(key)
		}
		return env.Unset(key)
	}
	if *cmd.SetSystem {
		return env.SetSystem(key, value)
	}
	return env.Set(key, value)
}

func listEnvVars() error {
	var envVars []env.EnvVar
	var err error

	if *cmd.SetSystem {
		color.Info("System Environment Variables:")
		envVars, err = env.ListSystem()
	} else {
		color.Info("User Environment Variables:")
		envVars, err = env.ListUser()
	}

	if err != nil {
		return err
	}

	fmt.Println()
	for _, e := range envVars {
		fmt.Printf("%s%s%s=%s\n", color.Green, e.Key, color.Reset, e.Value)
	}
	fmt.Printf("\nTotal: %d\n", len(envVars))
	return nil
}

func getEnvVar(key string) error {
	var value string
	var err error

	if *cmd.SetSystem {
		value, err = env.GetSystem(key)
	} else {
		value, err = env.GetUser(key)
	}

	if err != nil {
		return err
	}

	if value == "" {
		color.Warning("%s is not set", key)
		return nil
	}

	fmt.Printf("%s%s%s=%s\n", color.Green, key, color.Reset, value)
	return nil
}

func showPath() error {
	var paths []string
	var err error

	if *cmd.SetSystem {
		color.Info("System PATH:")
		paths, err = path.QuerySystemPath()
	} else {
		color.Info("User PATH:")
		paths, err = path.QueryUserPath()
	}

	if err != nil {
		return err
	}

	fmt.Println()
	for i, p := range paths {
		fmt.Printf("%s%3d%s  %s\n", color.Cyan, i+1, color.Reset, p)
	}
	fmt.Printf("\nTotal: %d\n", len(paths))
	return nil
}

func exportEnvVars(filename string) error {
	var envVars []env.EnvVar
	var err error

	if *cmd.SetSystem {
		envVars, err = env.ListSystem()
	} else {
		envVars, err = env.ListUser()
	}

	if err != nil {
		return err
	}

	if err := env.Export(filename, envVars); err != nil {
		return err
	}

	format := env.DetectFormat(filename)
	color.Success("Exported %d env vars to %s (format: %s)", len(envVars), filename, format)
	return nil
}
