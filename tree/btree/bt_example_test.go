package btree

import (
	"fmt"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ExampleNewWithPreAndIn() {
	bt, err := NewWithPreAndIn([]int{1,2,3,4}, []int{1,2,3,4})
	checkError(err)
	result := bt.PreOrderRecursiveVisit()
	for _, v := range result {
		fmt.Println(v)
	}
	result = bt.InOrderRecursiveVisit()
	for _, v := range result {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 1
	// 2
	// 3
	// 4
}

func ExampleNewWithPostAndIn() {
	post := []int{1,2,3,4,5}
	in := []int{1,2,3,4,5}
	bt, err := NewWithPostAndIn(post, in)
	checkError(err)
	postResult := bt.PostOrderRecursiveVisit()
	inResult := bt.InOrderRecursiveVisit()
	for _, v := range postResult {
		fmt.Println(v)
	}
	for _, v := range inResult {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 1
	// 2
	// 3
	// 4
	// 5
}