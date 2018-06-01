// Cody Shepherd
// Leetcode problem 838

package main

import "fmt"

func pushDominoes(dominoes string) string {
	var newd []rune
	n := len(dominoes)
	if n == 1 {
		return dominoes
	}
	changed := false
	for i, letter := range dominoes {
		if letter != '.' {
			//newd[i] = letter
			newd = append(newd, letter)
		} else {
			if i == 0 {
				if dominoes[i+1] == 'L' {
					//newd[i] = 'L'
					newd = append(newd, 'L')
					changed = true
				} else {
					//newd[i] = '.'
					newd = append(newd, '.')
				}
			} else if i == n-1 {
				if dominoes[i-1] == 'R' {
					//newd[i] = 'R'
					newd = append(newd, 'R')
					changed = true
				} else {
					//newd[i] = '.'
					newd = append(newd, '.')
				}
			} else {
				if dominoes[i-1] == 'R' && dominoes[i+1] == 'L' {
					//newd[i] = '.'
					newd = append(newd, '.')
				} else if dominoes[i-1] == 'R' {
					//newd[i] = 'R'
					newd = append(newd, 'R')
					changed = true
				} else if dominoes[i+1] == 'L' {
					//newd[i] == 'L'
					newd = append(newd, 'L')
					changed = true
				} else {
					//newd[i] = '.'
					newd = append(newd, '.')
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
