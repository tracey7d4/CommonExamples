package main

import "fmt"

func GCD(a,b int64) int64 {
	r := int64(1)
	for r !=0 {
		r = a%b
		if r == 0 {
			return b
		}
		a,b = b,r
	}
	return a/b
}

func LCM(a,b int64) int64 {
	return a*b/GCD(a,b)
}
func main() {
	fmt.Println(GCD(357,234))
	fmt.Println(LCM(9,6))
}