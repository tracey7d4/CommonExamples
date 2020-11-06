package main

import (
	"fmt"
)

func LCSof2Map(arr1, arr2 []int32) int {
	n := len(arr1)
	k := len(arr2)
	m := make(map[[2]int]int)
	return LCSof2(arr1, arr2, n, k, m)
}

func LCSof2(arr1, arr2 []int32, n, k int, m map[[2]int]int) int {
	if n == 0 || k == 0 {
		return 0
	}
	if _, ok := m[[2]int{n, k}]; ok {
		return m[[2]int{n, k}]
	}
	if arr1[n-1] == arr2[k-1]  {
		res := 1 + LCSof2(arr1, arr2, n-1, k-1, m)
		m[[2]int{n, k}] = res
		return res
	}
	res := maxOf2(LCSof2(arr1, arr2, n-1, k, m), LCSof2(arr1, arr2, n, k-1,m))
	m[[2]int{n, k}] = res
	return res
}

func maxOf2(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr1 := []int32{1, 2, 3}
	arr2 := []int32{2, 1, 3}

	fmt.Println(LCSof2Map(arr1, arr2))
}




