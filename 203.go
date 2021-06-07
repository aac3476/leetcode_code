package main


func removeElements(head *ListNode, val int) *ListNode {
	headNode :=
		&ListNode{
			Val:  0,
			Next: head,
		}
	head = headNode
	for head != nil {
		prev := head
		head = head.Next
		if head == nil {
			break
		}
		if head.Val == val {
			for ;head != nil && head.Val == val;head=head.Next{}
			prev.Next=head
		}
	}
	return headNode.Next
}

func main() {

}
