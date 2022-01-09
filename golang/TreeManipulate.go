package main

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

// 迭代法判断是否对称：
func isSymmetricIter(root *TreeNode) bool {

}