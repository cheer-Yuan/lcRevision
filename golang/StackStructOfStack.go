package main

//栈

type StackOfInt struct {
	stack []int
}

func (thisStack *StackOfInt) push(element int) {
	thisStack.stack = append(thisStack.stack, element)
}

func (thisStack *StackOfInt) pop() int {
	buff := thisStack.peek()
	thisStack.stack = thisStack.stack[:len(thisStack.stack)-1]
	return buff
}

func (thisStack *StackOfInt) peek() int {
	if len(thisStack.stack) == 0 {
		return 0
	}
	return thisStack.stack[len(thisStack.stack)-1]
}

func (thisStack *StackOfInt) size() int {
	return len(thisStack.stack)

}

func (thisStack *StackOfInt) isEmpty() bool {
	if len(thisStack.stack) == 0 {
		return true
	} else {
		return false
	}
}
