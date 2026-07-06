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

---

## 🎯 Interview Questions

### Q1: What is the difference between `var` and `:=` in Go?
**Answer:** `var` can be used both inside and outside functions (package level), and requires explicit type or value. `:=` is the short variable declaration — it can **only** be used inside functions, infers the type automatically, and is more idiomatic for local variables.

### Q2: What are zero values in Go? Give examples.
**Answer:** When a variable is declared without an explicit initial value, Go assigns it a **zero value**:
- `int` → `0`
- `float64` → `0.0`
- `bool` → `false`
- `string` → `""` (empty string)
- `pointer`, `slice`, `map`, `channel`, `interface`, `function` → `nil`

### Q3: What is the difference between a package and a module in Go?
**Answer:** A **package** is a directory of `.go` files — it's the unit of code organization. A **module** (defined by `go.mod`) is a collection of related packages released together — it's the unit of versioning and distribution. Think of a package as a chapter and a module as the entire book.

### Q4: How does Go control visibility/access of identifiers?
**Answer:** Go uses **capitalization** instead of keywords like `public`/`private`. Names starting with an **uppercase** letter are exported (accessible outside the package). Names starting with a **lowercase** letter are unexported (package-private).

### Q5: Why does Go only have a `for` loop and no `while` or `do-while`?
**Answer:** This is a deliberate design choice for **simplicity**. Go's `for` loop can replicate all loop types:
- `for i := 0; i < n; i++` (standard for)
- `for condition` (while)
- `for { }` (infinite/do-while)

### Q6: What is `iota` in Go?
**Answer:** `iota` is a constant generator used in `const` blocks. It starts at `0` and auto-increments by `1` for each constant in the block. It's commonly used to create enumerations. It resets to `0` in each new `const` block.

### Q7: Can you declare a variable without using it in Go?
**Answer:** **No.** Go treats unused local variables as a **compile-time error**. This is by design to keep code clean. However, unused **package-level** variables and unused **imports** also cause compile errors (imports can be silenced with `_`).

### Q8: What is the difference between `=` and `:=`?
**Answer:** `=` is **assignment** — the variable must already be declared. `:=` is **declaration + assignment** — it declares a new variable and assigns a value in one step. `:=` cannot be used outside functions.

### Q9: What is type conversion in Go? Does Go support implicit type conversion?
**Answer:** Go does **NOT** support implicit type conversion (unlike C/Java). You must explicitly convert types: `float64(myInt)`, `int(myFloat)`. This prevents subtle bugs from automatic coercion.

### Q10: What happens if you have an unused import in Go?
**Answer:** It's a **compile-time error**. You can use the blank identifier `_` to import a package solely for its side effects (e.g., `import _ "net/http/pprof"`), or use `goimports` to auto-manage imports.
