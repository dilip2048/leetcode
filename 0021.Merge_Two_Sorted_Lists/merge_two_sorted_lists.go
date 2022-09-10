package _019_Remove_Nth_Node_From_End_of_List

import _ "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newList := &ListNode{}
	out := newList
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			newList.Next = list1
			list1 = list1.Next
		} else {
			newList.Next = list2
			list2 = list2.Next
		}
		newList = newList.Next
	}
	if list1 != nil {
		newList.Next = list1
	} else if list2 != nil {
		newList.Next = list2
	}
	return out.Next
}
