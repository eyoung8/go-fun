package queue

import (
	"errors"
	ll "github.com/eyoung8/go-fun/linkedlist"
)

var (
	EmptyQueueErr = errors.New("Empty queue")
)

type Queue struct {
	list ll.LinkedList
}

func (q *Queue) Add(val interface{}) {
	q.list.AddFront(val)
}

func (q *Queue) Remove() (interface{}, error) {
	val, err := q.list.RemoveBack()
	if err != nil {
		return nil, EmptyQueueErr
	}
	return val, nil
}

func (q *Queue) Peek() (interface{}, error) {
	val := q.list.Back()
	if val == nil {
		return nil, EmptyQueueErr
	}
	return val, nil
}

func (q *Queue) IsEmpty() bool {
	return q.list.Size() == 0
}
