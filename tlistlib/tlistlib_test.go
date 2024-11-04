package tlistlib

import (
	"fmt"
	"testing"
)

func TestNodeString(t *testing.T) {
	n := &Node{"a", nil}

	want := "a"
	got := n.String()

	if got != want {
		t.Errorf("*Node.String() is incorrect: want '%s', got '%s'", want, got)
	}
}

func TestSingleElementTailedListStrings(t *testing.T) {
	tl := constructListFromNode(&Node{"a", nil})

	t.Run("Single Element Tailed List String", func(t *testing.T) {
		want := "a->nil"
		got := tl.String()

		if got != want {
			t.Errorf("*TailedListString.String() on a single element tailed list is incorrect: want '%s', got '%s'", want, got)
		}
	})

	t.Run("Single Element Tailed List Head", func(t *testing.T) {
		want := "a"
		got := tl.head.String()

		if got != want {
			t.Errorf("Head of the single element tailed list is incorrect: want '%s', got '%s'", want, got)
		}
	})

	t.Run("Single Element Tailed List Tail", func(t *testing.T) {
		want := "a"
		got := tl.tail.String()

		if got != want {
			t.Errorf("Tail of the single element tailed list is incorrect: want '%s', got '%s'", want, got)
		}
	})
}

func TestIsEmpty(t *testing.T) {
	tl := TailedList{}

	t.Run("Is Empty True", func(t *testing.T) {

		want := true
		got := tl.IsEmpty()

		if got != want {
			t.Errorf("Is Empty on an empty list failed: want '%v', got '%v'", want, got)
		}
	})

	tl = constructListFromNode(&Node{"a", nil})
	t.Run("Is Empty False", func(t *testing.T) {

		want := false
		got := tl.IsEmpty()

		if got != want {
			t.Errorf("Is Empty on a non-empty list failed: want '%v', got '%v'", want, got)
		}
	})
}

func TestAddAtEnd(t *testing.T) {
	tl := &TailedList{}

	tl.AddAtEnd("a")
	t.Run("Add At End Empty List", func(t *testing.T) {

		want := "a->nil"
		got := tl.String()
		if got != want {
			t.Errorf("Adding a value at the end of an empty list fails: want '%s', got '%s'", want, got)
		}

		want = "a"
		got = tl.head.String()
		if got != want {
			t.Errorf("New Head on adding a value at the end of an empty list fails: want '%s', got '%s'", want, got)
		}

		want = "a"
		got = tl.tail.String()
		if got != want {
			t.Errorf("New Tail on adding a value at the end of an empty list fails: want '%s', got '%s'", want, got)
		}
	})

	tl.AddAtEnd("b")
	t.Run("Add At End Non-Empty List", func(t *testing.T) {

		want := "a->b->nil"
		got := tl.String()
		if got != want {
			t.Errorf("Adding a value at the end of a non-empty list fails: want '%s', got '%s'", want, got)
		}

		want = "a"
		got = tl.head.String()
		if got != want {
			t.Errorf("New Head on adding a value at the end of a non-empty list fails: want '%s', got '%s'", want, got)
		}

		want = "b"
		got = tl.tail.String()
		if got != want {
			t.Errorf("New Tail on adding a value at the end of a non-empty list fails: want '%s', got '%s'", want, got)
		}
	})
}

func TestAddAtBeginningOnAnEmptyList(t *testing.T) {
	tl := &TailedList{}
	tl.AddAtBeginning("a")

	want := "a->nil"
	got := tl.String()
	if got != want {
		t.Errorf("Adding a value at the beginning of an empty list fails: want '%s', got '%s'", want, got)
	}

	want = "a"
	got = tl.head.String()
	if got != want {
		t.Errorf("New Head on adding a value at the beginning of an empty list fails: want '%s', got '%s'", want, got)
	}

	want = "a"
	got = tl.tail.String()
	if got != want {
		t.Errorf("New Tail on adding a value at the beginning of an empty list fails: want '%s', got '%s'", want, got)
	}
}

func TestAddAtBeginningOnANonEmptyList(t *testing.T) {
	tl := constructListFromNode(&Node{"a", nil})
	tl.AddAtBeginning("b")

	want := "b->a->nil"
	got := tl.String()
	if got != want {
		t.Errorf("Adding a value at the beginning of a non-empty list fails: want '%s', got '%s'", want, got)
	}

	want = "b"
	got = tl.head.String()
	if got != want {
		t.Errorf("New Head on adding a value at the beginning of a non-empty list fails: want '%s', got '%s'", want, got)
	}

	want = "a"
	got = tl.tail.String()
	if got != want {
		t.Errorf("New Tail on adding a value at the beginning of a non-empty list fails: want '%s', got '%s'", want, got)
	}
}

func TestRemoveValueOnEmptyList(t *testing.T) {
	tl := &TailedList{}

	err := tl.RemoveValue("a")
	if err == nil {
		t.Errorf("Removing a value from an empty list should have returned an error")
	} else {
		fmt.Println(err)
	}
}

func TestRemoveValidValueOnSingleElementList(t *testing.T) {
	n := &Node{"a", nil}
	tl := TailedList{n, n}

	err := tl.RemoveValue("a")
	if err != nil {
		t.Errorf("Removal failed on: %v", err)
	} else {
		want := "nil"
		got := tl.String()

		if want != got {
			t.Errorf("Removing a valid value from an empty list failed, want %v, got %v", want, got)
		}
	}
}

func TestRemoveInvalidValueOnSingleElementList(t *testing.T) {
	n := &Node{"a", nil}
	tl := TailedList{n, n}

	err := tl.RemoveValue("b")
	if err == nil {
		t.Errorf("Removing an invalid value from an empty list should have returned an error")
	} else {
		fmt.Println(err)
	}
}

func TestRemoveFirstOnEmptyList(t *testing.T) {
	tl := TailedList{}

	_, err := tl.RemoveFirst()
	if err == nil {
		t.Errorf("Removing first element from an empty list should have returned an error")
	} else {
		fmt.Println(err)
	}
}

func TestRemoveFirstOnSingleElementList(t *testing.T) {
	tl := TailedList{}
	tl.AddAtEnd("a")

	first, err := tl.RemoveFirst()
	if err != nil {
		t.Errorf("Removal failed on: %v", err)
	} else {
		want := "a"
		got := first.String()

		if want != got {
			t.Errorf("Removing the first element from a single element list failed, want %v, got %v", want, got)
		}
	}

	if !tl.IsEmpty() {
		t.Errorf("The list should be empty after removing the only element in it")
	}
}

func TestRemoveFirstOnList(t *testing.T) {
	tl := TailedList{}
	tl.AddAtEnd("a")
	tl.AddAtEnd("b")
	tl.AddAtEnd("c")

	removed := []string{"a", "b", "c"}
	partial := []string{"b->c->nil", "c->nil", "nil"}

	for i := 0; !tl.IsEmpty(); i++ {
		first, err := tl.RemoveFirst()
		if err != nil {
			t.Errorf("Removal failed on: %v", err)
		} else {
			wantNode := removed[i]
			gotNode := first.String()

			if wantNode != gotNode {
				t.Errorf("Element removed from the list is incorrect, want %v, got %v", wantNode, gotNode)
			}

			wantList := partial[i]
			gotList := tl.String()

			if wantList != gotList {
				t.Errorf("Remaining list after removing first element is incorrect, want %v, got %v", wantList, gotList)
			}
		}
	}

	if !tl.IsEmpty() {
		t.Errorf("The list should be empty after removing the only element in it")
	}
}

func TestRemoveLastOnEmptyList(t *testing.T) {
	tl := TailedList{}

	_, err := tl.RemoveLast()
	if err == nil {
		t.Errorf("Removing last element from an empty list should have returned an error")
	} else {
		fmt.Println(err)
	}
}

func TestRemoveLastOnSingleElementList(t *testing.T) {
	n := &Node{"a", nil}
	tl := constructListFromNode(n)

	last, err := tl.RemoveLast()
	if err != nil {
		t.Errorf("Removal failed on: %v", err)
	} else {
		want := "a"
		got := last.String()

		if want != got {
			t.Errorf("Removing the last element from a single element list failed, want %v, got %v", want, got)
		}
	}

	if !tl.IsEmpty() {
		t.Errorf("The list should be empty after removing the only element in it")
	}
}

func TestRemoveLastOnList(t *testing.T) {
	tl := TailedList{}
	tl.AddAtBeginning("a")
	tl.AddAtBeginning("b")
	tl.AddAtBeginning("c")
	tl.AddAtBeginning("d")
	tl.AddAtBeginning("e")

	removed := []string{"a", "b", "c", "d", "e"}
	partial := []string{"e->d->c->b->nil", "e->d->c->nil", "e->d->nil", "e->nil", "nil"}

	for i := 0; !tl.IsEmpty(); i++ {
		last, err := tl.RemoveLast()
		if err != nil {
			t.Errorf("Removal failed on: %v", err)
		} else {
			wantNode := removed[i]
			gotNode := last.String()

			if wantNode != gotNode {
				t.Errorf("Element removed from the list is incorrect, want %v, got %v", wantNode, gotNode)
			}

			wantList := partial[i]
			gotList := tl.String()

			if wantList != gotList {
				t.Errorf("Remaining list after removing last element is incorrect, want %v, got %v", wantList, gotList)
			}
		}
	}

	if !tl.IsEmpty() {
		t.Errorf("The list should be empty after removing the only element in it")
	}
}

func TestReverseEmptyList(t *testing.T) {
	tl := &TailedList{}

	err := tl.Reverse()
	if err == nil {
		t.Errorf("Attempt to reverse an empty list should have failed")
	} else {
		fmt.Println(err)
	}
}

func TestReverseSingleElementList(t *testing.T) {
	tl := &TailedList{}
	tl.AddAtBeginning("a")

	err := tl.Reverse()
	if err != nil {
		t.Errorf("Reverse failed on: %v", err)
	} else {
		want := "a->nil"
		got := tl.String()
		if want != got {
			t.Errorf("Reversing a single element list failed, want %v, got %v", want, got)
		}

		want = "a"
		got = tl.head.String()
		if want != got {
			t.Errorf("New head after reversing a single element list is incorrect, want %v, got %v", want, got)
		}

		want = "a"
		got = tl.tail.String()
		if want != got {
			t.Errorf("New tail after reversing a single element list is incorrect, want %v, got %v", want, got)
		}
	}
}

func TestReverseList(t *testing.T) {
	tl := &TailedList{}
	tl.AddAtEnd("a")
	tl.AddAtEnd("b")
	tl.AddAtEnd("c")
	tl.AddAtEnd("d")
	tl.AddAtEnd("e")

	err := tl.Reverse()
	if err != nil {
		t.Errorf("Reverse failed on: %v", err)
	} else {
		want := "e->d->c->b->a->nil"
		got := tl.String()
		if want != got {
			t.Errorf("Reversing a tailed list failed, want %v, got %v", want, got)
		}

		want = "e"
		got = tl.head.String()
		if want != got {
			t.Errorf("New head after reversing a single element list is incorrect, want %v, got %v", want, got)
		}

		want = "a"
		got = tl.tail.String()
		if want != got {
			t.Errorf("New tail after reversing a single element list is incorrect, want %v, got %v", want, got)
		}
	}
}

func TestReverseListTillEmpty(t *testing.T) {
	tl := &TailedList{}
	tl.AddAtEnd("a")
	tl.AddAtEnd("b")
	tl.AddAtEnd("c")
	tl.AddAtEnd("d")
	tl.AddAtEnd("e")

	partials := []string{"e->d->c->b->a->nil", "a->b->c->d->nil", "d->c->b->nil", "b->c->nil", "c->nil"}
	heads := []string{"e", "a", "d", "b", "c"}
	tails := []string{"a", "d", "b", "c", "c"}

	for i := 0; !tl.IsEmpty(); i++ {
		err := tl.Reverse()
		if err != nil {
			t.Fatalf("Reverse failed on: %v", err)
		} else {
			want := partials[i]
			got := tl.String()
			if want != got {
				t.Errorf("Reversing a tailed list failed, want %v, got %v", want, got)
			}

			want = heads[i]
			got = tl.head.String()
			if want != got {
				t.Errorf("New head after reversing a single element list is incorrect, want %v, got %v", want, got)
			}

			want = tails[i]
			got = tl.tail.String()
			if want != got {
				t.Errorf("New tail after reversing a single element list is incorrect, want %v, got %v", want, got)
			}
		}
		tl.RemoveFirst()
	}
}

func TestCopyEmptyList(t *testing.T) {
	tl := &TailedList{}
	tl2 := tl.Copy()

	want := "nil"
	got := tl2.String()

	if want != got {
		t.Errorf("Copying an empty tailed list failed, want %v, got %v", want, got)
	}
}

func TestCopySingleElementList(t *testing.T) {
	tl := constructListFromNode(&Node{"a", nil})
	tl2 := tl.Copy()

	want := "a->nil"
	got := tl2.String()

	if want != got {
		t.Errorf("Copying a single element tailed list failed, want %v, got %v", want, got)
	}

	want = "a"
	got = tl2.head.String()

	if want != got {
		t.Errorf("Head after copying a single element tailed list is incorrect, want %v, got %v", want, got)
	}

	want = "a"
	got = tl2.tail.String()

	if want != got {
		t.Errorf("Tail after copying a single element tailed list is incorrect, want %v, got %v", want, got)
	}

	if tl2.head != tl2.tail {
		t.Errorf("Head and Tail of a single element list should be the same node")
	}
}

func TestCopyList(t *testing.T) {
	tl := &TailedList{}
	data := []string{"a", "b", "c", "d", "e"}
	partials := []string{"a->nil", "a->b->nil", "a->b->c->nil", "a->b->c->d->nil", "a->b->c->d->e->nil"}
	heads := []string{"a", "a", "a", "a", "a"}
	tails := []string{"a", "b", "c", "d", "e"}

	for i := 0; i < len(data); i++ {
		tl.AddAtEnd(data[i])
		tl2 := tl.Copy()

		want := partials[i]
		got := tl2.String()
		if want != got {
			t.Errorf("Copying a list failed, want %v, got %v", want, got)
		}

		want = heads[i]
		got = tl.head.String()
		if want != got {
			t.Errorf("New head after copying a list is incorrect, want %v, got %v", want, got)
		}

		want = tails[i]
		got = tl.tail.String()
		if want != got {
			t.Errorf("New tail after copying a list is incorrect, want %v, got %v", want, got)
		}
	}
}
