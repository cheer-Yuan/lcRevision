package others

import "sort"

/* 组合
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

回溯：抽象成树
1 -> 2 : 1, 2 叶子
  -> 3 : 1, 3 叶子
	...
  -> n : 1, n 叶子
2 -> 3 : 2, 3 叶子
	...

		n-1, n 叶子

参数：n，k，起点
*/

func combine(n int, k int) [][]int {
	results := [][]int{}
	path := []int{}

	var traverse func(n, k, start int)
	traverse = func(n, k, start int) {
		if len(path) == k { // 取到n个数：结束本次回溯
			buff := make([]int, k)
			copy(buff, path) // 注意引用和值
			results = append(results, buff)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i) //加入路径
			traverse(n, k, i+1)
			path = path[0 : len(path)-1] // 删除处理过的路径
		}
	}

	traverse(n, k, 1)
	return results
}

/*
找出所有相加之和为n 的k个数的组合。组合中只允许含有 1 -9 的正整数，并且每种组合中不存在重复的数字。

说明：
所有数字都是正整数。
解集不能包含重复的组合。
*/

func combinationSum3(k int, n int) [][]int {

}


