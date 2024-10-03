package main

import (
	"example/linkedListLibrary"
	"fmt"
)

func main() {
	fmt.Println("let's help you with a linked list")
	basicAddStuff()
	basicRemovalStuff()
	removeAtEndStuff()
	removeAtBeginningStuff()
	basicCopyStuff()
}

func basicAddStuff() {
	list := linkedListLibrary.ConstructFromValues("1", "2", "3")
	fmt.Printf("here's your list: %s \n", list.GetStringForm())
	fmt.Println("")
}

func basicRemovalStuff() {
	list := linkedListLibrary.ConstructFromValues("w", "h", "a", "t")
	fmt.Printf("here's your list: %s \n", list.GetStringForm())

	list.RemoveValue("i")
	list.RemoveValue("w")
	fmt.Printf("here's your list: %s \n", list.GetStringForm())

	list.RemoveValue("a")
	fmt.Printf("here's your list: %s \n", list.GetStringForm())

	list.RemoveValue("t")
	fmt.Printf("here's your list: %s \n", list.GetStringForm())

	list.RemoveValue("h")
	fmt.Printf("here's your list: %s \n", list.GetStringForm())

	list.RemoveValue("b")
	fmt.Println("")
}

func removeAtEndStuff() {
	list := linkedListLibrary.ConstructFromValues("w", "h", "a", "t")
	fmt.Printf("here's your list: %s \n", list.GetStringForm())

	for !list.IsEmpty() {
		list.RemoveAtEnd()
		fmt.Printf("here's your list: %s \n", list.GetStringForm())
	}

	list.RemoveAtEnd()
	fmt.Printf("here's your list: %s \n", list.GetStringForm())
	fmt.Println("")
}

func removeAtBeginningStuff() {
	list := linkedListLibrary.ConstructFromValues("w", "h", "a", "t")
	fmt.Printf("here's your list: %s \n", list.GetStringForm())

	for !list.IsEmpty() {
		list.RemoveAtBeginning()
		fmt.Printf("here's your list: %s \n", list.GetStringForm())
	}

	list.RemoveAtBeginning()
	fmt.Printf("here's your list: %s \n", list.GetStringForm())
	fmt.Println("")
}

func basicCopyStuff() {
	list := linkedListLibrary.ConstructFromValues("w", "h", "a", "t")
	fmt.Printf("here's your list: %s \n", list.GetStringForm())

	list2 := list.Copy()
	fmt.Printf("here's your other list: %s \n", list2.GetStringForm())
}
