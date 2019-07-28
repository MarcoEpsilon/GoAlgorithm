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