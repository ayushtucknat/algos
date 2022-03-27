package main

import (
	"strconv"
	"testing"
)

type data struct {
	num string
	ans string
	k   int
}

var testData = []data{
	{num: "1432219", k: 3, ans: "1219"},
	{num: "10200", k: 1, ans: "200"},
	{num: "999", k: 1, ans: "99"},
	{num: "1999", k: 2, ans: "19"},
	{num: "9111", k: 2, ans: "11"},
	{num: "10991", k: 3, ans: "1"},
}

func Test_RemoveKDigits(t *testing.T) {
	for _, val := range testData {
		actual := removeKDigits(val.num, val.k)
		if actual != val.ans {
			t.Errorf("num: %s, k: %d, expected: %s, actual: %s", val.num, val.k, val.ans, actual)
		}
		t.Logf("{%s, %d, %s}", val.num, val.k, actual)
	}
}

func removeKDigits(s string, k int) string {
	stack := []rune{}
	for _, num := range []rune(s) {
		for ; k > 0 && len(stack) > 0 && stack[len(stack)-1] > num; k = k - 1 {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, num)
	}

	stack = stack[:len(stack)-k]
	num, _ := strconv.Atoi(string(stack))
	return strconv.Itoa(num)

}
