package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewList(t *testing.T) {
	list := New(1)

	assert.Equal(t, list.value, 1)
}

func TestListAdd(t *testing.T) {
	list := New(1)

	list.Add(2)
	assert.Equal(t, list.next.value, 2)
	assert.Equal(t, list.next.prev, list)

	list.Add(3)
	next := list.next
	assert.Equal(t, next.next.value, 3)
	assert.Equal(t, next.next.prev, next)
}

func TestListRemove(t *testing.T) {
	list := New(1)
	list.Add(2)
	list.Add(3)

	list.Remove(2)
	assert.Equal(t, list.next.value, 3)
	assert.Equal(t, list.next.prev, list)
}

func TestListContains(t *testing.T) {
	list := New(1)
	list.Add(2)
	list.Add(3)

	assert.True(t, list.Contains(2))
	assert.False(t, list.Contains(4))
}

func TestListSeek(t *testing.T) {
	list := New(1)
	list.Add(2)
	list.Add(3)

	assert.Equal(t, list.Seek(2).value, 2)
	assert.Equal(t, list.Seek(3).value, 3)
	assert.Nil(t, list.Seek(4))
}

func TestListLength(t *testing.T) {
	list := New(1)
	list.Add(2)
	list.Add(3)

	assert.Equal(t, list.Length(), 3)

	list.Add(4)
	assert.Equal(t, list.Length(), 4)
}
