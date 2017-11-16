package main

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	type Case struct {
		in  string
		out []string
	}
	cases := []Case{
		{"dataA", []string{"dataA"}},
		{"dataB", []string{"dataA", "dataB"}},
		{"dataC", []string{"dataA", "dataB", "dataC"}},
	}
	s := Stack{}
	for i, c := range cases {
		s.Push(c.in)
		if !reflect.DeepEqual(s.a, c.out) {
			t.Errorf("配列の%d番目がエラー: s.Push(%#v)は結果%#vが返ります。実際は%#vが返ってきています。\n", i, c.in, c.out, s.a)
		}
	}
}

func TestPop(t *testing.T) {
	type Case struct {
		out string
	}
	cases := []Case{
		{"dataC"},
		{"dataB"},
	}
	s := Stack{a: []string{"dataA", "dataB", "dataC"}}
	for i, c := range cases {
		if got := s.Pop(); got != c.out {
			t.Errorf("配列の%d番目がエラー: s.Pop()は結果%#vが返ります。実際は%#vが返ってきています。\n", i, c.out, got)
		}
	}
}
