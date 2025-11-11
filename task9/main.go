package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	input := make(chan int)
	output := make(chan int)

	go func() {
		defer close(output)
		for v := range input {
			output <- v * v
		}
	}()

	go func() {
		defer close(input)
		for _, v := range nums {
			input <- v
		}
	}()

	for v := range output {
		fmt.Println(v)
	}

}
