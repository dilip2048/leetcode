package _203_Remove_Linked_List_Elements

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

func removeElements(head *ListNode, val int) *ListNode {
	left := &ListNode{}
	right := head
	result := left
	for right != nil {
		if right.Val != val {
			left.Next = &ListNode{
				Val:  right.Val,
				Next: nil,
			}
			left = left.Next
		}
		right = right.Next
	}
	return result.Next
}
