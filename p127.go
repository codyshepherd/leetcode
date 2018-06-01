package main

import "fmt"

type Node struct {
	val       string
	neighbors []*Node
}

func newNode(s string) *Node {
	return &Node{
		val: s,
	}
}

func dist(a string, b string) int {
	if len(a) != len(b) {
		return -1
	}

	total := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			total++
		}
	}
	return total
}

func buildGraph(wordList []string) *Node {
	seen := make(map[string]*Node)
	var root *Node
	var n *Node
	var prev *Node
	for i, word := range wordList {
		prev = seen[word]
		if prev == nil {
			prev = newNode(word)
			seen[word] = prev
		}
		if i == 0 {
			root = prev
		}
		for j := i + 1; j < len(wordList); j++ {
			if dist(word, wordList[j]) == 1 {
				n = seen[wordList[j]]
				if n == nil {
					n = newNode(wordList[j])
					seen[wordList[j]] = n
				}
				prev.neighbors = append(prev.neighbors, n)
				n.neighbors = append(n.neighbors, prev)
			}
		}
	}
	return root
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	graphRoot := buildGraph(append([]string{beginWord}, wordList...))

	var q []*Node
	var dists []int
	visited := make(map[string]bool)

	length := 1

	q = append(q, graphRoot)
	dists = append(dists, 0)
	visited[graphRoot.val] = true
	var node *Node
	var l int
	var lastLen int

	for len(q) != 0 {
		node = q[0]
		l = dists[0]
		if l > lastLen {
			length++
		}
		fmt.Println("Dequeueing ", node.val, "length ", l)
		fmt.Println("length: ", length)
		q = q[1:]
		dists = dists[1:]

		for _, neighbor := range node.neighbors {
			if !visited[neighbor.val] {
				q = append(q, neighbor)
				dists = append(dists, l+1)
				visited[neighbor.val] = true
			}
		}
		if node.val == endWord {
			return length
		}
		lastLen = l
	}
	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("a", "c", []string{"a", "b", "c"}))
}
