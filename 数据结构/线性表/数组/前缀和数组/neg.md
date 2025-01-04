# 前缀和的定义

设有数

$$
    a_{1}, a_{2}, a_{3}, \cdots, a_{n}
$$

定义

$$
a_{1} + a_{2} + a_{3} + \cdots + a_{i}
$$

为 $s_{i}$，称这样的式子为前缀和

# 前缀和的性质

利用前缀和可以快速求出某一段连续子数组的和

$$
a_{i} + a_{i+1} + \cdots + a_{i+k} = s_{i + k} - s_{i-1}  
$$


# 涉及前缀和的相关题目
- [特殊数组II](https://leetcode.cn/problems/special-array-ii/description/)
- [货舱选址](https://www.acwing.com/problem/content/description/106/) 也可以用中位数贪心解决
