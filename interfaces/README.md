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
