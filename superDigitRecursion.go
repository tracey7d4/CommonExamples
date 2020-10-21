// https://www.hackerrank.com/challenges/recursive-digit-sum/problem

package main

import (
	"fmt"
	"strconv"
)

func superDigit(n string) int32 {
	sum := int64(0)
	for _, v := range n {
		a, _ := strconv.ParseInt(string(v), 10, 64)
		sum += a
	}
	if sum < 10 {
		return int32(sum)
	}
	stringSum := strconv.FormatInt(sum, 10)
	return superDigit(stringSum)
}

func main() {
	fmt.Println(superDigit("128"))
}