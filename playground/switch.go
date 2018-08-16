package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GO runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "liunx":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}
