package main

import (
	"fmt"
	"sync"
)

func main() {
	const n = 10

	var messages = make(chan string, n)

	var wg = sync.WaitGroup{}
	wg.Add(n) // Let the wait group know we have n things to do

	for i := 0; i < n; i++ {

		go func(msgNumber int) {
			messages <- fmt.Sprintf("Hello, World (%d)!", msgNumber)
			wg.Done() // Tell the wait group we're done with one thing
		}(i)
	}

	wg.Wait() // Wait for the wait group to be finished with all n things

	for i := 0; i < n; i++ {
		fmt.Println(<-messages)
	}
}
