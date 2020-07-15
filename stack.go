package gstl

import "errors"

type stack struct {
	valList []interface{}
	count   int
	size    int
}

func NewStack(size int) *stack {
	return &stack{[]interface{}{}, 0, size}
}

func (s *stack) GetTop() interface{} {
	if s.count == 0 {
		return nil
	}
	return s.valList[s.count-1]
}

func (s *stack) Pop() interface{} {
	if s.count <= 0 {
		return nil
	}
	v := s.valList[s.count-1]
	s.valList = s.valList[:s.count-1]
	s.count -= 1
	return v
}

func (s *stack) Push(v interface{}) error {
	if s.size <= s.count {
		return errors.New("stack is full ")
	}
	s.valList = append(s.valList, v)
	s.count += 1
	return nil
}

func (s *stack) GetCount() int {
	return s.count
}

func (s *stack) IsEmpty() bool {
	return s.count == 0
}

func (s *stack) IsFull() bool {
	return s.count == s.size
}
