package main

import "strconv"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type StackOfTreeNode struct {
	stack []*TreeNode
}

func (thisStack *StackOfTreeNode) push(element *TreeNode) {
	thisStack.stack = append(thisStack.stack, element)
}

func (thisStack *StackOfTreeNode) pop() *TreeNode {
	buff := thisStack.peek()
	thisStack.stack = thisStack.stack[:len(thisStack.stack)-1]
	return buff
}

func (thisStack *StackOfTreeNode) peek() *TreeNode {
	return thisStack.stack[len(thisStack.stack)-1]
}

func (thisStack *StackOfTreeNode) size() int {
	return len(thisStack.stack)
}

func (thisStack *StackOfTreeNode) isEmpty() bool {
	if len(thisStack.stack) == 0 {
		return true
	} else {
		return false
	}
}

type QueueOfTreeNode struct {
	queue []*TreeNode
}

func (thisQueue *QueueOfTreeNode) push(element *TreeNode) {
	thisQueue.queue = append(thisQueue.queue, element)
}

func (thisQueue *QueueOfTreeNode) pop() *TreeNode {
	buff := thisQueue.peek()
	thisQueue.queue = thisQueue.queue[1:len(thisQueue.queue)]
	return buff
}

func (thisQueue *QueueOfTreeNode) peek() *TreeNode {
	return thisQueue.queue[0]
}

func (thisQueue *QueueOfTreeNode) size() int {
	return len(thisQueue.queue)
}

func (thisQueue *QueueOfTreeNode) isEmpty() bool {
	if len(thisQueue.queue) == 0 {
		return true
	} else {
		return false
	}
}

//递归遍历二叉树, 存储到values列表
func TreeTraversePreorder(treeNode *TreeNode) []int {
	var traverse func(node *TreeNode)
	results := []int{}

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		results = append(results, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(treeNode)
	return results
}

func TreeTraverseInorder(treeNode *TreeNode) []int {
	var traverse func(node *TreeNode)
	results := []int{}

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		results = append(results, node.Val)
		traverse(node.Right)
	}

	traverse(treeNode)
	return results
}

func TreeTraversePostorder(treeNode *TreeNode) []int {
	var traverse func(node *TreeNode)
	results := []int{}

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		traverse(node.Right)
		results = append(results, node.Val)
	}

	traverse(treeNode)
	return results
}

// iterative traverse : 入栈：根，右，左
func TreeTraverseIterativeMidorder(root *TreeNode) []int {
	ListOfNodes := []int{}

	// 先加入根节点
	stack := new(StackOfTreeNode)
	if root != nil {
		stack.push(root)
	}

	for !stack.isEmpty() {
		node := stack.peek() // 处理栈中当前节点
		if node != nil {     // 如节点非空 ：
			_ = stack.pop()        // 删除该节点，避免重复
			if node.Right != nil { // 添加非空右节点
				stack.push(node.Right)
			}
			stack.push(node)
			stack.push(nil)       // 空节点作为提示符，向结果列表中添加该节点
			if node.Left != nil { // 添加非空左节点
				stack.push(node.Left)
			}
		} else { // 空节点：将空节点前的节点加入结果列表
			_ = stack.pop()
			ListOfNodes = append(ListOfNodes, stack.pop().Val)
		}
	}

	return ListOfNodes
}

// 迭代法前序遍历
func TreeTraverseIterativePreorder(root *TreeNode) []int {
	ListOfNodes := []int{}

	stack := new(StackOfTreeNode) // 创建一个节点的栈

	if root != nil {
		stack.push(root) // 栈：推入根节点
	}

	for !stack.isEmpty() { // 只要栈内非空：
		node := stack.peek() // 取栈上值
		if node != nil {
			_ = stack.pop()        // 从栈上删除该节点
			if node.Right != nil { // 添加非空右节点
				stack.push(node.Right) // 推右节点入栈
			}
			if node.Left != nil { // 添加非空右节点
				stack.push(node.Left) // 推左节点入栈
			}
			stack.push(node) // 需要在结果列表前的越往下
			stack.push(nil)
		} else {
			_ = stack.pop()
			ListOfNodes = append(ListOfNodes, stack.pop().Val)
		}
	}

	return ListOfNodes
}

// 迭代法后序遍历
func TreeTraverseIterativePostorder(root *TreeNode) []int {
	ListOfNodes := []int{}
	stack := new(StackOfTreeNode) // 创建一个节点的栈

	if root != nil { // 栈：推入根节点
		stack.push(root)
	}

	for !stack.isEmpty() { // 只要栈内非空：
		node := stack.peek() // 取栈上值
		if node != nil {
			_ = stack.pop() // 从栈上删除该节点
			stack.push(node)
			stack.push(nil)
			if node.Right != nil {
				stack.push(node.Right)
			}
			if node.Left != nil {
				stack.push(node.Left)
			}
		} else {
			_ = stack.pop()
			ListOfNodes = append(ListOfNodes, stack.pop().Val)
		}
	}

	return ListOfNodes
}

/*
层序遍历：广度优先遍历，使用队列实现，FIFO适合一层一层遍历的逻辑
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	queue := new(QueueOfTreeNode)
	queue.push(root) // 根入队列

	for !queue.isEmpty() {
		NumOfThisLayer := queue.size() // 记录本层节点数
		subresult := []int{}
		for i := 0; i < NumOfThisLayer; i++ {
			node := queue.pop()
			subresult = append(subresult, node.Val) // 本层节点加入结果集
			if node.Left != nil {                   // 节点的左右子节点加入队列，等待下一轮加入结果集
				queue.push(node.Left)
			}
			if node.Right != nil {
				queue.push(node.Right)
			}

		}
		result = append(result, subresult)
	}

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
	racine := postorder[len(postorder)-1] // 后序最末元素为根节点值
	// 3. 找到根在中序排列中的位置作为切割点
	index := 0
	for ; index < len(inorder) && inorder[index] != racine; index++ {
	} // 中序切割点
	// 6. 递归处理
	return &TreeNode{
		Val:   racine,
		Left:  buildTree1(inorder[:index], postorder[:index]),                   // 4,
		Right: buildTree1(inorder[index+1:], postorder[index:len(postorder)-1]), // 5, 后续数组除去最后一个元素
	}
}

/*
105. 从前序与中序遍历序列构造二叉树
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	racine := preorder[0]

	index := 0
	for ; index < len(inorder) && inorder[index] != racine; index++ {
	}

	return &TreeNode{
		Val:   racine,
		Left:  buildTree(preorder[1:index+1], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
}

/*654. 最大二叉树
给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：

二叉树的根是数组 nums 中的最大元素。
左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
返回有给定数组 nums 构建的 最大二叉树 。
*/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	index, root := 0, 0
	for i, val := range nums {
		if val > root {
			root = val
			index = i
		}
	}

	return &TreeNode{
		Val:   root,
		Left:  constructMaximumBinaryTree(nums[0:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
}

/*617. 合并二叉树
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为NULL 的节点将直接作为新二叉树的节点。
*/

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	var r1l, r2l, r1r, r2r *TreeNode
	value := 0
	if root1 != nil {
		value += root1.Val
		r1l, r1r = root1.Left, root1.Right
	} else {
		r1l, r1r = nil, nil
	}
	if root2 != nil {
		value += root2.Val
		r2l, r2r = root2.Left, root2.Right
	} else {
		r2l, r2r = nil, nil
	}

	return &TreeNode{
		Val:   value,
		Left:  mergeTrees(r1l, r2l),
		Right: mergeTrees(r1r, r2r),
	}
}

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
	} // 递归终止条件：当前节点为空

	root.Right, root.Left = root.Left, root.Right // 前序遍历，交换左右节点
	invertTreeRec(root.Left)                      // 递归操作左右子节点
	invertTreeRec(root.Right)
	return root
}

// 迭代翻转二叉树：后序遍历模板
func invertTreeIter(root *TreeNode) *TreeNode {

	stack := new(StackOfTreeNode) // 创建一个节点的栈

	if root != nil {
		stack.push(root) // 栈：推入根节点
	}

	for !stack.isEmpty() { // 只要栈内非空：
		node := stack.peek() // 取栈上值
		if node != nil {
			_ = stack.pop() // 从栈上删除该节点
			if node.Right != nil {
				stack.push(node.Right)
			}
			if node.Left != nil {
				stack.push(node.Left)
			}
			stack.push(node) // 需要在结果列表前的越往下
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
		if Left == nil && Right == nil { // 含有空节点的情况
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
		Left, Right := queue.pop(), queue.pop() // 弹出队列前两个节点
		if Left == nil && Right == nil {        // 前两个节点都为空节点，仍为对称，继续迭代
			continue
		}
		if Left == nil || Right == nil || Left.Val != Right.Val { // 前两个节点有一个非空，或两个值不一样，直接返回假
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
			return (2 << dL) - 1 // 整棵树没有空节点的情况：直接计算节点数
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
		if hL == -1 || hR == -1 || hL > hR+1 || hR > hL+1 {
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

	results := []string{} // 记录答案的切片

	traverse = func(treeNode *TreeNode, result string) {
		//终止条件：找到叶子节点：当前节点和左右子节点不为空。
		if treeNode != nil && treeNode.Left == nil && treeNode.Right == nil {
			result += strconv.Itoa(treeNode.Val) // 加入最后的节点
			results = append(results, result)    // 添加答案
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
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil { // 排除右叶子
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
			if leftLen > maxLen { // 第一次取到最大深度
				maxLen = leftLen   // 更新最大深度
				maxLeft = node.Val // 更新最左边的值
			}
		}

		if node.Left != nil {
			traverse(node.Left, leftLen+1)
		}
		if node.Right != nil {
			traverse(node.Right, leftLen+1)
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
		if node != nil && node.Left == nil && node.Right == nil && node.Val+sum == targetSum {
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
		if node != nil && node.Left == nil && node.Right == nil && node.Val+sum == targetSum {
			temp := []int{}
			for _, val := range path { // 防止值的变化
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

/*236. 二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

思路：自底向上回溯查找：后序遍历
*/
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	// 后序遍历
	if root == nil || root.Val == p.Val || root.Val == q.Val { // 如找到节点：传回该节点
		return root
	}

	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)

	if left != nil && right != nil { // 左右子节点都有值传来：本节点为最近公共祖先
		return root
	}
	if left == nil && right != nil {
		return right
	} else if left != nil && right == nil {
		return left
	} else {
		return nil
	}
}

/*
给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
*/

func searchBST(root *TreeNode, val int) *TreeNode {
	var search func(node *TreeNode, val int)

	var result *TreeNode

	search = func(node *TreeNode, val int) {
		if node != nil {
			if node.Val == val {
				result = node
			} else {
				if node.Left != nil {
					search(node.Left, val)
				}
				if node.Right != nil {
					search(node.Right, val)
				}
			}
		}
	}

	search(root, val)

	return result
}

/* 98. 验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。


思路： 二叉搜索树相关：转换为中序遍历的切片。二叉搜索树 == 中序遍历单调递增

*/

func isValidBST(root *TreeNode) bool {
	var inorderTraverse func(node *TreeNode)
	inOrderSlice := []int{}

	inorderTraverse = func(node *TreeNode) {
		if node != nil {
			inorderTraverse(node.Left)
			inOrderSlice = append(inOrderSlice, node.Val)
			inorderTraverse(node.Right)
		}
	}

	inorderTraverse(root)

	if len(inOrderSlice) <= 1 {
		return true
	}

	for index := 1; index < len(inOrderSlice); index++ {
		if inOrderSlice[index] > inOrderSlice[index-1] {
			continue
		} else {
			return false
		}
	}

	return true
}

/*530. 二叉搜索树的最小绝对差
给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

差值是一个正数，其数值等于两值之差的绝对值。

思路：中序遍历时记录前一个结点, 在中序遍历时进行差值比较
*/
func getMinimumDifference(root *TreeNode) int {
	var inOrderTraverse func(node *TreeNode)
	var prevNode *TreeNode
	var result int
	if root.Left != nil {
		result = MaxOf2(root.Val, root.Left.Val) - MinOf2(root.Val, root.Left.Val)
	} else {
		result = MaxOf2(root.Val, root.Right.Val) - MinOf2(root.Val, root.Right.Val)
	}

	inOrderTraverse = func(node *TreeNode) {
		if node != nil {
			inOrderTraverse(node.Left)
			if prevNode != nil {
				result = MinOf2(result, MaxOf2(node.Val, prevNode.Val)-MinOf2(node.Val, prevNode.Val))
			}
			prevNode = node
			inOrderTraverse(node.Right)
		}
	}

	inOrderTraverse(root)
	return result
}

/*501. 二叉搜索树中的众数
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
*/
func findMode(root *TreeNode) []int {
	var result []int       // 存储结果
	var prevNode *TreeNode // 存储上一个节点
	var iOT func(node *TreeNode)
	var count, maxCount int // 迭代的当前频次，最大频次
	result = append(result, root.Val)

	maxCount = 0
	iOT = func(node *TreeNode) { // 中序遍历迭代
		if node != nil {
			iOT(node.Left) // 左

			if prevNode == nil { // 中：处理节点
				count = 1
			} else if node.Val == prevNode.Val {
				count++
			} else {
				count = 1
			}
			prevNode = node
			if count == maxCount {
				result = append(result, node.Val)
			}
			if count > maxCount {
				maxCount = count
				result = []int{}
				result = append(result, node.Val)
			}

			iOT(node.Right)
		}
	}

	iOT(root)
	return result
}

/*235. 二叉搜索树的最近公共祖先
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

思路：对于有序树，若node.val 属于区间[p， q]，则说明为公共祖先。可使用前序遍历
单层逻辑：if node.val > p and q : 向左递归，相反情况则向右递归
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 无需处理：处理部分空缺

	if root.Val > p.Val && root.Val > q.Val { // 向左搜索
		left := lowestCommonAncestor(root.Left, p, q)
		if left != nil { // 搜索一条边的写法：如果返回值不为空，立刻返回
			return left
		}
	}

	if root.Val < p.Val && root.Val < q.Val { // 向右搜索
		right := lowestCommonAncestor(root.Right, p, q)
		if right != nil {
			return right
		}
	}

	return root
}

/*701.二叉搜索树中的插入操作
给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。
*/

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}

	if root == nil {
		return newNode
	}

	var traverse func(node *TreeNode, val int) *TreeNode

	// 不改变原有结构的插入方法，递归寻找可以插入的空节点
	traverse = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			return node
		}

		if val < node.Val && node.Left != nil {
			return traverse(node.Left, val)
		}
		if val > node.Val && node.Right != nil {
			return traverse(node.Right, val)
		}

		return node
	}

	current := traverse(root, val)

	if val > current.Val {
		current.Right = newNode
	} else {
		current.Left = newNode
	}

	return root
}

/*
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的key对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

思路：
第一种情况：没找到删除的节点，遍历到空节点直接返回了
找到删除的节点
第二种情况：左右孩子都为空（叶子节点），直接删除节点， 返回NULL为根节点
第三种情况：删除节点的左孩子为空，右孩子不为空，删除节点，右孩子补位，返回右孩子为根节点
第四种情况：删除节点的右孩子为空，左孩子不为空，删除节点，左孩子补位，返回左孩子为根节点
第五种情况：左右孩子节点都不为空，则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点。
*/	

func deleteNode(root *TreeNode, key int) *TreeNode {
	// 返回值：更新后的节点

	if root == nil {
		return root // 遇到空节点
	}

	// 删除操作
	if root.Val == key {
		if root.Left == nil && root.Right == nil { // 删除叶子
			return nil
		} else if root.Left == nil { // 左子节点为空
			return root.Right
		} else if root.Right == nil { // 右子节点为空
			return root.Left
		} else { // 左右子节点都不为空
			node := root.Right // 找右子树最左边的子节点
			for node.Left != nil {
				node = node.Left
			}
			node.Left = root.Left // 在这里放置要删除的节点的左子节点
			root = root.Right     // 返回要删除的节点的右子节点
			return root
		}
	}

	// 迭代处理左右节点
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}
