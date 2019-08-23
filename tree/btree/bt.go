package btree

import (
	"errors"
	"reflect"
	//  "fmt"
)

var (
	UnExpectedType = errors.New("UnExpectedType")
	TypeError = errors.New("TypeError")
	InvalidSequence = errors.New("InvalidSequence")
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
		for currentNode := bt.root; currentNode != nil; {
			value := reflect.ValueOf(currentNode.data)
			current := wrap.pre.Index(i)
			isleft := wrap.in.IsLeft(&value, &current)
			if isleft && currentNode.left == nil {
				currentNode.left = NewNode(wrap.pre.Index(i).Interface())
				break
			}
			if !isleft && currentNode.right == nil {
				currentNode.right = NewNode(wrap.pre.Index(i).Interface())
				break
			}
			if isleft {
				currentNode = currentNode.left
				continue
			} else {
				currentNode = currentNode.right
				continue
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
		for currentNode := bt.root; currentNode != nil; {
			value := wrap.post.Index(i)
			current := reflect.ValueOf(currentNode.data)
			isleft := wrap.in.IsLeft(&current, &value)
			if isleft && currentNode.left == nil {
				currentNode.left = NewNode(value.Interface())
				break
			}
			if !isleft && currentNode.right == nil {
				currentNode.right = NewNode(value.Interface())	
				break
			}
			if isleft {
				currentNode = currentNode.left
				continue
			} else {
				currentNode = currentNode.right
				continue
			}
		}
	}
	return bt, nil
}

// 通过前序和层次遍历结果构造二叉树
func NewWithPreAndLevel(pre interface{}, level interface{}) {

}

// 通过中序和层次遍历结果构造二叉树
func NewWithInAndLevel(in interface{}, level interface{}) {

}
// 通过后序和层次遍历结果构造二叉树
func NewWithPostAndLevel(post interface{}, level interface{}) {

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
func (bt BinaryTree) Display() {

}