package linkedList

import "errors"

var ErrListIsEmpty error
var ErrNegativeOrder error
var ErrOrderOutOfSize error
var ErrElementNotExists error

func init() {
	ErrListIsEmpty = errors.New("Error: list is empty")
	ErrNegativeOrder = errors.New("Error: element order can not be negative")
	ErrOrderOutOfSize = errors.New("Error: order out of list size")
	ErrElementNotExists = errors.New("Error: element does not exists in the list")
}
