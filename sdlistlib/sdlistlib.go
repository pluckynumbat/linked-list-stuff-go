// Semi generic doubly linked list with nodes that contain
// data of any type that implements the fmt.Stringer interface

package sdlistlib

import (
	"fmt"
)

var nilListError error = fmt.Errorf("The list is nil")
var emptyListError error = fmt.Errorf("The list is empty")

// Node's data can be anything that implements the Stringer interface
type Node[T fmt.Stringer] struct {
	prev *Node[T]
	data T
	next *Node[T]
}

// Node's implementation of the Stringer interface
func (node *Node[T]) String() string {
	if node == nil {
		return "nil"
	}

	return node.data.String()
}

type SemiGenericList[T fmt.Stringer] struct {
	head *Node[T]
	tail *Node[T]
}

// Semi Generic List's implementation of the Stringer interface
func (list *SemiGenericList[T]) String() string {
	if list.IsNil() {
		return "nil"
	}

	if list.IsEmpty() {
		return "empty"
	}

	str := "nil<-"

	runner := list.head

	for runner != nil {
		str += runner.String()
		if runner.next != nil {
			str += "<=>"
		}
		runner = runner.next
	}
	str += "->nil"

	return str
}

// Method to check whether a list is nil
func (list *SemiGenericList[T]) IsNil() bool {
	return list == nil
}

// Method to check whether a list is empty
func (list *SemiGenericList[T]) IsEmpty() bool {
	return list.IsNil() || list.head == nil
}

// Method to return a pointer to the head of the list
func (list *SemiGenericList[T]) Head() *Node[T] {
	if list.IsNil() {
		return nil
	}

	return list.head
}

// Method to return a pointer to the tail of the list
func (list *SemiGenericList[T]) Tail() *Node[T] {
	if list.IsNil() {
		return nil
	}

	return list.tail
}

// Method to add a new node to the beginning of the list
func (list *SemiGenericList[T]) AddAtBeginning(val T) error {
	if list.IsNil() {
		return nilListError
	}

	node := &Node[T]{nil, val, nil}

	if list.IsEmpty() {
		list.head = node
		list.tail = node
		return nil
	}

	node.next = list.head
	list.head.prev = node
	list.head = node
	return nil
}

// Method to add a new node to the end of the list
func (list *SemiGenericList[T]) AddAtEnd(val T) error {
	if list.IsNil() {
		return nilListError
	}

	node := &Node[T]{nil, val, nil}

	if list.IsEmpty() {
		list.head = node
		list.tail = node
		return nil
	}

	node.prev = list.tail
	list.tail.next = node
	list.tail = node
	return nil
}

// Method to reverse the semi generic doubly linked list
func (list *SemiGenericList[T]) Reverse() error {
	if list.IsNil() {
		return nilListError
	}

	if list.IsEmpty() {
		return nil
	}

	if list.head == list.tail {
		return nil
	}

	runner := list.head
	var nxt *Node[T]

	for runner != nil {
		nxt = runner.next

		runner.prev, runner.next = runner.next, runner.prev
		runner = nxt
	}

	list.head, list.tail = list.tail, list.head
	return nil
}

// Method to copy a semi generic doubly linked list
func (list *SemiGenericList[T]) Copy() (*SemiGenericList[T], error) {
	if list.IsNil() {
		return nil, nilListError
	}

	result := &SemiGenericList[T]{}
	runner := list.head

	for runner != nil {
		result.AddAtEnd(runner.data)
		runner = runner.next
	}

	return result, nil
}

// Method to remove a given value from the list, if present
// If present more than once, this will remove the first occurrence
func (list *SemiGenericList[T]) RemoveValue(val T) error {
	if list.IsNil() {
		return nilListError
	}

	if list.IsEmpty() {
		return emptyListError
	}

	runner := list.head

	for runner != nil {
		// String() is needed to compare the values
		if runner.data.String() == val.String() {

			// single element list
			if list.head == list.tail {
				list.head = nil
				list.tail = nil
				return nil
			}

			// first element of the list
			if runner == list.head {
				list.head = list.head.next
				list.head.prev = nil
				return nil
			}

			// last element of the list
			if runner == list.tail {
				list.tail = list.tail.prev
				list.tail.next = nil
				return nil
			}

			//any other element
			runner.prev.next = runner.next
			runner.next.prev = runner.prev
			return nil
		}
		runner = runner.next
	}

	return fmt.Errorf("The value: %v is not present in the list", val)
}
