package singlelist

import "testing"
// import "fmt"
func checkError(err error,t *testing.T) {
	switch err {
	case nil:
		return
	case UnExpectedType,OperationWithEmpty,TypeError,
	OutOfRangeIndex,NotSupportCompare:
		t.Error(err.Error())
	default:
		t.Error("haven't handle the error")
	}
}

func getSlice(link LinkList) []interface{} {
	var slice []interface{}
	head := link.head
	for ; head != nil; {
		slice = append(slice,head.data)
		head = head.next
	}
	return slice
}

func TestAppend(t *testing.T) {
	errMsg := "fucn Append's result is (%d = %d) unexpected"
	list := New()
	src := []int{1,2,3,4,5}
	for i := 0; i < len(src); i++ {
		err := list.Append(src[i])
		checkError(err,t)
	}
	result := getSlice(list)
	if len(result) != len(src) {
		t.Error(errMsg,len(result),len(src))
	}
	for i := 0; i < len(src); i++ {
		if result[i].(int) != src[i] {
			t.Errorf(errMsg,result[i],src[i])
		}
	}
}


func TestNewWith(t *testing.T) {
	errMsg := "func NewWith's result (%d = %d) is unexpected"
	list,err := NewWith([]int{1,2,3,4,5})
	checkError(err,t)
	trans := getSlice(list)
	result := []int{1,2,3,4,5}
	if len(trans) != len(result) {
		t.Errorf(errMsg,len(trans),len(result))
	}
	for i := 0; i < len(trans); i++ {
		if trans[i].(int) != result[i] {
			t.Errorf(errMsg,trans[i],result[i])
		}
	}
}


func TestLength(t *testing.T) {
	errMsg := "func Length's result (%d = %d) is unexpected"
	list,err := NewWith([]int{1,2,3,4,5})
	checkError(err,t)
	if list.Length() != 5 {
		t.Errorf(errMsg,list.Length,5)
	}
}
func TestIsEmpty(t *testing.T) {
	list := New()
	errMsg := "func IsEmpty's result is unexpected"
	if !list.IsEmpty() {
		t.Error(errMsg)
	}
}

func TestInsertByIndex(t *testing.T) {
	errMsg := "func InsertByIndex's result (%d = %d) is unexpected"
	list,err := NewWith([]int{1,2,4,5,6})
	checkError(err,t)
	err = list.InsertByIndex(0,0)
	checkError(err,t)
	err = list.InsertByIndex(3,3)
	checkError(err,t)
	trans := getSlice(list)
	result := []int{0,1,2,3,4,5,6}
	if len(trans) != len(result) {
		t.Errorf(errMsg,len(trans),len(result))
	}
	for i := 0; i < len(trans); i++ {
		if trans[i].(int) != result[i] {
			t.Errorf(errMsg,trans[i],result[i])
		}
	}
}

func TestDeleteByIndex(t *testing.T) {
	errMsg := "func DeleteByIndex's result (%d = %d) is unexpected"
	list,err := NewWith([]int{0,1,2,3,4,5})
	checkError(err,t)
	elem,err := list.DeleteByIndex(0)
	checkError(err,t)
	if elem.(int) != 0 {
		t.Errorf(errMsg,elem,0)
	}
	elem,err = list.DeleteByIndex(4)
	checkError(err,t)
	if elem.(int) != 5 {
		t.Errorf(errMsg,elem,5)
	}
}


func TestGet(t *testing.T) {
	errMsg := "func Get's result (%d = %d) is unexpected"
	list,err := NewWith([]int{0,1,2,3,4,5})
	checkError(err,t)
	result := []int{0,1,2,3,4,5}
	for i := 0; i < len(result); i++ {
		elem,err := list.Get(int64(i))
		checkError(err,t)
		if elem.(int) != result[i] {
			t.Errorf(errMsg,elem,result[i])
		}
	}
}

func TestModify(t *testing.T) {
	errMsg := "func Modify's result (%d = %d) is unexpected"
	list,err := NewWith([]int{1,2,3,4,5,6})
	checkError(err,t)
	modify := []int{6,5,4,3,2,1}
	for i := 0; i < len(modify); i++ {
		err = list.Modify(int64(i),modify[i])
		checkError(err,t)
	}
	trans := getSlice(list)
	for i := 0; i < len(modify); i++ {
		if trans[i].(int) != modify[i] {
			t.Errorf(errMsg,trans[i],modify[i])
		}
	}
}

func TestDeleteAllWithRecusive(t *testing.T) {
	list,err := NewWith([]int{2,3,2,7,5,3,1,2})
	checkError(err,t)
	err = list.DeleteAllWithRecusive(2)
	checkError(err,t)
	result := []int{3,7,5,3,1}
	trans := getSlice(list)
	errMsg := "func DeleteAllWithRecusive's result (%d = %d) is unexpected"
	if len(trans) != len(result) {
		t.Errorf(errMsg,len(trans),len(result))
	}
	for i := 0; i < len(result); i++ {
		if trans[i].(int) != result[i] {
			t.Errorf(errMsg,trans[i],result[i])
		}
	}
	err = list.DeleteAllWithRecusive(3)
	checkError(err,t)
	trans = getSlice(list)
	result = []int{7,5,1}
	if len(trans) != len(result) {
		t.Errorf(errMsg,len(trans),len(result))
	}
	for i := 0; i < len(trans); i++ {
		if trans[i].(int) != result[i] {
			t.Errorf(errMsg,trans[i],result[i])
		}
	}
}

func TestDeleteMin(t *testing.T) {
	list,err := NewWith([]int{1,4,3,5,2})
	checkError(err,t)
	minElem,err := list.DeleteMin()
	checkError(err,t)
	errMsg := "func DeleteMin's result (%d = %d) is unexpected"
	if minElem.(int) != 1 {
		t.Errorf(errMsg,minElem,1)
	}
	minElem,err = list.DeleteMin()
	checkError(err,t)
	if minElem.(int) != 2 {
		t.Errorf(errMsg,minElem,2)
	}
	minElem,err = list.DeleteMin()
	checkError(err,t)
	if minElem.(int) != 3 {
		t.Errorf(errMsg,minElem,3)
	}
}

func TestReverse(t *testing.T) {
	list,err := NewWith([]int{1,2,3,4,5})
	checkError(err,t)
	err = list.Reverse()
	checkError(err,t)
	result := []int{5,4,3,2,1}
	errMsg := "func Reverse's result (%d = %d) is unexpected"
	trans := getSlice(list)
	if len(trans) != len(result) {
		t.Errorf(errMsg,len(trans),len(result))
	}
	for i := 0; i < len(trans); i++ {
		if trans[i].(int) != result[i] {
			t.Errorf(errMsg,trans[i],result[i])
		}
	}
}

func TestReverseBySwap(t *testing.T) {
	list,err := NewWith([]int{1,2,3,4,5,6})
	checkError(err,t)
	err = list.ReverseBySwap()
	checkError(err,t)
	result := []int{6,5,4,3,2,1}
	trans := getSlice(list)
	errMsg := "func NewWith's result (%d = %d) is uexpected"
	if len(trans) != len(result) {
		t.Errorf(errMsg,len(result),len(trans))
	}
	for i := 0; i < len(trans); i++ {
		if trans[i].(int) != result[i] {
			t.Errorf(errMsg,trans[i],result[i])
		}
	}
}





func TestInsertSort(t *testing.T) {
	list,err := NewWith([]int{9,4,1,0,2,3,7,5,8,6})
	checkError(err,t)
	err = list.InsertSort()
	checkError(err,t)
	errMsg := "func Sort's result (%d = %d) is unexpected"
	result := []int{0,1,2,3,4,5,6,7,8,9}
	trans := getSlice(list)
	if len(trans) != len(result) {
		t.Errorf(errMsg,len(trans),len(result))
	}
	for i := 0; i < len(trans); i++ {
		if trans[i].(int) != result[i] {
			t.Errorf(errMsg,trans[i],result[i])
		}
	}
}

func TestDeleteRangeElem(t *testing.T) {
	list,err := NewWith([]int{1,2,3,4,5,6})
	checkError(err,t)
	err = list.DeleteRangeElem(2,5)
	checkError(err,t)
	result := []int{1,6}
	trans := getSlice(list)
	errMsg := "func DeleteRangeElem's result (%d = %d) is unexpected"
	if len(trans) != len(result) {
		t.Errorf(errMsg,len(result),len(trans))
	}
	for i := 0; i < len(trans); i++ {
		if trans[i].(int) != result[i] {
			t.Errorf(errMsg,trans[i],result[i])
		}
	}
	err = list.DeleteRangeElem(1,7)
	checkError(err,t)
	if !list.IsEmpty() {
		t.Error("func DeleteRangeElem's result is unexpected")
	}
}
