package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	c := &ListNode{4, nil}
	b := &ListNode{2, c}
	a := &ListNode{1, b}
	list1 := a

	c2 := &ListNode{4, nil}
	b2 := &ListNode{3, c2}
	a2 := &ListNode{1, b2}
	list2 := a2

	g := *mergeTwoLists(list1, list2)
	fmt.Println(g.Val)
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
