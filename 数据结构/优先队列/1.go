package main

import (
	"container/heap"
	"fmt"
)

type hp []int

func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool { return h[i] > h[j] }
func (h hp) Len() int           { return len(h) }
func (h *hp) Push(x any)        { *h = append(*h, x.(int)) }
func (h *hp) Pop() any {
	n := h.Len()
	old := *h // 解引用
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &hp{2, 1, 3, 5, 10, 8, 9}
	heap.Init(h)
	for h.Len() > 0 {
		x := heap.Pop(h)
		fmt.Println(x)
	}
}
