package main

import (
	"dining-gophers/sleep"
	"fmt"
	"sync"
)

const nForks = 5
const hunger = 5

func main() {
	var forks = make([]chan int, nForks)
	for i := 0; i < nForks; i++ {
		forks[i] = make(chan int, 1)
	}
	var wg = sync.WaitGroup{}
	wg.Add(nForks)

	for i := 0; i < nForks; i++ {
		var left = i
		var right = (i + 1) % nForks

		go func(gopher int) {
			for h := 0; h < hunger; h++ {
				think(gopher)
				fmt.Printf("Gopher %d finished thinking, time to pick up forks.\n", gopher)
				if left < right {
					fmt.Printf("Gopher %d picking up fork %d...\n", gopher, left)
					<-forks[left]
					fmt.Printf("Gopher %d picked up fork %d.\n", gopher, left)

					fmt.Printf("Gopher %d picking up fork %d...\n", gopher, right)
					<-forks[right]
					fmt.Printf("Gopher %d picked up fork %d.\n", gopher, right)

				} else {
					fmt.Printf("Gopher %d picking up fork %d...\n", gopher, right)
					<-forks[right]
					fmt.Printf("Gopher %d picked up fork %d.\n", gopher, right)

					fmt.Printf("Gopher %d picking up fork %d...\n", gopher, left)
					<-forks[left]
					fmt.Printf("Gopher %d picked up fork %d.\n", gopher, left)

				}
				fmt.Printf("Gopher %d picked up both forks, time to eat...\n", gopher)
				eat(gopher)
				fmt.Printf("Gopher %d finished eating, time to put down forks.\n", gopher)
				fmt.Printf("Gopher %d trying to put down fork %d.\n", gopher, left)
				forks[left] <- left
				fmt.Printf("Gopher %d put down fork %d.\n", gopher, left)
				fmt.Printf("Gopher %d trying to put down fork %d.\n", gopher, right)
				forks[right] <- right
				fmt.Printf("Gopher %d put down fork %d.\n", gopher, right)
			}
			wg.Done()
		}(i)

	}

	for i := 0; i < nForks; i++ {
		forks[i] <- i
	}

	fmt.Println("Waiting for all gophers to finish thinking/eating...")
	wg.Wait()
	fmt.Println("Everybody is finished!")
}

func think(gopher int) {
	fmt.Printf("Gopher %d thinking...\n", gopher)
	sleep.RandomSleep()
}

func eat(gopher int) {
	fmt.Printf("Gopher %d eating...\n", gopher)
	sleep.RandomSleep()
}
