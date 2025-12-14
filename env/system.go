package env

import (
	"fmt"
	"os/exec"
)

// SetSystem sets a system environment variable using setx /m command.
// Requires administrator privileges.
func SetSystem(key, value string) error {
	fmt.Printf("set %s=%s\n", key, value)
	cmd := exec.Command("setx", key, value, "/m")
	return cmd.Run()
}

// UnsetSystem removes a system environment variable.
// Requires administrator privileges.
func UnsetSystem(key string) error {
	fmt.Printf("unset %s\n", key)
	cmd := exec.Command("reg", "delete", "HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment", "/F", "/V", key)
	return cmd.Run()
}
