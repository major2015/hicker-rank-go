package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
    Val int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

/**
* recursion 递归思维解决子问题，往上解决反转
*/
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	newHead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

var head = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: nil,
				},
			},
		},
	},
}

func TestReverse(t *testing.T) {
	newhead := reverseList(head)
	assert.Equal(t, newhead.Val, 5)
}

func TestReverse1(t *testing.T) {
	newhead := reverseList1(head)
	assert.Equal(t, newhead.Val, 5)
}