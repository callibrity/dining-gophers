package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("Hello, World!")
	}()

	time.Sleep(1 * time.Second)
}
