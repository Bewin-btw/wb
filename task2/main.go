package main

import "fmt"

func main() {
	array := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(array))

	for _, n := range array {
		go func(x int) {
			ch <- x * x
		}(n)
	}

	for i := 0; i < len(array); i++ {
		fmt.Println(<-ch)
	}
}
