package main_test

import "testing"

func Test_singleNumber(t *testing.T) {
	if singleNumber([]int{2, 2, 1}) != 1 {
		t.Fail()
	}
	if singleNumber([]int{4,1,2,1,2}) != 4 {
		t.Fail()
	}
	if singleNumber([]int{1}) != 1 {
		t.Fail()
	}
}

func singleNumber(nums []int) int {
	mismatch := nums[0]
	for i := 1; i < len(nums); i++ {
		mismatch = mismatch ^ nums[i]
	}
	return mismatch
}
