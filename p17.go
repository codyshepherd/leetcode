package main

import "fmt"

func mapme(num rune) string {
	switch num {
	case '2':
		return "abc"
	case '3':
		return "def"
	case '4':
		return "ghi"
	case '5':
		return "jkl"
	case '6':
		return "mno"
	case '7':
		return "pqrs"
	case '8':
		return "tuv"
	case '9':
		return "wxyz"
	default:
		return ""
	}
}

func fixme(prefix string, lst []string) []string {
	first := lst[0]
	var rest []string
	if len(lst) > 1 {
		rest = lst[1:]
	}
	var ret []string

	for _, item := range first {
		newprefix := prefix + string(item)
		if len(lst) == 1 {
			ret = append(ret, newprefix)
		} else {
			ret = append(ret, fixme(newprefix, rest)...)
		}
	}
	return ret
}

func letterCombinations(digits string) []string {
	var buckets []string
	if digits == "" {
		return buckets
	}
	//var lens []int
	var final []string
	var rest []string

	for _, digit := range digits {
		opts := mapme(digit)
		buckets = append(buckets, opts)
		//lens := len(opts)
	}

	startbucket := buckets[0]
	if len(buckets) > 1 {
		rest = buckets[1:]
	}

	for _, prefix := range startbucket {
		if len(rest) == 0 {
			final = append(final, string(prefix))
		} else {
			final = append(final, fixme(string(prefix), rest)...)
		}
	}

	return final
}

func main() {
	str := "234"

	combos := letterCombinations(str)

	for _, combo := range combos {
		fmt.Println(combo)
	}
	fmt.Println("Len: ", len(combos))
}
