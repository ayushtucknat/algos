package main

import "fmt"

func main() {
	fmt.Printf("%v\n", longest_subsequence_length([]int{2, 100, 200, 3, 4, 1}))                                        //4
	fmt.Printf("%v\n", longest_subsequence_length([]int{0, 2, 100, 200, 3, 4, 1}))                                     //5
	fmt.Printf("%v\n", longest_subsequence_length([]int{2, -1, 100, 200, 3, 4, 1}))                                    //4
	fmt.Printf("%v\n", longest_subsequence_length([]int{2, 100, 101, 102, 103, 104, 105, 106, 200, 3, 4, 1, 5, 6, 7})) //7
	fmt.Printf("%v\n", longest_subsequence_length([]int{}))                                                            //0
	fmt.Printf("%v\n", longest_subsequence_length([]int{1}))                                                           //1
	fmt.Printf("%v\n", longest_subsequence_length([]int{1, 2}))                                                        //2

}

func longest_subsequence_length(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	set := createSet(nums)
	longest := 0
	for _, num := range nums {
		// check if first number of sequence
		if _, ok := set[num-1]; !ok {
			count := 1
			for i := num + 1; ; i, count = i+1, count+1 {
				_, ok = set[i]
				if !ok {
					if count > longest {
						longest = count
					}
					break
				}
			}
		}
	}
	return longest
}

func createSet(nums []int) map[int]struct{} {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	return set
}
