// Leetcode 132

// DISCLOSURE: my initial solution attempts (in comments below) either were not correct or
// took too long
// The solution given outside the comments here is taken from the leetcode
// discussion board
package main

import "fmt"

func min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}

func minCut(s string) int {
	n := len(s)
	cut := make([]int, n+1)

	for i := 0; i <= n; i++ {
		cut[i] = i - 1
	}

	for i := 0; i < n; i++ {
		for j := 0; i-j >= 0 && i+j < n && s[i-j] == s[i+j]; j++ {
			cut[i+j+1] = min(cut[i+j+1], 1+cut[i-j])
		}

		for j := 1; i-j+1 >= 0 && i+j < n && s[i-j+1] == s[i+j]; j++ {
			cut[i+j+1] = min(cut[i+j+1], 1+cut[i-j+1])
		}
	}
	return cut[n]
}

func main() {
	fmt.Println(minCut("aab")) // 1
	fmt.Println()
	fmt.Println(minCut("abba")) // 0
	fmt.Println()
	fmt.Println(minCut("abbadbbd")) // 1
	fmt.Println()
	fmt.Println(minCut("abbapdbbd")) // 2
	fmt.Println()
	fmt.Println(minCut("efe")) // 0
	fmt.Println()
	fmt.Println(minCut("aaabaa")) // 1
	fmt.Println()
	fmt.Println(minCut("daabaacc")) // 2
	fmt.Println()
	fmt.Println(minCut("daabaace")) // 3
	fmt.Println()
	fmt.Println(minCut("apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp")) // 452
	fmt.Println()
	fmt.Println(minCut("abcdefghijklmnopqrstuvwxyz")) // 25
	fmt.Println()
	fmt.Println(minCut("aafaabbaabbee")) // 2
	fmt.Println()
}

/*
func minCut(s string) int {
	windowSize := len(s)
	checked := make([]bool, len(s))
	totalParts := 0
	numCut := 0
	nonPalPortioni := 0
	nonPalPortionj := len(s)

	for windowSize > 0 && numCut < len(s) {
		//fmt.Printf("windowSize: %d  numCut: %d\n", windowSize, numCut)
		windowParts := []Part{}
		remainder := nonPalPortionj - windowSize
		i := nonPalPortioni
		for i <= remainder {
			lastInd := i + windowSize
			//fmt.Printf("Checking for palindrome: %s\n", string(s[i:lastInd]))
			if isPalindrome(string(s[i:lastInd])) {
				//fmt.Printf("found palidrome: %s  indices: %d, %d\n", string(s[i:lastInd]), i, lastInd)
				windowParts = append(windowParts, Part{i, lastInd})
				numCut += lastInd - i
				for j := i; j < lastInd; j++ {
					checked[j] = true
				}
				totalParts++
				break
			}
			i++
		}
		if i > remainder {
			windowSize--
		} else {
			a, b := findFirst(checked)
			if a < 0 {
				numCut = len(s)
				break
			} else {
				nonPalPortioni = a
				nonPalPortionj = b
			}
			windowSize = b - a + 1
		}
	}
	return totalParts - 1
}


type Part struct {
	start   int
	last    int
	checked []bool
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func findFirst(b []bool) (int, int) { // returns i, j s.t. b[i:j] is the leftmost false sequence
	i := 0
	for i < len(b) {
		if b[i] == false {
			break
		}
		i++
	}

	if i >= len(b) {
		return -1, -1
	}

	j := i
	for j < len(b) {
		if b[j] == true {
			break
		}
		j++
	}
	return i, j
}

func findParts(s string, winSize int, checked []bool) []Part {
	parts := []Part{}
	freeBlocks := []Part{}

	i := 0
	for i < len(s) {
		if checked[i] == false {
			j := i + 1
			for j < len(s) && checked[j] == false {
				j++
			}
			if j-i >= winSize {
				freeBlocks = append(freeBlocks, Part{i, j, nil})
			}
			i = j
		} else {
			i++
		}
	}

	for _, block := range freeBlocks {
		for i := block.start; i < block.last-winSize+1; i++ {
			j := i + winSize
			if isPalindrome(string(s[i:j])) {
				newc := make([]bool, len(checked))
				copy(newc, checked)

				for ind, _ := range newc {
					if ind >= i && ind < j {
						newc[ind] = true
					}
				}
				parts = append(parts, Part{i, j, newc})
			}
		}
	}
	return parts
}

func genChecked(l int, i int, j int) []bool {
	c := make([]bool, l)

	for k := 0; k < l; k++ {
		if k >= i && k < j {
			c[k] = true
		} else {
			c[k] = false
		}
	}
	return c
}

func finish(s string, n Node) Node {
	numCut := n.numCut
	twosFound := 0

	for i := 0; i < len(s)-1; i++ {
		if n.checked[i] == false && n.checked[i+1] == false && s[i] == s[i+1] {
			twosFound++
			numCut += 2
			i++
		}
	}

	onesLeft := len(s) - numCut

	return Node{1, n.checked, len(s), n.totalParts + twosFound + onesLeft}
}

type Node struct {
	winSize int
	//nexti          int
	checked    []bool
	numCut     int
	totalParts int
	//children   []Node
}

func minCut(s string) int {
	fmt.Printf("Input string: %s\n", s)
	queue := []Node{}

	windowSize := len(s)

	blankChecked := make([]bool, len(s))
	toplist := findParts(s, windowSize, blankChecked)
	minParts := len(s)

	for len(toplist) == 0 {
		windowSize--
		toplist = findParts(s, windowSize, blankChecked)
	}

	//fmt.Printf("toplist:\n")
	for _, part := range toplist {
		//fmt.Printf("%s\n", string(s[part.start:part.last]))
		queue = append(queue, Node{windowSize, part.checked, windowSize, 1})
	}

	//fmt.Println(queue)

	for len(queue) > 0 {
		next := queue[0]
		//fmt.Println("next:")
		//fmt.Println(next)
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []Node{}
		}

		if next.numCut == len(s) {
			//fmt.Printf("leaf node totalParts: %d\n", next.totalParts)
			if next.totalParts < minParts {
				minParts = next.totalParts
			}
			continue
		}

		if next.winSize <= 2 {
			queue = append(queue, finish(s, next))
		} else {
			nextList := findParts(s, next.winSize, next.checked)
			wsize := next.winSize
			for len(nextList) == 0 && wsize > 3 {
				wsize--
				nextList = findParts(s, wsize, next.checked)
			}

			if len(nextList) == 0 {
				queue = append(queue, finish(s, next))
			}

			for _, part := range nextList {
				queue = append(queue, Node{wsize, part.checked, next.numCut + wsize, next.totalParts + 1})
			}
		}
	}
	return minParts - 1
}
*/
