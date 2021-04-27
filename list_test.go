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

	elementValue, err := list.Head()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	head, _ := elementValue.(string)
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
	headValueInterface, err := list1.Head()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	headValue, _ := headValueInterface.(string)
	if headValue != testValueOne {
		t.Fail()
	}

	element := "c"
	list2 := NewList(element)
	headValueInterface2, err := list2.Head()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	headValue2, _ := headValueInterface2.(string)
	if headValue2 != element {
		t.Fail()
	}
}

func TestTail(t *testing.T) {
	log.Println("Test Tail running")
	element := "c"
	list1 := NewList(element)
	tailValueInterface, err := list1.Tail()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	tailValue, _ := tailValueInterface.(string)
	if tailValue != element {
		t.Fail()
	}

	list2 := createThreeElementTestList()
	tailValueInterface2, err := list2.Tail()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	tailValue2, _ := tailValueInterface2.(string)
	if tailValue2 != testValueThree {
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

	emptyList := &List{}
	_, err := emptyList.Tail()
	if err != ErrListIsEmpty {
		t.Fail()
	}

	list := createThreeElementTestList()
	err = list.RemoveTail()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	tailValueInterface, err := list.Tail()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	tail, _ := tailValueInterface.(string)

	if tail != testValueTwo {
		t.Fail()
	}

	err = list.RemoveTail()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	tailValueInterface2, err := list.Tail()
	tail2, _ := tailValueInterface2.(string)
	if tail2 != testValueOne {
		t.Fail()
	}

	err = list.RemoveTail()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	headShouldBeEmpty, err := list.Head()
	if err != ErrListIsEmpty {
		t.Fail()
	}

	if headShouldBeEmpty != nil {
		t.Fail()
	}
	tailShouldBeEmpty, err := list.Tail()
	if err != ErrListIsEmpty {
		t.Fail()
	}
	if tailShouldBeEmpty != nil {
		t.Fail()
	}
}

func TestRemoveHead(t *testing.T) {
	log.Println("Test RemoveHead running")

	emptyList := &List{}
	_, err := emptyList.Head()
	if err.Error() != ErrListIsEmpty.Error() {
		t.Fail()
	}

	list := createThreeElementTestList()
	err = list.RemoveHead()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	headValueInterface, err := list.Head()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	head, _ := headValueInterface.(string)
	if head != testValueTwo {
		t.Fail()
	}

	err = list.RemoveHead()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	headValueInterface2, err := list.Head()
	head2, _ := headValueInterface2.(string)
	if head2 != testValueThree {
		t.Fail()
	}

	err = list.RemoveHead()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	headShouldBeEmpty, err := list.Head()
	//head3, _ := headValueInterface3.(string)
	if err.Error() != ErrListIsEmpty.Error() {
		t.Fail()
	}
	if headShouldBeEmpty != nil {
		t.Fail()
	}
	tailShouldBeEmpty, err := list.Tail()
	//head3, _ := headValueInterface3.(string)
	if err.Error() != ErrListIsEmpty.Error() {
		t.Fail()
	}
	if tailShouldBeEmpty != nil {
		t.Fail()
	}
}

func TestRemoveByOrderInList(t *testing.T) {
	log.Println("Test RemoveByOrderNo running")
	list1 := createThreeElementTestList()
	err := list1.RemoveByOrderInList(0)
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	headValueInterfaceList1, err := list1.Head()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	headOfList1, _ := headValueInterfaceList1.(string)
	if headOfList1 != testValueTwo {
		t.Fail()
	}

	list2 := createThreeElementTestList()
	err = list2.RemoveByOrderInList(1)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	headValueInterfaceList2, err := list2.Head()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	headOfList2, _ := headValueInterfaceList2.(string)
	if headOfList2 != testValueOne {
		t.Fail()
	}
	tailValueIntefaceList2, err := list2.Tail()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	tailOfList2, _ := tailValueIntefaceList2.(string)
	if tailOfList2 != testValueThree {
		t.Fail()
	}

	list3 := createThreeElementTestList()
	err = list3.RemoveByOrderInList(2)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	headValueInterfaceList3, err := list3.Head()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	headOfList3, _ := headValueInterfaceList3.(string)
	if headOfList3 != testValueOne {
		t.Fail()
	}

	tailValueIntefaceList3, err := list3.Tail()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	tailOfList3, _ := tailValueIntefaceList3.(string)
	if tailOfList3 != testValueTwo {
		t.Fail()
	}
}

func createThreeElementTestList() *List {

	list := NewList(testValueOne)

	list.Add(testValueTwo)

	list.Add(testValueThree)

	return list
}
