package main

import (
	"sort"
	"testing"
)

func Test_twoSumSorted(t *testing.T) {

	t.Errorf("target: %d, sol: %v\n", 1, two_sum([]int{-5, -5, -3, -2, -1, 0, 0, 1, 3, 4, 5, 6, 6, 7, 10, 12, 13, 15}, 1))
	t.Errorf("target: %d, sol: %v\n", 0, two_sum([]int{-5, -5, -3, -2, -1, 0, 0, 1, 3, 4, 5, 6, 6, 7, 10, 12, 13, 15}, 0))
	t.Errorf("target: %d, sol: %v\n", -1, two_sum([]int{-5, -5, -3, -2, -1, 0, 0, 1, 3, 4, 5, 6, 6, 7, 10, 12, 13, 15}, -1))
	t.Errorf("target: %d, sol: %v\n", 10, two_sum([]int{-5, -5, -3, -2, -1, 0, 0, 1, 3, 4, 5, 6, 6, 7, 10, 12, 13, 15}, 10))
	t.Errorf("target: %d, sol: %v\n", 0, two_sum([]int{0}, 0))
}

func two_sum(nums []int, k int) [][]int {
	//sort the array
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	solutions := [][]int{}
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		// too big
		if sum > k {
			j = j - 1
		} else if sum < k { // to smalll
			i = i + 1
		} else { //found the sum
			solutions = append(solutions, []int{nums[i], nums[j]})
			i = i + 1
			for ; i < j; i = i + 1 {
				if nums[i] != nums[i-1] {
					break
				}
			}
		}
	}
	return solutions
}

// [-5, -3, -2, -1, 0, 1, 3, 4, 5, 6, 6, 7, 10, 12, 13, 15]
