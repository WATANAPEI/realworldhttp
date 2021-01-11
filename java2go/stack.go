package collection_test

//zero value of Stack is empty slice
type Stack struct {
	data []string
}

func (s *Stack) Push(x string) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() string {
	n := len(s.data) - 1
	res := s.data[n]
	s.data[n] = "" // to avoid memory leak
	s.data = s.data[:n]
	return res
}

func (s *Stack) Size() int {
	return len(s.data)
}
