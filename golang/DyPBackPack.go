package main

/*
01背包问题：每件物品不可重复放入
一维：
dp[j] ：容量为j的背包最大可能价值
*/

func backPack(Value, Weights []int, Volume int) int {
	dp := make([]int, Volume + 1)

	//init：一维初始化为零， 不需显式
	for index, i := range Weights {
		for j := Volume; j >= i; j-- {
			dp[Volume] = MaxOf2(dp[j], dp[j - i] + Value[index])
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
	if volume * 2 != sum {
		return false
	}
	dp := make([]int, volume + 1)


	for _, val := range nums {
		for j := volume; j >= val; j-- {
			dp[j] = MaxOf2(dp[j], dp[j - val] + val)
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

	dp := make([]int, leng + 1)


	for _, i := range stones {
		for j := leng; j >= i; j-- {
			dp[j] = MaxOf2(dp[j], dp[j - i] + i)
		}
	}

	result := sum - dp[leng] * 2

	if result > 0 {
		return result
	} else {
		return - result
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
	dp := make([]int, target + 1)

	dp[0] = 1

	for i := 0; i < target; i++ {
		for _, j := range nums {
			//由于物品在内层，用if确保引用>=0
			if i - j >= 0 {
				dp[i] += dp[i - j]
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

	if (target + Sum) % 2 == 1 || abs(target) > abs(Sum) {
		return 0
	}

	dp := make([]int, S + 1)

	dp[0] = 1

	for _, i := range nums {
		for j := S; j >= i; j-- {
			dp[j] += dp[j - i]
		}
	}

	return dp[S]
}
