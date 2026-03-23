package main

import "fmt"

// 1. Defining a Struct
// A struct is a collection of fields.
type User struct {
	FirstName string
	LastName  string
	Age       int
	IsActive  bool
}

// 2. Embedded Structs (Composition)
// Go uses embedding instead of traditional inheritance.
type Admin struct {
	User        // All fields from User are now part of Admin
	AccessLevel int
}

func main() {
	// 3. Initializing with Literals
	u1 := User{
		FirstName: "Harshit",
		LastName:  "Saxena",
		Age:       25,
		IsActive:  true,
	}
	fmt.Println("--- Simple Struct ---")
	fmt.Printf("User 1: %+v\n", u1)

	// 4. Zero Values
	// If you omit fields, they get their zero value.
	u2 := User{FirstName: "John"}
	fmt.Printf("User 2 (Partial): %+v\n", u2)

	// 5. Accessing Fields
	u1.Age = 26
	fmt.Println("Updated Age:", u1.Age)

	// 6. Embedded Struct usage
	admin := Admin{
		User: User{
			FirstName: "Super",
			LastName:  "Admin",
			Age:       99,
		},
		AccessLevel: 1,
	}
	// You can access embedded fields directly!
	fmt.Println("\n--- Embedded Struct ---")
	fmt.Printf("Admin Name: %s %s (Access Level: %d)\n", admin.FirstName, admin.LastName, admin.AccessLevel)

	// 7. Pointers to Structs
	// Very common in Go for performance and mutability.
	uPtr := &u1
	uPtr.Age = 30 // No need to dereference with (*uPtr).Age! Go handles it.
	fmt.Printf("User 1 Age via Pointer: %d\n", u1.Age)

	// 8. Anonymous Structs
	// Useful for one-time configurations or API responses.
	config := struct {
		APIKey string
		Port   int
	}{
		APIKey: "secret-key",
		Port:   8080,
	}
	fmt.Println("\n--- Anonymous Struct ---")
	fmt.Printf("Config: %+v\n", config)
}
