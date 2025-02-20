package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	var a, b, c int
	in := bufio.NewReader(os.Stdin)   // 创建一个输入缓冲区
	out := bufio.NewWriter(os.Stdout) // 创建一个输出缓冲区
	Fscan(in, &a, &b, &c)             // 将a, b, c输入到in这个缓冲区中
	ans := a + b + c
	Fprint(out, ans) // 将ans写入到out这个输出缓冲区中
	out.Flush()      //需要通过flush将数据输出到控制台中，否则会一直在out这个输出缓冲区中
}
