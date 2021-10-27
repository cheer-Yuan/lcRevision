package main

//递归遍历二叉树, 存储到values列表
func TreeTraversePreorder(treeNode *TreeNode, Values []int) {
	if (treeNode == nil) {return}

	//前序
	Values = append(Values, treeNode.value)		//根节点
	TreeTraversePreorder(treeNode.left, Values)			//左子树
	TreeTraversePreorder(treeNode.right, Values)		//右子树
}

func TreeTraverseInorder(treeNode *TreeNode, Values []int) {
	if (treeNode == nil) {return}

	//中序
	TreeTraverseInorder(treeNode.left, Values)			//左子树
	Values = append(Values, treeNode.value)		//根节点
	TreeTraverseInorder(treeNode.right, Values)		//右子树
}

func TreeTraversePostorder(treeNode *TreeNode, Values []int) {
	if (treeNode == nil) {return}

	//后序
	TreeTraversePostorder(treeNode.left, Values)			//左子树
	TreeTraversePostorder(treeNode.right, Values)		//右子树
	Values = append(Values, treeNode.value)		//根节点
}

func TreeTraverseRecur(root *TreeNode) []int {
	results := make([]int, 0)

	TreeTraversePreorder(root, results)

	return  results
}



// iterative traverse : 入栈：根，右，左

