package main

//递归遍历二叉树, 存储到values列表
func TreeTraversePreorder(treeNode *TreeNode) []int {
	var traverse func(node *TreeNode)
	results := []int{}

	traverse = func(node *TreeNode) {
		if (node == nil) {return}
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
		node := stack.peek()	// 处理栈中当前节点
		if node != nil {		// 如节点非空 ：
			_ = stack.pop()			// 删除该节点，避免重复
			if node.Right != nil {	// 添加非空右节点
				stack.push(node.Right)
			}
			stack.push(node)
			stack.push(nil)	// 空节点作为提示符，向结果列表中添加该节点
			if node.Left != nil {	// 添加非空左节点
				stack.push(node.Left)
			}
		} else {				// 空节点：将空节点前的节点加入结果列表
			_ = stack.pop()
			ListOfNodes = append(ListOfNodes, stack.pop().Val)
		}
	}

	return ListOfNodes
}

// 迭代法前序遍历
func TreeTraverseIterativePreorder(root *TreeNode) []int {
	ListOfNodes := []int{}

	stack := new(StackOfTreeNode)	// 创建一个节点的栈

	if root != nil {
		stack.push(root)			// 栈：推入根节点
	}

	for !stack.isEmpty() {			// 只要栈内非空：
		node := stack.peek()		// 取栈上值
		if node != nil {
			_ = stack.pop()        // 从栈上删除该节点
			if node.Right != nil { // 添加非空右节点
				stack.push(node.Right) // 推右节点入栈
			}
			if node.Left != nil { // 添加非空右节点
				stack.push(node.Left) // 推左节点入栈
			}
			stack.push(node)       // 需要在结果列表前的越往下
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
	stack := new(StackOfTreeNode)	// 创建一个节点的栈

	if root != nil {				// 栈：推入根节点
		stack.push(root)
	}

	for !stack.isEmpty() {			// 只要栈内非空：
		node := stack.peek()		// 取栈上值
		if node != nil {
			_ = stack.pop()			// 从栈上删除该节点
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
	queue.push(root)				// 根入队列

	for !queue.isEmpty() {
		NumOfThisLayer := queue.size()		// 记录本层节点数
		subresult := []int{}
		for i := 0; i < NumOfThisLayer; i++ {
			node := queue.pop()
			subresult = append(subresult, node.Val)		// 本层节点加入结果集
			if node.Left != nil {						// 节点的左右子节点加入队列，等待下一轮加入结果集
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


