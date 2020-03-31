package algorithm

import (
	"log"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//func reverseList(head *ListNode) *ListNode {
//
//}

func TestReverseLinkedList(t *testing.T) {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	head := &ListNode{1, node2}

	cur := new(ListNode)
	prev := new(ListNode)

	for {
		// first
		if cur.Val == 0 {
			cur = head
			prev = new(ListNode)
			head.Next = prev
			head = head.Next
		} else {
			prev = cur
			cur = cur.Next
		}

		if cur == nil {
			log.Println("end")
			break
		}

		log.Printf("prev: %v, cur: %v\n", prev.Val, cur.Val)
	}
}
