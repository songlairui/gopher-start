package main

import (
	"fmt"
	"time"
	"math/rand"
)

func fate(everyday chan bool, rainbow chan bool) {
	a, b := rand.Float64(), rand.Float64()
	if a == b {
		rainbow <- true
	} else {
		everyday <- false
	}
}

func main() {
	everyday, rainbow := make(chan bool, 10), make(chan bool, 10)
	for {
		select {
		case <-everyday:
			fmt.Println("where somebody?")
		case <-rainbow:
			fmt.Println("FOUND!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(2 * time.Second)
		}
		fate(everyday, rainbow)
	}
}
