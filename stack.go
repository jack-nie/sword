package algo

//Stack struct stack
type Stack struct {
	Data []int
}

// Pop pop an element from stack
func (s *Stack) Pop() int {
	if s.Empty() {
		panic("stack is empty!")
	}
	result := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return result
}

//Push push an element to stack
func (s *Stack) Push(value int) {
	s.Data = append(s.Data, value)
}

//Empty check if the stack is empty
func (s *Stack) Empty() bool {
	return len(s.Data) == 0
}

func (s *Stack) Size() int {
	return len(s.Data)
}
