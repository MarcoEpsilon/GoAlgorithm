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
	NotSupportCompare = errors.New("NotSupportCompare")
)

const (
	LessThan      = -1
	Eq            = 0
	GreaterThan   = 1
	UnCompareable = 2
)

func compare(left interface{}, right interface{}) (status int, err error) {
	if reflect.TypeOf(left).Kind() != reflect.TypeOf(right).Kind() {
		return UnCompareable, UnExpectedType
	}
	//ps:switch 的局限性导致大量重复代码
	switch right := right.(type) {
	case int:
		left := left.(int)
		if left < right {
			return LessThan, err
		} else if left > right {
			return GreaterThan, err
		} else {
			return Eq, err
		}
	case string:
		left := left.(string)
		if left < right {
			return LessThan, err
		} else if left > right {
			return GreaterThan, err
		} else {
			return Eq, err
		}
	case float32:
		left := left.(float32)
		if left < right {
			return LessThan, err
		} else if left > right {
			return GreaterThan, err
		} else {
			return Eq, err
		}
	//ignore int32,int64,uint32,float64 throw to default
	default:
		return UnCompareable, NotSupportCompare
	}
}

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

func (stack Stack) checkType(elem interface{}) (err error) {
	if stack.IsEmpty() {
		return nil
	}
	leftType := reflect.TypeOf(stack.top.data)
	rightType := reflect.TypeOf(elem)
	if leftType == rightType {
		return nil
	}
	return TypeError
}

func (stack Stack) Push(data interface{}) (err error) {
	err = stack.checkType(data)
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
			err := stack.Push(value.Index(i).Interface())
			if err != nil {
				return nil, err
			}
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

func checkSliceOrArray(push interface{},pop interface{}) (err error) {
	pushKind := reflect.TypeOf(push).Kind()
	popKind := reflect.TypeOf(pop).Kind()
	if pushKind == popKind && (pushKind == reflect.Slice || pushKind == reflect.Array) {
		return nil
	}
	return UnExpectedType
}

func IsResult(push interface{},pop interface{}) (bool,error) {
	help := New()
	j := 0
	popValue := reflect.ValueOf(pop)
	pushValue := reflect.ValueOf(push)
	for i := 0; i < popValue.Len(); i++ {
		if !help.IsEmpty() {
			top, _ := help.Top()
			status, err := compare(top,popValue.Index(i).Interface())
			if err != nil {
				return false, err
			}
			if status == Eq {
				help.Pop()
				continue
			}
		}
		for ; j < pushValue.Len(); j++ {
			leftValue := popValue.Index(i).Interface()
			rightValue := pushValue.Index(j).Interface()
			status, err := compare(leftValue,rightValue)
			if err != nil {
				return false, err
			}
			if status == Eq {
				j++
				break
			} else {
				_ = help.Push(rightValue)
			}
		}
	}
	if !help.IsEmpty() {
		return false, nil
	} else {
		return true, nil
	}
}

func (stack Stack) bottomBubble() (err error) {
	if stack.IsEmpty() {
		return nil
	}
	top, _ := stack.Pop()
	if !stack.IsEmpty() {
		waitErr := stack.bottomBubble()
		bottom, _ := stack.Pop()
		status, err := compare(bottom, top)
		if err != nil {
			return err
		}
		if status == LessThan {
			stack.Push(top)
			stack.Push(bottom)
		} else {
			stack.Push(bottom)
			stack.Push(top)
		}
		return waitErr
	} else {
		stack.Push(top)
	}
	return nil
}

func (stack Stack) SortWithBubble() (err error) {
	if stack.IsEmpty() {
		return nil
	}
	err = stack.bottomBubble()
	if err != nil {
		return err
	}
	top, _ := stack.Pop()
	err = stack.SortWithBubble()
	if err != nil {
		return err
	}
	stack.Push(top)
	return nil
}