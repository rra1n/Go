package main

import . "fmt"

func main() {
	var t, n int
	Scanln(&t, &n)
	a, b := make([]int, n+1), make([]int, n+1)
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, t+1)
	}
	for i := 1; i <= n; i++ {
		Scanln(&a[i], &b[i])
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= t; j++ {
			f[i][j] = f[i-1][j]
			if j >= a[i] {
				f[i][j] = max(f[i][j], f[i-1][j-a[i]]+b[i])
			}
		}
	}
	Print(f[n][t])
}
