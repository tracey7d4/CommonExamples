package main

import "fmt"

func bruteForce(arr []int32) bool {
	n := len(arr)
	for _, v:= range arr {
		count := 0
		for i := 0; i < n; i++ {
			if arr[i] == v {
				count += 1
			}
		}
		if count > n/2 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(bruteForce([]int32{2,2,3,3,4,2,2}))
	fmt.Println(bruteForce([]int32{2,2,3,3}))
}