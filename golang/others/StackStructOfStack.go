package others

//栈

type Stack struct {
	stack []interface{}
}

func (thisStack *Stack) push(element interface{}) {
	thisStack.stack = append(thisStack.stack, element)
}

func (thisStack *Stack) pop() interface{} {
	buff := thisStack.peek()
	thisStack.stack = thisStack.stack[:len(thisStack.stack)-1]
	return buff
}

func (thisStack *Stack) peek() interface{} {
	if len(thisStack.stack) == 0 {
		return 0
	}
	return thisStack.stack[len(thisStack.stack)-1]
}

func (thisStack *Stack) size() int {
	return len(thisStack.stack)

}

func (thisStack *Stack) isEmpty() bool {
	if len(thisStack.stack) == 0 {
		return true
	} else {
		return false
	}
}
