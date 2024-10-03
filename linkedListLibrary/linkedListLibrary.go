package linkedListLibrary

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

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) AddToEnd(inputData string) {

	node := constructNode(inputData)

	runner := list.head
	for runner.next != nil {
		runner = runner.next
	}
	runner.next = node
}

func (list *LinkedList) AddToBeginning(inputData string) {

	node := constructNode(inputData)

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

func (list *LinkedList) RemoveAtEnd() {
	if list.IsEmpty() { // no elements in the list
		fmt.Println("The list doesn't have anything")
		return
	}

	if list.head.next == nil { // only one element
		list.head = nil
		return
	}

	runner := list.head
	var prev *Node

	for runner.next != nil {
		prev = runner
		runner = runner.next
	}

	runner = nil
	prev.next = runner
}

func (list *LinkedList) RemoveAtBeginning() {
	if list.IsEmpty() { // no elements in the list
		fmt.Println("The list doesn't have anything")
		return
	}

	if list.head.next == nil { // only one element
		list.head = nil
		return
	}

	list.head = list.head.next
}
