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

func TestCopy(t *testing.T) {
	var dlist *DoublyLinkedList

	t.Run("Nil list", func(t *testing.T) {
		_, err := dlist.Copy()
		if err == nil {
			t.Errorf("Copying a nil list should have failed")
		} else {
			fmt.Println(err)
		}
	})

	dlist = &DoublyLinkedList{}
	t.Run("Empty list", func(t *testing.T) {
		dlistCopy, err := dlist.Copy()
		if err != nil {
			t.Errorf("Copying an empty list failed, error: %v", err)
		} else {
			want := "empty"
			got := dlistCopy.String()
			if want != got {
				t.Errorf("Copying an empty list returned incorrect results, want: %v, got %v", want, got)
			}
		}
	})

	t.Run("General Copy", func(t *testing.T) {
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

				dlistCopy, err2 := dlist.Copy()

				if err2 != nil {
					t.Errorf("Copying a list failed, error: %v", err)
				} else {

					want := partials[i]
					got := dlistCopy.String()

					if got != want {
						t.Errorf("Copying a list returned incorrect results, want: %v, got: %v", want, got)
					}

					want = head
					got = dlistCopy.head.String()
					if got != want {
						t.Errorf("Head of the copied list is incorrect, want: %v, got: %v", want, got)
					}

					want = tails[i]
					got = dlistCopy.tail.String()
					if got != want {
						t.Errorf("Tail of the copied list is incorrect, want: %v, got: %v", want, got)
					}
				}
			}
		}
	})
}

func TestReverse(t *testing.T) {
	var dlist *DoublyLinkedList

	t.Run("Nil list", func(t *testing.T) {
		err := dlist.Reverse()
		if err == nil {
			t.Errorf("Reversing a nil list should have failed")
		} else {
			fmt.Println(err)
		}
	})

	dlist = &DoublyLinkedList{}
	t.Run("Empty list", func(t *testing.T) {
		err := dlist.Reverse()
		if err != nil {
			t.Errorf("Reversing an empty list failed, error: %v", err)
		} else {
			want := "empty"
			got := dlist.String()
			if want != got {
				t.Errorf("Reversing an empty list returned incorrect results, want: %v, got %v", want, got)
			}
		}
	})

	t.Run("General Reverse", func(t *testing.T) {
		vals := []string{"a", "b", "c", "d", "e"}
		partials := []string{
			"nil<-a->nil",
			"nil<-b<=>a->nil",
			"nil<-c<=>a<=>b->nil",
			"nil<-d<=>b<=>a<=>c->nil",
			"nil<-e<=>c<=>a<=>b<=>d->nil",
		}
		heads := []string{"a", "b", "c", "d", "e"}
		tails := []string{"a", "a", "b", "c", "d"}

		for i, v := range vals {
			err := dlist.AddAtEnd(v)

			if err != nil {
				t.Errorf("Adding to a list failed, error: %v", err)
			} else {

				err2 := dlist.Reverse()

				if err2 != nil {
					t.Errorf("Reversing a list failed, error: %v", err)
				} else {

					want := partials[i]
					got := dlist.String()

					if got != want {
						t.Errorf("Reversing a list returned incorrect results, want: %v, got: %v", want, got)
					}

					want = heads[i]
					got = dlist.head.String()
					if got != want {
						t.Errorf("Head of the reversed list is incorrect, want: %v, got: %v", want, got)
					}

					want = tails[i]
					got = dlist.tail.String()
					if got != want {
						t.Errorf("Tail of the reversed list is incorrect, want: %v, got: %v", want, got)
					}
				}
			}
		}
	})
}

func TestRemoveValue(t *testing.T) {
	var dl *DoublyLinkedList

	t.Run("Nil list", func(t *testing.T) {
		err := dl.RemoveValue("a")
		if err == nil {
			t.Errorf("Removing a value from a nil list should have failed")
		} else {
			fmt.Println(err)
		}
	})

	dl = &DoublyLinkedList{}
	t.Run("Empty list", func(t *testing.T) {
		err := dl.RemoveValue("a")
		if err == nil {
			t.Errorf("Removing a value from an empty list should have failed")
		} else {
			fmt.Println(err)
		}
	})

	dl = &DoublyLinkedList{}
	t.Run("Invalid Value, Single Element list", func(t *testing.T) {
		err := dl.AddAtEnd("b")
		if err != nil {
			t.Errorf("Adding to a list failed, error: %v", err)
		} else {
			err2 := dl.RemoveValue("a")
			if err2 == nil {
				t.Errorf("Removing an invalid value from a single element list failed should have failed")
			} else {
				fmt.Println(err2)
			}
		}
	})

	dl = &DoublyLinkedList{}
	t.Run("Valid Value, Single Element list", func(t *testing.T) {
		err := dl.AddAtEnd("a")
		if err != nil {
			t.Errorf("Adding to a list failed, error: %v", err)
		} else {
			err2 := dl.RemoveValue("a")
			if err2 != nil {
				t.Errorf("Removing a value from a single element list failed, error:  %v", err2)
			} else {

				want := true
				got := dl.IsEmpty()

				if got != want {
					t.Errorf("Removing the only value from a single element list should result in an empty list")
				}
			}
		}
	})

	dl = &DoublyLinkedList{}
	t.Run("Invalid Value, Multi-Element list", func(t *testing.T) {
		vals := []string{"a", "b", "c", "d", "e"}

		for _, v := range vals {
			err := dl.AddAtEnd(v)
			if err != nil {
				t.Fatalf("Adding to a list failed, error: %v", err)
			}
		}

		err2 := dl.RemoveValue("f")
		if err2 == nil {
			t.Errorf("Removing an invalid value from a list failed should have failed")
		} else {
			fmt.Println(err2)
		}
	})

	dl = &DoublyLinkedList{}
	t.Run("Remove values till list is empty", func(t *testing.T) {
		vals := []string{"a", "b", "c", "d", "e"}

		for _, v := range vals {
			err := dl.AddAtEnd(v)
			if err != nil {
				t.Fatalf("Adding to a list failed, error: %v", err)
			}
		}

		removes := []string{"c", "b", "e", "a", "d"}
		partials := []string{
			"nil<-a<=>b<=>d<=>e->nil",
			"nil<-a<=>d<=>e->nil",
			"nil<-a<=>d->nil",
			"nil<-d->nil",
			"empty",
		}
		heads := []string{"a", "a", "a", "d", "nil"}
		tails := []string{"e", "e", "d", "d", "nil"}
		for i, v := range removes {
			err := dl.RemoveValue(v)
			if err != nil {
				t.Errorf("Removing from a list failed, error: %v", err)
			} else {
				want := partials[i]
				got := dl.String()

				if got != want {
					t.Errorf("Incorrect list after removing an element, want: %v, got: %v", want, got)
				}

				want = heads[i]
				got = dl.head.String()
				if got != want {
					t.Errorf("Incorrect head after removing an element, want: %v, got: %v", want, got)
				}

				want = tails[i]
				got = dl.tail.String()
				if got != want {
					t.Errorf("Incorrect tail after removing an element, want: %v, got: %v", want, got)
				}
			}
		}
	})
}

func TestRemoveFirst(t *testing.T) {
	var dl *DoublyLinkedList

	t.Run("Nil list", func(t *testing.T) {
		_, err := dl.RemoveFirst()
		if err == nil {
			t.Errorf("Removing the first element of a nil list should have failed")
		} else {
			fmt.Println(err)
		}
	})

	dl = &DoublyLinkedList{}
	t.Run("Empty list", func(t *testing.T) {
		_, err := dl.RemoveFirst()
		if err == nil {
			t.Errorf("Removing the first element of an empty list should have failed")
		} else {
			fmt.Println(err)
		}
	})

	dl = &DoublyLinkedList{}
	t.Run("Single Element list", func(t *testing.T) {
		err := dl.AddAtEnd("a")
		if err != nil {
			t.Errorf("Adding to a list failed, error: %v", err)
		} else {
			removed, err2 := dl.RemoveFirst()
			if err2 != nil {
				t.Errorf("Removing the first element from a single element list failed, error:  %v", err2)
			} else {

				want := "a"
				got := removed.String()
				if got != want {
					t.Errorf("Removed element from a list is incorrect, want: %v, got %v", want, got)
				}

				want = "empty"
				got = dl.String()
				if got != want {
					t.Errorf("Removing the only value from a single element list should result in an empty list")
				}
			}
		}
	})

	dl = &DoublyLinkedList{}
	t.Run("Remove first till list is empty", func(t *testing.T) {
		vals := []string{"a", "b", "c", "d", "e"}

		for _, v := range vals {
			err := dl.AddAtEnd(v)
			if err != nil {
				t.Fatalf("Adding to a list failed, error: %v", err)
			}
		}

		removed := []string{"a", "b", "c", "d", "e"}
		partials := []string{
			"nil<-b<=>c<=>d<=>e->nil",
			"nil<-c<=>d<=>e->nil",
			"nil<-d<=>e->nil",
			"nil<-e->nil",
			"empty",
		}
		heads := []string{"b", "c", "d", "e", "nil"}
		tails := []string{"e", "e", "e", "e", "nil"}
		for i := 0; !dl.IsEmpty(); i++ {
			rem, err := dl.RemoveFirst()
			if err != nil {
				t.Errorf("Removing from a list failed, error: %v", err)
			} else {
				want := removed[i]
				got := rem.String()

				if got != want {
					t.Errorf("Removed element from a list is incorrect, want: %v, got %v", want, got)
				}

				want = partials[i]
				got = dl.String()

				if got != want {
					t.Errorf("Incorrect list after removing the first element, want: %v, got: %v", want, got)
				}

				want = heads[i]
				got = dl.head.String()
				if got != want {
					t.Errorf("Incorrect head after removing the first element, want: %v, got: %v", want, got)
				}

				want = tails[i]
				got = dl.tail.String()
				if got != want {
					t.Errorf("Incorrect tail after removing the first element, want: %v, got: %v", want, got)
				}
			}
		}
	})
}

func TestRemoveLast(t *testing.T) {
	var dl *DoublyLinkedList

	t.Run("Nil list", func(t *testing.T) {
		_, err := dl.RemoveLast()
		if err == nil {
			t.Errorf("Removing the last element of a nil list should have failed")
		} else {
			fmt.Println(err)
		}
	})

	dl = &DoublyLinkedList{}
	t.Run("Empty list", func(t *testing.T) {
		_, err := dl.RemoveLast()
		if err == nil {
			t.Errorf("Removing the last element of an empty list should have failed")
		} else {
			fmt.Println(err)
		}
	})

	dl = &DoublyLinkedList{}
	t.Run("Single Element list", func(t *testing.T) {
		err := dl.AddAtEnd("a")
		if err != nil {
			t.Errorf("Adding to a list failed, error: %v", err)
		} else {
			removed, err2 := dl.RemoveLast()
			if err2 != nil {
				t.Errorf("Removing the last element from a single element list failed, error:  %v", err2)
			} else {

				want := "a"
				got := removed.String()
				if got != want {
					t.Errorf("Removed element from a list is incorrect, want: %v, got %v", want, got)
				}

				want = "empty"
				got = dl.String()
				if got != want {
					t.Errorf("Removing the only value from a single element list should result in an empty list")
				}
			}
		}
	})
}
