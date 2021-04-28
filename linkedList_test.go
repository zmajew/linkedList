package linkedList

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/zmajew/linkedList"
)

var (
	testValueOne   string = "one"
	testValueTwo   string = "two"
	testValueThree string = "three"
)

// Queue example
type Queue interface {
	GetHead() (interface{}, error)
	Add(value interface{})
	Size() int64
}

func NewQueue() Queue {
	return linkedList.NewList()
}

func ExampleQueue() {
	// Inicialse queue
	q := NewQueue()
	// Add two element on the queue
	q.Add("first element")
	q.Add("second element")

	//Get and remove head
	element1Interface, err := q.GetHead()
	if err != nil {
		fmt.Println(err)
		return
	}
	element1Value, _ := element1Interface.(string)
	fmt.Println(element1Value)
	// Add new element
	q.Add("third element")

	element2Interface, err := q.GetHead()
	if err != nil {
		fmt.Println(err)
		return
	}
	element2Value, _ := element2Interface.(string)
	fmt.Println(element2Value)

	fmt.Println(q.Size())

	element3Interface, err := q.GetHead()
	if err != nil {
		fmt.Println(err)
		return
	}
	element3Value, _ := element3Interface.(string)
	fmt.Println(element3Value)

	_, err = q.GetHead()
	fmt.Println(err == linkedList.ErrListIsEmpty)
	fmt.Println(q.Size())
	// Output:
	// first element
	// second element
	// 1
	// third element
	// true
	// 0
}

//Test
func TestMain(m *testing.M) {
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewList(t *testing.T) {
	log.Println("Test NewList running")
	element := testValueOne

	list := NewList()
	list.Add(element)

	head, _ := list.Head().(string)
	if head != element {
		t.Fail()
	}
}

func TestSize(t *testing.T) {
	log.Println("Test Size running")
	list := createThreeElementTestList()

	if list.Size() != 3 {
		t.Fail()
	}
}

func TestHead(t *testing.T) {
	log.Println("Test Head running")

	list1 := createThreeElementTestList()

	headValue, _ := list1.Head().(string)
	if headValue != testValueOne {
		t.Fail()
	}

	element := "c"
	list2 := NewList()
	list2.Add(element)
	headValue2, _ := list2.Head().(string)
	if headValue2 != element {
		t.Fail()
	}
}

func TestGetHead(t *testing.T) {
	log.Println("Test GetHead running")

	list1 := createThreeElementTestList()

	headValueInteface, err := list1.GetHead()
	if err != nil {
		t.Fail()
	}
	headValue, _ := headValueInteface.(string)
	if headValue != testValueOne {
		t.Fail()
	}

	newHeadValue, _ := list1.Head().(string)
	if newHeadValue != testValueTwo {
		t.Fail()
	}
}

func TestTail(t *testing.T) {
	log.Println("Test Tail running")
	element := "c"
	list1 := NewList()
	list1.Add(element)

	tailValue, _ := list1.Tail().(string)
	if tailValue != element {
		t.Fail()
	}

	list2 := createThreeElementTestList()

	tailValue2, _ := list2.Tail().(string)
	if tailValue2 != testValueThree {
		t.Fail()
	}
}

func TestGetTail(t *testing.T) {
	log.Println("Test GetTail running")

	list1 := createThreeElementTestList()

	tailValueInteface, err := list1.GetTail()
	if err != nil {
		t.Fail()
	}
	tailValue, _ := tailValueInteface.(string)
	if tailValue != testValueThree {
		t.Fail()
	}

	newTailValue, _ := list1.Tail().(string)
	if newTailValue != testValueTwo {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	log.Println("Test Add running")
	// create empty list
	list1 := &List{}
	element := "c"
	list1.Add(element)

	if list1.Size() != 1 {
		t.Fail()
	}

	list1.Add(element)
	if list1.Size() != 2 {
		t.Fail()
	}
}

func TestRemoveTail(t *testing.T) {
	log.Println("Test RemoveTail running")

	// Test on three element list
	list := createThreeElementTestList()
	// 1. Remove tail
	err := list.RemoveTail()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	tail, _ := list.Tail().(string)
	if tail != testValueTwo {
		t.Fail()
	}
	// 2. Remove tail second time
	err = list.RemoveTail()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	tail2, _ := list.Tail().(string)
	if tail2 != testValueOne {
		t.Fail()
	}
	// 3. Remove tail third time
	err = list.RemoveTail()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	// Head and tail should be empty
	if list.Head() != nil {
		t.Fail()
	}
	if list.Tail() != nil {
		t.Fail()
	}
}

func TestRemoveHead(t *testing.T) {
	log.Println("Test RemoveHead running")

	// Testing on the three element list
	list := createThreeElementTestList()
	// 1. Remove head
	err := list.RemoveHead()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	head, _ := list.Head().(string)
	if head != testValueTwo {
		t.Fail()
	}
	// 2. Remove head second time
	err = list.RemoveHead()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	head2, _ := list.Head().(string)
	if head2 != testValueThree {
		t.Fail()
	}
	// Remove head thir time
	err = list.RemoveHead()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	// Head and tail should be empty
	if list.Head() != nil {
		t.Fail()
	}

	if list.Tail() != nil {
		t.Fail()
	}
}

func TestRemoveByOrderInList(t *testing.T) {
	log.Println("Test RemoveByOrderNo running")
	// Creating three element list
	list1 := createThreeElementTestList()
	// removing first element of the list
	err := list1.RemoveByOrderInList(0)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	headOfList1, _ := list1.Head().(string)
	if headOfList1 != testValueTwo {
		t.Fail()
	}

	// Creating new three element list
	list2 := createThreeElementTestList()
	// removing second element of the list
	err = list2.RemoveByOrderInList(1)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	headOfList2, _ := list2.Head().(string)
	if headOfList2 != testValueOne {
		t.Fail()
	}
	tailOfList2, _ := list2.Tail().(string)
	if tailOfList2 != testValueThree {
		t.Fail()
	}
	// Testing of removing the third element on new list
	list3 := createThreeElementTestList()
	err = list3.RemoveByOrderInList(2)
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	headOfList3, _ := list3.Head().(string)
	if headOfList3 != testValueOne {
		t.Fail()
	}

	tailOfList3, _ := list3.Tail().(string)
	if tailOfList3 != testValueTwo {
		t.Fail()
	}
}

func createThreeElementTestList() *List {

	list := NewList()
	list.Add(testValueOne)

	list.Add(testValueTwo)

	list.Add(testValueThree)

	return list
}
