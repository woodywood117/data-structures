package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cmp = compare_int

func compare_int(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func TestNew(t *testing.T) {
	b := New[int](&cmp)
	assert.Nil(t, b.root)
	assert.Equal(t, &cmp, b.cmp)
}

func TestBSTInsert(t *testing.T) {
	b := New[int](&cmp)
	assert.Nil(t, b.Insert(5))
	assert.Nil(t, b.Insert(3))
	assert.Nil(t, b.Insert(2))
	assert.Nil(t, b.Insert(4))
	assert.Nil(t, b.Insert(7))
	assert.Nil(t, b.Insert(6))
	assert.Nil(t, b.Insert(8))
	assert.ErrorIs(t, ErrRepeatedValue, b.Insert(5))

	assert.Equal(t, 5, b.root.value)
	assert.Equal(t, 3, b.root.left.value)
	assert.Equal(t, 2, b.root.left.left.value)
	assert.Equal(t, 4, b.root.left.right.value)
	assert.Equal(t, 7, b.root.right.value)
	assert.Equal(t, 6, b.root.right.left.value)
	assert.Equal(t, 8, b.root.right.right.value)
}

func TestBSTRemoveNoChildren(t *testing.T) {
	b := New[int](&cmp)
	_ = b.Insert(2)
	_ = b.Insert(1)
	_ = b.Insert(3)
	b.Remove(1)
	b.Remove(3)
	assert.Nil(t, b.root.left)
	assert.Nil(t, b.root.right)
}

func TestBSTRemoveOneChild(t *testing.T) {
	// Left outer child
	b := New[int](&cmp)
	_ = b.Insert(3)
	_ = b.Insert(2)
	_ = b.Insert(1)
	b.Remove(2)
	assert.Equal(t, 1, b.root.left.value)

	// Left inner child
	b = New[int](&cmp)
	_ = b.Insert(3)
	_ = b.Insert(1)
	_ = b.Insert(2)
	b.Remove(1)
	assert.Equal(t, 2, b.root.left.value)

	// Right outer child
	b = New[int](&cmp)
	_ = b.Insert(1)
	_ = b.Insert(2)
	_ = b.Insert(3)
	b.Remove(2)
	assert.Equal(t, 3, b.root.right.value)

	// Right inner child
	b = New[int](&cmp)
	_ = b.Insert(1)
	_ = b.Insert(3)
	_ = b.Insert(2)
	b.Remove(3)
	assert.Equal(t, 2, b.root.right.value)
}

func TestBSTRemoveTwoChildren(t *testing.T) {
	// Right, no recursive children
	b := New[int](&cmp)
	_ = b.Insert(1)
	_ = b.Insert(3)
	_ = b.Insert(2)
	_ = b.Insert(4)
	b.Remove(3)
	assert.Equal(t, 4, b.root.right.value)
	assert.Equal(t, 2, b.root.right.left.value)
	assert.Nil(t, b.root.right.right)

	// Left, no recursive children
	b = New[int](&cmp)
	_ = b.Insert(4)
	_ = b.Insert(2)
	_ = b.Insert(1)
	_ = b.Insert(3)
	b.Remove(2)
	assert.Equal(t, 3, b.root.left.value)
	assert.Equal(t, 1, b.root.left.left.value)
	assert.Nil(t, b.root.left.right)

	// Recursive children
	b = New[int](&cmp)
	_ = b.Insert(1)
	_ = b.Insert(3)
	_ = b.Insert(2)
	_ = b.Insert(5)
	_ = b.Insert(4)
	_ = b.Insert(6)
	b.Remove(3)
	assert.Equal(t, 4, b.root.right.value)
	assert.Equal(t, 2, b.root.right.left.value)
	assert.Equal(t, 5, b.root.right.right.value)
	assert.Equal(t, 6, b.root.right.right.right.value)

	// Recursive children, recurse has right child
	b = New[int](&cmp)
	_ = b.Insert(1)
	_ = b.Insert(3)
	_ = b.Insert(2)
	_ = b.Insert(6)
	_ = b.Insert(4)
	_ = b.Insert(5)
	_ = b.Insert(7)
	b.Remove(3)
	assert.Equal(t, 4, b.root.right.value)
	assert.Equal(t, 2, b.root.right.left.value)
	assert.Equal(t, 6, b.root.right.right.value)
	assert.Equal(t, 5, b.root.right.right.left.value)
	assert.Equal(t, 7, b.root.right.right.right.value)
}

func TestBSTContains(t *testing.T) {
	b := New[int](&cmp)
	_ = b.Insert(1)
	_ = b.Insert(3)
	_ = b.Insert(2)
	_ = b.Insert(4)
	assert.True(t, b.Contains(1))
	assert.True(t, b.Contains(3))
	assert.True(t, b.Contains(2))
	assert.True(t, b.Contains(4))
	assert.False(t, b.Contains(5))
}
