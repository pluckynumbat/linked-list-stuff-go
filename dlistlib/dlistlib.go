package dlistlib
type Node struct {
	prev *Node
	data string
	next *Node

}
type DoublyLinkedList struct {
	head *Node
	tail *Node
}
