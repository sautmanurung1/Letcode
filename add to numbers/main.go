package main

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var curr *ListNode
	var carry int
	for l1 != nil || l2 != nil {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		carry = sum / 10
		sum = sum % 10
		if head == nil {
			head = &ListNode{Val: sum}
			curr = head
		} else {
			curr.Next = &ListNode{Val: sum}
			curr = curr.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return head
}
