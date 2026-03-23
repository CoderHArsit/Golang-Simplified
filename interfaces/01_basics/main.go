package main

import "fmt"

// 1. Defining an Interface
// We define a behavior: "Anything that can area and perimeter is a Shape"
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. Concrete Type: Rect
type Rect struct {
	Width, Height float64
}

// 3. Implementing the methods
// Notice NO 'implements' keyword!
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 4. Concrete Type: Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 5. Polymorphic Function
// This function doesn't care if it's a Rect or a Circle!
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	r := Rect{Width: 10, Height: 5}
	c := Circle{Radius: 7}

	fmt.Println("--- Interface Polymorphism ---")
	printShapeInfo(r) // Rect works
	printShapeInfo(c) // Circle works

	// 6. The Empty Interface (any)
	fmt.Println("\n--- Empty Interface & Type Assertions ---")
	var anything interface{} = "I am a string"

	// Type Assertion (Manual casting)
	str, ok := anything.(string)
	if ok {
		fmt.Printf("Value is a string: %q\n", str)
	}

	// 7. Type Switch
	checkType(42)
	checkType("Hello")
	checkType(true)
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("It's an int: %d\n", v)
	case string:
		fmt.Printf("It's a string: %s\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
