package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
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
RandPointSolution solution = new RandPointSolution(nums);
// pick(3) 应该返回索引 2,3 或者 4。每个索引的返回概率应该相等。
solution.pick(3);
// pick(1) 应该返回 0。因为只有nums[0]等于1。
solution.pick(1);

遍历nums，当我们第 i 次遇到值为target 的元素时，随机选择区间 [0,i) 内的一个整数，如果其等于 0，则将返回值置为该元素的下标，否则返回值不变。
*/

type Solution1 []int

func Constructor1(nums []int) Solution1 {
	return nums
}

func (this Solution1) Pick(target int) int {
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

/*
给你一个由若干单词组成的句子sentence ，单词间由空格分隔。每个单词仅由大写和小写英文字母组成。
请你将句子转换为 “山羊拉丁文（Goat Latin）”（一种类似于 猪拉丁文- Pig Latin 的虚构语言）。山羊拉丁文的规则如下：

如果单词以元音开头（'a', 'e', 'i', 'o', 'u'），在单词后添加"ma"。
例如，单词 "apple" 变为 "applema" 。
如果单词以辅音字母开头（即，非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"。
例如，单词 "goat" 变为 "oatgma" 。
根据单词在句子中的索引，在单词最后添加与索引相同数量的字母'a'，索引从 1 开始。
例如，在第一个单词后添加 "a" ，在第二个单词后添加 "aa" ，以此类推。
返回将 sentence 转换为山羊拉丁文后的句子。

输入：sentence = "I speak Goat Latin"
输出："Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
*/

func toGoatLatin(sentence string) string {
	length, result := len(sentence), ""
	count, newWord, voy, tempCon := 1, true, false, ""
	for i := 0; i <= length; i++ {
		if newWord && i < length {
			newWord = false
			if sentence[i] == 'a' || sentence[i] == 'e' || sentence[i] == 'i' || sentence[i] == 'o' || sentence[i] == 'u' || sentence[i] == 'A' || sentence[i] == 'E' || sentence[i] == 'I' || sentence[i] == 'O' || sentence[i] == 'U' {
				voy = true
			} else {
				tempCon = string(sentence[i])
				continue
			}
		}

		if i == length || sentence[i] == ' ' {
			newWord = true
			if voy {
				result += "ma"
				voy = false
			} else {
				result += tempCon
				result += "ma"
			}
			for j := 0; j < count; j++ {
				result += "a"
			}
			count++
		}

		if i < length {
			result += string(sentence[i])
		}
	}

	return result
}

/*按奇偶排序数组
给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。返回满足此条件的 任一数组 作为答案。

输入：nums = [3,1,2,4]
输出：[2,4,3,1]
解释：[4,2,3,1]、[2,4,1,3] 和 [4,2,1,3] 也会被视作正确答案。
*/

func sortArrayByParity(nums []int) []int {
	length := len(nums)

	for true {
		i, j := 0, length-1
		for i < j && nums[i]%2 == 0 {
			i++
		}
		for i < j && nums[j]%2 != 0 {
			j--
		}

		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp

		if i >= j {
			break
		}
	}

	return nums
}

/* 建立四叉树
给你一个 n * n 矩阵 grid ，矩阵由若干 0 和 1 组成。请你用四叉树表示该矩阵 grid 。你需要返回能表示矩阵的 四叉树 的根结点。注意，当 isLeaf 为 False 时，你可以把 True 或者 False 赋值给节点，两种值都会被判题机制 接受 。

四叉树数据结构中，每个内部节点只有四个子节点。此外，每个节点都有两个属性：
val：储存叶子结点所代表的区域的值。1 对应 True，0 对应 False；
isLeaf: 当这个节点是一个叶子结点时为 True，如果它有 4 个子节点则为 False 。
class Node {
 public boolean val;
 public boolean isLeaf;
 public Node topLeft;
 public Node topRight;
 public Node bottomLeft;
 public Node bottomRight;
}

我们可以按以下步骤为二维区域构建四叉树：
如果当前网格的值相同（即，全为 0 或者全为 1），将 isLeaf 设为 True ，将 val 设为网格相应的值，并将四个子节点都设为 Null 然后停止。如果当前网格的值不同，将 isLeaf 设为 False， 将 val 设为任意值，然后如下图所示，将当前网格划分为四个子网格。使用适当的子网格递归每个子节点。

四叉树格式：输出为使用层序遍历后四叉树的序列化形式，其中 null 表示路径终止符，其下面不存在节点。它与二叉树的序列化非常相似。唯一的区别是节点以列表形式表示 [isLeaf, val] 。如果 isLeaf 或者 val 的值为 True ，则表示它在列表[isLeaf, val] 中的值为 1 ；如果 isLeaf 或者 val 的值为 False ，则表示值为 0 。

输入：grid = [[0,1],[1,0]]
输出：[[0,1],[1,0],[1,1],[1,1],[1,0]]

具体地，我们用递归函数 dfs(r_0, c_0, r_1, c_1)处理给定的矩阵， 首先判定这一部分是否均为 0 或 1，如果是，那么这一部分对应的是一个叶节点，我们构造出对应的叶节点并结束递归；如果不是，那么这一部分对应的是一个非叶节点，我们需要将其分成四个部分。根据这两条分界线递归地调用 \text{dfs}dfs 函数得到四个部分对应的树，再将它们对应地挂在非叶节点的四个子节点上。
*/
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var dfs func([][]int, int, int) *Node // the start and fin of columns as int
	dfs = func(rows [][]int, r0, r1 int) *Node {
		value := rows[0][r0]
		for _, row := range rows {
			for j := r0; j < r1; j++ {
				if row[j] != value {
					midR, midC := len(rows)/2, (r1-r0)/2+r0 // to locate the subdivided grid
					// recursive construction
					return &Node{
						false,
						false,
						dfs(rows[:midR], r0, midC),
						dfs(rows[:midR], midC, r1),
						dfs(rows[midR:], r0, midC),
						dfs(rows[midR:], midC, r1),
					}
				}
			}
		}
		return &Node{value == 1, true, nil, nil, nil, nil}
	}

	return dfs(grid, 0, len(grid))
}

/* 最小差值 I
给你一个整数数组 nums，和一个整数 k 。在一个操作中，您可以选择 0 <= i < nums.length 的任何索引 i 。将 nums[i] 改为 nums[i] + x ，其中 x 是一个范围为 [-k, k] 的整数。对于每个索引 i ，最多 只能 应用 一次 此操作。
nums的分数是nums中最大和最小元素的差值。在对 nums 中的每个索引最多应用一次上述操作后，返回nums 的最低 分数 。

输入：nums = [0,10], k = 2
输出：6
解释：将 nums 改为 [2,8]。分数是 max(nums) - min(nums) = 8 - 2 = 6。
*/
func smallestRangeI(nums []int, k int) int {
	min, max := nums[0], nums[0]
	for _, i := range nums {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}

	if max-min <= 2*k {
		return 0
	} else {
		return (max - min) - 2*k
	}
}

/*两棵二叉搜索树中的所有元素
给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

中序遍历访问二叉搜索树 == 有序数组, 然后可以使用双指针方法来合并这两个有序数组

*/

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	result1, result2, result := TreeTraverseInorder(root1), TreeTraverseInorder(root2), []int{}

	if len(result1) == 0 {
		return result1
	} else if len(result2) == 0 {
		return result2
	}

	for i, j := 0, 0; i < len(result1) && j < len(result2); {
		if result1[i] <= result2[j] {
			result = append(result, result1[i])
			i++
		} else if result1[i] > result2[j] {
			result = append(result, result2[j])
			j++
		}

		if i == len(result1) && j < len(result2) {
			result = append(result, result2[j:]...)
		} else if j == len(result2) && i < len(result1) {
			result = append(result, result1[i:]...)
		}
	}

	return result
}

/*标签验证器
给定一个表示代码片段的字符串，你需要实现一个验证器来解析这段代码，并返回它是否合法。合法的代码片段需要遵守以下的所有规则：

代码必须被合法的闭合标签包围。否则，代码是无效的。
闭合标签（不一定合法）要严格符合格式：<TAG_NAME>TAG_CONTENT</TAG_NAME>。其中，<TAG_NAME>是起始标签，</TAG_NAME>是结束标签。起始和结束标签中的 TAG_NAME 应当相同。当且仅当 TAG_NAME 和 TAG_CONTENT 都是合法的，闭合标签才是合法的。
合法的TAG_NAME仅含有大写字母，长度在范围 [1,9] 之间。否则，该 TAG_NAME 是不合法的。
合法的TAG_CONTENT 可以包含其他合法的闭合标签，cdata （请参考规则7）和任意字符（注意参考规则1）除了不匹配的<、不匹配的起始和结束标签、不匹配的或带有不合法 TAG_NAME 的闭合标签。否则，TAG_CONTENT 是不合法的。
一个起始标签，如果没有具有相同 TAG_NAME 的结束标签与之匹配，是不合法的。反之亦然。不过，你也需要考虑标签嵌套的问题。
一个<，如果你找不到一个后续的>与之匹配，是不合法的。并且当你找到一个<或</时，所有直到下一个>的前的字符，都应当被解析为 TAG_NAME（不一定合法）。
cdata 有如下格式：<![CDATA[CDATA_CONTENT]]>。CDATA_CONTENT 的范围被定义成 <![CDATA[ 和后续的第一个 ]]>之间的字符。
CDATA_CONTENT 可以包含任意字符。cdata 的功能是阻止验证器解析CDATA_CONTENT，所以即使其中有一些字符可以被解析为标签（无论合法还是不合法），也应该将它们视为常规字符。

合法代码的例子:

输入: "<DIV>This is the first line <![CDATA[<div>]]></DIV>"

输出: True

解释:

代码被包含在了闭合的标签内： <DIV> 和 </DIV> 。

TAG_NAME 是合法的，TAG_CONTENT 包含了一些字符和 cdata 。

即使 CDATA_CONTENT 含有不匹配的起始标签和不合法的 TAG_NAME，它应该被视为普通的文本，而不是标签。

所以 TAG_CONTENT 是合法的，因此代码是合法的。最终返回True。


输入: "<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>"

输出: True

解释:

我们首先将代码分割为： start_tag|tag_content|end_tag 。

start_tag -> "<DIV>"

end_tag -> "</DIV>"

tag_content 也可被分割为： text1|cdata|text2 。

text1 -> ">>  ![cdata[]] "

cdata -> "<![CDATA[<div>]>]]>" ，其中 CDATA_CONTENT 为 "<div>]>"

text2 -> "]]>>]"


start_tag 不是 "<DIV>>>" 的原因参照规则 6 。
cdata 不是 "<![CDATA[<div>]>]]>]]>" 的原因参照规则 7 。
不合法代码的例子:

输入: "<A>  <B> </A>   </B>"
输出: False
解释: 不合法。如果 "<A>" 是闭合的，那么 "<B>" 一定是不匹配的，反之亦然。

输入: "<DIV>  div tag is not closed  <DIV>"
输出: False

输入: "<DIV>  unmatched <  </DIV>"
输出: False

输入: "<DIV> closed tags with invalid tag name  <b>123</b> </DIV>"
输出: False

输入: "<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>"
输出: False

输入: "<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>"
输出: False
注意:

为简明起见，你可以假设输入的代码（包括提到的任意字符）只包含数字, 字母, '<','>','/','!','[',']'和' '。

"<TRUE><![CDATA[wahaha]]]><![CDATA[]> wahaha]]></TRUE>"
*/

func isValid1(code string) bool {
	// parse labels
	parseOP, labelEnds := true, false
	nLabels, parseLabel, labelLength, labels := 0, false, 0, []string{}
	if code[0] != '<' || code[len(code)-1] != '>' {
		return false
	}

	for i := 0; i < len(code); i++ {
		if labelEnds {
			return false
		}

		if parseLabel { // if started parsing a label
			if code[i] == '/' && code[i-1] == '<' { // legal closing label
				continue
			}
			if code[i] == '>' { // exit label mode
				if labelLength == 0 || labelLength > 9 { // illegal length
					return false
				}
				if parseOP { // op : add a label, close : check a label
					labels = append(labels, code[i-labelLength:i])
				} else {
					if code[i-labelLength:i] == labels[len(labels)-1] {
						labels = labels[:len(labels)-1]
						if nLabels == 0 {
							labelEnds = true
						}
					} else {
						return false
					}
				}
				labelLength = 0
				parseLabel = false
			} else if code[i] < 'A' || code[i] > 'Z' { // invalid label content
				return false
			} else { // valid label content
				labelLength++
			}
		}

		// remove cdata : <![CDATA[ + ... + ]]>
		if code[i] == '<' { // count labels
			if i+9 < len(code) {
				if code[i+1:i+9] == "![CDATA[" && len(labels) != 0 {
					for j := i + 9; j < len(code)-3; j++ {
						if code[j:j+3] == "]]>" { // remove the cdata
							code = code[:i] + code[j+3:]
							i = i - 2
							break
						}
					}
				}
			}

			if code[i] == '<' {
				if code[i+1] != '/' { // meet op labels
					parseOP = true
					nLabels++
				} else { // meet closing labels
					parseOP = false
					nLabels--
				}
				parseLabel = true // start parsing labels
			}
		}

		if nLabels < 0 { // illegally surrounded by labels
			return false
		}
	}

	if nLabels != 0 { // not closed
		return false
	}
	return true
}

/* 重新排列日志文件
由于标签具有「最先开始的标签最后结束」的特性，因此我们可以考虑使用一个栈存储当前开放的标签。
如果当前的字符为 <，那么需要考虑下面的四种情况：
	如果下一个字符为 /，那么说明我们遇到了一个结束标签。我们需要定位下一个 > 的位置 j
	如果下一个字符为 !，那么说明我们遇到了一个 cdata，我们需要继续往后读 7 个字符，判断其是否为 [CDATA[。
	如果下一个字符为大写字母，那么说明我们遇到了一个开始标签。我们需要定位下一个 > 的位置
	除此之外，如果不存在下一个字符，或者下一个字符不属于上述三种情况，那么  code 是不合法的
如果当前的字符为其它字符，那么根据规则 11，栈中需要存在至少一个开放的标签
*/
//NOT valid
func isValid(code string) bool {
	labels := []string{}
	length := len(code)

	for i := 0; i < length; i++ {
		if code[i] == '<' {
			if i+1 >= length {
				return false
			} else if code[i+1] == '/' { // meet an ed label
				if len(labels) == 0 { // no op label
					return false
				} else {
					lenLabel, label := 1, ""
					for code[i+1+lenLabel] >= 'A' && code[i+1+lenLabel] <= 'Z' {
						label = label + string(code[i+1+lenLabel])
						lenLabel++
					}
					if code[i+1+lenLabel] == '>' { // pop the last op label
						if labels[len(labels)-1] == label {
							labels = labels[0 : len(labels)-1]
							i = i + 1 + lenLabel
							continue
						} else {
							return false // different op label
						}
					} else { // illegal ed label
						return false
					}
				}
			} else if code[i+1] == '!' { // meet a cdata
				if i+12 < len(code) {
					if code[i+1:i+9] == "![CDATA[" {
						for j := i + 9; j+2 < len(code); j++ {
							if code[j:j+3] == "]]>" {
								i = j + 2
								break
							}
						}
					}
				}
				return false // illegal cdata
			} else if code[i+1] >= 'A' && code[i+1] <= 'Z' { // meet an op label
				lenLabel, label := 1, ""
				for i+lenLabel < length && code[i+lenLabel] >= 'A' && code[i+lenLabel] <= 'Z' {
					label = label + string(1) // NOT VALID
					lenLabel++
				}
				if code[i+1+lenLabel] == '>' {
					labels = append(labels, label)
				} else {
					return false // illegal op label
				}
			}
		}
	}
	if len(labels) == 0 {
		return true
	} else {
		return false
	}
}

/*
给你一个日志数组 logs。每条日志都是以空格分隔的字串，其第一个字为字母与数字混合的 标识符 。

有两种不同类型的日志：
字母日志：除标识符之外，所有字均由小写字母组成
数字日志：除标识符之外，所有字均由数字组成
请按下述规则将日志重新排序：
所有 字母日志 都排在 数字日志 之前。
字母日志 在内容不同时，忽略标识符后，按内容字母顺序排序；在内容相同时，按标识符排序。
数字日志 应该保留原来的相对顺序。

输入：logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
输出：["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
解释：
字母日志的内容都不同，所以顺序为 "art can", "art zero", "own kit dig" 。
数字日志保留原来的相对顺序 "dig1 8 1 5 1", "dig2 3 6" 。
*/

func reorderLogFiles(logs []string) []string {
	var ampute func(str string) string
	ampute = func(str string) string {
		for i := 0; i < len(str); i++ {
			if str[i] == ' ' {
				return str[i+1:]
			}
		}
		return ""
	}

	var amputeComp func(str1, str2 string) string // return if str2 >=< str1 without tag
	amputeComp = func(str1, str2 string) string {
		temp1, temp2 := ampute(str1), ampute(str2)

		if temp1 > temp2 {
			return ">"
		} else if temp1 == temp2 {
			return "="
		} else {
			return "<"
		}
	}

	for index, strs := range logs {
		fmt.Println(strs, logs)
		if ampute(strs)[0] > '9' { // reorder if meet a letter
			switchIndex := index
			for i := index - 1; i >= 0; i-- { // compare with the previous, if ok, switch
				if ampute(logs[i])[0] <= '9' || amputeComp(strs, logs[i]) == "<" || amputeComp(strs, logs[i]) == "=" && strs < logs[i] {
					logs[i], logs[switchIndex] = logs[switchIndex], logs[i]
					switchIndex = i
				}
			}
		}
	}
	return logs
}

/*
共有 n 名小伙伴一起做游戏。小伙伴们围成一圈，按 顺时针顺序 从 1 到 n 编号。确切地说，从第 i 名小伙伴顺时针移动一位会到达第 (i+1) 名小伙伴的位置，其中 1 <= i < n ，从第 n 名小伙伴顺时针移动一位会回到第 1 名小伙伴的位置。

游戏遵循如下规则：
从第 1 名小伙伴所在位置 开始 。
沿着顺时针方向数 k 名小伙伴，计数时需要 包含 起始时的那位小伙伴。逐个绕圈进行计数，一些小伙伴可能会被数过不止一次。
你数到的最后一名小伙伴需要离开圈子，并视作输掉游戏。
如果圈子中仍然有不止一名小伙伴，从刚刚输掉的小伙伴的 顺时针下一位 小伙伴 开始，回到步骤 2 继续执行。
否则，圈子中最后一名小伙伴赢得游戏。
给你参与游戏的小伙伴总数 n ，和一个整数 k ，返回游戏的获胜者。

输入：n = 5, k = 2
输出：3
解释：游戏运行步骤如下：
1) 从小伙伴 1 开始。
2) 顺时针数 2 名小伙伴，也就是小伙伴 1 和 2 。
3) 小伙伴 2 离开圈子。下一次从小伙伴 3 开始。
4) 顺时针数 2 名小伙伴，也就是小伙伴 3 和 4 。
5) 小伙伴 4 离开圈子。下一次从小伙伴 5 开始。
6) 顺时针数 2 名小伙伴，也就是小伙伴 5 和 1 。
7) 小伙伴 1 离开圈子。下一次从小伙伴 3 开始。
8) 顺时针数 2 名小伙伴，也就是小伙伴 3 和 5 。
9) 小伙伴 5 离开圈子。只剩下小伙伴 3 。所以小伙伴 3 是游戏的获胜者。
*/
func findTheWinner1(n int, k int) int {
	if k == 1 {
		return n
	}

	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i
	}

	elim := 0
	for len(list) != 1 {
		elim += k - 1
		for elim >= len(list) {
			elim -= len(list)
		}
		list = append(list[:elim], list[elim+1:]...)
	}

	return list[0] + 1
}

/*
约瑟夫环——公式法（递推公式）
相当于把数组向前移动M位。若已知N-1个人时，胜利者的下标位置f(N−1,M)，则N个人的时候，就是往后移动M位，(因为有可能数组越界，超过的部分会被接到头上，所以还要模N)，既f(N,M)=(f(N−1,M)+M)%n
*/
func findTheWinner(n int, k int) int {
	result := 0
	for i := 2; i <= n; i++ {
		result += (result + k) % i
	}

	return result + 1
}

/*乘积小于 K 的子数组
给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。

输入：nums = [10,5,2,6], k = 100
输出：8
解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2],、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。

滑动窗口, 记录以每个数字为右边界所形成的有效子数组的个数


*/
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	i, j, prod, result := 0, 0, 1, 0
	for j < len(nums) {

		prod *= nums[j]
		for prod >= k && i < j {
			prod /= nums[i]
			i++
		}
		result += j - i + 1

		j++
	}

	return result
}

/*最近的请求次数
写一个RecentCounter类来计算特定时间范围内最近的请求。
请你实现 RecentCounter 类：
RecentCounter() 初始化计数器，请求数为 0 。
int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。
保证 每次对 ping 的调用都使用比之前更大的 t 值。

输入：
["RecentCounter", "ping", "ping", "ping", "ping"]
[[], [1], [100], [3001], [3002]]
输出：
[null, 1, 2, 3, 3]
解释：
RecentCounter recentCounter = new RecentCounter();
recentCounter.ping(1);     // requests = [1]，范围是 [-2999,1]，返回 1
recentCounter.ping(100);   // requests = [1, 100]，范围是 [-2900,100]，返回 2
recentCounter.ping(3001);  // requests = [1, 100, 3001]，范围是 [1,3001]，返回 3
recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002]，范围是 [2,3002]，返回 3
*/
/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

type RecentCounter []int

//func Constructor() (_ RecentCounter) {
//	return
//}

func (this *RecentCounter) Ping(t int) int { // ping : 收到一个请求。在每次收到请求的时候都要返回三秒内的请求数
	*this = append(*this, t) // 入队列
	for (*this)[0] < t-3000 {
		*this = (*this)[1:]
	}

	return len(*this)
}

/*最小基因变化
基因序列可以示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
假设我们需要调查从基因序列start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。
给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使start 变化为 end 所需的最少变化次数。如果无法完成此基因变化，返回 -1 。
注意：起始基因序start 默认是有效的，但是它并不一定会出现在基因库中。


示例 1：

输入：start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
输出：1
示例 2：

输入：start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
输出：2
示例 3：

输入：start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC","AACCCCCC"]
输出：3

提示：
start.length == 8
end.length == 8
0 <= bank.length <= 10
bank[i].length == 8
start、end 和 bank[i] 仅由字符 ['A', 'C', 'G', 'T'] 组成


广度优先搜索：

*/
func minMutation(start string, end string, bank []string) int {

}

/*数组中重复的数据
给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整数，并以数组形式返回。
你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。

输入：nums = [4,3,2,7,8,2,3,1]
输出：[2,3]

use the index of original list to simulate a hash map
nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 : mark some opposite value
*/
func findDuplicates(nums []int) []int {
	result := []int{}
	temp := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			temp = -nums[i]
		} else {
			temp = nums[i]
		}

		if nums[temp-1] < 0 {
			result = append(result, temp)
		}
		nums[temp-1] = -nums[temp-1]
	}
	return result
}

/*
由范围 [0,n] 内所有整数组成的 n + 1 个整数的排列序列可以表示为长度为 n 的字符串 s ，其中:
如果perm[i] < perm[i + 1]，那么s[i] == 'I'
如果perm[i] > perm[i + 1]，那么 s[i] == 'D'
给定一个字符串 s ，重构排列perm 并返回它。如果有多个有效排列perm，则返回其中 任何一个 。

示例 1：
输入：s = "IDID"
输出：[0,4,1,3,2]
*/
func diStringMatch(s string) []int {
	left, right := 0, len(s)
	result := []int{}
	for indexStr := 0; indexStr < len(s); indexStr++ {
		if s[indexStr] == 'I' {
			result = append(result, left)
			left++
		} else {
			result = append(result, right)
			right--
		}
	}
	result = append(result)

	if result[len(s)-1] == left {
		result = append(result, right)
	} else {
		result = append(result, left)
	}

	return result
}

/*
序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。
编码的字符串应尽可能紧凑。

示例 1：
输入：root = [2,1,3]
输出：[2,1,3]
示例 2：
输入：root = []
输出：[]

提示：
树中节点数范围是 [0, 104]
0 <= Node.val <= 104
题目数据 保证 输入的树是一棵二叉搜索树。

思路
仅对二叉搜索树做「先序遍历」或者「后序遍历」，即可达到序列化和反序列化的要求。
后序遍历得到的数组中，根结点的值位于数组末尾，左子树的节点均小于根节点的值，右子树的节点均大于根节点的值，可以根据这些性质设计递归函数恢复二叉搜索树。
*/

type Codec struct {
}

//func Constructor() (_ Codec) {
//	return
//}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	postR := []string{}

	var traverseMid func(*TreeNode)
	traverseMid = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverseMid(node.Left)
		traverseMid(node.Right)
		postR = append(postR, strconv.Itoa(node.Val))
	}

	traverseMid(root)
	return strings.Join(postR, " ")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	sep := strings.Split(data, " ")

	var construct func(int, int) *TreeNode
	construct = func(lower int, upper int) *TreeNode {
		if len(sep) == 0 {
			return nil // return when all seperated values removed
		}

		val, _ := strconv.Atoi(sep[len(sep)-1])
		if val < lower || val > upper { // leaf node, return nil for its leaves`
			return nil
		}
		sep = sep[:len(sep)-1] // remove the processed node

		return &TreeNode{Val: val, Right: construct(val, upper), Left: construct(lower, val)}
	}

	return construct(math.MinInt32, math.MaxInt32)
}

/*一次编辑
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。

输入:
first = "pale"
second = "ple"
输出: True

输入:
first = "pales"
second = "pal"
输出: False
*/
func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}

	len1, len2 := len(first), len(second)
	if len1+1 < len2 || len2+1 < len1 {
		return false
	}

	var length int
	if len1 < len2 {
		length = len1
	} else {
		length = len2
	}

	if length == 0 {
		return true
	}

	used := false
	for i := 0; i < length; i++ {
		if first[i] != second[i] {
			if used {
				return false
			}

			if len1 < len2 {
				second = second[:i] + second[i+1:]
				i--
			} else if len1 > len2 {
				first = first[:i] + first[i+1:]
			}

			used = true
		}
	}

	return true
}

/*贴纸拼词
我们有 n 种不同的贴纸。每个贴纸上都有一个小写的英文单词。
您想要拼写出给定的字符串 target，方法是从收集的贴纸中切割单个字母并重新排列它们。如果你愿意，你可以多次使用每个贴纸，每个贴纸的数量是无限的。
返回你需要拼出 target的最小贴纸数量。如果任务不可能，则返回 -1 。
注意：在所有的测试用例中，所有的单词都是从 1000 个最常见的美国英语单词中随机选择的，并且 target 被选择为两个随机单词的连接。

示例 1：
输入： stickers = ["with","example","science"], target = "thehat"
输出：3
解释：
我们可以使用 2 个 "with" 贴纸，和 1 个 "example" 贴纸。
把贴纸上的字母剪下来并重新排列后，就可以形成目标 “thehat“ 了。
此外，这是形成目标字符串所需的最小贴纸数量。

记忆化搜索 + 状态压缩
target有2^m个子序列，dp(子序列)为所需的最小贴纸数。对于某一子序列：使用交集遍历挑选最优的sticker，未被覆盖的其他字符用dp继续计算。
用二进制数来表示某一子序列。
*/
func minStickers(stickers []string, target string) int {
	length := len(target)

	masks := make([]int, 1<<length) // 2 ^ n, 长度等于子序列数量
	for i := range masks {
		masks[i] = -1
	}
	masks[0] = 0

	var dp func(int) int
	dp = func(mask int) int { // dp函数的作用是，输入mask：还需要解决哪些状态，返回：最少需要的卡片数
		if masks[mask] != -1 { // 如果没有需要解决的状态，就不需要卡片，那么直接return 0
			return masks[mask]
		}
		masks[mask] = length + 1 // 不会需要length+1张卡片，把这个设置为初始最大值

		for _, sticker := range stickers { // 对于每一个贴纸
			left := mask
			count := [26]int{}
			for _, caracter := range sticker { // 对于某个贴纸里的每一个字母
				count[caracter-'a']++ // 用列表统计该贴纸的每个字母数量
			}
			for index, caracter := range target { // 对于目标单词
				if mask>>index&1 == 1 && count[caracter-'a'] > 0 { // 如果第i位恰好没解决，且第i位的字母，卡片sticker还有，
					count[caracter-'a']-- // 那么就把sticker的这个字母减下来，然后把然后把这一位解决掉
					left ^= 1 << index
				}
			}
			if left < mask {
				masks[mask] = MinOf2(masks[mask], dp(left)+1)
			}
		}
		return masks[mask]
	}

	ans := dp(1<<length - 1)
	if ans <= length {
		return ans
	}
	return -1
}

/*最大三角形面积
给定包含多个点的集合，从其中取三个点组成三角形，返回能组成的最大三角形的面积。

示例:
输入: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
输出: 2

凸包：

*/
func largestTriangleArea(points [][]int) float64 {

}

/*乘法表中第k小的数
给定高度m 、宽度n 的一张 m * n的乘法表，以及正整数k，你需要返回表中第k 小的数字。

矩阵的二分查找
对于乘法表的第i行: i, 2i, 3i ... ni， 不超过x的数有min( floor(x/i), n)个，整个乘法表不超过x的数为：sum( min(...) )
i < floor(x/n) 时， floor(x/i) > n
*/
func findKthNumber(m int, n int, k int) int {
	left, right := 1, m*n // 二分查找的左右边界
	for left < right {
		x := left + (right-left)/2 // 二分的中点
		sum := x / n * n           // sum 的前n部分
		for i := x/n + 1; i < m; i++ {
			sum += x / i
		}
		if sum >= k { // 不超过x的数的数量大于等于k：向左移动右边界
			right = x
		} else { // 否则移动左边界
			left = x + 1
		}
	}
	return left
}

/*最少移动次数使数组元素相等 II
给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最少移动数。在一步操作中，你可以使数组中的一个元素加 1 或者减 1 。

示例 1：
输入：nums = [1,2,3]
输出：2
解释：
只需要两步操作（每步操作指南使一个元素加 1 或减 1）：
[1,2,3]  =>  [2,2,3]  =>  [2,2,2]

取排列后的数组的中位数
*/
func minMoves2(nums []int) int {
	sort.Ints(nums)
	n := nums[len(nums)/2]

	result := 0
	for _, i := range nums {
		if i < n {
			result = result + n - i
		} else {
			result = result + i - n
		}
	}

	return result
}

/*
给你一个区间数组 intervals ，其中intervals[i] = [starti, endi] ，且每个starti 都 不同 。区间 i 的 右侧区间 可以记作区间 j ，并满足 startj>= endi ，且 startj 最小化 。
返回一个由每个区间 i 的 右侧区间 的最小起始位置组成的数组。如果某个区间 i 不存在对应的 右侧区间 ，则下标 i 处的值设为 -1 。

输入：intervals = [[3,4],[2,3],[1,2]]
输出：[-1,0,1]
解释：对于 [3,4] ，没有满足条件的“右侧”区间。
对于 [2,3] ，区间[3,4]具有最小的“右”起点;
对于 [1,2] ，区间[2,3]具有最小的“右”起点。

1 <=intervals.length <= 2 * 104
intervals[i].length == 2
-106 <= starti <= endi <= 106
每个间隔的起点都不相同
*/
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	type pair struct{ value, index int }
	start, end := make([]pair, n), make([]pair, n) // 起始点and结束点从小到大排序
	for index, i := range intervals {              // 问题转化为求两个有序数组start和end。对end中每个元素找start中最小的大于它的值
		start[index] = pair{i[0], index}
		end[index] = pair{i[1], index}
	}

	sort.Slice(start, func(i, j int) bool { // 根据第一个元素排序
		return start[i].value < start[j].value
	})
	sort.Slice(end, func(i, j int) bool {
		return end[i].value < end[j].value
	})

	result := make([]int, n)
	j := 0 // 已排序，每次直接从j开始比较
	for _, e := range end {
		for j < n && start[j].value < e.value {
			j++
		}
		if j < n {
			result[e.index] = start[j].index
		} else {
			result[e.index] = -1
		}
	}

	return result
}

/*单值二叉树
如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
只有给定的树是单值二叉树时，才返回 true；否则返回 false。
*/
func isUnivalTree(root *TreeNode) bool {
	return root == nil || (root.Left == nil || root.Val == root.Left.Val && isUnivalTree(root.Left)) && (root.Right == nil || root.Val == root.Right.Val && isUnivalTree(root.Right))
}

/*
在无限长的数轴（即 x 轴）上，我们根据给定的顺序放置对应的正方形方块。
第 i 个掉落的方块（positions[i] = (left, side_length)）是正方形，其中left 表示该方块最左边的点位置(positions[i][0])，side_length 表示该方块的边长(positions[i][1])。
每个方块的底部边缘平行于数轴（即 x 轴），并且从一个比目前所有的落地方块更高的高度掉落而下。在上一个方块结束掉落，并保持静止后，才开始掉落新方块。
方块的底边具有非常大的粘性，并将保持固定在它们所接触的任何长度表面上（无论是数轴还是其他方块）。邻接掉落的边不会过早地粘合在一起，因为只有底边才具有粘性。
返回一个堆叠高度列表ans 。每一个堆叠高度ans[i]表示在通过positions[0], positions[1], ..., positions[i]表示的方块掉落结束后，目前所有已经落稳的方块堆叠的最高高度。


示例 1:
输入: [[1, 2], [2, 3], [6, 1]]
输出: [2, 5, 5]
解释:

第一个方块 positions[0] = [1, 2] 掉落：
_aa
_aa
-------
方块最大高度为 2 。

第二个方块 positions[1] = [2, 3] 掉落：
__aaa
__aaa
__aaa
_aa__
_aa__
--------------
方块最大高度为5。
大的方块保持在较小的方块的顶部，不论它的重心在哪里，因为方块的底部边缘有非常大的粘性。

第三个方块 positions[1] = [6, 1] 掉落：
__aaa
__aaa
__aaa
_aa
_aa___a
--------------
方块最大高度为5。
因此，我们返回结果[2, 5, 5]。

示例 2:
输入: [[100, 100], [200, 100]]
输出: [100, 100]
解释: 相邻的方块不会过早地卡住，只有它们的底部边缘才能粘在表面上。

使用map一对一地记录高度会超时，更改定义：hights[x]表明从x开始的点的高度
*/
func fallingSquares(positions [][]int) []int {
	length := len(positions)

	result := []int{}
	mapHeight := map[int]int{}
	mapHeight
	topAll := 0
	for index, cube := range positions {
		top := 0
		for length := cube[0]; length < cube[0]+cube[1]; length++ {
			if mapHeight[length] > top {
				top = mapHeight[length]
			}
		}
		for length := cube[0]; length < cube[0]+cube[1]; length++ {
			mapHeight[length] = top + cube[1]
		}
		if top+cube[1] > topAll {
			topAll = top + cube[1]
		}
		result = append(result, topAll)
	}

	return result
}

/*单词距离
有个内含单词的超大文本文件，给定任意两个不同的单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?

示例：
输入：words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
输出：1

目标：一次遍历
*/
func findClosest(words []string, word1 string, word2 string) int {
	index1, index2 := -1, -1 // 适配首位有/没有目标单词之一

	result := len(words)
	for index, i := range words {
		if i == word1 {
			index1 = index
		} else if i == word2 {
			index2 = index
		}
		if index1 != index2 && index1 >= 0 && index2 >= 0 {
			dist := index1 - index2
			if dist < 0 {
				dist = -dist
			}
			if dist < result {
				result = dist
			}
		}
	}

	return result
}

/*
有效括号字符串为空 ""、"(" + A + ")"或A + B ，其中A 和B都是有效的括号字符串，+代表字符串的连接。
例如，""，"()"，"(())()"和"(()(()))"都是有效的括号字符串。
如果有效字符串 s 非空，且不存在将其拆分为 s = A + B的方法，我们称其为原语（primitive），其中A 和B都是非空有效括号字符串。
给出一个非空有效字符串 s，考虑将其进行原语化分解，使得：s = P_1 + P_2 + ... + P_k，其中P_i是有效括号字符串原语。
对 s 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 s 。

示例 1：
输入：s = "(()())(())"
输出："()()()"
解释：
输入字符串为 "(()())(())"，原语化分解得到 "(()())" + "(())"，
删除每个部分中的最外层括号后得到 "()()" + "()" = "()()()"。

示例 2：
输入：s = "(()())(())(()(()))"
输出："()()()()(())"
解释：
输入字符串为 "(()())(())(()(()))"，原语化分解得到 "(()())" + "(())" + "(()(()))"，
删除每个部分中的最外层括号后得到 "()()" + "()" + "()(())" = "()()()()(())"。

示例 3
输入：s = "()()"
输出：""
解释：
输入字符串为 "()()"，原语化分解得到 "()" + "()"，
删除每个部分中的最外层括号后得到 "" + "" = ""。

提示：
1 <= s.length <= 105
s[i] 为 '(' 或 ')'
s 是一个有效括号字符串
*/
func removeOuterParentheses(s string) string {
	length := len(s)

	numP := 0
	result := ""
	for i := 0; i < length; i++ {

		if s[i] == '(' {

			numP++
			if numP == 1 {
				continue
			}
		} else {
			numP--
			if numP == 0 {
				continue
			}
		}

		result += string(s[i])

	}

	return result
}

/*
给定一个字符串queryIP。如果是有效的 IPv4 地址，返回 "IPv4" ；如果是有效的 IPv6 地址，返回 "IPv6" ；如果不是上述类型的 IP 地址，返回 "Neither" 。
有效的IPv4地址 是 “x1.x2.x3.x4” 形式的IP地址。 其中0 <= xi<= 255且xi不能包含 前导零。例如:“192.168.1.1”、 “192.168.1.0” 为有效IPv4地址， “192.168.01.1” 为无效IPv4地址; “192.168.1.00” 、 “192.168@1.1” 为无效IPv4地址。
一个有效的IPv6地址是一个格式为“x1:x2:x3:x4:x5:x6:x7:x8” 的IP地址，其中:

1 <= xi.length <= 4
xi是一个 十六进制字符串 ，可以包含数字、小写英文字母( 'a' 到 'f' )和大写英文字母( 'A' 到 'F' )。
在xi中允许前导零。
例如 "2001:0db8:85a3:0000:0000:8a2e:0370:7334" 和 "2001:db8:85a3:0:0:8A2E:0370:7334" 是有效的 IPv6 地址，而 "2001:0db8:85a3::8A2E:037j:7334" 和 "02001:0db8:85a3:0000:0000:8a2e:0370:7334" 是无效的 IPv6 地址。

示例 1：
输入：queryIP = "172.16.254.1"
输出："IPv4"
解释：有效的 IPv4 地址，返回 "IPv4"

示例 2：
输入：queryIP = "2001:0db8:85a3:0:0:8A2E:0370:7334"
输出："IPv6"
解释：有效的 IPv6 地址，返回 "IPv6"

示例 3：
输入：queryIP = "256.256.256.256"
输出："Neither"
解释：既不是 IPv4 地址，又不是 IPv6 地址
*/
func validIPAddress(queryIP string) string {
	pieces := strings.Split(queryIP, ".")
	if len(pieces) == 4 {
		for _, thisStr := range pieces {
			if len(thisStr) == 0 {
				return "Neither"
			}
			convert, err := strconv.Atoi(thisStr)
			if err == nil {
				if len(thisStr) == len(strconv.Itoa(convert)) && convert <= 255 {
					continue
				}
			}
			return "Neither"
		}
		return "IPv4"
	}

	pieces = strings.Split(queryIP, ":")
	if len(pieces) == 8 {
		for _, thisStr := range pieces {
			if len(thisStr) == 0 {
				return "Neither"
			}
			if len(thisStr) <= 4 {
				for i := 0; i < len(thisStr); i++ {
					if thisStr[i] >= 48 && thisStr[i] <= 57 || thisStr[i] >= 97 && thisStr[i] <= 102 || thisStr[i] >= 65 && thisStr[i] <= 70 {
						continue
					}
					return "Neither"
				}
			} else {
				return "Neither"
			}
		}
		return "IPv6"
	}

	return "Neither"
}

/*从根到叶的二进制数之和
给出一棵二叉树，其上每个结点的值都是0或1。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。
例如，如果路径为0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数01101，也就是13。
对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。
返回这些数字之和。题目数据保证答案是一个 32 位 整数。

后序遍历的访问顺序为：左子树——右子树——根节点。我们对根节点 root 进行后序遍历：
如果节点是叶子节点，返回它对应的数字 val。
如果节点是非叶子节点，返回它的左子树和右子树对应的结果之和。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	var dfs func(node *TreeNode, val int) int
	dfs = func(node *TreeNode, val int) int { // 有返回值，使用有序遍历
		if node == nil { // 空节点：返回零
			return 0
		}

		val = val<<1 | node.Val // 左移1位 ：增加一个二进制数后的进位变化。对此数按位取或：若等于1则+1

		if node.Left == nil && node.Right == nil { // 空节点：返回
			return val
		}

		return dfs(node.Left, val) + dfs(node.Right, val) // 实际上是后序遍历
	}

	return dfs(root, 0)
}

/*外星文字典
现有一种使用英语字母的外星文语言，这门语言的字母顺序与英语顺序不同。
给定一个字符串列表 words ，作为这门语言的词典，words 中的字符串已经 按这门新语言的字母顺序进行了排序 。
请你根据该词典还原出此语言中已知的字母顺序，并 按字母递增顺序 排列。若不存在合法字母顺序，返回 "" 。若存在多种可能的合法字母顺序，返回其中 任意一种 顺序即可。

字符串 s 字典顺序小于 字符串 t 有两种情况：
在第一个不同字母处，如果 s 中的字母在这门外星语言的字母顺序中位于 t 中字母之前，那么s 的字典顺序小于 t 。
如果前面 min(s.length, t.length) 字母都相同，那么 s.length < t.length 时，s 的字典顺序也小于 t 。
*/
func alienOrder(words []string) string {

}

/*连续整数求和
给定一个正整数 n，返回 连续正整数满足所有数字之和为 n 的组数 。

输入: n = 15
输出: 4
解释: 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5

思路：
找2个连续数字：15 - 1 = 14, 14 % 2 = 0, 14 / 2 = 7, 15 = 7 + 8 = 7 + 7 + 1
找3个连续数字：15 - 3 = 12, 12 % 3 = 0， 12 / 3 = 4 , 15 = 4 + 5 + 6 = 4 + 4 + 1 + 4 + 2
找4个连续数字：15 - 6 = 9, 9 % 4 != 0, pass
找5个连续数字：15 - 10 = 5, 5 mod 5 = 0, ...
*/
func consecutiveNumbersSum(n int) int {
	result := 0

}

/*
每个 有效电子邮件地址 都由一个 本地名 和一个 域名 组成，以 '@' 符号分隔。除小写字母之外，电子邮件地址还可以含有一个或多个'.' 或 '+' 。
例如，在alice@leetcode.com中，alice是 本地名 ，而leetcode.com是 域名 。
如果在电子邮件地址的 本地名 部分中的某些字符之间添加句点（'.'），则发往那里的邮件将会转发到本地名中没有点的同一地址。请注意，此规则 不适用于域名 。
例如，"alice.z@leetcode.com” 和 “alicez@leetcode.com”会转发到同一电子邮件地址。
如果在 本地名 中添加加号（'+'），则会忽略第一个加号后面的所有内容。这允许过滤某些电子邮件。同样，此规则 不适用于域名 。
例如 m.y+name@email.com 将转发到 my@email.com。
可以同时使用这两个规则。
给你一个字符串数组 emails，我们会向每个 emails[i] 发送一封电子邮件。返回实际收到邮件的不同地址数目。

*/
func numUniqueEmails(emails []string) int {
	sentAdr := map[string]bool{}

	result := 0
	for _, addres := range emails {
		thisAdr, parseLocal := "", true
		for i := 0; i < len(addres); i++ {
			if addres[i] == '@' {
				thisAdr += addres[i:]
				break
			}
			if parseLocal {
				if addres[i] == '.' {
					continue
				}
				if addres[i] == '+' {
					parseLocal = false
					continue
				}
				thisAdr += string(addres[i])
			}
		}
		if sentAdr[thisAdr] == true {
			continue
		} else {
			sentAdr[thisAdr] = true
			result++
		}
	}

	return result
}

/*
给定圆的半径和圆心的位置，实现函数 randPoint ，在圆中产生均匀随机点。

实现Solution类:
RandPointSolution(double radius, double x_center, double y_center)用圆的半径radius和圆心的位置 (x_center, y_center) 初始化对象
randPoint()返回圆内的一个随机点。圆周上的一点被认为在圆内。答案作为数组返回 [x, y] 。
*/
type RandPointSolution struct {
	radius   float64
	x_center float64
	y_center float64
}

func RandPointConstructor(radius float64, x_center float64, y_center float64) RandPointSolution {
	return RandPointSolution{
		radius:   radius,
		x_center: x_center,
		y_center: y_center,
	}
}

func (this *RandPointSolution) RandPoint() []float64 {
	for {
		x, y := rand.Float64()*2-1, rand.Float64()*2-1
		if x*x+y*y < 1 {
			return []float64{x*this.radius + this.x_center, y*this.radius + this.y_center}
		}
	}
}

/**
 * Your RandPointSolution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */

/*
给定一个由非重叠的轴对齐矩形的数组 rects ，其中 rects[i] = [ai, bi, xi, yi] 表示 (ai, bi) 是第 i 个矩形的左下角点，(xi, yi) 是第 i 个矩形的右上角角点。设计一个算法来随机挑选一个被某一矩形覆盖的整数点。矩形周长上的点也算做是被矩形覆盖。所有满足要求的点必须等概率被返回。
在一个给定的矩形覆盖的空间内任何整数点都有可能被返回。
请注意，整数点是具有整数坐标的点。

实现Solution类:
Solution(int[][] rects)用给定的矩形数组rects 初始化对象。
int[] pick() 返回一个随机的整数点 [u, v] 在给定的矩形所覆盖的空间内。
*/

type Solution struct {
	rects [][]int
}

func Constructor(rects [][]int) Solution {
	return Solution{
		rects: rects,
	}
}

func (this *Solution) Pick() []int {
	nums := []int{}
	nums = append(nums, 0)
	for i, rect := range this.rects {
		nums = append(nums, nums[i]+(rect[2]-rect[0]+1)*(rect[3]-rect[1]+1))
	}

	randi := rand.Intn(nums[len(nums)-1])
	iRect := sort.SearchInts(nums, randi+1) - 1
	points := randi - nums[iRect]
	col := points % (this.rects[iRect][3] - this.rects[iRect][1] + 1)
	row := points / (this.rects[iRect][3] - this.rects[iRect][1] + 1)
	if row < 0 {
		row = -row
	}
	return []int{this.rects[iRect][0] + row, this.rects[iRect][1] + col}
}

/**
 * Your RandPointSolution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */

/*将字符串翻转到单调递增
如果一个二进制字符串，是以一些 0（可能没有 0）后面跟着一些 1（也可能没有 1）的形式组成的，那么该字符串是 单调递增 的。
给你一个二进制字符串 s，你可以将任何 0 翻转为 1 或者将 1 翻转为 0 。
返回使 s 单调递增的最小翻转次数。

第i位的答案与第i-1位相关，联想使用动态规划。
分析：第i位可以是0或1
dp[i][0] = dp[i-1][0] + if(s[i] == 1) : 第i位若是0，则第i-1位必须是0。
dp[i][1] = min(dp[i-1][1], dp[i-1][0]) + if(s[i] == 0) : 若第i位为1，则第i-1位可以是0或1。
*/
func minFlipsMonoIncr(s string) int {
	dp := []int{0, 0}
	for _, thisS := range s {
		var dp0, dp1 int
		if thisS == '1' {
			dp0 = dp[0] + 1
			dp1 = MinOf2(dp[0], dp[1])
		} else {
			dp0 = dp[0]
			dp1 = MinOf2(dp[0], dp[1]) + 1
		}
		dp[0], dp[1] = dp0, dp1
	}

	return MinOf2(dp[0], dp[1])
}

/*
你有一个单词列表words和一个模式pattern，你想知道 words 中的哪些单词与模式匹配。
如果存在字母的排列 p，使得将模式中的每个字母 x 替换为 p(x) 之后，我们就得到了所需的单词，那么单词与模式是匹配的。
（回想一下，字母的排列是从字母到字母的双射：每个字母映射到另一个字母，没有两个字母映射到同一个字母。）
返回 words 中与给定模式匹配的单词列表。
你可以按任何顺序返回答案。

输入：words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
输出：["mee","aqq"]
解释：
"mee" 与模式匹配，因为存在排列 {a -> m, b -> e, ...}。
"ccc" 与模式不匹配，因为 {a -> c, b -> c, ...} 不是排列。
因为 a 和 b 映射到同一个字母。
*/
func findAndReplacePattern(words []string, pattern string) []string {
	result := []string{}
	for _, word := range words {
		bijection1, bijection2 := map[uint8]uint8{}, map[uint8]uint8{}
		ifTrue := true
		for index := 0; index < len(pattern); index++ {
			val1, ok1 := bijection1[pattern[index]]
			val2, ok2 := bijection2[word[index]]
			if !ok1 || !ok2 {
				if !ok1 {
					bijection1[pattern[index]] = word[index]
					val1, ok1 = bijection1[pattern[index]]
				}
				if !ok2 {
					bijection2[word[index]] = pattern[index]
					val2, ok2 = bijection2[word[index]]

				}
			}
			if val1 == word[index] && val2 == pattern[index] {
				continue
			} else {
				ifTrue = false
				break
			}
		}
		if ifTrue {
			result = append(result, word)
		}
	}
	return result
}
