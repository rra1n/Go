package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	var n, m int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	cnt := make(map[int]int)
	i, cost := 1, n+1
	x, y := 0, 0

	for j := 1; j <= n; j++ {
		cnt[a[j]]++
		for len(cnt) == m {
			if j-i+1 < cost {
				cost = j - i + 1
				x, y = i, j
			}
			cnt[a[i]]--
			if cnt[a[i]] == 0 {
				delete(cnt, a[i])
			}
			i++
		}
	}

	if x == 0 && y == 0 {
		Fprint(out, "0 0")
	} else {
		Fprintf(out, "%d %d", x, y)
	}
}
