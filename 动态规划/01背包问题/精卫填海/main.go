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

	var v, n, c int
	Fscan(in, &v, &n, &c)
	a, b := make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i], &b[i])
	}

	f := make([]int, v+1)
	for j := 1; j <= v; j++ {
		f[j] = 0x3f3f3f3f
	}
	for i := 1; i <= n; i++ {
		for j := v; j >= 0; j-- {
			k := max(0, j-a[i])
			f[j] = min(f[j], f[k]+b[i])
		}
	}

	if f[v] > c {
		Fprint(out, "Impossible")
	} else {
		Fprint(out, c-f[v])
	}
}
