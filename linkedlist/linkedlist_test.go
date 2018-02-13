package linkedlist

import (
	"testing"
)

func TestAddFront(t *testing.T) {
	var ll LinkedList
	i := 1
	ll.AddFront(i)
	if i != ll.Back() {
		t.Error("Single node linkedlist back does not match front when front inserted")
	}
	if i != ll.Front() {
		t.Error("Single node linkedlist failed to insert in front")
	}
	if ll.Size() != 1 {
		t.Errorf("Inserting one val should result in size=1 but instead resulted in %d", ll.Size())
	}
	j := 2
	ll.AddFront(j)
	if j != ll.Front() {
		t.Errorf("Front was not expected with val=%d expected=%d", ll.Front(), j)
	}
	if i != ll.Back() {
		t.Errorf("Back changed to %d after inserting to front", ll.Back())
	}
	if ll.Size() != 2 {
		t.Errorf("Expected size=2 after inserting 2 values but received: %d", ll.Size())
	}
}

func TestAddBack(t *testing.T) {
	var ll LinkedList
	i := 1
	ll.AddBack(i)
	if i != ll.Front() {
		t.Error("Single node linkedlist back does not match front when back inserted")
	}
	if i != ll.Back() {
		t.Error("Single node linkedlist failed to insert in back")
	}
	if ll.Size() != 1 {
		t.Errorf("Inserting one val should result in size=1 but instead resulted in %d", ll.Size())
	}
	j := 2
	ll.AddBack(j)
	if i != ll.Front() {
		t.Error("Front changed after inserting to back")
	}
	if j != ll.Back() {
		t.Errorf("Back did not match value inserted in back but was instead %d", ll.Back())
	}
	if ll.Size() != 2 {
		t.Errorf("Expected size=2 after inserting 2 values but received: %d", ll.Size())
	}
}

func TestFrontBackEmptyList(t *testing.T) {
	var ll LinkedList
	if ll.Front() != nil {
		t.Errorf("Front of empty list somehow returned non-nail value=%v", ll.Front())
	}
	if ll.Back() != nil {
		t.Errorf("Back of empty list somehow returned non-nail value=%v", ll.Front())
	}
}

func TestSize(t *testing.T) {
	var ll LinkedList
	if ll.Size() != 0 {
		t.Errorf("Size was expected to be 0 but was %d", ll.Size())
	}
	ll.AddFront(1)
	ll.AddFront(1)
	ll.AddFront(1)
	if ll.Size() != 3 {
		t.Errorf("Size was expected to be 3 but was %d", ll.Size())
	}
}

func TestPop(t *testing.T) {
	var ll LinkedList
	i := 1
	j := 2
	ll.AddFront(i)
	ll.AddFront(j)
	size := ll.Size()
	val := ll.Pop()
	if val != i {
		t.Errorf("Expected popped back value to be %d but was %v", i, val)
	}
	if ll.Size() != size-1 {
		t.Errorf("Expected size to be 1 less than before pop=%d but was %d", size-1, ll.Size())
	}
}

func TestPopEmptyList(t *testing.T) {
	var ll LinkedList
	val := ll.Pop()
	if val != nil {
		t.Errorf("Empty linked list somehow popped non-nil val=%v", val)
	}
}
