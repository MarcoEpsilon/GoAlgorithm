package linked

import (
	"errors"
	"reflect"
)

var (
	QueueIsEmpty = errors.New("QueueIsEmpty")
	UnExpectedType = errors.New("UnExpectedType")
	TypeError = errors.New("TypeError")
)

type Node = *node_inner
type node_inner struct {
	data interface{}
	next Node
}
type Queue = *queue_inner
type queue_inner struct {
	front Node
	rear Node
}

func New() (Queue) {
	return &queue_inner {
		front: nil,
		rear: nil,
	}
}

func (queue Queue) IsEmpty() bool {
	return queue.front == nil
}

func (queue Queue) Length() int {
	i := 0
	for current := queue.front; current != nil; current = current.next {
		i++
	}
	return i
}

func (queue Queue) checkType(data interface{}) (err error) {
	if queue.IsEmpty() {
		return nil
	}
	leftType := reflect.TypeOf(queue.front.data)
	rightType := reflect.TypeOf(data)
	if leftType == rightType {
		return nil
	}
	return UnExpectedType
}

func (queue Queue) EnQueue(data interface{}) (err error) {
	err = queue.checkType(data)
	if err != nil {
		return err
	}
	elem := &node_inner {
		data: data,
		next: nil,
	}
	if queue.rear == nil {
		queue.front = elem
		queue.rear = elem
		return nil
	}
	queue.rear.next = elem
	queue.rear = elem 
	return nil
}

func (queue Queue) DeQueue() (data interface{}, err error) {
	if queue.IsEmpty() {
		return nil, QueueIsEmpty
	}
	data = queue.front.data
	queue.front = queue.front.next
	if queue.front == nil {
		queue.rear = nil
	}
	return data, nil
}

func NewWith(elements interface{}) (queue Queue, err error) {
	value := reflect.ValueOf(elements)
	queue = New()
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			err = queue.EnQueue(value.Index(i).Interface())
			if err != nil {
				return nil, err
			}
		}
		return queue, nil
	default:
		return nil, TypeError
	}
}

