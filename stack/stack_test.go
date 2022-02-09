package stack

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewStack(t *testing.T) {
	s, err := New[int](10)
	assert.Nil(t, err)
	assert.Equal(t, -1, s.top)
	assert.Equal(t, 10, len(s.data))
}

func TestNewStackBadCapacity(t *testing.T) {
	s, err := New[int](0)
	assert.ErrorIs(t, ErrInvalidCapacity, err)
	assert.Nil(t, s)

	s, err = New[int](-1)
	assert.ErrorIs(t, ErrInvalidCapacity, err)
	assert.Nil(t, s)
}

func TestStackPush(t *testing.T) {
	s, _ := New[int](10)

	value := 1
	s.Push(&value)
	assert.Equal(t, 0, s.top)
	assert.Equal(t, value, *s.data[0])
}

func TestStackOverflow(t *testing.T) {
	s, _ := New[int](1)

	value := 1
	s.Push(&value)
	assert.ErrorIs(t, ErrStackOverflow, s.Push(&value))
}

func TestStackPop(t *testing.T) {
	s, _ := New[int](1)

	value := 1
	s.Push(&value)
	pop, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 1, *pop)
	assert.Equal(t, -1, s.top)
}

func TestStackMultiPop(t *testing.T) {
	s, _ := New[int](1)

	value := 1
	s.Push(&value)
	_, _ = s.Pop()

	value2 := 2
	s.Push(&value2)
	pop, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 2, *pop)
	assert.Equal(t, -1, s.top)
}

func TestStackEmptyPop(t *testing.T) {
	s, _ := New[int](1)

	_, err := s.Pop()
	assert.ErrorIs(t, ErrEmptyStack, err)
}

func TestStackPeek(t *testing.T) {
	s, _ := New[int](1)

	value := 1
	s.Push(&value)
	peek, err := s.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 1, *peek)
	assert.Equal(t, 0, s.top)
}

func TestStackEmptyPeek(t *testing.T) {
	s, _ := New[int](1)

	_, err := s.Peek()
	assert.ErrorIs(t, ErrEmptyStack, err)
}

func TestStackEmpty(t *testing.T) {
	s, _ := New[int](1)

	assert.True(t, s.IsEmpty())
}

func TestStackNotEmpty(t *testing.T) {
	s, _ := New[int](1)

	value := 1
	s.Push(&value)
	assert.False(t, s.IsEmpty())
}

func TestStackFull(t *testing.T) {
	s, _ := New[int](1)

	value := 1
	s.Push(&value)
	assert.True(t, s.IsFull())
}

func TestStackNotFull(t *testing.T) {
	s, _ := New[int](1)

	assert.False(t, s.IsFull())
}
