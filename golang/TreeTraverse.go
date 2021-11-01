package main
//
////递归遍历二叉树, 存储到values列表
//func TreeTraversePreorder(treeNode *TreeNode, Values []int) {
//	if (treeNode == nil) {return}
//
//	//前序
//	Values = append(Values, treeNode.value)		//根节点
//	TreeTraversePreorder(treeNode.left, Values)			//左子树
//	TreeTraversePreorder(treeNode.right, Values)		//右子树
//}
//
//func TreeTraverseInorder(treeNode *TreeNode, Values []int) {
//	if (treeNode == nil) {return}
//
//	//中序
//	TreeTraverseInorder(treeNode.left, Values)			//左子树
//	Values = append(Values, treeNode.value)		//根节点
//	TreeTraverseInorder(treeNode.right, Values)		//右子树
//}
//
//func TreeTraversePostorder(treeNode *TreeNode, Values []int) {
//	if (treeNode == nil) {return}
//
//	//后序
//	TreeTraversePostorder(treeNode.left, Values)			//左子树
//	TreeTraversePostorder(treeNode.right, Values)		//右子树
//	Values = append(Values, treeNode.value)		//根节点
//}
//
//func TreeTraverseRecur(root *TreeNode) []int {
//	results := make([]int, 0)
//
//	TreeTraversePreorder(root, results)
//
//	return  results
//}
//
//
//
//// iterative traverse : 入栈：根，右，左
//func TreeTraverseIterativeMidorder(root *TreeNode, ListOfNodes []*TreeNode) {
//	// 先加入根节点
//	stack := new(StackOfTreeNode)
//	if root != nil {
//		 stack.push(root)
//	 }
//
//	for !stack.isEmpty() {
//		node := stack.peek()	// 处理栈中当前节点
//		if node != nil {		// 如节点非空 ：
//			_ = stack.pop()			// 删除该节点，避免重复
//			stack.push(node.right)
//			stack.push(node)
//			stack.push(nil)	// 空节点作为提示符，向结果列表中添加该节点
//			stack.push(node.left)
//		} else {				// 空节点：将空节点前的节点加入结果列表
//			_ = stack.pop()
//			ListOfNodes = append(ListOfNodes, stack.pop())
//		}
//	}
//}
//
//func TreeTraverseIterativeInorder(root *TreeNode, ListOfNodes []*TreeNode) {
//	stack := new(StackOfTreeNode)
//	if root != nil {
//		stack.push(root)
//	}
//
//	for !stack.isEmpty() {
//		node := stack.peek()
//		if node != nil {
//			_ = stack.pop()
//			stack.push(node.right)
//			stack.push(node.left)
//			stack.push(node)			// 需要在结果列表前的越往下
//			stack.push(nil)
//		} else {
//			_ = stack.pop()
//			ListOfNodes = append(ListOfNodes, stack.pop())
//		}
//	}
//}
//
//func TreeTraverseIterativePostorder(root *TreeNode, ListOfNodes []*TreeNode) {
//	stack := new(StackOfTreeNode)
//	if root != nil {
//		stack.push(root)
//	}
//
//	for !stack.isEmpty() {+
//		node := stack.peek()
//		if node != nil {
//			_ = stack.pop()
//			stack.push(node)
//			stack.push(nil)
//			stack.push(node.right)
//			stack.push(node.left)
//		} else {
//			_ = stack.pop()
//			ListOfNodes = append(ListOfNodes, stack.pop())
//		}
//	}
//}
//
//func ()  {
//
//}
