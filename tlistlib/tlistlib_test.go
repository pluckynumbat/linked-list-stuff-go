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

