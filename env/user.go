package env

import (
	"fmt"
	"os"
	"os/exec"
)

// Set sets a user environment variable using setx command.
func Set(key, value string) error {
	if os.Getenv(key) == value {
		fmt.Printf("skip %s=%s\n", key, value)
		return nil
	}
	fmt.Printf("set %s=%s\n", key, value)
	cmd := exec.Command("setx", key, value)
	return cmd.Run()
}

// SetPS sets a user environment variable using PowerShell.
// This method handles keys and values containing '&' correctly.
func SetPS(key, value string) error {
	if os.Getenv(key) == value {
		fmt.Printf("skip %s=%s\n", key, value)
		return nil
	}
	fmt.Printf("set %s=%s\n", key, value)
	cmd := exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\""+key+"\", \""+value+"\", \"User\")")
	return cmd.Run()
}

// Unset removes a user environment variable.
func Unset(key string) error {
	fmt.Printf("unset %s\n", key)
	cmd := exec.Command("reg", "delete", "HKEY_CURRENT_USER\\Environment", "/F", "/V", key)
	return cmd.Run()
}
