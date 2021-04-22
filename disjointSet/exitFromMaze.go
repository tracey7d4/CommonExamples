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

type node struct {
	key    int
	parent int
	rank   int
}

func makeSet(i int) *node {
	return &node{i,i,0,
	}
}

func union(disjointSet []*node, firstID, secondID int) {
	if firstID == secondID {
		return
	}
	if disjointSet[firstID].rank < disjointSet[secondID].rank {
		disjointSet[firstID].parent = disjointSet[secondID].parent
	} else {
		disjointSet[secondID].parent = disjointSet[firstID].parent
		if disjointSet[secondID].rank == disjointSet[firstID].rank {
			disjointSet[firstID].rank += 1
		}
	}
}

func findID(disjointSet []*node, id int) int {
	var arr []int // for path compression
	for id != disjointSet[id].parent {
		arr = append(arr, id)
		id = disjointSet[id].parent
	}
	for _, v := range arr {
		disjointSet[v].parent = id
	}
	return id
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	n, _ := strconv.ParseInt(s[0], 10, 64) //number of vertices
	m, _ := strconv.ParseInt(s[1], 10, 64) // number of edges

	disjointSet := make([]*node, n) //disjointSet is an array of vertices
	for i := 0; i < int(n); i++ {
		disjointSet[i] = makeSet(i)
	}

	for i := 0; i < int(m); i++ {
		scanner.Scan()
		data := strings.Split(scanner.Text(), " ")
		j, _ := strconv.ParseInt(data[0], 10, 64) // first node in 1-based index
		k, _ := strconv.ParseInt(data[1], 10, 64) // second node in 1-based index
		//find set's ID of the 2 given nodes
		firstID := findID(disjointSet,int(j)-1)
		secondID := findID(disjointSet, int(k)-1)
		// if they belong to 2 different sets then do union
		if firstID != secondID {
			union(disjointSet, firstID, secondID)
		}
	}
	// get the required nodes that needed to check the path
	scanner.Scan()
	data := strings.Split(scanner.Text(), " ")
	j, _ := strconv.ParseInt(data[0], 10, 64)
	k, _ := strconv.ParseInt(data[1], 10, 64)
	destinationID := findID(disjointSet,int(j)-1)
	sourceID := findID(disjointSet,int(k)-1)
	if destinationID == sourceID {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

