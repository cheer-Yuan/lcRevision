package main

/*
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

示例 2：

输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"

思路：双指针
*/

func countSubstrings(s string) int {
	lens := len(s)
	var result int

	for i := 0; i < lens ; i++ {
		result += extend(s, i, i)
		result += extend(s, i, i + 1)
	}

	return result
}

func extend(s string, center1, center2 int) int {
	c1, c2 := center1, center2
	result := 0

	for c1 >= 0 && c2 <= len(s) - 1 && s[c1] == s[c2] {
		result++
		c1--
		c2++
	}

	return result
}

/*
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

示例 1：

输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。
示例 2：

输入：s = "cbbd"
输出：2
解释：一个可能的最长回文子序列为 "bb" 。

思路：
dp[i][j] : 在[i, j]范围内最长的回文子序列长度
s[i] == s[j]
	dp[i][j] = dp[i + 1][j - 1] + 2 // 扩展2位
s[i] != s[j]
	dp[i][j] = max(dp[i + 1][j], dp[i][j - 1])
初始化：i = h 的情况

  b b b a b
b 1
b -
b - - 1 cal
a - - - 1 cal
b - - - - 1

*/

func longestPalindromeSubseq(s string) int {
	lens := len(s)

	dp := make([][]int, lens)
	for i := 0; i < lens; i++ {
		dp[i] = make([]int, lens)
		dp[i][i] = 1
	}

	for i := lens - 2; i >= 0; i-- {
		for j := i + 1; j < lens; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i + 1][j - 1] + 2
			} else {
				dp[i][j] = MaxOf2(dp[i + 1][j], dp[i][j - 1])
			}
		}
	}

	return dp[0][lens - 1]
}


/*
给你一个字符串 s，找到 s 中最长的回文子串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

*/

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	lens := len(s)
	max := 0
	c1, c2 := 0, 0
	result := ""

	for i := 0; i < lens ; i++ {
		c1, c2 = extend1(s, i, i)
		if c2 - c1 > max {
			max = c2 - c1
			result = s[c1:c2 + 1]
		}
		c1, c2 = extend1(s, i, i + 1)
		if c2 - c1 > max {
			max = c2 - c1
			result = s[c1:c2 + 1]
		}
	}

	if result == "" {
		result = s[0:1]
	}

	return result
}

func extend1(s string, center1, center2 int) (int, int) {
	c1, c2 := center1, center2

	r1, r2 := 0, 0

	for c1 >= 0 && c2 <= len(s) - 1 && s[c1] == s[c2] {
		r1 = c1
		r2 = c2
		c1--
		c2++
	}

	return r1, r2
}