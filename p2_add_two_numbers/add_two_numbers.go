package p2_add_two_numbers

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) InitFromList(l []int) {
	next := ln
	for i, v := range l {
		next.Val = v
		if i == len(l)-1 {
			break
		}
		next.Next = &ListNode{}
		next = next.Next
	}
}

func recurseAdd(l1 *ListNode, l2 *ListNode, ln *ListNode) {
	if l1 == nil && l2 == nil {
		return
	}

	var first int
	var second int
	var l1n *ListNode
	var l2n *ListNode

	if l1 != nil {
		first = l1.Val
		l1n = l1.Next
	}

	if l2 != nil {
		second = l2.Val
		l2n = l2.Next
	}

	sum := first + second + ln.Val
	ln.Val = sum % 10

	ln.Next = &ListNode{}

	if sum >= 10 {
		ln.Next.Val = 1
	}

	if l1n == nil && l2n == nil && ln.Next.Val == 0 {
		ln.Next = nil
		return
	}

	recurseAdd(l1n, l2n, ln.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := &ListNode{}
	recurseAdd(l1, l2, ln)
	return ln
}
