// Given a sorted array arr1 and an array of positive integers arr2.
// Our task is finding the index of each values in array

package main

import "fmt"

func binarySearch(arr []int64, k int64) int64 {
	return index(arr, k, int64(len(arr))-1, 0)
}

// arr - given sorted array
// k - target value
// h - the highest value
// l - the lowest value. h and l are used to define the sub-array
func index(arr []int64, k, h, l int64) int64 {
	if h < l {
		return -1
	}
	mid := l + (h-l)/2
	if k == arr[mid] {
		return mid
	}
	if k > arr[mid] {
		return index(arr, k, h, mid+1)
	}
	return index(arr, k, mid-1, l)
}

func main() {
	arr := []int64{1, 5, 8, 12, 13}
	//arr2 := []int64{8, 1, 23, 1, 11}
	fmt.Println(binarySearch(arr, 8))
}
