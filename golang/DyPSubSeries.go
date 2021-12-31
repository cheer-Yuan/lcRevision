package main

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

思路：dp[i]：一定包括i的最长递增子序列长度
	if i <= i - 1: dp[i - 1]
	if i > i - 1: 位置i的最长升序子序列等于j从0到i-1各个位置的最长升序子序列 + 1：max(dp[1]...dp[i]) + 1
		对于每个nums[i]，遍历从0到i - 1
最后取最大的dp[i]
*/

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	dp := make([]int, length)

	dp[0] = 1

	result := 0

	for i := 1; i < length; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = MaxOf2(dp[i], dp[j] + 1)
			}
			if dp[i] > result {
				result = dp[i]
			}
		}
	}

	return result
}

/*
给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。

连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。


示例 1：

输入：nums = [1,3,5,4,7]
输出：3
解释：最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。

贪心：O 1 复杂度

尝试DP
dp[i]：以下标i为结尾的数组的连续递增的子序列长度为dp[i]。
	i <= i - 1 : 1
	i > i - 1 : dp[i] + 1
dp[0] = 1
*/



func findLengthOfLCIS(nums []int) int {
	length := len(nums)
	dp := make([]int, length)

	dp[0] = 1
	result := dp[0]
	for i := 1; i < length; i++ {
		if nums[i - 1] >= nums[i] {
			dp[i] = 1
		} else {
			dp[i] = dp[i - 1] + 1
		}
		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}



/*
最长连续公共子序列：

给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

输入：
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出：3
解释：
长度最长的公共子数组是 [3, 2, 1] 。

思路：
dp[i][j]:以A[i - 1]和B[j - 1] 结尾的最长连续公共子序列
	if A[i - 1] != B[j - 1]：0
				==		   ：dp[i - 1][j - 1] + 1

初始化：0
if A[0] == B[0] : dp[1][1] = 1 = dp[0][0] + 1 = 0 + 1
*/

func findLength(nums1 []int, nums2 []int) int {
	lengthA, lengthB := len(nums1), len(nums2)
	if lengthA == 0 || lengthB == 0 {
		return 0
	}
	dp := make([][]int, lengthA + 1)
	for i := 0; i <= lengthA; i++ {
		dp[i] = make([]int, lengthB + 1)
	}

	result := 0
	for i := 1; i <= lengthA; i++ {
		for j := 1; j <= lengthB; j++ {
			if nums1[i - 1] == nums2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
				if dp[i][j] > result {
					result = dp[i][j]
				}
			}
		}
	}

	return result
}



/*
最长公共子序列

给定两个字符串text1 和text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

示例 1：

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。

思路：
if i == j : 1
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	lengthA, lengthB := len(text1), len(text2)
	if lengthA == 0 || lengthB == 0 {
		return 0
	}
	dp := make([][]int, lengthA + 1)
	for i := 0; i <= lengthA; i++ {
		dp[i] = make([]int, lengthB + 1)
	}

	for i := 1; i <= lengthA; i++ {
		for j := 1; j <= lengthB; j++ {
			if text1[i - 1] == text2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				dp[i][j] = MaxOf2(dp[i][j - 1], dp[i - 1][j])
			}
		}
	}

	return dp[lengthA][lengthB]
}


/*
在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。

现在，可以绘制一些连接两个数字 nums1[i]和 nums2[j]的直线，这些直线需要同时满足满足：

nums1[i] == nums2[j]
且绘制的直线不与任何其他连线（非水平线）相交。
请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。

以这种方法绘制线条，并返回可以绘制的最大连线数。

思路：最长公共子序列问题。
dp[i][j] ：i - 1 & j - 1 最长公共子序列
	if i = j : dpij =
	if i != j : max (i j-1, i-1 j)
初始化： 0
*/

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	len1, len2 := len(nums1), len(nums2)

	dp := make([][]int, len1 + 1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2 + 1)
	}

	for index1 := 1; index1 <= len1; index1++ {
		for index2 := 1; index2 <= len2; index2++ {
			if nums1[index1 - 1] == nums2[index2 - 1] {
				dp[index1][index2] = dp[index1 - 1][index2 - 1] + 1
			} else {
				dp[index1][index2] = MaxOf2(dp[index1 - 1][index2], dp[index1][index2 - 1])
			}
		}
	}

	return dp[len1][len2]
}


/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例: 输入: [-2,1,-3,4,-1,2,1,-5,4] 输出: 6 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

思路：dp[i]：包括i的最大ss和
推导：对于nums[i]：dp[i] =
	加入：dp[i - 1] + nums[i]
	重新计算：nums[i]
初始化：0
返回最大值
*/

func maxSubArray1(nums []int) int {
	len := len(nums)
	dp := make([]int, len)
	dp[0] = nums[0]
	result := dp[0]

	for i := 1; i < len; i++ {
		dp[i] = MaxOf2(dp[i - 1] + nums[i], nums[i])
		if result < dp[i] {
			result = dp[i]
		}
	}

	return result
}


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