package stack

import (
	"testing"
)

func TestStackQueueBasicFunctionality(t *testing.T) {
	s := NewStackQueue()
	if s.IsEmpty() != true {
		t.Error("Was expecting queue to be empty")
	}
	s.Add(1)
	if s.IsEmpty() != false {
		t.Error("Was expecting queue to be non-empty")
	}
	s.Add(2)
	if s.IsEmpty() != false {
		t.Error("Was expecting queue to be non-empty")
	}
	s.Add(3)
	if s.IsEmpty() != false {
		t.Error("Was expecting queue to be non-empty")
	}
	if peek, _ := s.Peek(); peek != 1 {
		t.Errorf("Was expecting first element peeked to be 1, not %v", peek)
	}
	if remove, _ := s.Remove(); remove != 1 {
		t.Errorf("Was expecting first element removed to be 1, not %v", remove)
	}
	if remove, _ := s.Remove(); remove != 2 {
		t.Errorf("Was expecting first element removed to be 2, not %v", remove)
	}
	if remove, _ := s.Remove(); remove != 3 {
		t.Errorf("Was expecting first element removed to be 3, not %v", remove)
	}
	if s.IsEmpty() != true {
		t.Error("Was expecting queue to be empty")
	}
}
