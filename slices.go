package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 4)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "A"
	s[1] = "B"
	s[2] = "C"
	s[3] = "D"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	s = append(s, "E")
	s = append(s, "F", "G")
	fmt.Println("after appended:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"H", "I", "J"}
	fmt.Println("dcl:", t)

	t2 := []string{"H", "I", "J"}
	if slices.Equal(t, t2) {
		fmt.Println("t==t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
