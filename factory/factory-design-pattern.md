# Factory Design Pattern

*The Factory Design Pattern is a creational pattern that provides an interface for creating 
objects in a superclass but allows subclasses to alter the type of objects that will be created. 
The Factory Pattern solves the problem of creating objects without specifying the exact class of 
object that will be created.*


### Key Concepts:
- Creator (Factory): Defines the method for creating an object, but it’s up to the subclasses to decide which class to instantiate.
- Product: The interface or abstract class that the concrete products implement.
- ConcreteProduct: A specific class that implements the Product interface.
- Factory Method: A method that returns an instance of a product.

### Benefits of Factory Pattern:
- Loose Coupling: The client code doesn't need to know which class to instantiate. It only depends on the interface.
- Single Responsibility: The creation logic of objects is separated from the logic that uses the objects.
- Extensibility: New product types can be introduced without altering the client code.

```go
package main

import "fmt"

// Product interface that all concrete products will implement
type Shape interface {
    Draw()
}

// ConcreteProduct - Circle
type Circle struct {}

func (c *Circle) Draw() {
    fmt.Println("Drawing Circle")
}

// ConcreteProduct - Square
type Square struct {}

func (s *Square) Draw() {
    fmt.Println("Drawing Square")
}

// ConcreteProduct - Rectangle
type Rectangle struct {}

func (r *Rectangle) Draw() {
    fmt.Println("Drawing Rectangle")
}

// Creator interface defines the factory method
type ShapeFactory interface {
    CreateShape() Shape
}

// ConcreteCreator - CircleFactory
type CircleFactory struct {}

func (cf *CircleFactory) CreateShape() Shape {
    return &Circle{}
}

// ConcreteCreator - SquareFactory
type SquareFactory struct {}

func (sf *SquareFactory) CreateShape() Shape {
    return &Square{}
}

// ConcreteCreator - RectangleFactory
type RectangleFactory struct {}

func (rf *RectangleFactory) CreateShape() Shape {
    return &Rectangle{}
}

func main() {
    // Using the factory pattern to create objects
    factories := []ShapeFactory{
        &CircleFactory{},
        &SquareFactory{},
        &RectangleFactory{},
    }

    // Loop through and create shapes using respective factories
    for _, factory := range factories {
        shape := factory.CreateShape()
        shape.Draw() // Calling Draw method of the respective shape
    }
}
```


## Explanation:
- Shape Interface: Defines the Draw() method that all concrete shapes must implement.
- Concrete Products: Circle, Square, and Rectangle implement the Shape interface and provide the specific implementation for Draw().
- ShapeFactory Interface: Defines the CreateShape() method that factories must implement.
- Concrete Factories: CircleFactory, SquareFactory, and RectangleFactory are concrete implementations of the ShapeFactory interface, responsible for creating specific products.



#### The Factory Design Pattern provides loose coupling, centralized object creation, extensibility, improved testability, flexibility, single responsibility, and simplified code maintenance by abstracting and managing object creation.