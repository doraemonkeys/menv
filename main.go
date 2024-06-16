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

// setenv -file xxx.sh -startWith export
// setenv key value
// setenv -d -file xxx.sh
// setenv -d key
// setenv -sys key value
// -add path

var (
	envFilePath *string = flag.String("file", "", "env file path")
	delenv      *bool   = flag.Bool("d", false, "delete env")
	setSystem   *bool   = flag.Bool("sys", false, "set system env")
	startWith   *string = flag.String("startWith", "", "line start with")
	addPath     *string = flag.String("add", "", "add path")
)

func main() {
	var err error
	flag.Parse()
	args := flag.Args()
	if *delenv && *envFilePath == "" && len(args) != 1 ||
		(len(os.Args) <= 1) {
		flag.Usage()
		return
	}
	if *addPath != "" {
		err := addToPath(*addPath, *setSystem)
		if err != nil {
			panic(err)
		}
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
				if *setSystem {
					unsetSystemEnv(v.First)
				} else {
					unsetenv(v.First)
				}
			} else {
				if *setSystem {
					setSystemEnv(v.First, v.Second)
				} else {
					setenv(v.First, v.Second)
				}
			}
		}
		return
	}
	if *delenv {
		if *setSystem {
			err = unsetSystemEnv(args[0])
		} else {
			err = unsetenv(args[0])
		}
		if err != nil {
			panic(err)
		}
		return
	}

	if *setSystem {
		err = setSystemEnv(args[0], args[1])
	} else {
		err = setenv(args[0], args[1])
	}
	if err != nil {
		panic(err)
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
		if *startWith != "" && !strings.HasPrefix(trimmedLine, *startWith) {
			continue
		}
		// 去掉 "export " 前缀
		trimmedLine = strings.TrimPrefix(trimmedLine, *startWith+" ")
		// 按等号分割 key 和 value
		parts := strings.SplitN(trimmedLine, "=", 2)
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
// key, value 包含'&'没有问题
func setenvPS(key, value string) error {
	if os.Getenv(key) == value {
		fmt.Printf("skip %s=%s\n", key, value)
		return nil
	}
	fmt.Printf("set %s=%s\n", key, value)
	cmd := exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\""+key+"\", \""+value+"\", \"User\")")
	return cmd.Run()
}

func setSystemEnv(key, value string) error {
	fmt.Printf("set %s=%s\n", key, value)
	cmd := exec.Command("setx", key, value, "/m")
	return cmd.Run()
}

func unsetenv(key string) error {
	fmt.Printf("unset %s\n", key)
	cmd := exec.Command("reg", "delete", "HKEY_CURRENT_USER\\Environment", "/F", "/V", key)
	return cmd.Run()
}

func unsetSystemEnv(key string) error {
	fmt.Printf("unset %s\n", key)
	cmd := exec.Command("reg", "delete", "HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment", "/F", "/V", key)
	return cmd.Run()
}

// setx有一个问题是设置的最大长度是1024
// 所以path很长的时候，就不能用这个方法了。
func addToPath(add string, sys bool) error {
	if strings.ContainsRune(add, ';') {
		return errors.New("invalid path: " + add)
	}
	path := os.Getenv("PATH")
	if strings.Contains(path, add) {
		fmt.Printf("skip %s\n", add)
		return nil
	}
	fmt.Printf("add %s\n", add)
	if strings.HasSuffix(path, ";") {
		path += add
	} else {
		path += ";" + add
	}
	path += ";"
	var cmd *exec.Cmd
	if sys {
		cmd = exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\"PATH\", \""+path+"\", \"Machine\")")
	} else {
		cmd = exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\"PATH\", \""+path+"\", \"User\")")
	}
	return cmd.Run()
}
