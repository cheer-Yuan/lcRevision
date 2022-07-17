package main

/*
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

示例 1:
输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
*/
func translateNum(num int) int {
	var rec func(n int) int // 从后向前递归处理

	rec = func(n int) int {
		if n < 10 {
			return 1
		}
		mod := n % 100
		if mod > 25 || mod < 10 {
			return rec(n / 10) // 最后一位只有一种处理方法
		} else {
			return rec(n/10) + rec(n/100) // 有两种处理方法
		}
	}

	return rec(num)
}
