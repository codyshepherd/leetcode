//Cody Shepherd
// leetcode challenge 324
package main

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	sort.Ints(nums)
	middle := int(n / 2)
	if n%2 != 0 {
		middle += 1
	}

	fmt.Println("sorted ", nums)
	fmt.Println("middle ", middle)

	var low []int
	var high []int
	if n < 3 {
		low = nums[:1]
		high = nums[1:]
	} else {
		low = nums[:middle]
		high = nums[middle:]
	}

	i := len(low) - 1
	j := len(high) - 1

	lempty := false
	hempty := false

	var wnums []int
	for !lempty && !hempty {
		wnums = append(wnums, low[i])
		wnums = append(wnums, high[j])

		i -= 1
		j -= 1

		if i < 0 {
			lempty = true
		}
		if j < 0 {
			hempty = true
		}
	}
	for i >= 0 {
		wnums = append(wnums, low[i])
		i -= 1
	}

	for j >= 0 {
		wnums = append(wnums, high[j])
		j -= 1
	}
	_ = copy(nums, wnums)
}

func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)

	nums = []int{1}
	wiggleSort(nums)
	fmt.Println(nums)

	nums = []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums)
	fmt.Println(nums)

	nums = []int{5, 3, 1, 2, 6, 7, 8, 5, 5}
	wiggleSort(nums)
	fmt.Println(nums)

	nums = []int{2, 1}
	wiggleSort(nums)
	fmt.Println(nums)

	nums = []int{1, 1, 2, 2, 2, 1}
	wiggleSort(nums)
	fmt.Println(nums)

	nums = []int{1, 1, 2, 1, 2, 2, 1}
	wiggleSort(nums)
	fmt.Println(nums)

	nums = []int{4, 5, 5, 6}
	wiggleSort(nums)
	fmt.Println(nums)

}
