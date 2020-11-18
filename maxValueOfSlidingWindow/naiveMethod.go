package main

import (
	"fmt"
)

func main() {
	arr := []int32{7, 6, 5, 4, 8, 9, 1}
	fmt.Println(maxOfWindow(arr, 3))
}

func maxOfWindow(arr []int32, m int32) []int32 {
	n := int32(len(arr))
	res := make([]int32, 0, n)
	for i := int32(0); i <= n - m; i++ {
		max := arr[i]
		for j := i+1; j < i+m; j++ {
			if max < arr[j] {
				max = arr[j]
			}
		}
		res = append(res, max)
	}
	return res
}