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