# Chapter 2: Data Structures 🏗️

This chapter covers how Go handles data collections and custom types.

## 🗒️ Topics Covered
1. **Arrays**: Fixed-size sequences of elements.
2. **Slices**: Powerful, flexible, and dynamic views into arrays.
3. **Maps**: Fast key-value storage.
4. **Structs**: Building custom types by grouping data.

## 🚀 How to Run
To run any example:
```bash
go run data-structures/01_arrays_and_slices/main.go
# or
go run data-structures/02_slices_deep_dive/main.go
# or
go run data-structures/03_maps/main.go
# or
go run data-structures/04_structs/main.go
# or
go run data-structures/05_structs_deep_dive/main.go
```

## 🧠 Deep Dive: Arrays vs Slices

### 1. The Core Difference
| Feature | Array | Slice |
| :--- | :--- | :--- |
| **Size** | Fixed (part of the type) | Dynamic (can grow/shrink) |
| **Type** | `[n]T` (e.g., `[5]int`) | `[]T` (e.g., `[]int`) |
| **Memory** | Allocated on the stack (usually) | Allocated on the heap (usually) |
| **Passing** | Passed by **Value** (copied!) | Passed by **Reference** (header is copied) |
| **Structure** | Continuous block of elements | A **Descriptor** pointing to an underlying array |

### 2. The Slice Header
Under the hood, a slice is a 24-byte structure (on 64-bit systems) called a **Header**:
- **Pointer**: Address of the first element in the underlying array.
- **Length (`len`)**: Number of elements currently in the slice.
- **Capacity (`cap`)**: Number of elements in the underlying array starting from the pointer.

### 3. Memory Allocation & `append`
When you use `append()`:
1.  Go checks if `len + 1 <= cap`.
2.  **If yes**: It simply adds the element and increments the length.
3.  **If no**: It allocates a **new, larger array** (typically double the previous capacity), copies the old elements over, and returns the new slice pointer.

**Pro-Tip**: Use `make([]T, len, cap)` to pre-allocate memory if you know how many items you'll have. This prevents expensive reallocations.

## 🏗️ Deep Dive: Structs

### 1. Memory Alignment (Padding)
The order of fields in a struct matters for performance and memory usage! Go aligns fields to memory boundaries (typically 8 bytes on 64-bit systems).
- **Poor Alignment**: Mixing large and small types (e.g., `bool`, `int64`, `bool`) can cause "padding" bytes to be added.
- **Good Alignment**: Grouping the largest types first can reduce the total size of the struct.

### 2. Struct Tags (Metadata)
Struct tags are strings associated with struct fields. They are used by packages like `encoding/json` to control how fields are serialized.
```go
type User struct {
    Name string `json:"user_name"` // Maps "Name" to "user_name" in JSON
    Age  int    `json:"-"`         // Ignores this field in JSON
}
```

### 3. Equality
Two structs are equal (`==`) if:
- They are of the same **type**.
- All of their **fields** are comparable and equal.
- *Note*: If a struct contains a slice or a map, it is **not comparable** and will cause a compile error if you try to use `==`.

---

## 🎯 Interview Questions

### Q1: What is the difference between an array and a slice in Go?
**Answer:** An **array** has a fixed size that's part of its type (`[5]int`), is passed by **value** (copied), and lives on the stack. A **slice** is a dynamic view into an underlying array (`[]int`), is passed by **reference** (header copied), and can grow using `append()`.

### Q2: What is a slice header? What does it contain?
**Answer:** A slice header is a 24-byte struct (on 64-bit systems) containing three fields:
- **Pointer** — address of the first element in the underlying array
- **Length (`len`)** — number of elements currently in the slice
- **Capacity (`cap`)** — total elements available from the pointer to the end of the underlying array

### Q3: What happens when you `append()` to a slice that's at capacity?
**Answer:** Go allocates a **new, larger underlying array** (typically 2x the previous capacity for small slices), copies all existing elements to the new array, appends the new element, and returns a new slice header pointing to the new array. The old array becomes eligible for GC.

### Q4: Are slices passed by reference or by value in Go?
**Answer:** Technically, **by value** — Go copies the slice header (pointer, length, capacity). But since the header contains a pointer to the underlying array, modifications to the elements are visible to the caller. However, `append()` that triggers reallocation won't be seen by the caller unless the new slice is returned.

### Q5: What happens if you read a key that doesn't exist in a Go map?
**Answer:** Go returns the **zero value** for that type (e.g., `0` for `int`, `""` for `string`). To distinguish between a missing key and a zero value, use the **comma-ok idiom**: `val, ok := m["key"]` — `ok` is `false` if the key doesn't exist.

### Q6: Are Go maps ordered?
**Answer:** **No.** Map iteration order is **intentionally randomized** by the Go runtime. If you need ordered keys, you must sort the keys yourself (e.g., collect keys into a slice, sort it, then iterate).

### Q7: Can you compare two structs with `==` in Go?
**Answer:** Only if **all fields are comparable types**. Structs with fields of type `slice`, `map`, or `func` are **not comparable** and will cause a compile error. For those, use `reflect.DeepEqual()` (but it's slower).

### Q8: What is memory alignment (padding) in Go structs? Why does field order matter?
**Answer:** Go aligns struct fields to memory boundaries (typically 8 bytes on 64-bit). Poor field ordering (e.g., `bool, int64, bool`) causes the compiler to insert **padding bytes**, increasing struct size. Grouping fields by decreasing size (largest first) minimizes padding and reduces memory usage.

### Q9: What are struct tags and how are they used?
**Answer:** Struct tags are string metadata attached to struct fields using backtick syntax. They're used by packages like `encoding/json` to control serialization. Example: `` `json:"user_name,omitempty"` `` maps the field to `"user_name"` in JSON and omits it if empty. Tags are accessed at runtime via the `reflect` package.

### Q10: What is the difference between `make()` and `new()` in Go?
**Answer:**
- `new(T)` — allocates zeroed memory for type `T` and returns a **pointer** (`*T`). Works for any type.
- `make(T)` — only works for **slices, maps, and channels**. It initializes the internal data structure and returns the type itself (not a pointer). Use `make([]int, 0, 10)` to pre-allocate a slice with capacity 10.
