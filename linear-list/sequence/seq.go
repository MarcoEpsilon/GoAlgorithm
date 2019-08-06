package sequence

import (
	"errors"
	"reflect"
	"sort"
	// "fmt"
)

var (
	// index out of range
	OutOfRangeIndex    = errors.New("OutOfRangeIndex")
	TypeError          = errors.New("TypeError")
	UnExpectedType     = errors.New("UnExpectedType")
	OperationWithEmpty = errors.New("OperationWithEmpty")
	NotSupportCompare  = errors.New("NotSupportCompare")
	IsNotSorted        = errors.New("IsNotSorted")
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

//constant for seqlist max length
const maxSIZE = 1 << 16

type seqList struct {
	length int64
	data   [maxSIZE]interface{}
}
type SeqList = *seqList

// create SeqList
func New(element interface{}, length int64) (seq SeqList, err error) {
	var seqlist seqList
	if length > maxSIZE || length < 0 {
		return nil, OutOfRangeIndex
	}
	seqlist.length = length
	for i := int64(0); i < length; i++ {
		seqlist.data[i] = element
	}
	return &seqlist, err
}
func NewWith(elements interface{}) (seq SeqList, err error) {
	var seqlist seqList
	seqlist.length = 0
	rawValue := reflect.ValueOf(elements)
	switch rawValue.Kind() {
	case reflect.Slice, reflect.Array:
		err = seqlist.checkAdding(int64(rawValue.Len()))
		if err != nil {
			return nil, err
		}
		for i := 0; i < rawValue.Len(); i++ {
			seqlist.data[i] = rawValue.Index(i).Interface()
		}
		seqlist.length = int64(rawValue.Len())
		return &seqlist, err
	default:
		return nil, UnExpectedType
	}
}

// return the length of seqList
func (seq SeqList) Length() int64 {
	return seq.length
}

// judge whether is empty seqlist
func (seq SeqList) IsEmpty() bool {
	return seq.length == 0
}

//check the operation's index
func (seq SeqList) checkIndex(index int64) (err error) {
	if index < 0 || index >= seq.length {
		return OutOfRangeIndex
	}
	return err
}

// check the rest cap of seqlist
func (seq SeqList) checkAdding(cap int64) (err error) {
	if cap < 0 || (seq.length+cap) >= maxSIZE {
		return OutOfRangeIndex
	}
	return err
}

// check the type whether is identical
func (seq SeqList) checkType(element interface{}) (err error) {
	if seq.IsEmpty() {
		return err
	}
	expectedType := reflect.TypeOf(seq.data[0]).Kind()
	rawType := reflect.TypeOf(element).Kind()
	if expectedType != rawType {
		err = TypeError
	}
	return err
}

/*
	时间复杂度:O(1)
	空间复杂度:O(1)
*/
func (seq SeqList) Get(index int64) (value interface{}, err error) {
	err = seq.checkIndex(index)
	if err != nil {
		return nil, err
	}
	return seq.data[index], err
}

/*
	时间复杂度: O(1)
	空间复杂度: O(1)
*/
func (seq SeqList) Modify(index int64, element interface{}) (err error) {
	err = seq.checkIndex(index)
	if err != nil {
		return err
	}
	err = seq.checkType(element)
	if err != nil {
		return err
	}
	seq.data[index] = element
	return err
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (seq SeqList) InsertByIndex(index int64, element interface{}) (err error) {
	err = seq.checkType(element)
	if err != nil {
		return err
	}
	err = seq.checkIndex(index)
	if err != nil {
		return err
	}
	err = seq.checkAdding(1)
	if err != nil {
		return err
	}
	for i := seq.length; i != index; i-- {
		seq.data[i] = seq.data[i-1]
	}
	seq.data[index] = element
	return err
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (seq SeqList) DeleteByIndex(index int64) (err error) {
	err = seq.checkIndex(index)
	if err != nil {
		return err
	}
	for i := index; i < seq.length-1; i++ {
		seq.data[i] = seq.data[i+1]
	}
	seq.length = seq.length - 1
	return err
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
	fixme: now the operator "<" is supported with
	number,string type
*/
func (seq SeqList) DeleteMin() (err error) {
	if seq.length <= 0 {
		return OperationWithEmpty
	}
	minIndex := int64(0)
	for i := int64(1); i < seq.length; i++ {
		status, err := compare(seq.data[i], seq.data[minIndex])
		if err != nil {
			return err
		}
		if status == LessThan {
			minIndex = i
		}
	}
	// O(n)操作
	return seq.DeleteByIndex(minIndex)
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (seq SeqList) ReverseWithRange(start int64, end int64) (err error) {
	if seq.checkIndex(start) != nil || seq.checkIndex(end) != nil {
		return OutOfRangeIndex
	}
	if start >= end {
		return OperationWithEmpty
	}
	var k int64 = (end + start + 1) / 2
	for i := int64(start); i < k; i++ {
		temp := seq.data[i]
		seq.data[i] = seq.data[end+start-i]
		seq.data[end+start-i] = temp
	}
	return err
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (seq SeqList) Reverse() (err error) {
	return seq.ReverseWithRange(0, seq.length-1)
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (seq SeqList) deleteAllWithIdenticalTag(elem interface{}) (err error) {
	if seq.IsEmpty() {
		return OperationWithEmpty
	}
	//set identical element's count
	var k int64 = 0
	for i := int64(0); i < seq.length; i++ {
		status, err := compare(seq.data[i], elem)
		if err != nil {
			return err
		}
		if status != Eq {
			seq.data[i-k] = seq.data[i]
		} else {
			k++
		}
	}
	//update seqlist's length
	seq.length = seq.length - k
	return err
}
func (seq SeqList) deleteAllWithDifferentTag(elem interface{}) (err error) {
	if seq.IsEmpty() {
		return OperationWithEmpty
	}
	// set different element's count
	var k int64 = 0
	for i := int64(0); i < seq.length; i++ {
		status, err := compare(seq.data[i], elem)
		if err != nil {
			return err
		}
		if status != Eq {
			seq.data[k] = seq.data[i]
			k++
		}
	}
	seq.length = k
	return err
}
func (seq SeqList) DeleteAll(elem interface{}) (err error) {
	//return seq.deleteAllWithIdenticalTag(elem)
	return seq.deleteAllWithDifferentTag(elem)
}

/*
	为seqlist实现sort.Interface接口
*/
func (seq SeqList) Len() int {
	return int(seq.Length())
}
func (seq SeqList) Less(i int, j int) bool {
	status, err := compare(seq.data[i], seq.data[j])
	if err != nil {
		panic("not Compareable")
	}
	if status == LessThan {
		return true
	}
	return false
}

func (seq SeqList) Swap(i, j int) {
	temp := seq.data[i]
	seq.data[i] = seq.data[j]
	seq.data[j] = temp
}

// notice: no check for the type
func (seq SeqList) ConvertToSorted() {
	sort.Sort(seq)
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (seq SeqList) DeleteRangeElemWithSorted(start interface{}, end interface{}) (err error) {
	if seq.IsEmpty() {
		return OperationWithEmpty
	}
	status, err := compare(start, end)
	if err != nil {
		return err
	}
	if status == GreaterThan {
		return OperationWithEmpty
	}
	if !sort.IsSorted(seq) {
		return IsNotSorted
	}
	err = seq.checkType(start)
	if err != nil {
		return err
	}
	//start index,end index
	var sindex int64
	var eindex int64
	for sindex = 0; sindex < seq.length; sindex++ {
		status, err = compare(seq.data[sindex], start)
		if err != nil {
			return err
		}
		if status != LessThan {
			break
		}
	}
	for eindex = sindex; eindex < seq.length; eindex++ {
		status, err = compare(seq.data[eindex], end)
		if err != nil {
			return err
		}
		if status == GreaterThan {
			break
		}
	}
	//compute the total count of that should delete
	//notice seq.data[eindex] > end,so ....
	total := eindex - sindex
	if total <= 0 {
		return OperationWithEmpty
	}
	//begin delete
	for i := eindex; i < seq.length; i++ {
		seq.data[i-total] = seq.data[i]
	}
	//update seqlist length
	seq.length = seq.length - total
	return err
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (seq SeqList) DeleteRangeElem(start interface{}, end interface{}) (err error) {
	if seq.IsEmpty() {
		return OperationWithEmpty
	}
	status, err := compare(start, end)
	if err != nil {
		return err
	}
	if status == GreaterThan {
		return OperationWithEmpty
	}
	//counter for elements of between start and end
	var k int64 = 0
	for i := int64(0); i < seq.length; i++ {
		left, err := compare(seq.data[i], start)
		if err != nil {
			return err
		}
		right, err := compare(seq.data[i], end)
		if err != nil {
			return err
		}
		if left == LessThan || right == GreaterThan {
			seq.data[i-k] = seq.data[i]
		} else {
			k++
		}
	}
	if k == 0 {
		err = OperationWithEmpty
	}
	//update the seqlist's length
	seq.length = seq.length - k
	return err
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (seq SeqList) DeleteRepeatElemWithSorted() (err error) {
	if seq.IsEmpty() {
		return OperationWithEmpty
	}
	if !sort.IsSorted(seq) {
		return IsNotSorted
	}
	//the begining of not unrepeat element
	var k int64 = 0
	for i := int64(1); i < seq.length; i++ {
		status, err := compare(seq.data[i], seq.data[k])
		if err != nil {
			return err
		}
		if status != Eq {
			k++
			seq.data[k] = seq.data[i]
		}
	}
	seq.length = k + 1
	return err
}
func (seq SeqList) Copy() (newseq SeqList) {
	seqlist := seqList{
		length: seq.length,
		data:   seq.data,
	}
	return &seqlist
}

/*
	设left.length = m,right.length = n
	则
		时间复杂度: O(m + n)
		空间复杂度: O(1)
*/
func MergeSortedSeqList(left SeqList, right SeqList) (seq SeqList, err error) {
	if left.IsEmpty() {
		if right.IsEmpty() {
			return right.Copy(), OperationWithEmpty
		} else {
			return right.Copy(), err
		}
	} else {
		if right.IsEmpty() {
			return left.Copy(), err
		}
	}
	if !sort.IsSorted(left) || !sort.IsSorted(right) {
		return nil, IsNotSorted
	}
	var newseq seqList
	newseq.length = 0
	//left seqlist index i,right seqlist index j
	var i int64 = 0
	var j int64 = 0
	for i < left.length && j < right.length {
		status, err := compare(left.data[i], right.data[j])
		if err != nil {
			return nil, err
		}
		if status == LessThan {
			newseq.data[i+j] = left.data[i]
			i++
		} else {
			newseq.data[i+j] = right.data[j]
			j++
		}
	}
	for ; i < left.length; i++ {
		newseq.data[i+j] = left.data[i]
	}
	for ; j < right.length; j++ {
		newseq.data[i+j] = right.data[j]
	}
	//update seqlist's length
	newseq.length = left.length + right.length
	return &newseq, err
}

/*
	example the seqlist is: b1b2b3....bma1a2....an
		index = m
	after the operation the seqlist is a1a2...anb1...bm
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (seq SeqList) LocalSwapFrom(index int64) (err error) {
	err = seq.checkIndex(index)
	if err != nil {
		return err
	}
	err = seq.ReverseWithRange(0, seq.length-1)
	if err != nil {
		return err
	}
	err = seq.ReverseWithRange(0, seq.length-(index+1)-1)
	if err != nil {
		//ignore the operation with empty
		if err != OperationWithEmpty {
			return err
		}
	}
	err = seq.ReverseWithRange(seq.length-(index+1), seq.length-1)
	if err != nil {
		//ignore the operation with empty
		if err != OperationWithEmpty {
			return err
		}
	}
	err = nil
	return err
}

/*
	时间复杂度: O(logn)
	空间复杂度: O(1)
*/
func (seq SeqList) BinarySearchWithSorted(elem interface{}) (index int64, err error) {
	if seq.IsEmpty() {
		return -1, OperationWithEmpty
	}
	if !sort.IsSorted(seq) {
		return -1, IsNotSorted
	}
	var low int64 = 0
	var high int64 = seq.length - 1
	var middle int64
	for low <= high {
		middle = (high + low) / 2
		status, err := compare(elem, seq.data[middle])
		if err != nil {
			return -1, err
		}
		if status == Eq {
			return middle, err
		} else if status == LessThan {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1, OperationWithEmpty
}

/*
	example: we have a0a1a2....apap+1....as
	n = p
	the result is apap+1...asa0a1....ap-1
*/
func (seq SeqList) CycleMoveLeftN(n int64) (err error) {
	if n < 0 || n > seq.length {
		return OutOfRangeIndex
	}
	if n == 0 || n == seq.length {
		return OperationWithEmpty
	}
	err = seq.ReverseWithRange(0, seq.length-1)
	if err != nil {
		return err
	}
	err = seq.ReverseWithRange(0, seq.length-n-1)
	if err != nil {
		//ignore empty empty operation
		if err != OperationWithEmpty {
			return err
		}
	}
	err = seq.ReverseWithRange(seq.length-n, seq.length-1)
	if err != nil {
		//ignore empty operation
		if err != OperationWithEmpty {
			return err
		}
	}
	return err
}

/*
	left: a1a2....an
	right: b1b2b3...bn
*/
func FindMiddleWithTwoSortedEqualLengthSequence(left SeqList, right SeqList) (elem interface{}, err error) {
	if left.length != right.length {
		return nil, OutOfRangeIndex
	}
	if left.IsEmpty() {
		return nil, OperationWithEmpty
	}
	if !sort.IsSorted(left) || !sort.IsSorted(right) {
		return nil, IsNotSorted
	}
	// find the middle of left
	var leftLow, rightLow = int64(0), int64(0)
	var leftHigh, rightHigh = left.length - 1, right.length - 1
	for leftLow < leftHigh && rightLow < rightHigh {
		leftMiddle := (leftHigh + leftLow) / 2
		rightMiddle := (rightHigh + rightLow) / 2
		leftMiddleElem := left.data[leftMiddle]
		rightMiddleElem := right.data[rightMiddle]
		status, err := compare(leftMiddleElem, rightMiddleElem)
		if err != nil {
			return nil, err
		}
		if status == Eq {
			return leftMiddleElem, err
		} else if status == LessThan {
			// 元素为奇数
			if (leftLow+leftHigh)%2 == 0 {
				leftLow = leftMiddle
				rightHigh = rightMiddle
			} else {
				leftLow = leftMiddle + 1
				rightHigh = rightMiddle
			}
		} else {
			if (rightLow+rightHigh)%2 == 0 {
				rightLow = rightMiddle
				leftHigh = leftMiddle
			} else {
				rightLow = rightMiddle + 1
				leftHigh = leftMiddle
			}
		}
	}
	leftMiddleElem := left.data[leftLow]
	rightMiddleElem := right.data[rightLow]
	status, err := compare(leftMiddleElem, rightMiddleElem)
	if err != nil {
		return nil, err
	}
	if status == LessThan {
		return leftMiddleElem, err
	} else {
		return rightMiddleElem, err
	}
}

func (seq SeqList) FindMajorElement() (elem interface{}, err error) {
	if seq.IsEmpty() {
		return nil, OperationWithEmpty
	}
	//the major element occur count
	var count int64 = 1
	majorElement := seq.data[0]
	for i := int64(1); i < seq.length; i++ {
		status, err := compare(majorElement, seq.data[i])
		if err != nil {
			return nil, err
		}
		if status == Eq {
			count++
		} else {
			if count > 0 {
				count--
			} else {
				majorElement = seq.data[i]
				count = 1
			}
		}
	}
	if count > 0 {
		count = 0
		for i := int64(0); i < seq.length; i++ {
			status, err := compare(seq.data[i], majorElement)
			if err != nil {
				return nil, err
			}
			if status == Eq {
				count++
			}
		}
		if count > seq.length/2 {
			return majorElement, err
		}
	} else {
		return nil, OperationWithEmpty
	}
	return nil, OperationWithEmpty
}
