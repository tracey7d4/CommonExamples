// https://www.hackerrank.com/challenges/coin-change/problem

package main

import "fmt"

// idea: we are make change of amount n using infinite number of coins of denominations arr[a,b,c,d]
// = how many way of change by using 0, 1, 2,... denomination a
// ways(arr[a,b,c,d],n)= ways(arr[b,c,d], n-0*a) + ways(arr[b,c,d], n- 1*a) + ways(arr[b,c,d], n- 2*a) and so on

func ways(n int64, c []int64) int64 {
	m := make(map[[2]int64]int64)
	return helper(c, int64(n), m)
}

// arr - the slice of coins
// n - the amount that needs to be satisfied
// m - map with the key being []int{len of the array, amount} and value being number of combinations

func helper(arr []int64, n int64, m map[[2]int64]int64) int64 {
	if n == 0 {
		return 1
	}
	k := int64(len(arr))
	if k == 0 {
		return 0
	}
	if n < 0 {
		return 0
	}

	if _, ok := m[[2]int64{k, n}]; ok {
		return m[[2]int64{k, n}]
	}

	count := int64(0)
	i := int64(0)
	for n-i*arr[0] >= 0 {
		count += helper(arr[1:], n-i*arr[0], m)
		i++
	}
	m[[2]int64{k, n}] = count

	return count
}

func main() {
	fmt.Println(ways(4, []int64{1, 2, 3}))
	fmt.Println(ways(10, []int64{2, 5, 3, 6}))
}