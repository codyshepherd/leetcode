/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

import (
	"strconv"
	"unicode"
)

type Token string

const (
	POPEN  Token = "POPEN"
	PCLOSE       = "PCLOSE"
	COMMA        = "COMMA"
	DASH         = "DASH"
	NUM          = "NUM"
	EOF          = "EOF"
)

var mapme = map[byte]Token{
	'[': POPEN,
	']': PCLOSE,
	',': COMMA,
	'-': DASH,
	'0': NUM,
	'1': NUM,
	'2': NUM,
	'3': NUM,
	'4': NUM,
	'5': NUM,
	'6': NUM,
	'7': NUM,
	'8': NUM,
	'9': NUM,
}

func TakeWhile(p func(rune) bool, s string) string {
	for i, r := range s {
		if !p(r) {
			return s[0:i]
		}
	}
	return s
}

//cursor, number
func num(s string, cursor int) (int, int) {
	num, _ := strconv.Atoi(TakeWhile(func(r rune) bool {
		if unicode.IsNumber(r) {
			return true
		} else {
			return false
		}
	}, s[cursor:]))

	for cursor < len(s) && unicode.IsNumber(rune(s[cursor])) {
		cursor++
	}
	return cursor, num

}

//Token, cursor, arg to num token
func token(s string, cursor int) (Token, int, int) {
	var lastarg int
	if cursor >= len(s) {
		return EOF, cursor, lastarg
	}
	t := mapme[s[cursor]]
	if t == NUM {
		cursor, lastarg = num(s, cursor)
	} else if t == DASH {
		cursor, lastarg = num(s, cursor+1)
		lastarg = -lastarg
	} else {
		cursor++
	}

	return t, cursor, lastarg
}

func trim(s string) string {
	parens := 1
	var ss string
	for i, v := range s {
		if v == ']' {
			parens--
		}
		if v == '[' {
			parens++
		}
		if parens == 0 {
			ss = s[:i]
			break
		}
	}
	return ss
}

func parse(s string, cursor int) *NestedInteger {
	var thislist *NestedInteger = new(NestedInteger)

	t, c, n := token(s, cursor)

	for t != EOF {
		switch t {
		case POPEN:
			substring := trim(s[c:])
			l := len(substring)
			if len(substring) == 0 {
				inner := new(NestedInteger)
				thislist.Add(*inner)
			} else {
				inner := parse(substring, 0)
				thislist.Add(*inner)
			}
			c += l
		case COMMA:
		case NUM, DASH:
			inner := new(NestedInteger)
			inner.SetInteger(n)
			thislist.Add(*inner)
		}
		t, c, n = token(s, c)
	}
	return thislist
}

func deserialize(s string) *NestedInteger {
	if s[0] != '[' {
		_, _, n := token(s, 0)
		nl := new(NestedInteger)
		nl.SetInteger(n)
		return nl
	}
	substring := trim(s[1:])

	var first *NestedInteger
	if len(substring) == 0 {
		first = new(NestedInteger)
	} else {
		first = parse(substring, 0)
	}
	return first
}