package env

import (
	"os/exec"

	"github.com/doraemonkeys/menv/color"
)

// SetSystem sets a system environment variable using setx /m command.
// Requires administrator privileges.
func SetSystem(key, value string) error {
	cmd := exec.Command("setx", key, value, "/m")
	if err := cmd.Run(); err != nil {
		return err
	}
	color.Success("set  %s=%s [system]", key, value)
	return nil
}

// UnsetSystem removes a system environment variable.
// Requires administrator privileges.
func UnsetSystem(key string) error {
	cmd := exec.Command("reg", "delete", "HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment", "/F", "/V", key)
	if err := cmd.Run(); err != nil {
		return err
	}
	color.Success("unset %s [system]", key)
	return nil
}
