package main

import "math"

/*7. 整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围[−231, 231− 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。

判断反转后的数字是否超过 3232 位有符号整数的范围 : −2 ^ 31 ≤ rev ⋅ 10 + digit ≤ 2 ^ 31 −1 若该不等式不成立则返回 0。
INT_MAX = 2147483647 = INT_MAX / 10 * 10 + 7
所以比较 rev 与 INTMAX / 10，digit 与 7 之间的大小关系
*/
func reverse(x int) int {
	rev, digit := 0, 0

	for x != 0 {
		if rev > math.MaxInt32/10 || rev < math.MinInt32/10 {
			return 0
		}

		digit = x % 10 // 取尾数
		x /= 10        // 准备下一个尾数
		rev *= 10      // 扩大一位
		rev += digit   // 加上尾数
	}

	return rev
}
