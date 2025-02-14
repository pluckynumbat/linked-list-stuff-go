package listlib

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	data string
	next *Node
}

func constructNode(inputData string) *Node {
	node := new(Node) //used 'new' allocator
	node.data = inputData
	node.next = nil
	return node
}

func (n *Node) GetData() (string, error) {
	if n == nil {
		return "", fmt.Errorf("the node is nil")
	}

	return n.data, nil
}

func constructList(newHead *Node) LinkedList {
	linkedList := LinkedList{newHead} // used composite literal
	linkedList.head = newHead
	return linkedList
}

func ConstructFromValues(values ...string) LinkedList {
	if len(values) == 0 {
		fmt.Println("Warning, creating an empty list")
		return LinkedList{}
	}

	newList := constructList(constructNode(values[0]))

	// TODO: make this use AddToBeginning instead of AddToEnd to make it more efficient
	for i := 1; i < len(values); i++ {
		newList.AddToEnd(values[i])
	}

	return newList
}

func (list *LinkedList) IsNil() bool {
	return list == nil
}

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) Head() *Node {
	if list.IsNil() {
		return nil
	}
	return list.head
}

func (list *LinkedList) Copy() LinkedList {
	if list.IsEmpty() {
		fmt.Println("Warning, copying an empty list")
		return LinkedList{}
	}

	newList := constructList(constructNode(list.head.data))

	runner := list.head.next

	// AddToBeginning(), followed by a Reverse() should be more efficient than AddToEnd()
	for runner != nil {
		newList.AddToBeginning(runner.data)
		runner = runner.next
	}

	newList.Reverse()

	return newList

}

func (list *LinkedList) AddToEnd(inputData string) {

	node := constructNode(inputData)

	if list.IsEmpty() {
		list.head = node
		return
	}

	runner := list.head
	for runner.next != nil {
		runner = runner.next
	}
	runner.next = node
}

func (list *LinkedList) AddToBeginning(inputData string) {

	node := constructNode(inputData)

	if list.IsEmpty() {
		list.head = node
		return
	}

	node.next = list.head
	list.head = node
}

func (list *LinkedList) Reverse() {
	var prev, nxt *Node
	prev = nil
	nxt = nil

	curr := list.head

	for curr != nil {
		nxt = curr.next
		curr.next = prev
		prev = curr
		curr = nxt
	}
	list.head = prev
}

func (list *LinkedList) GetStringForm() string {
	result := ""

	runner := list.head
	for runner != nil {
		result += runner.data + "->"
		runner = runner.next
	}
	result += "nil"
	return result
}

func (list *LinkedList) RemoveValue(val string) {
	fmt.Println("")
	if list.IsEmpty() { // no elements in the list
		fmt.Println("The list doesn't have anything")
		return
	}

	if strings.Compare(list.head.data, val) == 0 { // the value is at the head
		fmt.Printf("found the value '%s'! removing it\n", val)
		list.head = list.head.next
		return
	}

	//we can start with head.next since we've already covered the value at head case above
	runner := list.head.next
	prev := list.head

	var found bool
	for runner != nil {
		if strings.Compare(runner.data, val) == 0 {
			fmt.Printf("found the value '%s'! removing it\n", val)
			prev.next = runner.next
			found = true
			break
		}
		prev = runner
		runner = runner.next
	}

	if !found {
		fmt.Printf("value '%s' wasn't found in the list\n", val)
	}

}

func (list *LinkedList) RemoveAtEnd() *Node {
	if list.IsEmpty() { // no elements in the list
		fmt.Println("The list doesn't have anything")
		return nil
	}

	var removed *Node

	if list.head.next == nil { // only one element
		removed = list.head
		list.head = nil
		return removed
	}

	runner := list.head
	var prev *Node

	for runner.next != nil {
		prev = runner
		runner = runner.next
	}

	removed = runner
	runner = nil
	prev.next = runner

	return removed
}

func (list *LinkedList) RemoveAtBeginning() *Node {
	if list.IsEmpty() { // no elements in the list
		fmt.Println("The list doesn't have anything")
		return nil
	}

	var removed *Node

	if list.head.next == nil { // only one element
		removed = list.head
		list.head = nil
		return removed
	}

	removed = list.head
	list.head = list.head.next

	return removed
}
