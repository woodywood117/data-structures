package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewList(t *testing.T) {
	list := New(1)

	assert.Equal(t, 1, list.head.value)
}

func TestListAdd(t *testing.T) {
	list := New(1)

	list.Add(2)
	assert.Equal(t, 2, list.head.next.value)
	assert.Equal(t, list.head, list.head.next.prev)

	list.Add(3)
	next := list.head.next
	assert.Equal(t, 3, next.next.value)
	assert.Equal(t, next, next.next.prev)
}

func TestListRemove(t *testing.T) {
	list := New(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	list.Remove(3)
	assert.Equal(t, 4, list.head.next.next.value)
	assert.Equal(t, list.head, list.head.next.prev)

	list.Remove(2)
	assert.Equal(t, 4, list.head.next.value)
	assert.Equal(t, list.head, list.head.next.prev)

	list.Remove(1)
	assert.Equal(t, 4, list.head.value)
	assert.Nil(t, list.head.prev)
}

func TestRecreateHead(t *testing.T) {
	list := New(1)
	list.Remove(1)
	list.Remove(1)

	list.Add(2)
	assert.Equal(t, 2, list.head.value)
}

func TestListContains(t *testing.T) {
	list := New(1)
	list.Add(2)
	list.Add(3)

	assert.True(t, list.Contains(2))
	assert.False(t, list.Contains(4))
}

func TestListLength(t *testing.T) {
	list := New(1)
	list.Add(2)
	list.Add(3)

	assert.Equal(t, 3, list.Length())

	list.Add(4)
	assert.Equal(t, 4, list.Length())
}
