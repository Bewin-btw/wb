package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Worker %d: %d\n", id, data)
	}
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Invalid num of workers\n")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Printf("Invalid num of workers\n")
		os.Exit(1)
	}

	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	counter := 1
	for {
		ch <- counter
		counter++
		time.Sleep(500 * time.Millisecond)
	}
}
