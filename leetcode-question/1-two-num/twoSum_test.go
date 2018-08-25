package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if fmt.Sprint(twoSum([]int{1, 2}, 3)) != "[0 1]" {
		t.Error(`twoSum([]int{1, 2}, 3)) = [1 2]`)
	}
	if fmt.Sprint(twoSum([]int{-21, 12, 3, 5, 7, -3}, 4)) != `[4 5]` {
		t.Error(`twoSum([]int{-21, 12, 3, 5, 7, -3}, 4)) = [4 5]`)
	}
}
