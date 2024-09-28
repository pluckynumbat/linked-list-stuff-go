package linkedListLibrary

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	data string
	next *Node
}

func ConstructNode(inputData string) *Node {
	node := new(Node) //used 'new' allocator
	node.data = inputData
	node.next = nil
	return node
}

func ConstructList(newHead *Node) LinkedList {
	linkedList := LinkedList{newHead} // used composite literal
	linkedList.head = newHead
	return linkedList
}

func (list *LinkedList) AddToEnd(inputData string) {

	node := ConstructNode(inputData)

	runner := list.head
	for runner.next != nil {
		runner = runner.next
	}
	runner.next = node
}

func (list *LinkedList) AddToBeginning(inputData string) {

	node := ConstructNode(inputData)

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

func (list *LinkedList) PrintList() {
	fmt.Println("Printing the list: ")
	var runner *Node
	runner = list.head

	for runner != nil {
		fmt.Printf("%s->", runner.data)
		runner = runner.next
	}
	fmt.Print("null")
	fmt.Println("")
}

func (list *LinkedList) GetStringForm() string {
	result := ""

	runner := list.head
	for runner != nil {
		result += runner.data + "->"
		runner = runner.next
	}
	result += "null"
	return result
}
