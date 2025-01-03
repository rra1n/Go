package main

import "fmt"

func main() {
	const N = 20010
	a := make([]int, N)

	f := make([][]int, N)
	for i := 0; i < N; i++ {
		f[i] = make([]int, N)
	}

	var m, n int
	fmt.Scanln(&m)
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Scanln(&a[i])
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= a[i] {

				f[i][j] = max(f[i-1][j], f[i-1][j-a[i]]+a[i])
			}
		}
	}

	fmt.Println(m - f[n][m])
}
