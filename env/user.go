package env

import (
	"os"
	"os/exec"

	"github.com/doraemonkeys/menv/color"
)

// Set sets a user environment variable using setx command.
func Set(key, value string) error {
	if os.Getenv(key) == value {
		color.Warning("skip %s=%s", key, value)
		return nil
	}
	cmd := exec.Command("setx", key, value)
	if err := cmd.Run(); err != nil {
		return err
	}
	color.Success("set  %s=%s", key, value)
	return nil
}

// SetPS sets a user environment variable using PowerShell.
// This method handles keys and values containing '&' correctly.
func SetPS(key, value string) error {
	if os.Getenv(key) == value {
		color.Warning("skip %s=%s", key, value)
		return nil
	}
	cmd := exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\""+key+"\", \""+value+"\", \"User\")")
	if err := cmd.Run(); err != nil {
		return err
	}
	color.Success("set  %s=%s", key, value)
	return nil
}

// Unset removes a user environment variable.
func Unset(key string) error {
	cmd := exec.Command("reg", "delete", "HKEY_CURRENT_USER\\Environment", "/F", "/V", key)
	if err := cmd.Run(); err != nil {
		return err
	}
	color.Success("unset %s", key)
	return nil
}
