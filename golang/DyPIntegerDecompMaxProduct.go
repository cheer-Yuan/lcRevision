package main

/*
题设：

给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1: 输入: 2 输出: 1

示例 2: 输入: 10 输出: 36 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。 说明: 你可以假设 n 不小于 2 且不大于 58。
*/

/*
数学方式：
考虑对于任意一个因数 f >= 4, 可以拆成 2 和 f - 2，有 2(f-2) = 2f - 4 >= f, 于是仅剩3和2备选。又有3*3>2*2*2，于是优先拆成最多个3
 */


/*
动规
dp[i] = 乘积，考虑一维数组，
两种情况：i * (i - j)， i * dp[i - j]
*/

func integerBreak(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	case 4:
		return 4
	}

	dp := make([]int, n + 1)
	dp[0], dp[1], dp[2], dp[3] = 0, 0, 1, 2

	for i := 4; i <= n; i++ {
		buff := 0
		for j := 2; j < i - 1; j++ {
			ifdevide := 0
			if (i - j) > dp[i - j] {
				ifdevide = j * (i - j)
			} else {
				ifdevide = j * dp[i - j]
			}
			if ifdevide > buff {
				buff = ifdevide
			}
		}
		dp[i] = buff
 	}

 	return dp[n]
}