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

func TestSingleElementTailedListString(t *testing.T) {
	tl := constructListFromNode(&Node{"a", nil})

	want := "a->nil"
	got := tl.String()

	if got != want {
		t.Errorf("*TailedListString.String() on a single element tailed list is incorrect: want '%s', got '%s'", want, got)
	}
}

func TestSingleElementTailedListHead(t *testing.T) {
	tl := constructListFromNode(&Node{"a", nil})

	want := "a"
	got := tl.head.String()

	if got != want {
		t.Errorf("Head of the single element tailed list is incorrect: want '%s', got '%s'", want, got)
	}
}

func TestSingleElementTailedListTail(t *testing.T) {
	tl := constructListFromNode(&Node{"a", nil})

	want := "a"
	got := tl.tail.String()

	if got != want {
		t.Errorf("Tail of the single element tailed list is incorrect: want '%s', got '%s'", want, got)
	}
}

func TestIsEmptyTrue(t *testing.T) {
	tl := &TailedList{}

	want := true
	got := tl.IsEmpty()

	if got != want {
		t.Errorf("Is Empty on an empty list failed: want '%v', got '%v'", want, got)
	}
}

func TestIsEmptyFalse(t *testing.T) {
	tl := constructListFromNode(&Node{"a", nil})

	want := false
	got := tl.IsEmpty()

	if got != want {
		t.Errorf("Is Empty on an empty list failed: want '%v', got '%v'", want, got)
	}
}

func TestAddAtEndOnAnEmptyList(t *testing.T) {
	tl := &TailedList{}
	tl.AddAtEnd("a")

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
}

func TestAddAtEndOnANonEmptyList(t *testing.T) {
	tl := constructListFromNode(&Node{"a", nil})
	tl.AddAtEnd("b")

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

