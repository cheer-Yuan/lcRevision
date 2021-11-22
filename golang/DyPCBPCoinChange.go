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


/*
力扣题目链接(opens new window)

给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

你可以认为每种硬币的数量是无限的。

示例 1： 输入：coins = [1, 2, 5], amount = 11 输出：3 解释：11 = 5 + 5 + 1

示例 2： 输入：coins = [2], amount = 3 输出：-1

示例 3： 输入：coins = [1], amount = 0 输出：0

示例 4： 输入：coins = [1], amount = 1 输出：1

示例 5： 输入：coins = [1], amount = 2 输出：2

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 10^4

思路：硬币无限-->完全背包
经典背包问题

dp[j]：最少用x个硬币
对于coin[i]：考虑：
	用i，dp[j] = dp[j]
	不用i，dp[j] = dp[j - i] + 1

dp[j] = min( , )

初始化：dp = max amount，dp[0] = 0 用于递推
*/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = 10001
	}

	for _, i := range coins {
		for j := i; j <= amount; j++ {
			dp[j] = MinOf2(dp[j], dp[j - i] + 1)
		}
	}

	if dp[amount] == 10001 {
		return  -1
	}

	return dp[amount]
}

