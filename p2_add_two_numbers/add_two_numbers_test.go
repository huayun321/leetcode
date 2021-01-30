package p2_add_two_numbers

import (
	"testing"
)

func TestListNode_InitFromList(t *testing.T) {
	type args struct {
		l []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "[0]", args: args{l: []int{0}}},
		{name: "[2,4,3]", args: args{l: []int{2, 4, 3}}},
		{name: "[5,6,4]", args: args{l: []int{5, 6, 4}}},
		{name: "[9,9,9,9,9,9]", args: args{l: []int{9, 9, 9, 9, 9, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ln := &ListNode{}
			ln.InitFromList(tt.args.l)
			var i int

			for ln != nil {
				if ln.Val != tt.args.l[i] {
					t.Errorf("listnode init from int arr not match")
				}
				ln = ln.Next
				i++
			}
		})
	}
}

func Test_addTwoNumbers(t *testing.T) {
	ln243 := &ListNode{}
	ln243.InitFromList([]int{2, 4, 3})
	ln564 := &ListNode{}
	ln564.InitFromList([]int{5, 6, 4})
	ln708 := &ListNode{}
	ln708.InitFromList([]int{7, 0, 8})

	ln0 := &ListNode{}
	ln0.InitFromList([]int{0})

	ln79 := &ListNode{}
	ln79.InitFromList([]int{9, 9, 9, 9, 9, 9, 9})
	ln49 := &ListNode{}
	ln49.InitFromList([]int{9, 9, 9, 9})
	ln899 := &ListNode{}
	ln899.InitFromList([]int{8, 9, 9, 9, 0, 0, 0, 1})

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[2,4,3] [5,6,4]",
			args: args{
				l1: ln243,
				l2: ln564,
			},
			want: ln708,
		},
		{
			name: "[0] [0]",
			args: args{
				l1: ln0,
				l2: ln0,
			},
			want: ln0,
		},
		{
			name: "[9,9,9,9,9,9] [9,9,9,9]",
			args: args{
				l1: ln79,
				l2: ln49,
			},
			want: ln899,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			want := tt.want

			for want != nil {
				t.Logf("==== got %v \n", got.Val)
				if want.Val != got.Val {
					t.Errorf("want got not match")
				}
				want = want.Next
				got = got.Next
			}

			for got != nil {
				t.Logf("==== got %v \n", got.Val)
				if want.Val != got.Val {
					t.Errorf("want got not match")
				}
				want = want.Next
				got = got.Next
			}
		})
	}
}
