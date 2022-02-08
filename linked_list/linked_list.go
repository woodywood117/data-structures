package linked_list

import "reflect"

type List[T any] struct {
	prev  *List[T]
	value T
	next  *List[T]
}

func New[T any](initial_value T) *List[T] {
	return &List[T]{value: initial_value}
}

func (l *List[T]) Add(value T) {
	if l.next == nil {
		l.next = &List[T]{value: value}
		l.next.prev = l
	} else {
		l.next.Add(value)
	}
}

func (l *List[T]) Remove(value T) {
	if reflect.DeepEqual(l.value, value) {
		if l.prev != nil {
			l.prev.next = l.next
		}
		if l.next != nil {
			l.next.prev = l.prev
		}
	} else {
		l.next.Remove(value)
	}
}

func (l *List[T]) Contains(value T) bool {
	if reflect.DeepEqual(l.value, value) {
		return true
	} else if l.next == nil {
		return false
	} else {
		return l.next.Contains(value)
	}
}

func (l *List[T]) Seek(value T) *List[T] {
	if reflect.DeepEqual(l.value, value) {
		return l
	} else if l.next == nil {
		return nil
	} else {
		return l.next.Seek(value)
	}
}

func (l *List[T]) Length() int {
	if l.next == nil {
		return 1
	} else {
		return 1 + l.next.Length()
	}
}
