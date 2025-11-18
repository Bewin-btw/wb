package main

import "fmt"

func Intersection(a []int, b []int) []int {
	m := make(map[int]bool)
	res := []int{}

	for _, v := range a {
		m[v] = true
	}

	for _, v := range b {
		if m[v] {
			res = append(res, v)
			m[v] = false
		}
	}
	return res
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	fmt.Println(Intersection(A, B))
}
