package main

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

示例 2：

输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。 偷窃到的最高金额 = 2 + 9 + 1 = 12 。

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 400

思路：
dp[i]：i号房及之前可以偷到的最高金额。来源：dp[i - 1]
	偷i - 1：必定不偷i：dp[i] = dp[i - 1]
	不偷dp[i - 1]：dp[i] = dp[i - 2] + i
综上：dp[i] = max(dp[i - 1], dp[i - 2] + i)
初始化：dp[1] = nums[1], dp[0] = 0
*/

// 可优化空间复杂度
func rob1(nums []int) int {
	leng := len(nums)
	dp := make([]int, leng + 1)
	dp[0], dp[1] = 0, nums[0]

	if leng == 1 {
		return nums[0]
	}

	for i := 2; i <= leng; i++ {
		dp[i] = MaxOf2(dp[i - 1], dp[i - 2] + nums[i - 1])
	}

	return dp[leng]
}

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

分别考虑包含头/尾情况
*/

func rob2(nums []int) int {
	leng := len(nums)
	if leng == 1 {
		return nums[0]
	} else if leng == 2 {
		return MaxOf2(nums[0], nums[1])
	}

	TtAmount0, TtAmount1 := 0, 0
	dp1, dp2 := make([]int, leng), make([]int, leng)

	// 不考虑尾元素
	dp1[0], dp1[1] = 0, nums[0]
	for i := 2; i <= leng - 1; i++ {
		dp1[i] = MaxOf2(dp1[i - 1], dp1[i - 2] + nums[i - 1])
	}
	TtAmount0 = dp1[leng - 1]

	//不考虑首元素
	dp2[0], dp2[1] = 0, nums[1]
	for i := 2; i <= leng - 1; i++ {
		dp2[i] = MaxOf2(dp2[i - 1], dp2[i - 2] + nums[i])
	}
	TtAmount1 = dp2[leng - 1]

	return MaxOf2(TtAmount0, TtAmount1)
}

/*
在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

输入: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

输出: 9
解释:小偷一晚能够盗取的最高金额= 4 + 5 = 9.

思路：
本题一定是要后序遍历，因为通过递归函数的返回值来做下一步计算。
对于某个节点：dp[0]: 不偷，dp[1]：偷
	不偷 --> 左右都可以偷：dp[0] = max(Left[0], Left[1]) + max(Right[0], Right[1])
	偷：不偷左右子节点 --> dp[1] = node.val + Left[0] + Right[0]
*/



/*type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}*/

func rob(root *TreeNode) int {
	result := robNode(root)
	return MaxOf2(result[0], result[1])
}

// 参与遍历树
func robNode(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}

	//遍历
	left := robNode(node.Left)
	right := robNode(node.Right)

	//偷当前：左右不偷
	return []int{MaxOf2(left[0], left[1]) + MaxOf2(right[0], right[1]), node.Val + left[0] + right[0]}

}