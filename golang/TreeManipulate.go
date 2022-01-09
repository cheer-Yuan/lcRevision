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

func invertTreeRec(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	buff := root.Left
	root.Left = root.Right
	root.Right = buff

	invertTreeRec(root.Left)
	invertTreeRec(root.Right)
	return root
}

func invertTreeIter(root *TreeNode) *TreeNode {

	stack := new(StackOfTreeNode)	// 创建一个节点的栈

	if root != nil {
		stack.push(root)			// 栈：推入根节点
	}

	for !stack.isEmpty() {			// 只要栈内非空：
		node := stack.peek()		// 取栈上值
		if node != nil {
			_ = stack.pop()        // 从栈上删除该节点
			stack.push(node.Right) // 推右节点入栈
			stack.push(node.Left)  // 推左节点入栈
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

*/

//func isSymmetric(root *TreeNode) bool {
//
//}
