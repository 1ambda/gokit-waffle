package main

import (
	"fmt"
	"time"
)

var (
	Version   = "undefined"
	BuildTime = "undefined"
	GitHash   = "undefined"
	Started   = time.Now()
)

func main() {
	fmt.Println(GitHash)
	fmt.Println(BuildTime)
	fmt.Println(Version)
	fmt.Println(Started)
}
