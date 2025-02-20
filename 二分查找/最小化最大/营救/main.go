package main

import (
	"bufio"
	. "fmt"
	"os"
)

type pair struct{ v, w int }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var n, m, s, t, u, v, w int
	Fscanln(in, &n, &m, &s, &t)
	g := make([][]pair, n+1)
	for i := 1; i <= m; i++ {
		Fscanln(in, &u, &v, &w)
		g[u] = append(g[u], pair{v, w})
		g[v] = append(g[v], pair{u, w})
	}
	f := func(x int) bool {
		vis := make([]bool, n+1)
		q := []int{s}
		vis[s] = true
		for len(q) > 0 {
			u := q[0]
			if u == t {
				return true
			}
			q = q[1:]
			for _, pii := range g[u] {
				if !vis[pii.v] && pii.w <= x {
					vis[pii.v] = true
					q = append(q, pii.v)
				}
			}
		}
		return false
	}
	l, r := 0, 10010
	for l+1 < r {
		mid := (l + r) >> 1
		if f(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	Fprint(out, r)
	out.Flush()
}
