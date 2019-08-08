package linked

import "fmt"

func checkExampleError(err error) {
	if err != nil {
		panic(err)
	}
}
func ExampleOperation() {
	queue, err := NewWith([]int{1,2,3,4,5})
	checkExampleError(err)
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Length())
	for ; !queue.IsEmpty(); {
		elem, _ := queue.DeQueue()
		fmt.Println(elem)
	}
	// Output:
	// false
	// 5
	// 1
	// 2
	// 3
	// 4
	// 5
}