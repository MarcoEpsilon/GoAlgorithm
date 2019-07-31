package singlelist

import "fmt"
func ExampleReverseDoWithRecusive() {
	list,err := NewWith([]int{1,2,3,4,5})
	if err != nil {
		panic(err)
	}
	list.ReverseDoWithRecusive(func (elem interface{}) {fmt.Println(elem)})
	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
}

func getTailNode(n Node) (Node) {
	if n == nil {
		return nil
	}
	for ; n.next != nil; n = n.next {
	}
	return n
}
func ExampleIsCycle() {
	list,err := NewWith([]int{1,2,3,4})
	if err != nil {
		panic(err)
	}
	listTail := getTailNode(list.head)
	listTail.next = list.head
	cycle := list.IsCycle()
	fmt.Println(cycle)
	// Output:
	// true
}

func ExampleFindLoopNode() {
	list,err := NewWith([]int{1,2,3,4,5,6})
	if err != nil {
		panic(err)
	}
	subList,err := NewWith([]int{7,8,9})
	if err != nil {
		panic(err)
	}
	listTail := getTailNode(list.head)
	subTail := getTailNode(subList.head)
	listTail.next = subList.head
	subTail.next = list.head.next.next
	elem,err := list.FindLoopNode()
	if err != nil {
		panic(err)
	}
	fmt.Println(elem.data)
	// Output:
	// 3
}

func ExampleLoopNodeIndex() {
	list,err := NewWith([]int{1,2,3,4,5,6})
	if err != nil {
		panic(err)
	}
	subList,err := NewWith([]int{7,8,9})
	if err != nil {
		panic(err)
	}
	listTail := getTailNode(list.head)
	subTail := getTailNode(subList.head)
	listTail.next = subList.head
	subTail.next = list.head.next.next
	index,err := list.LoopStartIndex()
	if err != nil {
		panic(err)
	}
	fmt.Println(index)
	// Output:
	// 2
}


func ExampleLoopLength() {
	list,err := NewWith([]int{1,2,3,4,5,6})
	if err != nil {
		panic(err)
	}
	subList,err := NewWith([]int{7,8,9})
	if err != nil {
		panic(err)
	}
	listTail := getTailNode(list.head)
	subTail := getTailNode(subList.head)
	listTail.next = subList.head
	subTail.next = list.head.next.next
	length,err := list.LoopLength()
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	// Output:
	// 7
}