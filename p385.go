package main

import (
	"strconv"
	"unicode"
)

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
func TakeWhile(p func(rune) bool, s string) string {
	for i, r := range s {
		if !p(r) {
			return s[0:i]
		}
	}
	return s
}

const BOPEN byte = '['
const BCLOSED byte = ']'
const COMMA byte = ','
const DASH byte = '-'

type single int
type list []single

func tokens(s string) list {
	switch s[0] {
	case BOPEN:
		nl := new(list)

	}
}

func deserialize(s string) *NestedInteger {
	switch s[0] {
	case BOPEN:
		var outer *NestedInteger = new(NestedInteger)

	}
}

/*  Do below, but "look ahead". i.e. look at next "token", add()ing.
Create new object & recurse only on brackets
*/

func deserialize(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	} else if s == "[]" {
		var we *NestedInteger = new(NestedInteger)
		return we
	}

	switch s[0] {
	case BOPEN:
		if s[1] == '[' {
			var we *NestedInteger = new(NestedInteger)
			ni := deserialize(s[1:])
			if ni != nil {
				we.Add(*ni)
			}
			return we
		} else if s[1] == ']' {
			var we *NestedInteger = new(NestedInteger)
			return we
		}
		return deserialize(s[1:])
	case BCLOSED:
		return nil
	case COMMA:
		return deserialize(s[1:])
	case DASH:
		var busta *NestedInteger = new(NestedInteger)
		value, _ := strconv.Atoi(TakeWhile(func(r rune) bool {
			if unicode.IsNumber(r) {
				return true
			} else {
				return false
			}
		}, s[1:]))
		i := 1
		for unicode.IsNumber(rune(s[i])) {
			i++
			if i >= len(s) {
				break
			}
		}
		var inner *NestedInteger = new(NestedInteger)
		inner.SetInteger(-value)
		if i < len(s) {
			busta.Add(*inner)
			ni := deserialize(s[i:])
			if ni != nil {
				busta.Add(*ni)
			}
		} else {
			return inner
		}

		return busta
	default:
		var busta *NestedInteger = new(NestedInteger)
		value, _ := strconv.Atoi(TakeWhile(func(r rune) bool {
			if unicode.IsNumber(r) {
				return true
			} else {
				return false
			}
		}, s))
		i := 0
		for unicode.IsNumber(rune(s[i])) {
			i++
			if i >= len(s) {
				break
			}
		}
		var inner *NestedInteger = new(NestedInteger)
		inner.SetInteger(value)
		if i < len(s) {
			busta.Add(*inner)
			ni := deserialize(s[i:])
			if ni != nil {
				busta.Add(*ni)
			}
		} else {
			return inner
		}

		return busta
	}
}
