package main

import "fmt"

// mergeSort O(nlgn)

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(mergeSort(left), mergeSort(right))
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
		if left[0] < right[0] {
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
	arr := []int{2, 3, 5, 6, 4, 8}

	fmt.Println(mergeSort(arr))
}
