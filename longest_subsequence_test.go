package main_test

import (
	"math"
	"testing"
)

func Test_longestSubsequence(t *testing.T) {
	sample := []struct {
		data   []int
		answer int
	}{
		{data: []int{1, 2, 4, 3}, answer: 3},
		{data: []int{10, 9, 2, 5, 3, 7, 101, 18}, answer: 4},
		{data: []int{0, 1, 0, 3, 2, 3}, answer: 4},
		{data: []int{7, 7, 7, 7, 7, 7, 7}, answer: 1},
	}
	for _, set := range sample {
		actual := longestSubsequence(set.data)
		if actual != set.answer {
			t.Errorf("data: %v, expected: %d, actual: %d", set.data, set.answer, actual)
		}
	}
}

func longestSubsequence(data []int) int {
	size := len(data)
	if size == 0 {
		return 0
	}

	mem := make([]float64, size)
	mem[size-1] = 1.0

	maxLen := 1

	for j := size - 2; j >= 0; j-- {
		max := 1.0
		for i := j + 1; i < size; i++ {
			if data[i] > data[j] {
				max = math.Max(max, mem[i]+1)
			}
		}
		mem[j] = max
		maxLen = int(math.Max(max, float64(maxLen)))
	}

	return maxLen
}
