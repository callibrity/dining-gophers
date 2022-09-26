package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const FORK = 1

func main() {
	var forks = []chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	var wg = sync.WaitGroup{}
	wg.Add(len(forks))
	for i := range forks {
		var left = i
		var right = (i + 1) % len(forks)
		p := i
		go func() {
			if left < right {
				runPhilosopher(p, forks[left], forks[right])
			} else {
				runPhilosopher(p, forks[right], forks[left])
			}
			wg.Done()
		}()
	}

	for i, c := range forks {
		c <- i
	}

	fmt.Printf("Awaiting wait group...\n")
	wg.Wait()

	fmt.Printf("Everyone is done eating!")
}

func runPhilosopher(philosopher int, first chan int, second chan int) {
	for i := 1; i < 5; i++ {
		think(philosopher)
		fmt.Printf("Philosopher %d picking up first fork...\n", philosopher)
		<-first
		fmt.Printf("Philosopher %d picked up first fork.\n", philosopher)
		fmt.Printf("Philosopher %d picking up second fork...\n", philosopher)
		<-second
		fmt.Printf("Philosopher %d picked up second fork...\n", philosopher)
		eat(philosopher)
		fmt.Printf("Philosopher %d putting down first fork...\n", philosopher)
		first <- FORK
		fmt.Printf("Philosopher %d putting down second fork...\n", philosopher)
		second <- FORK
	}
	fmt.Printf("Philosopher %d is no longer hungry.\n", philosopher)
}

func think(i int) {
	thinkTime := rand.Intn(500) + 500
	fmt.Printf("Philosopher %d thinking for %dms...\n", i, thinkTime)
	time.Sleep(time.Duration(thinkTime) * time.Millisecond)
}

func eat(i int) {
	eatTime := rand.Intn(500) + 500
	fmt.Printf("Philosopher %d eating for %dms...\n", i, eatTime)
	time.Sleep(time.Duration(eatTime) * time.Millisecond)
}
