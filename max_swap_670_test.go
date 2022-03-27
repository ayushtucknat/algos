package main

import (
	"strconv"
	"testing"
)

var sampleMaxSwap = []struct {
	num      int
	expected int
}{
	{num: 2736, expected: 7236},
	{num: 2376, expected: 7326},
	{num: 9973, expected: 9973},
	{num: 99307, expected: 99703},
	{num: 1993, expected: 9913},
}

func Test_MaxSwap(t *testing.T) {
	for _, data := range sampleMaxSwap {
		actual := maxSwap(data.num)
		if actual != data.expected {
			t.Errorf("data: %d, expected: %d, actual: %d", data.num, data.expected, actual)
		}

	}
}

func maxSwap(num int) int {
	data := []rune(strconv.Itoa(num))
	swap := 0
	for i := 0; i < len(data)-1; i++ {
		if data[i] < data[i+1] {
			swap = i
			break
		}
	}
	max := swap
	for i := swap + 1; i < len(data); i++ {
		if data[i] >= data[max] {
			max = i
		}
	}

	for i := 0; i <= swap; i++ {
		if data[i] < data[max] {
			data[i], data[max] = data[max], data[i]
			break
		}
	}

	num, _ = strconv.Atoi(string(data))
	return num
}
