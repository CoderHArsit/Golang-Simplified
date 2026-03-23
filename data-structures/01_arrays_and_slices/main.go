package main

import "fmt"

func main() {
	// 1. Arrays (Fixed Size)
	// You rarely use arrays directly in Go, but they are the foundation for slices.
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Cherry"
	fmt.Println("Array:", fruits)

	// 2. Slices (Dynamic Size)
	// Slices are much more common. They are declared without a size.
	numbers := []int{10, 20, 30}
	fmt.Println("Initial Slice:", numbers)

	// 3. Append to a Slice
	// 'append' returns a NEW slice (it might reallocate memory if the capacity is full)
	numbers = append(numbers, 40, 50)
	fmt.Println("After Append:", numbers)

	// 4. make() - Creating a slice with a specific length and capacity
	// make([]type, length, capacity)
	// Length = items currently in the slice.
	// Capacity = items the underlying array can hold before resizing.
	mySlice := make([]string, 2, 5)
	mySlice[0] = "Go"
	mySlice[1] = "Is"
	fmt.Printf("Make Slice: Len=%d, Cap=%d, Value=%v\n", len(mySlice), cap(mySlice), mySlice)

	// 5. Slice Expressions [low:high]
	// It's [inclusive:exclusive]
	fullList := []string{"A", "B", "C", "D", "E"}
	part := fullList[1:4] // Index 1, 2, 3 (B, C, D)
	fmt.Println("Partial Slice [1:4]:", part)

	// 6. Slices are REFERENCE Types
	// Modifying a slice actually modifies the underlying array!
	part[0] = "Z" // Changes "B" in fullList
	fmt.Println("Modified Partial:", part)
	fmt.Println("Original List changed too!:", fullList)

	// 7. Iterating with Range
	fmt.Println("--- Iterating with Range ---")
	for index, value := range fullList {
		fmt.Printf("Index %d has value %s\n", index, value)
	}
}
