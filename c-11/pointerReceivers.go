package main

import "fmt"

type circle struct {
	width  int
	height int
	radius int
}

// a method of cirlce
func (c *circle) doubleRadius() {
	c.radius *= 2
}

// another method of circle
func (c *circle) setWidthHeight(w, h int) {
	c.width = w
	c.height = h
}

func main() {
	c := circle{
		width:  8,
		height: 8,
		radius: 2,
	}
	c.doubleRadius()
	c.setWidthHeight(10, 10)
	fmt.Println(c)
}
