package linkedListLibrary

import (
	"strings"
	"testing"
)

func TestStringForm(t *testing.T) {
	list := constructList(constructNode("1"))
	want := "1->nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Errorf("GetStringForm Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestAddToEnd(t *testing.T) {
	list := constructList(constructNode("1"))
	list.AddToEnd("2")
	want := "1->2->nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Errorf("AddToEnd Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestAddToBeginning(t *testing.T) {
	list := constructList(constructNode("1"))
	list.AddToBeginning("2")
	want := "2->1->nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Errorf("AddToBeginning Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestReverse(t *testing.T) {
	list := constructList(constructNode("2"))
	list.AddToEnd("1")
	list.AddToBeginning("3")
	list.Reverse()
	want := "1->2->3->nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Errorf("Reverse Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestConstructFromValues(t *testing.T) {
	list1 := constructList(constructNode("t"))
	list1.AddToBeginning("a")
	list1.AddToEnd("e")
	list1.AddToBeginning("w")
	list1.AddToEnd("r")

	list2 := ConstructFromValues("w", "a", "t", "e", "r")

	want := list1.GetStringForm()
	have := list2.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("ConstructFromValues Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestConstructEmptyList(t *testing.T) {
	list := ConstructFromValues()

	want := "nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Errorf("ConstructFromValues Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestRemoveInvalidValue(t *testing.T) {
	list := ConstructFromValues("bat", "cat", "mat")
	want := list.GetStringForm()

	list.RemoveValue("rat")
	have := list.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestRemoveHeadValue(t *testing.T) {
	list1 := ConstructFromValues("w", "h", "a", "t")
	list1.RemoveValue("w")

	list2 := ConstructFromValues("h", "a", "t")

	want := list1.GetStringForm()
	have := list2.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestRemoveLastValue(t *testing.T) {
	list1 := ConstructFromValues("p", "a", "t", "h")
	list1.RemoveValue("h")

	list2 := ConstructFromValues("p", "a", "t")

	want := list2.GetStringForm()
	have := list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestRemoveOnlyValue(t *testing.T) {
	list1 := ConstructFromValues("1")
	list1.RemoveValue("1")

	list2 := ConstructFromValues()

	want := list2.GetStringForm()
	have := list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestRemoveAllValues(t *testing.T) {
	list1 := ConstructFromValues("t", "o", "w", "e", "r")
	list1.RemoveValue("o")
	list1.RemoveValue("e")

	list2 := ConstructFromValues("t", "w", "r")

	want := list2.GetStringForm()
	have := list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Fatalf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}

	list1.RemoveValue("r")
	list1.RemoveValue("t")

	list3 := ConstructFromValues("w")

	want = list3.GetStringForm()
	have = list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Fatalf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}

	list1.RemoveValue("w")

	list4 := ConstructFromValues()

	want = list4.GetStringForm()
	have = list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Fatalf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}

}
