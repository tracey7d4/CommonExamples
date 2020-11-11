// http://courses.csail.mit.edu/6.006/spring09/recitations/DP-problems.pdf

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	data := strings.Split(s, "")
	n := len(data)
	digit := make([]int64, 0, n/2+1)
	op := make([]string, 0, n/2)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			d, _ := strconv.ParseInt(data[i], 10, 32)
			digit = append(digit, d)
		} else {
			op = append(op, data[i])
		}
	}
	fmt.Println(getMaxValue(digit, op))
}

func getMaxValue(digit []int64, op []string) int64 {
	n := len(digit)
	maximum := newMatrix(n)
	minimum := newMatrix(n)
	for i, v := range digit {
		maximum[i][i] = v
		minimum[i][i] = v
	}
	s := 1
	for s < n {
		for i := 0; i < n-s; i++ {
			j := i + s
			minimum[i][j], maximum[i][j] = minAndMax(maximum, minimum, i, j, op)
		}
		s++
	}

	return maximum[0][n-1]
}

func newMatrix(n int) [][]int64{
	m := make([][]int64,n)
	for i, _ := range m {
		m[i] = make([]int64,n)
	}
	return m
}

func calc(a, b int64, op string) int64 {
	if op == "+" {
		return a + b
	}
	if op == "-" {
		return a - b
	}
	return a * b
}

func minAndMax(maximum, minimum [][]int64, i, j int, op []string) (int64, int64) {
	maxA := int64(math.MinInt64)
	minA := int64(math.MaxInt64)
	k := i
	for k < j {
		a := calc(maximum[i][k], maximum[k+1][j], op[k])
		b := calc(maximum[i][k], minimum[k+1][j], op[k])
		c := calc(minimum[i][k], maximum[k+1][j], op[k])
		d := calc(minimum[i][k], minimum[k+1][j], op[k])
		minA = min1([]int64{minA, a, b, c, d})
		maxA = max2([]int64{maxA, a, b, c, d})
		k++
	}
	return minA, maxA
}

func min1(arr []int64) int64 {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min
}

func max2(arr []int64) int64 {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

//func printM(a [][]int64) {
//	for _,v := range a {
//		fmt.Println(v)
//	}
//}