# 同余的概念

设有两个整数 $x$ , $y$ 和一个除数 $p$，若满足以下关系：

$$
x \equiv y \ (mod \ p)
$$

则称 $x$ 和 $y$ 同余。

# 同余的一般推导

若 $x$ 和 $y$ 均大于0，那么前面所写的同余表示形式是正确的。然而，通常余数是不能为负数的，所以一般情况下我们写作：

$$
(x \ mod \ p) + p = (y \ mod\ p)
$$

这样一来，即便 $x$ 为负数时，通过此式计算得出的余数依然是非负数。

# 同余的其他性质

1. 设有数列：

$$
a_{1}, a_{2}, a_{3}, \cdots, a_{n}
$$

记：

$$
s_{i} = a_{1} + a_{2} + a_{3} + \cdots + a_{i - 1}
$$

这里的 $s_{i}$ 表示数列 $a_{n}$ 的下标的从1到 $i - 1$ 项的和。

在这些数字当中找出连续的若干个数，使得这些数的和 $s$ 除以整数 $p$ 后的余数为 $mod$，其充要条件是：

$$
(s_{i} \ mod \ p) - (s_{j} \ mod \ p) = mod
$$

简单来讲：

区间 $\left[l,r - 1\right]$ 的所有数的和对 $p$ 取模后得到余数 $mod$ 等价于 $\left[0,r - 1\right]$ 的所有数的和 $s_{r}$ 对 $p$ 取模后得到的余数减去 $\left[0,l - 1\right]$ 的 $s_{l}$ 对 $p$ 取模后得到的余数。

特殊地：

我们通常会考虑能被整除的相关问题。

# 涉及同余的相关题目

## 和能被整除

### 当存在负数时要用一般的取模公式，和前缀和结合时，一般考虑用哈希表的键存储余数，而值存储需要用到的数

- [和可以被k整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/description/)
- [使得数组和能被p整除](https://leetcode.cn/problems/make-sum-divisible-by-p/description/)
- [连续的子数组和](https://leetcode.cn/problems/continuous-subarray-sum/)

### 长度能被整除 