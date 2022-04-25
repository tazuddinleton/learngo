package linkedlist

import "fmt"

// List[T any] describes the structure of a singly linked list
type List[T any] struct {
	head T
	tail *List[T]
}

// New[T any] creates a new singly linked list
func New[T any](head T, tail *List[T]) *List[T] {
	return &List[T]{head, tail}
}

// Apply allows (1,2,3,4) syntax to create linked list with items 1,2,3,4
func Apply[T any](ts ...T) *List[T] {
	if len(ts) == 0 {
		return nil
	}
	return New(ts[0], Apply((ts[1:])...))
}

func (l *List[T]) getTail() *List[T] {
	if l == nil {
		return nil
	}
	return l.tail
}

func (l *List[T]) getHead() T {
	return l.head
}

// OO style set head, mutating in place
func (l *List[T]) setHead(h T) {
	l.head = h
}

// SetHead sets given head
func SetHead[T any](l *List[T], h T) *List[T] {
	if l == nil {
		return nil
	}
	return New(h, l.tail)
}

// todo: Find out why mutating in place doesn't mutate original list
// OO Style drop, mutating in place
func (l *List[T]) drop(n int) {
	fmt.Println(&l)
	for i := 0; i < n; i++ {
		l = l.tail
	}
}

// Drop deletes first n items
func Drop[T any](l *List[T], n int) *List[T] {
	if n <= 0 {
		return l
	}
	return Drop(l.tail, n-1)
}

// DropWhile deletes items while a condition is true
func DropWhile[T any](l *List[T], f func(t T) bool) *List[T] {
	if f(l.head) {
		return DropWhile(l.tail, f)
	}
	return l
}

// Append appends a list to another list
func Append[T any](l1 *List[T], l2 *List[T]) *List[T] {
	if l1 == nil {
		return l2
	}

	return New(l1.head, Append(l1.tail, l2))
}

// Last[T any] returns the last element
func Last[T any](l *List[T]) T {
	if l.tail != nil {
		return Last(l.tail)
	}
	return l.head
}

// Droplast[T any] deletes the last element
func DropLast[T any](l *List[T]) *List[T] {
	var loop func(*List[T], *List[T]) *List[T]
	loop = func(l1, l2 *List[T]) *List[T] {
		if l2 == nil || l2.tail == nil {
			return l1
		}
		return loop(Append(l1, New(l2.head, nil)), l2.tail)
	}
	return loop(New(l.head, nil), l.tail)
}
