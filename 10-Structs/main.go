package main

import "fmt"

// Go let us create any specific type we want (like in C and C++)
type AnimalName string

// We can use to create a specific struct name (also similar to C and C++)
type SportCar struct {
	IsSport bool
}

type Car struct {
	Name  string
	Year  int
	Color string
	// This is an example of Inheritance
	SportCar
}

// We alco can attach functions to the struct - so we can work like classes
func (c Car) info() string {
	return c.Name + " - " + c.Color
}

func main() {
	//We can create a new variable with the type we want using ont of the sintax below
	car1 := Car{"Mercedes", 2023, "Gray", SportCar{true}}
	car2 := Car{"Ferrari", 2023, "Red", SportCar{false}}
	var car3 Car
	car3.Name = "Red Bull"
	car3.Year = 2023
	car3.Color = "Blue"
	car3.SportCar.IsSport = true

	// We can access the role struct at once or part of the struct like bellow
	fmt.Println(car1.Name)
	fmt.Println(car2)
	fmt.Println(car3.Year)
	fmt.Println(car3.info())
}
