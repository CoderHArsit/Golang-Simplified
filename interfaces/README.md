# Chapter 4: Interfaces ⚙️

An interface defines a set of **behaviors** (methods) that a type must have.

## 🗒️ Key Concepts

### 1. Implicit Satisfaction
In Go, you don't say `class User implements Shape`. If your `User` struct has an `Area()` method, it **automatically** becomes a `Shape`.

### 2. Polymorphism
Interfaces allow you to write functions that work with multiple different types as long as they share a common behavior.

### 3. The Empty Interface (`interface{}` or `any`)
Since an empty interface has zero methods, **every type satisfies it**. This is Go's version of the "Object" type in other languages. (Note: As of Go 1.18, `any` is an alias for `interface{}`).

### 4. Type Assertions
How to get the original concrete type back from an interface value:
```go
val, ok := i.(string) // ok is true if the interface actually contains a string
```

## 🚀 How to Run
```bash
go run interfaces/01_basics/main.go
```

---

## 🎯 Interview Questions

### Q1: How are interfaces implemented in Go? How does it differ from Java/C#?
**Answer:** Go interfaces are satisfied **implicitly** — a type implements an interface simply by having all the required methods. There's no `implements` keyword. This enables a form of **structural typing** (duck typing at compile time), making Go interfaces more flexible and decoupled.

### Q2: What is the empty interface (`interface{}` / `any`) and when would you use it?
**Answer:** The empty interface has **zero methods**, so every type satisfies it. It's Go's way of accepting "any" value (like `Object` in Java). Common uses: `fmt.Println(a ...any)`, generic containers, and JSON parsing (`map[string]interface{}`). Since Go 1.18, `any` is an alias for `interface{}`.

### Q3: What is a type assertion? What happens if it fails?
**Answer:** A type assertion extracts the concrete type from an interface value: `val := i.(string)`. If the assertion is wrong, it **panics**. Use the **comma-ok** form to avoid panics: `val, ok := i.(string)` — `ok` is `false` if the type doesn't match.

### Q4: What is a type switch in Go?
**Answer:** A type switch checks the concrete type of an interface value against multiple types:
```go
switch v := i.(type) {
case string:  fmt.Println("string:", v)
case int:     fmt.Println("int:", v)
default:      fmt.Println("unknown")
}
```
It's cleaner than multiple type assertions and commonly used when handling `interface{}` values.

### Q5: What is interface composition (embedding)?
**Answer:** Go allows you to build larger interfaces by **embedding** smaller ones. Example: `io.ReadWriter` embeds both `io.Reader` and `io.Writer`. This promotes the Go idiom: **"Accept interfaces, return structs"** and keeps interfaces small and focused.

### Q6: What is a nil interface vs. an interface holding a nil pointer?
**Answer:** This is a famous Go gotcha:
- A **nil interface** has both its type and value as nil → `i == nil` is `true`.
- An interface holding a **nil pointer** has a type but nil value → `i == nil` is `false`!
This is why you should return `nil` directly (not a typed nil pointer) for error interfaces.

### Q7: What is the "Accept interfaces, return structs" principle?
**Answer:** Functions should accept **interfaces** as parameters (for flexibility and testability) but return **concrete structs** (to avoid premature abstraction). This makes code easier to mock in tests and follows the principle of least coupling.

### Q8: How do you check if a type implements an interface at compile time?
**Answer:** Use a compile-time assertion:
```go
var _ MyInterface = (*MyStruct)(nil)
```
This doesn't execute at runtime — it just ensures `MyStruct` satisfies `MyInterface` at compile time. If it doesn't, the code won't compile.

### Q9: What are some important interfaces in Go's standard library?
**Answer:**
- `error` — `Error() string` (error handling)
- `fmt.Stringer` — `String() string` (custom print formatting)
- `io.Reader` — `Read(p []byte) (n int, err error)` (reading data)
- `io.Writer` — `Write(p []byte) (n int, err error)` (writing data)
- `sort.Interface` — `Len()`, `Less()`, `Swap()` (custom sorting)

### Q10: Can an interface have fields in Go?
**Answer:** **No.** Go interfaces can only define **method signatures**, not fields. If you need shared fields, use struct embedding. Interfaces define behavior, structs define data — this is a core Go design principle.
