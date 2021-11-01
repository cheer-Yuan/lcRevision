package main

type TreeNode struct {
	value int
	left, right *TreeNode
}


type StackOfTreeNode struct {
	stack []*TreeNode
}



func (thisStack *StackOfTreeNode) push (element *TreeNode) {
	thisStack.stack = append(thisStack.stack, element)
}

func (thisStack *StackOfTreeNode) pop() *TreeNode {
	buff := thisStack.peek()
	thisStack.stack = thisStack.stack[:len(thisStack.stack) -1]
	return buff
}

func (thisStack *StackOfTreeNode) peek() *TreeNode {
	return thisStack.stack[len(thisStack.stack) - 1]
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