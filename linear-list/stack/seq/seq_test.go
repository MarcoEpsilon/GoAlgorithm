package seq

import (
	"testing"
)
func checkError(err error, t *testing.T) {
	switch err {
	case nil:
		return
	default:
		t.Error(err)
	}
}
func TestIsResult(t *testing.T) {
	errMsg := "func TestIsResult's result is unexpected"
	is, err := IsResult([]int{1,2,3,4,5},[]int{3,2,5,4,1})
	checkError(err, t)
	if !is {
		t.Error(errMsg)
	}
}


func TestSortWithBubble(t *testing.T) {
	errMsg := "func TestSortWithBubble's result (%d = %d) is unexpected"
	stack, err := NewWith([]int{0,3,6,2,1,7,4,8,9,5})
	checkError(err, t)
	result := []int{0,1,2,3,4,5,6,7,8,9}
	err = stack.SortWithBubble()
	checkError(err, t)
	if stack.Length() != len(result) {
		t.Errorf(errMsg, stack.Length(), len(result))
		panic("SortWithBubble react")
	}
	i := 0
	for ; !stack.IsEmpty(); {
		top, _ := stack.Pop()
		if top.(int) != result[i] {
			t.Errorf(errMsg, top, result[i])
		}
		i++	
	}
}