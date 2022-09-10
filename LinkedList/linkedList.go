package LinkedList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (listNode *ListNode) Insert(data int) {
	newNode := &ListNode{
		Val:  data,
		Next: nil,
	}
	temp := listNode
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newNode
}

func (listNode *ListNode) PrintLinkedList() {
	temp := listNode
	for temp.Next != nil {
		fmt.Printf("%v->", temp.Val)
		temp = temp.Next
	}
	fmt.Printf("%v\n", temp.Val)
}
