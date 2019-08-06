package seq

import "fmt"

func checkExampleError(err error) {
	if err != nil {
		panic(err)
	}
}
func ExampleOperation() {
	queue := New()
	init := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(init); i++ {
		err := queue.EnQueue(init[i])
		checkExampleError(err)
	}
	for !queue.IsEmpty() {
		elem, err := queue.DeQueue()
		checkExampleError(err)
		fmt.Println(elem)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}
