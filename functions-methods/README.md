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
# or
go run functions-methods/03_method_receivers/main.go
```

## 🧠 Deep Dive: Method Receivers

### Value Receiver `(u User)` — The Photocopy
- Gets a **copy** of the struct.
- Changes inside the method are **lost** when the method returns.
- Use for: reading data, simple getters.

### Pointer Receiver `(u *User)` — The Key to the Locker
- Gets the **memory address** of the struct.
- Changes inside the method **modify the original**.
- Use for: setters, mutations, or when the struct is large (avoids expensive copies).

### Method Sets Rule
| Type | Can call |
| :--- | :--- |
| `T` (value) | Only value receiver methods |
| `*T` (pointer) | Both value AND pointer receiver methods |

**Go's Auto-Conversion**: When you call `acc.Deposit(500)` on a value, Go automatically converts it to `(&acc).Deposit(500)` for you.


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

---

## 🎯 Interview Questions

### Q1: How does Go handle multiple return values? Why is this important?
**Answer:** Go functions can return **multiple values**, which is the idiomatic way to return results alongside errors: `func divide(a, b float64) (float64, error)`. This eliminates the need for exceptions/try-catch and forces the caller to explicitly handle errors.

### Q2: What is the difference between a value receiver and a pointer receiver?
**Answer:**
- **Value receiver `(u User)`** — operates on a **copy**; cannot modify the original.
- **Pointer receiver `(u *User)`** — operates on the **original**; can modify and is more efficient for large structs.
- Rule: A `*T` value can call both value and pointer receiver methods, but a `T` value can only directly call value receiver methods (Go auto-converts for convenience).

### Q3: What are variadic functions? Give an example.
**Answer:** Variadic functions accept a variable number of arguments using `...` syntax. The arguments are received as a slice inside the function.
```go
func sum(nums ...int) int { /* nums is []int */ }
sum(1, 2, 3)        // pass individual values
sum(mySlice...)      // spread a slice
```
`fmt.Println` is a famous variadic function.

### Q4: What is a closure in Go?
**Answer:** A closure is an anonymous function that **captures and references** variables from its enclosing scope. The captured variables survive beyond the enclosing function's return. Common uses: counters, callbacks, and middleware patterns.
```go
func counter() func() int {
    count := 0
    return func() int { count++; return count }
}
```

### Q5: What is `defer` and in what order do deferred calls execute?
**Answer:** `defer` schedules a function call to run when the surrounding function returns. Multiple defers execute in **LIFO (Last In, First Out)** order. Key gotcha: arguments to deferred functions are evaluated **immediately** (at the `defer` statement), not when the deferred function runs.

### Q6: What is `init()` in Go? Can you have multiple `init()` functions?
**Answer:** `init()` is a special function that runs **automatically before `main()`**. Yes, you can have **multiple** `init()` functions — even in the same file! They run in the order they appear. Execution order: package-level variables → `init()` → `main()`. `init()` takes no arguments and returns nothing.

### Q7: What are named return values? When should you use them?
**Answer:** Named return values are pre-declared in the function signature and can be returned with a bare `return` statement:
```go
func divide(a, b float64) (result float64, err error) {
    if b == 0 { err = errors.New("division by zero"); return }
    result = a / b
    return
}
```
Use them for short functions or when they improve readability. Avoid in long functions where bare `return` is confusing.

### Q8: What is the difference between `func` as a type and a regular function declaration?
**Answer:** In Go, functions are **first-class citizens** — they can be assigned to variables, passed as arguments, and returned from other functions. `type MathFunc func(int, int) int` declares a function type. This enables higher-order functions, callbacks, and strategy patterns.

### Q9: What is the method set rule in Go?
**Answer:**
| Type | Can call |
|------|----------|
| `T` (value) | Only value receiver methods |
| `*T` (pointer) | Both value AND pointer receiver methods |

This matters for **interface satisfaction**: a value of type `T` cannot satisfy an interface requiring pointer receiver methods, but `*T` can satisfy interfaces with either.

### Q10: What happens if you use a loop variable directly inside a goroutine launched from a `for` loop?
**Answer:** Classic bug! The goroutine captures the **variable** (not the value), so all goroutines may see the **last** value of the loop variable. Fix: pass the variable as a function argument or create a local copy:
```go
for i := 0; i < 5; i++ {
    go func(id int) { fmt.Println(id) }(i) // ✅ Pass as argument
}
```
Note: Go 1.22+ fixed this by making loop variables per-iteration by default.
