package others

/*
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0开头。
*/

type ListNode struct {
    Val int
	Next *ListNode
}

func newNode(val int, nextNode *ListNode) *ListNode {
	return &ListNode{Val: val, Next: nextNode}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	L1, L2 := []int{}, []int{}
	buff := l1
	for buff != nil {
		L1 = append(L1, buff.Val)
		buff = buff.Next
	}
	buff = l2
	for buff != nil {
		L2 = append(L2, buff.Val)
		buff = buff.Next
	}

	r := ListNode{0, nil}		// 根节点
	Result := &r						// 变量存储根节点地址，用于迭代

	i, j, avance := 0, 0, 0
	for i < len(L1) || j < len(L2) || avance != 0 {
		a, b, sum := 0, 0, 0
		if i < len(L1) {
			a = L1[i]
			i++
		}
		if j < len(L2) {
			b = L2[j]
			j++
		}
		sum = a + b + avance
		if sum >= 10 {
			sum -= 10
			avance = 1
		} else {
			avance = 0
		}

		Result.Val = sum									// 更新本节点数值
		if i < len(L1) || j < len(L2) || avance != 0 {		// 新建节点，连接到本节点
			Result.Next = newNode(0, nil)		// 更新本节点地址
		}
		Result = Result.Next
	}
	return &r
}