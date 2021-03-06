package others

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
				dp[i] = MaxOf2(dp[i], dp[j]+1)
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
		if nums[i-1] >= nums[i] {
			dp[i] = 1
		} else {
			dp[i] = dp[i-1] + 1
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
	dp := make([][]int, lengthA+1)
	for i := 0; i <= lengthA; i++ {
		dp[i] = make([]int, lengthB+1)
	}

	result := 0
	for i := 1; i <= lengthA; i++ {
		for j := 1; j <= lengthB; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
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
	dp := make([][]int, lengthA+1)
	for i := 0; i <= lengthA; i++ {
		dp[i] = make([]int, lengthB+1)
	}

	for i := 1; i <= lengthA; i++ {
		for j := 1; j <= lengthB; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = MaxOf2(dp[i][j-1], dp[i-1][j])
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

	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}

	for index1 := 1; index1 <= len1; index1++ {
		for index2 := 1; index2 <= len2; index2++ {
			if nums1[index1-1] == nums2[index2-1] {
				dp[index1][index2] = dp[index1-1][index2-1] + 1
			} else {
				dp[index1][index2] = MaxOf2(dp[index1-1][index2], dp[index1][index2-1])
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
		dp[i] = MaxOf2(dp[i-1]+nums[i], nums[i])
		if result < dp[i] {
			result = dp[i]
		}
	}

	return result
}

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
	dp := make([]bool, leng+1)
	dp[0] = true

	for i := 1; i <= leng; i++ {
		for j := 1; j <= i; j++ {
			if dp[j-1] == false {
				continue
			}

			IfAppear := false
			for _, temp := range wordDict {
				if s[j-1:i] == temp {
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


/*环绕字符串中唯一的子字符串
把字符串 s 看作是“abcdefghijklmnopqrstuvwxyz”的无限环绕字符串，所以s 看起来是这样的：
"...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".
现在给定另一个字符串 p 。返回s 中唯一 的 p 的 非空子串的数量。

输入: p = "a"
输出: 1
解释: 字符串 s 中只有一个"a"子字符。

输入: p = "cac"
输出: 2
解释: 字符串 s 中的字符串“cac”只有两个子串“a”、“c”。.

输入: p = "zab"
输出: 6
解释: 在字符串 s 中有六个子串“z”、“a”、“b”、“za”、“ab”、“zab”。

思路： 周期字符串，首字母+长度->确定子串。
定义dp[a]：p中含有的以a结尾的最长子串长度
*/
func findSubstringInWraproundString(p string) int {
	dp := [26]int{}
	k := 0
	for i, ch := range p {
		if i > 0 && (byte(ch)-p[i-1]+26)%26 == 1 { 		// 从第i位开始寻找以它开头的连续字符串
			k++
		} else {
			k = 1										// 不是连续字符串
		}
		dp[ch-'a'] = max(dp[ch-'a'], k)					// 记录较大值
	}

	result := 0
	for _, v := range dp {
		result += v
	}
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}