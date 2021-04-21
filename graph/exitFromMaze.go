package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// input format: The first line contains 2 integers n - number of vertices and m - number of edges.
// The next m lines show those edges between 2 vertices u and v
// The last line contains 2 vertices for us to check whether there is a path between them
// Output: output YES if there is a path, and NO otherwise

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 1024*1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	n, _ := strconv.ParseInt(s[0], 10, 64) //number of vertices
	m, _ := strconv.ParseInt(s[1], 10, 64) // number of edges

	graph := make([]*vertex, n) //graph is an array of vertices
	for i := 0; i < int(n); i++ {
		graph[i] = newVertex()
	}

	for i := 0; i < int(m); i++ {
		scanner.Scan()
		data := strings.Split(scanner.Text(), " ")
		j, _ := strconv.ParseInt(data[0], 10, 64)
		k, _ := strconv.ParseInt(data[1], 10, 64)
		graph[j-1].children[int(k)-1] = graph[k-1]
		graph[k-1].children[int(j)-1] = graph[j-1]
	}
	scanner.Scan()
	data := strings.Split(scanner.Text(), " ")
	j, _ := strconv.ParseInt(data[0], 10, 64)
	k, _ := strconv.ParseInt(data[1], 10, 64)
	isConnected(graph, int(j)-1, int(k)-1)
}

func isConnected(node []*vertex, j, k int) {
	q := newQueue()
	q.push(node[j])
	node[j].flag = true
	for !q.isEmpty() {
		s := q.pop()
		for key, value := range s.children {
			if key == k {
				fmt.Println("YES")
				return
			}
			if !value.flag {
				value.flag = true
				q.push(value)
			}
		}
	}
	fmt.Println("NO")
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