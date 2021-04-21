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
	key    int
	parent int
	rank   int
}

func newVertex(i int) *vertex {
	return &vertex{i+1,i+1,0,
	}
}
func mergeSet(disjointSet []*vertex, destinationID, sourceID int) {
	if destinationID == sourceID {
		return
	}
	if disjointSet[destinationID-1].rank < disjointSet[sourceID-1].rank {
		disjointSet[destinationID-1].parent = disjointSet[sourceID-1].parent
	} else {
		disjointSet[sourceID-1].parent = disjointSet[destinationID-1].parent
		if disjointSet[sourceID-1].rank == disjointSet[destinationID-1].rank {
			disjointSet[destinationID-1].rank += 1
		}
	}
}

func findID(disjointSet []*vertex, id int) int {
	var arr []int // for path compression
	for id != disjointSet[id-1].parent {
		arr = append(arr, id-1)
		id = disjointSet[id-1].parent
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

	disjointSet := make([]*vertex, n) //disjointSet is an array of vertices
	for i := 0; i < int(n); i++ {
		disjointSet[i] = newVertex(i)
	}

	for i := 0; i < int(m); i++ {
		scanner.Scan()
		data := strings.Split(scanner.Text(), " ")
		j, _ := strconv.ParseInt(data[0], 10, 64)
		k, _ := strconv.ParseInt(data[1], 10, 64)
		destinationID := findID(disjointSet,int(j))
		sourceID := findID(disjointSet, int(k))
		if destinationID != sourceID {
			mergeSet(disjointSet,destinationID,sourceID)
		}
	}

	scanner.Scan()
	data := strings.Split(scanner.Text(), " ")
	j, _ := strconv.ParseInt(data[0], 10, 64)
	k, _ := strconv.ParseInt(data[1], 10, 64)
	destinationID := findID(disjointSet,int(j))
	sourceID := findID(disjointSet,int(k))
	if destinationID == sourceID {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

