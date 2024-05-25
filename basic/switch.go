package main

import (
	"fmt"
	"time"
)

func main() {
	pickerOne(2)
	pickerOne(3)
	pickerTwo(time.Tuesday)

	whatAmi := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm an bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type %T\n", t)

		}
	}

	whatAmi(true)
	whatAmi(1)
	whatAmi("hey")
}

func pickerOne(num int) {
	fmt.Println("Write ", num, "as ")
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}

func pickerTwo(t time.Weekday) {
	switch t {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekday")
	}
}
