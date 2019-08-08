package seq

import (
	"errors"
	"reflect"
	// "fmt"
)

var (
	StackIsEmpty   = errors.New("StackIsEmpty")
	StackIsFull    = errors.New("StackIsFull")
	TypeError      = errors.New("TypeError")
	UnExpectedType = errors.New("UnExceptedType")
	OperationInvalid = errors.New("OperationInvalid")
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

const (
	stackMaxSize = 2 << 10
)

type Stack = *stack_inner
type stack_inner struct {
	elems [stackMaxSize]interface{}
	top   int
}

func New() Stack {
	return &stack_inner{
		elems: [stackMaxSize]interface{}{},
		top:   -1,
	}
}

func NewWith(elems interface{}) (Stack, error) {
	value := reflect.ValueOf(elems)
	stack := New()
	switch reflect.TypeOf(elems).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			stack.top++
			stack.elems[i] = value.Index(i).Interface()
		}
	default:
		return nil, UnExpectedType
	}
	return stack, nil
}

func (stack Stack) IsEmpty() bool {
	return stack.top == -1
}

func (stack Stack) IsFull() bool {
	return stack.top == stackMaxSize-1
}

func (stack Stack) Length() int {
	return stack.top + 1
}

func (stack Stack) Pop() (elem interface{}, err error) {
	if stack.IsEmpty() {
		return nil, StackIsEmpty
	}
	elem = stack.elems[stack.top]
	stack.top--
	return elem, nil
}

func (stack Stack) checkType(elem interface{}) (err error) {
	if stack.IsEmpty() {
		return nil
	}
	leftKind := reflect.TypeOf(stack.elems[stack.top])
	rightKind := reflect.TypeOf(elem)
	if leftKind == rightKind {
		return nil
	} else {
		return UnExpectedType
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

func (stack Stack) Top() (elem interface{}, err error) {
	if stack.IsEmpty() {
		return nil, StackIsEmpty
	}
	return stack.elems[stack.top], nil
}

func (stack Stack) moveBottomToTop() {
	if stack.IsEmpty() {
		return
	}
	// get the current function's stack element
	top, _ := stack.Pop()
	if !stack.IsEmpty() {
		//deeply
		stack.moveBottomToTop()
		//get the stack bottom element
		bottom, _ := stack.Pop()
		// reverse bottom and top
		_ = stack.Push(top)
		_ = stack.Push(bottom)
	} else {
		// return the last element of stack
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
func Copy(stack Stack) (Stack) {
	return &stack_inner {
		elems: stack.elems,
		top: stack.top,
	}
}

func IsResult(push []int,pop []int) (is bool, err error) {
	help := New()
	if len(push) != len(pop) || len(push) == 0 {
		return false, OperationInvalid
	}
	j := 0
	for i := 0; i < len(pop); i++ {
		if !help.IsEmpty() {
			top, err := help.Top()
			if err != nil {
				return false, err
			}
			status, err := compare(top,pop[i])
			if err != nil {
				return false,err	
			}
			if status == Eq {
				_, err = help.Pop()
				if err != nil {
					return false, err
				}
				continue
			}
		}
		for ; j < len(push); j++ {
			status,err := compare(pop[i],push[j])
			if err != nil {
				return false, err
			}
			if status != Eq {
				_ = help.Push(push[j])
			} else {
				j = j + 1
				break
			}
		}
	}
	if help.IsEmpty() {
		return true, nil
	} else {
		return false, nil
	}
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

func (stack Stack) bottomBubble() (err error) {
	if stack.IsEmpty() {
		return
	}
	top, _ := stack.Pop()
	if !stack.IsEmpty() {
		waitErr := stack.bottomBubble()
		bottom, _ := stack.Pop()
		status, err := compare(top,bottom)
		if err != nil {
			return err
		}
		if status == LessThan {
			stack.Push(bottom)
			stack.Push(top)
		} else {
			stack.Push(top)
			stack.Push(bottom)
		}
		return waitErr
	} else {
		stack.Push(top)
		return nil
	}
}