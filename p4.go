package main

/*
Full Disclosure: I had help from the leetcode discussion on this problem, because
I couldn't figure out a solution after 15+ hours.
*/

import (
	"fmt"
)

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func med(n1 int, n2 int) float64 {
	return float64(n1) + (float64(n2)-float64(n1))/2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		if len(nums2) == 1 {
			return float64(nums2[0])
		} else if len(nums2)%2 == 0 {
			mid := len(nums2) / 2
			return (float64(nums2[mid]) + float64(nums2[mid-1])) / 2.0
		} else {
			mid := len(nums2) / 2
			return float64(nums2[mid])
		}
	}
	if len(nums2) == 0 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		} else if len(nums1)%2 == 0 {
			mid := len(nums1) / 2
			return (float64(nums1[mid]) + float64(nums1[mid-1])) / 2.0
		} else {
			mid := len(nums1) / 2
			return float64(nums1[mid])
		}
	}

	if len(nums1) > len(nums2) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}

	tlen := len(nums1) + len(nums2)

	min_index := 0
	max_index := len(nums1)

	var p1 int
	var p2 int

	for {
		p1 = (min_index + max_index) / 2
		p2 = ((tlen + 1) / 2) - p1
		//fmt.Println("p1: ", p1)
		//fmt.Println("p2: ", p2)
		if p1 < len(nums1) && nums2[p2-1] > nums1[p1] {
			min_index = p1 + 1
		} else if p1 > 0 && nums1[p1-1] > nums2[p2] {
			max_index = p1 - 1
		} else {
			var lo int
			var hi int
			if p1 == 0 {
				lo = nums2[p2-1]
			} else if p2 == 0 {
				lo = nums1[p1-1]
			} else {
				lo = max(nums1[p1-1], nums2[p2-1])
			}
			if tlen%2 != 0 {
				return float64(lo)
			} else {
				if p1 == len(nums1) {
					hi = nums2[p2]
				} else if p2 == len(nums2) {
					hi = nums1[p1]
				} else {
					hi = min(nums1[p1], nums2[p2])
				}
				return med(lo, hi)
			}
		}
	}
}

func main() {
	n1 := []int{1, 2}
	n2 := []int{3, 4}
	fmt.Println(n1, "\n", n2, ":")
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1, 3}
	n2 = []int{2}
	fmt.Println(n1, "\n", n2, ":")
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{}
	n2 = []int{1}
	fmt.Println(n1, "\n", n2, ":")
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{}
	n2 = []int{2, 3}
	fmt.Println(n1, "\n", n2, ":")
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1, 1}
	n2 = []int{1, 2}
	fmt.Println(n1, "\n", n2, ":")
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1, 2, 3, 5, 6}
	n2 = []int{4}
	fmt.Println(n1, "\n", n2, ":")
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{2, 3}
	n2 = []int{1}
	fmt.Println(n1, "\n", n2, ":")
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1, 2, 3, 4, 7, 8}
	n2 = []int{5, 6}
	fmt.Println(n1, "\n", n2, ":")
	fmt.Println(findMedianSortedArrays(n1, n2))

}
