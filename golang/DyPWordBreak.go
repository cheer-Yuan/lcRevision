package main

/*
给你一个字符串 s 和一个字符串列表 wordDict 作为字典，判定s 是否可以由空格拆分为一个或多个在字典中出现的单词。

说明：拆分时可以重复使用字典中的单词。

示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"

思路：可以分割 --> 装满背包
重复使用：完全背包

dp[j] = dp[j - len(i)] and if j-i ... j 出现在字典中
i < j

初始化： dp[0] = true, else = false
*/

func wordBreak(s string, wordDict []string) bool {
	leng := len(s)
	dp := make([]bool, leng + 1)
	dp[0] = true

	for i := 1; i <= leng; i++ {
		for j := 1; j <= i; j++ {
			if dp[j - 1] == false {
				continue
			}

			IfAppear := false
			for _, temp := range wordDict {
				if s[j - 1 : i] == temp {
					IfAppear = true
					break
				}
			}

			if IfAppear {
				dp[i] = true
				break
			}
		}
	}

	return dp[leng]
}