package main

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