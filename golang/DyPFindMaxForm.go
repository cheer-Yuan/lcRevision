package main

/*
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。

请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。

如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

示例 1：

输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3 输出：4

解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。 其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。

示例 2： 输入：strs = ["10", "0", "1"], m = 1, n = 1 输出：2 解释：最大的子集是 {"0", "1"} ，所以答案是 2 。



思路：本题中strs 数组里的元素就是物品，每个物品都是一个。而m 和 n相当于是一个背包，两个维度的背包。

dp[i][j]: 最多有i个0和j个1的strs的最大子集的大小
推导： 遍历每一个string元素，统计该元素的1以及0的数量 然后更新dp数组。对于某个有M个0和N个1的str，dp[i][j] = dp[i - M][j - N] + 1
选or不选：取最大值
综上，dp[i][j] = max(dp[i][j], dp[i - M][j - N] + 1)

*/




func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m + 1)
	for i, _ := range dp {
		dp[i] = make([]int, n + 1)
	}

	for _, str := range strs {
		M, N := 0, 0

		// count 0s and 1s
		for _, i := range str {
			if i == '0' {
				M++
			} else if i == '1' {
				N++
			}
		}

		for i := m; i >= M; i-- {
			for j := n; j >= N; j-- {
				dp[i][j] = MaxOf2(dp[i][j], dp[i - M][j - N] + 1)
			}

		}

	}

	return dp[m][n]
}