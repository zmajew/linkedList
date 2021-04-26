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

func (l *List) Head() interface{} {
	return l.list.FirstElement.Value
}

func (l *List) Tail() interface{} {
	tail := l.list.FirstElement
	for {
		if tail.NextElement == nil {
			break
		}
		tail = tail.NextElement
	}

	return tail.Value
}

func (l *List) Add(value interface{}) {
	element := newElement(value)
	if l.list == nil {
		l.list = newList(element)
		return
	}
	l.list.add(element)
}

type list struct {
	FirstElement *element
}

func newList(element *element) *list {
	return &list{
		FirstElement: element,
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
	if l.size() == 0 {
		l.FirstElement = newElement
		return
	}
	element := l.FirstElement

	for {
		if element.NextElement == nil {
			element.NextElement = newElement
			break
		}
		element = element.NextElement
	}
	element.NextElement = newElement
}

func (l *list) remove(elementToRemove *element) error {
	if l.listIsEmpty() {
		return fmt.Errorf("Error: List is empty")
	}

	if elementToRemove == l.FirstElement {
		l.FirstElement = l.FirstElement.NextElement
		return nil
	}

	element := l.FirstElement
	for {
		if element.NextElement == elementToRemove {
			element.NextElement = elementToRemove.NextElement
			elementToRemove = nil
			return nil
		}
		element = element.NextElement
		if element == elementToRemove {
			element.NextElement = elementToRemove.NextElement
			elementToRemove = nil
			return nil
		}
		if element.NextElement == nil {
			return fmt.Errorf("List does not containt element to be removed")
		}
	}
}

func (l *List) RemoveByOrderNo(order int64) error {
	if order < 0 {
		return fmt.Errorf("Error: element order can not be negative integer")
	}
	if l.list.listIsEmpty() {
		return fmt.Errorf("Error: List is empty")
	}

	size := l.list.size()
	if size < order+1 {
		return fmt.Errorf("Error: Element's order %d out of the list size %d", order, size)
	}

	element := l.list.FirstElement
	var i int64
	for {
		if i == order {
			l.list.remove(element)
			return nil
		}
		element = element.NextElement
		i++
	}
}

type element struct {
	Value       interface{}
	NextElement *element
}

func newElement(value interface{}) *element {
	return &element{Value: value}
}

func (t *element) PrintElement() {
	fmt.Println(t.Value)
}

// func main() {
// 	jedan := &Element{
// 		Value: "jedan",
// 	}
// 	list := NewList(jedan)
// 	list.PrintList()
// 	fmt.Println("velicina", list.Size())

// 	dva := &Element{
// 		Value: "dva",
// 	}
// 	list.Add(dva)

// 	tri := &Element{
// 		Value: "tri",
// 	}
// 	list.Add(tri)

// 	list.PrintList()
// 	fmt.Println("velicina", list.Size())

// 	//tail := list.Tail()

// 	// err := list.Remove(dva)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	err := list.RemoveByOrderNo(0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	list.PrintList()
// 	fmt.Println("velicina", list.Size())
// }
