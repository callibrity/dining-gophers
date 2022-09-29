package main

import "fmt"

func main() {
	var ch = make(chan string)

	go func() {
		ch <- "Hello, World!"
	}()

	var msg = <-ch
	fmt.Println(msg)
}
