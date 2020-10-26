package main

import "fmt"

func BoyerMoore(arr []int32) bool {
	count := 0
	currentNumber := int32(0)
	for _,v := range arr {
		if count == 0 {
			currentNumber = v
			count = 1
		}
		if v == currentNumber {
			count += 1
		} else {
			count -= 1
		}
	}
	fmt.Println(currentNumber)
	count = 0
	for _, v := range arr {
		if v == currentNumber {
			count += 1
		}
	}
	if count > len(arr)/2 {
		return true
	}
	return false
}

func main() {
	fmt.Println(BoyerMoore([]int32{2,2,3,3,4,2,2}))
	fmt.Println(BoyerMoore([]int32{7,7,1,5,5,5}))
}