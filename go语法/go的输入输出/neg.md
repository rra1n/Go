# Go的输入输出

Go语言配备了标准库 __fmt__ 来进行输入输出。

常见的 __输出__ 方式有

``` go
package main

import (
	. "fmt" // 前面加点可以省略fmt这个前缀
)

func main() {
	Print("hello") // 输出不换行
	Print("\n")    // 输出换行符
	Println("hello") // 输出并且换行
	Printf("你好，这个数字是%d\n", 3) // 格式化输出，并且换行
}

```

常见的 __输入__ 方式

```go
package main

import (
	. "fmt"
)

func main() {
	var s string
	Scanln(&s)
	Scan(&s)
	Scanf("%s", &s)
	Println(s)
}
```

上述的输入输入方式会反复调用底层系统，当数据量较大的时候，需要时间会比较久，__可以用bufio库中的输入输出来减少时间__

使用bufio的输入输出示例如下:

```go
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
	out.Flush()      //需要通过flush将数据输出到控制台中，否则会一直在out这个输出缓冲区中。// 也可以在定义好out后，defer out.flush
}
```

类似于 __fmt__ 中的输入输出方式，同样有 $Fscanln，Fscan，Fscanf$ 等输入方式和 $Fprintln，Fprint，Fprinf$ 这些输出方式