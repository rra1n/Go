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
	var n, m int
	Fscan(in, &n, &m)
	a, f := make([]int, n+1), make([][]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m+1)
	}
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= a[i] {
				f[i][j] += f[i-1][j-a[i]]
			}
		}
	}
	Fprint(out, f[n][m])
}
