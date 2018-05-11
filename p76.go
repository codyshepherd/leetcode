package main

import (
	"fmt"
	"math"
)

func tabulate(s string) map[rune]int {
	table := make(map[rune]int)
	for _, letter := range s {
		table[letter]++
	}
	return table
}

func minWindow(s string, t string) string {
	table := tabulate(t)
	l := len(t)
	found := 0
	catchup := 0
	low := -1
	high := len(s)
	dist := math.MaxInt64

	for i, val := range s {
		seen := table[val]
		if seen > 0 {
			//fmt.Println("Found ", string(val))
			//fmt.Println(s[catchup : i+1])
			found++
		}
		table[val]--

		for found == l {
			//fmt.Println("All found")
			//fmt.Println(s[catchup : i+1])
			newdist := i + 1 - catchup
			if newdist < dist {
				dist = newdist
				low = catchup
				high = i
			}

			lowletter := rune(s[catchup])
			//fmt.Println("lowletter ", string(lowletter))
			table[lowletter]++
			if table[lowletter] > 0 {
				found--
			}
			catchup++
		}
	}

	if low < 0 {
		return ""
	}

	//fmt.Println("low ", low, "high ", high)
	return s[low : high+1]
}

func main() {
	input := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(input, t))

	input = "a"
	t = "aa"
	fmt.Println(minWindow(input, t))

	input = "aab"
	t = "aab"
	fmt.Println(minWindow(input, t))

	input = "cabwefgewcwaefgcf"
	t = "cae"
	fmt.Println(minWindow(input, t))

	input = "aa"
	t = "aa"
	fmt.Println(minWindow(input, t))

	input = "bbaa"
	t = "aba"
	fmt.Println(minWindow(input, t))

}
