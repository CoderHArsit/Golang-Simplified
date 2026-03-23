package main

import (
	"encoding/json"
	"fmt"
	"unsafe" // Used here only to demonstrate memory size. Avoid in production!
)

// 1. Memory Alignment & Padding
// Note the order of fields in these two structs.
type BadStructure struct {
	A bool  // 1 byte
	B int64 // 8 bytes
	C bool  // 1 byte
}

type GoodStructure struct {
	B int64 // 1st: Large (8 bytes)
	A bool  // 2nd: Small (1 byte)
	C bool  // 3rd: Small (1 byte)
}

// 2. Struct Tags (JSON)
type Product struct {
	Name      string  `json:"product_name"`
	Price     float64 `json:"price_usd"`
	Quantity  int     `json:"-"` // Hidden from JSON
}

func main() {
	// 3. Memory Alignment Demo
	bad := BadStructure{}
	good := GoodStructure{}
	fmt.Println("--- Struct Memory Deep Dive ---")
	fmt.Printf("Size of BadStructure (ordered bool-int64-bool): %d bytes\n", unsafe.Sizeof(bad))
	fmt.Printf("Size of GoodStructure (ordered int64-bool-bool): %d bytes\n", unsafe.Sizeof(good))
	// Result: Padding makes the first one much larger!

	// 4. Struct Equality
	fmt.Println("\n--- Struct Equality ---")
	type Point struct{ X, Y int }
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{2, 1}
	fmt.Printf("Are p1 and p2 equal? %v\n", p1 == p2)
	fmt.Printf("Are p1 and p3 equal? %v\n", p1 == p3)

	// 5. JSON Serialization with Tags
	fmt.Println("\n--- Struct Tags & JSON ---")
	prod := Product{
		Name:     "Laptop",
		Price:    1200.50,
		Quantity: 5,
	}
	jsonData, _ := json.Marshal(prod)
	fmt.Printf("Serialized JSON: %s\n", string(jsonData))
	// Notice 'Quantity' is missing and names are custom!

	// 6. Type Promotion via Embedding
	// Methods and fields are "promoted" to the outer struct.
	type Logger struct{}
	type App struct {
		Logger // Embedding
	}
	_ = App{} // We can access Logger methods directly on App instance
}
