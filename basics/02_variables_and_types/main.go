package main

import "fmt"

func main() {
	// 1. Variable Declaration with 'var'
	// Syntax: var <name> <type> = <value>
	var age int = 25
	var name string = "Harshit"
	var isLearning bool = true

	// 2. Type Inference
	// Go can often infer the type from the value
	var country = "India"

	// 3. Short Variable Declaration
	// Only works inside functions!
	// Syntax: <name> := <value>
	version := 1.21

	// 4. Multiple Declarations
	var (
		language = "Go"
		creator  = "Google"
	)

	// 5. Constants
	// Use 'const' for values that never change
	const Pi = 3.14159

	// 6. Zero Values
	// Variables declared without a value get a "zero value"
	var score int      // 0
	var price float64  // 0.0
	var label string   // ""
	var active bool    // false

	// Printing the results
	fmt.Println("--- Learning Basics ---")
	fmt.Printf("Name: %s, Age: %d, Country: %s\n", name, age, country)
	fmt.Printf("Learning %s? %v (Version: %.2f)\n", language, isLearning, version)
	fmt.Printf("Creator: %s, Pi: %.2f\n", creator, Pi)
	fmt.Printf("Zero Values -> Score: %d, Price: %.2f, Label: %q, Active: %v\n", score, price, label, active)
}
