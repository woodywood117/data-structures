package linked_list

import "reflect"

type node[T any] struct {
	prev  *node[T]
	value T
	next  *node[T]
}

type List[T any] struct {
	head *node[T]
}

func New[T any](initial_value T) *List[T] {
	head_node := &node[T]{value: initial_value}
	return &List[T]{head: head_node}
}

func (n *node[T]) add(value T) {
	if n.next == nil {
		n.next = &node[T]{value: value}
		n.next.prev = n
	} else {
		n.next.add(value)
	}
}

func (l *List[T]) Add(value T) {
	if l.head == nil {
		l.head = &node[T]{value: value}
	} else {
		l.head.add(value)
	}
}

func (n *node[T]) remove(value T) {
	if reflect.DeepEqual(n.value, value) {
		if n.prev != nil {
			n.prev.next = n.next
		}
		if n.next != nil {
			n.next.prev = n.prev
		}
	} else {
		n.next.remove(value)
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
		}
	} else if l.head.next != nil {
		l.head.next.remove(value)
	}
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
