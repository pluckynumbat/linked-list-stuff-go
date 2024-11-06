package dlistlib

import (
	"fmt"
)

const NilListError string = "The list is nil"

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
		return fmt.Errorf(NilListError)
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
		return fmt.Errorf(NilListError)
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
