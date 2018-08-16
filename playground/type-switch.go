package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is an int\n", v)
	case string:
		fmt.Printf("%v is a %v long string\n", v, len(v))
	default:
		fmt.Printf("typeof v is %T\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(false)
	do(nil)
}
