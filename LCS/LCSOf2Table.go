package main

import "fmt"

func LCSof2Table(arr1, arr2 []int32) int {
	n := len(arr1)
	m := len(arr2)

	var arr [101][101]int
// when i = 0 or j = 0 - means one of the given string is empty
// i = 1: consider 1st element in the arr1 (index 0) = arr[0]
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 || j == 0 {
				arr[i][j] = 0
			} else {
				if arr1[i-1] == arr2[j-1] {
					arr[i][j] = 1 + arr[i-1][j-1]
				} else {
					arr[i][j] = maxab(arr[i-1][j], arr[i][j-1])
				}
			}
		}
	}

	return arr[n][m]
}

func maxab(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr1 := []int32{1, 2, 3}
	arr2 := []int32{2, 1, 3}

	fmt.Println(LCSof2Table(arr1, arr2))
}