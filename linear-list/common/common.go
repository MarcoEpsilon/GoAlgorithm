package common

import (
	stack "algorithm/linear-list/stack/linked"
	"errors"
	"reflect"
)
var (
	QueueIsEmpty = errors.New("QueueIsEmpty")
	TypeError = errors.New("TypeError")
	UnExpectedType = errors.New("UnExpectedType")
)
type Queue = *queue_inner
// use two stack achieve queue
type queue_inner struct {
	ened stack.Stack
	deing stack.Stack
}

func New() Queue {
	return &queue_inner {
		ened: stack.New(),
		deing: stack.New(),
	}
}

func NewWith(elements interface{}) (Queue, error) {
	ened, err := stack.NewWith(elements)
	return &queue_inner {
		ened: ened,
		deing: stack.New(),
	}, err
}

func (queue Queue) IsEmpty() bool {
	return queue.ened.IsEmpty() && queue.deing.IsEmpty()
}

func (queue Queue) Length() (int) {
	return queue.ened.Length() + queue.deing.Length() 
}

func (queue Queue) checkType(data interface{}) (err error) {
	if queue.IsEmpty() {
		return nil
	}
	leftType := reflect.TypeOf(data)
	if !queue.ened.IsEmpty() {
		top, _ := queue.ened.Top()
		rightType := reflect.TypeOf(top)
		if leftType != rightType {
			return TypeError
		}
	} else if !queue.deing.IsEmpty() {
		top, _ := queue.deing.Top()
		rightType := reflect.TypeOf(top)
		if leftType != rightType {
			return TypeError
		}
	}
	return nil
}

func (queue Queue) EnQueue(data interface{}) (err error) {
	err = queue.checkType(data)
	if err != nil {
		return err
	}
	return queue.ened.Push(data)
}

func (queue Queue) DeQueue() (element interface{}, err error) {
	if queue.IsEmpty() {
		return nil, QueueIsEmpty
	}
	if !queue.deing.IsEmpty() {
		return queue.deing.Pop()
	}
	for ; !queue.ened.IsEmpty(); {
		top, _ := queue.ened.Pop()
		queue.deing.Push(top)
	}
	return queue.deing.Pop()
}