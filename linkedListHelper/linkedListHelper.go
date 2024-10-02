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

