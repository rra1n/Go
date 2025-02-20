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
	a, b := make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &b[i])
	}
	l, r := 0, 2000000000
	f := func(x int) bool {
		cnt := 0
		for i := 1; i <= n; i++ {
			if x-a[i] > b[i] {
				return false
			}
			if x-a[i] > 0 {
				cnt += x - a[i]
			}
		}
		return cnt <= m
	}
	for l+1 < r {
		mid := (l + r) >> 1
		if f(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	Fprint(out, l)
}
