package main

import "fmt"

func main() {
	message := make(chan string, 4)

	message <- "Hello"
	message <- "World"
	message <- "I'm software developer"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
