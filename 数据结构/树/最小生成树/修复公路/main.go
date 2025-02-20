package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type UnionFind struct {
	rank   []int
	parent []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size+1)
	rank := make([]int, size+1)
	for i := 1; i <= size; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{parent: parent, rank: rank}
}

func (uf *UnionFind) Union(x, y int) bool {
	rx, ry := uf.Find(x), uf.Find(y)
	if rx == ry {
		return false
	}
	if uf.rank[rx] < uf.rank[ry] {
		rx, ry = ry, rx
	}
	uf.parent[ry] = rx
	uf.rank[rx] += uf.rank[ry]
	return true
}

func (uf *UnionFind) Find(x int) int {
	root := x
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	for uf.parent[x] != root {
		fa := uf.parent[x]
		uf.parent[x] = root
		x = fa
	}
	return root
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstLine := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(firstLine[0])
	M, _ := strconv.Atoi(firstLine[1])

	edges := make([][3]int, M)
	for i := 0; i < M; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		t, _ := strconv.Atoi(line[2])
		edges[i] = [3]int{t, x, y}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	uf := NewUnionFind(N)
	maxTime := 0
	count := 0

	for _, edge := range edges {
		t, x, y := edge[0], edge[1], edge[2]
		if uf.Union(x, y) {
			maxTime = max(maxTime, t)
			count++
			if count == N-1 {
				break
			}
		}
	}

	if count == N-1 {
		fmt.Println(maxTime)
	} else {
		fmt.Println(-1)
	}
}
