package stack

type StackQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func (q *StackQueue) Add(val interface{}) {
	q.stack2.Push(val)
	for v, _ := q.stack1.Pop(); v != nil; v, _ = q.stack1.Pop() {
		q.stack2.Push(v)
	}
	q.stack1 = q.stack2
	q.stack2 = &Stack{}
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
