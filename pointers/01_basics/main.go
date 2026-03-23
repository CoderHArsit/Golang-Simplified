package main

import "fmt"

func main() {
	// 1. Basic Pointer Usage
	num := 42
	ptr := &num // ptr points to num's address

	fmt.Println("--- Basic Pointers ---")
	fmt.Printf("num value: %v\n", num)
	fmt.Printf("num address (&num): %v\n", &num)
	fmt.Printf("ptr value: %v (should match above)\n", ptr)
	fmt.Printf("ptr dereferenced (*ptr): %v\n", *ptr)

	// 2. Modifying via Pointer
	*ptr = 100 // Changes num!
	fmt.Println("\n--- Modifying via Pointer ---")
	fmt.Printf("num changed to: %v\n", num)

	// 3. Passing by Value vs Pointer
	val := 10
	fmt.Println("\n--- Pass by Value vs Pass by Pointer ---")
	
	updateByValue(val)
	fmt.Printf("After updateByValue: %v (unchanged)\n", val)

	updateByPointer(&val)
	fmt.Printf("After updateByPointer: %v (CHANGED!)\n", val)

	// 4. Pointer to a Pointer (Yes, you can!)
	doublePtr := &ptr
	fmt.Println("\n--- Pointer to a Pointer ---")
	fmt.Printf("Value via doublePtr (**doublePtr): %v\n", **doublePtr)

	// 5. The 'new' keyword
	// new(T) allocates memory and returns a *T
	pNew := new(int)
	*pNew = 777
	fmt.Println("\n--- The 'new' keyword ---")
	fmt.Printf("Value from new(int): %v\n", *pNew)

	// 6. Nil Pointers
	var nilPtr *int
	fmt.Println("\n--- Nil Pointers ---")
	fmt.Printf("nilPtr: %v\n", nilPtr)
	if nilPtr == nil {
		fmt.Println("Safe check: nilPtr is indeed nil. Do NOT dereference!")
	}
}

// updateByValue gets a COPY of the data. Original is safe.
func updateByValue(n int) {
	n = 999
}

// updateByPointer gets the ADDRESS. Modifying the address modifies the original.
func updateByPointer(n *int) {
	*n = 999
}
