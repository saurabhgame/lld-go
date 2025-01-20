# Builder Design Pattern


*The Builder Design Pattern is a creational design pattern that provides a 
flexible solution to constructing complex objects. It separates the construction
of an object from its representation, allowing the same construction process to
create different representations.*


### When to Use the Builder Pattern

1. When the object to be constructed has multiple optional fields or parameters.
2. When the object construction process is complex and needs to be broken into discrete steps.
3. To avoid a "telescoping constructor" problem, where a class has multiple constructors with varying numbers of parameters.


### Key Components

1. roduct: The complex object that is being built.
- Example: Car, Pizza, Computer.

2. Builder: Abstract interface or class defining the building steps.
- Example: Methods like SetEngine(), SetSeats(), SetColor().

3. Concrete Builder: Implements the steps defined in the builder interface.
- Example: CarBuilder.

4. Director (Optional): Orchestrates the building process using the builder.
- Example: A class that defines a process to construct specific configurations of the object (e.g., "BuildSportsCar").

5. Client: The code that uses the builder to construct the object.



### Advantages

- Readable and Flexible Code: Easy to construct objects with clear method chaining.

- Step-by-Step Construction: Can construct an object incrementally.

- Reusability: Same building logic can create different representations of the object.

- Encapsulation: Hides the construction details from the client.

``` go 
package main

import "fmt"

// Car represents the product being built
type Car struct {
	Brand  string
	Model  string
	Color  string
	Engine string
	Seats  int
}

// CarBuilder is a builder for creating Car objects
type CarBuilder struct {
	brand  string
	model  string
	color  string
	engine string
	seats  int
}

// NewCarBuilder initializes the CarBuilder
func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

// SetBrand sets the car brand
func (b *CarBuilder) SetBrand(brand string) *CarBuilder {
	b.brand = brand
	return b
}

// SetModel sets the car model
func (b *CarBuilder) SetModel(model string) *CarBuilder {
	b.model = model
	return b
}

// SetColor sets the car color
func (b *CarBuilder) SetColor(color string) *CarBuilder {
	b.color = color
	return b
}

// SetEngine sets the car engine
func (b *CarBuilder) SetEngine(engine string) *CarBuilder {
	b.engine = engine
	return b
}

// SetSeats sets the number of seats
func (b *CarBuilder) SetSeats(seats int) *CarBuilder {
	b.seats = seats
	return b
}

// Build creates a Car object from the builder
func (b *CarBuilder) Build() Car {
	return Car{
		Brand:  b.brand,
		Model:  b.model,
		Color:  b.color,
		Engine: b.engine,
		Seats:  b.seats,
	}
}

func main() {
	// Create a Sports Car
	sportsCar := NewCarBuilder().
		SetBrand("Ferrari").
		SetModel("488 Spider").
		SetColor("Red").
		SetEngine("V8").
		SetSeats(2).
		Build()

	// Create an SUV
	suv := NewCarBuilder().
		SetBrand("Toyota").
		SetModel("Land Cruiser").
		SetColor("Black").
		SetEngine("V6").
		SetSeats(7).
		Build()

	// Print the results
	fmt.Printf("Sports Car: %+v\n", sportsCar)
	fmt.Printf("SUV: %+v\n", suv)
}
```

```go 
Sports 
Car: {Brand:Ferrari Model:488 Spider Color:Red Engine:V8 Seats:2}
SUV: {Brand:Toyota Model:Land Cruiser Color:Black Engine:V6 Seats:7}  
```



*Highlight the Problem Solved: Explain how it avoids the "telescoping constructor" problem.*