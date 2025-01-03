package main

import "fmt"

func main() {
	const N = 110

	var m, n int
	fmt.Scanln(&m, &n)

	f := make([][]int, N)
	for i := 0; i < N; i++ {
		f[i] = make([]int, N)
	}
	w, t := make([]int, N), make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Scanln(&t[i], &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= t[i] {
				f[i][j] = max(f[i-1][j], f[i-1][j-t[i]]+w[i])
			}
		}
	}

	fmt.Println(f[n][m])
}
