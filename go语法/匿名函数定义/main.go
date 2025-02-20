package main

import . "fmt"

func add(x, y int) int { // 定义了一个加法函数，传入两个形式参数，返回类型为整数
	return x + y
}

func main() {
	f1 := func(x int) int { // 定义一个匿名函数，函数名字为f1，可以在main函数中反复调用，但是不能在其他函数中调用
		if x < 0 {
			x = -x
		}
		return x
	}
	f2 := func(x int) int { // 定义了一个匿名函数，这里的()表示即刻使用，传入的是2，返回值传入f2中，如果没有返回值，可以不用声明f2
		return -x
	}(2)
	func() { // 定义了一个匿名函数，没有返回值，立即执行
		Println("hello")
	}()
	Println(f2)
	Print(f1(-2))
}
