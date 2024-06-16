package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/doraemonkeys/doraemon"
)

// xxx.sh
// export xxx=xxx

// setenv -file xxx.sh
// setenv key value
// setenv -d -file xxx.sh
// setenv -d key
// setenv -su key value

var (
	envFilePath *string = flag.String("file", "", "env file path")
	delenv      *bool   = flag.Bool("d", false, "delete env")
	superuser   *bool   = flag.Bool("su", false, "super user")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if *delenv && *envFilePath == "" && len(args) != 1 ||
		(len(os.Args) <= 1) {
		flag.Usage()
		return
	}
	if *envFilePath != "" {
		content, err := os.ReadFile(*envFilePath)
		if err != nil {
			panic(err)
		}
		envMap, err := parseEnvFile(content)
		if err != nil {
			panic(err)
		}
		for _, v := range envMap {
			if *delenv {
				if *superuser {
					unsetenvSuperUser(v.First)
				} else {
					unsetenv(v.First)
				}
			} else {
				if *superuser {
					setenvSuperUser(v.First, v.Second)
				} else {
					setenv(v.First, v.Second)
				}
			}
		}
		return
	}
	if *delenv {
		if *superuser {
			unsetenvSuperUser(args[0])
		} else {
			unsetenv(args[0])
		}
		return
	}

	if *superuser {
		setenvSuperUser(args[0], args[1])
	} else {
		setenv(args[0], args[1])
	}
}

func parseEnvFile(content []byte) ([]doraemon.Pair[string, string], error) {
	// envMap := make(map[string]string)
	envMap := []doraemon.Pair[string, string]{}
	// 将文件内容按行分割
	lines := bytes.Split(content, []byte("\n"))
	for _, line := range lines {
		// 将行前后的空白字符去掉
		trimmedLine := strings.TrimSpace(string(line))
		// 忽略空行和注释行
		if len(trimmedLine) == 0 || strings.HasPrefix(trimmedLine, "#") {
			continue
		}
		// 检查是否以 "export " 开头
		if strings.HasPrefix(trimmedLine, "export ") {
			// 去掉 "export " 前缀
			kv := strings.TrimPrefix(trimmedLine, "export ")
			// 按等号分割 key 和 value
			parts := strings.SplitN(kv, "=", 2)
			if len(parts) != 2 {
				return nil, errors.New("invalid format: missing '=' in line: " + trimmedLine)
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if strings.Contains(value, "=") {
				return nil, errors.New("invalid format: multiple '=' in line: " + trimmedLine)
			}
			// envMap[key] = value
			envMap = append(envMap, doraemon.Pair[string, string]{First: key, Second: value})
		}
	}
	return envMap, nil
}

// setx key value
func setenv(key, value string) error {
	if os.Getenv(key) == value {
		fmt.Printf("skip %s=%s\n", key, value)
		return nil
	}
	fmt.Printf("set %s=%s\n", key, value)
	cmd := exec.Command("setx", key, value)
	return cmd.Run()
}

// [System.Environment]::SetEnvironmentVariable("MY_ENV_VAR", "HelloWorld", "User")
func setenvPS(key, value string) error {
	if os.Getenv(key) == value {
		fmt.Printf("skip %s=%s\n", key, value)
		return nil
	}
	fmt.Printf("set %s=%s\n", key, value)
	cmd := exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\""+key+"\", \""+value+"\", \"User\")")
	return cmd.Run()
}

func setenvSuperUser(key, value string) error {
	fmt.Printf("set %s=%s\n", key, value)
	cmd := exec.Command("setx", key, value, "/m")
	return cmd.Run()
}

func unsetenv(key string) error {
	fmt.Printf("unset %s\n", key)
	cmd := exec.Command("reg", "delete", "HKEY_CURRENT_USER\\Environment", "/F", "/V", key)
	return cmd.Run()
}

func unsetenvSuperUser(key string) error {
	fmt.Printf("unset %s\n", key)
	cmd := exec.Command("reg", "delete", "HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment", "/F", "/V", key)
	return cmd.Run()
}
