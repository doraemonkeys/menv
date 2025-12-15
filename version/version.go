package version

import "fmt"

var (
	Version   = "0.1.1"
	BuildHash = "unknown hash" // will insert in build time
	BuildTime = "unknown time" // will insert in build time
)

func PrintVersion() {
	fmt.Println("menv", "v"+Version)
	fmt.Println("BuildTime:", BuildTime)
	fmt.Println("BuildHash:", BuildHash)
}
