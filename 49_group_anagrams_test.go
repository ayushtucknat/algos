package main_test

import (
	"fmt"
	"sort"
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
		actual := groupAnagramFastest(example.sample)
		// FixMe: asserstion are not correct
		for _, result := range example.expected {
			if !contains(actual, result) {
				t.Fail()
			}
		}

	}

}

func contains(actual [][]string, result []string) bool {
	for _, data := range actual {
		matched := true
		for _, elem := range data {
			if !containsInArray(result, elem) {
				matched = false
				break

			}
		}
		if matched {
			return matched
		}
	}
	return false
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

func groupAnagramLinear(strs []string) [][]string {
	scrollIndex := make(map[string]int)
	anagrams := make([][]string, 0)
	len := 0
	for _, word := range strs {
		sorted := sortWord(word)
		if index, ok := scrollIndex[sorted]; ok {
			anagrams[index] = append(anagrams[index], word)
		} else {
			scrollIndex[sorted] = len
			anagrams = append(anagrams, []string{word})
			len++
		}

	}
	return anagrams
}

func groupAnagramFastest(strs []string) [][]string {
	anagrams := make(map[string][]string, 0)
	for _, word := range strs {
		occurances := make([]int, 26)
		for _, char := range []rune(word) {
			occurances[char-'a'] += 1
		}
		hash := fmt.Sprint(occurances)
		if arr, ok := anagrams[hash]; ok {
			anagrams[hash] = append(arr, word)

		} else {
			anagrams[hash] = []string{word}
		}
	}

	result := make([][]string, 0)

	for _, value := range anagrams {
		result = append(result, value)
	}
	return result
}

func sortWord(word string) string {
	chars := []rune(word)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func containsInArray(array []string, data string) bool {
	for _, element := range array {
		if element == data {
			return true
		}
	}
	return false
}
