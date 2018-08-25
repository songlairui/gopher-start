package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	pool := make(map[int]int)
	for i, v := range nums {
		candy := target - v
		pairIdx, ok := pool[candy]
		if ok {
			return []int{pairIdx, i}
		} else {
			pool[v] = i
		}
	}
	return []int{}
}

func main() {
	fmt.Println("[0 1]" == fmt.Sprint(twoSum([]int{1, 2}, 3)))
	fmt.Println(twoSum([]int{-21, 12, 3, 5, 7, -3}, 4))
	fmt.Println(fmt.Sprint(twoSum([]int{-21, 12, 3, 5, 7, -3}, 4)) == `[4 5]`)
}
