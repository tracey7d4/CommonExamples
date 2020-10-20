package main

import "fmt"

func fibonachi(n int) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := int64(0)
	b := int64(1)
	var fi int64
	for i := 2; i <= n; i++ {
		fi = a + b
		a,b = b, fi
	}
	return fi
}

func lastDigitFibonacci(n int64) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := 0
	b := 1
	var fi int
	for i := int64(2); i <= n; i++ {
		fi = (a + b)%10
		a,b = b, fi
	}
	return fi
}

// Sum of the first 𝑛 Fibonacci numbers.
// Note: sum_{i=0 ->n)F_i = F_(n+2)-1

func lastDigitSum(n int64) int {
	return lastDigitFibonacci(n+2)-1
}

//Partial sum of Fibonacci numbers: 𝐹𝑚 + 𝐹𝑚+1 + · · · + 𝐹𝑛.
// Answer: 𝐹𝑚 + 𝐹𝑚+1 + · · · + 𝐹𝑛 = F_(n+2) - F_(m+1)
func lastDigitPartialSum(m,n int64) int {
	a := lastDigitFibonacci(n+2)
	b := lastDigitFibonacci(m+1)
	if a > b {
		return a-b
	}
	return a+10-b
}

// Sum of square of Fibonacci numbers
// Note: compute: Fn*F(n+1)
func lastDigitSquareSumFib(n int64) int {
	return (lastDigitFibonacci(n+1)*lastDigitFibonacci(n))%10
}


func main() {
	fmt.Println(lastDigitSum(100))
}

