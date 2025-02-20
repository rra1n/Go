package main

import (
	"bufio"
	. "fmt"
	"os"
)

type pair struct{ w, v int }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, p, s int
	Fscan(in, &n, &p, &s)
	pii := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &pii[i].w, &pii[i].v)
	}
	l, r := 0, 20000000000
	check := func(x int) bool {
		a := []pair{}
		for i := 1; i <= n; i++ {
			if pii[i].w <= x {
				a = append(a, pii[i])
			}
		}
		f := make([]int, s+1)
		for i := 0; i < len(a); i++ {
			for j := s; j >= a[i].w; j-- {
				f[j] = max(f[j], f[j-a[i].w]+a[i].v)
			}
		}
		return f[s] >= p
	}
	for l+1 < r {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	if r != 20000000000 {
		Fprint(out, r)
	} else {
		Fprint(out, "No Solution!")
	}
}
