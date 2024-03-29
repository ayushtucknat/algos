package main

import (
	"www.example.com/algos/algo_expert"
)

func main() {
	input := []int{5, 7, 1, 1, 2, 3, 22}
	expected := 20
	actual := algo_expert.NonConstructibleChange(input)
	if actual != expected {
		panic("Expected " + string(expected) + ", got " + string(actual))
	}
}
