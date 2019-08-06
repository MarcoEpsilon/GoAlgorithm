//seq_test.go
package sequence

import "testing"
import (
// "fmt"
)

func checkError(err error, t *testing.T) {
	switch err {
	case nil:
		return
	case OutOfRangeIndex:
		t.Error("OutOfRangeIndex")
	case TypeError:
		t.Error("TypeError")
	case UnExpectedType:
		t.Error("UnExpectedType")
	case OperationWithEmpty:
		t.Error("OpeationWithEmpty")
	case NotSupportCompare:
		t.Error("NotSupportCompare")
	case IsNotSorted:
		t.Error("IsNotSorted")
	default:
		t.Error("should unreached")
		return
	}
}

func TestNew(t *testing.T) {
	_, err := New(1, 100)
	checkError(err, t)
}
func TestNewWith(t *testing.T) {
	_, err := NewWith([]int{1, 2, 3, 4})
	checkError(err, t)
}

func TestLength(t *testing.T) {
	seq, err := New(1, 100)
	checkError(err, t)
	length := seq.Length()
	if length != 100 {
		t.Error("func Length returned dont' match result")
	}
}

func TestIsEmpty(t *testing.T) {
	seq, err := New(1, 0)
	checkError(err, t)
	if empty := seq.IsEmpty(); !empty {
		t.Error("func IsEmpty returned don't match result")
	}
	seq, err = New("hello", 1)
	checkError(err, t)
	if empty := seq.IsEmpty(); empty {
		t.Error("func isEmpty returned don't match result")
	}
}

func TestGet(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	seq, err := NewWith(slice)
	checkError(err, t)
	for i := 0; i < len(slice); i++ {
		elem, err := seq.Get(int64(i))
		checkError(err, t)
		if slice[i] != elem.(int) {
			t.Error("func Get result is unexpected")
		}
	}
}

func TestModify(t *testing.T) {
	replace := []int{5, 6, 7, 8}
	seq, err := NewWith([]int{1, 2, 3, 4, 5})
	checkError(err, t)
	for i := 0; i < len(replace); i++ {
		err = seq.Modify(int64(i), replace[i])
		checkError(err, t)
	}
	for i := 0; i < len(replace); i++ {
		if seq.data[i].(int) != replace[i] {
			t.Error("func Modify result is unexpected")
		}
	}
}

func TestInsertByIndex(t *testing.T) {
	errMsg := "func InsertByIndex don't match result"
	seq, err := New(10, 10)
	checkError(err, t)
	err = seq.InsertByIndex(6, 5)
	checkError(err, t)
	if seq.data[6] != 5 {
		t.Error(errMsg)
	}
	/*err = seq.InsertByIndex(2,"hello")
	checkError(err,t)
	if seq.data[2] != "hello" {
		t.Error(errMsg)
	}*/
}

func TestDeleteByIndex(t *testing.T) {
	seq, err := NewWith([]int{1, 2, 3, 4, 5})
	checkError(err, t)
	err = seq.DeleteByIndex(2)
	checkError(err, t)
	if seq.data[2] != 4 {
		t.Error("func DeleteByIndex's result unexpected")
	}
}

func TestDeleteMin(t *testing.T) {
	seq, err := NewWith([]int{6, 1, 3, 4, 5})
	checkError(err, t)
	err = seq.DeleteMin()
	checkError(err, t)
	if seq.data[1] != 3 {
		t.Error("func DeleteMin's result unexpected")
	}
}

func TestReverseWithRange(t *testing.T) {
	errMsg := "func ReverseWithRange's result unexpected"
	seq, err := NewWith([]int{1, 2, 3, 4, 5})
	checkError(err, t)
	err = seq.ReverseWithRange(0, 1)
	checkError(err, t)
	result := []int{2, 1, 3, 4, 5}
	for i := 0; i < len(result); i++ {
		v, err := seq.Get(int64(i))
		checkError(err, t)
		if v.(int) != result[i] {
			t.Error(errMsg)
		}
	}
}

func TestDeleteAll(t *testing.T) {
	errMsg := "func DeleteAll's result is unexpected"
	seq, err := NewWith([]int{2, 3, 4, 3, 3, 5, 3, 2})
	checkError(err, t)
	err = seq.DeleteAll(3)
	checkError(err, t)
	result := []int{2, 4, 5, 2}
	if int64(len(result)) != seq.length {
		t.Error(errMsg)
	}
	for i := 0; i < len(result); i++ {
		if seq.data[i] != result[i] {
			t.Error(errMsg)
		}
	}
	err = seq.DeleteAll(2)
	if err != nil {
		t.Error(errMsg)
	}
	result = []int{4, 5}
	if int64(len(result)) != seq.length {
		t.Error(errMsg)
	}
	for i := 0; i < len(result); i++ {
		if seq.data[i] != result[i] {
			t.Error(errMsg)
		}
	}
}

func TestConvertToSorted(t *testing.T) {
	seq, err := NewWith([]int{3, 1, 2, 4, 6, 8, 5, 7})
	checkError(err, t)
	errMsg := "func ConvertToSorted's result is unexpected"
	seq.ConvertToSorted()
	result := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < len(result); i++ {
		if seq.data[i] != result[i] {
			t.Error(errMsg)
		}
	}
}

func TestDeleteRangeElemWithSorted(t *testing.T) {
	errMsg := "func DeleteRangeElemWithSorted's result is unexpected"
	seq, err := NewWith([]int{9, 0, 3, 4, 5, 1, 2, 4, 5, 6, 7})
	checkError(err, t)
	seq.ConvertToSorted()
	err = seq.DeleteRangeElemWithSorted(1, 4)
	checkError(err, t)
	result := []int{0, 5, 5, 6, 7, 9}
	if int64(len(result)) != seq.length {
		t.Error(errMsg)
	}
	for i := 0; i < len(result); i++ {
		if seq.data[i].(int) != result[i] {
			t.Error(errMsg)
		}
	}
	err = seq.DeleteRangeElemWithSorted(5, 9)
	checkError(err, t)
	if seq.length != 1 || seq.data[0].(int) != 0 {
		t.Error(errMsg)
	}
}

func (seq SeqList) TestDeleteRangeElem(t *testing.T) {
	errMsg := "func DeleteRangeElem's result is unexpected"
	seq, err := NewWith([]int{4, 0, 1, 5, 7, 8, 9, 2})
	checkError(err, t)
	err = seq.DeleteRangeElem(2, 7)
	checkError(err, t)
	result := []int{0, 1, 8, 9}
	if int64(len(result)) != seq.length {
		t.Error(errMsg)
	}
	for i := 0; i < len(result); i++ {
		if seq.data[i].(int) != result[i] {
			t.Error(errMsg)
		}
	}
}

func TestDeleteRepeatElemWithSorted(t *testing.T) {
	errMsg := "func DeleteReplaceElemWithSorted's result is unexpected"
	seq, err := NewWith([]int{2, 3, 5, 6, 7, 2, 3, 4, 7, 8})
	checkError(err, t)
	seq.ConvertToSorted()
	err = seq.DeleteRepeatElemWithSorted()
	checkError(err, t)
	result := []int{2, 3, 4, 5, 6, 7, 8}
	if int64(len(result)) != seq.length {
		t.Error(errMsg)
	}
	for i := 0; i < len(result); i++ {
		if seq.data[i].(int) != result[i] {
			t.Error(errMsg)
		}
	}
}

func TestMergeSortedSeqList(t *testing.T) {
	errMsg := "func MergeSortedSeqList's result is unexpected"
	left, err := NewWith([]int{3, 6, 7, 1, 4})
	checkError(err, t)
	right, err := NewWith([]int{2, 8, 5, 0, 9})
	checkError(err, t)
	left.ConvertToSorted()
	right.ConvertToSorted()
	result := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	seq, err := MergeSortedSeqList(left, right)
	checkError(err, t)
	if int64(len(result)) != seq.length {
		t.Error(errMsg)
	}
	for i := 0; i < len(result); i++ {
		if seq.data[i].(int) != result[i] {
			t.Error(errMsg)
		}
	}
}

func TestLocalSwapFrom(t *testing.T) {
	errMsg := "func LocalSwapFrom's result (%d = %d) is unexpected"
	seq, err := NewWith([]int{0, 1, 2, 3, 4, 5, 6})
	checkError(err, t)
	err = seq.LocalSwapFrom(3)
	checkError(err, t)
	result := []int{4, 5, 6, 0, 1, 2, 3}
	if int64(len(result)) != seq.length {
		t.Error(errMsg, len(result), seq.length)
	}
	for i := 0; i < len(result); i++ {
		if seq.data[i].(int) != result[i] {
			t.Errorf(errMsg, seq.data[i].(int), result[i])
		}
	}
}

func TestBinarySearchWithSorted(t *testing.T) {
	errMsg := "func BinarySearchWithSorted's result is unexpected"
	seq, err := NewWith([]int{0, 6, 4, 1, 2, 7, 8})
	checkError(err, t)
	seq.ConvertToSorted()
	index, err := seq.BinarySearchWithSorted(2)
	checkError(err, t)
	if index != 2 {
		t.Error(errMsg)
	}
	index, err = seq.BinarySearchWithSorted(4)
	checkError(err, t)
	if index != 3 {
		t.Error(errMsg)
	}
}

func TestCycleMoveLeftN(t *testing.T) {
	errMsg := "func CycleMoveLeftN's result (%d = %d) is unexpected"
	seq, err := NewWith([]int{1, 2, 3, 4, 5, 6, 7})
	checkError(err, t)
	err = seq.CycleMoveLeftN(3)
	checkError(err, t)
	result := []int{4, 5, 6, 7, 1, 2, 3}
	if int64(len(result)) != seq.length {
		t.Errorf(errMsg, len(result), seq.length)
	}
	for i := 0; i < len(result); i++ {
		if seq.data[i].(int) != result[i] {
			t.Errorf(errMsg, seq.data[i], result[i])
		}
	}
}

func TestFindMiddleWithTwoSortedEqualLengthSequence(t *testing.T) {
	left, err := NewWith([]int{11, 13, 15, 17, 19})
	checkError(err, t)
	right, err := NewWith([]int{2, 4, 6, 8, 20})
	checkError(err, t)
	elem, err := FindMiddleWithTwoSortedEqualLengthSequence(left, right)
	checkError(err, t)
	errMsg := "func FindMiddleWithSortedEqualLengthSequence's result (%d = %d) unexpected"
	if elem.(int) != 11 {
		t.Errorf(errMsg, elem, 11)
	}
}

func TestFindMajorElement(t *testing.T) {
	seq, err := NewWith([]int{0, 5, 5, 3, 5, 7, 5, 5})
	checkError(err, t)
	elem, err := seq.FindMajorElement()
	checkError(err, t)
	errMsg := "func FindMajorElement's result (%d = %d) is unexpected"
	if elem.(int) != 5 {
		t.Errorf(errMsg, elem, 5)
	}
}
