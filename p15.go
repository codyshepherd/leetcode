package main

import "fmt"

func next(nums []int, i int, j int, k int) ([]int, int, int, int) {
	l := len(nums)
	var lst []int
	if i == l-3 && k == l {
		return nil, 0, 0, 0
	} else if j == l-2 && k == l {
		lst = append(lst, nums[i+1])
		lst = append(lst, nums[i+2])
		lst = append(lst, nums[i+3])
		return lst, i + 1, i + 2, i + 4
	} else if k == l {
		lst = append(lst, nums[i])
		lst = append(lst, nums[j+1])
		lst = append(lst, nums[j+2])
		return lst, i, j + 1, j + 3
	} else {
		lst = append(lst, nums[i])
		lst = append(lst, nums[j])
		lst = append(lst, nums[k])
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

func check(nums []int, lst [][]int) bool {
	var bits int = 0
	for _, num := range nums {
		bits = bits ^ num
	}

	var newbits int = 0
	for _, trip := range lst {
		newbits = 0
		for _, t := range trip {
			newbits = newbits ^ t
		}
		if bits == newbits {
			return false
		}
	}

	//fmt.Println("Not Found")
	return true //true == no duplicate found
}

func threeSum(nums []int) [][]int {
	var ll [][]int
	if len(nums) < 3 {
		return ll
	}
	trip, i, j, k := next(nums, 0, 1, 2)
	for trip != nil {
		//fmt.Println(trip, sum(trip))
		if sum(trip) == 0 && check(trip, ll) {
			ll = append(ll, trip)
		}
		trip, i, j, k = next(nums, i, j, k)
	}
	return ll
}

func main() {

	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	//Expected output: [[-4,-2,6],[-4,0,4],[-4,1,3],[-4,2,2],[-2,-2,4],[-2,0,2]]
	ret := threeSum(nums)
	fmt.Println(ret)

}
