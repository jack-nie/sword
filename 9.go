package algo

//StackQueue implemented by two stack
type StackQueue struct {
	stack1 *Stack
	stack2 *Stack
}

//AppendTail append value to tail
func (s *StackQueue) AppendTail(value int) {
	s.stack1.Push(value)
}

//DeleteHead delete head from queue
func (s *StackQueue) DeleteHead() int {
	if s.stack2.Size() <= 0 {
		for s.stack1.Size() > 0 {
			data := s.stack1.Pop()
			s.stack2.Push(data)
		}
	}
	return s.stack2.Pop()
}
