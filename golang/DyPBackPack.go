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