# 完全背包的概念

给定一个体积为 $v$ 的背包，和体积为 $v_{1}, v_{2}, \cdots, v_{i}$ 的 $i$ 个物品，求总体积不超过这个背包所能获得的方案数\最大值，每个物品有无限个。

# 思路

和[01背包不同](https://github.com/Ra1nowo/Go/blob/main/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92/01%E8%83%8C%E5%8C%85%E9%97%AE%E9%A2%98/neg.md)的是，01背包每个物品只能选一次，而在完全背包下，__每个物品可以重复选择无穷多次__，只要保证不大于背包容量，因此对于每个子问题，我们需要考虑选几个物品，所以可以写出最原始的版本代码如下：

``` go
package main

import "fmt"

func main() {
	const N = 1010
	var n, m int
	v, w := make([]int, N), make([]int, N)
	f := make([][]int, N)
	for i := range f {
		f[i] = make([]int, N)
	}
	fmt.Scan(&n, &m)

	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k*v[i] <= j; k++ {
				f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
			}
		}
	}

	fmt.Println(f[n][m])
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

```

上述代码的时间复杂度是 $O(n^3)$ 的，实际上可以优化这份代码的时间复杂度，优化过程如下：

考虑状态转移方程

$$
f_{i,j} = \max(f_{i-1,j}, f_{i-1,j-v_{i}} + w_{i}, f_{i-1,j-2v_{i}} + 2w_{i}, \cdots)
$$
$$
f_{i,j-v_{i}} = \max(f_{i-1,j-v_{i}}, f_{i-1,j-2v_{i}} + w_{i}, \cdots)
$$

可以发现

$$
f_{i,j} = \max(f_{i-1,j}, f_{i,j-v_{i}} + w_{i})
$$

因此我们可以用这个转移方程来优化时间复杂度，这样就不需要枚举第 $i$ 个物品选几个了

优化代码如下：

``` go
package main

import "fmt"

func main() {
	const N = 1010
	var n, m int
	v, w := make([]int, N), make([]int, N)
	f := make([][]int, N)
	for i := range f {
		f[i] = make([]int, N)
	}
	fmt.Scan(&n, &m)

	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= v[i] {
				f[i][j] = max(f[i-1][j], f[i][j-v[i]]+w[i])

			}
		}
	}

	fmt.Println(f[n][m])
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
```

这样时间复杂度就被降低到 $O(n^2)$ 了