package btree

import (
	"errors"
	"reflect"
	//"fmt"
)

var (
	UnExpectedType = errors.New("UnExpectedType")
	TypeError = errors.New("TypeError")
	InvalidSequence = errors.New("InvalidSequence")
	InvalidPreAndIn = errors.New("InvalidPreAndIn")
	InvalidPostAndIn = errors.New("InvalidPostAndIn")
	InvalidLevelAndIn = errors.New("InvalidLevelAndIn")
)

const (
	LessThan = -1
	Eq = 0
	MoreThan = 1
	UnCompareable = 2
)
func compare(left interface{}, right interface{}) (int, error) {
	if !isSameType(left, right) {
		return UnCompareable, UnExpectedType
	}
	switch left := left.(type) {
	case int:
		right := right.(int)
		if left < right {
			return LessThan, nil
		} else if left > right {
			return MoreThan, nil
		} else {
			return Eq, nil
		}
	case string:
		right := right.(string)
		if left < right {
			return LessThan, nil
		} else if left > right {
			return MoreThan, nil
		} else {
			return Eq, nil
		}
	default:
		return UnCompareable, UnExpectedType
	}
}
type BinaryTreeNode = *binarytree_inner_node
type binarytree_inner_node struct {
	left BinaryTreeNode
	right BinaryTreeNode
	data interface{}
}
type BinaryTree = *binarytree_inner
type binarytree_inner struct {
	root BinaryTreeNode
}

func New() BinaryTree {
	return &binarytree_inner {
		root: nil,
	}
}

func NewNode(data interface{}) BinaryTreeNode {
	return &binarytree_inner_node {
		left: nil,
		right: nil,
		data: data,
	}
}
type Sequence = *sequence
type sequence struct {
	reflect.Value
}
type PreSequence = Sequence
type InSequence = Sequence
type PostSequence = Sequence
type LevelSequence = Sequence
type WrapPreAndIn = *wrapPreAndIn
type wrapPreAndIn struct {
	pre PreSequence
	in InSequence
}
type WrapPostAndIn = *wrapPostAndIn
type wrapPostAndIn struct {
	post PostSequence
	in InSequence
}
type WrapLevelAndIn = *wrapLevelAndIn
type wrapLevelAndIn struct {
	level LevelSequence
	in InSequence
}
// we should require eq? maybe better
func (inseq InSequence) IsLeft(root *reflect.Value, left *reflect.Value) bool {
	panicMsg := "Not Support Comapre"
	for i := 0; i < inseq.Len(); i++ {
		first, err := compare(left.Interface(), inseq.Index(i).Interface())
		if err != nil {
			panic(panicMsg)
		}
		if first == Eq {
			return true
		}
		first, err = compare(root.Interface(), inseq.Index(i).Interface())
		if err != nil {
			panic(panicMsg)
		}
		if first == Eq {
			return false
		}
	}
	panic("Invalid InOrder Sequence")
}


func isSameType(left interface{}, right interface{}) (bool) {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)
	return leftType == rightType
}

func isCouldType(left interface{}, right interface{}) (bool, error) {
	if !isSameType(left, right) {
		return false, UnExpectedType
	}
	leftType := reflect.TypeOf(left)
	//rightType := reflect.TypeOf(right)
	switch leftType.Kind() {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(left).Len() != reflect.ValueOf(right).Len() {
			return false, InvalidSequence
		}
		return true, nil
	default:
		return false, TypeError
	}
}

func NewWrapPreAndIn(pre interface{}, in interface{}) (WrapPreAndIn, error) {
	if is, err := isCouldType(pre, in); !is {
		return nil, err
	}
	return &wrapPreAndIn {
		pre: PreSequence(&sequence{reflect.ValueOf(pre)}),
		in: InSequence(&sequence{reflect.ValueOf(in)}),
	}, nil
}

// 通过前序和中序遍历结果构造二叉树
func NewWithPreAndIn(pre interface{}, in interface{}) (BinaryTree, error) {
	wrap, err := NewWrapPreAndIn(pre, in)
	if err != nil {
		return nil, err
	}
	bt := New()
	if wrap.pre.Len() == 0 {
		return bt, nil
	}
	bt.root = NewNode(wrap.pre.Index(0).Interface())
	for i := 1; i != wrap.pre.Len(); i++ {
		waitValue := wrap.pre.Index(i)
		for currentNode := bt.root; currentNode != nil; {
			currentValue := reflect.ValueOf(currentNode.data)
			isleft := wrap.in.IsLeft(&currentValue, &waitValue)
			if isleft && currentNode.left == nil {
				//这表明该节点构造序列不合理
				if isleft && currentNode.right != nil {
					return nil, InvalidPreAndIn
				}
				currentNode.left = NewNode(waitValue.Interface())
				break
			}
			if !isleft && currentNode.right == nil {
				currentNode.right = NewNode(waitValue.Interface())
				break
			}
			if isleft {
				currentNode = currentNode.left
			} else {
				currentNode = currentNode.right
			}
		}
	}
	return bt, nil
}

func NewWrapPostAndIn(post interface{}, in interface{}) (WrapPostAndIn, error) {
	if is, err := isCouldType(post, in); !is {
		return nil, err
	}
	return &wrapPostAndIn {
		post: PostSequence(&sequence{reflect.ValueOf(post)}),
		in: InSequence(&sequence{reflect.ValueOf(in)}),
	}, nil
}

// 通过后序和中序遍历结果构造二叉树
func NewWithPostAndIn(post interface{}, in interface{}) (BinaryTree, error) {
	wrap, err := NewWrapPostAndIn(post, in)
	if err != nil {
		return nil, err
	}
	bt := New()
	if wrap.post.Len() == 0 {
		return bt, nil
	}
	bt.root = NewNode(wrap.post.Index(wrap.post.Len() - 1).Interface())
	for i := wrap.post.Len() - 2; i >= 0; i-- {
		waitValue := wrap.post.Index(i)
		for currentNode := bt.root; currentNode != nil; {
			currentValue := reflect.ValueOf(currentNode.data)
			isleft := wrap.in.IsLeft(&currentValue, &waitValue)
			if isleft && currentNode.left == nil {
				currentNode.left = NewNode(waitValue.Interface())
				break
			}
			if !isleft && currentNode.right == nil {
				if currentNode.left != nil {
					return nil, InvalidPostAndIn
				}
				currentNode.right = NewNode(waitValue.Interface())
				break
			}
			if isleft {
				currentNode = currentNode.left
			} else {
				currentNode = currentNode.right
			}
		}
	}
	return bt, nil
}

func NewWrapLevelAndIn(level interface{}, in interface{}) (WrapLevelAndIn, error) {
	if is, err := isCouldType(level, in); !is {
		return nil, err
	}
	return &wrapLevelAndIn {
		level: LevelSequence(&sequence{reflect.ValueOf(level)}),
		in: InSequence(&sequence{reflect.ValueOf(in)}),
	}, nil
}

// 通过中序和层次遍历结果构造二叉树
func NewWithLevelAndIn(level interface{}, in interface{}) (BinaryTree, error) {
	wrap, err := NewWrapLevelAndIn(level, in)
	if err != nil {
		return nil, err
	}
	bt := New()
	if wrap.level.Len() == 0 {
		return bt, nil
	}
	bt.root = NewNode(wrap.level.Index(0).Interface())
	for i := 1; i < wrap.level.Len(); i++ {
		waitValue := wrap.level.Index(i)
		for currentNode := bt.root; currentNode != nil; {
			currentValue := reflect.ValueOf(currentNode.data)
			isleft := wrap.in.IsLeft(&currentValue, &waitValue)
			if isleft && currentNode.left == nil {
				if currentNode.right != nil {
					return nil, InvalidLevelAndIn
				}
				currentNode.left = NewNode(waitValue.Interface())
				break
			}
			if !isleft && currentNode.right == nil {
				currentNode.right = NewNode(waitValue.Interface())
				break
			}
			if isleft {
				currentNode = currentNode.left
			} else {
				currentNode = currentNode.right
			}
		}
	}
	return bt, nil
}

func (btn BinaryTreeNode) preOrderRecursiveVisit(seq *[]interface{}) {
	if btn == nil {
		return
	}
	*seq = append(*seq, btn.data)
	btn.left.preOrderRecursiveVisit(seq)
	btn.right.preOrderRecursiveVisit(seq)
}
func (bt BinaryTree) PreOrderRecursiveVisit() (seq []interface{}){
	seq = make([]interface{}, 0)
	bt.root.preOrderRecursiveVisit(&seq)
	return seq
}

func (btn BinaryTreeNode) inOrderRecursiveVisit(seq *[]interface{}) {
	if btn == nil {
		return
	}
	btn.left.inOrderRecursiveVisit(seq)
	*seq = append(*seq, btn.data)
	btn.right.inOrderRecursiveVisit(seq)
}

func (bt BinaryTree) InOrderRecursiveVisit() (seq []interface{}) {
	seq = make([]interface{}, 0)
	bt.root.inOrderRecursiveVisit(&seq)
	return seq
}

func (btn BinaryTreeNode) postOrderRecursiveVisit(seq *[]interface{}) {
	if btn == nil {
		return
	}
	btn.left.postOrderRecursiveVisit(seq)
	btn.right.postOrderRecursiveVisit(seq)
	*seq = append(*seq, btn.data)
}

func (bt BinaryTree) PostOrderRecursiveVisit() (seq []interface{}) {
	seq = make([]interface{}, 0)
	bt.root.postOrderRecursiveVisit(&seq)
	return seq
}
// 前序遍历
// should use stack, we use slice to achieve
func (bt BinaryTree) PreOrderVisit() (seq []interface{}) {
	seq = make([]interface{}, 0)
	childs := make([]BinaryTreeNode, 0)
	if bt.root == nil {
		return seq
	}
	childs = append(childs, bt.root)
	for ; len(childs) != 0; {
		child := childs[len(childs) - 1]
		childs = childs[0:len(childs) - 1]
		seq = append(seq, child.data)
		if child.right != nil {
			childs = append(childs, child.right)
		}
		if child.left != nil {
			childs = append(childs, child.left)
		}
	}
	return seq
}

// 中序遍历
func (bt BinaryTree) InOrderVisit() (seq []interface{}) {
	seq = make([]interface{}, 0)
	childs := make([]BinaryTreeNode, 0)
	current := bt.root
	for ; len(childs) != 0 || current != nil; {
		if current != nil {
			childs = append(childs, current)
			current = current.left
		} else {
			child := childs[len(childs) - 1]
			childs = childs[0:len(childs) - 1]
			seq = append(seq, child.data)
			current = child.right
		}
	}
	return seq
}

// 后序遍历
func (bt BinaryTree) PostOrderVisit() (seq []interface{}) {
	if bt.root == nil {
		return
	}
	seq = make([]interface{}, 0)
	childs := make([]BinaryTreeNode, 0)
	var pre BinaryTreeNode = nil
	childs = append(childs, bt.root)
	for ; len(childs) != 0; {
		current := childs[len(childs) - 1]
		if pre == nil || pre.left == current || pre.right == current {
			if current.left != nil {
				childs = append(childs, current.left)
			} else if current.right != nil {
				childs = append(childs, current.right)
			}
		} else if current.left == pre {
			if current.right != nil {
				childs = append(childs, current.right)
			}
		} else {
			seq = append(seq, current.data)
			childs = childs[0:len(childs) - 1]
		}
		pre = current
	}
	return seq
}

//层序遍历
func (bt BinaryTree) LevelOrderVisit() (seq []interface{}) {
	seq = make([]interface{}, 0)
	childs := make([]BinaryTreeNode, 0)
	childs = append(childs, bt.root)
	for ; len(childs) != 0; {
		seq = append(seq, childs[0].data)
		if childs[0].left != nil {
			childs = append(childs, childs[0].left)
		}
		if childs[0].right != nil {
			childs = append(childs, childs[0].right)
		}
		childs = childs[1:]
	}
	return seq
}

// 层序遍历 从下至上 右至左
func (bt BinaryTree) LevelOrderWithBTRLVisit() (seq []interface{}) {
	//把seq当做栈处理
	seq = make([]interface{}, 0)
	if bt.root == nil {
		return seq
	}
	childs := make([]BinaryTreeNode, 0)
	childs = append(childs, bt.root)
	for ; len(childs) != 0; {
		child := childs[0]
		childs = childs[1:len(childs)]
		//从右至左
		if child.left != nil {
			childs = append(childs, child.left)
		}
		if child.right != nil {
			childs = append(childs, child.right)
		}
		//从下至上
		seq = append(append(make([]interface{}, 0), child.data), seq...)
	}
	return seq
}

// 求树高
func (bt BinaryTree) Height() int {
	if bt.root == nil {
		return 0
	}
	//经过每一层高度加1
	height := 0
	childs := make([]BinaryTreeNode, 0)
	childs = append(childs, bt.root)
	lastCount := 1
	childCount := 0
	for ; len(childs) != 0; {
		child := childs[0]
		childs = childs[1:]
		if child.left != nil {
			childs = append(childs, child.left)
			childCount++
		}
		if child.right != nil {
			childs = append(childs, child.right)
			childCount++
		}
		lastCount--
		if lastCount == 0 {
			height++
			lastCount = childCount
			childCount = 0
		}
	}
	return height
}
// 判断二叉树是否是完全二叉树
// 利用层序遍历
func (bt BinaryTree) IsComplete() (bool) {
	if bt.root == nil {
		return true
	}
	childs := make([]BinaryTreeNode, 0)
	childs = append(childs, bt.root)
	for ; len(childs) != 0; {
		child := childs[0]
		childs = childs[1:]
		if child != nil {
			childs = append(childs, child.left)
			childs = append(childs, child.right)
		} else {
			for ; len(childs) != 0; {
				rest := childs[0]
				childs = childs[1:]
				if rest != nil {
					return false
				}
			}
		}
	}
	return true
}

func (btn BinaryTreeNode) twoDegreeNodeCount() int {
	if btn == nil {
		return 0
	}
	if btn.left != nil && btn.right != nil {
		return btn.left.twoDegreeNodeCount() + btn.right.twoDegreeNodeCount() + 1
	} else {
		return btn.left.twoDegreeNodeCount() + btn.right.twoDegreeNodeCount()
	}
}

// 统计度为2的节点树
func (bt BinaryTree) TwoDegreeNodeCount() int {
	return bt.root.twoDegreeNodeCount()
}

func (btn BinaryTreeNode) swapLeftAndRight() {
	if btn == nil {
		return
	}
	btn.left.swapLeftAndRight()
	btn.right.swapLeftAndRight()
	btn.left, btn.right = btn.right, btn.left
}

// 交换二叉树左右子树
func (bt BinaryTree) SwapLeftAndRight() {
	bt.root.swapLeftAndRight()
}

func (btn BinaryTreeNode) preOrderNWithRecursive(n *int) (BinaryTreeNode) {
	if btn == nil {
		return nil
	}
	*n = *n - 1
	if *n <= 0 {
		return btn
	}
	left := btn.left.preOrderNWithRecursive(n)
	if left != nil {
		return left
	}
	right := btn.right.preOrderNWithRecursive(n)
	return right
} 
// 求先序遍历第N个节点值
func (bt BinaryTree) PreOrderN(n int) interface{} {
	if n <= 0 {
		return nil
	}
	return bt.root.preOrderNWithRecursive(&n)
}

// 某节点祖先序列
func (bt BinaryTree) AncestorsOf(x interface{}) (seq []interface{}) {
	seq = make([]interface{}, 0)
	if bt.root == nil {
		return seq
	}
	childs := append(make([]BinaryTreeNode, 0), bt.root)
	var pre BinaryTreeNode = nil
	for ; len(childs) != 0; {
		current := childs[len(childs) - 1]
		if pre == nil || pre.left == current || pre.right == current {
			if current.left != nil {
				childs = append(childs, current.left)
			} else if current.right != nil {
				childs = append(childs, current.right)
			}
		} else if current.left == pre {
			if current.right != nil {
				childs = append(childs, current.right)
			}
		} else {
			childs = childs[:len(childs) - 1]
			status, err := compare(x, current.data)
			if err != nil {
				panic(err)
			}
			if status == Eq {
				break
			}
		}
		pre = current
	}
	for _, node := range childs {
		seq = append(seq, node.data)
	}
	return seq
}

// 求两节点最近公共祖先
func (bt BinaryTree) NearCommonAncestorOf(x interface{}, y interface{}) interface{} {
	if bt.root == nil {
		return nil
	}
	childs := append(make([]BinaryTreeNode, 0), bt.root)
	var directXParent BinaryTreeNode = nil
	var directYParent BinaryTreeNode = nil
	var pre BinaryTreeNode = nil
	for ; len(childs) != 0; {
		current := childs[len(childs) - 1]
		if pre == nil || pre.left == current || pre.right == current {
			if current.left != nil {
				childs = append(childs, current.left)
			} else if current.right != nil {
				childs = append(childs, current.right)
			}
		} else if current.left == pre {
			if current.right != nil {
				childs = append(childs, current.right)
			}
		} else {
			childs = childs[:len(childs) - 1]
			if directXParent == nil {
				status, err := compare(x, current.data)
				if err != nil {
					panic(err)
				}
				if status == Eq {
					directXParent = childs[len(childs) - 1]
				}
			}
			if directYParent == nil {
				status, err := compare(y, current.data)
				if err != nil {
					panic(err)
				}
				if status == Eq {
					directYParent = childs[len(childs) - 1]
				}
			}
			if directXParent != nil && directYParent != nil {
				break
			}
		}
		pre = current
	}
	if directXParent != nil && directYParent != nil {
		if directXParent == directYParent {
			return directXParent.data
		}
		for i, v := range childs {
			if v == directXParent || v == directYParent {
				if i >= 1 {
					return childs[i - 1].data
				}
			}
		}
	}
	return nil
}

// 求二叉树最大宽度
// 利用层序遍历求得最大宽度
func (bt BinaryTree) Width() (int) {
	if bt.root == nil {
		return 0
	}
	childs := append(make([]BinaryTreeNode, 0), bt.root)
	maxLevelCount := 1
	parentCount := 1
	childCount := 0
	for ; len(childs) != 0; {
		child := childs[0]
		childs = childs[1:]
		parentCount--
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
			if childCount > maxLevelCount {
				maxLevelCount = childCount
			}
			childCount = 0
		}
	}
	return maxLevelCount
}

func isSimilar(left BinaryTreeNode, right BinaryTreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left.left == nil && left.right == nil {
		if right.left == nil && right.right == nil {
			return true
		}
		return false
	} else {
		return isSimilar(left.left, right.left) && isSimilar(left.right, right.right)
	}
}
func IsSimilar(left BinaryTree, right BinaryTree) bool {
	if left != nil && right != nil {
		return isSimilar(left.root, right.root)	
	} else {
		return left == right
	}
}