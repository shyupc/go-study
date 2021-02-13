package stack

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (interface{}, bool) {
	if !s.IsEmpty() {
		v := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return v, true
	}
	return 0, false
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
