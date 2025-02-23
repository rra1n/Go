package main

import (
	"bufio"
	. "fmt"
	"os"
)

type MinHeap struct{ data []int }

func (h *MinHeap) up(i int) {
	if i == 0 {
		return
	}
	fa := (i - 1) / 2
	if h.data[fa] > h.data[i] {
		h.data[fa], h.data[i] = h.data[i], h.data[fa]
		h.up(fa)
	}
}

func (h *MinHeap) down(i int) {
	n := len(h.data)
	l, r, j := 2*i+1, 2*i+2, i
	if l < n && h.data[l] < h.data[j] {
		j = l
	}
	if r < n && h.data[r] < h.data[j] {
		j = r
	}
	if j != i {
		h.data[j], h.data[i] = h.data[i], h.data[j]
		h.down(j)
	}
}

func (h *MinHeap) push(x int) {
	h.data = append(h.data, x)
	h.up(len(h.data) - 1)
}

func (h *MinHeap) top() int {
	return h.data[0]
}

func (h *MinHeap) pop() {
	n := len(h.data)
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	h.down(0)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)

	heap := &MinHeap{data: make([]int, 0)} // 初始化最小堆

	for i := 0; i < n; i++ {
		var op int
		Fscan(in, &op)
		switch op {
		case 1: // 插入操作
			var x int
			Fscan(in, &x)
			heap.push(x)
		case 2: // 查询最小值
			Fprintln(out, heap.top())
		case 3: // 删除最小值
			heap.pop()
		}
	}
}
