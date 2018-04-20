package main

import (
	"fmt"
)

func next(nums []int, i int, j int, k int) ([]int, int, int, int) {
	l := len(nums)
    var lst = []int{0,0,0}
	if i == l-3 && k == l {
		return nil, 0, 0, 0
	} else if j == l-2 && k == l {
		//lst = append(lst, nums[i+1])
		//lst = append(lst, nums[i+2])
		//lst = append(lst, nums[i+3])
        return nums[i+1:i+4], i + 1, i + 2, i + 4
	} else if k == l {
		lst[0] = nums[i]
		lst[1] = nums[j+1]
		lst[2] = nums[j+2]
		return lst, i, j + 1, j + 3
	} else {
		lst[0] = nums[i]
		lst[1] = nums[j]
		lst[2] = nums[k]
		return lst, i, j, k + 1
	}
}

func sum(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func isIn(target int, nums []int) bool {
	for _, num := range nums {
		if target == num {
			return true
		}
	}
	return false
}

func check(nums []int, lst [][]int) bool {

	for _, trip := range lst {
		all := true
		for _, num := range nums {
			if !isIn(num, trip) {
				all = false
			}
		}
		if all {
			return false
		}
	}

	return true //true == no duplicate found
}

func allZeros(nums []int) bool {
	for _, num := range nums {
		if num != 0{
			return false
		}
	}
	return true
}

func threeSum(nums []int) [][]int {
	var ll [][]int
    if len(nums) < 3 {
		return ll
	}
	if allZeros(nums) {
		ll = append(ll, []int{0,0,0})
        return ll
	}
	zeros := false
    seen := make(map[]bool)

	trip, i, j, k := next(nums, 0, 1, 2)
	for trip != nil {
		if trip[0] == 0 && trip[1] == 0 && trip[2] == 0 {
			if !zeros {
				ll = append(ll, trip)
				zeros = true
			}
			continue
		}
		if sum(trip) == 0 && check(trip, ll) {
			ll = append(ll, trip)
		}
		trip, i, j, k = next(nums, i, j, k)
	}
	return ll
}

func main() {

	nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	//Expected output: [[-4,-2,6],[-4,0,4],[-4,1,3],[-4,2,2],[-2,-2,4],[-2,0,2]]
	ret := threeSum(nums)
	fmt.Println(ret)

}
