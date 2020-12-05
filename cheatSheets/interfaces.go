/////////////////////////////////
// Implementing Interfaces in Go
// Go Playground: https://play.golang.org/p/SMjFrOYL5f3
/////////////////////////////////

package main

import (
	"fmt"
	"math"
)

// declaring an interface type called shape
// an interface contains only the signatures of the methods, but not their implementation
type shape interface {
	area() float64
	perimeter() float64
}

// declaring 2 struct types that represent geometrical shapes: rectangle and circle

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

// method that calculates circle's area
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// method that calculates rectangle's area
func (r rectangle) area() float64 {
	return r.height * r.width
}

// method that calculates circle's perimeter
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// method that calculates rectangle's perimeter
func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// any type that implements the interface is also of type of the interface
// rectangle and circle values are also of type shape
func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

// declaring an empty interface
type emptyInterface interface {
}

// declaring a new struct type which has one field of type empty interface
type person struct {
	info interface{}
}

func main() {

	// declaring a circle and a rectangle values
	c1 := circle{radius: 5}
	r1 := rectangle{width: 3, height: 2}

	// circle and rectangle both implements the geometry interface  because they implement all methods of the interface
	// an interface is implicitly implemented in Go
	print(c1)
	print(r1)

	a := c1.area()
	fmt.Println("Circle Area:", a)

	// declaring an interface value that holds a circle type value
	var s shape = circle{radius: 2.5}

	fmt.Printf("%T\n", s) //interface dynamic type is circle

	// no direct access to interface's dynamic values
	// s.volume() -> error

	// there is access only to the methods that are defined inside the interface
	fmt.Printf("Circle Area:%v\n", s.area())

	// an interface value hides its dynamic value.
	// use type assertion to extract and return the dynamic value of the interface value.
	fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())

	// checking if the assertion succeded or not
	ball, ok := s.(circle)
	if ok == true {
		fmt.Printf("Ball Volume:%v\n", ball.volume())
	}

	//** TYPE SWITCHES **//

	// it permits several type assertions in series
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)

	}

	// declaring an empty interface value
	var empty interface{}

	// an empty interface may hold values of any type
	// storing an int in the empty interface
	empty = 5
	fmt.Println(empty) // => 5

	// storing a string in the empty interface
	empty = "Go"
	fmt.Println(empty) // => Go

	// storing a slice in the empty interface
	empty = []int{2, 34, 4}
	fmt.Println(empty) // => [2 34 4]

	// fmt.Println(len(empty)) -> error, and it doesn't work!

	// retrieving the dynamic value using an assertion
	fmt.Println(len(empty.([]int))) // => 3

	// declaring person value
	you := person{}

	// assigning any value to empty interface field
	you.info = "You name"
	you.info = 40
	you.info = []float64{4.5, 6., 8.1}

	fmt.Println(you.info)
}
