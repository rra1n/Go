package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	const N = 100010
	a := make([]int, N)
	pre := make([]int, N)
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a[1 : n+1])

	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + a[i]
	}

	ans := 4000000000
	for i := 1; i <= n; i++ {
		ans = min(ans, (i+1)*a[i]-pre[i]-(n-i+1)*a[i]+(pre[n]-pre[i]))
	}

	fmt.Println(ans)
}
