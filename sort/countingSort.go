package main

import "fmt"

func main() {
	arr := []int32{1, 4, 1, 2, 7, 5, 2}
	fmt.Println(countingSort(arr))
}

func findMax(arr []int32) int32 {
	max := int32(0)
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}

func countingSort(arr []int32) []int32 {
	n := findMax(arr)
	countArr := make([]int32, n+1) // array which index represents data range
	for _, v := range arr {
		countArr[v] += 1
	}
	for i := int32(1); i < n+1; i++ {
		countArr[i] += countArr[i-1]
	}
	res := make([]int32, len(arr))
	for _, v := range arr {
		index := countArr[v] - 1
		res[index] = v
		countArr[v]--
	}
	return res
}
