package main

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