package main


func getIntersectionNode(headA, headB *ListNode) *ListNode {
     lenA,lenB := 1,1
     A,B := headA,headB
     for A.Next != nil {
          lenA ++
          A = A.Next
     }
     for B.Next != nil {
          lenB ++
          B = B.Next
     }
     var fast,slow,result *ListNode
     var walk int
     if lenA > lenB {
          walk = lenA - lenB
          fast = headA
          slow = headB
     } else {
          walk = lenB - lenA
          fast = headB
          slow = headA
     }
     for ;walk>0;walk-- {
          fast = fast.Next
     }
     if fast == slow {
          return fast
     }
     for true {
          fast = fast.Next
          slow = slow.Next
          if fast == slow {
               return fast
          }
          if fast.Next == nil {
               break
          }
     }
     return result
}
