package main

import "fmt"

func LCSof3Table(arr1, arr2, arr3 []int32) int {
	n := len(arr1)
	m := len(arr2)
	var l = len(arr3)
	arr := make([][101][101]int, 101)

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= l; k++ {
				if i == 0 || j == 0 || k == 0 {
					arr[i][j][k] = 0
				} else {
					if arr1[i-1] == arr2[j-1] && arr1[i-1] == arr3[k-1] {
						arr[i][j][k] = 1 + arr[i-1][j-1][k-1]
					} else {
						arr[i][j][k] = maxabc(arr[i-1][j][k], arr[i][j-1][k], arr[i][j][k-1])
					}
				}
			}
		}
	}

	return arr[n][m][l]
}

func maxabc(a, b, c int) int {
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

	fmt.Println(LCSof3Table(arr1, arr2, arr3))
}
