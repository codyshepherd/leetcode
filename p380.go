package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type RandomizedSet struct {
	table [][]int
	list  []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	t := make([][]int, 19991)
	rs := RandomizedSet{table: t}
	return rs
}

//Taken from https://gobyexample.com/collection-functions
func Any(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// Taken from https://gobyexample.com/collection-functions
func Find(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	l := len(this.table)
	v := int(math.Abs(float64(val)))
	ind := v % l
	entrylist := this.table[ind]

	if Any(entrylist, func(i int) bool {
		if i == val {
			return true
		}
		return false
	}) {
		return false
	}
	this.table[ind] = append(entrylist, val)
	this.list = append(this.list, val)
	sort.IntSlice.Sort(this.list)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	l := len(this.table)
	v := int(math.Abs(float64(val)))
	ind := v % l
	entrylist := this.table[ind]

	if len(entrylist) == 0 {
		return false
	} else {
		for i, value := range entrylist {
			if value == val {
				this.table[ind] = append(entrylist[:i], entrylist[i+1:]...)
				index := Find(this.list, val)
				this.list = append(this.list[:index], this.list[index+1:]...)
				return true
			}
		}
		return false
	}
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	i := rand.Int() % len(this.list)
	return this.list[i]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	obj := Constructor()

	inst := []string{"RandomizedSet", "insert", "insert", "remove", "insert", "insert", "insert", "remove", "remove", "insert", "remove", "insert", "insert", "insert", "insert", "insert", "getRandom", "insert", "remove", "insert", "insert"}
	vals := []int{0, 3, -2, 2, 1, -3, -2, -2, 3, -1, -3, 1, -2, -2, -2, 1, 0, -2, 0, -3, 1}

	for i, val := range inst {
		if val == "RandomizedSet" {

		} else if val == "insert" {
			fmt.Println(obj.Insert(vals[i]))
		} else if val == "remove" {
			fmt.Println(obj.Remove(vals[i]))
		} else if val == "getRandom" {
			fmt.Println(obj.GetRandom())
		}
	}

}
