//Given an array of numbers, arrange them in a way that yields the largest value.

//Ans: Consider X and Y, check XY or YX bigger, sort descendingly
// Use modified merged sort big-O = nlogn

package main

import (
	"fmt"
	"strconv"
)

func largestNumber(arr []int) string {
	arr = mergedSort(arr)
	var s string
	for _, v := range arr {
		s = s + strconv.Itoa(v)
	}
	return s
}

func mergedSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(mergedSort(left), mergedSort(right))
}

func merge(left, right []int) []int {
	res := make([]int,0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}
		xy,_ := strconv.Atoi(strconv.Itoa(left[0])+strconv.Itoa(right[0]))
		yx,_ := strconv.Atoi(strconv.Itoa(right[0])+strconv.Itoa(left[0]))
		if xy > yx {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}

func main() {
	arr := []int{4, 6, 9, 23, 95}
	b := mergedSort(arr)
	fmt.Println(b)
	fmt.Println(largestNumber(arr))
}
