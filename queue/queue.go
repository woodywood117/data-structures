package queue

import (
	"errors"
	"github.com/woodywood117/data-structures/linked_list"
)

var ErrEmptyQueue = errors.New("queue is empty")

type Queue[T any] struct {
	data *linked_list.List[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{linked_list.New[T]()}
}

func (q *Queue[T]) Enqueue(value T) {
	q.data.Add(value)
}

func (q *Queue[T]) Dequeue() (T, error) {
	value, err := q.data.PopHead()
	if err != nil {
		return *new(T), ErrEmptyQueue
	}
	return value, nil
}
