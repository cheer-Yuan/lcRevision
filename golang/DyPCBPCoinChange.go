package main

/*
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。

请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。

假设每一种面额的硬币有无限个。 

题目数据保证结果符合 32 位带符号整数。

思路：硬币无限个->完全背包

dp[j]：凑成总金额j的货币组合数为dp[j]
和的组合问题
dp[j] += dp[j - i]
初始化：设coins = [1, 2, 3] amount = 5
对于coin = 1
dp[1] = 1 = 0 + dp[1 - 1] = dp[0] = 1
*/





func change(amount int, coins []int) int {
	dp := make([]int, amount + 1)
	dp[0] = 1

	for _, i := range coins {
		for j := i; j <= amount; j++ {
			dp[j] += dp[j - i]
		}
	}

	return dp[amount]
}