package main

import (
	"fmt"
	"sort"
)

func sortArray(arr []int32) bool {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	n := len(arr)
	for i := 0; i < n/2; i++ {
		if arr[i] == arr[n/2] {
			if arr[i+n/2] == arr[n/2] {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(sortArray([]int32{1,2,2,2,3,4,2}))
	fmt.Println(sortArray([]int32{1,2,2,2,3}))
	fmt.Println(sortArray([]int32{1,1,1,2}))
}