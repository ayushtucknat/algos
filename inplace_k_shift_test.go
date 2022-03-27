package main

import (
	"errors"
	"fmt"
	"testing"
)

func Test_inplace_k_shift(t *testing.T) {
	data := []int{1, 2, 3, 5, 6, 7, 8, 10}
	// 7, 8, 10, 1, 2, 3, 5, 6
	data, err := shift(data, -1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v\n", data)
}

func shift(data []int, k int) ([]int, error) {
	size := len(data)

	k = k % size
	if k < 0 {
		return nil, errors.New("stupid k")
	}
	if k == 0 {
		return data, nil
	}
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	for i, j := k, size-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	return data, nil
}
