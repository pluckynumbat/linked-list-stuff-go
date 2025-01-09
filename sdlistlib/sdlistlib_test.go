package sdlistlib

import (
	"fmt"
	"testing"
)

type prInt int       // printable int!
type prBool bool     // printable bool
type prString string // printable string, yeah

func (p prInt) String() string {
	return fmt.Sprintf("%v", int(p))
}

func (p prBool) String() string {
	return fmt.Sprintf("%v", bool(p))
}

func (p prString) String() string {
	return fmt.Sprintf("%v", string(p))
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

func TestReverse(t *testing.T) {

	var list *SemiGenericList[prInt]

	t.Run("nil list", func(t *testing.T) {
		err := list.Reverse()
		if err == nil {
			t.Errorf("Reverse() on a nil list should return an error")
		} else {
			fmt.Println(err)
		}
	})

	list = &SemiGenericList[prInt]{}

	t.Run("empty list", func(t *testing.T) {
		err := list.Reverse()
		if err != nil {
			t.Errorf("Reverse() failed with error: %v", err)
		} else {
			want := "empty"
			got := list.String()

			if got != want {
				t.Errorf("Reverse() gave incorrect results, want: %v, got :%v", want, got)
			}
		}
	})

	var tests = []struct {
		name    string
		val     prInt
		expList string
		expHead string
		expTail string
	}{
		{"reverse 1 element list", 1, "nil<-1->nil", "1", "1"},
		{"reverse 2 element list", 2, "nil<-2<=>1->nil", "2", "1"},
		{"reverse 3 element list", 3, "nil<-3<=>1<=>2->nil", "3", "2"},
		{"reverse 4 element list", 4, "nil<-4<=>2<=>1<=>3->nil", "4", "3"},
		{"reverse 5 element list", 5, "nil<-5<=>3<=>1<=>2<=>4->nil", "5", "4"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := list.AddAtEnd(test.val)
			if err != nil {
				t.Errorf("AddAtEnd() failed with error: %v", err)
			} else {

				err2 := list.Reverse()
				if err2 != nil {
					t.Errorf("Reverse() failed with error: %v", err2)
				} else {

					got := list.String()
					want := test.expList
					if got != want {
						t.Errorf("list after reversing is incorrect, want: %v, got: %v", want, got)
					}

					head := list.Head()
					got = head.String()
					want = test.expHead
					if got != want {
						t.Errorf("head after reversing the list is incorrect, want: %v, got: %v", want, got)
					}

					if head != nil {
						headPrev := head.prev
						if headPrev != nil {
							t.Errorf("head's prev pointer should always be nil, got: %v", headPrev)
						}
					}

					tail := list.Tail()
					got = tail.String()
					want = test.expTail
					if got != want {
						t.Errorf("tail after reversing the list is incorrect, want: %v, got: %v", want, got)
					}

					if tail != nil {
						tailNext := tail.next
						if tailNext != nil {
							t.Errorf("head's prev pointer should always be nil, got: %v", tailNext)
						}
					}
				}
			}
		})
	}
}

func TestCopy(t *testing.T) {

	var list *SemiGenericList[prInt]

	t.Run("nil list", func(t *testing.T) {
		_, err := list.Copy()
		if err == nil {
			t.Errorf("Copy() on a nil list should return an error")
		} else {
			fmt.Println(err)
		}
	})

	list = &SemiGenericList[prInt]{}

	t.Run("empty list", func(t *testing.T) {
		copy, err := list.Copy()
		if err != nil {
			t.Errorf("Copy() failed with error: %v", err)
		} else {
			want := "empty"
			got := copy.String()

			if got != want {
				t.Errorf("Reverse() gave incorrect results, want: %v, got :%v", want, got)
			}
		}
	})

	var tests = []struct {
		name    string
		val     prInt
		expList string
		expHead string
		expTail string
	}{
		{"copy 1 element list", 1, "nil<-1->nil", "1", "1"},
		{"copy 2 element list", 2, "nil<-2<=>1->nil", "2", "1"},
		{"copy 3 element list", 3, "nil<-3<=>2<=>1->nil", "3", "1"},
		{"copy 4 element list", 4, "nil<-4<=>3<=>2<=>1->nil", "4", "1"},
		{"copy 5 element list", 5, "nil<-5<=>4<=>3<=>2<=>1->nil", "5", "1"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := list.AddAtBeginning(test.val)
			if err != nil {
				t.Errorf("AddAtBeginning() failed with error: %v", err)
			} else {

				copy, err2 := list.Copy()
				if err2 != nil {
					t.Errorf("Copy() failed with error: %v", err2)
				} else {

					got := copy.String()
					want := test.expList
					if got != want {
						t.Errorf("copied list is incorrect, want: %v, got: %v", want, got)
					}

					head := copy.Head()
					got = head.String()
					want = test.expHead
					if got != want {
						t.Errorf("head of the copied list is incorrect, want: %v, got: %v", want, got)
					}

					if head != nil {
						headPrev := head.prev
						if headPrev != nil {
							t.Errorf("head's prev pointer should always be nil, got: %v", headPrev)
						}
					}

					tail := copy.Tail()
					got = tail.String()
					want = test.expTail
					if got != want {
						t.Errorf("tail of the copied list is incorrect, want: %v, got: %v", want, got)
					}

					if tail != nil {
						tailNext := tail.next
						if tailNext != nil {
							t.Errorf("head's prev pointer should always be nil, got: %v", tailNext)
						}
					}
				}
			}
		})
	}
}

func TestRemoveValue(t *testing.T) {

	var list *SemiGenericList[prInt]
	t.Run("nil list", func(t *testing.T) {
		err := list.RemoveValue(0)
		if err == nil {
			t.Errorf("RemoveValue() on a nil list should return an error")
		} else {
			fmt.Println(err)
		}
	})

	list = &SemiGenericList[prInt]{}
	t.Run("empty list", func(t *testing.T) {
		err := list.RemoveValue(0)
		if err == nil {
			t.Errorf("RemoveValue() on an empty list should return an error")
		} else {
			fmt.Println(err)
		}
	})

	// some setup
	var val prInt
	for val = 1; val <= 5; val++ {
		err := list.AddAtEnd(val)
		if err != nil {
			t.Fatalf("AddAtEnd() failed, error: %v", err)
		}
	}

	var tests = []struct {
		name    string
		value   prInt
		isValid bool
		expList string
		expHead string
		expTail string
		isEmpty bool
	}{
		{"5 element list, invalid value", 6, false, "nil<-1<=>2<=>3<=>4<=>5->nil", "1", "5", false},
		{"5 element list, valid value", 1, true, "nil<-2<=>3<=>4<=>5->nil", "2", "5", false},
		{"4 element list, invalid value", 7, false, "nil<-2<=>3<=>4<=>5->nil", "1", "5", false},
		{"4 element list, valid value", 4, true, "nil<-2<=>3<=>5->nil", "2", "5", false},
		{"3 element list, invalid value", 8, false, "nil<-2<=>3<=>5->nil", "2", "5", false},
		{"3 element list, valid value", 5, true, "nil<-2<=>3->nil", "2", "3", false},
		{"2 element list, invalid value", 9, false, "nil<-2<=>3->nil", "2", "3", false},
		{"2 element list, valid value", 3, true, "nil<-2->nil", "2", "2", false},
		{"1 element list, invalid value", 0, false, "nil<-2->nil", "2", "2", false},
		{"1 element list, valid value", 2, true, "empty", "nil", "nil", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := list.RemoveValue(test.value)

			if !test.isValid {
				if err == nil {
					t.Error("RemoveValue() should have returned an error when trying to remove an invalid element")
				} else {
					fmt.Println(err)
				}
			} else { // valid element being removed
				if err != nil {
					t.Errorf("RemoveValue() failed unexpectedly, error %v", err)
				} else {
					got := list.String()
					want := test.expList
					if got != want {
						t.Errorf("list after removal is incorrect, want: %v, got: %v", want, got)
					}

					head := list.Head()
					got = head.String()
					want = test.expHead
					if got != want {
						t.Errorf("head after removal is incorrect, want: %v, got: %v", want, got)
					}

					if head != nil {
						headPrev := head.prev
						if headPrev != nil {
							t.Errorf("head's prev pointer should always be nil, got: %v", headPrev)
						}
					}

					tail := list.Tail()
					got = tail.String()
					want = test.expTail
					if got != want {
						t.Errorf("tail after removal is incorrect, want: %v, got: %v", want, got)
					}

					if tail != nil {
						tailNext := tail.next
						if tailNext != nil {
							t.Errorf("head's prev pointer should always be nil, got: %v", tailNext)
						}
					}

					gotEmpty := list.IsEmpty()
					wantEmpty := test.isEmpty
					if gotEmpty != wantEmpty {
						t.Errorf("IsEmpty()'s value after removal is incorrect, want: %v, got: %v", wantEmpty, gotEmpty)
					}
				}
			}
		})
	}
}

func TestRemoveFirst(t *testing.T) {
	var list *SemiGenericList[prInt]

	t.Run("nil list", func(t *testing.T) {
		_, err := list.RemoveFirst()
		if err == nil {
			t.Errorf("RemoveFirst() on a nil list should return an error")
		} else {
			fmt.Println(err)
		}
	})

	list = &SemiGenericList[prInt]{}
	t.Run("empty list", func(t *testing.T) {
		_, err := list.RemoveFirst()
		if err == nil {
			t.Errorf("RemoveFirst() on an empty list should return an error")
		} else {
			fmt.Println(err)
		}
	})

	// some setup
	var val prInt
	for val = 1; val <= 5; val++ {
		err := list.AddAtBeginning(val)
		if err != nil {
			t.Fatalf("AddAtBeginning() failed, error: %v", err)
		}
	}

	var tests = []struct {
		name    string
		expRem  string
		expList string
		expHead string
		expTail string
		isEmpty bool
	}{
		{"5 element list", "5", "nil<-4<=>3<=>2<=>1->nil", "4", "1", false},
		{"4 element list", "4", "nil<-3<=>2<=>1->nil", "3", "1", false},
		{"3 element list", "3", "nil<-2<=>1->nil", "2", "1", false},
		{"2 element list", "2", "nil<-1->nil", "1", "1", false},
		{"1 element list", "1", "empty", "nil", "nil", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rem, err := list.RemoveFirst()
			if err != nil {
				t.Errorf("RemoveFirst() failed unexpectedly, error %v", err)
			} else {
				want := test.expRem
				got := rem.String()
				if got != want {
					t.Errorf("item removed is incorrect, want: %v, got: %v", want, got)
				}

				want = test.expList
				got = list.String()
				if got != want {
					t.Errorf("list after removing first element is incorrect, want: %v, got: %v", want, got)
				}

				head := list.Head()
				want = test.expHead
				got = head.String()
				if got != want {
					t.Errorf("head after removing first element is incorrect, want: %v, got: %v", want, got)
				}

				if head != nil {
					headPrev := head.prev
					if headPrev != nil {
						t.Errorf("head's prev pointer should always be nil, got: %v", headPrev)
					}
				}

				tail := list.Tail()
				want = test.expTail
				got = list.Tail().String()
				if got != want {
					t.Errorf("tail after removing first element is incorrect, want: %v, got: %v", want, got)
				}

				if tail != nil {
					tailNext := tail.next
					if tailNext != nil {
						t.Errorf("tail's next pointer should always be nil, got: %v", tailNext)
					}
				}

				wantEmpty := test.isEmpty
				gotEmpty := list.IsEmpty()

				if gotEmpty != wantEmpty {
					t.Errorf("IsEmpty()'s value after removal is incorrect, want: %v, got: %v", wantEmpty, gotEmpty)
				}
			}
		})
	}
}

func TestRemoveLast(t *testing.T) {
	var list *SemiGenericList[prInt]

	t.Run("nil list", func(t *testing.T) {
		_, err := list.RemoveLast()
		if err == nil {
			t.Errorf("RemoveLast() on a nil list should return an error")
		} else {
			fmt.Println(err)
		}
	})

	list = &SemiGenericList[prInt]{}
	t.Run("empty list", func(t *testing.T) {
		_, err := list.RemoveLast()
		if err == nil {
			t.Errorf("RemoveLast() on an empty list should return an error")
		} else {
			fmt.Println(err)
		}
	})

	// some setup
	var val prInt
	for val = 1; val <= 5; val++ {
		err := list.AddAtEnd(val)
		if err != nil {
			t.Fatalf("AddAtEnd() failed, error: %v", err)
		}
	}

	var tests = []struct {
		name    string
		expRem  string
		expList string
		expHead string
		expTail string
		isEmpty bool
	}{
		{"5 element list", "5", "nil<-1<=>2<=>3<=>4->nil", "1", "4", false},
		{"4 element list", "4", "nil<-1<=>2<=>3->nil", "1", "3", false},
		{"3 element list", "3", "nil<-1<=>2->nil", "1", "2", false},
		{"2 element list", "2", "nil<-1->nil", "1", "1", false},
		{"1 element list", "1", "empty", "nil", "nil", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rem, err := list.RemoveLast()
			if err != nil {
				t.Errorf("RemoveLast() failed unexpectedly, error %v", err)
			} else {
				want := test.expRem
				got := rem.String()
				if got != want {
					t.Errorf("item removed is incorrect, want: %v, got: %v", want, got)
				}

				want = test.expList
				got = list.String()
				if got != want {
					t.Errorf("list after removing last element is incorrect, want: %v, got: %v", want, got)
				}

				head := list.Head()
				want = test.expHead
				got = head.String()
				if got != want {
					t.Errorf("head after removing last element is incorrect, want: %v, got: %v", want, got)
				}

				if head != nil {
					headPrev := head.prev
					if headPrev != nil {
						t.Errorf("head's prev pointer should always be nil, got: %v", headPrev)
					}
				}

				tail := list.Tail()
				want = test.expTail
				got = list.Tail().String()
				if got != want {
					t.Errorf("tail after removing last element is incorrect, want: %v, got: %v", want, got)
				}

				if tail != nil {
					tailNext := tail.next
					if tailNext != nil {
						t.Errorf("tail's next pointer should always be nil, got: %v", tailNext)
					}
				}

				wantEmpty := test.isEmpty
				gotEmpty := list.IsEmpty()

				if gotEmpty != wantEmpty {
					t.Errorf("IsEmpty()'s value after removal is incorrect, want: %v, got: %v", wantEmpty, gotEmpty)
				}
			}
		})
	}
}

func TestNodeGetData(t *testing.T) {
	var n1, n2 *Node[prInt]
	n2 = &Node[prInt]{nil, 1, nil}

	prIntTests := []struct {
		name   string
		node   *Node[prInt]
		expVal prInt
		expErr error
	}{
		{"type prInt : nil node : test get data", n1, 0, nilNodeError},
		{"type prInt : non-nil node : test get data", n2, 1, nil},
	}

	for _, test := range prIntTests {
		t.Run(test.name, func(t *testing.T) {
			wantVal := test.expVal
			gotVal, err := test.node.GetData()
			if err != test.expErr {
				t.Errorf("Unexpected error for node's GetData(), want: %v, got: %v", test.expErr, err)
			}

			if gotVal != test.expVal {
				t.Errorf("Incorrect result fornode's GetData(), want: %v, got: %v", wantVal, gotVal)
			}
		})
	}

	var n3, n4 *Node[prBool]
	n4 = &Node[prBool]{nil, true, nil}

	prBoolTests := []struct {
		name   string
		node   *Node[prBool]
		expVal prBool
		expErr error
	}{
		{"type prBool : nil node : test get data", n3, false, nilNodeError},
		{"type prBool : non-nil node : test get data", n4, true, nil},
	}

	for _, test := range prBoolTests {
		t.Run(test.name, func(t *testing.T) {
			wantVal := test.expVal
			gotVal, err := test.node.GetData()
			if err != test.expErr {
				t.Errorf("Unexpected error for node's GetData(), want: %v, got: %v", test.expErr, err)
			}

			if gotVal != test.expVal {
				t.Errorf("Incorrect result fornode's GetData(), want: %v, got: %v", wantVal, gotVal)
			}
		})
	}

	var n5, n6 *Node[prString]
	n6 = &Node[prString]{nil, "a", nil}

	prStringTests := []struct {
		name   string
		node   *Node[prString]
		expVal prString
		expErr error
	}{
		{"type prString : nil node : test get data", n5, "", nilNodeError},
		{"type prString : non-nil node : test get data", n6, "a", nil},
	}

	for _, test := range prStringTests {
		t.Run(test.name, func(t *testing.T) {
			wantVal := test.expVal
			gotVal, err := test.node.GetData()
			if err != test.expErr {
				t.Errorf("Unexpected error for node's GetData(), want: %v, got: %v", test.expErr, err)
			}

			if gotVal != test.expVal {
				t.Errorf("Incorrect result fornode's GetData(), want: %v, got: %v", wantVal, gotVal)
			}
		})
	}
}
