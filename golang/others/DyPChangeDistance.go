package others

/*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：
如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

思路：
i in len s, j in len t
dp[i][j] ： i-1和j-1的公共子序列长度

双指针贪心：O n
进阶情况优化思路：对于每个字母寻找在目标字符串中的位置是重复动作，考虑建立直接查询关系

dp[i]][j], T字符串中从字母i往后的每个字母，在dict j中出现的最早位置，
i : 字符串T长度 + 1，从后向前遍历  j：字典长度，从前向后遍历
初始化：最后一行，值初始化为T长度

查询s : 先搜索第一个字符s[0]: 考察dp[0][对应s[0]] 的值，
	== 0 false（因为第一行存储了包括T[0]以及之后所有字符在dict中的位置）
	!= 0 --> s[0]出现在T中的第dp位 --> 考察s[1] --> 考察第dp + i行，因为只需要考察单词T中第dp位后的字符
*/

func isSubsequence(s string, t string) bool {
	lenS, lenT := len(s), len(t)
	dp := make([][26]int, lenT + 1)

	//初始化最后一行
	for i := 0; i < 26; i++ {
		dp[lenT][i] = 9999
	}

	//dp初始化矩阵
	for indexT := lenT - 1; indexT >= 0; indexT-- {
		for indexDict := 0; indexDict < 26; indexDict++ {
			if int(t[indexT]) == 97 + indexDict {
				dp[indexT][indexDict] = indexT		//该字母出现： 将值初始化为当前字母下标
			} else {
				dp[indexT][indexDict] = dp[indexT + 1][indexDict]	//未出现，继承下方格子的值
			}
		}
	}

	//查询字符串s
	indexLine := 0
	for indexS := 0; indexS < lenS; indexS++ {
		if dp[indexLine][s[indexS] - 97] == 9999 {
			return false
		} else {
			indexLine = dp[indexLine][s[indexS] - 97] + 1
		}
	}

	return true
}



/*
给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。

字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE"是"ABCDE"的一个子序列，而"AEC"不是）

题目数据保证答案符合 32 位带符号整数范围。

输入：s = "rabbbit", t = "rabbit"
输出：3

输入：s = "babgbag", t = "bag"
输出：5

思路：类似上题，i in t, j in s
dp[i][j]：前 i 个字符的 s 子串中，出现前 j 个字符的 t 子串的次数

画表：
  0 b a b g b a g
0 1 1 1 1 1 1 1 1 // 空为任意子串
b 0 1 1 2 2 3 3 3
a 0 0 1 1 1 1 4 4
g 0 0 0 0 1 1 1 5

🔺分析两种情况
s[i - 1] != t[j - 1]：
	不进行匹配，dp[i][j] = dp[i][j - 1]
s[i - 1] = t[j - 1]：
	1. 不使用 i - 1 匹配，同上一种情况：dp[i][j] = dp[i][j - 1]
	2. 使用 i - 1 匹配，加上dp[i - 1][j - 1]的值
	综上，dp[i][j] = dp[i][j - 1] + dp[i - 1][j - 1]

*/

func numDistinct(s string, t string) int {
	lenS, lenT := len(s), len(t)
	dp := make([][]int, lenT + 1)
	for i := 0; i <= lenT; i++ {
		dp[i] = make([]int, lenS + 1)
	}

	//初始化
	for j := 0; j <= lenS; j++ {
		dp[0][j] = 1
	}

	for i := 1; i <= lenT; i++ {
		for j := i; j <= lenS; j++ {
			if s[j - 1] != t[i - 1] {
				dp[i][j] = dp[i][j - 1]
			} else {
				dp[i][j] = dp[i][j - 1] + dp[i - 1][j - 1]
			}
		}
	}

	return dp[lenT][lenS]
}

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
				dp[i][j] = MinOf2(dp[i - 1][j] + 1, MinOf2(dp[i][j - 1] + 1, dp[i - 1][j - 1] + 1))
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
i 1 1 2 3 4 5 6 6 7 8 //
n 2 2 2 3 4 5 6 6 7 7 // != : add ->  : j > i/ change -> dp[i][j-1] j <= i /
t 3 3 3 3 4 5 5 6 7 8
e 4 3 4
n 5
t 6
i 7
o 8
n 9

w1 i == w2 j : 不操作 	dp[i][j] = dp[i - 1][j - 1]
w1 i != w2 j :
	删除: dp[i][j] = dp[i - 1][j] + 1
	增加 == 删除w2一个字符 ：dp[i][j-1]+1
	替换 == w1和w2都删除一个字符：dp[i-1][j-1]+1
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
				dp[i][j] = MinOf2(dp[i - 1][j] + 1, MinOf2(dp[i][j - 1] + 1, dp[i - 1][j - 1] + 2))
			}
		}
	}

	return dp[len1][len2]
}