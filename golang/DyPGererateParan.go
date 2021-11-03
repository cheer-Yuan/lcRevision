package main

import (
	"strings"
)

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
		for p := 0; p <= i - 1; p++ {
			q := i - p - 1
			for _, parantP := range  results[p]{
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