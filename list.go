package list

import "fmt"

type List struct {
	list *list
}

func (l *List) Size() int64 {
	return l.list.size()
}

func NewList(value interface{}) *List {
	element := newElement(value)
	return &List{
		list: newList(element),
	}
}

func (l *List) Head() (interface{}, error) {
	if l.list.listIsEmpty() {
		return nil, ErrListIsEmpty
	}
	return l.list.FirstElement.Value, nil
}

func (l *List) Tail() (interface{}, error) {
	// tail := l.list.FirstElement
	// for {
	// 	if tail.NextElement == nil {
	// 		break
	// 	}
	// 	tail = tail.NextElement
	// }
	if l.list.listIsEmpty() {
		return nil, ErrListIsEmpty
	}
	return l.list.LastElement.Value, nil
}

func (l *List) Add(value interface{}) {
	element := newElement(value)
	if l.list == nil {
		l.list = newList(element)
		return
	}
	l.list.add(element)
}

func (l *List) RemoveHead() error {
	return l.list.removeHead()
}

func (l *List) RemoveTail() error {
	return l.list.removeTail()
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
}

func newList(element *element) *list {
	return &list{
		FirstElement: element,
		LastElement:  element,
	}
}

func (l *list) listIsEmpty() bool {
	if l == nil {
		return true
	}

	return l.FirstElement == nil
}

func (l *list) size() int64 {
	var size int64

	element := l.FirstElement
	for {
		if element == nil {
			return size
		}
		size++
		element = element.NextElement
	}
}

func (l *list) add(newElement *element) {
	if l.listIsEmpty() {
		l.FirstElement = newElement
		l.LastElement = newElement
		return
	}

	newElement.PreviousElement = l.LastElement
	l.LastElement.NextElement = newElement
	l.LastElement = newElement
}

func (l *list) removeHead() error {
	if l.listIsEmpty() {
		return ErrListIsEmpty
	}
	headToRemove := l.FirstElement
	l.FirstElement = headToRemove.NextElement
	headToRemove = nil
	return nil
}

func (l *list) removeTail() error {
	if l.listIsEmpty() {
		return ErrListIsEmpty
	}
	tailToRemove := l.LastElement
	if tailToRemove.PreviousElement == nil {
		l.FirstElement = nil
		tailToRemove = nil
		return nil
	}
	tailToRemove.PreviousElement.NextElement = nil
	l.LastElement = tailToRemove.PreviousElement
	tailToRemove = nil
	return nil
}

func (l *list) remove(elementToRemove *element) error {
	if elementToRemove == l.FirstElement {
		return l.removeHead()
	}
	if elementToRemove == l.LastElement {
		return l.removeTail()
	}

	elementToRemove.PreviousElement.NextElement = elementToRemove.NextElement
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

func (t *element) PrintElement() {
	fmt.Println(t.Value)
}
