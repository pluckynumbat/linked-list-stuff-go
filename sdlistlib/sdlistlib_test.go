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

