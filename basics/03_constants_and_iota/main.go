package main

import "fmt"

func main() {
	// 1. Basic Constants
	// Constants are declared with the 'const' keyword and cannot be changed.
	const Pi = 3.14159
	const CompanyName = "Learning Go Inc."

	// 2. Typed vs Untyped Constants
	const UntypedInt = 42          // Inferred as int
	const TypedInt int32 = 42       // Explicitly int32

	// 3. Iota - The Enumerator
	// 'iota' is used to create a sequence of related constants.
	// It starts at 0 and increments for each line.
	const (
		Sunday = iota // 0
		Monday        // 1
		Tuesday       // 2
		Wednesday     // 3
		Thursday      // 4
		Friday        // 5
		Saturday      // 6
	)

	// 4. Using Iota for Bitwise Flags (Advanced)
	const (
		Read    = 1 << iota // 1 << 0 = 1
		Write               // 1 << 1 = 2
		Execute             // 1 << 2 = 4
	)

	// 5. Skipping values with '_'
	const (
		_ = iota // Skip 0
		KB = 1 << (10 * iota) // 1 << (10 * 1) = 1024
		MB                    // 1 << (10 * 2) = 1048576
	)

	fmt.Println("--- Learning Constants & Iota ---")
	fmt.Printf("Pi: %f, Welcome to %s\n", Pi, CompanyName)
	fmt.Printf("TypedInt: %T (%v), UntypedInt: %T (%v)\n", TypedInt, TypedInt, UntypedInt, UntypedInt)
	fmt.Printf("Days: Sun(%d), Mon(%d), Tue(%d)\n", Sunday, Monday, Tuesday)
	fmt.Printf("Perms: Read(%d), Write(%d), Execute(%d)\n", Read, Write, Execute)
	fmt.Printf("Storage: KB(%d bytes), MB(%d bytes)\n", KB, MB)
}
