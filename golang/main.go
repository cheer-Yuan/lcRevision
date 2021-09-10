package main

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
	i := node{index: 2, next: nil}
	h := node{index: 1, next: &i}
	g := node{index: 0, next: &h}
	f := node{index: 2, next: &g}
	e := node{index: 1, next: &f}
	d := node{index: 0, next: &e}
	c := node{index: 2, next: &d}
	b := node{index: 1, next: &c}
	a := node{index: 0, next: &b}
	i.next = &e


}

