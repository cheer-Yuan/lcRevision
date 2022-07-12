package main

import "strings"

/*504. 七进制数
给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。

示例 1:
输入: num = 100
输出: "202"

示例 2:
输入: num = -7
输出: "-10"
*/
//func convertToBase7(num int) string {
//	result := ""s
//	acl := 7
//	temp := []int{}
//
//	if num > 0 {
//		for num >= 0 {
//			temp = append(temp, num%acl)
//			num -= acl
//		}
//	}
//	for num != 0 {
//		acl *= 7
//
//	}
//}

/*405. 数字转换为十六进制数got
给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用 补码运算 方法。
注意:
十六进制中所有字母(a-f)都必须是小写。
十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符'0'来表示；对于其他情况，十六进制字符串中的第一个字符将不会是0字符。
给定的数确保在32位有符号整数范围内。
不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。

示例 1：
输入:
26
输出:
"1a"

示例 2：
输入:
-1
输出:
"ffffffff"
*/
func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	result := &strings.Builder{} // 构造string
	for i := 7; i >= 0; i-- {    // 32bit == 4位16进制数
		val := num >> (4 * i) & 0xf      // devide by 位数 * i 再按位和16取与
		if val > 0 || result.Len() > 0 { // 拼接
			var digit byte // 取ascii码
			if val < 10 {
				digit = '0' + byte(val)
			} else {
				digit = 'a' + byte(val-10)
			}
			result.WriteByte(digit)
		}
	}

	return result.String()
}
