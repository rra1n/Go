package main

import . "fmt"

func main() {
	q := []int{1, 2, 3, 4, 5}
	for len(q) > 0 {
		u := q[0]
		Printf("队列长度为%d, 队首元素为%d。", len(q), u)
		q = q[1:]
	}

}
