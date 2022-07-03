package others

/*
01背包问题：每件物品不可重复放入
一维：
dp[j] ：容量为j的背包最大可能价值
*/

func backPack(Value, Weights []int, Volume int) int {
	dp := make([]int, Volume+1)

	//init：一维初始化为零， 不需显式
	for index, i := range Weights {
		for j := Volume; j >= i; j-- {
			dp[Volume] = MaxOf2(dp[j], dp[j-i]+Value[index])
		}
	}

	return dp[Volume]
}

/*
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意: 每个数组中的元素不会超过 100 数组的大小不会超过 200

示例 1: 输入: [1, 5, 11, 5] 输出: true 解释: 数组可以分割成 [1, 5, 5] 和 [11].

示例 2: 输入: [1, 2, 3, 5] 输出: false 解释: 数组不能分割成两个元素和相等的子集.


转化为背包问题：
Volume = Sum / 2

*/

func canPartition(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}

	volume := sum / 2
	if volume*2 != sum {
		return false
	}
	dp := make([]int, volume+1)

	for _, val := range nums {
		for j := volume; j >= val; j-- {
			dp[j] = MaxOf2(dp[j], dp[j-val]+val)
		}
	}

	if dp[volume] == volume {
		return true
	} else {
		return false
	}
}

/*
有一堆石头，每块石头的重量都是正整数。

每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎； 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。 最后，最多只会剩下一块石头。返回此石头最小的可能重量。如果没有石头剩下，就返回 0。

示例： 输入：[2,7,4,1,8,1] 输出：1 解释： 组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]， 组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]， 组合 2 和 1，得到 1，所以数组转化为 [1,1,1]， 组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。

1 <= stones.length <= 30
1 <= stones[i] <= 100
	可得出max length =  30 * 100 / 2

*/

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, i := range stones {
		sum += i
	}
	leng := sum / 2

	dp := make([]int, leng+1)

	for _, i := range stones {
		for j := leng; j >= i; j-- {
			dp[j] = MaxOf2(dp[j], dp[j-i]+i)
		}
	}

	result := sum - dp[leng]*2

	if result > 0 {
		return result
	} else {
		return -result
	}
}

/*
由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。

题目数据保证答案符合 32 位整数范围。

输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。

思路：
本质上是排列问题 --> 外层背包容量，内层物品价值（如果物品价值在外层：大价值一定出现于小价值后 --> 组合问题）

初始化：对于nums = [1, 2, 3]，j = 1, dp[1] = 1 = O + dp[0] --> dp[0] = 1
*/

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	dp[0] = 1

	for i := 0; i < target; i++ {
		for _, j := range nums {
			//由于物品在内层，用if确保引用>=0
			if i-j >= 0 {
				dp[i] += dp[i-j]
			}
		}
	}

	return dp[target]
}

/*
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

示例：

输入：nums: [1, 1, 1, 1, 1], S: 3
输出：5

数组非空，且长度不会超过 20 。
初始的数组的和不会超过 1000 。
保证返回的最终结果能被 32 位整数存下。


1. 转化为背包问题：
假设加法的总和为x，那么减法对应的总和就是sum - x
所以我们要求的是 x - (sum - x) = S
x = (S + sum) / 2
问题转化为x的组合问题，01背包（每个数只用一次）


思路：

二维：

dp[i][j] 表示在数组 nums 的前 i 个数中选取元素，使得这些元素之和等于 j 的方案数。假设数组 nums 的长度为n，则最终答案为 dp[n][pos]

状态转移：

对于num[i]：
如果j < num[i]：必定不选num[i]：dp[i][j] = dp[i - 1][j]
如果j > num[i]：不选num[i]：dp[i][j] = dp[i - 1][j]，若选：关注概念
i.e. num:[1, 1, 1, 2, 3, 4], i = 4, j = 4，选num[4] ,子集和为4的数量，等于在前3位中选子集和为4-num[4]的数量
得：若选，dp[i][j] = dp[i - 1][j - num[i]]
综上： dp[i][j] = dp[i - 1][j] + dp[i - 1][j - num[i]]

优化至一维：dp[j] += dp[j - nums[i]]

*/

func findTargetSumWays(nums []int, target int) int {
	Sum := 0

	for _, i := range nums {
		Sum += i
	}

	S := (target + Sum) / 2

	if (target+Sum)%2 == 1 || abs(target) > abs(Sum) {
		return 0
	}

	dp := make([]int, S+1)

	dp[0] = 1

	for _, i := range nums {
		for j := S; j >= i; j-- {
			dp[j] += dp[j-i]
		}
	}

	return dp[S]
}

/* 完全bb
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
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, i := range coins {
		for j := i; j <= amount; j++ {
			dp[j] += dp[j-i]
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
	dp := make([]int, amount+1)

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = 10001
	}

	for _, i := range coins {
		for j := i; j <= amount; j++ {
			dp[j] = MinOf2(dp[j], dp[j-i]+1)
		}
	}

	if dp[amount] == 10001 {
		return -1
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
	dp := make([]int, n+1)

	dp[0], dp[1] = 1, 1

	if n > 1 {
		for i := 2; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
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
	dp := make([]int, n+1)

	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = 10001
	}

	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			if j >= i*i {
				dp[j] = MinOf2(dp[j], dp[j-i*i]+1)
			}
		}
	}

	return dp[n]
}
