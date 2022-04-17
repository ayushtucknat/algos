package main_test

import (
	"sort"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	if isAnagram("anagram", "nagaram") != true {
		t.Fail()
	}
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	source := []rune(s)
	sort.SliceStable(source, func(i, j int) bool {
		return source[i] < source[j]
	})

	target := []rune(t)
	sort.SliceStable(target, func(i, j int) bool {
		return target[i] < target[j]
	})

	for index, char := range source {
		if target[index] != char {
			return false
		}
	}

	return true
}
