package main

import (
	"fmt"
	"strings"
)

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
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
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
				dp[i][j] = MaxOf2(dp[i][j], dp[i-M][j-N]+1)
			}

		}

	}

	return dp[m][n]
}

/*数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

有效括号组合需满足：左括号必须以正确的顺序闭合。*/

func generateParenthesis(n int) []string {
	var results [][]string
	results = append(results, []string{""})
	results = append(results, []string{"()"})

	if n == 0 {
		return results[0]
	} else if n == 1 {
		return results[1]
	}

	for i := 2; i <= n; i++ {
		list := []string{}
		for p := 0; p <= i-1; p++ {
			q := i - p - 1
			for _, parantP := range results[p] {
				for _, parantQ := range results[q] {
					result := strings.Join([]string{"(", parantP, ")", parantQ}, "")
					list = append(list, result)
				}
			}
		}
		results = append(results, list)
	}
	return results[n]
}

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

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2], dp[3] = 0, 0, 1, 2

	for i := 4; i <= n; i++ {
		buff := 0
		for j := 2; j < i-1; j++ {
			ifdevide := 0
			if (i - j) > dp[i-j] {
				ifdevide = j * (i - j)
			} else {
				ifdevide = j * dp[i-j]
			}
			if ifdevide > buff {
				buff = ifdevide
			}
		}
		dp[i] = buff
	}

	return dp[n]
}

func ReverseString(s string) string {
	a := []rune(s)

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func SubPanlidrome(s string) string {
	RevS := ReverseString(s)
	Len := len(s)

	var Sub [][]int

	for i := 0; i < Len; i++ {
		a := make([]int, Len)
		Sub = append(Sub, a)
	}

	for i := 0; i < Len; i++ {
		if s[i] == RevS[0] {
			Sub[i][0] = 1
		}
		if s[0] == RevS[i] {
			Sub[0][i] = 1
		}
	}

	fmt.Println(Sub)

	maxL, maxI := 0, 0
	for i := 1; i < Len; i++ {
		for j := 1; j < Len; j++ {
			if s[i] == RevS[j] {
				Sub[i][j] = Sub[i-1][j-1] + 1
				if Sub[i][j] > maxL {
					if Len-j-1+Sub[i][j]-1 == i {
						maxL = Sub[i][j]
						maxI = i
					}
				}
			}
		}
	}

	var Result []string
	for i := 0; i < maxL; i++ {
		Result = append(Result, string(s[maxI-maxL+1+i]))
	}

	if maxL == 0 {
		return string(s[0])
	}

	return strings.Join(Result, "")
}

/*给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。*/

/*
dp[n]：有n个节点的树的种数
dp[n] = sum of f(1) ... f(i) ,,, f(n) ：以i为根的树的种数
f(i) = 左边<i，右边>i = dp[i - 1] * dp[n - i]

dp[3] = 2 + 1 + 2, 2 = dp[1 - 1] * dp[2] = 1 * 2 = 2, dp[0] = dp[1] = 1
*/

func numOfSearchTree(n int) int {
	NumList := make([]int, n+1)

	NumList[0], NumList[1] = 1, 1

	for i := 2; i <= n; i++ {
		f := 0
		for j := 1; j <= i; j++ {
			f += NumList[j-1] * NumList[i-j]
		}
		NumList[i] = f
	}

	return NumList[n]
}
