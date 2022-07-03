package others

/* return the first node entering the ring, without changing the list, and null if no ring */
func RingEntrance(head *node) *node {
	sP, fP := head, head
	for fP != nil && fP.next != nil{

		// avance the slow and fast pointer
		sP = sP.next
		fP = fP.next.next

		if sP == fP {

			//search for the entrance
			p1, p2 := head, fP
			for p1 != p2 {
				p1 = p1.next
				p2 = p2.next
			}
			return p2
		}
	}
	return nil
}