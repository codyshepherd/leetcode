package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	numrows := len(matrix)

	for i := 0; i < numrows; i++ {
		if i < numrows-1 {
			nextstart := matrix[i+1][0]
			if target < nextstart {
				for _, num := range matrix[i] {
					if num == target {
						return true
					}
				}
			}
		} else {
			for _, num := range matrix[i] {
				if num == target {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	var test = [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 3

	fmt.Println(searchMatrix(test, target))
}
