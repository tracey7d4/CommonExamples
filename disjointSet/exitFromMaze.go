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
	const maxCapacity = 1024*1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	n, _ := strconv.ParseInt(s[0], 10, 64)
	m, _ := strconv.ParseInt(s[1], 10, 64)
	node := make([]*vertex, n)
	for i := 0; i < int(n); i++ {
		node[i] = newVertex()
	}

	for i := 0; i < int(m); i++ {
		scanner.Scan()
		data := strings.Split(scanner.Text(), " ")
		j, _ := strconv.ParseInt(data[0], 10, 64)
		k, _ := strconv.ParseInt(data[1], 10, 64)
		node[j-1].children[int(k)-1] = node[k-1]
		node[k-1].children[int(j)-1] = node[j-1]
	}
	scanner.Scan()
	data := strings.Split(scanner.Text(), " ")
	j, _ := strconv.ParseInt(data[0], 10, 64)
	k, _ := strconv.ParseInt(data[1], 10, 64)
	isConnected(node, int(j)-1, int(k)-1)
}

type vertex struct {
	children map[int]*vertex
	flag	bool
}

func newVertex() *vertex{
	return &vertex{
		make(map[int]*vertex),
		false,
	}
}

func isConnected(node []*vertex, j, k int) {
	q := newQueue()
	q.push(node[j])
	node[j].flag = true
	for !q.isEmpty() {
		s := q.pop()
		for key, value := range s.children {
			if key == k {
				fmt.Println("1")
				return
			}
			if !value.flag {
				value.flag = true
				q.push(value)
			}
		}
	}
	fmt.Println("0")
}

type queue []*vertex
func newQueue() *queue{
	return &queue{}
}
func (q *queue) isEmpty() bool {
	return len(*q) == 0
}
func (q *queue) push(v *vertex) {
	*q = append(*q,v)
}
func (q * queue) pop() *vertex {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}