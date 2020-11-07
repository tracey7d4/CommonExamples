package main

import (
	"fmt"
)

func LCSof2Map(arr1, arr2 []int32) int {
	n := len(arr1)
	m := len(arr2)
	mapOf2 := make(map[[2]int]int)
	return LCSof2(arr1, arr2, n, m, mapOf2)
}

func LCSof2(arr1, arr2 []int32, n, m int, mapOf2 map[[2]int]int) int {
	if n == 0 || m == 0 {
		return 0
	}
	if _, ok := mapOf2[[2]int{n, m}]; ok {
		return mapOf2[[2]int{n, m}]
	}
	if arr1[n-1] == arr2[m-1]  {
		res := 1 + LCSof2(arr1, arr2, n-1, m-1, mapOf2)
		mapOf2[[2]int{n, m}] = res
		return res
	}
	res := maxOf2(LCSof2(arr1, arr2, n-1, m, mapOf2), LCSof2(arr1, arr2, n, m-1, mapOf2))
	mapOf2[[2]int{n, m}] = res
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




