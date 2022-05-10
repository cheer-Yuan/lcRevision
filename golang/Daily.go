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
		i, j := 0, length - 1
		for i < j && nums[i] % 2 == 0 {
			i++
		}
		for i < j && nums[j] % 2 != 0 {
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
	Val bool
	IsLeaf bool
	TopLeft *Node
	TopRight *Node
	BottomLeft *Node
	BottomRight *Node
	}

func construct(grid [][]int) *Node {
	  var dfs func([][]int, int, int) *Node		// the start and fin of columns as int
	  dfs = func(rows [][]int, r0, r1 int) *Node {
		value := rows[0][r0]
		for _, row := range rows{
			for j := r0; j < r1; j++ {
				if row[j] != value {
					midR, midC := len(rows) / 2, (r1 - r0) / 2 + r0	// to locate the subdivided grid
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

	if max - min <= 2 * k {
		return 0
	} else {
		return (max - min) - 2 * k
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
	if code[0] != '<' || code[len(code) - 1] != '>' {
		return false
	}

	for i := 0; i < len(code) ; i++  {
		if labelEnds {
			return false
		}

		if parseLabel {										// if started parsing a label
			if code[i] == '/' && code[i - 1] == '<'  {								// legal closing label
				continue
			}
			if code[i] == '>' {								// exit label mode
				if labelLength == 0 || labelLength >9 {		// illegal length
					return false
				}
				if parseOP {								// op : add a label, close : check a label
					labels = append(labels, code[i - labelLength:i])
				} else {
					if code[i - labelLength:i] == labels[len(labels) - 1] {
						labels = labels[:len(labels) - 1]
						if nLabels == 0 {
							labelEnds  = true
						}
					} else {
						return false
					}
				}
				labelLength = 0
				parseLabel = false
			} else if code[i] < 'A' || code[i] > 'Z' {		// invalid label content
				return false
			} else {										// valid label content
				labelLength ++
			}
		}

		// remove cdata : <![CDATA[ + ... + ]]>
		if code[i] == '<' {									// count labels
			if i + 9 < len(code) {
				if code[i + 1:i + 9] == "![CDATA[" && len(labels) != 0 {
					for j := i + 9; j < len(code) - 3; j++ {
						if code [j:j + 3] == "]]>" {		// remove the cdata
							code = code[:i] + code[j + 3:]
							i = i - 2
							break
						}
					}
				}
			}

			if code[i] == '<' {
				if code[i + 1] != '/' {							// meet op labels
					parseOP = true
					nLabels ++
				} else {										// meet closing labels
					parseOP = false
					nLabels --
				}
				parseLabel = true								// start parsing labels
			}
		}

		if nLabels < 0 {									// illegally surrounded by labels
			return false
		}
	}

	if nLabels != 0 {										// not closed
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
			if i + 1 >= length {
				return false
			} else if code[i + 1] == '/' {							// meet an ed label
				if len(labels) == 0 {								// no op label
					return false
				} else {
					lenLabel, label := 1, ""
					for code[i + 1 + lenLabel] >= 'A' && code[i + 1 + lenLabel] <= 'Z' {
						label = label + string(code[i + 1 + lenLabel])
						lenLabel++
					}
					if code[i + 1 + lenLabel] == '>' {				// pop the last op label
						if labels[len(labels) - 1] == label {
							labels = labels[0:len(labels) - 1]
							i = i + 1 + lenLabel
							continue
						} else {
							return false							// different op label
						}
					} else {										// illegal ed label
						return false
					}
				}
			} else if code[i + 1] == '!' {							// meet a cdata
				if i + 12 < len(code) {
					if code[i + 1:i + 9] == "![CDATA[" {
						for j := i + 9; j + 2 < len(code); j++ {
							if code[j:j + 3] == "]]>" {
								i = j + 2
								break
							}
						}
					}
				}
				return false										// illegal cdata
			} else if code[i + 1] >= 'A' && code[i + 1] <= 'Z' {	// meet an op label
				lenLabel, label := 1, ""
				for i + lenLabel < length && code[i + lenLabel] >= 'A' && code[i + lenLabel] <= 'Z' {
					label = label + string(1)		// NOT VALID
					lenLabel++
				}
				if code[i + 1 + lenLabel] == '>' {
					labels = append(labels, label)
				} else {
					return false									// illegal op label
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
				return str[i + 1:]
			}
		}
		return ""
	}

	var amputeComp func(str1, str2 string) string			// return if str2 >=< str1 without tag
	amputeComp = func(str1, str2 string) string {
		temp1, temp2 := ampute(str1), ampute(str2)

		if temp1 > temp2 {
			return ">"
		} else if temp1 ==  temp2 {
			return "="
		} else {
			return "<"
		}
	}

	for index, strs := range logs{
		fmt.Println(strs, logs)
		if ampute(strs)[0] > '9' {								// reorder if meet a letter
			switchIndex := index
			for i := index - 1; i >= 0; i-- {					// compare with the previous, if ok, switch
				 if ampute(logs[i])[0] <= '9' || amputeComp(strs, logs[i]) == "<" || amputeComp(strs, logs[i]) == "=" && strs < logs[i]{
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
		list = append(list[:elim], list[elim + 1:]...)
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

func (this *RecentCounter) Ping(t int) int {		// ping : 收到一个请求。在每次收到请求的时候都要返回三秒内的请求数
	*this = append(*this, t)						// 入队列
	for (*this)[0] < t - 3000 {
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

		if nums[temp - 1] < 0 {
			result = append(result, temp)
		}
		nums[temp - 1] = -nums[temp - 1]
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

	if result[len(s) - 1] == left {
		result = append(result, right)
	} else {
		result = append(result, left)
	}

	return result
}


