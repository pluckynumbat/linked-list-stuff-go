package dlistlib

import (
	"fmt"
	"testing"
)

func TestNodeString(t *testing.T) {

	var node *Node
	t.Run("Nil Node String", func(t *testing.T) {
		got := node.String()
		want := "nil"

		if got != want {
			t.Errorf("Incorrect string for Node: want %v, got %v", want, got)
		}
	})

	node = &Node{nil, "a", nil}
	t.Run("Non-nil Node String", func(t *testing.T) {
		got := node.String()
		want := "a"

		if got != want {
			t.Errorf("Incorrect string for Node: want %v, got %v", want, got)
		}
	})
}

func TestListString(t *testing.T) {

	var dlist *DoublyLinkedList

	t.Run("Nil List String", func(t *testing.T) {
		got := dlist.String()
		want := "nil"

		if got != want {
			t.Errorf("Incorrect string for a nil list: want %v, got %v", want, got)
		}
	})

	dlist = &DoublyLinkedList{}
	t.Run("Empty List String", func(t *testing.T) {
		got := dlist.String()
		want := "empty"

		if got != want {
			t.Errorf("Incorrect string for an empty list: want %v, got %v", want, got)
		}
	})

	node := &Node{nil, "a", nil}
	dlist = &DoublyLinkedList{node, node}
	t.Run("Single Element List String", func(t *testing.T) {
		got := dlist.String()
		want := "nil<-a->nil"

		if got != want {
			t.Errorf("Incorrect string for a single element list: want %v, got %v", want, got)
		}
	})

	var head, tail *Node
	head = &Node{nil, "a", nil}
	tail = &Node{nil, "b", nil}

	head.next = tail
	tail.prev = head
	dlist = &DoublyLinkedList{head, tail}
	t.Run("Two Element List String", func(t *testing.T) {
		got := dlist.String()
		want := "nil<-a<=>b->nil"

		if got != want {
			t.Errorf("Incorrect string for a 2 element list: want %v, got %v", want, got)
		}
	})
}

func TestIsNil(t *testing.T) {
	var dlist *DoublyLinkedList

	t.Run("Is Nil True", func(t *testing.T) {
		want := true
		got := dlist.IsNil()

		if got != want {
			t.Errorf("Incorrect value of IsNil on nil list: want %v, got %v", want, got)
		}
	})

	dlist = &DoublyLinkedList{}

	t.Run("Is Nil False Struct Literal", func(t *testing.T) {
		want := false
		got := dlist.IsNil()

		if got != want {
			t.Errorf("Incorrect value of IsNil on a non nil list: want %v, got %v", want, got)
		}
	})

	dlist = new(DoublyLinkedList)

	t.Run("Is Nil False New Allocator", func(t *testing.T) {
		want := false
		got := dlist.IsNil()

		if got != want {
			t.Errorf("Incorrect value of IsNil on a non nil list: want %v, got %v", want, got)
		}
	})
}

func TestIsEmpty(t *testing.T) {
	var dlist *DoublyLinkedList

	dlist = &DoublyLinkedList{}

	t.Run("Is Empty True Struct Literal", func(t *testing.T) {
		want := true
		got := dlist.IsEmpty()

		if got != want {
			t.Errorf("Incorrect value of IsEmpty on an empty list: want %v, got %v", want, got)
		}
	})

	dlist = new(DoublyLinkedList)

	t.Run("Is Empty True New Allocator", func(t *testing.T) {
		want := true
		got := dlist.IsEmpty()

		if got != want {
			t.Errorf("Incorrect value of IsEmpty on an empty list: want %v, got %v", want, got)
		}
	})

	node := &Node{nil, "a", nil}
	dlist = &DoublyLinkedList{node, node}

	t.Run("Is Empty False Struct Literal", func(t *testing.T) {
		want := false
		got := dlist.IsEmpty()

		if got != want {
			t.Errorf("Incorrect value of IsEmpty on a non empty list: want %v, got %v", want, got)
		}
	})

	dlist = new(DoublyLinkedList)
	dlist.head = node
	dlist.tail = node

	t.Run("Is Empty False New Allocator", func(t *testing.T) {
		want := false
		got := dlist.IsEmpty()

		if got != want {
			t.Errorf("Incorrect value of IsEmpty on a non empty list: want %v, got %v", want, got)
		}
	})
}

func TestAddAtBeginning(t *testing.T) {
	var dlist *DoublyLinkedList

	t.Run("Nil List", func(t *testing.T) {
		err := dlist.AddAtBeginning("a")
		if err == nil {
			t.Errorf("Adding to a nil list should have failed")
		} else {
			fmt.Println(err)
		}
	})

	dlist = &DoublyLinkedList{}
	t.Run("Empty List", func(t *testing.T) {
		err := dlist.AddAtBeginning("a")
		if err != nil {
			t.Errorf("Adding to an empty list failed, error: %v", err)
		} else {

			want := "nil<-a->nil"
			got := dlist.String()

			if got != want {
				t.Errorf("List after adding to it is incorrect, want: %v, got: %v", want, got)
			}

			want = "a"
			got = dlist.head.String()
			if got != want {
				t.Errorf("Head after adding to an empty list is incorrect, want: %v, got: %v", want, got)
			}

			got = dlist.tail.String()
			if got != want {
				t.Errorf("Tail after adding to an empty list is incorrect, want: %v, got: %v", want, got)
			}
		}
	})

	dlist = &DoublyLinkedList{}
	t.Run("General Add At Beginning", func(t *testing.T) {
		vals := []string{"a", "b", "c", "d", "e"}
		partials := []string{
			"nil<-a->nil",
			"nil<-b<=>a->nil",
			"nil<-c<=>b<=>a->nil",
			"nil<-d<=>c<=>b<=>a->nil",
			"nil<-e<=>d<=>c<=>b<=>a->nil",
		}
		heads := []string{"a", "b", "c", "d", "e"}
		tail := "a"

		for i, v := range vals {
			err := dlist.AddAtBeginning(v)

			if err != nil {
				t.Errorf("Adding to a list failed, error: %v", err)
			} else {

				want := partials[i]
				got := dlist.String()

				if got != want {
					t.Errorf("List after adding to it is incorrect, want: %v, got: %v", want, got)
				}

				want = heads[i]
				got = dlist.head.String()
				if got != want {
					t.Errorf("Head after adding to a list is incorrect, want: %v, got: %v", want, got)
				}

				want = tail
				got = dlist.tail.String()
				if got != want {
					t.Errorf("Tail after adding to a list is incorrect, want: %v, got: %v", want, got)
				}
			}
		}
	})
}

func TestAddAtEnd(t *testing.T) {
	var dlist *DoublyLinkedList

	t.Run("Nil List", func(t *testing.T) {
		err := dlist.AddAtEnd("a")
		if err == nil {
			t.Errorf("Adding to a nil list should have failed")
		} else {
			fmt.Println(err)
		}
	})

	dlist = &DoublyLinkedList{}
	t.Run("Empty List", func(t *testing.T) {
		err := dlist.AddAtEnd("a")
		if err != nil {
			t.Errorf("Adding to an empty list failed, error: %v", err)
		} else {

			want := "nil<-a->nil"
			got := dlist.String()
			if got != want {
				t.Errorf("List after adding to it is incorrect, want: %v, got: %v", want, got)
			}

			want = "a"
			got = dlist.head.String()
			if got != want {
				t.Errorf("Head after adding to an empty list is incorrect, want: %v, got: %v", want, got)
			}

			got = dlist.tail.String()
			if got != want {
				t.Errorf("Tail after adding to an empty list is incorrect, want: %v, got: %v", want, got)
			}
		}
	})

	dlist = &DoublyLinkedList{}
	t.Run("General Add At End", func(t *testing.T) {
		vals := []string{"a", "b", "c", "d", "e"}
		partials := []string{
			"nil<-a->nil",
			"nil<-a<=>b->nil",
			"nil<-a<=>b<=>c->nil",
			"nil<-a<=>b<=>c<=>d->nil",
			"nil<-a<=>b<=>c<=>d<=>e->nil",
		}
		head := "a"
		tails := []string{"a", "b", "c", "d", "e"}

		for i, v := range vals {
			err := dlist.AddAtEnd(v)

			if err != nil {
				t.Errorf("Adding to a list failed, error: %v", err)
			} else {

				want := partials[i]
				got := dlist.String()

				if got != want {
					t.Errorf("List after adding to it is incorrect, want: %v, got: %v", want, got)
				}

				want = head
				got = dlist.head.String()
				if got != want {
					t.Errorf("Head after adding to a list is incorrect, want: %v, got: %v", want, got)
				}

				want = tails[i]
				got = dlist.tail.String()
				if got != want {
					t.Errorf("Tail after adding to a list is incorrect, want: %v, got: %v", want, got)
				}
			}
		}
	})
}
