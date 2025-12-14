package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

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
		fmt.Println("  -check            Check PATH for invalid directories")
		fmt.Println("  -fix              Auto-remove invalid paths (use with -check)")
		fmt.Println("  -i                Interactive confirmation")
		fmt.Println("  -d                Delete environment variable")
		fmt.Println("  -sys              Target system env (default: user)")
		fmt.Println("  -file <path>      Read env vars from file")
		fmt.Println("  -startWith <str>  Filter lines starting with string")
		fmt.Println("  -export <path>    Export env vars to file (sh/bat/json)")
		fmt.Println("  -backup <path>    Backup env vars to JSON file")
		fmt.Println("  -restore <path>   Restore env vars from backup file")
		fmt.Println("  -search <keyword> Search env vars by keyword")
		fmt.Println("                    Use with -path to search in PATH")
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
		fmt.Println("  menv -backup backup.json           # Backup user env vars")
		fmt.Println("  menv -backup backup.json -sys      # Backup system env vars")
		fmt.Println("  menv -restore backup.json          # Restore user env vars")
		fmt.Println("  menv -restore backup.json -sys     # Restore system env vars")
		fmt.Println("  menv -search java                  # Search env vars for 'java'")
		fmt.Println("  menv -search java -path            # Search PATH for 'java'")
		fmt.Println("  menv -check                        # Check user PATH for invalid dirs")
		fmt.Println("  menv -check -sys                   # Check system PATH for invalid dirs")
		fmt.Println("  menv -check -fix                   # Check and remove invalid paths")
		fmt.Println("  menv -check -fix -i                # Check and remove with confirmation")
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

	// Handle -search flag: search env vars or PATH
	if *cmd.Search != "" {
		if *cmd.ShowPath {
			return searchPath(*cmd.Search)
		}
		return searchEnvVars(*cmd.Search)
	}

	// Handle -path flag: display PATH
	if *cmd.ShowPath {
		return showPath()
	}

	// Handle -export flag: export env vars to file
	if *cmd.ExportPath != "" {
		return exportEnvVars(*cmd.ExportPath)
	}

	// Handle -backup flag: backup env vars to file
	if *cmd.BackupPath != "" {
		return backupEnvVars(*cmd.BackupPath)
	}

	// Handle -restore flag: restore env vars from file
	if *cmd.RestorePath != "" {
		return restoreEnvVars(*cmd.RestorePath)
	}

	// Handle PATH modification commands
	if handled, err := handlePathCommands(args); handled {
		return err
	}

	// Validate arguments
	if *cmd.DelEnv && *cmd.EnvFilePath == "" && len(args) != 1 || len(os.Args) <= 1 {
		flag.Usage()
		return nil
	}

	// Handle -file flag: process env file
	if *cmd.EnvFilePath != "" {
		return processEnvFile()
	}

	// Handle -d flag: delete env var
	if *cmd.DelEnv {
		return deleteEnvVar(args[0])
	}

	// Default: set env var
	return setEnvVar(args)
}

func handlePathCommands(args []string) (handled bool, err error) {
	if *cmd.AddPath != "" {
		if len(args) != 0 {
			return true, fmt.Errorf("unexpected arguments: %v", args)
		}
		return true, path.Add(*cmd.AddPath, *cmd.SetSystem)
	}

	if *cmd.RemovePath != "" {
		if len(args) != 0 {
			return true, fmt.Errorf("unexpected arguments: %v", args)
		}
		return true, path.Remove(*cmd.RemovePath, *cmd.SetSystem)
	}

	if *cmd.CleanPath {
		if len(args) != 0 {
			return true, fmt.Errorf("unexpected arguments: %v", args)
		}
		return true, path.Clean(*cmd.SetSystem)
	}

	if *cmd.CheckPath {
		if len(args) != 0 {
			return true, fmt.Errorf("unexpected arguments: %v", args)
		}
		return true, checkPath()
	}

	return false, nil
}

func deleteEnvVar(key string) error {
	if *cmd.SetSystem {
		return env.UnsetSystem(key)
	}
	return env.Unset(key)
}

func setEnvVar(args []string) error {
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

func backupEnvVars(filename string) error {
	count, err := env.Backup(filename, *cmd.SetSystem)
	if err != nil {
		return err
	}

	source := "user"
	if *cmd.SetSystem {
		source = "system"
	}
	color.Success("Backed up %d %s env vars to %s", count, source, filename)
	return nil
}

func restoreEnvVars(filename string) error {
	backup, err := env.LoadBackup(filename)
	if err != nil {
		return err
	}

	target := "user"
	if *cmd.SetSystem {
		target = "system"
	}
	color.Info("Restoring %d env vars from %s backup (created: %s) to %s...",
		len(backup.EnvVars), backup.Source, backup.CreatedAt.Format("2006-01-02 15:04:05"), target)

	count, err := env.Restore(filename, *cmd.SetSystem)
	if err != nil {
		return err
	}

	color.Success("Restored %d env vars", count)
	return nil
}

func searchEnvVars(keyword string) error {
	var results []env.EnvVar
	var err error

	if *cmd.SetSystem {
		color.Info("Searching system env vars for '%s':", keyword)
		results, err = env.SearchSystem(keyword)
	} else {
		color.Info("Searching user env vars for '%s':", keyword)
		results, err = env.SearchUser(keyword)
	}

	if err != nil {
		return err
	}

	if len(results) == 0 {
		color.Warning("No matches found")
		return nil
	}

	fmt.Println()
	for _, e := range results {
		fmt.Printf("%s%s%s=%s\n", color.Green, e.Key, color.Reset, e.Value)
	}
	fmt.Printf("\nFound: %d\n", len(results))
	return nil
}

func searchPath(keyword string) error {
	var results []string
	var err error

	if *cmd.SetSystem {
		color.Info("Searching system PATH for '%s':", keyword)
		results, err = path.SearchSystemPath(keyword)
	} else {
		color.Info("Searching user PATH for '%s':", keyword)
		results, err = path.SearchUserPath(keyword)
	}

	if err != nil {
		return err
	}

	if len(results) == 0 {
		color.Warning("No matches found")
		return nil
	}

	fmt.Println()
	for i, p := range results {
		fmt.Printf("%s%3d%s  %s\n", color.Cyan, i+1, color.Reset, p)
	}
	fmt.Printf("\nFound: %d\n", len(results))
	return nil
}

func checkPath() error {
	scope := "user"
	if *cmd.SetSystem {
		scope = "system"
	}

	color.Info("Checking %s PATH for invalid directories...", scope)

	invalid, err := path.Check(*cmd.SetSystem)
	if err != nil {
		return err
	}

	if len(invalid) == 0 {
		color.Success("All paths are valid!")
		return nil
	}

	fmt.Println()
	for _, p := range invalid {
		fmt.Printf("%s%3d%s  %s%s%s\n", color.Cyan, p.Index, color.Reset, color.Red, p.Path, color.Reset)
	}
	fmt.Printf("\nFound %d invalid path(s)\n", len(invalid))

	if !*cmd.FixPath {
		color.Info("Use -fix to remove invalid paths")
		return nil
	}

	if *cmd.Interactive {
		if !confirmAction(fmt.Sprintf("Remove %d invalid path(s)?", len(invalid))) {
			color.Warning("Cancelled")
			return nil
		}
	}

	return path.RemoveInvalidPaths(invalid, *cmd.SetSystem)
}

func confirmAction(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [y/N]: ", prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	input = strings.TrimSpace(strings.ToLower(input))
	return input == "y" || input == "yes"
}
