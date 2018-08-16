package main

import "fmt"

var a bool

func add(x int, y int) int {
	return x + y
}

func main() {
	done := make(chan bool)
	defer fmt.Println("world", done)

	fmt.Println("hello")
	fmt.Println(add(42, 13))
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
