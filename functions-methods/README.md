# Chapter 5: Functions & Methods 🛠️

Functions are the building blocks of Go programs, and Methods are functions attached to specific types (like Structs).

## 🗒️ Key Concepts

### 1. Multiple Return Values
Go functions can return multiple values, which is the standard way to handle errors.
```go
func divide(a, b float64) (float64, error) { ... }
```

### 2. Variadic Functions
Functions that accept any number of trailing arguments using the `...` syntax (e.g., `fmt.Println`).

### 3. Methods: Value vs. Pointer Receivers
This is a critical Go concept:
- **Value Receiver**: Operates on a **copy** of the struct. Cannot modify the original!
- **Pointer Receiver**: Operates on the **actual** struct address. Allows modification and is more efficient for large structs.

### 4. Closures
Anonymous functions that "capture" variables from their surrounding scope. Useful for state management and callbacks.

## 🚀 How to Run
```bash
go run functions-methods/01_basics/main.go
# or
go run functions-methods/02_functions_deep_dive/main.go
```

## 🧠 Deep Dive: Functions

### 1. First-Class Functions
In Go, functions are "first-class citizens". This means you can:
- Assign them to **variables**.
- Pass them as **arguments** to other functions (Callbacks).
- **Return** them from other functions.

### 2. Higher-Order Functions
These are functions that treat other functions as data. They are common in functional programming patterns like `Filter`, `Map`, or `Reduce`.

### 3. The `defer` Keyword
`defer` is used to ensure a function call is performed later in a program's execution, usually for purposes of cleanup. 
- **LIFO Order**: Multiple `defer` calls are executed in Last-In-Function-Out order.
- **Argument Evaluation**: Arguments to a deferred function are evaluated **immediately**, not when the call executes!

### 4. `init()` Functions
Go has a special `init()` function that runs automatically before `main()`. You can have multiple `init()` functions across different files in the same package.
```
