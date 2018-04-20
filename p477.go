package main

import (
	"fmt"
	"math/bits"
)

func totalHammingDistance(nums []int) int {
	sum := 0

	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			result := n ^ nums[j]

			sum += bits.OnesCount(uint(result))
			/*
				for i := result; i > 0; i >>= 1 {
					sum += i & 1
				}*/
			/*
				for i := 0; i < 30; i++ {
					if (result-1)%2 == 0 {
						sum++
					}
					result >>= 1
				}*/
			/*
				if (result-1)%2 == 0 {
					result--
					sum++
				}
				for t := 2; t <= result; t *= 2 {
					if (result ^ t) == (result - t) {
						result -= t
						sum++
					}
				}*/
		}
	}
	return sum
}

func main() {

	a := []int{4, 14, 2}

	fmt.Println(totalHammingDistance(a))
}
