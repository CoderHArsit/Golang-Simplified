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
