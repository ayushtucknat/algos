package main

import "testing"

func Test_SearchPivotedArray(t *testing.T) {
	sampleData := []struct {
		array    []int
		val      int
		expected int
	}{
		{array: []int{4, 5, 6, 7, 1, 2, 3}, val: 5, expected: 1},
		{array: []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5}, val: 5, expected: 9},
		{array: []int{4, 5, 6, 7, 1, 2, 3}, val: 15, expected: -1},
		{array: []int{7, 1}, val: 1, expected: 1},
		{array: []int{7, 1}, val: 7, expected: 0},
		{array: []int{7}, val: 7, expected: 0},
		{array: []int{7}, val: 1, expected: -1},
		{array: []int{}, val: 5, expected: -1},

		// => failed case {array: []int{1, 0, 1, 1, 1}, val: 0, expected: 1},
		{array: []int{4, 5, 6, 7, 0, 1, 2}, val: 0, expected: 4},
		{array: []int{4, 5, 6, 7, 8, 1, 2, 3}, val: 8, expected: 4},
	}

	for _, data := range sampleData {
		actual := searchPivotedArray(data.array, data.val)
		if actual != data.expected {
			t.Errorf("array: %v, val: %d, expected: %d, actual: %v", data.array, data.val, data.expected, actual)
		}
	}
}

func searchPivotedArray(nums []int, val int) int {

	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == val {
			return mid
		}

		// left sorted portion
		if nums[l] <= nums[mid] {
			if val > nums[mid] || val < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		} else { // right sorted portion
			if val > nums[r] || val < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

	}

	return -1
}
