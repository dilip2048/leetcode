package LinkedList

import "testing"

func TestLinkedList(t *testing.T) {
	linkedList := ListNode{
		Val:  0,
		Next: nil,
	}
	linkedList.Insert(1)
	linkedList.Insert(2)
	linkedList.Insert(3)
	linkedList.PrintLinkedList()

}
