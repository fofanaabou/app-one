package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println("Hello Abou")
	fmt.Println(msg)
}
