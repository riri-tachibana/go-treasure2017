package main

import "fmt"

type Stack struct {
	a []string
}

func (s *Stack) Push(target string) {
	s.a = append(s.a, target)
}

func (s *Stack) Pop() string {
	if len(s.a) > 0 {
		result := s.a[len(s.a)-1]
		s.a = s.a[0 : len(s.a)-1]
		return result
	}
	return ""
}

func main() {
	s := Stack{}
	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")
	fmt.Println(s.Pop()) // "dataC"
	fmt.Println(s.Pop()) // "dataB"
	s.Push("dataD")
	fmt.Println(s.Pop()) // "dataD"
	fmt.Println(s.Pop()) // "dataA"
	fmt.Println(s.Pop()) // ""
}
