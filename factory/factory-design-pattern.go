package main

import "fmt"

// Shape interface
type Shape interface {
	Draw()
}

// Circle struct
type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing Circle")
}

// Square struct
type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Drawing Square")
}

// Rectangle struct
type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("Drawing Rectangle")
}

func main() {
	// Create a list (slice) of Shape interfaces
	shapes := []Shape{
		&Circle{},    // Add a Circle to the list
		&Square{},    // Add a Square to the list
		&Rectangle{}, // Add a Rectangle to the list
	}

	// Iterate through the slice and call the Draw method on each item
	for _, shape := range shapes {
		shape.Draw()
	}
}
