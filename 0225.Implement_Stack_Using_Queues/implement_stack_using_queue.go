package _0225_Implement_Stack_Using_Queues

type MyStack struct {
	Q1 []int
	Q2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {

	// add x to queue
	// add all the other elements one by one to the back of the queue

	s.Q1 = append(s.Q1, x)
	qSize := len(s.Q1)
	for qSize > 1 {
		//add
		s.Q1 = append(s.Q1, s.Q1[0])
		//pop
		s.Q1 = s.Q1[1:]
		qSize--
	}

	/*
		Using two queues approach
		// Push to Q2
			s.Q2 = append(s.Q2, x)

			// Move all the elements from Q1 to Q2 one by one
			for len(s.Q1) > 0 {
				s.Q2 = append(s.Q2, s.Q1[0])
				s.Q1 = s.Q1[1:]
			}

			// Swap Q1 and Q2
			s.Q1, s.Q2 = s.Q2, s.Q1
	*/
}

func (s *MyStack) Pop() int {
	x := s.Q1[0]
	s.Q1 = s.Q1[1:]
	return x
}

func (s *MyStack) Top() int {
	return s.Q1[0]
}

func (s *MyStack) Empty() bool {
	return len(s.Q1) == 0
}
