package _002_Add_Two_Numbers

// ListNode is the structure of node of the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// this method will add two linked list and form a new linked list
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var head *ListNode
	temp := head
	for l1 != nil && l2 != nil {
		remainder := (l1.Val + l2.Val + carry) % 10
		if (l1.Val + l2.Val + carry) > 9 {
			carry = 1
		} else {
			carry = 0
		}
		if head == nil {
			head = &ListNode{
				Val:  remainder,
				Next: nil,
			}
		} else {
			temp = head
			for temp.Next != nil {
				temp = temp.Next
			}
			temp.Next = &ListNode{
				Val:  remainder,
				Next: nil,
			}
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	// check if l1 is not nil. if not nil, traverse and put all the numbers to the final linked list
	for l1 != nil {
		remainder := (l1.Val + carry) % 10
		if (l1.Val + carry) > 9 {
			carry = 1
		} else {
			carry = 0
		}
		if head == nil {
			head = &ListNode{
				Val:  remainder,
				Next: nil,
			}
		} else {
			temp = head
			for temp.Next != nil {
				temp = temp.Next
			}
			temp.Next = &ListNode{
				Val:  remainder,
				Next: nil,
			}
		}
		l1 = l1.Next
	}
	// check if l1 is not nil. if not nil, traverse and put all the numbers to the final linked list
	for l2 != nil {
		remainder := (l2.Val + carry) % 10
		if (l2.Val + carry) > 9 {
			carry = 1
		} else {
			carry = 0
		}
		if head == nil {
			head = &ListNode{
				Val:  remainder,
				Next: nil,
			}
		} else {
			temp = head
			for temp.Next != nil {
				temp = temp.Next
			}
			temp.Next = &ListNode{
				Val:  remainder,
				Next: nil,
			}
		}
		l2 = l2.Next
	}

	if carry == 1 {
		temp = head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return head
}
