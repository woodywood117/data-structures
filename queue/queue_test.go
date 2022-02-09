package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	q := New[int]()
	assert.NotNil(t, q.data)
}

func TestQueueDequeue(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, err := q.Dequeue()
	assert.Equal(t, 1, v)
	assert.Nil(t, err)

	v, err = q.Dequeue()
	assert.Equal(t, 2, v)
	assert.Nil(t, err)

	v, err = q.Dequeue()
	assert.Equal(t, 3, v)
	assert.Nil(t, err)

	_, err = q.Dequeue()
	assert.ErrorIs(t, ErrEmptyQueue, err)
}
