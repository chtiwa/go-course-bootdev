package main

import "fmt"

type car struct {
	color string
}

// a method of car
func (c *car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
}
