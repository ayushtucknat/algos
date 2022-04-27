package main_test

import (
	"strings"
	"testing"
)

func Test_groupAnagram(t *testing.T) {
	data := []struct {
		sample   []string
		expected [][]string
	}{
		{
			sample:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			sample:   []string{"", ""},
			expected: [][]string{{"", ""}},
		},
		{
			sample:   []string{"", "b"},
			expected: [][]string{{"b"}, {""}},
		},
		{
			sample:   []string{"eat", "tea", "tan", "ate", "nat", "bat", "ac", "bd", "aac", "bbd", "aacc", "bbdd", "acc", "bdd"},
			expected: [][]string{{"bdd"}, {"bat"}, {"nat", "tan"}, {"ac"}, {"ate", "eat", "tea"}, {"bd"}, {"aac"}, {"bbd"}, {"aacc"}, {"bbdd"}, {"acc"}},
		},
	}

	for _, example := range data {
		actual := groupAnagram(example.sample)
		for i := range actual {

			for j := range actual[i] {
				if !contains(actual[i], example.expected[i][j]) {
					t.Fail()
				}
			}

		}

	}

}

func groupAnagram(strs []string) [][]string {
	data := make(map[int]bool, 0)
	results := [][]string{}
	for i := len(strs) - 1; i >= 0; i-- {
		if _, ok := data[i]; ok {
			continue
		}
		data[i] = true
		anagrams := []string{strs[i]}
		for j := i - 1; j >= 0; j-- {
			if strs[j] == "" && strs[i] == "" {
				anagrams = append(anagrams, strs[j])
				data[j] = true
				continue
			}
			if _, ok := data[j]; ok || strs[j] == "" || len(strs[j]) != len(strs[i]) {
				continue
			}
			word := strs[j]
			for _, char := range []rune(strs[i]) {
				word = strings.Replace(word, string(char), "", 1)
			}
			if word == "" {
				data[j] = true
				anagrams = append(anagrams, strs[j])
			}
		}
		results = append(results, anagrams)
	}

	return results
}

func contains(array []string, data string) bool {
	for _, element := range array {
		if element == data {
			return true
		}
	}
	return false
}
