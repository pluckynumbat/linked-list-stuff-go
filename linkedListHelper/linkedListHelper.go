package main

import (
	"example/linkedListLibrary"
	"fmt"
)

func main() {
	fmt.Println("let's help you with a linked list")
	linkedList := linkedListLibrary.ConstructList(linkedListLibrary.ConstructNode("2"))
	linkedList.AddToBeginning("1")
	linkedList.AddToEnd("3")
	fmt.Printf("here's your list: %s", linkedList.GetStringForm())
	fmt.Println("")
}
