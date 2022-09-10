package _019_Remove_Nth_Node_From_End_of_List

import (
	"testing"
)

func TestMergeList1(t *testing.T) {
	linkedList1 := ListNode{
		Val:  1,
		Next: nil,
	}
	linkedList1.Insert(2)
	linkedList1.Insert(4)

	linkedList2 := ListNode{
		Val:  1,
		Next: nil,
	}
	linkedList2.Insert(3)
	linkedList2.Insert(4)

	expectedLinkedList := &ListNode{
		Val:  1,
		Next: nil,
	}
	expectedLinkedList.Insert(1)
	expectedLinkedList.Insert(2)
	expectedLinkedList.Insert(3)
	expectedLinkedList.Insert(4)
	expectedLinkedList.Insert(4)

	ll := mergeTwoLists(&linkedList1, &linkedList2)
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
