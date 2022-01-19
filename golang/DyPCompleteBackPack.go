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


/*

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

dp[i]： 爬到i层楼梯的方法数 = 爬到 i - 1 层的方法数（再爬一层） + 爬到 i - 2 层的方法数（再爬2层）
dp[i] = dp[i - 1] + dp[i - 2]
初始化：dp[2] = 2 = dp[1] + dp[0] = 1 + dp[0] , dp[0] = 1
*/

func climbStairs(n int) int {
	dp := make([]int, n + 1)

	dp[0], dp[1] = 1, 1

	if n > 1 {
		for i := 2; i <= n; i++ {
			dp[i] = dp[i - 1] + dp[i - 2]
		}
	}

	return dp[n]
}


/*
一步一个台阶，两个台阶，三个台阶，.......，直到 m个台阶。问有多少种不同的方法可以爬到楼顶

思路：
物品价值=步数，可无限重复，考虑完全背包 --> 内循环从前向后遍历
dp[i] = 对于一步某个及以下的台阶数，可以爬到i层的方法数
排列问题：外层背包，内层价值

*/


/*
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

示例 1： 输入：n = 12 输出：3 解释：12 = 4 + 4 + 4

示例 2： 输入：n = 13 输出：2 解释：13 = 4 + 9

提示：

1 <= n <= 10^4

思路：
物品：完全平方数。可重复 --> 完全背包
dp[j]：对于完全平方数<=i，组成j的最小的个数
推理同CoinChange
*/

func numSquares(n int) int {
	dp := make([]int, n + 1)

	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = 10001
	}

	for i := 1; i * i <= n; i++ {
		for j := i * i; j <= n; j++ {
			if j >= i * i {
				dp[j] = MinOf2(dp[j], dp[j - i * i] + 1)
			}
		}
	}

	return dp[n]
}
