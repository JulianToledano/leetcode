package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	// l2 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}

	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	// l1 := ListNode{5, nil}
	// l2 := ListNode{5, nil}

	r := addTwoNumbers(&l1, &l2)
	fmt.Println(r)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := ListNode{l1.Val + l2.Val, nil}
	if l1.Next == nil && l2.Next == nil {
		if r.Val/10 > 0 {
			r = *recurIter(&r)
		} else {
			return &r
		}
	} else {
		var s int
		if r.Val/10 > 0 {
			s = r.Val / 10
			r.Val %= 10
		}
		if l1.Next == nil {
			l2.Next.Val += s
			r.Next = recurIter(l2.Next)
		} else if l2.Next == nil {
			l1.Next.Val += s
			r.Next = recurIter(l1.Next)
		} else {
			l2.Next.Val += s
			r.Next = addTwoNumbers(l1.Next, l2.Next)
		}
	}
	return &r
}

func recurIter(l *ListNode) *ListNode {
	if l.Val/10 > 0 {
		s := l.Val / 10
		l.Val %= 10
		if l.Next == nil {
			l.Next = &ListNode{s, nil}
		} else {
			l.Next.Val += s
			recurIter(l.Next)
		}
	}
	return l
}
