package linkedlist

func RemoveDuplicatesWithBuffer(ll *LinkedList) {
	if ll == nil {
		return
	}
	vals := make(map[interface{}]bool)
	n := ll.head
	for n != nil {
		_, ok := vals[n.val]
		vals[n.val] = true
		next := n.next
		if ok {
			ll.deleteNode(n)
		}
		n = next
	}
}

func KthToLast(ll *LinkedList, k int) interface{} {
	if ll == nil {
		return nil
	}
	i := 0
	n := ll.head
	for n != nil {
		i++
		n = n.next
	}
	pos := i - k - 1
	if pos < 0 {
		return nil
	}
	val, _ := ll.Get(pos)
	return val
}
