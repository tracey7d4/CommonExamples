package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()
	s := strings.Split(data, "")
	fmt.Println(checkBrackets(s))
}

func checkBrackets(s []string) string {
	var brackets stack
	for i , v := range s {
		if v == "[" || v == "{" || v == "(" {
			brackets.push(v,i+1)
		} else if v == ")" || v == "]" || v == "}" {
			res := brackets.pop()
			if res == "false" || res != matchS(v) {
				return "Unbalanced at index "+strconv.Itoa(i+1)
			}
		}
	}
	if !brackets.isEmpty() {
		return "Unbalanced at index "+strconv.Itoa(brackets.index[len(brackets.index)-1])
	}
	return "Balanced"
}

type stack struct {
	value []string
	index []int
}

// isEmpty: check if stack is empty
func (s *stack) isEmpty() bool {
	return len((*s).value) == 0
}

// Push a new key onto the stack
func (s *stack) push(key string, i int) {
	(*s).value = append((*s).value, key)
	(*s).index = append((*s).index, i)
}

// Remove top element of stack
func (s *stack)	pop() string {
	if s.isEmpty() {
		return "false"
	}
	i := len((*s).value)-1
	res := (*s).value[i]
	(*s).value = (*s).value[:i]
	(*s).index = (*s).index[:i]
	return res
}

func matchS(s string) string {
	if s == ")" {
		return "("
	}
	if s == "]" {
		return "["
	}
	return "{"
}