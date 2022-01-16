package main

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
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	racine := preorder[0]

	index := 0
	for ; index < len(inorder) && inorder[index] != racine; index++ {}

	return &TreeNode{
		Val: racine,
		Left: buildTree(preorder[1:index + 1], inorder[:index]),
		Right: buildTree(preorder[index + 1:], inorder[index + 1:]),
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
		Val: root,
		Left: constructMaximumBinaryTree(nums[0:index]),
		Right: constructMaximumBinaryTree(nums[index + 1:]),
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
		Val: value,
		Left: mergeTrees(r1l, r2l),
		Right: mergeTrees(r1r, r2r),
	}
}