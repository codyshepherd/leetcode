/* Cody Shepherd
Full Disclosure: I worked on this for 10+ hours and was not able to
get the solution, so I imitated one from the leetcode discussion board
*/

package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {

	i := 0 // pattern
	j := 0 // string
	star := -1
	ss := 0

	for j < len(s) {
		if i < len(p) && (p[i] == '?' || p[i] == s[j]) {
			i++
			j++
		} else if i < len(p) && p[i] == '*' {
			star = i
			i++
			ss = j
		} else if star >= 0 {
			i = star + 1
			ss++
			j = ss
		} else {
			return false
		}
	}
	for i < len(p) && p[i] == '*' {
		i++
	}

	if i >= len(p) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isMatch("aa", "a")) // false
	fmt.Println()
	fmt.Println(isMatch("aa", "*")) // true
	fmt.Println()
	fmt.Println(isMatch("cb", "?a")) // false
	fmt.Println()
	fmt.Println(isMatch("adceb", "*a*b")) // true
	fmt.Println()
	fmt.Println(isMatch("acdcb", "a*c?b")) // false
	fmt.Println()
	fmt.Println(isMatch("abefcdgiescdfimde", "ab*cd?i*de")) // true
	fmt.Println()
	fmt.Println(isMatch("aaaa", "***a")) // true
	fmt.Println()
	fmt.Println(isMatch("aaba", "?***")) // true
	fmt.Println()
	fmt.Println(isMatch("ab", "*?*?*")) // true
	fmt.Println()
	fmt.Println(isMatch("hi", "*?")) // true
	fmt.Println()
	fmt.Println(isMatch("ho", "**ho")) // true
	fmt.Println()
	fmt.Println(isMatch("mississippi", "m*iss*")) // true
	fmt.Println()
	fmt.Println(isMatch("mississippi", "m*iss*iss*")) // true
	fmt.Println()
	fmt.Println(isMatch("a", "aa")) // false
	fmt.Println()
	fmt.Println(isMatch("abbaaa", "*?aa?a")) // false
	fmt.Println()
	fmt.Println(isMatch("abbaaa", "*?aa*a")) // true
	fmt.Println()
	fmt.Println(isMatch("mississippi", "m??*ss*?i*pi")) // false
	fmt.Println()
	fmt.Println(isMatch("mississippi", "m*i*si*si*pi")) // true
	fmt.Println()
}

/*
func lookBackwards(s string, litSeqTok Token, start int, numLeft int, atLeast int) int {
	//fmt.Println("lookBackwards listArr arg:")
	//fmt.Println(litSeqTok.LitSeq)
	lenSeq := len(litSeqTok.LitSeq)
	seqsFound := 0
	charsChecked := 0
	listStartChar := litSeqTok.LitSeq[0]
	i := len(s) - 1
	for i >= start {
		charsChecked++
		if s[i] == listStartChar {
			/*
				j := i - 1
				for j >= start && s[j] == listStartChar {
					j--
				}
				i = j + 1
			//fmt.Printf("lookBackwards found startchar %c\n", s[i])
			if (len(s))-i < len(litSeqTok.LitSeq) || charsChecked < atLeast {
				i--
				continue
			} else if string(s[i:i+lenSeq]) == litSeqTok.LitSeq {
				seqsFound++
				//fmt.Printf("lookBackwards found sequence %s   number left to find: %d\n", litSeqTok.LitSeq, numLeft-seqsFound)
			}

			if seqsFound == numLeft {
				return i
			}
		}
		i--
	}
	return -1
}

func isMatch(s string, p string) bool {

	tseq := []Token{}
	toks := make(map[string]int)
	minchars := 0

	i := 0
	for i < len(p) {
		c := p[i]
		switch c {
		case ANYCHAR:
			tseq = append(tseq, Token{"ANYCHAR", string(ANYCHAR)})
			toks[string(ANYCHAR)]++
			minchars++
		case ANYSEQ:
			tseq = append(tseq, Token{"ANYSEQ", string(ANYSEQ)})
			toks[string(ANYSEQ)]++
		default:
			j := i + 1
			litSeq := string(c)
			for j < len(p) && p[j] != ANYCHAR && p[j] != ANYSEQ {
				litSeq += string([]byte{p[j]})
				j++
			}
			tseq = append(tseq, Token{"LITSEQ", litSeq})
			minchars += len(litSeq)
			toks[litSeq]++
			i = j - 1
		}
		//last := len(tseq) - 1
		//fmt.Printf("[%s %s] ", tseq[last].Type, tseq[last].LitSeq)
		i++
	}
	//fmt.Println()
	//fmt.Printf("minchars: %d\n", minchars)

	// check for correct number of discrete chars
	if len(s) < minchars {
		return false
	}

	ind := 0
	consumed := 0
	for i, t := range tseq {
		fmt.Printf("Checking token %s   ind: %d\n", t.LitSeq, ind)
		if t.Type == "ANYCHAR" {
			if ind >= len(s) {
				return false
			} else {
				fmt.Printf("Match token [ANYCHAR] with %c\n", s[ind])
				ind++
				consumed++
			}
		} else if t.Type == "LITSEQ" {
			if ind >= len(s) {
				return false
			} else {
				if ind < len(s) && ind+len(t.LitSeq)-1 < len(s) && t.LitSeq == string(s[ind:ind+len(t.LitSeq)]) {
					fmt.Printf("Match token LITSEQ [%s] with [%s]\n", t.LitSeq, string(s[ind:ind+len(t.LitSeq)]))
					ind += len(t.LitSeq)
					consumed += len(t.LitSeq)
					toks[t.LitSeq] -= 1
				} else {
					return false
				}
			}
		} else { // case: token ANYSEQ
			if i >= len(tseq)-1 {
				return true
			} else {
				nextTok := tseq[i+1]
				if nextTok.Type == "ANYSEQ" {
					continue
				}
				nextLit := nextTok
				j := i + 1
				for j < len(tseq) && nextLit.Type != "LITSEQ" {
					nextLit = tseq[j]
					j++
				}
				if nextLit.Type != "LITSEQ" { // no literal left in pattern sequence
					if nextTok.Type == "ANYCHAR" {
						if ind < len(s) {
							return true
						} else {
							return false
						}
					} else { // only ANYSEQ tokens left in pattern
						return true
					}
				} else {
					numLeft := toks[nextLit.LitSeq]
					atLeast := minchars - consumed
					if nextTok.Type == "ANYCHAR" {
						atLeast--
					}
					fmt.Printf("numLeft: %d  atLeast: %d\n", numLeft, atLeast)
					litInd := lookBackwards(s, nextLit, ind, numLeft, atLeast)
					fmt.Printf("Current ind: %d  Next literal %s at index %d\n", ind, nextLit.LitSeq, litInd)
					if litInd < ind {
						return false
					} else {
						toAdd := litInd - ind
						if nextTok.Type == "ANYCHAR" {
							if litInd == ind {
								return false
							}
							toAdd--
						}
						ind += toAdd
						//consumed += toAdd
						//fmt.Printf("ANYSEQ consumed %d characters\n", toAdd)
					}
				}
			}
		}
	}
	//fmt.Printf("Consumed: %d  strlen: %d\n", consumed, len(s))
	if consumed < minchars || ind < len(s) {
		return false
	}
	return true
}


const ANYCHAR byte = '?'
const ANYSEQ byte = '*'

func lookBackwards(s string, c byte, start int) int {
	if c == ANYSEQ {
		return len(s) - 1
	} else if c == ANYCHAR {
		return start
	}
	i := len(s) - 1
	for i >= start {
		if s[i] == c {
			return i
		}
		i--
	}
	return i
}

func isMatch(s string, p string) bool {

	if len(s) == 0 {
		if len(p) == 0 {
			return true
		} else {
			for i := 0; i < len(p); i++ {
				if p[i] != ANYSEQ {
					return false
				}
			}
			return true
		}
	} else if len(p) == 0 {
		return false
	}

	match := true
	i := 0
	j := 0
	var next byte

	for i < len(p) && j < len(s) {
		//fmt.Printf("Start of loop s: %c p %c\n", s[j], p[i])
		if p[i] == ANYCHAR {
			//fmt.Printf("skipping char: %c\n", s[j])
			i++
			j++
		} else if p[i] == ANYSEQ {
			//fmt.Println("Skipping sequence")
			if i < len(p)-1 {
				next = p[i+1]
				for next == ANYSEQ && i < len(p)-2 {
					i++
					next = p[i+1]
				}
				if next == ANYCHAR {
					i++
					continue
				}
				//fmt.Printf("Nextchar: %c\n", next)
				j = lookBackwards(s, next, j)
				if j < i {
					return false
				}
				i++
			} else {
				return true
			}
		} else if p[i] == s[j] {
			//fmt.Printf("Match s:%c p:%c\n", s[j], p[i])
			i++
			j++
		} else {
			//fmt.Println("B3")
			return false
		}

		if i >= len(p) && j < len(s) {
			//fmt.Println("B1")
			return false
		} else if j >= len(s) && i < len(p) {
			//fmt.Println("B2")
			for i < len(p) {
				if p[i] != ANYSEQ {
					return false
				}
				i++
			}
		}
	}
	return match
}
*/
