package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	// defer out.Flush()

	var n, a, b int
	Fscanln(in, &n, &a, &b)

	w := make([]int, n)
	for i := 0; i < n; i++ {
		Fscanln(in, &w[i])
	}

	l, r := 0, 500010

	for l+1 < r {
		mid := (l + r) / 2
		x := 0
		for i := 0; i < n; i++ {
			if w[i] > mid*a {
				x += (w[i] - mid*a + b - 1) / b
			}
		}
		if x > mid {
			l = mid
		} else {
			r = mid
		}
	}
	Fprint(out, r)
	out.Flush()
}
