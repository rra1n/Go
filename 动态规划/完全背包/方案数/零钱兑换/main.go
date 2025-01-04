package main

func change(amount int, coins []int) int {
	f := make([][]int, len(coins)+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}

	f[0][0] = 1
	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			f[i][j] = f[i-1][j]
			if j >= coins[i-1] {
				f[i][j] += f[i][j-coins[i-1]]
			}
		}
	}
	return f[len(coins)][amount]
}
