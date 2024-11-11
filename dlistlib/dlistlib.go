package dlistlib

import (
	"fmt"
)

var NilListError error = fmt.Errorf("The list is nil")
var EmptyListError error = fmt.Errorf("The list is empty")

type Node struct {
	prev *Node
	data string
	next *Node
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return n.data
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dl *DoublyLinkedList) String() string {
	if dl.IsNil() {
		return "nil"
	}

	if dl.IsEmpty() {
		return "empty"
	}

	str := "nil<-"

	runner := dl.head
	for runner != nil {
		str += runner.data
		if runner.next != nil {
			str += "<=>"
		}
		runner = runner.next
	}
	str += "->nil"

	return str
}

func (dl *DoublyLinkedList) IsNil() bool {
	return dl == nil
}

func (dl *DoublyLinkedList) IsEmpty() bool {
	return dl.head == nil
}

func (dl *DoublyLinkedList) AddAtBeginning(val string) error {

	if dl.IsNil() {
		return NilListError
	}
	node := &Node{nil, val, nil}
	if dl.IsEmpty() {
		dl.head = node
		dl.tail = node
		return nil
	}

	node.next = dl.head
	dl.head.prev = node
	dl.head = node
	return nil
}

func (dl *DoublyLinkedList) AddAtEnd(val string) error {

	if dl.IsNil() {
		return NilListError
	}
	node := &Node{nil, val, nil}

	if dl.IsEmpty() {
		dl.head = node
		dl.tail = node
		return nil
	}

	node.prev = dl.tail
	dl.tail.next = node
	dl.tail = node
	return nil
}

func (dl *DoublyLinkedList) Reverse() error {
	if dl.IsNil() {
		return NilListError
	}

	if dl.IsEmpty() {
		return nil
	}

	runner := dl.head
	var nxt *Node

	for runner != nil {

		nxt = runner.next
		runner.next, runner.prev = runner.prev, runner.next
		runner = nxt
	}

	dl.head, dl.tail = dl.tail, dl.head
	return nil
}

func (dl *DoublyLinkedList) Copy() (*DoublyLinkedList, error) {
	if dl.IsNil() {
		return nil, NilListError
	}

	dlcopy := &DoublyLinkedList{}
	if dl.IsEmpty() {
		return dlcopy, nil
	}

	runner := dl.head

	for runner != nil {
		dlcopy.AddAtEnd(runner.data)
		runner = runner.next
	}
	return dlcopy, nil
}

func (dl *DoublyLinkedList) RemoveValue(val string) error {
	if dl.IsNil() {
		return NilListError
	}

	if dl.IsEmpty() {
		return EmptyListError
	}

	var notFoundError error = fmt.Errorf("The value: %v is not present in the list", val)

	runner := dl.head

	for runner != nil {

		//value is found
		if runner.data == val {

			//single element list
			if runner == dl.head && runner == dl.tail {
				dl.head, dl.tail = nil, nil
				return nil
			}

			//value is at head
			if runner == dl.head {
				dl.head = dl.head.next
				dl.head.prev = nil
				return nil
			}

			//value is at tail
			if runner == dl.tail {
				dl.tail = dl.tail.prev
				dl.tail.next = nil
				return nil
			}

			//value is elsewhere in the list
			runner.prev.next = runner.next
			runner.next.prev = runner.prev
			return nil
		}

		runner = runner.next
	}

	return notFoundError
}

// Method to remove the first element (head) of a doubly linked list
func (dl *DoublyLinkedList) RemoveFirst() (*Node, error) {
	if dl.IsNil() {
		return nil, NilListError
	}

	if dl.IsEmpty() {
		return nil, EmptyListError
	}

	removed := dl.head
	dl.head = dl.head.next

	if dl.head == nil { //only element in the list was removed
		dl.tail = nil
	} else {
		dl.head.prev = nil
	}

	return removed, nil
}

// Method to remove the last element (tail) of a doubly linked list
func (dl *DoublyLinkedList) RemoveLast() (*Node, error) {
	if dl.IsNil() {
		return nil, NilListError
	}

	if dl.IsEmpty() {
		return nil, EmptyListError
	}

	removed := dl.tail
	dl.tail = dl.tail.prev

	if dl.tail == nil { //only element in the list was removed
		dl.head = nil
	} else {
		dl.tail.next = nil
	}

	return removed, nil
}
