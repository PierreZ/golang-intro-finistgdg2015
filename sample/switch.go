package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Print("Go runs on ")
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.", os)
	}

}
