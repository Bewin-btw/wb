package main

import (
	"fmt"
	"sort"
)

func Grouping(nums []float64) [][]float64 {
	if len(nums) < 1 {
		return nil
	}
	sort.Float64s(nums)
	res := [][]float64{}
	current := int(nums[0]) / 10
	term := []float64{nums[0]}
	for i := 1; i < len(nums); i++ {
		if int(nums[i])/10 == current {
			term = append(term, nums[i])
			continue
		}
		res = append(res, term)
		term = []float64{nums[i]}
		current = int(nums[i]) / 10
	}
	res = append(res, term)
	return res
}

func MapGrouping(nums []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, t := range nums {
		key := int(t) / 10 * 10
		groups[key] = append(groups[key], t)
	}

	return groups
}

func main() {
	init := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(Grouping(init))
	fmt.Println(MapGrouping(init))
}
