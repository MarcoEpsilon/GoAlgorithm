package common

import "testing"
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func TestStackForQueue(t *testing.T) {
	queue, err := NewWith([]int{1,2,3,4,5})
	checkError(err)
	result := []int{1,2,3,4,5}
	errMsg := "module StackForQueue's result (%d = %d) is unexpected"
	for i := 0; !queue.IsEmpty(); i++ {
		top, _ := queue.DeQueue()
		if top.(int) != result[i] {
			t.Errorf(errMsg, top, result[i])
		}
	}
} 