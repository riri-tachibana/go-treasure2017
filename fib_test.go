package main

import "testing"

func TestFib(t *testing.T) {
	type Case struct {
		n, out int
	}
	cases := []Case{
		{0, 0},
		{1, 1},
		{10, 55},
	}
	for i, c := range cases {
		if got := fib(c.n); got != c.out {
			t.Errorf("配列の%d番目がエラー: fib(%d)は結果%dが返ります。実際は%dが返ってきています。\n", i, c.n, c.out, got)
		}
	}
}
