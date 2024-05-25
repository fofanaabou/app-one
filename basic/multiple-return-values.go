package main

import "fmt"

func getValues() (int, int) {
	return 5, 7
}

func main() {
	a, b := getValues()
	fmt.Println(a)
	fmt.Println(b)

	_, c := getValues()
	fmt.Println(c)
}
