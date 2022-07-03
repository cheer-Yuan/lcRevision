package others

func checkValidParanthese(parenthese string) bool {
	if parenthese == "" {
		return true
	}

	a := Stack{stack: make([]int, 0)}
	for i := 0; i < len(parenthese); i++ {
		if parenthese[i] == '[' {
			a.push(']')
		} else if parenthese[i] == '{' {
			a.push('}')
		} else if parenthese[i] == '(' {
			a.push(')')
		} else {
			if a.isEmpty() {
				return false
			}
			if int(parenthese[i]) == a.peek() {
				a.pop()
			} else {
				return false
			}
		}
	}
	return a.isEmpty()
}
