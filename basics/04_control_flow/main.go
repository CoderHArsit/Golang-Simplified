package main

import "fmt"

func main() {
	// 1. If / Else
	// No parentheses around conditions, but braces are REQUIRED.
	age := 18
	if age >= 18 {
		fmt.Println("You can vote! 🗳️")
	} else if age == 17 {
		fmt.Println("Wait one more year!")
	} else {
		fmt.Println("Too young to vote.")
	}

	// 2. If with Initialization
	// You can declare a variable right in the 'if' statement! 
	// The scope of 'num' is limited to the if/else block.
	if num := 10; num > 0 {
		fmt.Printf("%d is positive!\n", num)
	}

	// 3. For Loop (Go's ONLY loop)
	fmt.Println("--- Simple For ---")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Iteration: %d\n", i)
	}

	// 4. For as While
	fmt.Println("--- For as While ---")
	count := 1
	for count <= 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}

	// 5. Infinite Loop (with break)
	fmt.Println("--- Infinite Loop with Break ---")
	x := 0
	for {
		if x >= 3 {
			break
		}
		fmt.Printf("X is %d\n", x)
		x++
	}

	// 6. Switch Statement
	// No 'break' needed! Go automatically breaks after each case.
	fmt.Println("--- Switch ---")
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the work week!")
	case "Friday":
		fmt.Println("TGIF! 🎉")
	case "Saturday", "Sunday": // Multiple matches
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Just another day...")
	}

	// 7. Switch without an expression (cleaner than multiple if/else)
	n := 15
	switch {
	case n%15 == 0:
		fmt.Println("FizzBuzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	case n%5 == 0:
		fmt.Println("Buzz")
	}
}
