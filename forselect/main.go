package main

import (
	"fmt"
	"time"
)

func main() {

	var c1 = make(chan string)
	var c2 = make(chan string)

	go func() {
		//fmt.Printf("Hello, %s!\n", <-c1)
		//fmt.Printf("Hello, %s!\n", <-c2)
		for {
			select {
			case name := <-c1:
				fmt.Printf("Hello, %s!\n", name)
			case name := <-c2:
				fmt.Printf("Hello, %s!\n", name)
			}
		}
	}()

	c2 <- "Foo"
	c1 <- "Bar"

	time.Sleep(1 * time.Second)

}
