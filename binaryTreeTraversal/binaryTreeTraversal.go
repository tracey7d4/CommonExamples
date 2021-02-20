package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	key   string
	left  *node
	right *node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 2048 * 2048
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	arr := make([]node, n)
	for i := int64(0); i < n; i++ {
		scanner.Scan()
		data := strings.Split(scanner.Text(), " ")
		arr[i].key = data[0]
		j, _ := strconv.ParseInt(data[1], 10, 64)
		k, _ := strconv.ParseInt(data[2], 10, 64)
		if j == -1 {
			arr[i].left = nil
		} else {
			arr[i].left = &arr[j]
		}
		if k == -1 {
			arr[i].right = nil
		} else {
			arr[i].right = &arr[k]
		}
	}
	// inOrder
	fmt.Println("In-Order traversal")
	fmt.Println(inOrderTraversal(arr))

	// preOrder
	fmt.Println("pre-Order traversal")
	fmt.Println(preOrderTraversal(arr))

	// postOrder
	fmt.Println("post-Order traversal")
	fmt.Println(postOrderTraversal(arr))

	// breadth-first search
	fmt.Println("breadth-first traversal")
	fmt.Println(bfs(arr))
}

/*-------------------------------------------------------------
							STACK
We use stack on inOrder, preOrder, and postOrder traversal
--------------------------------------------------------------*/

type stack []node

func (s *stack) push(n node) {
	*s = append(*s, n)
}

func (s *stack) pop() *node {
	a := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return &a
}

func (s *stack) peek() node {
	return (*s)[len(*s)-1]
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}
//--------------------------------------------------------------

func inOrderTraversal(arr []node) []string {
	s := &stack{}
	curr := &arr[0]
	answer := make([]string,0,len(arr))
	for curr != nil || !s.isEmpty() {
		for curr != nil {
			s.push(*curr)
			curr = curr.left
		}
		curr = s.pop()
		answer = append(answer,curr.key)
		curr = curr.right
	}
	return answer
}

func preOrderTraversal(arr []node) []string {
	s := &stack{}
	s.push(arr[0])
	answer := make([]string,0, len(arr))
	for !s.isEmpty() {
		c := s.peek()
		answer = append(answer,c.key)
		s.pop()
		if c.right != nil {
			s.push(*c.right)
		}
		if c.left != nil {
			s.push(*c.left)
		}
	}
	return answer
}

func postOrderTraversal(arr []node) []string {
	s := &stack{}
	s1 := &stack{}
	s.push(arr[0])

	for !s.isEmpty() {
		c := s.peek()
		s1.push(c)
		s.pop()
		if c.left != nil {
			s.push(*c.left)
		}
		if c.right != nil {
			s.push(*c.right)
		}
	}
	answer := make([]string, 0, len(arr))
	for !s1.isEmpty() {
		c := s1.pop()
		answer = append(answer, c.key)
	}
	return answer
}

/*-------------------------------------------------------------
							QUEUE
We use queue on breadth first traversal
--------------------------------------------------------------*/
type queue []node
func newQueue() *queue{
	return &queue{}
}
func (q *queue) isEmpty() bool {
	return len(*q) == 0
}
func (q *queue) push(v node) {
	*q = append(*q,v)
}
func (q * queue) pop() node {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func bfs(arr []node) []string {
	q := newQueue()
	q.push(arr[0])
	answer := make([]string, 0, len(arr))
	answer = append(answer, arr[0].key)
	for !q.isEmpty() {
		c := q.pop()
		if c.left != nil {
			q.push(*c.left)
			answer = append(answer,c.left.key)
		}
		if c.right != nil {
			q.push(*c.right)
			answer = append(answer,c.right.key)
		}
	}
	return answer
}