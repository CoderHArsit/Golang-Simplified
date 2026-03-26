package main

import (
	"errors"
	"fmt"
	"os"
)

// 1. Custom Error using errors.New
var ErrDivideByZero = errors.New("cannot divide by zero")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// 2. Custom Error using fmt.Errorf (with context)
func openFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		// %w wraps the original error so we can inspect it later
		return nil, fmt.Errorf("openFile(%q) failed: %w", name, err)
	}
	return data, nil
}

func main() {
	// 3. The Classic: if err != nil
	fmt.Println("--- Basic Error Handling ---")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 4. Comparing Errors with errors.Is()
	fmt.Println("\n--- errors.Is() ---")
	if errors.Is(err, ErrDivideByZero) {
		fmt.Println("Caught a specific error: ErrDivideByZero")
	}

	// 5. Wrapped Errors
	fmt.Println("\n--- Wrapped Errors ---")
	_, err = openFile("nonexistent.txt")
	if err != nil {
		fmt.Println("Wrapped error:", err)

		// Unwrap to check the original cause
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			fmt.Printf("Original cause → Op: %s, Path: %s\n", pathErr.Op, pathErr.Path)
		}
	}

	// 6. Defer Basics
	fmt.Println("\n--- Defer Order (LIFO) ---")
	fmt.Println("Start")
	defer fmt.Println("Defer 1 (first defer, runs last)")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3 (last defer, runs first)")
	fmt.Println("End of main logic")
	// Output order: Start → End → Defer 3 → Defer 2 → Defer 1
}
