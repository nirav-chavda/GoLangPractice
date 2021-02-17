package main

import (
	"fmt"
	"math"
)

/* define an interface */
type Shape interface {
	area() float32
}

/* define a circle */
type Circle struct {
	x,y,radius float32
}

/* define a rectangle */
type Rectangle struct {
	width, height float32
}

/* define a method for circle (implementation of Shape.area())*/
func (circle Circle) area() float32  {
	return circle.radius * circle.radius * math.Pi
}

/* define a method for rectangle (implementation of Shape.area())*/
func (rectangle Rectangle) area() float32 {
	return rectangle.height * rectangle.width
}

/* define a method for shape */
func getArea(shape Shape) float32 {
	return shape.area()
}

func main() {
	var circle = Circle{10,20,5}
	var rectangle = Rectangle{10,20}

	fmt.Println(getArea(circle), " ", getArea(rectangle))
}
