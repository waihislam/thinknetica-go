// Реализуация двусвязного списка вместе с базовыми операциями.
package list

import "fmt"

// List - двусвязный список.
type List struct {
	root *Elem
}

// Elem - элемент списка.
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

// New создаёт список и возвращает указатель на него.
func New() *List {
	var l List
	l.root = &Elem{}
	l.root.next = l.root
	l.root.prev = l.root
	return &l
}

// Push вставляет элемент в начало списка.
func (l *List) Push(e Elem) *Elem {
	e.prev = l.root
	e.next = l.root.next
	l.root.next = &e
	if e.next != l.root {
		e.next.prev = &e
	}
	return &e
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (l *List) String() string {
	el := l.root.next
	var s string
	for el != l.root {
		s += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

// Pop удаляет первый элемент списка.
func (l *List) Pop() *List {
	if l.root.next == l.root {
		// The list is empty
		return l
	}
	first := l.root.next
	l.root.next = first.next
	first.next.prev = l.root
	first.next = nil
	first.prev = nil
	return l
}

// Reverse разворачивает список.
func (l *List) Reverse() *List {
	if l.root.next == l.root || l.root.next.next == l.root {
		return l
	}

	current := l.root.next

	for current.next != l.root {
		next := *current
		current.next = next.prev
		current.prev = next.next
		current = next.next
	}
	next := *current
	l.root.next = current
	current.next, current.prev = next.prev, l.root

	return l
}
