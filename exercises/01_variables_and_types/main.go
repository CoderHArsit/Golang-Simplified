package main

import "fmt"

func main() {
	// Challenge: 
	// 1. Declare an integer variable 'age' and assign it 25.
	// 2. Declare a string variable 'name' and assign it your name.
	// 3. Declare a float variable 'height' and assign it your height in meters.
	// 4. Print all variables in a formatted string.

	age := 25
	name := "Harshit"
	height := 1.75

	fmt.Printf("My name is %s, I am %d years old, and my height is %.2f meters.\n", name, age, height)
}
