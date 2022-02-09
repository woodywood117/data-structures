package linked_list

import (
	"errors"
	"reflect"
)

var ErrEmptyList = errors.New("empty list")

type node[T any] struct {
	prev  *node[T]
	value T
	next  *node[T]
}

type List[T any] struct {
	head *node[T]
	tail *node[T]
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (n *node[T]) add(value T) *node[T] {
	if n.next == nil {
		n.next = &node[T]{value: value}
		n.next.prev = n
		return n.next
	} else {
		return n.next.add(value)
	}
}

func (l *List[T]) Add(value T) {
	if l.head == nil {
		l.head = &node[T]{value: value}
	} else {
		l.tail = l.head.add(value)
	}
}

func (n *node[T]) remove(value T) *node[T] {
	if reflect.DeepEqual(n.value, value) {
		if n.prev != nil {
			n.prev.next = n.next
		}
		if n.next != nil {
			n.next.prev = n.prev
		}
		return n
	} else {
		return n.next.remove(value)
	}
}

func (l *List[T]) Remove(value T) {
	if l.head == nil {
		return
	}

	if reflect.DeepEqual(l.head.value, value) {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil
		}
	} else if l.head.next != nil {
		removed := l.head.next.remove(value)
		if removed == l.tail {
			l.tail = removed.prev
		}
	}
}

func (l *List[T]) PopHead() (T, error) {
	if l.head == nil {
		return *new(T), ErrEmptyList
	}

	value := l.head.value
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}

	return value, nil
}

func (l *List[T]) PopTail() (T, error) {
	if l.tail == nil {
		return *new(T), ErrEmptyList
	}

	value := l.tail.value
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}

	return value, nil
}

func (n *node[T]) contains(value T) bool {
	if reflect.DeepEqual(n.value, value) {
		return true
	} else if n.next == nil {
		return false
	} else {
		return n.next.contains(value)
	}
}

func (l *List[T]) Contains(value T) bool {
	return l.head.contains(value)
}

func (n *node[T]) length() int {
	if n.next == nil {
		return 1
	} else {
		return 1 + n.next.length()
	}
}

func (l *List[T]) Length() int {
	return l.head.length()
}
