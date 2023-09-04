package algo_expert_test

import (
	"testing"
	"www.example.com/algos/algo_expert"
)

func Test_coinChange(t *testing.T) {
	input := []int{5, 7, 1, 1, 2, 3, 22}
	expected := 20
	actual := algo_expert.NonConstructibleChange(input)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
