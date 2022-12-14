package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	c := &ListNode{3, nil}
	b := &ListNode{4, c}
	a := &ListNode{2, b}
	list1 := a

	c2 := &ListNode{4, nil}
	b2 := &ListNode{6, c2}
	a2 := &ListNode{5, b2}
	list2 := a2
	fmt.Println(*addTwoNumbers(list1, list2).Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	m := 0
	head := new(ListNode)
	tail := head
	for l1 != nil || l2 != nil {
		p := new(ListNode)
		if l1 != nil {
			p.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			p.Val += l2.Val
			l2 = l2.Next
		}

		p.Val += m
		m = p.Val / 10
		p.Val = p.Val % 10
		tail.Next = p
		tail = p
	}
	if m > 0 {
		tail.Next = &ListNode{m, nil}
	}
	return head.Next
}
