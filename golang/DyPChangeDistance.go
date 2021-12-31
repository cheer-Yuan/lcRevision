package main


/*
给定两个单词word1和word2，找到使得word1和word2相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。

示例：

输入: "sea", "eat"
输出: 2
解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"

思路：dp[i][j]：前 i 位 w1 与 前 j 位 w2 变化为相同所需的步骤数

🔺分析两种情况
s[i - 1] == t[j - 1]：se vs e : 等于 s vs null：
	dp[i][j] = dp[i - 1][j - 1]
s[i - 1] != t[j - 1]：想象此时的实际操作：删除 i - 1 or j - 1 or 都删
	dp[i][j] = min(dp[i - 1][j] + 1, dp[i][j - 1] + 1, dp[i - 1][j - 1] + 2)

  0 e a t
0 0 1 2 3
s 1 2 3 4
e 2 1 2 3
a 3 2 1 2
*/

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1 + 1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2 + 1)
	}

	for i := 1; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= len2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i - 1] == word2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[]
			}
		}
	}

}