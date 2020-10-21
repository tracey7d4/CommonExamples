package main

import "fmt"

func FibonacciNumber(n int) int64 {
	m := make(map[int]int64)
	return FibonacciDP(n,m)
}

func FibonacciDP(n int, m map[int]int64) int64 {
	if n <= 1 {
		return int64(n)
	}
	if _,ok := m[n]; ok {
		return m[n]
	}
	fn := FibonacciDP(n-1,m) + FibonacciDP(n-2,m)
	m[n] = fn
	return fn
}

func main() {
	fmt.Println(FibonacciNumber(50))
}


