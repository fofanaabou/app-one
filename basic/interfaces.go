package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type Rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type triangle struct {
	a, b, c float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perim() float64 {
	return 2 * (r.width + r.height)

}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (t triangle) area() float64 {
	return t.a * t.b * t.c
}

func (t triangle) perim() float64 {
	return t.a + t.b + t.c
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := Rectangle{width: 2, height: 5}
	c := circle{radius: 5}
	t := triangle{a: 3, b: 4, c: 5}

	measure(r)
	measure(c)
	measure(t)
}
