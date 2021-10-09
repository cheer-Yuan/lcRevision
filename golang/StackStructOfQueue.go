package main

//队列
type QueueOfInt struct {
	queue []int
}

func (thisQueue *QueueOfInt) push (element int) {
	thisQueue.queue = append(thisQueue.queue, element)
}

func (thisQueue *QueueOfInt) pop() int {
	buff := thisQueue.peek()
	thisQueue.queue = thisQueue.queue[1 : len(thisQueue.queue)]
	return buff
}

func (thisQueue *QueueOfInt) peek() int {
	return thisQueue.queue[0]
}

func (thisQueue *QueueOfInt) size() int {
	return len(thisQueue.queue)

}

func (thisQueue *QueueOfInt) isEmpty() bool {
	if len(thisQueue.queue) == 0 {
		return true
	} else {
		return false
	}
}


