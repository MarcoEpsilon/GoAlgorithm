package linked

import "fmt"

func checkExampleError(err error) {
	if err != nil {
		panic(err)
	}
}

func ExampleNewWith() {
	stack, err := NewWith([]int{1, 2, 3, 4, 5})
	checkExampleError(err)
	for ; !stack.IsEmpty(); {
		elem, err := stack.Pop()
		checkExampleError(err)
		fmt.Println(elem)
	}
	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
}

func ExampleOperation() {
	stack := New()
	err := stack.Push(1)
	checkExampleError(err)
	err = stack.Push(2)
	top, err := stack.Top()
	checkExampleError(err)
	fmt.Println(top)
	elem, err := stack.Pop()
	checkExampleError(err)
	fmt.Println(elem)
	// Output:
	// 2
	// 2
}

func ExampleReverse() {
	stack, err := NewWith([]int{1, 2, 3, 4, 5})
	checkExampleError(err)
	stack.Reverse()
	for ; !stack.IsEmpty(); {
		top, _ := stack.Pop()
		fmt.Println(top)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleIsResult() {
	is, err := IsResult([]int{1,2,3,4,5},[]int{3,2,5,4,1})
	checkExampleError(err)
	fmt.Println(is)
	// Output:
	// true
}

func ExampleSortWithBubble() {
	stack, err := NewWith([]int{5,0,1,4,3,2})
	checkExampleError(err)
	err = stack.SortWithBubble()
	checkExampleError(err)
	for ; !stack.IsEmpty(); {
		top, _ := stack.Pop()
		fmt.Println(top)
	}
	// Ouput:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
}
