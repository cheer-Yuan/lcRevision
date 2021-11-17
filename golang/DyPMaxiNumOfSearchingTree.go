package main

/*给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。*/

/*
dp[n]：有n个节点的树的种数
dp[n] = sum of f(1) ... f(i) ,,, f(n) ：以i为根的树的种数
f(i) = 左边<i，右边>i = dp[i - 1] * dp[n - i]

dp[3] = 2 + 1 + 2, 2 = dp[1 - 1] * dp[2] = 1 * 2 = 2, dp[0] = dp[1] = 1
*/

func numOfSearchTree(n int) int {
	NumList := make([]int, n + 1)

	NumList[0], NumList[1] = 1, 1

	for i := 2; i <= n; i++ {
		f := 0
		for j := 1; j <= i; j++ {
			f += NumList[j - 1] * NumList[i - j]
		}
		NumList[i] = f
	}

	return NumList[n]
}