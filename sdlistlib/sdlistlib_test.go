package sdlistlib

import (
	"fmt"
	"testing"
)

type prInt int // printable int!

func (p prInt) String() string {
	return fmt.Sprintf("%v", int(p))
}

func TestNodeString(t *testing.T) {

	var n1, n2 *Node[prInt]
	n2 = &Node[prInt]{nil, 1, nil}

	tests := []struct {
		name string
		node *Node[prInt]
		want string
	}{
		{"nil node : test string", n1, "nil"},
		{"non-nil node : test string", n2, "1"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.want
			got := test.node.String()
			if got != want {
				t.Errorf("Got incorrect results for String(), want: %v, got: %v", want, got)
			}
		})
	}
}

func TestListString(t *testing.T) {

	var l1, l2, l3, l4 *SemiGenericList[prInt]

	l2 = &SemiGenericList[prInt]{}

	node := &Node[prInt]{nil, 1, nil}
	l3 = &SemiGenericList[prInt]{node, node}

	var n1, n2 *Node[prInt]
	n1 = &Node[prInt]{nil, 1, nil}
	n2 = &Node[prInt]{nil, 2, nil}
	n1.next = n2
	n2.prev = n1
	l4 = &SemiGenericList[prInt]{n1, n2}

	tests := []struct {
		name string
		list *SemiGenericList[prInt]
		want string
	}{
		{"nil list : test string", l1, "nil"},
		{"empty list : test string", l2, "empty"},
		{"1 element list : test string", l3, "nil<-1->nil"},
		{"2 element list : test string", l4, "nil<-1<=>2->nil"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.want
			got := test.list.String()
			if got != want {
				t.Errorf("Got incorrect results for String(), want: %v, got: %v", want, got)
			}
		})
	}
}

func TestListStates(t *testing.T) {
	var l1, l2, l3, l4, l5 *SemiGenericList[prInt]
	l2 = &SemiGenericList[prInt]{}
	l3 = new(SemiGenericList[prInt])

	node := &Node[prInt]{nil, 1, nil}

	l4 = &SemiGenericList[prInt]{}
	l4.head = node
	l4.tail = node

	l5 = new(SemiGenericList[prInt])
	l5.head = node
	l5.tail = node

	tests := []struct {
		name string
		list *SemiGenericList[prInt]
		fn   func() bool
		want bool
	}{
		{"nil list : test nil", l1, l1.IsNil, true},
		{"struct literal empty list : test nil", l2, l2.IsNil, false},
		{"struct literal empty list : test empty", l2, l2.IsEmpty, true},
		{"new allocator empty list : test nil", l3, l3.IsNil, false},
		{"new allocator empty list : test empty", l3, l3.IsEmpty, true},
		{"struct literal non-empty list : test nil", l4, l4.IsNil, false},
		{"struct literal non-empty list : test empty", l4, l4.IsEmpty, false},
		{"new allocator non-empty list : test nil", l3, l3.IsNil, false},
		{"new allocator non-empty list : test empty", l5, l5.IsEmpty, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.want
			got := test.fn()
			if got != want {
				t.Errorf("Got incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}

func TestHead(t *testing.T) {

	var l1, l2, l3, l4, l5 *SemiGenericList[prInt]
	l2 = &SemiGenericList[prInt]{}

	n1 := &Node[prInt]{nil, 1, nil}
	l3 = &SemiGenericList[prInt]{n1, n1}

	n2 := &Node[prInt]{nil, 2, nil}
	n1.next = n2
	n2.prev = n1
	l4 = &SemiGenericList[prInt]{n1, n2}

	n3 := &Node[prInt]{nil, 3, nil}
	n2.next = n3
	n3.prev = n2
	l5 = &SemiGenericList[prInt]{n1, n3}

	tests := []struct {
		name string
		list *SemiGenericList[prInt]
		want string
	}{
		{"nil list : test head", l1, "nil"},
		{"empty list : test head", l2, "nil"},
		{"1 element list : test head", l3, "1"},
		{"2 element list : test head", l4, "1"},
		{"3 element list : test head", l5, "1"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.want
			got := test.list.Head().String()
			if got != want {
				t.Errorf("Got incorrect results for Head(), want: %v, got: %v", want, got)
			}
		})
	}
}

func TestTail(t *testing.T) {

	var l1, l2, l3, l4, l5 *SemiGenericList[prInt]
	l2 = &SemiGenericList[prInt]{}

	n1 := &Node[prInt]{nil, 1, nil}
	l3 = &SemiGenericList[prInt]{n1, n1}

	n2 := &Node[prInt]{nil, 2, nil}
	n1.next = n2
	n2.prev = n1
	l4 = &SemiGenericList[prInt]{n1, n2}

	n3 := &Node[prInt]{nil, 3, nil}
	n2.next = n3
	n3.prev = n2
	l5 = &SemiGenericList[prInt]{n1, n3}

	tests := []struct {
		name string
		list *SemiGenericList[prInt]
		want string
	}{
		{"nil list : test head", l1, "nil"},
		{"empty list : test head", l2, "nil"},
		{"1 element list : test head", l3, "1"},
		{"2 element list : test head", l4, "2"},
		{"3 element list : test head", l5, "3"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.want
			got := test.list.Tail().String()
			if got != want {
				t.Errorf("Got incorrect results for Head(), want: %v, got: %v", want, got)
			}
		})
	}
}

func TestAddAtBeginning(t *testing.T) {

	var list *SemiGenericList[prInt]

	t.Run("nil list", func(t *testing.T) {
		err := list.AddAtBeginning(1)
		if err == nil {
			t.Errorf("AddAtBeginning() on a nil list should return an error")
		} else {
			fmt.Println(err)
		}
	})

	list = &SemiGenericList[prInt]{}

	var tests = []struct {
		name    string
		val     prInt
		expList string
		expHead string
		expTail string
	}{
		{"add to empty list", 1, "nil<-1->nil", "1", "1"},
		{"add to 1 element list", 2, "nil<-2<=>1->nil", "2", "1"},
		{"add to 2 element list", 3, "nil<-3<=>2<=>1->nil", "3", "1"},
		{"add to 3 element list", 4, "nil<-4<=>3<=>2<=>1->nil", "4", "1"},
		{"add to 4 element list", 5, "nil<-5<=>4<=>3<=>2<=>1->nil", "5", "1"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := list.AddAtBeginning(test.val)
			if err != nil {
				t.Errorf("AddAtBeginning() failed with error: %v", err)
			} else {

				got := list.String()
				want := test.expList
				if got != want {
					t.Errorf("list after adding to it is incorrect, want: %v, got: %v", want, got)
				}

				got = list.Head().String()
				want = test.expHead
				if got != want {
					t.Errorf("head after adding to the list is incorrect, want: %v, got: %v", want, got)
				}

				got = list.Tail().String()
				want = test.expTail
				if got != want {
					t.Errorf("tail after adding to the list is incorrect, want: %v, got: %v", want, got)
				}
			}
		})
	}
}

func TestAddAtEnd(t *testing.T) {

	var list *SemiGenericList[prInt]

	t.Run("nil list", func(t *testing.T) {
		err := list.AddAtEnd(1)
		if err == nil {
			t.Errorf("AddAtEnd() on a nil list should return an error")
		} else {
			fmt.Println(err)
		}
	})

	list = &SemiGenericList[prInt]{}

	var tests = []struct {
		name    string
		val     prInt
		expList string
		expHead string
		expTail string
	}{
		{"add to empty list", 1, "nil<-1->nil", "1", "1"},
		{"add to 1 element list", 2, "nil<-1<=>2->nil", "1", "2"},
		{"add to 2 element list", 3, "nil<-1<=>2<=>3->nil", "1", "3"},
		{"add to 3 element list", 4, "nil<-1<=>2<=>3<=>4->nil", "1", "4"},
		{"add to 4 element list", 5, "nil<-1<=>2<=>3<=>4<=>5->nil", "1", "5"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := list.AddAtEnd(test.val)
			if err != nil {
				t.Errorf("AddAtEnd() failed with error: %v", err)
			} else {

				got := list.String()
				want := test.expList
				if got != want {
					t.Errorf("list after adding to it is incorrect, want: %v, got: %v", want, got)
				}

				got = list.Head().String()
				want = test.expHead
				if got != want {
					t.Errorf("head after adding to the list is incorrect, want: %v, got: %v", want, got)
				}

				got = list.Tail().String()
				want = test.expTail
				if got != want {
					t.Errorf("tail after adding to the list is incorrect, want: %v, got: %v", want, got)
				}
			}
		})
	}
}
