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

func minDistance1(word1 string, word2 string) int {
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
				dp[i][j] = MinOf2(dp[i - 1][j] + 1, MinOf2(dp[i][j - 1] + 1, dp[i - 1][j - 1] + 2))
			}
		}
	}

	return dp[len1][len2]
}

/*
给你两个单词word1 和word2，请你计算出将word1转换成word2 所使用的最少操作数。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

思路：
dp[i][j]：前i位 in w1 to 前j位 in w2 的最少步骤

  0 e x e c u t i o n
0 0 1 2 3 4 5 6 7 8 9
i 1 1 2 3 4 5 6 6 7 8 // w1 i == w2 j : 不操作 	dp[i][j] = dp[i - 1][j - 1]
n 2 2 2 3 4 5 6 6 7 7 // != : add -> dp[i][j-1]+1 : j > i/ change -> dp[i][j-1] j <= i /
t 3 3 3 3 4 5 5 6 7 8
e 4 3 4
n 5
t 6
i 7
o 8
n 9

*/


func minDistance(word1 string, word2 string) int {

}