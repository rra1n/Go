package main

import (
	"fmt"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	t := redblacktree.NewWithIntComparator()
	t.Put(10, 20)
	t.Put(20, 30)
	t.Put(20, 100)

	x, _ := t.Ceiling(20)
	it := t.IteratorAt(x)
	if it.Prev() {
		fmt.Println(it.Key())
	} else {
		fmt.Println(it.Key())
	}
}
