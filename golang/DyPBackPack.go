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
