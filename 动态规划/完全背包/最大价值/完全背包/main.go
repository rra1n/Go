package main

import "fmt"

func main() {
	const N = 1010
	var n, m int
	v, w := make([]int, N), make([]int, N)
	f := make([][]int, N)
	for i := range f {
		f[i] = make([]int, N)
	}
	fmt.Scan(&n, &m)

	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= v[i] {
				f[i][j] = max(f[i-1][j], f[i][j-v[i]]+w[i])

			}
		}
	}

	fmt.Println(f[n][m])
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
