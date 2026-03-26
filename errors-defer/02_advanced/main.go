package main

import (
	"fmt"
	"os"
)

// ============================================================
// ADVANCED: Custom Error Types, Panic/Recover, Defer with Files
// ============================================================

// 1. Custom Error Type (Struct implementing the error interface)
type ValidationError struct {
	Field   string
	Message string
}

// Implementing the error interface
func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on '%s': %s", v.Field, v.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "cannot be negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically high"}
	}
	return nil
}

// 2. Defer with File Handling (Real-World Pattern)
func writeToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	// defer ensures the file is ALWAYS closed, even if Write fails!
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("write to file: %w", err)
	}

	fmt.Printf("Successfully wrote to %s\n", filename)
	return nil
}

// 3. Panic & Recover
// panic = "this should NEVER happen, abort everything"
// recover = "catch the panic inside a defer, save the program"
func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("About to panic...")
	panic("something went terribly wrong!")
	// Code after panic never runs
}

func main() {
	// 4. Custom Error Types
	fmt.Println("--- Custom Error Types ---")
	err := validateAge(-5)
	if err != nil {
		fmt.Println("Error:", err)

		// Type assert to get the specific error details
		var ve *ValidationError
		if ok := fmt.Errorf(""); ok != nil {
			// Using type assertion directly
			ve, ok := err.(*ValidationError)
			if ok {
				fmt.Printf("  Field: %s, Message: %s\n", ve.Field, ve.Message)
			}
			_ = ve
		}
		_ = ve
	}

	err = validateAge(200)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = validateAge(25)
	if err == nil {
		fmt.Println("Age 25 is valid! ✅")
	}

	// 5. Defer with File Handling
	fmt.Println("\n--- Defer with Files ---")
	err = writeToFile("/tmp/go_learning_test.txt", "Hello from Go! 🚀")
	if err != nil {
		fmt.Println("File error:", err)
	}

	// Read it back to verify
	data, err := os.ReadFile("/tmp/go_learning_test.txt")
	if err == nil {
		fmt.Printf("Read back: %s\n", string(data))
	}

	// 6. Panic & Recover Demo
	fmt.Println("\n--- Panic & Recover ---")
	safeFunction()
	fmt.Println("Program continued after panic! 🎉")
}
