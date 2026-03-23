package main

import "fmt"

// 1. init() runs BEFORE main()
// Use it for setup or global variable initialization.
func init() {
	fmt.Println("--- [INIT] Running pre-setup ---")
}

// 2. Function as a Type
// We can define a custom type for a function signature!
type mathOp func(int, int) int

func add(a, b int) int { return a + b }
func mult(a, b int) int { return a * b }

// 3. Higher-Order Function
// Takes another function as an argument.
func calculate(a, b int, op mathOp) int {
	return op(a, b)
}

// 4. Recursion
func factorial(n int) int {
	if n == 0 { // Base case
		return 1
	}
	return n * factorial(n-1) // Recursive call
}

func main() {
	fmt.Println("--- Functions Deep Dive ---")

	// 5. First-Class Functions (Anonymous)
	myFunc := func(msg string) {
		fmt.Printf("Anonymous says: %s\n", msg)
	}
	myFunc("Hello from a variable!")

	// 6. Callback Example
	resAdd := calculate(10, 5, add)
	resMult := calculate(10, 5, mult)
	fmt.Printf("Calculation (Add): %d, (Mult): %d\n", resAdd, resMult)

	// 7. Recursion Demo
	fmt.Printf("Factorial of 5: %d\n", factorial(5))

	// 8. Defer Mastery (LIFO)
	// Notice the order they print!
	fmt.Println("\n--- Defer Stack ---")
	defer fmt.Println("Defer 1 (I was called first, but I finish last!)")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3 (Executed first!)")

	// 9. Defer Argument Evaluation
	// Captures the value of 'x' at the time of the defer call, NOT when it runs.
	x := 10
	defer fmt.Printf("Value of x in defer: %d (even though it's 20 now!)\n", x)
	x = 20

	fmt.Println("Main logic is ending now...")
}
