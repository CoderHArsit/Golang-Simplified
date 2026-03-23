# Chapter 3: Pointers 📍

Pointers are variables that store the memory address of another value.

## 🗒️ Why Pointers?
1. **Efficiency**: Passing a large struct by pointer avoids copying the entire structure.
2. **Mutability**: Passing by pointer allows a function to modify the original value.

## 🧠 In-Depth Concepts

### 1. The Operators
- **`&` (Address-of)**: Gets the memory address of a variable.
- **`*` (Dereference)**: Accesses the value stored at a memory address.

### 2. Pointer Declaration
```go
var p *int // p is a pointer to an integer
```

### 3. Safety First
Unlike C, Go **does not allow pointer arithmetic** (like `p++`). This prevents many common memory safety bugs.

### 4. Nil Pointers
A pointer that doesn't point to anything is `nil`. Always check for `nil` before dereferencing to avoid panics!

## 🚀 How to Run
```bash
go run pointers/01_basics/main.go
```
