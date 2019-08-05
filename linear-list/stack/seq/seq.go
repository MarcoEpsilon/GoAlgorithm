package seq

import (
	"errors"
	"reflect"
)

var (
	StackIsEmpty 	= 	errors.New("StackIsEmpty")
	StackIsFull		=   errors.New("StackIsFull")
	TypeError		=   errors.New("TypeError")
	UnExceptedType	=	errors.New("UnExceptedType")
)

const (
	stackMaxSize	=	2 << 10
)
type Stack = *stack_inner
type stack_inner struct {
	elems	[stackMaxSize]interface{}
	top		int
}

func New() (Stack) {
	return &stack_inner {
		elems: [stackMaxSize]interface{}{},
		top: -1,
	}
}

func NewWith(elems interface{}) (Stack,error) {
	value := reflect.ValueOf(elems)
	stack := New()
	switch reflect.TypeOf(elems).Kind() {
	case reflect.Slice,reflect.Array:
		for i := 0; i < value.Len(); i++ {
			stack.top++
			stack.elems[i] = value.Index(i).Interface()
		}
	default:
		return nil,UnExceptedType
	}
	return stack,nil
}

func (stack Stack) IsEmpty() (bool) {
	return stack.top == -1
}

func (stack Stack) IsFull() (bool) {
	return stack.top == stackMaxSize - 1
}

func (stack Stack) Length() (int) {
	return stack.top + 1
}

func (stack Stack) Pop() (elem interface{},err error) {
	if stack.IsEmpty() {
		return nil,StackIsEmpty
	}
	elem = stack.elems[stack.top]
	stack.top--
	return elem,nil
}

func (stack Stack) checkType(elem interface{}) (err error) {
	if stack.IsEmpty() {
		return nil
	}
	leftKind := reflect.TypeOf(stack.elems[stack.top]).Kind()
	rightKind := reflect.TypeOf(elem).Kind()
	if leftKind == rightKind {
		return nil
	} else {
		return UnExceptedType
	}
}

func (stack Stack) Push(elem interface{}) (err error) {
	if stack.IsFull() {
		return StackIsFull
	}
	err = stack.checkType(elem)
	if err != nil {
		return err
	}
	stack.top += 1
	stack.elems[stack.top] = elem
	return nil
}

func (stack Stack) Top() (elem interface{},err error) {
	if stack.IsEmpty() {
		return nil, StackIsEmpty
	}
	return stack.elems[stack.top], nil	
}

func (stack Stack) moveBottomToTop() {
	if stack.IsEmpty() {
		return
	}
	top, _ := stack.Pop()
	if !stack.IsEmpty() {
		bottom, _ = stack.Pop()
		stack.moveBottomToTop()
		_ = stack.Push(bottom)
		_ = stack.Push(top)
	} else {
		stack.Push(top)
	}
}
func (stack Stack) reverse() {
	
}

func (stack Stack) Reverse() {
	
}