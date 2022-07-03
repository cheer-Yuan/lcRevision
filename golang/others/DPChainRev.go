package others

type node struct {
	index int
	next *node
}

func ChainRev(head node) node {
	var cur, pre *node
	cur = &head
	pre = nil

	for cur != nil {
		temp := cur.next
		cur.next = pre
		pre = cur
		cur = temp
	}

	return *pre
}