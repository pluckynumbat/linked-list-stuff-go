package tlistlib

import (
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

func TestAddToEndOnAnEmptyList(t *testing.T) {
	tl := &TailedList{}
	tl.AddToEnd("a")

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

