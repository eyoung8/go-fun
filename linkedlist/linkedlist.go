package linkedlist

import (
	"errors"
)

var (
	outOfBoundsErr = errors.New("Out of bounds")
	emptyListErr   = errors.New("Empty list")
)

type LinkedList struct {
	head *node
	tail *node
	size int
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) AddFront(val interface{}) {
	n := node{
		val:  val,
		next: l.head,
	}

	if l.head == nil {
		l.addFrontBack(&n)
		return
	}
	l.head.prev = &n
	l.head = &n
	l.size++
}

func (l *LinkedList) AddBack(val interface{}) {
	n := node{
		val:  val,
		prev: l.tail,
	}

	if l.tail == nil {
		l.addFrontBack(&n)
		return
	}

	l.tail.next = &n
	l.tail = &n
	l.size++
}

func (l *LinkedList) addFrontBack(n *node) {
	l.head = n
	l.tail = n
	l.size++
}

func (l *LinkedList) Insert(val interface{}, pos int) error {
	newNode := node{
		val: val,
	}

	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		return nil
	}

	if l.size < pos {
		return outOfBoundsErr
	}

	n := l.head
	for i := 0; i != pos; i++ {
		n = n.next
	}

	newNode.next = n.next
	newNode.prev = n
	n.next = &newNode
	l.size++
	return nil
}

func (l *LinkedList) Pop() interface{} {
	if l.size < 1 {
		return nil
	}
	val := l.tail.val
	newTail := l.tail.prev
	if newTail == nil {
		l.head = nil
	} else {
		newTail.next = nil
	}
	l.tail = newTail
	l.size--
	return val
}

func (l *LinkedList) Front() interface{} {
	if l.head != nil {
		return l.head.val
	}
	return nil
}

func (l *LinkedList) Back() interface{} {
	if l.tail != nil {
		return l.tail.val
	}
	return nil
}

func (l *LinkedList) Get(pos int) (interface{}, error) {
	if pos > l.size {
		return nil, outOfBoundsErr
	}
	n := l.head
	for i := 0; i != pos; i++ {
		n = n.next
	}
	return n.val, nil
}

func (l *LinkedList) Remove(pos int) (interface{}, error) {
	if pos > l.size {
		return nil, outOfBoundsErr
	}
	n := l.head
	for i := 0; i != pos; i++ {
		n = n.next
	}

	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}

	l.size--
	return n.val, nil
}

func (l *LinkedList) RemoveFront() (interface{}, error) {
	if l.head != nil {
		if l.head.next != nil {
			l.head.next.prev = nil
		}
		val := l.head.val
		l.head = l.head.next
		l.size--
		return val, nil
	}
	return nil, emptyListErr
}

func (l *LinkedList) RemoveBack() (interface{}, error) {
	if l.tail != nil {
		if l.tail.prev != nil {
			l.tail.prev.next = nil
		}
		val := l.tail.val
		l.tail = l.tail.prev
		l.size--
		return val, nil
	}
	return nil, emptyListErr
}

func (l *LinkedList) deleteNode(n *node) error {
	if n == nil || l == nil {
		return errors.New("Nil pointer")
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	if l.head == n {
		l.head = n.next
	}
	if l.tail == n {
		l.tail = n.prev
	}
	l.size--
	return nil
}
