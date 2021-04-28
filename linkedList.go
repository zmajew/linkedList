package linkedList

import (
	"sync"
)

type List struct {
	list *list
}

func (l *List) Size() int64 {
	return l.list.size()
}

func NewList() *List {
	return &List{
		list: newList(),
	}
}

func (l *List) Head() interface{} {
	if l.list.listIsEmpty() {
		return nil
	}
	return l.list.FirstElement.Value
}

func (l *List) Tail() interface{} {
	if l.list.listIsEmpty() {
		return nil
	}
	return l.list.LastElement.Value
}

func (l *List) Add(value interface{}) {
	element := newElement(value)
	if l.list == nil {
		l.list = newList()
	}
	l.list.add(element)
}

func (l *List) RemoveHead() error {
	return l.list.removeHead()
}

func (l *List) RemoveTail() error {
	return l.list.removeTail()
}

func (l *List) ListIsEmpty() bool {
	return l.list.listIsEmpty()
}

func (l *List) RemoveByOrderInList(order int64) error {
	if order < 0 {
		return ErrNegativeOrder
	}
	if l.list.listIsEmpty() {
		return ErrListIsEmpty
	}
	if l.list.size() < order+1 {
		return ErrOrderOutOfSize
	}

	element := l.list.FirstElement
	var i int64
	for {
		if i == order {
			return l.list.remove(element)
		}
		element = element.NextElement
		i++
	}
}

type list struct {
	FirstElement *element
	LastElement  *element
	mu           sync.Mutex
}

func newList() *list {
	return &list{}
}

func (l *list) listIsEmpty() bool {
	if l == nil {
		return true
	}

	return l.FirstElement == nil
}

func (l *list) size() int64 {
	var size int64
	l.mu.Lock()
	element := l.FirstElement
	for {
		if element == nil {
			l.mu.Unlock()
			return size
		}
		size++
		element = element.NextElement
	}
}

func (l *list) add(newElement *element) {
	l.mu.Lock()
	if l.listIsEmpty() {
		l.FirstElement = newElement
		l.LastElement = newElement
		l.mu.Unlock()
		return
	}

	newElement.PreviousElement = l.LastElement
	l.LastElement.NextElement = newElement
	l.LastElement = newElement
	l.mu.Unlock()
}

func (l *list) removeHead() error {
	l.mu.Lock()
	if l.listIsEmpty() {
		l.mu.Unlock()
		return ErrListIsEmpty
	}
	headToRemove := l.FirstElement
	l.FirstElement = headToRemove.NextElement
	headToRemove = nil
	l.mu.Unlock()
	return nil
}

func (l *list) removeTail() error {
	l.mu.Lock()
	if l.listIsEmpty() {
		l.mu.Unlock()
		return ErrListIsEmpty
	}
	tailToRemove := l.LastElement
	if tailToRemove.PreviousElement == nil {
		l.FirstElement = nil
		tailToRemove = nil
		l.mu.Unlock()
		return nil
	}
	tailToRemove.PreviousElement.NextElement = nil
	l.LastElement = tailToRemove.PreviousElement
	tailToRemove = nil
	l.mu.Unlock()
	return nil
}

func (l *list) remove(elementToRemove *element) error {
	l.mu.Lock()
	if elementToRemove == l.FirstElement {
		l.mu.Unlock()
		return l.removeHead()
	}
	if elementToRemove == l.LastElement {
		l.mu.Unlock()
		return l.removeTail()
	}

	elementToRemove.PreviousElement.NextElement = elementToRemove.NextElement
	l.mu.Unlock()
	return nil
}

type element struct {
	Value           interface{}
	PreviousElement *element
	NextElement     *element
}

func newElement(value interface{}) *element {
	return &element{Value: value}
}
