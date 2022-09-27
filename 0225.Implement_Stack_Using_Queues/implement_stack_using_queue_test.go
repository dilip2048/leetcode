package _0225_Implement_Stack_Using_Queues

import "testing"

func TestStack1(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	popped := obj.Pop()
	if popped != 4 {
		t.Fail()
	}
	top := obj.Top()
	if top != 3 {
		t.Fail()
	}
	isEmpty := obj.Empty()
	if isEmpty != false {
		t.Fail()
	}
}
