package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
		//fallthrough
	case "linux":
		fmt.Println("Linux")
		//fallthrough
	default:
		fmt.Printf("%s.\n", os)
	}
}
