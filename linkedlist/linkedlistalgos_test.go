package linkedlist

import (
	"testing"
)

func TestRemoveDuplicatesWithBuffer(t *testing.T) {
	RemoveDuplicatesWithBuffer(nil)
	var ll LinkedList
	ll.AddFront(1)
	ll.AddFront(2)
	ll.AddFront(3)
	ll.AddFront(2)
	RemoveDuplicatesWithBuffer(&ll)
	if ll.Size() != 3 {
		t.Errorf("Expected size to be 3 but was %d", ll.Size())
	}
	var ll2 LinkedList
	RemoveDuplicatesWithBuffer(&ll2)
	if ll2.Size() != 0 {
		t.Errorf("Empty linked list should be 0 size always but was %d", ll2.Size())
	}
	var ll3 LinkedList
	ll3.AddFront(1)
	ll3.AddFront(1)
	ll3.AddFront(1)
	ll3.AddFront(1)
	ll3.AddFront(1)
	RemoveDuplicatesWithBuffer(&ll3)
	if ll3.Size() != 1 {
		t.Errorf("Expected size to be 1 but was %d", ll3.Size())
	}
}

func TestKthToLast(t *testing.T) {
	if x := KthToLast(nil, 0); x != nil {
		t.Errorf("Empty linkedlist should return nil but got %v", x)
	}
	var ll LinkedList
	ll.AddBack(1)
	ll.AddBack(2)
	ll.AddBack(3)
	ll.AddBack(4)
	ll.AddBack(5)
	if x := KthToLast(&ll, 0); x != 5 {
		t.Errorf("Expected x to be 5, the 0th to last value but got %v", x)
	}
	if x := KthToLast(&ll, 1); x != 4 {
		t.Errorf("Expected x to be 4, the 1st to last value but got %v", x)
	}
	if x := KthToLast(&ll, 2); x != 3 {
		t.Errorf("Expected x to be 3, the 2nd to last value but got %v", x)
	}
	if x := KthToLast(&ll, 3); x != 2 {
		t.Errorf("Expected x to be 2, the 3rd to last value but got %v", x)
	}
	if x := KthToLast(&ll, 4); x != 1 {
		t.Errorf("Expected x to be 1, the 4th to last value but got %v", x)
	}
	if x := KthToLast(&ll, 5); x != nil {
		t.Errorf("Expected x to be nil as k is as long as the list but got %v", x)
	}
}
