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
