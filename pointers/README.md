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

---

## 🎯 Interview Questions

### Q1: What is a pointer in Go?
**Answer:** A pointer is a variable that stores the **memory address** of another variable. It is declared using `*T` syntax (e.g., `*int`). You get the address of a variable with `&` and access the value at an address with `*` (dereferencing).

### Q2: Does Go support pointer arithmetic like C?
**Answer:** **No.** Go intentionally prohibits pointer arithmetic (`p++`, `p + offset`) to prevent memory safety bugs like buffer overflows and dangling pointers. If you need low-level memory manipulation, use the `unsafe` package (strongly discouraged in production).

### Q3: What is a nil pointer and what happens if you dereference it?
**Answer:** A nil pointer doesn't point to any valid memory address (its zero value is `nil`). Dereferencing a nil pointer causes a **runtime panic**: `invalid memory address or nil pointer dereference`. Always check `if p != nil` before dereferencing.

### Q4: Is Go pass-by-value or pass-by-reference?
**Answer:** Go is **always pass-by-value**. Everything is copied when passed to a function. However, you can achieve pass-by-reference semantics by passing a **pointer** to the value. Slices, maps, and channels appear to be "by reference" because their internal structures contain pointers, but the header itself is still copied.

### Q5: When should you use a pointer vs a value?
**Answer:**
- **Use a pointer** when: you need to modify the original value, the data is large (avoids expensive copies), or you need to represent "no value" with `nil`.
- **Use a value** when: the data is small (like `int`, `bool`), you want immutability, or you want simpler, safer code.

### Q6: What is the `&` operator and the `*` operator?
**Answer:**
- `&` (address-of) — returns the memory address of a variable: `p := &x`
- `*` (dereference) — accesses the value stored at a pointer's address: `fmt.Println(*p)`
- `*` is also used in type declarations: `var p *int` means "p is a pointer to an int"

### Q7: Can a function return a pointer to a local variable in Go?
**Answer:** **Yes!** Unlike C, this is safe in Go. The compiler performs **escape analysis** — if it detects that a local variable's address escapes the function (is returned or stored), it allocates that variable on the **heap** instead of the stack. The garbage collector handles cleanup.

### Q8: What is the difference between a value receiver and a pointer receiver on methods?
**Answer:**
- **Value receiver `(v T)`** — gets a copy; cannot modify the original struct.
- **Pointer receiver `(v *T)`** — gets the address; can modify the original struct and avoids copying (more efficient for large structs).
- Rule: If any method uses a pointer receiver, **all methods** on that type should use pointer receivers for consistency.

### Q9: What does `new()` do in Go?
**Answer:** `new(T)` allocates zeroed memory for a value of type `T` and returns a **pointer** to it (`*T`). It's rarely used in idiomatic Go — most developers prefer `&T{}` (address-of composite literal) instead, as it also allows field initialization.

### Q10: How does Go's garbage collector interact with pointers?
**Answer:** Go has a **concurrent, tri-color mark-and-sweep garbage collector**. It traces pointers from roots (stack, globals) to determine which heap objects are reachable. Unreachable objects are freed automatically. You don't need to manually free memory — the GC handles it. This is why pointer arithmetic is disallowed (it would break GC's ability to trace references).
