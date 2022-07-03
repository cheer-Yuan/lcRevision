package others

import (
	"sort"
)

/*火柴拼正方形
你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i个火柴棒的长度。你要用 所有的火柴棍拼成一个正方形。你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。
如果你能使这个正方形，则返回 true ，否则返回 false 。
*/
func makesquare(matchsticks []int) bool {
	// if总长度不能mod4：false
	totalLen := 0
	for _, i := range matchsticks {
		totalLen += i
	}
	if totalLen % 4 != 0 {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))		// 从小到大排序
	edges := [4]int{}
	var dfs func(int) bool									// param :
	dfs = func(i int) bool {
		if i == len(matchsticks) {
			return true										// 所有火柴棒已经放完
		}

		for index := range edges {						// 用这一只火柴尝试所有边：
			edges[index] += matchsticks[i]					// 将某根火柴放入某个边
			if edges[index]	<= totalLen / 4 && dfs(i + 1) {	// 如果可以放入，dfs下一支火柴，也尝试所有边（dfs）
				return true										// 所有火柴都可以放入，true
			}
			edges[index] -= matchsticks[i]				// 从这边一条边拿掉，尝试下一条边
		}

		return false									// 这一直火柴不适合任意一条边，false
	}

	return dfs(matchsticks[0])
}