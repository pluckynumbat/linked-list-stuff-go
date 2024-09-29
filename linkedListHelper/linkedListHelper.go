package main

import (
	"example/linkedListLibrary"
	"fmt"
)

func main() {
	fmt.Println("let's help you with a linked list")
	basicAddStuff()

}

func basicAddStuff() {
	list := linkedListLibrary.ConstructFromValues("1", "2", "3")
	fmt.Printf("here's your list: %s \n", list.GetStringForm())
	fmt.Println("")
}
