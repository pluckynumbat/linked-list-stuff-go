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

