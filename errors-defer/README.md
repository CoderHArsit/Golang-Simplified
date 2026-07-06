# Chapter 6: Errors & Defer 🛡️

Go handles errors explicitly — there are no try/catch blocks. This forces you to deal with errors at every step.

## 🗒️ Key Concepts

### 1. Errors are Values
In Go, an error is just a value that implements the `error` interface:
```go
type error interface {
    Error() string
}
```

### 2. The `if err != nil` Pattern
This is the most common pattern in Go code. Every function that can fail returns an `error` as its last return value.
```go
result, err := doSomething()
if err != nil {
    // handle the error
}
```

### 3. `defer` — Guaranteed Cleanup
`defer` schedules a function call to run when the surrounding function returns. Perfect for closing files, database connections, and HTTP responses.
- Executes in **LIFO** (Last In, First Out) order.
- Arguments are evaluated **immediately**, not at execution time.

### 4. `panic` and `recover`
- **`panic`**: Stops normal execution. Use only for truly unrecoverable errors (like a corrupt config at startup).
- **`recover`**: Catches a panic inside a `defer` and lets the program continue.
- **Rule**: Don't use `panic` for normal error handling. Use `error` values instead.

### 5. Custom Errors
You can create your own error types with additional context using `fmt.Errorf`, `errors.New`, or by implementing the `error` interface on a struct.

### 6. Error Wrapping (Go 1.13+)
Wrap errors for context while preserving the original:
```go
return fmt.Errorf("failed to open config: %w", err)
```
Then unwrap with `errors.Is()` and `errors.As()`.

## 🚀 How to Run
```bash
go run errors-defer/01_basics/main.go
# or
go run errors-defer/02_advanced/main.go
```

---

## 🎯 Interview Questions

### Q1: How does Go handle errors? Why doesn't it use try/catch?
**Answer:** Go handles errors **explicitly** by returning an `error` value as the last return value. There's no try/catch/finally. This is by design — it forces developers to handle errors at every call site, making error flow visible and preventing silently swallowed exceptions.

### Q2: What is the `error` interface in Go?
**Answer:** `error` is a built-in interface with a single method:
```go
type error interface {
    Error() string
}
```
Any type that implements `Error() string` is an error. You can create simple errors with `errors.New("msg")` or `fmt.Errorf("msg: %v", detail)`.

### Q3: What is error wrapping and how do you use `%w`?
**Answer:** Error wrapping (Go 1.13+) adds context while preserving the original error:
```go
return fmt.Errorf("failed to open config: %w", err)
```
Use `errors.Is(err, target)` to check if any error in the chain matches a specific error, and `errors.As(err, &target)` to check if any error in the chain matches a specific type.

### Q4: What is the difference between `errors.Is()` and `errors.As()`?
**Answer:**
- `errors.Is(err, target)` — checks if any error in the chain **equals** a specific error value (for sentinel errors like `io.EOF`).
- `errors.As(err, &target)` — checks if any error in the chain **matches a specific type** and extracts it (for custom error types with extra fields).

### Q5: What is a sentinel error? Give examples.
**Answer:** A sentinel error is a pre-defined, package-level error variable used for comparison:
```go
var ErrNotFound = errors.New("not found")
```
Standard library examples: `io.EOF`, `sql.ErrNoRows`, `os.ErrNotExist`. Convention: prefix with `Err` (e.g., `ErrTimeout`).

### Q6: Explain `defer` execution order and argument evaluation.
**Answer:**
- **LIFO order**: Multiple defers execute in reverse order (last defer runs first).
- **Immediate argument evaluation**: Arguments are evaluated when `defer` is called, NOT when the deferred function executes.
```go
x := 10
defer fmt.Println(x) // Prints 10, not 20!
x = 20
```

### Q7: When should you use `panic` vs returning an error?
**Answer:**
- **Return an error** for expected failures: file not found, network timeout, invalid input — anything the caller can handle.
- **Use `panic`** only for truly unrecoverable situations: nil pointer in invariant code, impossible state, or corrupt initialization at startup.
- Rule: **Libraries should never panic**. Only `main` or init code may panic in rare cases.

### Q8: How does `recover` work? Where must it be called?
**Answer:** `recover()` catches a panic and returns the panic value. It **must** be called inside a `defer` function — calling it outside a deferred function has no effect. After recovery, the program continues from the deferred function's return point.
```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
```

### Q9: How do you create a custom error type with additional context?
**Answer:** Implement the `error` interface on a struct:
```go
type ValidationError struct {
    Field   string
    Message string
}
func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}
```
Use `errors.As()` to extract the custom error and access its fields.

### Q10: What is the idiomatic way to handle errors in Go?
**Answer:** The standard pattern:
```go
result, err := doSomething()
if err != nil {
    return fmt.Errorf("context: %w", err) // wrap and propagate
}
// use result
```
**Don'ts**: Don't ignore errors with `_`, don't panic for recoverable errors, don't log-and-return (choose one). **Do**: wrap errors with context, use sentinel errors for expected cases, and handle errors at the appropriate level.
