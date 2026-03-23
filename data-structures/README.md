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
