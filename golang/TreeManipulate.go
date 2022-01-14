package main

import (
	"strconv"
)

/*
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/
// 递归翻转二叉树：
func invertTreeRec(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}		// 递归终止条件：当前节点为空

	root.Right, root.Left = root.Left, root.Right	// 前序遍历，交换左右节点
	invertTreeRec(root.Left)						// 递归操作左右子节点
	invertTreeRec(root.Right)
	return root
}

// 迭代翻转二叉树：后序遍历模板
func invertTreeIter(root *TreeNode) *TreeNode {

	stack := new(StackOfTreeNode)	// 创建一个节点的栈

	if root != nil {
		stack.push(root)			// 栈：推入根节点
	}

	for !stack.isEmpty() {			// 只要栈内非空：
		node := stack.peek()		// 取栈上值
		if node != nil {
			_ = stack.pop()        // 从栈上删除该节点
			if node.Right != nil {
				stack.push(node.Right)
			}
			if node.Left != nil {
				stack.push(node.Left)
			}
			stack.push(node)       // 需要在结果列表前的越往下
			stack.push(nil)
		} else {
			_ = stack.pop()
			node = stack.pop()
			node.Left, node.Right = node.Right, node.Left
		}
	}

	return root
}


/*
给定一个二叉树，检查它是否是镜像对称的。
*/

// 递归判断是否对称
func isSymmetricRec(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 递归：左节点的左子节点-又节点的右子节点
	var compare func(Left, Right *TreeNode) bool
	compare = func(Left, Right *TreeNode) bool {
		if Left == nil && Right == nil {			// 含有空节点的情况
			return true
		} else if Left == nil && Right != nil {
			return false
		} else if Left != nil && Right == nil {
			return false
		} else if Left.Val != Right.Val {
			return false
		}

		// 递归操作
		outside := compare(Left.Left, Right.Right)
		inside := compare(Left.Right, Right.Left)
		return outside && inside
	}

	return compare(root.Left, root.Right)
}

// 迭代法判断是否对称：使用队列或栈
func isSymmetricIter(root *TreeNode) bool {
	// 空树肯定对称
	if root == nil {
		return true
	}

	// 推根节点的左右节点入队列
	var queue QueueOfTreeNode
	queue.push(root.Left)
	queue.push(root.Right)

	// 迭代判断模块
	for !queue.isEmpty() {
		Left, Right := queue.pop(), queue.pop()		// 弹出队列前两个节点
		if Left == nil && Right == nil {			// 前两个节点都为空节点，仍为对称，继续迭代
			continue
		}
		if Left == nil || Right == nil || Left.Val != Right.Val {	// 前两个节点有一个非空，或两个值不一样，直接返回假
			return false
		}

		// 入队列操作: 外+外，内+内
		queue.push(Left.Left)
		queue.push(Right.Right)
		queue.push(Left.Right)
		queue.push(Right.Left)
	}

	return true
}


/*
给定一个二叉树，找出其最大深度&最小深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
*/

// 递归求最大深度：后序递归求根节点的高度
func maxDepth(root *TreeNode) int {
	// 递归函数：输入根节点，返回高度
	var getdepth func(treeNode *TreeNode) int

	getdepth = func(treeNode *TreeNode) int {
		if treeNode == nil {
			return 0
		}

		// 递归逻辑：求左右子树深度，取最大值加1
		lD := getdepth(treeNode.Left)
		rD := getdepth(treeNode.Right)
		return MaxOf2(lD, rD) + 1
	}

	return getdepth(root)
}

func minDepth(root *TreeNode) int {
	// 递归函数：输入根节点，返回高度
	var getdepth func(treeNode *TreeNode) int

	getdepth = func(treeNode *TreeNode) int {
		if treeNode == nil {
			return 0
		}

		lD := getdepth(treeNode.Left)
		rD := getdepth(treeNode.Right)

		// 递归逻辑：左子树为空则最小深度为右子树深度+1，右子树同理
		if treeNode.Left == nil && treeNode.Right != nil {
			return rD + 1
		}
		if treeNode.Left != nil && treeNode.Right == nil {
			return lD + 1
		}
		return MinOf2(lD, rD) + 1
	}

	return getdepth(root)
}


/*
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~2h个节点。
*/
// 利用完全二叉树性质求节点数
func countNodes(root *TreeNode) int {
	var countSubTree func(node *TreeNode) int

	// 递归求左右子树的深度
	countSubTree = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		L, R := node.Left, node.Right
		dL, dR := 0, 0

		for L != nil {
			L = L.Left
			dL++
		}
		for R != nil {
			R = R.Right
			dR++
		}
		if dL == dR {
			return (2 << dL) - 1	// 整棵树没有空节点的情况：直接计算节点数
		}

		// 有空节点的情况, 分别计算左右子树的节点数
		return countSubTree(node.Left) + countSubTree(node.Right) + 1
	}

	return countSubTree(root)
}


/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
*/
// 后序遍历，递归逻辑：分别求出其左右子树的高度，然后如果差值小于等于1，则返回当前二叉树的高度，否则则返回-1，表示已经不是二叉平衡树了。
func isBalanced(root *TreeNode) bool {
	var balancedHeight func(treeNode *TreeNode) int

	balancedHeight = func(treeNode *TreeNode) int {
		if treeNode == nil {
			return 0
		}

		hL := balancedHeight(treeNode.Left)
		hR := balancedHeight(treeNode.Right)

		// 非平衡，返回-1
		if hL == -1 || hR == -1 || hL > hR + 1 || hR > hL + 1 {
			return -1
		}

		return 1 + MaxOf2(hL, hR)
	}

	// 高度为-1 ：非平衡，否则为平衡
	if balancedHeight(root) == -1 {
		return false
	} else {
		return true
	}
}


/*
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
*/
// 前序遍历递归与回溯
func binaryTreePaths(root *TreeNode) []string {
	// 递归函数：传入根节点。参数：子节点，记录过的路径
	var traverse func(treeNode *TreeNode, result string)

	results := []string{}	// 记录答案的切片

	traverse = func(treeNode *TreeNode, result string) {
		//终止条件：找到叶子节点：当前节点和左右子节点不为空。
		if treeNode != nil && treeNode.Left == nil && treeNode.Right == nil {
			result += strconv.Itoa(treeNode.Val)	// 加入最后的节点
			results = append(results, result)	// 添加答案
		}
		result += strconv.Itoa(treeNode.Val) + "->" // 加入现节点

		// 迭代左右子节点
		if treeNode.Left != nil {
			traverse(treeNode.Left, result)
		}
		if treeNode.Right != nil {
			traverse(treeNode.Right, result)
		}
	}

	traverse(root, "")
	return results
}


/*
计算给定二叉树的所有左叶子之和。
*/
// 递归后序遍历，因为有返回值：左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	var sumLL func(node *TreeNode) int

	sumLL = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lS, lR := sumLL(node.Left), sumLL(node.Right) // 左，右

		// 中
		sum := 0
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {	// 排除右叶子
			sum += node.Left.Val
		}
		return lS + lR + sum
	}

	return sumLL(root)
}


/*
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
*/
// 递归法：求深度：遍历一整棵树：前序，无返回值
func findBottomLeftValue(root *TreeNode) int {
	maxLen, maxLeft := 0, 0

	var traverse func(node *TreeNode, leftLen int)

	traverse = func(node *TreeNode, leftLen int) {
		// 终止条件：遇到叶子节点时更新深度
		if node.Left == nil && node.Right == nil {
			if leftLen > maxLen {	// 第一次取到最大深度
				maxLen = leftLen	// 更新最大深度
				maxLeft = node.Val	// 更新最左边的值
			}
		}

		if node.Left != nil {
			traverse(node.Left, leftLen + 1)
		}
		if node.Right != nil {
			traverse(node.Right, leftLen + 1)
		}
	}

	traverse(root, 1)

	return maxLeft
}

/*112. 路径总和
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:  给定如下二叉树，以及目标和 sum = 22，
*/
// 无返回值，整棵树 -> 前序遍历
func hasPathSum1(root *TreeNode, targetSum int) bool {
	var traverse func(node *TreeNode, sum int)

	result := false

	traverse = func(node *TreeNode, sum int) {
		if node != nil && node.Left == nil && node.Right == nil && node.Val + sum == targetSum {
			result = true
		}

		if node != nil {
			sum += node.Val

			if node.Left != nil {
				traverse(node.Left, sum)
			}
			if node.Right != nil {
				traverse(node.Right, sum)
			}
		}

	}
	traverse(root, 0)

	return result
}

/*路径总和 II
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。
*/
func hasPathSum(root *TreeNode, targetSum int) [][]int {
	var traverse func(node *TreeNode, path []int, sum int)

	result := [][]int{}

	traverse = func(node *TreeNode, path []int, sum int) {
		if node != nil && node.Left == nil && node.Right == nil && node.Val + sum == targetSum {
			temp := []int{}
			for _, val := range path {		// 防止值的变化
				temp = append(temp, val)
			}
			temp = append(temp, node.Val)
			result = append(result, temp)
		}

		if node != nil {
			sum += node.Val
			path = append(path, node.Val)
			if node.Left != nil {
				traverse(node.Left, path, sum)
			}
			if node.Right != nil {
				traverse(node.Right, path, sum)
			}
		}

	}
	traverse(root, []int{}, 0)

	return result
}

/*
106. 从中序与后序遍历序列构造二叉树

首先回忆一下如何根据两个顺序构造一个唯一的二叉树，相信理论知识大家应该都清楚，就是以 后序数组的最后一个元素为切割点，先切中序数组，根据中序数组，反过来在切后序数组。一层一层切下去，每次后序数组最后一个元素就是节点元素。

第一步：如果数组大小为零的话，说明是空节点了。

第二步：如果不为空，那么取后序数组最后一个元素作为节点元素。

第三步：找到后序数组最后一个元素在中序数组的位置，作为切割点

第四步：切割中序数组，切成中序左数组和中序右数组 （顺序别搞反了，一定是先切中序数组）

第五步：切割后序数组，切成后序左数组和后序右数组

第六步：递归处理左区间和右区间
*/

func buildTree1(inorder []int, postorder []int) *TreeNode {
	// 1. 数组长度为零：空节点
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 2. 取后序最后节点为根
	racine := postorder[len(postorder) - 1]	// 后序最末元素为根节点值
	// 3. 找到根在中序排列中的位置作为切割点
	index := 0
	for ; index < len(inorder) && inorder[index] != racine; index++ {}	// 中序切割点
	// 6. 递归处理
	return &TreeNode{
		Val: racine,
		Left: buildTree1(inorder[:index], postorder[:index]),	// 4,
		Right: buildTree1(inorder[index + 1:], postorder[index:len(postorder) - 1]),	// 5, 后续数组除去最后一个元素
	}
}


/*
105. 从前序与中序遍历序列构造二叉树
*/
func buildTree(preorder []int, inorder []int) *TreeNode {

}