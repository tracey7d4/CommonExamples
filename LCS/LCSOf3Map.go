package main

import (
	"fmt"
)

func LCSof3Map(arr1, arr2, arr3 []int32) int {
	n := len(arr1)
	m := len(arr2)
	l := len(arr3)
	mx := make(map[[3]int]int)
	return LCSof3(arr1, arr2, arr3, n, m, l, mx)
}

func LCSof3(arr1, arr2, arr3 []int32, n, m, l int, mx map[[3]int]int) int {
	if n == 0 || m == 0 || l == 0 {
		return 0
	}
	if _, ok := mx[[3]int{n, m, l}]; ok {
		return mx[[3]int{n, m, l}]
	}
	if arr1[n-1] == arr2[m-1] && arr1[n-1] == arr3[l-1] {
		res := 1 + LCSof3(arr1, arr2, arr3, n-1, m-1, l-1, mx)
		mx[[3]int{n, m, l}] = res
		return res
	}
	res := maxMap(LCSof3(arr1, arr2, arr3, n-1, m, l, mx),
		LCSof3(arr1, arr2, arr3, n, m-1, l, mx),
		LCSof3(arr1, arr2, arr3, n, m, l-1, mx))
	mx[[3]int{n, m, l}] = res
	return res
}

func maxMap(a, b, c int) int {
	max := a
	if a < b {
		max = b
	}
	if max < c {
		return c
	}
	return max
}

func main() {
	arr1 := []int32{1, 2, 3}
	arr2 := []int32{2, 1, 3}
	arr3 := []int32{1, 3, 5}

	fmt.Println(LCSof3Map(arr1, arr2, arr3))
}




