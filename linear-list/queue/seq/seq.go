// cycle queue
package seq

import (
	"errors"
	"reflect"
)

var (
	QueueIsEmpty   = errors.New("QueueIsEmpty")
	QueueIsFull    = errors.New("QueueIsFull")
	UnExpectedType = errors.New("UnExpectedType")
	TypeError      = errors.New("TypeError")
)

// for test we set as small as possible
const queueSize = 10

type Queue = *queue_innner
type queue_innner struct {
	elems [queueSize]interface{}
	front int
	rear  int
}

func New() Queue {
	return &queue_innner{
		elems: [queueSize]interface{}{},
		front: 0,
		rear:  0,
	}
}

func (queue Queue) IsEmpty() bool {
	if queue.rear == queue.front {
		return true
	}
	return false
}

func (queue Queue) Length() int {
	return (queue.rear - queue.front + queueSize) % queueSize
}

func (queue Queue) IsFull() bool {
	if (queue.rear+1)%queueSize == queue.front {
		return true
	}
	return false
}
func (queue Queue) checkType(elem interface{}) (err error) {
	if queue.IsEmpty() {
		return nil
	}
	randElem := reflect.TypeOf(queue.elems[queue.front])
	rightElem := reflect.TypeOf(elem)
	if randElem == rightElem {
		return nil
	}
	return TypeError
}
func (queue Queue) EnQueue(elem interface{}) (err error) {
	if queue.IsFull() {
		return QueueIsFull
	}
	err = queue.checkType(elem)
	if err != nil {
		return err
	}
	queue.elems[queue.rear] = elem
	queue.rear = (queue.rear + 1) % queueSize
	return nil
}

func NewWith(elements interface{}) (Queue, error) {
	value := reflect.ValueOf(elements)
	queue := New()
	switch reflect.TypeOf(elements).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			err := queue.EnQueue(value.Index(i).Interface())
			if err != nil {
				return nil, err
			}
		}
		return queue, nil
	default:
		return nil, UnExpectedType
	}
}

func (queue Queue) DeQueue() (elem interface{}, err error) {
	if queue.IsEmpty() {
		return nil, QueueIsEmpty
	}
	elem = queue.elems[queue.front]
	queue.front = (queue.front + 1) % queueSize
	return elem, nil
}
