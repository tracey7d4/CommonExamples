package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type stacks struct {
	value    []int32
	maxArray []int32
	minArray []int32
	max      int32
	min      int32
}

func (stk *stacks) pushStack(n int32) {
	fmt.Printf("Done pushing %v\n", n)
	(*stk).value = append((*stk).value, n)
	if (*stk).max < n {
		(*stk).max = n
	}
	if (*stk).min > n {
		(*stk).min = n
	}
	(*stk).maxArray = append((*stk).maxArray, (*stk).max)
	(*stk).minArray = append((*stk).minArray, (*stk).min)
}

func (stk *stacks) pop() {
	n := len((*stk).value)
	if n == 0 { //stack is empty
		fmt.Println("Stack is empty")
		return
	}
	fmt.Printf("Pop %v\n", (*stk).value[n-1] )
	(*stk).value = (*stk).value[:n-1] // pop original array

	// update max and min value after pop operation
	if n > 1 {
		(*stk).max = (*stk).maxArray[n-2]
		(*stk).min = (*stk).minArray[n-2]
	} else {
		(*stk).max = -2                   // reset max value after pop operation (n == 1)
		(*stk).min = int32(math.MaxInt32) // reset min value after pop operation (n == 1)
	}
	// pop maxArray and minArray
	(*stk).maxArray = (*stk).maxArray[:n-1]
	(*stk).minArray = (*stk).minArray[:n-1]
}

func (stk *stacks) printStack() {
	n := len((*stk).value)
	if n == 0 {
		fmt.Println("Stack is empty")
	} else {
		str := fmt.Sprintf("%v", (*stk).value)
		str = strings.Trim(str, "[]")
		fmt.Println(str)
	}
}

// constraint : value is positive int32
func main() {
	fmt.Println("Please type in your request")
	fmt.Println("	- Syntax `push (value)` is to push a value in, ex. push 3")
	fmt.Println("	- Syntax `pop` is to pop a value out")
	fmt.Println("	- Syntax `print` is to print out stack at current time")
	fmt.Println("	- Syntax `max` is to print out max value of stack at current time")
	fmt.Println("	- Syntax `min` is to print out min value of stack at current time")

	stk := new(stacks)
	stk.max = -2
	stk.min = math.MaxInt32

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		v := scanner.Text()
		d := strings.Split(v, " ")
		if len(d) == 2 {
			k, _ := strconv.ParseInt(d[1], 10, 64)
			stk.pushStack(int32(k))
		} else {
			switch v {
			case "pop":
				stk.pop()
			case "print":
				stk.printStack()
			case "max":
				if stk.max != -2 {
					fmt.Println(stk.max)
				} else {
					fmt.Println("Stack is empty")
				}
			case "min":
				if len(stk.minArray) != 0 {
					fmt.Println(stk.min)
				} else {
					fmt.Println("Stack is empty")
				}
			case "stop":
				fmt.Println("See you later!")
				return
			default:
				fmt.Println("Wrong Syntax. Please try again")
			}
		}
	}
}
