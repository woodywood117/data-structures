package stack

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewStack(t *testing.T) {
	s, err := New[int](10)
	assert.Nil(t, err)
	assert.Equal(t, s.top, -1)
	assert.Equal(t, len(s.data), 10)
}

func TestNewStackBadCapacity(t *testing.T) {
	s, err := New[int](0)
	assert.ErrorIs(t, err, ErrInvalidCapacity)
	assert.Nil(t, s)

	s, err = New[int](-1)
	assert.ErrorIs(t, err, ErrInvalidCapacity)
	assert.Nil(t, s)
}

func TestStackPush(t *testing.T) {
	s, _ := New[int](10)

	value := 1
	s.Push(&value)
	assert.Equal(t, s.top, 0)
	assert.Equal(t, *s.data[0], value)
}

func TestStackOverflow(t *testing.T) {
	s, _ := New[int](1)

	value := 1
	s.Push(&value)
	assert.ErrorIs(t, s.Push(&value), ErrStackOverflow)
}

func TestStackPop(t *testing.T) {
	s, _ := New[int](1)

	value := 1
	s.Push(&value)
	pop, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, *pop, 1)
	assert.Equal(t, s.top, -1)
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
	assert.Equal(t, *pop, 2)
	assert.Equal(t, s.top, -1)
}

func TestStackEmptyPop(t *testing.T) {
	s, _ := New[int](1)

	_, err := s.Pop()
	assert.ErrorIs(t, err, ErrEmptyStack)
}

func TestStackPeek(t *testing.T) {
	s, _ := New[int](1)

	value := 1
	s.Push(&value)
	peek, err := s.Peek()
	assert.Nil(t, err)
	assert.Equal(t, *peek, 1)
	assert.Equal(t, s.top, 0)
}

func TestStackEmptyPeek(t *testing.T) {
	s, _ := New[int](1)

	_, err := s.Peek()
	assert.ErrorIs(t, err, ErrEmptyStack)
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
