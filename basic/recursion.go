package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(6))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return 2
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
