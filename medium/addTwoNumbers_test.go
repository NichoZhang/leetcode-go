package main

import "testing"

func TestFunc1(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{},
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{},
			},
		},
	}

	t.Log(addTwoNumbers(l1, l2))
}
