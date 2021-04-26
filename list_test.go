package list

import (
	"log"
	"os"
	"testing"
)

var (
	testValueOne   string = "one"
	testValueTwo   string = "two"
	testValueThree string = "three"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewList(t *testing.T) {
	log.Println("Test NewList running")
	element := testValueOne

	list := NewList(element)

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
	list2 := NewList(element)
	headValue2, _ := list2.Head().(string)
	if headValue2 != element {
		t.Fail()
	}
}

func TestTail(t *testing.T) {
	log.Println("Test Tail running")
	element := "c"
	list1 := NewList(element)
	value, _ := list1.Tail().(string)
	if value != element {
		t.Fail()
	}

	list2 := createThreeElementTestList()
	valueThree, _ := list2.Tail().(string)
	if valueThree != testValueThree {
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

func TestRemoveByOrderNo(t *testing.T) {
	log.Println("Test RemoveByOrderNo running")
	list1 := createThreeElementTestList()
	err := list1.RemoveByOrderNo(0)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	head, _ := list1.Head().(string)
	if head != testValueTwo {
		t.Fail()
	}
}

func createThreeElementTestList() *List {

	list := NewList(testValueOne)

	list.Add(testValueTwo)

	list.Add(testValueThree)

	return list
}
