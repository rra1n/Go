package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t, m int
	Fscan(in, &t, &m)
	a, b := make([]int, m+1), make([]int, m+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &a[i], &b[i])
	}
	f := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]int, t+1)
	}
	for i := 1; i <= m; i++ {
		for j := 0; j <= t; j++ {
			f[i][j] = f[i-1][j]
			if j >= a[i] {
				f[i][j] = max(f[i][j], f[i][j-a[i]]+b[i])
			}
		}
	}
	Fprint(out, f[m][t])
}
