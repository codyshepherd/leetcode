package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	//fmt.Println(nums)

	var dists = make([]int, nums[len(nums)-1]+1)
	var val int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			val = abs(nums[j] - nums[i])
			//fmt.Println("val: ", val)
			dists[val]++
			//fmt.Printf("i: %d  j: %d  nums[j]-nums[i]: %d\n", i, j, val)
		}
	}
	//fmt.Println(dists)
	i := 0
	var to_sub int
	for k != 0 {
		for dists[i] == 0 {
			i++
		}
		to_sub = min(dists[i], k)
		k -= to_sub
		dists[i] -= to_sub
	}
	return i
}

func main() {
	nums := []int{1, 3, 1}
	k := 1
	fmt.Printf("answer: %d\n", smallestDistancePair(nums, k))

	nums = []int{62, 100, 4}
	k = 2
	fmt.Printf("answer: %d\n", smallestDistancePair(nums, k))

	nums = []int{1, 6, 1}
	k = 3
	fmt.Printf("answer: %d\n", smallestDistancePair(nums, k))

	nums = []int{2, 2, 0, 1, 1, 0, 0, 1, 2, 0}
	k = 2
	fmt.Printf("answer: %d\n", smallestDistancePair(nums, k))
}
