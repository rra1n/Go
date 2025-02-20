// https://atcoder.jp/contests/abc218/tasks/abc218_f
package main

import (
	"bufio"
	. "fmt"
	"os"
)

type edge struct{ to, i int }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, s, t int
	Fscan(in, &n, &m)

	g := make([][]edge, n+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &s, &t)
		g[s] = append(g[s], edge{t, i})
	}

	from := make([]edge, n+1)
	dis := make([]int, n+1)

	bfs := func(ban int) int {
		for i := range dis {
			dis[i] = -1
		}
		dis[1] = 0
		q := []int{1}

		for len(q) > 0 {
			v := q[0]
			q = q[1:]

			for _, e := range g[v] {
				w := e.to
				if e.i != ban && dis[w] < 0 {
					dis[w] = dis[v] + 1
					if ban < 0 {
						from[w] = edge{v, e.i}
					}
					q = append(q, w)
				}
			}
		}
		return dis[n]
	}

	res := bfs(-1)
	ans := make([]int, m+1)
	for i := range ans {
		ans[i] = res
	}

	if res != -1 {
		for v := n; v != 1; v = from[v].to {
			ans[from[v].i] = bfs(from[v].i)
		}
	}

	for i := 1; i <= m; i++ {
		Fprintln(out, ans[i])
	}
}
