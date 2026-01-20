# Go Slice - 🚀

## সূচিপত্র
1. [Slice কী এবং কেন দরকার?](#slice-কী-এবং-কেন-দরকার)
2. [Slice vs Array - মূল পার্থক্য](#slice-vs-array---মূল-পার্থক্য)
3. [Slice-এর Internal Structure](#slice-এর-internal-structure)
4. [কোড #১: Slice Basics এবং Underlying Array](#কোড-১-slice-basics-এবং-underlying-array)
5. [কোড #২: Slice তৈরির বিভিন্ন উপায়](#কোড-২-slice-তৈরির-বিভিন্ন-উপায়)
6. [কোড #৩: Slice Copy এবং Append Behavior](#কোড-৩-slice-copy-এবং-append-behavior)
7. [কোড #৪: Function-এ Slice Pass করা](#কোড-৪-function-এ-slice-pass-করা)
8. [কোড #৫: Variadic Function](#কোড-৫-variadic-function)
9. [আরো উদাহরণ এবং Best Practices](#আরো-উদাহরণ-এবং-best-practices)

---

## Slice কী এবং কেন দরকার?

### সমস্যা: Array-র সীমাবদ্ধতা

ধরো তুমি একটি program লিখছো যেখানে user কতগুলো number input দেবে তা আগে থেকে জানো না। Array দিয়ে এটা করা কঠিন:

```go
// Array - Fixed size
var numbers [5]int  // শুধু 5টি number রাখা যাবে
// যদি user 10টি number দিতে চায়? 😰
```

### সমাধান: Slice - Dynamic Array

```go
// Slice - Dynamic size
var numbers []int  // যতগুলো খুশি number রাখা যাবে ✨
numbers = append(numbers, 1)
numbers = append(numbers, 2)
numbers = append(numbers, 3)
// ... যত খুশি যোগ করতে পারো!
```

### Slice-এর মূল বৈশিষ্ট্য:
- ✅ **Dynamic size** - runtime-এ size বাড়ানো যায়
- ✅ **Flexible** - slicing operation সহজ
- ✅ **Efficient** - memory-তে array শেয়ার করে
- ✅ **Built-in functions** - `append()`, `copy()`, `len()`, `cap()`

---

## Slice vs Array - মূল পার্থক্য

### তুলনা টেবিল:

| বৈশিষ্ট্য | Array | Slice |
|---------|-------|-------|
| **Declaration** | `[5]int` (size আছে) | `[]int` (size নেই) |
| **Size** | Fixed (compile time-এ ঠিক) | Dynamic (runtime-এ পরিবর্তন হয়) |
| **Type** | Value type | Reference type |
| **Memory** | Stack-এ পুরো array | Pointer + metadata |
| **Copy** | পুরো array copy হয় | শুধু pointer copy হয় |
| **Function pass** | Expensive (সব data copy) | Cheap (pointer copy) |

### উদাহরণ দিয়ে বুঝি:

```go
package main
import "fmt"

func main() {
    // ===== ARRAY Example =====
    arr1 := [3]int{1, 2, 3}
    arr2 := arr1  // পুরো array copy হয়
    arr2[0] = 100
    
    fmt.Println("arr1:", arr1)  // Output: [1 2 3] (unchanged)
    fmt.Println("arr2:", arr2)  // Output: [100 2 3]
    
    // ===== SLICE Example =====
    slice1 := []int{1, 2, 3}
    slice2 := slice1  // শুধু pointer copy হয়
    slice2[0] = 100
    
    fmt.Println("slice1:", slice1)  // Output: [100 2 3] (changed!)
    fmt.Println("slice2:", slice2)  // Output: [100 2 3]
}
```

**ব্যাখ্যা:**
- Array: `arr1` এবং `arr2` দুটি আলাদা memory location
- Slice: `slice1` এবং `slice2` একই underlying array point করছে

---

## Slice-এর Internal Structure

### Slice আসলে কী?

Slice হলো একটি **struct** যার মধ্যে ৩টি field আছে:

```go
type slice struct {
    array unsafe.Pointer  // underlying array-র pointer
    len   int            // বর্তমান length
    cap   int            // total capacity
}
```

### Visual Representation:

```
Slice:
┌──────────────┬─────────┬──────────┐
│   Pointer    │  len=3  │  cap=5   │
└──────┬───────┴─────────┴──────────┘
       │
       └─────> Underlying Array:
               ┌───┬───┬───┬───┬───┐
               │ 1 │ 2 │ 3 │ ? │ ? │
               └───┴───┴───┴───┴───┘
                 0   1   2   3   4
```

### বিস্তারিত উদাহরণ:

```go
package main
import "fmt"

func main() {
    // একটি slice তৈরি করি
    s := []int{10, 20, 30, 40, 50}
    
    fmt.Printf("Slice: %v\n", s)
    fmt.Printf("Length: %d\n", len(s))
    fmt.Printf("Capacity: %d\n", cap(s))
    fmt.Printf("Address of first element: %p\n", &s[0])
    
    // এখন একটি sub-slice তৈরি করি
    sub := s[1:4]  // [20, 30, 40]
    
    fmt.Printf("\nSub-slice: %v\n", sub)
    fmt.Printf("Length: %d\n", len(sub))
    fmt.Printf("Capacity: %d\n", cap(sub))
    fmt.Printf("Address of first element: %p\n", &sub[0])
    
    // লক্ষ্য করো: sub[0] এবং s[1] একই memory location!
}
```

**Output:**
```
Slice: [10 20 30 40 50]
Length: 5
Capacity: 5
Address of first element: 0xc000014078

Sub-slice: [20 30 40]
Length: 3
Capacity: 4
Address of first element: 0xc000014080  (s[1] এর address)
```

**কেন Capacity = 4?**
```
Original: [10 | 20 | 30 | 40 | 50]
           0    1    2    3    4
           
sub starts at index 1
sub can grow till index 4
So capacity = 5 - 1 = 4
```

---

## কোড #১: Slice Basics এবং Underlying Array

```go
package main

import "fmt"

func main() {
    arr := [6]string{"This", "is", "a", "Go", "Intro", "slice"}
    fmt.Println(arr)

    //slice is a reference to an array
    //slice = pointer + length + capacity

    sl := arr[1:4] //[is a go]

    //pointer =  &arr[1] means the slice starts at index 1 of the array
    //length = 4-1 = 3 means the slice has 3 elements
    //capacity = len(arr) - 1 = 6 - 1 = 5 means the slice can grow up to 5 elements before needing to allocate new memory

    fmt.Println(sl)

    sl1 := sl[1:2]

    fmt.Println(sl1)
    fmt.Println(len(sl1)) //length
    fmt.Println(cap(sl1)) //capacity
    fmt.Println(*&sl1[0])
}
```

### 🔍 বিস্তারিত ব্যাখ্যা:

#### Step 1: Array তৈরি করা
```go
arr := [6]string{"This", "is", "a", "Go", "Intro", "slice"}
```

**Memory-তে:**
```
arr: 
Index: 0      1     2    3     4        5
      [This | is  | a  | Go | Intro | slice]
Memory: 0x100 0x108 0x110 0x118 0x120  0x128  (hypothetical addresses)
```

#### Step 2: প্রথম Slice তৈরি
```go
sl := arr[1:4]  // Start index 1, End BEFORE index 4
```

**Slicing Syntax বুঝি:**
```go
array[start:end]
// start = inclusive (এটা included)
// end = exclusive (এটা included না)
```

**Memory Representation:**
```
arr: [This | is  | a  | Go | Intro | slice]
       0     1     2    3     4       5
              ↑         ↑
              start     end (not included)
              
sl points to: [is | a | Go]
              ↑
              sl.pointer = &arr[1]
              sl.len = 3
              sl.cap = 5  (index 1 থেকে শেষ পর্যন্ত)
```

**কেন Capacity = 5?**
```go
cap(sl) = len(arr) - start_index
cap(sl) = 6 - 1 = 5
```

#### বিস্তারিত উদাহরণ - Capacity বুঝি:

```go
package main
import "fmt"

func main() {
    arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    // বিভিন্ন slice এবং তাদের capacity
    s1 := arr[0:3]
    fmt.Printf("s1=%v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
    // s1=[0 1 2], len=3, cap=10 (0 থেকে 9 পর্যন্ত = 10)
    
    s2 := arr[5:8]
    fmt.Printf("s2=%v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
    // s2=[5 6 7], len=3, cap=5 (5 থেকে 9 পর্যন্ত = 5)
    
    s3 := arr[9:10]
    fmt.Printf("s3=%v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
    // s3=[9], len=1, cap=1 (9 থেকে 9 পর্যন্ত = 1)
}
```

#### Step 3: আরেকটি Sub-Slice তৈরি
```go
sl1 := sl[1:2]
```

**Memory:**
```
sl:  [is  | a  | Go | Intro | slice]
       0     1    2     3       4     (sl এর index)
       ↑     ↑
       start end

sl1 = [a]
sl1.pointer = &arr[2]  (original array-র index 2)
sl1.len = 1
sl1.cap = 4  (index 2 থেকে 5 পর্যন্ত = 4)
```

**Nested Slicing Example:**
```go
package main
import "fmt"

func main() {
    original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    level1 := original[2:8]  // [2,3,4,5,6,7]
    fmt.Printf("level1: %v, len=%d, cap=%d\n", level1, len(level1), cap(level1))
    // cap = 10-2 = 8
    
    level2 := level1[1:4]    // [3,4,5]
    fmt.Printf("level2: %v, len=%d, cap=%d\n", level2, len(level2), cap(level2))
    // cap = 8-1 = 7
    
    level3 := level2[2:]     // [5]
    fmt.Printf("level3: %v, len=%d, cap=%d\n", level3, len(level3), cap(level3))
    // cap = 7-2 = 5
    
    // সবগুলো একই underlying array শেয়ার করছে
    level3[0] = 999
    fmt.Println("original:", original)  // [0,1,2,3,4,999,6,7,8,9]
}
```

#### Step 4: Pointer এবং Address
```go
fmt.Println(*&sl1[0])  // "a"
```

**ব্যাখ্যা:**
- `sl1[0]` = value পাবে ("a")
- `&sl1[0]` = address পাবে (pointer)
- `*&sl1[0]` = address থেকে আবার value (dereference)

**Pointer Example:**
```go
package main
import "fmt"

func main() {
    s := []string{"Apple", "Banana", "Cherry"}
    
    // Direct access
    fmt.Println(s[0])  // Apple
    
    // Pointer access
    ptr := &s[0]
    fmt.Println(*ptr)  // Apple
    
    // Modify through pointer
    *ptr = "Apricot"
    fmt.Println(s)     // [Apricot Banana Cherry]
}
```

---

## কোড #২: Slice তৈরির বিভিন্ন উপায়

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 5} //slice literal

	fmt.Println("slice :", s, "length :", len(s), "capacity :", cap(s))

	//slice declare with make
	//make function with len

	lenSlice := make([]int, 5)

	lenSlice[4] = 10

	fmt.Println(lenSlice)

	//make function with len and capacity
	sliceData := make([]int, 3, 5) //here 3 is length and 5 is capacity and[] is slice literal

	sliceData[0] = 5
	sliceData[1] = 7
	sliceData[2] = 10
	fmt.Println(sliceData)

	//empty slice or nil slice
	var emptySlice []int

	fmt.Println(emptySlice)

	emptySlice = append(emptySlice, 1, 2, 3)

	fmt.Println(emptySlice, len(emptySlice), cap(emptySlice))
}
```

### 🔍 ৪টি পদ্ধতিতে Slice তৈরি:

---

#### পদ্ধতি #১: Slice Literal

```go
s := []int{1, 2, 5}
```

**বৈশিষ্ট্য:**
- সবচেয়ে সহজ এবং common
- Go automatically capacity = length করে দেয়
- Compile time-এ value জানা থাকলে এটা use করো

**বিস্তারিত উদাহরণ:**
```go
package main
import "fmt"

func main() {
    // বিভিন্ন type-এর slice literal
    
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Printf("numbers: %v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))
    
    names := []string{"Alice", "Bob", "Charlie"}
    fmt.Printf("names: %v, len=%d, cap=%d\n", names, len(names), cap(names))
    
    // Empty slice literal (nil নয়!)
    empty := []int{}
    fmt.Printf("empty: %v, len=%d, cap=%d, nil=%v\n", 
        empty, len(empty), cap(empty), empty == nil)
    // Output: empty: [], len=0, cap=0, nil=false
}
```

---

#### পদ্ধতি #২: `make()` শুধু Length দিয়ে

```go
lenSlice := make([]int, 5)
```

**কী হচ্ছে:**
```
make([]int, 5) creates:
┌─────────┬────────┬──────────┐
│ Pointer │ len=5  │  cap=5   │
└────┬────┴────────┴──────────┘
     │
     └─> [0 | 0 | 0 | 0 | 0]
          0   1   2   3   4
```

**বিস্তারিত উদাহরণ:**
```go
package main
import "fmt"

func main() {
    // make দিয়ে slice তৈরি
    s := make([]int, 5)
    
    fmt.Println("Initial:", s)  // [0 0 0 0 0]
    
    // সব index accessible
    s[0] = 10
    s[1] = 20
    s[2] = 30
    s[3] = 40
    s[4] = 50
    
    fmt.Println("After:", s)  // [10 20 30 40 50]
    
    // s[5] = 60  // ❌ Panic: index out of range
    
    // কিন্তু append করা যাবে
    s = append(s, 60)
    fmt.Println("Appended:", s)  // [10 20 30 40 50 60]
    fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))  // len=6, cap=10 (doubled!)
}
```

**কখন ব্যবহার করবে:**
```go
// যখন আগে থেকে জানো কতগুলো element লাগবে
userInput := make([]string, 10)  // 10 জন user-এর input নেবো

// Loop-এ direct assignment
for i := 0; i < 10; i++ {
    userInput[i] = fmt.Sprintf("User %d", i)
}
```

---

#### পদ্ধতি #৩: `make()` Length এবং Capacity দিয়ে

```go
sliceData := make([]int, 3, 5)
```

**এটা কেন দরকার?**

Performance optimization! যখন জানো যে slice বাড়বে, তখন আগে থেকে capacity দিয়ে দিলে বারবার memory allocation হবে না।

**Memory Visualization:**
```
make([]int, 3, 5) creates:
┌─────────┬────────┬──────────┐
│ Pointer │ len=3  │  cap=5   │
└────┬────┴────────┴──────────┘
     │
     └─> [0 | 0 | 0 | ? | ?]
          0   1   2   3   4
          ←── accessible ──→
          ←────── allocated ──────→
```

**বিস্তারিত উদাহরণ:**
```go
package main
import "fmt"

func main() {
    // Length < Capacity
    s := make([]int, 3, 5)
    
    fmt.Printf("Initial: %v, len=%d, cap=%d\n", s, len(s), cap(s))
    // [0 0 0], len=3, cap=5
    
    // প্রথম 3টা index accessible
    s[0] = 10
    s[1] = 20
    s[2] = 30
    
    // s[3] = 40  // ❌ Panic! Length এর বাইরে
    
    // কিন্তু append করতে পারবে (capacity-র মধ্যে)
    s = append(s, 40)
    fmt.Printf("After 1st append: %v, len=%d, cap=%d\n", s, len(s), cap(s))
    // [10 20 30 40], len=4, cap=5
    
    s = append(s, 50)
    fmt.Printf("After 2nd append: %v, len=%d, cap=%d\n", s, len(s), cap(s))
    // [10 20 30 40 50], len=5, cap=5
    
    // এখন capacity পূর্ণ, আরো append করলে নতুন array
    s = append(s, 60)
    fmt.Printf("After 3rd append: %v, len=%d, cap=%d\n", s, len(s), cap(s))
    // [10 20 30 40 50 60], len=6, cap=10 (capacity doubled!)
}
```

**Performance Comparison:**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    // Without pre-allocation
    start := time.Now()
    s1 := []int{}
    for i := 0; i < 1000000; i++ {
        s1 = append(s1, i)  // Multiple reallocations
    }
    fmt.Println("Without pre-alloc:", time.Since(start))
    
    // With pre-allocation
    start = time.Now()
    s2 := make([]int, 0, 1000000)
    for i := 0; i < 1000000; i++ {
        s2 = append(s2, i)  // No reallocation
    }
    fmt.Println("With pre-alloc:", time.Since(start))
    // প্রায় 2-3 গুণ দ্রুত!
}
```

---

#### পদ্ধতি #৪: Nil Slice

```go
var emptySlice []int
```

**Nil Slice vs Empty Slice:**

```go
package main
import "fmt"

func main() {
    // Nil slice
    var nilSlice []int
    fmt.Printf("nilSlice: %v, len=%d, cap=%d, nil=%v\n", 
        nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
    // Output: nilSlice: [], len=0, cap=0, nil=true
    
    // Empty slice (not nil)
    emptySlice := []int{}
    fmt.Printf("emptySlice: %v, len=%d, cap=%d, nil=%v\n", 
        emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
    // Output: emptySlice: [], len=0, cap=0, nil=false
    
    // Another empty slice
    madeSlice := make([]int, 0)
    fmt.Printf("madeSlice: %v, len=%d, cap=%d, nil=%v\n", 
        madeSlice, len(madeSlice), cap(madeSlice), madeSlice == nil)
    // Output: madeSlice: [], len=0, cap=0, nil=false
}
```

**Internal Difference:**
```
Nil Slice:
┌─────────┬────────┬──────────┐
│  nil    │ len=0  │  cap=0   │  → No memory allocated
└─────────┴────────┴──────────┘

Empty Slice:
┌─────────┬────────┬──────────┐
│ 0x...   │ len=0  │  cap=0   │  → Memory allocated (empty array)
└────┬────┴────────┴──────────┘
     └─> [] (empty array exists)
```

**Practical Example:**
```go
package main
import (
    "encoding/json"
    "fmt"
)

func main() {
    type Response struct {
        Data []int `json:"data"`
    }
    
    // Nil slice marshal করলে
    r1 := Response{Data: nil}
    json1, _ := json.Marshal(r1)
    fmt.Println(string(json1))  // {"data":null}
    
    // Empty slice marshal করলে
    r2 := Response{Data: []int{}}
    json2, _ := json.Marshal(r2)
    fmt.Println(string(json2))  // {"data":[]}
}
```

**Nil Slice-এ Append:**
```go
package main
import "fmt"

func main() {
    var s []int  // nil slice
    
    fmt.Println("Before:", s, "nil?", s == nil)  // [] nil? true
    
    // Nil slice-এও append করা যায়!
    s = append(s, 1)
    s = append(s, 2)
    s = append(s, 3)
    
    fmt.Println("After:", s, "nil?", s == nil)   // [1 2 3] nil? false
    fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))  // len=3, cap=4
}
```

---

## কোড #৩: Slice Copy এবং Append Behavior

```go
package main

import "fmt"

func main() {
	var data []int

	data = append(data, 1)
	data = append(data, 2)
	data = append(data, 3)

	data2 := data

	data = append(data, 4)
	data2 = append(data2, 5)

	data[0] = 100

	fmt.Println(data)   // [100 2 3 4]
	fmt.Println(data2)  // [1 2 3 5]
}
```

### 🔍 Step-by-Step Memory Analysis:

#### Step 1-3: প্রথম তিনটি Append
```go
var data []int       // nil slice
data = append(data, 1)
data = append(data, 2)
data = append(data, 3)
```

**Memory Evolution:**
```
After 1st append:
data → [1]
       len=1, cap=1

After 2nd append (capacity exceeded, new array):
data → [1 | 2]
       len=2, cap=2

After 3rd append (capacity exceeded, new array):
data → [1 | 2 | 3 | ?]
       len=3, cap=4 (Go doubles capacity)
```

#### Step 4: Slice Copy (Shallow Copy)
```go
data2 := data
```

**Critical Point:** এটা shallow copy! দুটো slice একই array point করছে।

```
Memory after copy:
data  → [1 | 2 | 3 | ?]
         ↑
data2 ─┘ (same underlying array)

Both slices:
  pointer = same
  len = 3
  cap = 4
```

**Detailed Example:**
```go
package main
import "fmt"

func main() {
    s1 := []int{10, 20, 30}
    s2 := s1  // Shallow copy
    
    fmt.Printf("s1: %p, %v\n", &s1[0], s1)
    fmt.Printf("s2: %p, %v\n", &s2[0], s2)
    // একই address!
    
    // Element modify করলে দুটোতেই change
    s2[1] = 999
    fmt.Println("s1:", s1)  // [10 999 30]
    fmt.Println("s2:", s2)  // [10 999 30]
}
```

#### Step 5: First Append (data)
```go
data = append(data, 4)
```

**যেহেতু capacity আছে (cap=4, len=3), নতুন array লাগবে না:**

```
data → [1 | 2 | 3 | 4]  (same array, length increased)
        ↑
data2 ─┘ (still pointing to same array)
```

**এখন পর্যন্ত দুটো একই array শেয়ার করছে!**

#### Step 6: Second Append (data2)
```go
data2 = append(data2, 5)
```

**Problem:** `data2` এর length=3, কিন্তু `data` ইতিমধ্যে index 3 use করছে!

**Solution:** Go নতুন array allocate করে `data2` এর জন্য।

```
Before:
data  → [1 | 2 | 3 | 4]
data2 → [1 | 2 | 3 | 4]  (same array)

After data2 = append(data2, 5):
data  → [1 | 2 | 3 | 4]     (old array)

data2 → [1 | 2 | 3 | 5]     (NEW array, copied from old)
```

**Detailed Append Behavior:**
```go
package main
import "fmt"

func main() {
    original := []int{1, 2, 3}
    fmt.Printf("original: %p, cap=%d\n", &original[0], cap(original))
    
    copy1 := original
    copy2 := original
    
    // copy1 append (within capacity)
    copy1 = append(copy1, 4)
    fmt.Printf("copy1 after append: %p, cap=%d\n", &copy1[0], cap(copy1))
    // Same address if capacity was enough
    
    // copy2 append (forces reallocation)
    for i := 0; i < 10; i++ {
        copy2 = append(copy2, i)
    }
    fmt.Printf("copy2 after many appends: %p, cap=%d\n", &copy2[0], cap(copy2))
    // Different address due to reallocation
}
```

#### Step 7: Modify First Element
```go
data[0] = 100
```

**যেহেতু এখন আলাদা array, শুধু `data` পরিবর্তিত হবে:**

```
data  → [100 | 2 | 3 | 4]

data2 → [1 | 2 | 3 | 5]  (unchanged)
```

**Final Output:**
```
data:  [100 2 3 4]
data2: [1 2 3 5]
```

---

### 📊 Capacity Growth Pattern

Go-তে slice-এর capacity কীভাবে বাড়ে:

```go
package main
import "fmt"

func main() {
    s := []int{}
    
    for i := 0; i < 20; i++ {
        prevCap := cap(s)
        s = append(s, i)
        if cap(s) != prevCap {
            fmt.Printf("After append %d: len=%d, cap=%d (grew from %d)\n", 
                i, len(s), cap(s), prevCap)
        }
    }
}
```

**Output:**
```
After append 0: len=1, cap=1 (grew from 0)
After append 1: len=2, cap=2 (grew from 1)
After append 2: len=3, cap=4 (grew from 2)
After append 4: len=5, cap=8 (grew from 4)
After append 8: len=9, cap=16 (grew from 8)
After append 16: len=17, cap=32 (grew from 16)
```

**Pattern:** Capacity প্রায় দ্বিগুণ হয় (2x growth strategy)

---

## কোড #৪: Function-এ Slice Pass করা

```go
package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changeSlice(a)

	fmt.Println(x)  // [1 2 3 4 10 6 7]
	fmt.Println(y)  // [10 6 7 11]
}

func changeSlice(sliceNum []int) []int {
	sliceNum[0] = 10
	return append(sliceNum, 11)
}
```

### 🔍 Function এবং Slice: Deep Dive

#### Initial Setup
```go
x := []int{1, 2, 3, 4, 5}
x = append(x, 6)
x = append(x, 7)
```

**Result:**
```
x = [1 | 2 | 3 | 4 | 5 | 6 | 7]
     0   1   2   3   4   5   6
```

#### Creating Sub-Slice
```go
a := x[4:]  // From index 4 to end
```

**Memory:**
```
x: [1 | 2 | 3 | 4 | 5 | 6 | 7]
                    ↑
                    a points here
a: [5 | 6 | 7]
    0   1   2  (a's indices)
```

**Critical:** `a` এবং `x` একই underlying array শেয়ার করছে!

#### Function Call: Pass by Value?

```go
y := changeSlice(a)
```

**Important Concept:** Go-তে slice **pass by value**, কিন্তু value হলো **pointer, length, capacity**!

```
When calling changeSlice(a):

Original 'a':                  Function parameter 'sliceNum':
┌─────────┬────────┬──────┐   ┌─────────┬────────┬──────┐
│ ptr     │ len=3  │ cap=3│   │ ptr     │ len=3  │ cap=3│
└────┬────┴────────┴──────┘   └────┬────┴────────┴──────┘
     │                              │
     └──────> Same Array! <─────────┘
              [5 | 6 | 7]
```

**Detailed Example:**
```go
package main
import "fmt"

func modifySlice(s []int) {
    fmt.Printf("Inside function - before: %p, %v\n", &s[0], s)
    s[0] = 999
    fmt.Printf("Inside function - after: %p, %v\n", &s[0], s)
}

func main() {
    original := []int{1, 2, 3}
    fmt.Printf("Before call: %p, %v\n", &original[0], original)
    
    modifySlice(original)
    
    fmt.Printf("After call: %p, %v\n", &original[0], original)
}
```

**Output:**
```
Before call: 0xc000014078, [1 2 3]
Inside function - before: 0xc000014078, [1 2 3]  (same address!)
Inside function - after: 0xc000014078, [999 2 3]
After call: 0xc000014078, [999 2 3]  (modified!)
```

#### Inside Function: Modify Element
```go
sliceNum[0] = 10
```

**Effect:**
```
x: [1 | 2 | 3 | 4 | 10 | 6 | 7]
                    ↑
                    Modified!
                    
a: [10 | 6 | 7]  (also sees the change)

sliceNum: [10 | 6 | 7]  (also sees the change)
```

#### Inside Function: Append
```go
return append(sliceNum, 11)
```

**Append creates a NEW slice:**

```
Old array (shared):
x: [1 | 2 | 3 | 4 | 10 | 6 | 7]

New array (for returned slice):
y: [10 | 6 | 7 | 11]
```

**Detailed Append in Function:**
```go
package main
import "fmt"

func appendInFunction(s []int) []int {
    fmt.Printf("Before append: %p, len=%d, cap=%d\n", &s[0], len(s), cap(s))
    s = append(s, 100)
    fmt.Printf("After append: %p, len=%d, cap=%d\n", &s[0], len(s), cap(s))
    return s
}

func main() {
    original := []int{1, 2, 3}
    fmt.Printf("Original: %p, len=%d, cap=%d\n", &original[0], len(original), cap(original))
    
    result := appendInFunction(original)
    
    fmt.Printf("Original after: %p, %v\n", &original[0], original)
    fmt.Printf("Result: %p, %v\n", &result[0], result)
}
```

**Key Learning:**
```
✅ Element modification → affects original (shared array)
❌ Append → doesn't affect original (new array created)
```

---

### 🎯 Common Pitfall Example

```go
package main
import "fmt"

func addElement(s []int, val int) {
    s = append(s, val)  // ❌ Wrong! Caller won't see this
}

func addElementCorrect(s []int, val int) []int {
    return append(s, val)  // ✅ Correct! Return new slice
}

func main() {
    nums := []int{1, 2, 3}
    
    // Wrong way
    addElement(nums, 4)
    fmt.Println(nums)  // [1 2 3] - unchanged!
    
    // Correct way
    nums = addElementCorrect(nums, 4)
    fmt.Println(nums)  // [1 2 3 4] - changed!
}
```

---

## কোড #৫: Variadic Function

```go
package main

import "fmt"

func printVar(nums ...int) {
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
}

func main() {
	printVar(5, 5, 6, 7, 8, 11, 15)
}
```

### 🔍 Variadic Functions বিস্তারিত

#### Variadic Parameter কী?

```go
func functionName(params ...Type) {
    // params হলো []Type (slice)
}
```

**`...` অপারেটরের দুই ব্যবহার:**

**১. Function declaration-এ:**
```go
func sum(numbers ...int) int {
    // numbers হলো []int slice
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}
```

**২. Function call-এ:**
```go
slice := []int{1, 2, 3}
sum(slice...)  // Slice কে expand করে individual arguments বানায়
```

#### Detailed Examples:

**Example 1: Basic Variadic**
```go
package main
import "fmt"

func printAll(items ...string) {
    fmt.Printf("Type: %T\n", items)  // Type: []string
    fmt.Printf("Value: %v\n", items)
    fmt.Printf("Length: %d\n", len(items))
    
    for i, item := range items {
        fmt.Printf("  [%d] = %s\n", i, item)
    }
}

func main() {
    // Different number of arguments
    printAll()  // Zero arguments
    printAll("Hello")  // One argument
    printAll("Go", "is", "awesome")  // Multiple arguments
}
```

**Example 2: Variadic with Regular Parameters**
```go
package main
import "fmt"

func greet(greeting string, names ...string) {
    for _, name := range names {
        fmt.Printf("%s, %s!\n", greeting, name)
    }
}

func main() {
    greet("Hello", "Alice", "Bob", "Charlie")
    // Output:
    // Hello, Alice!
    // Hello, Bob!
    // Hello, Charlie!
}
```

**❌ Error Example:**
```go
func wrong(items ...string, last string) {  // ❌ Compile error!
    // Variadic parameter must be LAST
}
```

**Example 3: Passing Existing Slice**
```go
package main
import "fmt"

func multiply(factor int, numbers ...int) []int {
    result := make([]int, len(numbers))
    for i, n := range numbers {
        result[i] = n * factor
    }
    return result
}

func main() {
    // Method 1: Individual arguments
    r1 := multiply(2, 1, 2, 3, 4, 5)
    fmt.Println(r1)  // [2 4 6 8 10]
    
    // Method 2: Slice expansion
    nums := []int{10, 20, 30}
    r2 := multiply(3, nums...)  // ... expands slice
    fmt.Println(r2)  // [30 60 90]
}
```

#### Original Code Analysis

```go
func printVar(nums ...int) {
	fmt.Println(nums)       // [5 5 6 7 8 11 15]
	fmt.Println(len(nums))  // 7
	fmt.Println(cap(nums))  // 7
}

func main() {
	printVar(5, 5, 6, 7, 8, 11, 15)
}
```

**What happens internally:**
```go
// When you call:
printVar(5, 5, 6, 7, 8, 11, 15)

// Go converts it to:
nums := []int{5, 5, 6, 7, 8, 11, 15}
printVar_internal(nums)
```

---

### 🎯 Real-World Variadic Examples

**Example: Custom Logger**
```go
package main
import (
    "fmt"
    "time"
)

func log(level string, messages ...string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("[%s] %s: ", timestamp, level)
    for i, msg := range messages {
        if i > 0 {
            fmt.Print(" | ")
        }
        fmt.Print(msg)
    }
    fmt.Println()
}

func main() {
    log("INFO", "Server started")
    log("ERROR", "Database connection failed", "Retrying...", "Attempt 1")
    log("DEBUG", "User login", "UserID: 123", "IP: 192.168.1.1")
}
```

**Output:**
```
[2024-01-15 10:30:45] INFO: Server started
[2024-01-15 10:30:45] ERROR: Database connection failed | Retrying... | Attempt 1
[2024-01-15 10:30:45] DEBUG: User login | UserID: 123 | IP: 192.168.1.1
```

**Example: fmt.Printf Style Function**
```go
package main
import "fmt"

func customPrintf(format string, args ...interface{}) {
    fmt.Printf(">>> ")
    fmt.Printf(format, args...)
    fmt.Println()
}

func main() {
    customPrintf("Hello, %s!", "World")
    customPrintf("Number: %d, Float: %.2f", 42, 3.14159)
    customPrintf("Mixed: %v, %v, %v", 10, "text", true)
}
```

---

## আরো উদাহরণ এবং Best Practices

### 📚 Advanced Slicing Techniques

#### 1. Full Slice Expression (3-index slice)

```go
package main
import "fmt"

func main() {
    original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    // Regular slice: [low:high]
    s1 := original[2:5]
    fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
    // s1: [2 3 4], len=3, cap=8 (can grow into original[5:])
    
    // Full slice: [low:high:max]
    s2 := original[2:5:5]
    fmt.Printf("s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
    // s2: [2 3 4], len=3, cap=3 (limited capacity!)
    
    // Benefit: Prevents unintended sharing
    s1 = append(s1, 100)  // Modifies original[5]
    s2 = append(s2, 200)  // Creates new array (capacity exceeded)
    
    fmt.Println("original:", original)  // [0 1 2 3 4 100 6 7 8 9]
}
```

**Syntax:**
```go
slice[low:high:max]
// length = high - low
// capacity = max - low
```

---

#### 2. Copy Function - Deep Copy

```go
package main
import "fmt"

func main() {
    source := []int{1, 2, 3, 4, 5}
    
    // Method 1: Using copy()
    dest1 := make([]int, len(source))
    n := copy(dest1, source)
    fmt.Printf("Copied %d elements: %v\n", n, dest1)
    
    // Method 2: Using append with nil
    dest2 := append([]int(nil), source...)
    fmt.Println("dest2:", dest2)
    
    // Method 3: Using append with empty slice
    dest3 := append([]int{}, source...)
    fmt.Println("dest3:", dest3)
    
    // Verify independence
    dest1[0] = 999
    fmt.Println("source:", source)  // [1 2 3 4 5] (unchanged)
    fmt.Println("dest1:", dest1)    // [999 2 3 4 5]
}
```

**Copy Behavior:**
```go
package main
import "fmt"

func main() {
    // Copy copies min(len(dst), len(src)) elements
    
    src := []int{1, 2, 3, 4, 5}
    
    // Destination smaller
    dst1 := make([]int, 3)
    n1 := copy(dst1, src)
    fmt.Printf("Copied %d: %v\n", n1, dst1)  // Copied 3: [1 2 3]
    
    // Destination larger
    dst2 := make([]int, 10)
    n2 := copy(dst2, src)
    fmt.Printf("Copied %d: %v\n", n2, dst2)  // Copied 5: [1 2 3 4 5 0 0 0 0 0]
    
    // Overlapping copy (same slice)
    nums := []int{1, 2, 3, 4, 5}
    copy(nums[1:], nums[:4])
    fmt.Println(nums)  // [1 1 2 3 4]
}
```

---

#### 3. Slice Tricks and Patterns

**Inserting Element:**
```go
package main
import "fmt"

func insert(slice []int, index, value int) []int {
    // Expand slice by 1
    slice = append(slice, 0)
    // Shift elements right
    copy(slice[index+1:], slice[index:])
    // Insert value
    slice[index] = value
    return slice
}

func main() {
    s := []int{1, 2, 4, 5}
    s = insert(s, 2, 3)
    fmt.Println(s)  // [1 2 3 4 5]
}
```

**Removing Element:**
```go
package main
import "fmt"

func remove(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

func removeUnordered(slice []int, index int) []int {
    // Faster but doesn't preserve order
    slice[index] = slice[len(slice)-1]
    return slice[:len(slice)-1]
}

func main() {
    s := []int{1, 2, 3, 4, 5}
    
    s = remove(s, 2)
    fmt.Println(s)  // [1 2 4 5]
    
    s = []int{1, 2, 3, 4, 5}
    s = removeUnordered(s, 2)
    fmt.Println(s)  // [1 2 5 4]
}
```

**Filtering:**
```go
package main
import "fmt"

func filter(slice []int, predicate func(int) bool) []int {
    result := slice[:0]  // Reuse same backing array
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Keep only even numbers
    evens := filter(nums, func(n int) bool {
        return n%2 == 0
    })
    
    fmt.Println(evens)  // [2 4 6 8 10]
}
```

**Reversing:**
```go
package main
import "fmt"

func reverse(slice []int) {
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
}

func main() {
    s := []int{1, 2, 3, 4, 5}
    reverse(s)
    fmt.Println(s)  // [5 4 3 2 1]
}
```

---

### ⚠️ Common Mistakes and Solutions

#### Mistake 1: Forgetting to Assign Append Result
```go
// ❌ Wrong
s := []int{1, 2, 3}
append(s, 4)  // Result is lost!
fmt.Println(s)  // [1 2 3]

// ✅ Correct
s := []int{1, 2, 3}
s = append(s, 4)
fmt.Println(s)  // [1 2 3 4]
```

#### Mistake 2: Slice Header Copy Confusion
```go
// ❌ Wrong assumption
func modify(s []int) {
    s = append(s, 100)  // Caller won't see this
}

s := []int{1, 2, 3}
modify(s)
fmt.Println(s)  // [1 2 3] - unchanged!

// ✅ Correct approach 1: Return slice
func modify(s []int) []int {
    return append(s, 100)
}

s = modify(s)

// ✅ Correct approach 2: Use pointer
func modify(s *[]int) {
    *s = append(*s, 100)
}

modify(&s)
```

#### Mistake 3: Range Loop Variable Reuse
```go
// ❌ Wrong
s := []int{1, 2, 3}
var pointers []*int

for _, v := range s {
    pointers = append(pointers, &v)  // v is reused!
}

for _, p := range pointers {
    fmt.Println(*p)  // 3 3 3 (all point to same variable!)
}

// ✅ Correct
for i := range s {
    pointers = append(pointers, &s[i])  // Point to slice elements
}

// Or
for _, v := range s {
    v := v  // Create new variable
    pointers = append(pointers, &v)
}
```

#### Mistake 4: Slice Header vs Underlying Array
```go
// ❌ Dangerous
func getData() []int {
    data := []int{1, 2, 3}
    return data[1:]  // Returns slice, but shares array
}

s1 := getData()
s2 := getData()
s1[0] = 100  // Might affect s2 if arrays are reused!

// ✅ Safe
func getData() []int {
    data := []int{1, 2, 3}
    result := make([]int, 2)
    copy(result, data[1:])
    return result
}
```

---

### 🎯 Performance Tips

#### 1. Pre-allocate When Size is Known
```go
// ❌ Slow
func generateNumbers(n int) []int {
    var result []int
    for i := 0; i < n; i++ {
        result = append(result, i)  // Multiple reallocations
    }
    return result
}

// ✅ Fast
func generateNumbers(n int) []int {
    result := make([]int, 0, n)  // Pre-allocate capacity
    for i := 0; i < n; i++ {
        result = append(result, i)  // No reallocation
    }
    return result
}

// ✅ Fastest
func generateNumbers(n int) []int {
    result := make([]int, n)  // Pre-allocate length
    for i := 0; i < n; i++ {
        result[i] = i  // Direct assignment
    }
    return result
}
```

#### 2. Reuse Backing Array for Filtering
```go
// ❌ Allocates new slice
func filterEven(nums []int) []int {
    var result []int
    for _, n := range nums {
        if n%2 == 0 {
            result = append(result, n)
        }
    }
    return result
}

// ✅ Reuses same backing array
func filterEven(nums []int) []int {
    result := nums[:0]  // Same backing array, length 0
    for _, n := range nums {
        if n%2 == 0 {
            result = append(result, n)
        }
    }
    return result
}
```

#### 3. Avoid Unnecessary Copies
```go
// ❌ Copies large slices
func process(data []LargeStruct) {
    for _, item := range data {  // item is a copy!
        // Process item
    }
}

// ✅ Uses pointers/indices
func process(data []LargeStruct) {
    for i := range data {  // No copy
        // Process data[i]
    }
}
```

---

### 📋 Quick Reference Card

```go
// ===== DECLARATION =====
var s []int              // nil slice
s := []int{}             // empty slice (not nil)
s := []int{1, 2, 3}      // slice literal
s := make([]int, 5)      // length 5, capacity 5
s := make([]int, 3, 5)   // length 3, capacity 5

// ===== OPERATIONS =====
len(s)                   // length
cap(s)                   // capacity
s[i]                     // access element
s[i] = x                 // modify element

// ===== SLICING =====
s[:]                     // full slice
s[2:]                    // from index 2 to end
s[:5]                    // from start to index 4
s[2:5]                   // from index 2 to 4
s[2:5:7]                 // low:high:max (limit capacity)

// ===== APPEND =====
s = append(s, 1)         // append single element
s = append(s, 1, 2, 3)   // append multiple elements
s = append(s, s2...)     // append another slice

// ===== COPY =====
dst := make([]int, len(src))
copy(dst, src)           // copy src to dst

// ===== ITERATION =====
for i, v := range s      // index and value
for i := range s         // index only
for _, v := range s      // value only

// ===== COMPARISON =====
s == nil                 // check if nil
len(s) == 0              // check if empty
```

---

### 🎓 উপসংহার

**Slice-এর মূল বিষয়গুলো:**

1. **Slice = Pointer + Length + Capacity**
2. **Reference Type** - copy করলে pointer copy হয়
3. **Dynamic Size** - `append()` দিয়ে বাড়ানো যায়
4. **Shares Underlying Array** - sub-slice একই array ব্যবহার করে
5. **Append May Reallocate** - capacity শেষ হলে নতুন array
6. **Function Parameter** - modification দেখা যায়, কিন্তু append না

**মনে রাখার মতো:**
- ✅ সবসময় `s = append(s, x)` format use করো
- ✅ Pre-allocate করো যখন size জানা থাকে
- ✅ Deep copy চাইলে `copy()` function use করো
- ⚠️ Sub-slice থেকে সাবধান - underlying array shared
- ⚠️ Function-এ append করলে return করতে হবে

এই guide follow করলে Go Slice-এ expert হয়ে যাবে! 💪🚀