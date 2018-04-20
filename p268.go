package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
	for _, num := range nums {
		sum -= num
	}
	return sum
}

func main() {
	a := []int{3, 0, 1}
	//a_ans := 2
	b := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	//b_ans := 8

	fmt.Println(missingNumber(a))
	fmt.Println(missingNumber(b))
}
