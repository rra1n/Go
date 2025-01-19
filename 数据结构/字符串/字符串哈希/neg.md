# 字符串哈希

字符串哈希是一种 __前缀哈希__，例如我们有一个字符串 "ACCCCLJY"，我们可以利用字符串哈希的方式取出某个前缀所对应的 $key$ ，比如求出"ACC"的前缀值和"ACCCCLJ" 的前缀值。

怎么求前缀的哈希值呢？

我们将字符串看作是一个 $p$ 进制的数，一般我们取 $p=131$ ，比如将字符串"ABS"看作一个 $p$ 进制的数，那么这个字符串的哈希值为 

$$
A * P^1 + B * P^2 + C * P^3 
$$

其中 $ABC$ 是这些字符的 $ASCII$ 值 

一般地，一个长度为 $n$ 的字符串 $s$ 的前缀哈希为（这里下标从1开始）

$$
s_{0}*p^1+s_{1}*p^2 + \cdots + s_{n-1}*p^n
$$

如果我们想要求出某一段 $\left[l,r\right]$ 的 $key$ 应该怎么做呢？

__利用前缀和的思想__，我们可以先求 $\left[0,l-1\right]$ 的前缀哈希和 $\left[0,r\right]$ 的前缀哈希。

那么是不是直接相减就行呢？ 

__其实并不是的__

$s$ 在 $\left[0,l-1\right]$ 中的哈希值是

$$
s_{0}p^{l-1}+s_{1}p^{l-2}+ \cdots +s_{l-1}p^{0}
$$

而 $s$ 在 $\left[0,r\right]$ 中的哈希值是

$$
s_{0}p^{r}+s_{1}p^{r-1}+ \cdots + s_{l-1}p^{r-l+1} + \cdots + s_{r}p^{0}
$$

要求的子串的范围是 $\left[l,r\right]$，对应的哈希值为

$$
s_{l}p^{l}+s_{l+1}p^{l+1}+ \cdots +s_{r}p^{r}
$$

对比发现，要求的子串的范围是 $\left[l,r\right]$，对应的哈希值应该为

$$
h\left[l,r\right] = h\left[0,r\right] - h\left[0,l-1\right]*p^{r-l+1}
$$

我们可以拿两个数组来记录 $p$ 的幂次方和每个前缀的哈希值

模板如下

``` go 
func strStr(s string, t string) int {
    n := len(s)
    m := len(t)
    if n < m {
       return -1
    }
    const P = 131
    const MOD = 1e9 + 7
    // 计算前缀哈希值
    h := make([]uint64, n+1)
    p := make([]uint64, n+1)
    p[0] = 1
    for i := 0; i < n; i++ {
       h[i+1] = (h[i]*uint64(P) + uint64(s[i])) % uint64(MOD)
       p[i+1] = p[i] * uint64(P) % uint64(MOD)
    }
    // 计算目标字符串的哈希值
    var ht uint64
    for i := 0; i < m; i++ {
       ht = (ht*uint64(P) + uint64(t[i])) % uint64(MOD)
    }
    for i := 0; i <= n-m; i++ {
       st := (h[i+m] - h[i]*p[m]%uint64(MOD) + uint64(MOD)) % uint64(MOD)
       if st == ht {
          return i
       }
    }
    return -1
}


```