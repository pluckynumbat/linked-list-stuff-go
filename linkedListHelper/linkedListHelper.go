package main

import (
	"example/linkedListLibrary"
	"fmt"
)

func main() {
	fmt.Println("let's help you with a linked list")
	basicLinkedListStuff()
}

func basicLinkedListStuff() {
	linkedList := linkedListLibrary.ConstructList(linkedListLibrary.ConstructNode("3"))
	linkedList.PrintList()
	linkedList.AddToEnd("4")
	linkedList.AddToEnd("5")
	linkedList.PrintList()
	linkedList.AddToBeginning("2")
	linkedList.AddToBeginning("1")
	linkedList.PrintList()
	linkedList.Reverse()
	linkedList.PrintList()
}
