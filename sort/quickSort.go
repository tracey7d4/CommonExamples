// 3-way partition procedure
// arr[left..i-1] - elements less than pivot
// arr[i..j] - elements equal to pivot
// arr[j+1 .. right] - elements greater than pivot

package main

import (
	"fmt"
)

func quickSort(arr []int32) []int32 {
	sort(arr, 0, int32(len(arr))-1)
	return arr
}

func sort(arr []int32, left, right int32) {
	//n := int32(len(arr))
	start := left
	stop := right
	if stop-start < 1 {
		return
	}
	splitIndex := start
	//pivot := rand.Int31n(n)
	pivot := arr[start+(stop-start)/2]
	for splitIndex <= stop {
		if arr[splitIndex] < pivot {
			arr[start], arr[splitIndex] = arr[splitIndex], arr[start]
			splitIndex++
			start++
		} else if arr[splitIndex] > pivot {
			arr[stop], arr[splitIndex] = arr[splitIndex], arr[stop]
			stop--
		} else {
			splitIndex++
		}
	}
	i, j := start-1, stop+1
	sort(arr, left, i)
	sort(arr, j, right)
}

func main() {
	arr := []int32{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 5, 4}
	fmt.Println(arr)
	fmt.Println(quickSort(arr))
}
