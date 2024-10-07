package main

import (
	"example/linkedListLibrary"
	"fmt"
)

func main() {
	fmt.Println("let's help you with a linked list")
	basicCopyStuff()
}

func basicCopyStuff() {
	list := linkedListLibrary.ConstructFromValues("w", "h", "a", "t")
	fmt.Printf("here's your list: %s \n", list.GetStringForm())

	list2 := list.Copy()
	fmt.Printf("here's your other list: %s \n", list2.GetStringForm())
}
