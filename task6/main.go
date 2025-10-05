package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func exitByCondition() {
	fmt.Println(" _1.Exit by condition_ ")

	go func() {
		counter := 0
		for counter < 5 {
			fmt.Printf("routine working, counter: %d\n", counter)
			counter++
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("routine ended by condition")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println()
}

func exitByChannel() {
	fmt.Println(" _2.Using channel notification_ ")

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("routine received a signal to stop(nil struct)")
				return
			default:
				fmt.Println("routine is working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(done)
	time.Sleep(500 * time.Millisecond)
	fmt.Println()
}

func exitByContextCancel() {
	fmt.Println(" _3.Using context (WithCancel)_ ")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("routine received a signal to stop")
				return
			default:
				fmt.Println("routine is working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
	fmt.Println()
}

func exitByContextTimeout() {
	fmt.Println(" _4.Using context (timeout)_ ")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("routine stopped by timeout")
				return
			default:
				fmt.Println("routine is working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println()
}

func exitByRuntimeGoexit() {
	fmt.Println(" _5.Using runtime.Goexit()_ ")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("inner routine is calling Goexit()")
			runtime.Goexit()
			fmt.Println("this part is not going to be executed")
		}()

		time.Sleep(2 * time.Second)
		fmt.Println("outer routine is done successfully")
	}()

	wg.Wait()
	fmt.Println()
}

func exitByPanicRecover() {
	fmt.Println(" _6.With panic and recover_")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recovering after panic: %v\n", r)
			}
		}()

		fmt.Println("routine starts working")
		time.Sleep(1 * time.Second)
		panic("aritificial panic!")
		fmt.Println("this code is not going to be executed")
	}()

	wg.Wait()
	fmt.Println()
}

func exitByChannelClose() {
	fmt.Println(" _7.Using channel closing_")

	dataChan := make(chan int)

	go func() {
		for data := range dataChan {
			fmt.Printf("proccessing of data: %d\n", data)
		}
		fmt.Println("channel is closed, routine stopping")
	}()

	for i := 0; i < 3; i++ {
		dataChan <- i
		time.Sleep(500 * time.Millisecond)
	}

	close(dataChan)
	time.Sleep(500 * time.Millisecond)
	fmt.Println()
}

func exitByTimer() {
	fmt.Println(" _8.Using timer_ ")

	timer := time.NewTimer(2 * time.Second)

	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("time ended, routine is stopping")
				return
			default:
				fmt.Println("routine is working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println()
}

func exitCombined() {
	fmt.Println(" _9.Context + channel_ ")

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("receive signal from ctx")
				close(done)
				return
			case <-time.After(500 * time.Millisecond):
				fmt.Println("routine is working")
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()

	<-done
	fmt.Println("routine successfully stopped")
	fmt.Println()
}

func main() {
	exitByCondition()
	exitByChannel()
	exitByContextCancel()
	exitByContextTimeout()
	exitByRuntimeGoexit()
	exitByPanicRecover()
	exitByChannelClose()
	exitByTimer()
	exitCombined()

	fmt.Println("All of the examples are done")
}
