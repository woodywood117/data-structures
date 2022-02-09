package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewList(t *testing.T) {
	list := New[int]()

	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
}

func TestListAdd(t *testing.T) {
	list := New[int]()

	list.Add(1)
	list.Add(2)
	assert.Equal(t, 2, list.head.next.value)
	assert.Equal(t, 2, list.tail.value)
	assert.Equal(t, list.head, list.head.next.prev)

	list.Add(3)
	next := list.head.next
	assert.Equal(t, 3, next.next.value)
	assert.Equal(t, 3, list.tail.value)
	assert.Equal(t, next, next.next.prev)
}

func TestListRemove(t *testing.T) {
	list := New[int]()
	list.Add(1)
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

	list.Remove(4)
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
}

func TestListRemoveTail(t *testing.T) {
	list := New[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	list.Remove(4)
	assert.Equal(t, 3, list.tail.value)

	list.Remove(2)
	assert.Equal(t, 3, list.tail.value)
}

func TestListPopHead(t *testing.T) {
	list := New[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	head_value, err := list.PopHead()
	assert.Nil(t, err)
	assert.Equal(t, 1, head_value)
	assert.Equal(t, 2, list.head.value)
	assert.Equal(t, list.head, list.head.next.prev)

	list = New[int]()
	list.Add(1)
	list.PopHead()
	_, err = list.PopHead()
	assert.ErrorIs(t, ErrEmptyList, err)
}

func TestListPopTail(t *testing.T) {
	list := New[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	tail_value, err := list.PopTail()
	assert.Nil(t, err)
	assert.Equal(t, 3, tail_value)
	assert.Equal(t, 2, list.tail.value)
	assert.Equal(t, list.tail, list.tail.prev.next)

	list2 := New[int]()
	list2.Add(1)
	list2.PopTail()
	assert.Nil(t, list2.head)
	_, err = list2.PopTail()
	assert.ErrorIs(t, ErrEmptyList, err)
}

func TestRecreateHead(t *testing.T) {
	list := New[int]()
	list.Remove(1)

	list.Add(2)
	assert.Equal(t, 2, list.head.value)
}

func TestListContains(t *testing.T) {
	list := New[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	assert.True(t, list.Contains(2))
	assert.False(t, list.Contains(4))
}

func TestListLength(t *testing.T) {
	list := New[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	assert.Equal(t, 3, list.Length())

	list.Add(4)
	assert.Equal(t, 4, list.Length())
}
