package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
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



type QueueOfTreeNode struct {
	queue []*TreeNode
}

func (thisQueue *QueueOfTreeNode) push (element *TreeNode) {
	thisQueue.queue = append(thisQueue.queue, element)
}

func (thisQueue *QueueOfTreeNode) pop() *TreeNode {
	buff := thisQueue.peek()
	thisQueue.queue = thisQueue.queue[1 :len(thisQueue.queue)]
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