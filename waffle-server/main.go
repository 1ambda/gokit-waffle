package main

import (
	"fmt"
	"time"

	arg "github.com/alexflint/go-arg"
)

// Build Flags
var (
	Version   = "undefined"
	BuildTime = "undefined"
	GitHash   = "undefined"
	Started   = time.Now()
)

// Env Variables
var EnvVar struct {
	Mode string `arg:"env"`
}

func main() {
	fmt.Println(GitHash)
	fmt.Println(BuildTime)
	fmt.Println(Version)
	fmt.Println(Started)

	// Setup default values
	EnvVar.Mode = "LOCAL" // [LOCAL, DEV, PROD]
	arg.MustParse(&EnvVar)
	fmt.Println(EnvVar.Mode)
}
