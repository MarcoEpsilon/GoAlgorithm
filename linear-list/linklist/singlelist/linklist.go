package singlelist

import (
	"errors"
	"reflect"
	// "fmt"
	//"math"
)

var (
	OperationWithEmpty = errors.New("OperationWithEmpty")
	UnExpectedType	   = errors.New("UnexpectedType")
	TypeError		   = errors.New("TypeError")
	OutOfRangeIndex	   = errors.New("OutOfRangeIndex")
	NotSupportCompare  = errors.New("NotSupportCompare")
)
const (
	LessThan		=	-1
	Eq				=	0
	GreaterThan		=	1
	UnCompareable	=	2
)
type Node = *node
type node struct {
	data 	interface{}
	next 	Node
}
/*
	we use no head node linklist
*/
type LinkList = *linkList
type linkList struct {
	head	Node
}

func compare(left interface{},right interface{}) (status int, err error) {
	if reflect.TypeOf(left).Kind() != reflect.TypeOf(right).Kind() {
		return UnCompareable,UnExpectedType
	}
	//ps:switch 的局限性导致大量重复代码
	switch right := right.(type) {
	case int:
		left  := left.(int)
		if left < right {
			return LessThan,err
		} else if left > right {
			return GreaterThan,err
		} else {
			return Eq,err
		}
	case string:
		left := left.(string)
		if left < right {
			return LessThan,err
		} else if left > right {
			return GreaterThan,err
		} else {
			return Eq,err
		}
	case float32:
		left := left.(float32)
		if left < right {
			return LessThan,err
		} else if left > right {
			return GreaterThan,err
		} else {
			return Eq,err
		}
	//ignore int32,int64,uint32,float64 throw to default
	default:
		return UnCompareable,NotSupportCompare
	}
}

/*
	create linkList with empty
*/
func New() (list LinkList) {
	return &linkList {
		head: nil,
	}
}

func (list LinkList) checkType(element interface{}) (err error) {
	if list.head == nil {
		return nil
	} else {
		leftKind := reflect.TypeOf(list.head.data).Kind()
		rightKind := reflect.TypeOf(element).Kind()
		if leftKind != rightKind {
			return UnExpectedType
		}
	}
	return nil
}
/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (list LinkList) Append(element interface{}) (err error) {
	err = list.checkType(element)
	if err != nil {
		return err
	}
	if list.head == nil {
		list.head = &node {
			data: element,
			next: nil,
		}
	} else {
		head := list.head
		for ; head.next != nil; {
			head = head.next
		}
		head.next = &node {
			data:element,
			next:nil,
		}
	}
	return nil
}
/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (list LinkList) Length() int64 {
	var length int64 = 0
	head := list.head
	for ; head != nil; {
		length++
		head = head.next
	}
	return length
}

/*
	时间复杂度: O(1)
	空间复杂度: O(1)
*/
func (list LinkList) IsEmpty() bool {
	return list.head == nil
}




func NewWith(element interface{}) (list LinkList,err error) {
	elementValue := reflect.ValueOf(element)
	list = New()
	switch elementValue.Kind() {
	case reflect.Slice,reflect.Array:
		for i := 0; i < elementValue.Len(); i++ {
			list.Append(elementValue.Index(i).Interface())
		}
		return list,err
	default:
		return nil,TypeError
	}
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (list LinkList) InsertByIndex(index int64,data interface{}) (err error) {
	if index == 0 {
		err = list.checkType(data)
		if err != nil {
			return err
		}
		head := &node {
			data: data,
			next:list.head,
		}
		list.head = head
		return nil
	}
	i := int64(0)
	head := list.head
	for ; head != nil; {
		if i == index - 1 {
			next := head.next
			newNode := &node {
				data: data,
				next: next,
			}
			head.next = newNode
			return nil
		}
		head = head.next
		i++
	}
	return OutOfRangeIndex
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (list LinkList) DeleteByIndex(index int64) (elem interface{},err error) {
	if index == 0 {
		head := list.head
		if head == nil {
			return nil,OutOfRangeIndex
		} else {
			next := head.next
			list.head = next
			return head.data,nil
		}
	}
	head := list.head
	i := int64(0)
	for ; head != nil; {
		if i == index - 1 {
			next := head.next
			if next == nil {
				return nil,OutOfRangeIndex
			} else {
				head.next = next.next
				return next.data,nil
			}
		}
		head = head.next
		i++
	}
	return nil,OutOfRangeIndex
}
/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (list LinkList) Get(index int64) (elem interface{},err error) {
	i := int64(0)
	for head := list.head; head != nil; head = head.next {
		if i == index {
			return head.data,nil
		}
		i++
	}
	return nil,OutOfRangeIndex
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (list LinkList) Modify(index int64,elem interface{}) (err error) {
	err = list.checkType(elem)
	if err != nil {
		return err
	}
	i := int64(0)
	for head := list.head; head != nil; head = head.next {
		if i == index {
			head.data = elem
			return nil
		}
		i++
	}
	return OutOfRangeIndex
}

func deleteAllWithRecusive(pre Node,elem interface{}) (err error) {
	if pre == nil || pre.next == nil {
		return OperationWithEmpty
	}
	current := pre.next
	status,err := compare(current.data,elem)
	if err != nil {
		return err
	}
	if status == Eq {
		pre.next = current.next
		err = deleteAllWithRecusive(pre,elem)
		if err == OperationWithEmpty {
			return nil
		}
		return err
	} 
	return deleteAllWithRecusive(current,elem)
}

/*
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func (list LinkList) DeleteAllWithRecusive(elem interface{}) (err error) {
	if list.head == nil {
		return OperationWithEmpty
	}
	head := list.head
	status,err := compare(head.data,elem) 
	if err != nil {
		return err
	}
	if status == Eq {
		list.head = head.next
		err = list.DeleteAllWithRecusive(elem)
		if err == OperationWithEmpty {
			return nil
		}
		return err
	}
	return deleteAllWithRecusive(head,elem)
}

func (list LinkList) DeleteAll(elem interface{}) (err error) {
	if list.IsEmpty() {
		return OperationWithEmpty
	}
	var pre Node = nil
	deleted := false
	for current := list.head; current != nil; current = current.next {
		currentData := current.data
		status,err := compare(currentData,elem)
		if err != nil {
			return err
		}
		if status == Eq {
			if pre == nil {
				list.head = current.next
			} else {
				pre.next = current.next
			}
			deleted = true
		} else {
			pre = current
		}
	}
	if deleted {
		return nil
	} else {
		return OperationWithEmpty
	}
}

func reverseDoWithRecusive(curr Node,fun func(elem interface{})) {
	if curr == nil {
		return
	}
	reverseDoWithRecusive(curr.next,fun)
	fun(curr.data)
}

func (list LinkList) ReverseDoWithRecusive(fun func(elem interface{})) {
	reverseDoWithRecusive(list.head,fun)
}

func (list LinkList) DeleteMin() (elem interface{},err error) {
	if list.IsEmpty() {
		return nil,OperationWithEmpty
	}
	var minNodePre Node = nil
	for head := list.head; head.next != nil; head = head.next {
		next := head.next
		if minNodePre == nil {
			leftData := list.head.data
			rightData := next.data
			status,err := compare(rightData,leftData)
			if err != nil {
				return nil,err
			}
			if status == LessThan {
				minNodePre = head
			}
		} else {
			leftData := minNodePre.next.data
			rightData := next.data
			status,err := compare(rightData,leftData)
			if err != nil {
				return nil,err
			}
			if status == LessThan {
				minNodePre = head
			}
		}
	}
	var data interface{}
	// delete head node
	if minNodePre == nil {
		next := list.head.next
		data = list.head.data
		list.head = next
	} else {
		next := minNodePre.next
		data = next.data
		minNodePre.next = next.next
	}
	return data,nil
}

func (list LinkList) Reverse() (err error) {
	if list.IsEmpty() {
		return OperationWithEmpty
	}
	var newHead Node = list.head
	head := list.head.next
	newHead.next = nil
	for ; head != nil; {
		next := head.next
		temp := newHead
		newHead = head
		newHead.next = temp
		head = next
	}
	list.head = newHead
	return err
}

func (list LinkList) ReverseBySwap() (err error) {
	if list.IsEmpty() {
		return OperationWithEmpty
	}
	pre := list.head
	current := pre.next
	pre.next = nil
	for ; current != nil; {
		next := current.next
		current.next = pre
		pre = current
		current = next
	}
	list.head = pre
	return nil
}

func (list LinkList) InsertSort() (err error) {
	newHead := list.head
	head := list.head.next
	newHead.next = nil
	for ; head != nil; {
		var pre Node = nil
		for ; pre == nil || pre.next != nil; {
			if pre == nil {
				leftData := head.data
				rightData := newHead.data
				status,err := compare(leftData,rightData)
				if err != nil {
					return err
				}
				if status == LessThan {
					next := head.next
					head.next = newHead
					newHead = head
					head = next
					break
				} else {
					pre = newHead
					continue
				}
			} else {
				preNext := pre.next
				leftData := head.data
				rightData := preNext.data
				status,err := compare(leftData,rightData)
				if err != nil {
					return err
				}
				if status == LessThan {
					next := head.next
					head.next = preNext
					pre.next = head
					head = next
					break
				} else {
					pre = preNext
					continue
				}
			}
		}
		if pre != nil && pre.next == nil {
			next := head.next
			head.next = pre.next
			pre.next = head
			head = next
		}
	}
	list.head = newHead
	return err
}


func (list LinkList) DeleteRangeElem(start interface{},end interface{}) (err error) {
	if list.IsEmpty() {
		return OperationWithEmpty
	}
	status,err := compare(start,end)
	if err != nil {
		return err
	}
	if status == GreaterThan {
		return OutOfRangeIndex
	}
	err = list.checkType(start)
	if err != nil {
		return
	}
	var pre Node = nil
	for pre == nil || pre.next != nil {
		if pre == nil {
			leftData := list.head.data
			startStatus,err := compare(leftData,start)
			if err != nil {
				return err
			}
			endStatus,err := compare(leftData,end)
			if err != nil {
				return err
			}
			if startStatus != LessThan && endStatus != GreaterThan {
				list.head = list.head.next
				if list.head == nil {
					break
				}
			} else {
				pre = list.head
			}
		} else {
			preNext := pre.next
			leftData := preNext.data
			startStatus,err := compare(leftData,start)
			if err != nil {
				return err
			}
			endStatus,err := compare(leftData,end)
			if err != nil {
				return err
			}
			if startStatus != LessThan && endStatus != GreaterThan {
				pre.next = preNext.next
			} else {
				pre = preNext
			}
		}
	}
	return err
}

func NodeToLinkList(n Node) (list LinkList) {
	return &linkList {
		head: n,
	}
}

func (list LinkList) Copy() (LinkList) {
	return &linkList {
		head: list.head,
	}
}

func FindCommonNode(left LinkList,right LinkList) (n Node,err error) {
	if left.IsEmpty() || right.IsEmpty() {
		return nil,OperationWithEmpty
	}
	leftLength := left.Length()
	rightLength := right.Length()
	handle := func(short Node,long Node,diff int64) (node Node,err error) {
		for k := int64(0); k < diff; k++ {
			long = long.next
		}
		for {
			if short == long {
				return short,nil
			}
			short = short.next
			long = long.next
			if short == nil {
				return nil,OperationWithEmpty
			}
		}
	}
	if leftLength <= rightLength {
		return handle(left.head,right.head,rightLength - leftLength)
	} else {
		return handle(right.head,left.head,leftLength - rightLength)
	}
}

/*
	this algorithm should be used to integer
*/
func (list LinkList) SplitToEvenAndOdd() (left LinkList,right LinkList) {
	if list.IsEmpty() {
		return nil,nil
	}
	left = New()
	right = New()
	appendNode := func (mainList LinkList,elem Node) {
		if mainList.head == nil {
			mainList.head = elem
			return
		}
		var head Node = mainList.head
		for ; head.next != nil; {
			head = head.next
		}
		head.next = elem
	}
	var isEven = func (elem interface{}) (bool,error) {
		if reflect.TypeOf(elem).Kind() != reflect.Int {
			return false,TypeError
		}
		num := elem.(int)
		if num % 2 == 0 {
			return true,nil
		} else {
			return false,nil
		}
	}
	for head := list.head; head != nil; {
		status,err := isEven(head.data)
		if err != nil {
			return nil,nil
		}
		next := head.next
		head.next = nil
		if status == true {
			appendNode(left,head)
		} else {
			appendNode(right,head)
		}
		head = next
	}
	list.head = nil
	return left,right
}


func (list LinkList) SplitToNaturalAndReverse() (left LinkList,right LinkList) {
	if list.IsEmpty() {
		return nil,nil
	}
	left = New()
	right = New()
	appendNode := func(mainList LinkList,elem Node) {
		if mainList.head == nil {
			mainList.head = elem
			return
		}
		head := mainList.head
		for ; head.next != nil; {
			head = head.next
		}
		head.next = elem
	}
	pushFront := func(mainList LinkList,elem Node) {
		if mainList.head == nil {
			mainList.head = elem
			return
		}
		next := mainList.head
		elem.next = next
		mainList.head = elem
	}
	i := 0
	for head := list.head; head != nil; {
		next := head.next
		head.next = nil
		if i % 2 == 0 {
			appendNode(left,head)
		} else {
			pushFront(right,head)
		}
		head = next
		i++
	}
	list.head = nil
	return left,right
}

func (list LinkList) DeleteRepeatWithSorted() (err error) {
	if list.IsEmpty() {
		return OperationWithEmpty
	}
	pre := list.head
	current := pre.next
	for ; current != nil; {
		next := current.next
		preData := pre.data
		currentData := current.data
		status,err := compare(preData,currentData)
		if err != nil {
			return err
		}
		if status == Eq {
			pre.next = next
			current = next
		} else {
			pre = current
			current = next
		}
	}
	return nil
}

func ReverseMergeSortedLinkList(left LinkList,right LinkList) (list LinkList,err error) {
	if left.IsEmpty() && right.IsEmpty() {
		return nil,OperationWithEmpty
	}
	list = New()
	leftCurrent := left.head
	rightCurrent := right.head
	pushFront := func (mainList LinkList,elem Node) {
		if mainList.head == nil {
			mainList.head = elem
			return
		}
		next := mainList.head
		elem.next = next
		mainList.head = elem
	}
	for ; leftCurrent != nil && rightCurrent != nil; {
		leftData := leftCurrent.data
		rightData := rightCurrent.data
		status,err := compare(leftData,rightData)
		if err != nil {
			return nil,err
		}
		if status != GreaterThan {
			leftNext := leftCurrent.next
			leftCurrent.next = nil
			pushFront(list,leftCurrent)
			leftCurrent = leftNext
		} else {
			rightNext := rightCurrent.next
			rightCurrent.next = nil
			pushFront(list,rightCurrent)
			rightCurrent = rightNext
		}
	}
	for ; leftCurrent != nil; {
		leftNext := leftCurrent.next
		leftCurrent.next = nil
		pushFront(list,leftCurrent)
		leftCurrent = leftNext
	}
	for ; rightCurrent != nil; {
		rightNext := rightCurrent.next
		rightCurrent.next = nil
		pushFront(list,rightCurrent)
		rightCurrent = rightNext
	}
	left.head = nil
	right.head = nil
	return list,nil
}


func GetCommonWithSortedLinkList(left LinkList,right LinkList) (list LinkList,err error) {
	if left.IsEmpty() || right.IsEmpty() {
		return nil,OperationWithEmpty
	}
	leftCurrent := left.head
	rightCurrent := right.head
	list = New()
	for ; leftCurrent != nil && rightCurrent != nil; {
		leftData := leftCurrent.data
		rightData := rightCurrent.data
		status,err := compare(leftData,rightData)
		if err != nil {
			return nil,err
		}
		if status == Eq {
			list.Append(leftData)
			leftCurrent = leftCurrent.next
			rightCurrent = rightCurrent
		} else if status == LessThan {
			leftCurrent = leftCurrent.next
		} else {
			rightCurrent = rightCurrent.next
		}
	}
	if list.IsEmpty() {
		return nil,OperationWithEmpty
	}
	return list,nil
}


func (list LinkList) IsSubSequenceOf(main LinkList) (bool,error) {
	if list.IsEmpty() || main.IsEmpty() {
		return false,OperationWithEmpty
	}
	mainCurrent := main.head
	listCurrent := list.head
	for ; mainCurrent != nil; {
		mainData := mainCurrent.data
		listData := listCurrent.data
		status,err := compare(mainData,listData)
		if err != nil {
			return false,err
		}
		if status == Eq {
			mainCurrent = mainCurrent.next
			listCurrent = listCurrent.next
			if listCurrent == nil {
				return true,nil
			}
		} else {
			mainCurrent = mainCurrent.next
			listCurrent = list.head
		}
	}
	return false,nil
}


func (list LinkList) FindLastN(n int) (elem interface{},err error) {
	if n <= 0 {
		return nil,OutOfRangeIndex
	}
	//slow pointer && quick pointer
	slow := list.head
	quick := list.head
	// move quick pointer
	for step := 0; step < n - 1; step++ {
		if quick == nil {
			return nil,OutOfRangeIndex
		}
		quick = quick.next
	}
	for ; quick.next != nil; {
		slow = slow.next
		quick = quick.next
	}
	return slow.data,nil
}

func findLastNWithRecusive(current Node,index int) (Node,int) {
	if current != nil {
		ret,n := findLastNWithRecusive(current.next,index)
		if n == index - 1 {
			return current,n + 1
		} else {
			return ret,n + 1
		}
	} else {
		return nil,0
	}
}

func (list LinkList) FindLastNWithRecusive(n int) (elem interface{},err error) {
	if n <= 0 {
		return nil,OperationWithEmpty
	}
	ret,_ := findLastNWithRecusive(list.head,n)
	if ret != nil {
		return ret.data,nil
	} else {
		return nil,nil
	}
}

func (list LinkList) NeighborReverse() (err error) {
	if list.IsEmpty() {
		return OperationWithEmpty
	}
	swap := func(pre Node) (Node,Node) {
		if pre.next == nil {
			return nil,nil
		}
		curr := pre.next
		next := curr.next
		curr.next = pre
		return curr,next
	}
	var pre Node = nil
	current := list.head
	for {
		newHead,next := swap(current)
		if next == nil {
			if pre != nil {
				pre.next = current
			}
			break
		}
		if pre == nil {
			list.head = newHead
			pre = newHead.next
		} else {
			pre.next = newHead
			pre = newHead.next
		}
		current = next
	}
	return nil
}
// ReverseN(2) == NeighborReverse
func (list LinkList) ReverseN(n int) (err error) {
	if n <= 0 {
		return OutOfRangeIndex
	}
	if list.IsEmpty() {
		return OperationWithEmpty
	}
	var newListEnd Node = nil
	var preNode Node = nil
	currentNode := list.head
	group := list.Length() / int64(n)
	for k := int64(0); k < group; k++ {
		newEndNode := currentNode
		for i := 0; i < n; i++ {
			nextCurrent := currentNode.next
			currentNode.next = preNode
			preNode = currentNode
			currentNode = nextCurrent
		}
		if newListEnd == nil {
			list.head = preNode
			newListEnd = newEndNode
		} else {
			newListEnd.next = preNode
			newListEnd = newEndNode
		}
		preNode = nil
	}
	//link the rest nodes
	newListEnd.next = currentNode
	return nil
}

func (list LinkList) IsCycle() (bool) {
	fast := list.head
	slow := list.head
	for {
		if fast == nil || fast.next == nil {
			return false
		}
		fast = fast.next.next
		slow = slow.next
		if slow == fast {
			return true
		}
	}
}

func (list LinkList) FindLoopNode() (Node,error) {
	fast := list.head
	slow := list.head
	for {
		if fast == nil || fast.next == nil {
			return nil,OperationWithEmpty
		}
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			for slow = list.head; slow != fast; {
				slow = slow.next
				fast = fast.next
			}
			return slow,nil
		}
	}
}


func (list LinkList) LoopStartIndex() (int,error) {
	fast := list.head
	slow := list.head
	for {
		if fast == nil || fast.next == nil {
			return 0,OperationWithEmpty
		}
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			i := 0
			for slow = list.head; slow != fast; {
				slow = slow.next
				fast = fast.next
				i++
			}
			return i,nil
		}
	}
}

func (list LinkList) LoopLength() (int,error) {
	fast := list.head
	slow := list.head
	for {
		if fast == nil || fast.next == nil {
			return 0,OperationWithEmpty
		}
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			i := 1
			slow = fast.next
			for ; slow != fast; i++ {
				slow = slow.next
			}
			return i,nil
		}
	}
}