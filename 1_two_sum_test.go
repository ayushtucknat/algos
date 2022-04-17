package main_test

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	sample := []struct {
		data   []int
		target int
		ans    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, data := range sample {
		twoSum(data.data, data.target)
	}
}

func twoSum(nums []int, target int) []int {
	iteratedMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if index, ok := iteratedMap[diff]; ok {
			return []int{index, i}
		}
		iteratedMap[nums[i]] = i
	}
	return []int{}
}
