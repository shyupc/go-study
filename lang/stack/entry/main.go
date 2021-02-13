package main

import (
	"fmt"

	"github.com/shyupc/go-study/lang/stack"
)

func main() {
	s := stack.Stack{1}

	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())

	s.Push("abc")
	fmt.Println(s.Pop())
}
