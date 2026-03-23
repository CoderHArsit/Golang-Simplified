# Chapter 1: Basics 🏁

This chapter covers the absolute fundamentals of Go.

## 🗒️ Topics Covered
1. **Hello World**: How to structure a simple Go program and run it.
2. **Packages**: The fundamental building blocks of Go programs.
3. **Variables**: Declaring and initializing variables using `var` and `:=`.
4. **Data Types**: Integers, floats, strings, and booleans.
5. **Constants**: Defining immutable values with `const`.

## 📦 Understanding Packages
In Go, every file belongs to a **package**. It's how Go organizes and reuses code.
- **`package main`**: Tells Go this is an executable program, not a library.
- **`main()` function**: The entry point of a `main` package.
- **Multiple Files**: All files in a single directory must share the same package name.
- **Visibility**: 
    - Names starting with **Uppercase** (e.g., `fmt.Println`) are exported (public).
    - Names starting with **lowercase** are unexported (private).
- **Imports**: Use `import "package_name"` to bring in functionality.

## Library vs Package
In Go, these terms are often used interchangeably, but there's a technical distinction:
1. **Package**: A single directory containing one or more `.go` files. It's the unit of **code organization**.
2. **Library / Module**: A collection of related packages released together as a unit. It's the unit of **distribution and versioning** (defined by your `go.mod` file).

**Analogy**:
- A **Package** is like a **chapter** in a book.
- A **Module/Library** is the **entire book** itself.

## 🔄 Why Only One Loop?
Go **only** has a `for` loop. It DOES NOT have `while` or `do-while`. This is a deliberate choice to keep the language simple and consistent.

| Type | Other Languages | Go Equivalent |
| :--- | :--- | :--- |
| **Standard** | `for (i=0; i<3; i++)` | `for i := 0; i < 3; i++` |
| **While** | `while (condition)` | `for condition` |
| **Infinite** | `while (true)` | `for { ... }` |
| **Do-While** | `do { ... } while (cond)` | `for { ... if !cond { break } }` |

## 🚀 How to Run
To run any example:
```bash
go run basics/01_hello_world/main.go
# or
go run basics/02_variables_and_types/main.go
# or
go run basics/03_constants_and_iota/main.go
# or
go run basics/04_control_flow/main.go
```

## 📝 Key Takeaways
- Every Go file must start with a `package` declaration.
- The `main` package and `main` function are the entry point of a program.
- Go is statically typed, meaning variable types are checked at compile time.
- The short variable declaration `:=` can only be used inside functions.
