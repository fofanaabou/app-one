package main

import "fmt"

func main() {
	fmt.Println("------ for loop----")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println("------Do for N time----")
	for i := range 3 {
		fmt.Println(i)
	}

	fmt.Println("------without condition-----------")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("------continue next iteration---------")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
