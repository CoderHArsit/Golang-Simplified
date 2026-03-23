package main

import "fmt"

func main() {
	// 1. Initializing a Map
	// Maps are key-value pairs (hash tables).
	// Syntax: make(map[keyType]valueType)
	userAges := make(map[string]int)

	// 2. Inserting/Updating
	userAges["Alice"] = 25
	userAges["Bob"] = 30
	userAges["Charlie"] = 35

	fmt.Println("--- Reading Maps ---")
	fmt.Printf("Alice's age: %v\n", userAges["Alice"])

	// 3. Map Literals
	languages := map[string]string{
		"Go":     "Awesome",
		"Python": "Readable",
		"Java":   "Verbose",
	}
	fmt.Printf("Java is: %v\n", languages["Java"])

	// 4. Checking for Existence (The Comma Ok idiom)
	// Missing keys return the "zero value" (0 for int, "" for string).
	// To tell the difference between "missing" and "set to zero", use the second return value.
	age, exists := userAges["Dave"]
	if exists {
		fmt.Printf("Dave's age: %d\n", age)
	} else {
		fmt.Println("Dave is not in the system!") // Since Dave isn't there, it returns 0 and false
	}

	// 5. Deleting from a Map
	delete(userAges, "Bob")
	fmt.Println("\n--- After Deleting Bob ---")
	fmt.Println(userAges)

	// 6. Iterating with Range
	// IMPORTANT: Map iteration order is RANDOM in Go. Do not rely on it!
	fmt.Println("\n--- Iterating over Languages ---")
	for key, value := range languages {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// 7. Maps are Reference Types
	// Modifying a map inside a function will affect the original map.
	modifyMap(userAges)
	fmt.Println("\n--- Final Table ---")
	fmt.Println(userAges)
}

func modifyMap(m map[string]int) {
	m["NewUser"] = 100
}
