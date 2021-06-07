package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


var nodeHead *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	l := 0
	nodeHead = head
	for head != nil {
		head = head.Next
		l++
	}
	return buildTree(0,l-1)
}


func buildTree (start int, end int) *TreeNode{
	if start > end{
		return nil
	}
	mid := (start + end) /2
	node := &TreeNode{
		Val:   nodeHead.Val,
		Left:  buildTree(start,mid - 1),
	}
	nodeHead = nodeHead.Next
	node.Right = buildTree(mid + 1, end)
	return node
}




func makeLinkList(list []int) *ListNode{
	start := &ListNode{Val: list[0],Next: nil}
	last := start
	for i:=1;i<len(list);i++{
		last.Next = &ListNode{
			Val: list[i],
			Next: nil,
		}
		last = last.Next
	}
	return start
}

func main() {
	sortedListToBST(makeLinkList([]int{-10, -3, 0, 5, 9}))
}
