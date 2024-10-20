package tlistlib

type TailedList struct {
	head, tail *Node
}

type Node struct {
	data string
	next *Node
}
