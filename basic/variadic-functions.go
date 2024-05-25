package main

import "fmt"

func sum(nums ...int) {

	fmt.Println(nums, "")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(2, 4, 5)
	sum(6, 8, 2, 5)
	sum(7, 9, 2, 4, 5, 6, 3, 9)

	nums := []int{17, 29, 42, 44, 56, 36, 63, 69}
	sum(nums...)
}
