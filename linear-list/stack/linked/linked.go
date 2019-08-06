package linked

import (
	"errors"
	"reflect"
)

// error handle
var (
	StackIsEmpty   = errors.New("StackIsEmpty")
	UnExpectedType = errors.New("UnExpectedType")
	TypeError      = errors.New("TypeError")
)

type Node = *node
type node struct {
	data interface{}
	next Node
}
type Stack = *stack_inner
type stack_inner struct {
	top Node
}

func New() Stack {
	return &stack_inner{
		top: nil,
	}
}

func (stack Stack) IsEmpty() bool {
	return stack.top == nil
}

func (stack Stack) Length() int {
	i := 0
	for current := stack.top; current != nil; current = current.next {
		i++
	}
	return i
}

func (stack Stack) Pop() (elem interface{}, err error) {
	if stack.IsEmpty() {
		return nil, StackIsEmpty
	}
	elem = stack.top.data
	stack.top = stack.top.next
	return elem, nil
}

func (stack Stack) checkeType(elem interface{}) (err error) {
	if stack.IsEmpty() {
		return StackIsEmpty
	}
	leftType := reflect.TypeOf(stack.top.data)
	rightType := reflect.TypeOf(elem)
	if leftType == rightType {
		return nil
	}
	return TypeError
}

func (stack Stack) Push(data interface{}) (err error) {
	err = stack.checkeType(data)
	if err != nil {
		return err
	}
	elem := &node{
		data: data,
		next: stack.top,
	}
	stack.top = elem
	return nil
}
func NewWith(elems interface{}) (Stack, error) {
	value := reflect.ValueOf(elems)
	stack := New()
	switch reflect.TypeOf(elems).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			stack.Push(value.Index(i).Interface())
		}
	default:
		return nil, UnExpectedType
	}
	return stack, nil
}

func (stack Stack) Top() (elem interface{}, err error) {
	if stack.IsEmpty() {
		return nil, StackIsEmpty
	}
	return stack.top.data, nil
}

func (stack Stack) moveBottomToTop() {
	if stack.IsEmpty() {
		return
	}
	top, _ := stack.Pop()
	if !stack.IsEmpty() {
		stack.moveBottomToTop()
		bottom, _ := stack.Pop()
		_ = stack.Push(top)
		_ = stack.Push(bottom)
	} else {
		stack.Push(top)
	}
}

func (stack Stack) reverse() {
	if stack.IsEmpty() {
		return
	}
	stack.moveBottomToTop()
	top, _ := stack.Pop()
	stack.reverse()
	_ = stack.Push(top)
}

func (stack Stack) Reverse() {
	stack.reverse()
}
