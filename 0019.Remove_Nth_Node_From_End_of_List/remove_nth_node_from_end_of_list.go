package _019_Remove_Nth_Node_From_End_of_List

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head
	var count int
	for temp != nil {
		temp = temp.Next
		count++
	}
	fmt.Println(count)
	temp = head
	for i := 0; i < count-n; i++ {
		temp = temp.Next
	}
	//delete node now
	temp.Next = temp.Next.Next
	return head
}
