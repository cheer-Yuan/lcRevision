package others

//使用队列实现栈的下列操作：
//
//push(x) -- 元素 x 入栈
//pop() -- 移除栈顶元素
//top() -- 获取栈顶元素
//empty() -- 返回栈是否为空

type StackByQueue struct {
	queue QueueOfInt
	queueBuff QueueOfInt
}

// 放入stack
func (thisStack *StackByQueue) push(element int) {
	thisStack.queue.queue = append(thisStack.queue.queue, element)
}

// 删除最上方元素， 返回值给peek使用
func (thisStack *StackByQueue) pop() int {
	for thisStack.queue.size() != 1 {
		thisStack.queueBuff.queue = append(thisStack.queueBuff.queue, thisStack.queue.pop())
	}
	buff := thisStack.queue.pop()
	thisStack.queue.queue = thisStack.queueBuff.queue
	for thisStack.queueBuff.size() != 0 {
		thisStack.queueBuff.pop()
	}
	return buff
}

// 查询最上方元素
func (thisStack *StackByQueue) peek() int {
	buff := thisStack.pop()
	thisStack.push(buff)
	return buff
}

func (thisStack *StackByQueue) isEmpty() bool {
	if thisStack.queue.isEmpty() {
		return true
	} else {
		return false
	}
}