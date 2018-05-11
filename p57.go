//COdy Shepherd
// leetcode challeng 57

package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for an interval.
 */
type Interval struct {
	Start int
	End   int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{newInterval}
	}

	var toCollapse []int

	for j, intv := range intervals {
		if (intv.Start <= newInterval.Start && intv.End >= newInterval.Start) ||
			(intv.Start > newInterval.Start && intv.End < newInterval.End) ||
			(intv.Start <= newInterval.End && intv.End >= newInterval.End) {
			toCollapse = append(toCollapse, j)
		}
	}
	//fmt.Println("toCollapse ", toCollapse)
	if len(toCollapse) == 0 {
		if newInterval.Start > intervals[len(intervals)-1].End {
			return append(intervals, newInterval)
		} else if newInterval.End < intervals[0].Start {
			return append([]Interval{newInterval}, intervals...)
		} else {
			i := 0
			for i < len(intervals) {
				if intervals[i].End < newInterval.Start && intervals[i+1].Start > newInterval.End {
					break
				}
				i += 1
			}
			//fmt.Println("i ", i)
			//fmt.Println("Intervals: ", intervals)
			newlist := make([]Interval, i+1)
			copy(newlist, intervals[:i+1])
			//fmt.Println("newlist ", newlist)
			newlist = append(newlist, newInterval)
			//fmt.Println(intervals[i+1:])
			newlist = append(newlist, intervals[i+1:]...)
			return newlist
		}
	}
	last := len(toCollapse) - 1
	newlist := intervals[:toCollapse[0]]
	middle := Interval{min(intervals[toCollapse[0]].Start, newInterval.Start), max(intervals[toCollapse[last]].End, newInterval.End)}
	newlist = append(newlist, middle)
	newlist = append(newlist, intervals[toCollapse[last]+1:]...)
	return newlist
}

func makeIntervals(s string) []Interval {
	lst := strings.Split(strings.Trim(s, "[]"), ",")
	var intlist []Interval
	var b int
	var e int
	for i, st := range lst {
		item := strings.Split(strings.Trim(st, "[]"), ",")
		if (i+1)%2 == 0 {
			e, _ = strconv.Atoi(item[0])
			inv := Interval{Start: b, End: e}
			intlist = append(intlist, inv)
		} else {
			b, _ = strconv.Atoi(item[0])
		}
	}

	return intlist
}

func main() {
	ints := makeIntervals("[[1,2],[3,5],[6,7],[8,10],[12,16]]")
	fmt.Println(ints)
	newints := insert(ints, Interval{4, 8})
	fmt.Println(newints)

	ints = makeIntervals("[[1,3],[6,9]]")
	fmt.Println(ints)
	newints = insert(ints, Interval{2, 5})
	fmt.Println(newints)

	ints = makeIntervals("[[1,5]]")
	fmt.Println(ints)
	newints = insert(ints, Interval{6, 8})
	fmt.Println(newints)

	ints = makeIntervals("[[3,5],[12,15]]")
	fmt.Println(ints)
	newints = insert(ints, Interval{6, 6})
	fmt.Println(newints)
}
