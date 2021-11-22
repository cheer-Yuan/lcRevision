package main

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

