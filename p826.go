// Cody Shepherd
// Leetcode problem 826

package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	diff   int
	profit int
}

type Pairs []Pair

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p Pairs) Less(i, j int) bool {
	if p[i].diff < p[j].diff {
		return true
	}
	return false
}

func bestAss(worker int, p Pairs) int {
	maxprofit := 0
	for _, val := range p {
		if val.diff > worker {
			return maxprofit
		}
		if val.profit > maxprofit {
			maxprofit = val.profit
		}
	}
	return maxprofit
}

func sum(arr []int) int {
	s := 0
	for _, val := range arr {
		s += val
	}
	return s
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	table := make(map[int]int) // map workers to max profit

	arr := make([]Pair, len(difficulty))

	for i, _ := range difficulty {
		arr[i] = Pair{difficulty[i], profit[i]}
	}

	sort.Sort(Pairs(arr))

	final := make([]int, len(worker))

	for i, val := range worker {
		prof, ok := table[val]
		if ok {
			final[i] = prof
		} else {
			final[i] = bestAss(val, arr)
			table[val] = final[i]
		}
	}

	return sum(final)
}

func main() {
	difficulty := []int{2, 4, 6, 8, 10}
	profit := []int{10, 20, 30, 40, 50}
	worker := []int{4, 5, 6, 7}
	fmt.Println(maxProfitAssignment(difficulty, profit, worker))

	difficulty = []int{5, 50, 92, 21, 24, 70, 17, 63, 30, 53}
	profit = []int{68, 100, 3, 99, 56, 43, 26, 93, 55, 25}
	worker = []int{96, 3, 55, 30, 11, 58, 68, 36, 26, 1}
	fmt.Println(maxProfitAssignment(difficulty, profit, worker))
}
