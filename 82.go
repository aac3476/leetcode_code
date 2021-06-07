package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Val: 0, Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}


func main() {
	var node1 = new(ListNode)
	var node2 = new(ListNode)
	var node3 = new(ListNode)
	var node4 = new(ListNode)
	var node5 = new(ListNode)

	node1.Val = 1
	node2.Val = 1
	node3.Val = 3
	node4.Val = 3
	node5.Val = 5

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	// node4.Next = node5

	res := deleteDuplicates(node1)

	fmt.Print(res)
}
