package stack

type StackQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func NewStackQueue() StackQueue {
	s := StackQueue{}
	s.stack1 = &Stack{}
	s.stack2 = &Stack{}
	return s
}

func (q *StackQueue) Add(val interface{}) {
	for v, _ := q.stack1.Pop(); v != nil; v, _ = q.stack1.Pop() {
		q.stack2.Push(v)
	}
	q.stack1.Push(val)
	for v, _ := q.stack2.Pop(); v != nil; v, _ = q.stack2.Pop() {
		q.stack1.Push(v)
	}
}

func (q *StackQueue) Remove() (interface{}, error) {
	return q.stack1.Pop()
}

func (q *StackQueue) Peek() (interface{}, error) {
	return q.stack1.Peek()
}

func (q *StackQueue) IsEmpty() bool {
	return q.stack1.IsEmpty()
}
