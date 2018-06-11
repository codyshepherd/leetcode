// Cody Shepherd
// Leetcode problem 838

package main

import "fmt"

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	if n == 1 {
		return dominoes
	}
	newd := make([]rune, n)

	table := make(map[rune][]int)

	for i, val := range dominoes {
		if val == 'L' {
			table['L'] = append(table['L'], i)
		}
		if val == 'R' {
			table['R'] = append(table['R'], i)
		}
	}

	var nnl int
	l := table['L']
	r := table['R']
	var nl int
	var nr int
	i := 0
	counter := 0
	for i < n {
		if len(l) > 0 {
			nl = l[0]
			l = l[1:]
		} else {
			nl = n - 1
		}
		if len(r) > 0 {
			nr = r[0]
			r = r[1:]
		} else {
			nr = n - 1
		}
		for i < nl && i < nr {
			newd[i] = 'L'
			i += 1
		}
		newd[i] = 'L'
		i += 1
		if i < n {
			newd[i] = '.'
			i += 1
		}
		if len(l) > 0 {
			nnl = l[0]
		} else {
			nnl = n - 1
		}
		if (nnl+1-nr)%2 == 0 {
			for i > nr && i < (nnl-nr)/2+nr {
				newd[i] = 'R'
				i += 1
			}
		} else {
			counter = ((nnl - 1) - nr) / 2
			for counter > 0 {
				newd[i] = 'R'
				i += 1
				counter -= 1
			}
			newd[i] = '.'
		}
	}
	return string(newd)
}

func pushDominoes1(dominoes string) string {
	//var newd []rune
	n := len(dominoes)
	if n == 1 {
		return dominoes
	}
	newd := make([]rune, n)
	changed := false
	for i, letter := range dominoes {
		if letter != '.' {
			newd[i] = letter
			//newd = append(newd, letter)
		} else {
			if i == 0 {
				if dominoes[i+1] == 'L' {
					newd[i] = 'L'
					//newd = append(newd, 'L')
					changed = true
				} else {
					newd[i] = '.'
					//newd = append(newd, '.')
				}
			} else if i == n-1 {
				if dominoes[i-1] == 'R' {
					newd[i] = 'R'
					//newd = append(newd, 'R')
					changed = true
				} else {
					newd[i] = '.'
					//newd = append(newd, '.')
				}
			} else {
				if dominoes[i-1] == 'R' && dominoes[i+1] == 'L' {
					newd[i] = '.'
					//newd = append(newd, '.')
				} else if dominoes[i-1] == 'R' {
					newd[i] = 'R'
					//newd = append(newd, 'R')
					changed = true
				} else if dominoes[i+1] == 'L' {
					newd[i] = 'L'
					//newd = append(newd, 'L')
					changed = true
				} else {
					newd[i] = '.'
					//newd = append(newd, '.')
				}
			}
		}
	}
	if !changed {
		return dominoes
	}
	return pushDominoes(string(newd))
}

func main() {
	st := ".L.R...LR..L.."
	fmt.Println(pushDominoes(st))
}
