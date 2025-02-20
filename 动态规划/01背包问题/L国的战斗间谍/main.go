package main

import . "fmt"

type pair struct{ x, y int }

func (pii pair) add() {
	Println(pii.x + pii.y)
}

func main() {
	pii := pair{2, 3}
	pii.add()
}
