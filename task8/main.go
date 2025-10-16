package main

import (
	"fmt"
)

// & - and, 	&^ - and not, 	^ - not, 	| or

func SetBit(value int64, i uint, bit int) (int64, error) {
	i -= 1
	if i > 63 || (bit != 0 && bit != 1) {
		return 0, fmt.Errorf("bit should be 0-63, got: %d", i)
	}
	if bit == 1 {
		return value | (1 << i), nil
	} else {
		return value &^ (1 << i), nil
	}
}

func main() {
	var num int64 = 5
	var position uint = 1

	fmt.Printf("initial value: %d\n", num)
	fmt.Printf("binary: %b\n", num)
	fmt.Printf("set %dth bit to 0\n", position)

	result, err := SetBit(num, position, 0)
	if err != nil {
		fmt.Printf("error: %e", err)
	}

	fmt.Printf("result: %d\n", result)
	fmt.Printf("binary: %b\n", result)
}
