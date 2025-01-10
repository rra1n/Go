# 01背包的概念

给定一个体积为 $v$ 的背包，和体积为 $v_{1}, v_{2}, \cdots, v_{i}$ 的 $i$ 个物品，求总体积不超过这个背包所能获得的方案数\最大值，**每个物品只能选择一次**。

# 思路

01背包每个物品只能选一次，只要保证不大于背包容量，因此对于每个子问题，我们需要考虑这个物品选或者不选，所以可以写出最原始的版本代码如下：

# 状态转移方程

考虑状态转移方程：

$$
f_{i,j} = \max(f_{i-1,j}, f_{i-1, j-v_{i}} + w_{i})
$$

代码如下：

```go
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
                f[i][j] = max(f[i][j], f[i-1][j-v[i]] + w[i])
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

上述代码的时间复杂度为 $O(n^2)$

# 相关习题
- [采药](https://www.luogu.com.cn/problem/P1048)
- [装箱问题](https://www.luogu.com.cn/problem/P1049)
- [和为目标值的最长子序列](https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/)