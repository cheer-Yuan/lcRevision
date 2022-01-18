package main

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

	inorderTraverse = func(node *TreeNode)  {
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
		if inOrderSlice[index] > inOrderSlice[index - 1] {
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
				result = MinOf2(result, MaxOf2(node.Val, prevNode.Val) - MinOf2(node.Val, prevNode.Val))
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
		var result []int						// 存储结果
		var prevNode *TreeNode					// 存储上一个节点
		var iOT func(node *TreeNode)
		var count, maxCount int					// 迭代的当前频次，最大频次
		result = append(result, root.Val)

		maxCount = 0
		iOT = func(node *TreeNode) {			// 中序遍历迭代
			if node != nil {
				iOT(node.Left)					// 左

				if prevNode == nil {			// 中：处理节点
					count = 1
				} else if node.Val == prevNode.Val {
					count ++
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

	if root.Val > p.Val && root.Val > q.Val {	// 向左搜索
		left := lowestCommonAncestor(root.Left, p, q)
		if left != nil {	// 搜索一条边的写法：如果返回值不为空，立刻返回
			return left
		}
	}

	if root.Val < p.Val && root.Val < q.Val {	// 向右搜索
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
		Val: val,
		Left: nil,
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
		return root		// 遇到空节点
	}

	// 删除操作
	if root.Val == key {
		if root.Left == nil && root.Right == nil {	// 删除叶子
			return nil
		} else if root.Left == nil {	// 左子节点为空
			return root.Right
		} else if root.Right == nil {	// 右子节点为空
			return root.Left
		} else {						// 左右子节点都不为空
			node := root.Right			// 找右子树最左边的子节点
			for node.Left != nil {
				node = node.Left
			}
			node.Left = root.Left			// 在这里放置要删除的节点的左子节点
			root = root.Right			// 返回要删除的节点的右子节点
			return root
		}
	}


	// 迭代处理左右节点
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right,key)
	}

	return root
}