package notes

//栈
type StackOfInt struct {
	stack []int
}

func (thisStack StackOfInt) push (element int) {
	thisStack.stack = append(thisStack.stack, element)
}

func (thisStack StackOfInt) pop() int {
	buff := thisStack.peek()
	thisStack.stack = thisStack.stack[:len(thisStack.stack) -1]
	return buff
}

func (thisStack StackOfInt) peek() int {
	return thisStack.stack[len(thisStack.stack) - 1]
}

func (thisStack StackOfInt) size() int {
	return len(thisStack.stack)

}

func (thisStack StackOfInt) isEmpty() bool {
	if len(thisStack.stack) == 0 {
		return true
	} else {
		return false
	}
}

//使用栈实现队列的下列操作：
//
//push(x) -- 将一个元素放入队列的尾部。
//pop() -- 从队列首部移除元素。
//peek() -- 返回队列首部的元素。
//empty() -- 返回队列是否为空。

//用栈实现队列
type QueueByStack struct {
	inStack StackOfInt
	outStack StackOfInt
}

func (thisQueue QueueByStack) push(element int) {
	thisQueue.inStack.stack = append(thisQueue.inStack.stack, element)
}

// pop 返回值给peek使用
func (thisQueue QueueByStack) pop() int {
	if !thisQueue.outStack.isEmpty() {
		return thisQueue.outStack.pop()
	} else {
		for !thisQueue.inStack.isEmpty() {
			thisQueue.outStack.push(thisQueue.inStack.peek())
			return thisQueue.inStack.pop()
		}
	}
}

func (thisQueue QueueByStack) peek() int {
	buff := thisQueue.pop()
	thisQueue.push(buff)
	return buff
}

func (thisQueue QueueByStack) isEmpty() bool {
	if thisQueue.inStack.size() == 0 && thisQueue.outStack.size() == 0 {
		return true
	} else {
		return false
	}
}
