package bst

import "errors"

var ErrRepeatedValue = errors.New("repeated value")

type BST[T any] struct {
	root *node[T]
	cmp  *func(a, b T) int // returns -1 if a < b, 0 if a == b, 1 if a > b
}

type node[T any] struct {
	parent *node[T]
	value  T
	left   *node[T]
	right  *node[T]
}

func New[T any](cmp *func(a, b T) int) *BST[T] {
	return &BST[T]{nil, cmp}
}

func (bst *BST[T]) Insert(value T) error {
	if bst.root == nil {
		bst.root = &node[T]{nil, value, nil, nil}
		return nil
	}
	n := bst.root
	for {
		if (*bst.cmp)(value, n.value) == -1 {
			if n.left == nil {
				n.left = &node[T]{n, value, nil, nil}
				return nil
			}
			n = n.left
		} else if (*bst.cmp)(value, n.value) == 1 {
			if n.right == nil {
				n.right = &node[T]{n, value, nil, nil}
				return nil
			}
			n = n.right
		} else {
			return ErrRepeatedValue
		}
	}
}

func (bst *BST[T]) Remove(value T) {
	n := bst.root
	for n != nil {
		if (*bst.cmp)(value, n.value) == -1 {
			n = n.left
		} else if (*bst.cmp)(value, n.value) == 1 {
			n = n.right
		} else {
			if n.left == nil && n.right == nil { // node has no children
				if n.parent.left == n {
					n.parent.left = nil
					return
				} else {
					n.parent.right = nil
					return
				}
			} else if n.left != nil && n.right != nil { // node has left and right children
				smallest := n.right
				for smallest.left != nil {
					smallest = smallest.left
				}
				bst.Remove(smallest.value)
				smallest.parent = n.parent
				smallest.left = n.left
				smallest.right = n.right
				if n.parent.left == n {
					n.parent.left = smallest
					return
				} else {
					n.parent.right = smallest
					return
				}
			} else if n.left != nil { // node has left child
				if n.parent.left == n {
					n.left.parent = n.parent
					n.parent.left = n.left
					return
				} else {
					n.left.parent = n.parent
					n.parent.right = n.left
					return
				}
			} else if n.right != nil { // node has right child
				if n.parent.left == n {
					n.right.parent = n.parent
					n.parent.left = n.right
					return
				} else {
					n.right.parent = n.parent
					n.parent.right = n.right
					return
				}
			}
		}
	}
}

func (bst *BST[T]) Contains(value T) bool {
	n := bst.root
	for n != nil {
		if (*bst.cmp)(value, n.value) == -1 {
			n = n.left
		} else if (*bst.cmp)(value, n.value) == 1 {
			n = n.right
		} else {
			return true
		}
	}
	return false
}
