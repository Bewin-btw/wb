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

// я выбрал контексты из за идиоматичности, также там есть встроенные дедлайны и таймауты

func worker(ctx context.Context, id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %d: channel closed, finishing\n", id)
				return
			}
			fmt.Printf("Worker %d: %d\n", id, data)
		case <-ctx.Done():
			fmt.Printf("Worker %d: received cancel signal, finishing\n", id)
			return
		}
	}
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

	ch := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- counter:
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
