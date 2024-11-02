package dlistlib

type Node struct {
	prev *Node
	data string
	next *Node
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return "<-" + n.data + "->"
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dl *DoublyLinkedList) String() string {
	if dl == nil || dl.head == nil {
		return "nil"
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
