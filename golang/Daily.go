package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*539. 最小时间差
给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
*/
func findMinDifference(timePoints []string) int {
	leng := len(timePoints)
	if leng > 60*24 { // 必有重复时间
		return 0
	}

	hash := map[int]bool{}

	var converse func(time string) int // 时间转换
	converse = func(time string) int {
		b1, b2, b3, b4 := time[0]-'0', time[1]-'0', time[3]-'0', time[4]-'0'

		return int(b1)*10*60 + int(b2)*60 + int(b3)*10 + int(b4)
	}

	for _, i := range timePoints { // 初始化hash表
		key := converse(i)
		if hash[key] == true {
			return 0
		} else {
			hash[key] = true
		}
	}

	result, temp, earliestTime := 1440, 0, 1440 // 初始化最小时间间隔
	for i := 0; i < 1440; i++ {
		if hash[i] == true {

			if result == 1440 { // 第一个时间
				temp = i
				result--
				earliestTime = i
				continue
			}

			if earliestTime > i { // 更新最早的时间
				earliestTime = i
			}
			difference := i - temp // 用最早时间计算环形时间差
			difference2 := 1440 - i + earliestTime
			if difference2 < difference {
				difference = difference2
			}

			if difference < result {
				result = difference
			}

			temp = i
		}
	}
	return result
}

/* 随机数索引
给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。
注意：
数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。

int[] nums = new int[] {1,2,3,3,3};
Solution solution = new Solution(nums);
// pick(3) 应该返回索引 2,3 或者 4。每个索引的返回概率应该相等。
solution.pick(3);
// pick(1) 应该返回 0。因为只有nums[0]等于1。
solution.pick(1);

遍历nums，当我们第 i 次遇到值为target 的元素时，随机选择区间 [0,i) 内的一个整数，如果其等于 0，则将返回值置为该元素的下标，否则返回值不变。
*/

type Solution []int

func Constructor(nums []int) Solution {
	return nums
}

func (this Solution) Pick(target int) int {
	count, result := 0, 0
	for index, i := range this {
		if i == target {
			count++
			if rand.Intn(count) == 0 {
				result = index
			}
		}
	}
	return result
}

/*二进制间距
给定一个正整数 n，找到并返回 n 的二进制表示中两个 相邻 1 之间的 最长距离 。如果不存在两个相邻的 1，返回 0 。
如果只有 0 将两个 1 分隔开（可能不存在 0 ），则认为这两个 1 彼此 相邻 。两个 1 之间的距离是它们的二进制表示中位置的绝对差。例如，"1001" 中的两个 1 的距离为 3 。
*/

func binaryGap(n int) int {
	pow, accumulat := 0, 1
	for accumulat < n {
		accumulat = accumulat * 2
		pow++
	}

	if accumulat == n {
		return 0
	}

	ifContinue := true
	accumulat /= 2
	rest, dist, maxDist := n-accumulat, 1, 0
	for ifContinue {
		fmt.Println(rest, accumulat)
		accumulat /= 2
		if rest >= accumulat {
			maxDist = MaxOf2(dist, maxDist)
			dist = 1
			rest -= accumulat
		} else {
			dist++
		}
		if rest <= 0 {
			ifContinue = false
		}
	}
	return maxDist
}

/*三维形体投影面积
在n x n的网格grid中，我们放置了一些与 x，y，z 三轴对齐的1 x 1 x 1立方体。
每个值v = grid[i][j]表示 v个正方体叠放在单元格(i, j)上。
现在，我们查看这些立方体在 xy、yz和 zx平面上的投影。
投影就像影子，将 三维 形体映射到一个 二维 平面上。从顶部、前面和侧面看立方体时，我们会看到“影子”。
返回 所有三个投影的总面积 。

row and col max
*/

func projectionArea(grid [][]int) int {
	n, result := len(grid), 0
	maxRow, maxCol := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				result++
			}
			if grid[i][j] > maxRow[i] {
				maxRow[i] = grid[i][j]
			}
			if grid[i][j] > maxCol[j] {
				maxCol[j] = grid[i][j]
			}
		}
	}

	for i := 0; i < n; i++ {
		result += maxRow[i]
		result += maxCol[i]
	}

	return result
}

/*安装栅栏
在一个二维的花园中，有一些用 (x, y) 坐标表示的树。由于安装费用十分昂贵，你的任务是先用最短的绳子围起所有的树。只有当所有的树都被绳子包围时，花园才能围好栅栏。你需要找到正好位于栅栏边界上的树的坐标。
输入: [[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]
输出: [[1,1],[2,0],[4,2],[3,3],[2,4]]

Graham 算法：
我们还需要考虑另一种重要的情况，如果共线的点在凸壳的最后一条边上，我们需要从距离初始点最远的点开始考虑起。所以在将数组排序后，我们从尾开始遍历有序数组并将共线且朝有序数组尾部的点反转顺序，因为这些点是形成凸壳过程中尾部的点，所以在经过了这些处理以后，我们得到了求凸壳时正确的点的顺序。
现在我们从有序数组最开始两个点开始考虑。我们将这条线上的点放入栈中。然后我们从第三个点开始遍历有序数组trees。如果当前点与栈顶的点相比前一条线是一个「左拐」或者是同一条线段上，我们都将当前点添加到栈顶，表示这个点暂时被添加到凸壳上。
检查左拐或者右拐使用的还是 cross 函数。对于向量pq ,qr，计算向量的叉积 (p,q,r)=pq × qr ，如果叉积小于 0，可以知道向量 pq ,qr 顺时针旋转，则此时向右拐；如果叉积大于 0，可以知道向量 pq ,qr  逆时针旋转，表示是左拐；如果叉积等于 00，则 p,q,rp,q,r 在同一条直线上。
如果当前点与上一条线之间的关系是右拐的，说明上一个点不应该被包括在凸壳里，因为它在边界的里面（正如动画中点 44），所以我们将它从栈中弹出并考虑倒数第二条线的方向。重复这一过程，弹栈的操作会一直进行，直到我们当前点在凸壳中出现了右拐。这表示这时凸壳中只包括边界上的点而不包括边界以内的点。在所有点被遍历了一遍以后，栈中的点就是构成凸壳的点。
*/

func outerTrees(trees [][]int) [][]int {
	length := len(trees)
	if length < 4 {
		return trees
	}

	// 找到 y 最小的点, put at trees[0]， 我们可以肯定它一定在凸包上
	minY := 0
	for i, point := range trees {
		if point[1] < trees[minY][1] {
			minY = i
		}
	}
	trees[0], trees[minY] = trees[minY], trees[0]

	// 按照极坐标的角度大小进行排序, 极角顺序更小的点排在数组的前面。如果有两个点相对于第一个点的极角大小相同，则按照与点 bottom 的距离排序。
	amputedTree := trees[1:]
	sort.Slice(amputedTree, func(i, j int) bool {
		a, b := amputedTree[i], amputedTree[j]
		diff := cross(trees[0], a, b)
		return diff > 0 || diff == 0 && distance2D(trees[0], a) < distance2D(trees[0], b)
	})

	// 对于凸包最后且在同一条直线的元素按照距离从大到小进行排序 (如果共线的点在凸壳的最后一条边上，我们需要从距离初始点最远的点开始考虑起。)
	commonPoint := length - 1
	for commonPoint > 0 && cross(trees[0], trees[commonPoint], trees[length-1]) == 0 {
		commonPoint-- //check how many points are on the same vector
	}
	for i, j := commonPoint+1, length-1; i < j; i++ {
		trees[i], trees[j] = trees[j], trees[i] // change smallTObig to bigTOsmall
	}

	// 如果当前元素与栈顶的两个元素构成的向量顺时针旋转，则弹出栈顶元素
	stack := []int{0, 1}
	for i := 2; i < length; i++ {
		for len(stack) > 1 && cross(trees[stack[len(stack)-2]], trees[stack[len(stack)-1]], trees[i]) < 0 {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	result := make([][]int, len(stack))
	for i, point := range stack {
		result[i] = trees[point]
	}

	return result
}

/*太平洋大西洋水流问题
有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。“太平洋”处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。
这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵heights，heights[r][c]表示坐标 (r, c) 上单元格 高于海平面的高度 。
岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。
返回 网格坐标 resulT的 2D列表 ，其中result[i] = [ri, ci]表示雨水可以从单元格 (ri, ci) 流向 太平洋和大西洋 。

输入: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
输出: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]

回溯深度优先搜索: 从矩阵的边界开始反向搜索寻找雨水流向边界的单元格，反向搜索时，每次只能移动到高度相同或更大的单元格
搜索过程中需要记录每个单元格是否可以从太平洋反向到达以及是否可以从大西洋反向到达。反向搜索结束之后，遍历每个网格，如果一个网格既可以从太平洋反向到达也可以从大西洋反向到达，则该网格满足太平洋和大西洋都可以到达，将该网格添加到答案中。
*/

func pacificAtlantic(heights [][]int) [][]int {
	// initializing
	result := [][]int{}
	directions := []struct{ x, y int }{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	nR, nC := len(heights), len(heights[0])
	pac, atl := make([][]bool, nR), make([][]bool, nR)
	for i := range pac {
		pac[i] = make([]bool, nC)
		atl[i] = make([]bool, nC)
	}

	// recursive function
	var dfs func(int, int, [][]bool)
	dfs = func(x, y int, ocean [][]bool) {
		if ocean[x][y] { // meet an already iterated block : return
			return
		}
		ocean[x][y] = true             // an eligible block
		for _, i := range directions { // recursive detection for an eligible block
			if nx, ny := x+i.x, y+i.y; 0 <= nx && nx < nR && 0 <= ny && ny < nC && heights[nx][ny] >= heights[x][y] {
				dfs(nx, ny, ocean)
			}
		}
	}

	// dectection
	for i := 0; i < nR; i++ {
		for j := 0; j < nC; j++ {
			if i == 0 || j == 0 { // initiate and search pacific
				dfs(i, j, pac)
			}
			if i == nR-1 || j == nC-1 {
				dfs(i, j, atl)
			}
		}
	}

	// pick eligible blocks
	for i := 0; i < nR; i++ {
		for j := 0; j < nC; j++ {
			if atl[i][j] && pac[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

/*
给定一个长度为 n 的整数数组nums。
假设arrk是数组nums顺时针旋转 k 个位置后的数组，我们定义nums的 旋转函数F为：
F(k) = 0 * arrk[0] + 1 * arrk[1] + ... + (n - 1) * arrk[n - 1]
返回F(0), F(1), ..., F(n-1)中的最大值。
生成的测试用例让答案符合32 位 整数。

输入: nums = [4,3,2,6]
输出: 26
解释:
F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26
所以 F(0), F(1), F(2), F(3) 中的最大值是 F(3) = 26 。

F(0) = 0 * [0] + 1 * [1] + .. + (n - 1) * [n - 1]
F(1) = 1 * [0] + 2 * [1] + ... + 0 * [n - 1] = F(0) + [0] + [1] + ... + [n - 1] - n * [n - 1]
递推公式： F (k) = F(k - 1) + sum - n * [n - k]

*/

func maxRotateFunction(nums []int) int {
	result, sum, temp := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		temp += i * nums[i]
		result = temp
	}
	for i := 1; i < len(nums); i++ {
		temp = temp + sum - len(nums)*nums[len(nums)-i]
		if temp > result {
			result = temp
		}
	}
	return result
}
