package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- counter
				counter++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-sigChan
	fmt.Println("\nShutting down")
	cancel()
	close(ch)
	wg.Wait()
}
