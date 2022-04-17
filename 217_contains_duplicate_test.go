package main_test

import "testing"

func Test_containsDuplicate(t *testing.T) {
	if containsDuplicate([]int{1, 2, 3, 1}) != true {
		t.Fail()
	}
}

func containsDuplicate(nums []int) bool {
	data := make(map[int]bool, 0)
	for _, num := range nums {
		if _, ok := data[num]; ok {
			return true
		}
		data[num] = true
	}
	return false
}
