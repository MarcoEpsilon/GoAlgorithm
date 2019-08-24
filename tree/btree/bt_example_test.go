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

func ExampleLevelOrderVisit() {
	pre := []int{1,2,3,4,5}
	in := []int{1,2,3,4,5}
	bt, err := NewWithPreAndIn(pre, in)
	checkError(err)
	level := bt.LevelOrderVisit()
	for _, v := range level {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleNewWithLevelAndIn() {
	level := []int{1,2,3,4,5}
	in := []int{1,2,3,4,5}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	levelResult := bt.LevelOrderVisit()
	inResult := bt.InOrderRecursiveVisit()
	for _, v := range levelResult {
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

func ExamplePreAndInVisit() {
	pre := []int{2,1,3,4,5}
	in := []int{1,2,3,5,4}
	bt, err := NewWithPreAndIn(pre, in)
	checkError(err)
	preResult := bt.PreOrderVisit()
	inResult := bt.InOrderVisit()
	for _, v := range preResult {
		fmt.Println(v)
	}
	for _, v := range inResult {
		fmt.Println(v)
	}
	// Output:
	// 2
	// 1
	// 3
	// 4
	// 5
	// 1
	// 2
	// 3
	// 5
	// 4
}

// fail example:
/*
 post: 5 3 2 1 4
 in: 2 3 1 4 5
 fail to constructor
*/
func ExamplePostAndInVisit() {
	post := []int{3,2,1,5,4}
	in := []int{2,3,1,4,5}
	bt, err := NewWithPostAndIn(post, in)
	checkError(err)
	postResult := bt.PostOrderRecursiveVisit()
	inResult := bt.InOrderVisit()
	for _, v := range postResult {
		fmt.Println(v)
	}
	for _, v := range inResult {
		fmt.Println(v)
	}
	// Output:
	// 3
	// 2
	// 1
	// 5
	// 4
	// 2
	// 3
	// 1
	// 4
	// 5
}

func ExampleLevelAndInVisit() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	levelResult := bt.LevelOrderVisit()
	inResult := bt.InOrderVisit()
	for _, v := range levelResult {
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
	// 6
	// 7
	// 7
	// 4
	// 2
	// 5
	// 1
	// 6
	// 3
}