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

