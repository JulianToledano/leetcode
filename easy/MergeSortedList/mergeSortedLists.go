package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{2, nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{4, nil}

	l2 := ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}

	mergeTwoLists(&l1, &l2)
	fmt.Println(l1)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	if l2.Val <= l1.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
	return nil
}
