package others

//使用栈实现队列的下列操作：
//
//push(x) -- 将一个元素放入队列的尾部。
//pop() -- 从队列首部移除元素。
//peek() -- 返回队列首部的元素。
//empty() -- 返回队列是否为空。

//用栈实现队列
type QueueByStack struct {
	inStack  Stack
	outStack Stack
}

// 放入queue
func (thisQueue *QueueByStack) push(element int) {
	thisQueue.inStack.push(element)
}

// 删除最上方元素， 返回值给peek使用
func (thisQueue *QueueByStack) pop() int {
	// 双队列，返回出队列最上方的值
	if !thisQueue.outStack.isEmpty() {
		return thisQueue.outStack.pop()
	} else {
		for !thisQueue.inStack.isEmpty() {
			thisQueue.outStack.push(thisQueue.inStack.peek())
			thisQueue.inStack.pop()
		}
		//fmt.Println(len(thisQueue.outStack.stack))
		return thisQueue.outStack.pop()

	}
}

// 查询最上方元素
func (thisQueue *QueueByStack) peek() int {
	buff := thisQueue.pop()
	thisQueue.outStack.push(buff)
	return buff
}

// 查询栈是否为空
func (thisQueue *QueueByStack) isEmpty() bool {
	if thisQueue.inStack.size() == 0 && thisQueue.outStack.size() == 0 {
		return true
	} else {
		return false
	}
}
