package tlistlib

import "fmt"

type TailedList struct {
	head, tail *Node
}

type Node struct {
	data string
	next *Node
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	return n.data
}

func (tl *TailedList) String() string {
	if tl == nil {
		return ""
	}

	str := ""

	runner := tl.head
	for runner != nil {
		str += runner.String()
		str += "->"
		runner = runner.next
	}
	str += "nil"
	return str
}

func constructListFromNode(node *Node) (tlist TailedList) {
	tlist = TailedList{}
	tlist.head = node
	tlist.tail = node

	return
}

func (tl *TailedList) IsEmpty() bool {
	return tl.head == nil
}

func (tl *TailedList) AddAtEnd(val string) {

	n := &Node{val, nil}

	if tl.IsEmpty() {
		tl.head = n
		tl.tail = n
		return
	}

	tl.tail.next = n
	tl.tail = n
}

func (tl *TailedList) AddAtBeginning(val string) {

	n := &Node{val, nil}

	if tl.IsEmpty() {
		tl.head = n
		tl.tail = n
		return
	}

	n.next = tl.head
	tl.head = n
}

func (tl *TailedList) RemoveValue(val string) error {

	notFoundErr := fmt.Errorf("the value %v isn't present in the list", val)

	//no elements in the list
	if tl.IsEmpty() {
		return fmt.Errorf("the list is empty")
	}

	//single element list
	if tl.head == tl.tail {
		if val == tl.head.data {
			tl.head = nil
			tl.tail = nil
			return nil
		} else {
			return notFoundErr
		}
	}

	runner := tl.head
	var prev *Node

	for runner != nil {
		if runner.data == val {

			//head has the value
			if runner == tl.head {
				tl.head = runner.next
				return nil
			}

			prev.next = runner.next

			//tail has the value
			if runner == tl.tail {
				tl.tail = prev
			}
		}

		prev = runner
		runner = runner.next
	}

	return notFoundErr

}
