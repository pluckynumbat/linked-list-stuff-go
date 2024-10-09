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
	fmt.Println("here's your list:", list.GetStringForm())

	list2 := list.Copy()
	fmt.Println("here's your other list:", list2.GetStringForm())
}
