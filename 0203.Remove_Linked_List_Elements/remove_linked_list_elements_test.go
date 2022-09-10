package _203_Remove_Linked_List_Elements

import "testing"

func TestMergeList1(t *testing.T) {
	linkedList1 := ListNode{
		Val:  1,
		Next: nil,
	}
	linkedList1.Insert(1)
	linkedList1.Insert(2)
	linkedList1.Insert(4)

	expectedLinkedList := &ListNode{
		Val:  1,
		Next: nil,
	}
	expectedLinkedList.Insert(1)
	expectedLinkedList.Insert(4)

	ll := removeElements(&linkedList1, 2)
	for expectedLinkedList.Next != nil {
		if ll.Val != expectedLinkedList.Val {
			t.Fail()
		}
		ll = ll.Next
		expectedLinkedList = expectedLinkedList.Next
	}
	if ll.Next != nil {
		t.Fail()
	}
}
