package stack 

type node struct {
	value interface{}
	next *node
}

type Stack struct {
	count int
	head *node
}

func (s Stack) Len() int {
	return s.count
}

func (s *Stack) Push(o interface{}) {
	s.head = &node{o, s.head}
	s.count++
}

func (s Stack) Peek() (val interface{}, ok bool) {
	if s.head != nil {
		val, ok = s.head.value, true
	}	
	return
}
	
func (s *Stack) Pop() (val interface{}, ok bool) {
	if s.head != nil {
		val, s.head, ok = s.head.value, s.head.next, true
		s.count--
	}
	return
}

func (s Stack) IsEmpty() bool {
	return s.count == 0
}
	