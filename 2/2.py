# -*- coding: utf-8 -*-  
"""
Create on 04-26 11:30 2020
@Author ywx 
@File 2.py
"""
import math

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        num1 = 0
        num2 = 0
        i = 0
        while True:
            num1 = num1 + l1.val * pow(10,i)
            if l1.next == None:
                break
            l1 = l1.next
            i = i + 1
        i = 0
        while True:
            num2 = num2 + l2.val * pow(10, i)
            if l2.next == None:
                break
            l2 = l2.next
            i = i + 1
        num1 = num1 + num2
        l3 = ListNode(0)
        l4 = l3
        while True:
            l4.next = None
            l4.val = num1%10
            if num1/10 < 1:
                return l3
            num1 = math.floor(num1//10)
            l5 = ListNode(0)
            l4.next = l5
            l4 = l5


def makeL(a):
    l = ListNode(0)
    s = l
    for i in range(len(a)):
        if i == len(a) -1 :
            l.next = None
            l.val = a[i]
            return s
        l.val = a[i]
        l.next = ListNode(0)
        l = l.next

s = Solution()
a = s.addTwoNumbers(l1 = makeL([1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]),l2 = makeL([5,6,4]))
pass