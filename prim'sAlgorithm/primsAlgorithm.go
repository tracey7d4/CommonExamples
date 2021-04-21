package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	children map[int32]int32
	flag     bool
}

func newNode() *node {
	return &node {
		make(map[int32]int32),
		false,
	}
}

func prims(n int32, edges [][]int32, start int32) int64 {
	graph := make([]*node,n)
	for i := int32(0); i < n; i++ {
		graph[i] = newNode()
	}

	for _, v := range edges { //edges are given in 1-based index, --> transform to 0-based index
		_, ok := graph[v[0]-1].children[v[1]-1]
		if !ok || graph[v[0]-1].children[v[1]-1] > v[2] {
			graph[v[0]-1].children[v[1]-1] = v[2] //v[2] is weight of edge
			graph[v[1]-1].children[v[0]-1] = v[2]
		}
	}

	s := start -1
	graph[s].flag = true
	weight := make(map[int32]int32)
	ans := int64(0)
	count := int32(1)
	for count < n {
		for key, value := range graph[s].children {
			if !graph[key].flag {
				if _, ok := weight[key]; !ok {
					weight[key] = value
				} else {
					if value < weight[key] {
						weight[key] = value
					}
				}
			}
		}
		i, w := getMin(weight)
		ans += w
		s = i
		graph[i].flag = true
		delete(weight,i)
		count++
	}
	return ans
}

func getMin(weight map[int32]int32) (int32,int64) {
	min := int64(100001)
	index := int32(0)
	for i, v := range weight {
		if min > int64(v) {
			min = int64(v)
			index = i
		}
	}
	return index, min
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var edges [][]int32
	for i := 0; i < int(m); i++ {
		edgesRowTemp := strings.Split(readLine(reader), " ")

		var edgesRow []int32
		for _, edgesRowItem := range edgesRowTemp {
			edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
			checkError(err)
			edgesItem := int32(edgesItemTemp)
			edgesRow = append(edgesRow, edgesItem)
		}

		if len(edgesRow) != int(3) {
			panic("Bad input")
		}

		edges = append(edges, edgesRow)
	}

	startTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	start := int32(startTemp)

	fmt.Println(prims(n, edges, start))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
