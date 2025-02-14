package listlib

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
	list := &LinkedList{}
	list.AddToEnd("1")
	list.AddToEnd("2")
	want := "1->2->nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Errorf("AddToEnd Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestAddToBeginning(t *testing.T) {
	list := &LinkedList{}
	list.AddToBeginning("1")
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

func TestIsEmptyTrue(t *testing.T) {
	list := ConstructFromValues()

	want := true
	have := list.IsEmpty()
	if want != have {
		t.Errorf("ConstructFromValues Fails: wanted '%t', got '%t'", want, have)
	}
}

func TestIsEmptyFalse(t *testing.T) {
	list := ConstructFromValues("hi")

	want := false
	have := list.IsEmpty()
	if want != have {
		t.Errorf("ConstructFromValues Fails: wanted '%t', got '%t'", want, have)
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

func TestRemoveAtEnd(t *testing.T) {
	list1 := ConstructFromValues("p", "a", "t", "h")
	list1.RemoveAtEnd()

	list2 := ConstructFromValues("p", "a", "t")

	want := list2.GetStringForm()
	have := list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestRemoveAtEndOnlyOneLeft(t *testing.T) {
	list1 := ConstructFromValues("a")
	list1.RemoveAtEnd()

	list2 := ConstructFromValues()

	want := list2.GetStringForm()
	have := list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestRemoveAtEndEmptyList(t *testing.T) {
	list1 := ConstructFromValues()
	list1.RemoveAtEnd()

	list2 := ConstructFromValues()

	want := list2.GetStringForm()
	have := list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestRemoveAtEndTillEmpty(t *testing.T) {
	list1 := ConstructFromValues("w", "h", "a", "t")

	partial := [...]LinkedList{ConstructFromValues("w", "h", "a"), ConstructFromValues("w", "h"), ConstructFromValues("w"), ConstructFromValues()}

	i := 0
	for i = 0; !list1.IsEmpty(); i++ {
		list1.RemoveAtEnd()

		want := partial[i].GetStringForm()
		have := list1.GetStringForm()

		if strings.Compare(want, have) != 0 {
			t.Fatalf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
		}
	}

	//also try removing from empty
	list1.RemoveAtEnd()

	list2 := ConstructFromValues()
	want := list2.GetStringForm()
	have := list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Fatalf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}

}

func TestRemoveAtBeginning(t *testing.T) {
	list1 := ConstructFromValues("w", "h", "a", "t")
	list1.RemoveAtBeginning()

	list2 := ConstructFromValues("h", "a", "t")

	want := list2.GetStringForm()
	have := list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestRemoveAtBeginningOnlyOneLeft(t *testing.T) {
	list1 := ConstructFromValues("a")
	list1.RemoveAtBeginning()

	list2 := ConstructFromValues()

	want := list2.GetStringForm()
	have := list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestRemoveAtBeginningEmptyList(t *testing.T) {
	list1 := ConstructFromValues()
	list1.RemoveAtBeginning()

	list2 := ConstructFromValues()

	want := list2.GetStringForm()
	have := list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestRemoveAtBeginningTillEmpty(t *testing.T) {
	list1 := ConstructFromValues("w", "h", "a", "t")

	partial := [...]LinkedList{
		ConstructFromValues("h", "a", "t"),
		ConstructFromValues("a", "t"),
		ConstructFromValues("t"),
		ConstructFromValues(),
	}

	i := 0
	for i = 0; !list1.IsEmpty(); i++ {
		list1.RemoveAtBeginning()

		want := partial[i].GetStringForm()
		have := list1.GetStringForm()

		if strings.Compare(want, have) != 0 {
			t.Fatalf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
		}
	}

	//also try removing from empty
	list1.RemoveAtBeginning()

	list2 := ConstructFromValues()
	want := list2.GetStringForm()
	have := list1.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Fatalf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
	}

}

func TestCopy(t *testing.T) {

	cases := []LinkedList{
		ConstructFromValues("w", "h", "a", "t"),
		ConstructFromValues("a"),
		ConstructFromValues(),
	}

	for i := 0; i < len(cases); i++ {

		list := cases[i].Copy()

		want := cases[i].GetStringForm()

		have := list.GetStringForm()

		if strings.Compare(want, have) != 0 {
			t.Fatalf("RemoveValue Fails: wanted '%s', got '%s'", want, have)
		}
	}
}

func TestRemovedAtEndValues(t *testing.T) {
	list := ConstructFromValues("w", "h", "a", "t")
	removed := []string{"t", "a", "h", "w"}

	for i := 0; !list.IsEmpty(); i++ {
		n := list.RemoveAtEnd()

		have := n.data
		want := removed[i]

		if strings.Compare(want, have) != 0 {
			t.Fatalf("RemoveAtEnd Fails: wanted '%s', got '%s'", want, have)
		}
	}

	n := list.RemoveAtEnd()

	if n != nil {
		t.Fatalf("RemoveAtEnd Fails: wanted nil, got %v", n)
	}
}

func TestRemovedAtBeginningValues(t *testing.T) {
	list := ConstructFromValues("w", "h", "a", "t")
	removed := []string{"w", "h", "a", "t"}

	for i := 0; !list.IsEmpty(); i++ {
		n := list.RemoveAtBeginning()

		have := n.data
		want := removed[i]

		if strings.Compare(want, have) != 0 {
			t.Fatalf("RemoveAtBeginning Fails: wanted '%s', got '%s'", want, have)
		}
	}

	n := list.RemoveAtBeginning()

	if n != nil {
		t.Fatalf("RemoveAtBeginning Fails: wanted nil, got %v", n)
	}
}

func TestNodeGetData(t *testing.T) {
	list := ConstructFromValues("w", "h", "a", "t")
	removed := []string{"t", "a", "h", "w"}

	for i := 0; !list.IsEmpty(); i++ {
		n := list.RemoveAtEnd()

		have, err := n.GetData()
		if err != nil {
			t.Fatalf("GetData returned an error: %v", err)
		}
		want := removed[i]

		if strings.Compare(want, have) != 0 {
			t.Fatalf("GetData Fails: wanted '%s', got '%s'", want, have)
		}
	}

	n := list.RemoveAtEnd()
	have, err := n.GetData()

	if err == nil || strings.Compare(have, "") != 0 {
		t.Fatalf("GetData should have returned an error on a nil node")
	}
}

func TestIsNil(t *testing.T) {
	var nilList *LinkedList

	var tests = []struct {
		name  string
		input *LinkedList
		want  bool
	}{
		{"nil list pointer", nilList, true},
		{"new allocator empty list", new(LinkedList), false},
		{"composite literal empty list", &LinkedList{}, false},
		{"composite literal  non-empty list", &LinkedList{&Node{"a", nil}}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.input.IsNil()
			if got != test.want {
				t.Errorf("IsNil gave incorrect results, want: %v, got: %v", test.want, got)
			}
		})
	}
}

func TestHead(t *testing.T) {

	var nilList *LinkedList

	oneNodeList := ConstructFromValues("a")
	twoNodeList := ConstructFromValues("a", "b")
	threeNodeList := ConstructFromValues("a", "b", "c")

	var tests = []struct {
		name  string
		input *LinkedList
		want  string
	}{
		{"nil list pointer", nilList, ""},
		{"new allocator empty list", new(LinkedList), ""},
		{"composite literal empty list", &LinkedList{}, ""},
		{"composite literal single element list", &oneNodeList, "a"},
		{"composite literal 2 element list", &twoNodeList, "a"},
		{"composite literal 3 element list", &threeNodeList, "a"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.input.Head().GetData()
			if err != nil {
				if test.input.Head() != nil { // if the Head is nil, error is expected
					t.Fatalf("Error while getting data from head of the list, want: %v, got: %v", test.want, got)
				}
			}
			if got != test.want {
				t.Errorf("Head gave incorrect results, want: %v, got: %v", test.want, got)
			}
		})
	}
}
