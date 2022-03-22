package _002_add_two_numbers

import (
	"testing"
)

func convertArrayToLinkedList(a []int) *ListNode {
	var head *ListNode
	for i := 0; i < len(a); i++ {
		if head == nil {
			// create a node. this is the one way to create a new node. another
			// way to create new node is using new keyword i.e new(ListNode)
			head = &ListNode{
				Val:  a[i],
				Next: nil,
			}
		} else {
			temp := head
			// traverse till last node and then add a new node.
			for temp.Next != nil {
				temp = temp.Next
			}
			temp.Next = &ListNode{
				Val:  a[i],
				Next: nil,
			}
		}
	}
	return head
}

func convertListToArray(list *ListNode) []int {
	var slice []int
	for list != nil {
		slice = append(slice, list.Val)
		list = list.Next
	}
	return slice[0 : cap(slice)-1]
}

func TestAddTwoNumbers(t *testing.T) {
	a := []int{2, 4, 3}
	b := []int{5, 6, 4}
	final := []int{7, 0, 8}
	l1 := convertArrayToLinkedList(a)
	l2 := convertArrayToLinkedList(b)
	l := addTwoNumbers(l1, l2)
	sum := convertListToArray(l)
	for i := 0; i < len(sum); i++ {
		if sum[i] != final[i] {
			t.Fail()
		}
	}
}
