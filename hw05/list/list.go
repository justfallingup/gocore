// Реализация двусвязного списка вместе с базовыми операциями.
package list

import "fmt"

// List - двусвязный список.
type List struct {
	root *Elem
}

// Elem - элемент списка.
type Elem struct {
	Val interface{}
	next, prev *Elem
}

// New создаёт список и возвращает указатель на него.
func New() *List{
	var l List
	l.root = &Elem{}
	l.root.prev = l.root
	l.root.next = l.root
	return &l
}

// Push вставляет элемент в начало списка.
func (l *List) Push(e Elem) *List {
	e.prev = l.root
	e.next = l.root.next
	l.root.next = &e
	if e.next != l.root {
		e.next.prev = &e
	}
	return l
}

// Pop удаляет первый элемент списка.
func (l *List) Pop() *Elem {
	el := l.root.next
	if el == l.root {
		return &Elem{}
	}
	l.root.next = el.next
	el.next.prev = l.root
	el.prev, el.next = nil, nil
	return el
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

// Reverse разворачивает список.
func (l *List) Reverse() *List {
	el := l.root.next
	if el == l.root {
		return l
	}
	for el != l.root {
		el.prev, el.next = el.next, el.prev
		if el.prev == l.root {
			break
		}
		el = el.prev
	}
	l.root.next = el
	return l
}