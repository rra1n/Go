package main

import (
	"bufio"
	. "fmt"
	"os"
)

func f(a []int, l, r int) int {
	pivot := a[l]
	for l < r {
		for l < r && a[r] >= pivot {
			r--
		}
		a[l] = a[r]
		for l < r && a[l] <= pivot {
			l++
		}
		a[r] = a[l]
	}
	a[l] = pivot
	return l
}

func g(a []int, l, r int) {
	if l < r {
		i := f(a, l, r)
		g(a, l, i-1)
		g(a, i+1, r)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	g(a, 1, n)
	for i := 1; i <= n; i++ {
		Fprintln(out, a[i])
	}
}
