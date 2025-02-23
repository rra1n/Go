package main

import (
	"bufio"
	. "fmt"
	"os"
)

func getNext(t string) []int {
	m, i, j := len(t), 0, -1
	next := make([]int, m)
	next[0] = -1
	for i < m-1 {
		if j == -1 || t[i] == t[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func kmp(s, t string) bool {
	next := getNext(t)
	n, m, i, j := len(s), len(t), 0, 0
	for i < n && j < m {
		if j == -1 || s[i] == t[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	return j == m
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t string
	Fscan(in, &t)
	next := getNext(t)
	for _, x := range next {
		Fprintln(out, x)
	}
}
