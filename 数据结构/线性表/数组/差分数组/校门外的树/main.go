package main

import "fmt"

func main() {
	const N = 1e4 + 10
	diff := make([]int, N)
	var n, m int
	fmt.Scanln(&n, &m)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Scanln(&a, &b)
		diff[a]--
		diff[b+1]++
	}
	s, x := 0, 0
	for i := 0; i <= n; i++ {
		s += diff[i]
		if s < 0 {
			x++
		}
	}
	fmt.Println(n - x + 1)
}
