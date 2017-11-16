package main

import "fmt"

func fib(n int) int {
	if n < 30 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(10))
}
