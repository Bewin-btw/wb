package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

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
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <number_of_workers> <timeout_seconds>\n", os.Args[0])
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Printf("Invalid number of workers\n")
		os.Exit(1)
	}

	timeoutSec, err := strconv.Atoi(os.Args[2])
	if err != nil || timeoutSec <= 0 {
		fmt.Printf("Invalid timeout value\n")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	ch := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	var writerWg sync.WaitGroup
	writerWg.Add(1)

	go func() {
		defer writerWg.Done()
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

	<-ctx.Done()
	fmt.Printf("\nProgram finished after %d seconds\n", timeoutSec)

	writerWg.Wait()
	close(ch)
	wg.Wait()
	fmt.Println("all workers finished gracefully")
}
