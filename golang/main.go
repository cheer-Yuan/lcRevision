package main

import "fmt"

func main()  {
	//string1 := []string{" ", " ", "a", "b", "c", " ", " ", "d", " ", " "}
	//string2 := []string{"a", "b", "c", " ", " ", "d"}
	//string3 := []string{"a", "b", "c", " ", "d"}
	//
	//fmt.Println(WordRev(string1))
	//fmt.Println(WordRev(string2))
	//fmt.Println(WordRev(string3))

///*	reverse a chain, show the new head*/
//	c := node{index: 2, next: nil}
//	b := node{index: 1, next: &c}
//	a := node{index: 0, next: &b}
//	fmt.Println(ChainRev(a).index)

	/* chain: find the entrance of a ring */
	//i := node{index: 8, next: nil}
	//h := node{index: 7, next: &i}
	//g := node{index: 6, next: &h}
	//f := node{index: 5, next: &g}
	//e := node{index: 4, next: &f}
	//d := node{index: 3, next: &e}
	//c := node{index: 2, next: &d}
	//b := node{index: 1, next: &c}
	//a := node{index: 0, next: &b}
	//i.next = &e
	//entrance := RingEntrance(&a)
	//fmt.Println(entrance)

	///* DPSum3Num */
	//a := []int{-1, 0, 1, 2, -1, -4}


	/*stack : make a queue by stack*/
	//a := QueueByStack{
	//	inStack:  StackOfInt{stack: make([]int, 0)}	,
	//	outStack: StackOfInt{stack: make([]int, 0)},
	//}
	//a.push(3)
	//a.push(2)
	//a.push(1)
	//// in queue : 1->2->3
	//fmt.Println(a.peek())
	//fmt.Println(a.pop())
	//fmt.Println(a.peek())
	//fmt.Println(a.pop())
	//fmt.Println(a.pop())
	////OK

	/*stack : make a stack by queue*/
	//a := StackByQueue{
	//	queue: QueueOfInt{make([]int, 0)},
	//	queueBuff: QueueOfInt{make([]int, 0)},
	//}
	//a.push(3)
	//a.push(2)
	//a.push(1)
	//// in queue : 1->2->3
	//fmt.Println(a.peek())
	//fmt.Println(a.pop())
	//fmt.Println(a.peek())
	//fmt.Println(a.pop())
	//fmt.Println(a.pop())
	////OK

	/*stack : check if is a valid parathense : (){}[]*/
	//	示例 4:
	//输入: "([)]"
	//输出: false
	//
	//	示例 5:
	//输入: "{[]}"
	//输出: true
	//fmt.Println(
	//	checkValidParanthese("{}"),
	//	checkValidParanthese("()[]{}"),
	//	checkValidParanthese("{]"),
	//	checkValidParanthese("([)]"),
	//	checkValidParanthese("{[]}"),
	//	checkValidParanthese("]["),
	//	checkValidParanthese("{}["),
	//	checkValidParanthese("()[]{}}"),
	//)
	////OK

	fmt.Println(integerBreak(5))
}

