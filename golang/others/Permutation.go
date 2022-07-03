package others

import "fmt"

func printFullPermute(s string, start int) {
	var buff string

	if len(s) <= 1 || start == len(s){			//当所有字符交换完成时输出
		fmt.Println(s)
	} else {
		for i := start; i < len(s); i++ {
			buff = swap(s, start, i)			//交换-固定已经交换掉的字符-递归执行之后的字符
			printFullPermute(buff, start+1)
			buff = swap(s, start, i)
		}
	}
}

func printFullPermuteNoRepeat(s string, start int) {
	var buff string

	if len(s) <= 1 || start == len(s){			//当所有字符交换完成时输出
		fmt.Println(s)
	} else {
		for i := start; i < len(s); i++ {
			buff = swap(s, start, i)			//重复的字符不交换
			printFullPermute(buff, start+1)
			buff = swap(s, start, i)
		}
	}
}