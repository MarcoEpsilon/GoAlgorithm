package btree

import (
	"fmt"
	//"bytes"
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

func ExampleLevelOrderBTRLVisit() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	levelBTRLResult := bt.LevelOrderWithBTRLVisit()
	inResult := bt.InOrderVisit()
	for _, v := range levelBTRLResult {
		fmt.Println(v)
	}
	for _, v := range inResult {
		fmt.Println(v)
	}
	// Output:
	// 7
	// 6
	// 5
	// 4
	// 3
	// 2
	// 1
	// 7
	// 4
	// 2
	// 5
	// 1
	// 6
	// 3
}

func ExampleHeight() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	height := bt.Height()
	fmt.Println(height)
	// Output:
	// 4
}

func ExampleIsComplete() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	iscomplete := bt.IsComplete()
	fmt.Println(iscomplete)
	// Output:
	// false
}

func ExampleTwoDegreeNodeCount() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	fmt.Println(bt.TwoDegreeNodeCount())
	// Output:
	// 2
}

func ExampleSwapLeftAndRight() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	bt.SwapLeftAndRight()
	//levelResult := []int{1,3,2,6,5,4,7}
	//inResult := []int{3,6,1,5,2,4,7}
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
	// 3
	// 2
	// 6
	// 5
	// 4
	// 7
	// 3
	// 6
	// 1
	// 5
	// 2
	// 4
	// 7
}

func ExamplePreOrderN() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	preResult := []int{1,2,4,7,5,3,6}
	for i, v := range preResult {
		curr := bt.PreOrderN(i + 1)
		fmt.Println(curr.(int) == v)
	}
	// true
	// true
	// true
	// true
	// true
	// true
	// true
}

func ExampleAncestorsOf() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	ancestors := bt.AncestorsOf(4)
	for _, v := range ancestors {
		fmt.Println(v)
	}
	ancestors = bt.AncestorsOf(5)
	for _, v := range ancestors {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 1
	// 2
}

func ExampleNearCommonAncestorOf() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	ancestor := bt.NearCommonAncestorOf(4, 5)
	fmt.Println(ancestor)
	ancestor = bt.NearCommonAncestorOf(5, 6)
	fmt.Println(ancestor)
	// Output:
	// 2
	// 1
}

func ExampleWidth() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	width := bt.Width()
	fmt.Println(width)
	// Output:
	// 3
}

func ExampleIsSimilar() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	similar := IsSimilar(bt, bt)
	fmt.Println(similar)
	fmt.Println(IsSimilar(bt, nil))
	// Output:
	// true
	// false
}

// 带权路径长度
// 统计叶子节点之和
// 节点类型为 int
// 思路:通过累加层序遍历孩子都为空的节点
func (bt BinaryTree) exampleWPL() (int) {
	if bt.root == nil {
		return 0
	}
	weight := 0
	levelCount := 1
	parentCount := 1
	childCount := 0
	childs := append(make([]BinaryTreeNode, 0), bt.root)
	for ; len(childs) != 0; {
		child := childs[0]
		childs = childs[1:]
		parentCount--
		if child.left == nil && child.right == nil {
			weight += child.data.(int) * levelCount
		}
		if child.left != nil {
			childs = append(childs, child.left)
			childCount++
		}
		if child.right != nil {
			childs = append(childs, child.right)
			childCount++
		}
		if parentCount == 0 {
			parentCount = childCount
			childCount = 0
			levelCount++
		}
	}
	return weight
}

func ExampleWPL() {
	level := []int{1,2,3,4,5,6,7}
	in := []int{7,4,2,5,1,6,3}
	bt, err := NewWithLevelAndIn(level, in)
	checkError(err)
	weight := bt.exampleWPL()
	fmt.Println(weight)
	// Output:
	// 61
}
/*
// 已知表达式内容为字符串
// 中序遍历
func (bt BinaryTree) inOrderConvertToExp() string {
	if bt.root == nil {
		return ""
	}
	buf := new(bytes.Buffer)
	leftChilds := append(make([]BinaryTreeNode, 0), bt.root)
	isLeft := true
	for ; len(leftChilds) != 0; {
		current := leftChilds[len(leftChilds) - 1]
		if isLeft {
			for ; current.left != nil && current.left.left != nil; {
				leftChilds = append(leftChilds, current.left)
			}
		}
		leftChild := leftChilds[len(leftChilds) - 1]
		leftChilds = leftChilds[:len(leftChilds) - 1]
		if isLeft && leftChild.left == nil {
			buf.WriteString("( ")
		} else {
			buf.WriteString(") ")
		}
		// 访问左结点
		if isLeft && leftChild.left != nil {
			buf.WriteString(leftChild.left.data.(string) + " ")
		}
		isLeft = false
		// 访问根结点
		buf.WriteString(leftChild.data.(string) + " ")
		// 添加右子树
		if leftChild.right != nil {
			isLeft = true
			leftChilds = append(leftChilds, leftChild.right)
		}
	}
	return buf.String()
}
// 将已知中序表达式二叉树转换为 中序表达式(已括号表示优先级)
func ExampleInConvert() {
	pre := []string{"*","+","a","b","*","c","-","d"}
	in := []string{"a","+","b","*","c","*","-","d"}
	bt, err := NewWithPreAndIn(pre, in)
	checkError(err)
	result := bt.inOrderConvertToExp()
	fmt.Println(result)
	// Output:
	// ( a + b ) * ( c * ( - d ) )
}
*/

