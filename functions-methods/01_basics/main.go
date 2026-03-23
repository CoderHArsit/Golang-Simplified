package main

import (
	"errors"
	"fmt"
)

// 1. Basic Function with multiple returns
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// 2. Named Return Values (Best for documentation)
func rectangleInfo(w, h float64) (area float64, perim float64) {
	area = w * h
	perim = 2 * (w + h)
	return // Naked return (returns area and perim)
}

// 3. Variadic Function
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// --- Methods Section ---

type User struct {
	Name string
	Age  int
}

// 4. Value Receiver Method
// Gets a COPY of User. (Safe for reading)
func (u User) Greet() string {
	return fmt.Sprintf("Hello! My name is %s", u.Name)
}

// 5. Pointer Receiver Method
// Gets the ADDRESS of User. (Allows updating!)
func (u *User) HaveBirthday() {
	u.Age++
}

func main() {
	// 6. Using Functions
	fmt.Println("--- Advanced Functions ---")
	result, err := divide(10, 2)
	fmt.Printf("Divide result: %v, Error: %v\n", result, err)

	a, p := rectangleInfo(10, 5)
	fmt.Printf("Rect Area: %.2f, Perim: %.2f\n", a, p)

	fmt.Printf("Variadic Sum: %d\n", sum(1, 2, 3, 4, 5))

	// 7. Using Methods
	fmt.Println("\n--- Value vs Pointer Receivers ---")
	harshit := User{Name: "Harshit", Age: 25}
	fmt.Println(harshit.Greet())

	harshit.HaveBirthday() // Actually updates his age!
	fmt.Printf("New Age: %d\n", harshit.Age)

	// 8. Anonymous Functions & Closures
	// A closure "remembers" its environment.
	fmt.Println("\n--- Closures ---")
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println("Call 1:", counter()) // 1
	fmt.Println("Call 2:", counter()) // 2
	fmt.Println("Call 3:", counter()) // 3
}
