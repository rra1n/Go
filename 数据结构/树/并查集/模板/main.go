package main

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

}
