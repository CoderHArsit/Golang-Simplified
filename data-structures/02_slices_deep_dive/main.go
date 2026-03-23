package main

import "fmt"

func main() {
	// 1. Arrays are Values
	// When you assign an array to another, ALL data is copied.
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1
	arr2[0] = 99
	fmt.Println("--- Arrays are Value Types ---")
	fmt.Println("Array 1:", arr1) // Still [1 2 3]
	fmt.Println("Array 2:", arr2) // [99 2 3] - Completely separate!

	// 2. Slices are References (Headers)
	// When you assign a slice, only the HEADER (pointer, len, cap) is copied.
	slice1 := []int{10, 20, 30}
	slice2 := slice1
	slice2[0] = 999
	fmt.Println("\n--- Slices are Reference Types ---")
	fmt.Println("Slice 1:", slice1) // [999 20 30] - Changed!
	fmt.Println("Slice 2:", slice2) // [999 20 30] - Shared underlying array!

	// 3. The 'append' Growth Pattern
	// Let's watch how Go reallocates memory as we grow a slice.
	fmt.Println("\n--- Watching Memory Growth ---")
	growth := make([]int, 0)
	for i := 0; i < 10; i++ {
		growth = append(growth, i)
		fmt.Printf("Element %d -> Len: %d, Cap: %d\n", i, len(growth), cap(growth))
	}
	// Notice how Capacity doubles (or grows by a set factor) when Length exceeds it.

	// 4. Sub-slicing and Capacity
	// Creating a slice from another slice shares the SAME underlying array.
	source := []int{0, 1, 2, 3, 4, 5}
	sub := source[2:4] // Index 2, 3 (values 2, 3)
	fmt.Println("\n--- Sub-slicing Memory ---")
	fmt.Printf("Sub-slice: %v, Len: %d, Cap: %d\n", sub, len(sub), cap(sub))
	// Capacity is measured from the START of the sub-slice to the END of the source.

	// 5. Caution: Memory Leak with Sub-slicing
	// If you have a huge array and take a tiny sub-slice, 
	// the entire huge array stays in memory because the sub-slice still points to it!
	// Solution: Use 'copy()' to allocate a fresh, smaller underlying array.
}
