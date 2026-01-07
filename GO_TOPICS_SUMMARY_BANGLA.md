# 🚀 Go Programming - সম্পূর্ণ টপিক সামারি (বাংলায়)

এই ডকুমেন্টে Go প্রোগ্রামিং এর সকল গুরুত্বপূর্ণ টপিক সংক্ষিপ্ত বুলেট পয়েন্টে দেওয়া হয়েছে - সহজে মনে রাখার জন্য।

---

## 📦 1. fmt Package - Print Methods

- **Print()** → নতুন লাইন যোগ করে না
- **Println()** → শেষে newline যোগ করে, স্বয়ংক্রিয় space দেয়
- **Printf()** → ফরম্যাট করা আউটপুট (`%s`, `%d`, `%f`, `%v`, `%T`)
- **Sprint()** → প্রিন্ট না করে string রিটার্ন করে
- **Sprintln()** → Sprint + newline
- **Sprintf()** → Printf এর মতো কিন্তু string রিটার্ন করে
- **Fprint()** → নির্দিষ্ট writer (file) এ লেখে
- **Fprintf()** → ফরম্যাট করে writer এ লেখে
- **Fprintln()** → Fprint + newline

**মনে রাখুন:**
- `Print` = Console এ
- `Sprint` = String return
- `Fprint` = File/Writer এ

---

## 🔢 2. Variables & Constants

- **Variable** → মান পরিবর্তন করা যায় (`var name string` বা `name := "value"`)
- **Constant** → মান পরিবর্তন করা যায় না (`const pi = 3.14`)
- **Short declaration** → `:=` (শুধু function এর ভিতরে)
- **Zero values** → int: 0, string: "", bool: false

---

## 🧭 3. Scope (স্কোপ)

### ৫ ধরনের Scope:

1. **Package Scope (Global)**
   - পুরো package এ accessible
   - Uppercase = Exported (অন্য package থেকে access)
   - Lowercase = Unexported (শুধু নিজের package)

2. **File Scope**
   - শুধু সেই file এর জন্য (import aliases)

3. **Block Scope (Local)**
   - Function scope → শুধু function এর ভিতরে
   - Loop scope → শুধু loop এর ভিতরে
   - If/Else scope → শুধু block এর ভিতরে
   - Switch scope → শুধু switch এর ভিতরে

4. **Lexical Scope**
   - Inner function outer function এর variable access করতে পারে
   - Closure এর ভিত্তি

5. **Shadowing**
   - Inner scope এ একই নামের variable outer কে hide করে
   - সতর্ক থাকুন - bug এর কারণ হতে পারে

**মনে রাখুন:**
- Global variable সব জায়গায় কিন্তু কম ব্যবহার করুন
- Local variable বেশি ব্যবহার করুন
- Shadowing থেকে সাবধান

---

## 📦 4. Package Scope

- **Package** → কোড organize করার উপায়
- **main package** → executable program
- **import** → অন্য package ব্যবহার করা
- **Exported** → Uppercase দিয়ে শুরু (public)
- **Unexported** → Lowercase দিয়ে শুরু (private)

---

## 🎭 5. Variable Shadowing

- Inner scope এ একই নামের variable তৈরি করলে outer variable hide হয়
- `:=` ব্যবহারে সাবধান - নতুন variable তৈরি হয়
- `=` ব্যবহার করলে existing variable update হয়

**উদাহরণ:**
```go
x := 10        // outer
{
    x := 20    // shadowing - নতুন variable
}
// x এখনো 10
```

---

## ⚙️ 6. init() Function

- `main()` এর আগে execute হয়
- Automatically call হয়
- একবারই চলে
- Setup/initialization এর জন্য ব্যবহার

**Execution flow:**
```
Program start → init() → main()
```

---

## 🔧 7. Functions

### Function Basics:
- **Declaration:** `func name(params) returnType { }`
- **Multiple returns:** `func name() (int, string) { }`
- **Named returns:** `func name() (result int) { }`
- **Variadic:** `func name(nums ...int) { }` (যত খুশি parameter)

### Function Types:

1. **Anonymous Function (Noob Function)**
   - নাম নেই
   - Variable এ রাখা যায়
   - Function expression

2. **First-order Function**
   - Normal function
   - Parameter বা return value হিসেবে function নেয় না

3. **Higher-order Function**
   - Parameter হিসেবে function নেয়
   - অথবা function return করে

**মনে রাখুন:**
- Function = First-class citizen (variable এর মতো)
- Callback এর জন্য higher-order function

---

## 🧠 8. Memory (Internal Memory)

### Memory Segments:

1. **Code Segment**
   - Program code থাকে
   - Constants থাকে
   - Read-only

2. **Data Segment (Global)**
   - Global variables
   - Static variables

3. **Stack**
   - Local variables
   - Function calls
   - LIFO (Last In First Out)
   - Fast কিন্তু limited size

4. **Heap**
   - Dynamic allocation
   - Pointers, slices, maps
   - Garbage collector manage করে
   - Slow কিন্তু flexible

**Call Stack:**
- Function call হলে stack frame তৈরি হয়
- Parameters, local variables, return address থাকে
- Function শেষ হলে stack frame remove হয়

---

## 🔒 9. Closure

- **Closure** = Function + Captured Variables
- Inner function outer function এর variable মনে রাখে
- Outer function শেষ হলেও variable থাকে
- Variable Heap এ যায় (escape analysis)

**মনে রাখুন:**
```
Closure = Function যা তার environment মনে রাখে
```

**উদাহরণ:**
```go
func outer() func() {
    money := 100  // Heap এ যাবে
    return func() {
        money += 10  // closure
    }
}
```

---

## 🏗️ 10. Struct

- **Struct** = User-defined data type
- একাধিক related data একসাথে রাখা
- Real-world object represent করে

**তৈরি করা:**
```go
type Person struct {
    Name string
    Age  int
}
```

**ব্যবহার:**
```go
p := Person{Name: "Habib", Age: 25}
fmt.Println(p.Name)  // Dot notation
```

**মনে রাখুন:**
- Struct = Container
- Field = Data
- Dot (.) দিয়ে access

---

## 🎯 11. Receiver Function

- **Receiver Function** = Struct এর সাথে যুক্ত function
- Struct এর behaviour define করে
- Method এর মতো কাজ করে

**Syntax:**
```go
func (p Person) printDetails() {
    fmt.Println(p.Name)
}
```

**Call করা:**
```go
person1.printDetails()  // Dot notation
```

**২ ধরনের Receiver:**
1. **Value Receiver** → Copy নিয়ে কাজ করে
2. **Pointer Receiver** → Original modify করে

---

## 📊 12. Array

- **Fixed size** → পরিবর্তন করা যায় না
- **Same type** → সব element একই type
- **Zero-indexed** → প্রথম element index 0

**Declaration:**
```go
var arr [5]int              // Zero values
arr := [5]int{1, 2, 3, 4, 5}
arr := [...]int{1, 2, 3}    // Auto size
```

**মনে রাখুন:**
- Array = Fixed size
- Value type = Copy হয়

---

## 👉 13. Pointers

### Pointer Basics:
- **Pointer** = Memory address ধারক
- **`&`** = Address operator (address নেয়)
- **`*`** = Dereference operator (value নেয়)

**উদাহরণ:**
```go
x := 10
ptr := &x      // x এর address
val := *ptr    // address থেকে value
```

### Pass by Value vs Reference:

1. **Pass by Value**
   - Copy তৈরি হয়
   - Original unchanged

2. **Pass by Reference (Pointer)**
   - Original modify করা যায়
   - Memory efficient

**কখন Pointer ব্যবহার করবেন:**
- বড় struct পাঠানোর সময়
- Original value modify করতে চাইলে
- Performance এর জন্য

**সতর্কতা:**
- Nil pointer check করুন
- Unnecessary pointer এড়িয়ে চলুন

---

## 🔪 14. Slice

### Slice = Dynamic Array

**৩টি মূল অংশ:**
1. **Pointer** → underlying array এর দিকে
2. **Length** → বর্তমানে কতগুলো element
3. **Capacity** → maximum কতগুলো element ধরতে পারবে

### তৈরি করার উপায়:

```go
// 1. Slice literal
s := []int{1, 2, 3}

// 2. make() - শুধু length
s := make([]int, 5)

// 3. make() - length + capacity
s := make([]int, 3, 5)

// 4. Array থেকে
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]  // [2, 3, 4]
```

### গুরুত্বপূর্ণ Operations:

```go
// Append
s = append(s, 4)
s = append(s, 5, 6, 7)

// Slicing
s[1:3]   // index 1 থেকে 2
s[:3]    // শুরু থেকে 2
s[2:]    // index 2 থেকে শেষ

// Copy
copy(dest, src)
```

### মনে রাখুন:
- **Slice vs Array:** Slice = Dynamic, Array = Fixed
- **Length vs Capacity:** len() = বর্তমান, cap() = সর্বোচ্চ
- **append() সবসময় assign করুন:** `s = append(s, x)`
- **Reference type:** Underlying array শেয়ার করে

**সাধারণ ভুল:**
- Index out of range (capacity আছে কিন্তু length নেই)
- append() এর return value ignore করা
- Slice sharing সম্পর্কে অসচেতন

---

## 🖥️ 15. Operating System Concepts

### Process (প্রসেস):
- Running program
- Isolated memory space
- Heavy weight
- PID আছে

### Thread (থ্রেড):
- Process এর মধ্যে execution unit
- Shared memory
- Lightweight
- Context switching দ্রুত

### Goroutine (গোরুটিন):
- Go এর ultra-lightweight thread
- 2KB stack (threads এর 1-2MB)
- Go runtime manage করে
- M:N scheduling (M goroutines, N OS threads)

### PCB (Process Control Block):
- Process এর সব info রাখে
- Process state, registers, memory info
- Context switching এ ব্যবহার হয়

### Context Switching:
- Process/thread পরিবর্তন করা
- State save + restore
- Overhead আছে (cache miss, TLB flush)

### Concurrency vs Parallelism:

**Concurrency:**
- একসাথে progress (dealing with)
- Single core এ possible
- Time-sharing, interleaved
- I/O bound tasks এর জন্য ভালো

**Parallelism:**
- একসাথে execution (doing)
- Multiple cores লাগবে
- Truly simultaneous
- CPU bound tasks এর জন্য ভালো

**মনে রাখুন:**
```
Concurrency = একজন শেফ তিনটা dish রান্না করছে
Parallelism = তিনজন শেফ একসাথে কাজ করছে
```

---

## 🔄 16. Defer

### Defer কী?
- Function call postpone করে
- Function শেষ হওয়ার ঠিক আগে execute হয়
- Resource cleanup এর জন্য perfect

### মূল নিয়ম:

1. **LIFO Order (Stack)**
   - শেষে defer করা প্রথমে execute হয়
   ```go
   defer fmt.Println("1")
   defer fmt.Println("2")
   defer fmt.Println("3")
   // Output: 3, 2, 1
   ```

2. **Arguments Immediate Evaluate**
   - Defer statement এ arguments তৎক্ষণাৎ evaluate হয়
   ```go
   x := 10
   defer fmt.Println(x)  // 10 store হয়ে যায়
   x = 20
   // Output: 10
   ```

3. **Execution Timing**
   - Return statement এর ঠিক আগে
   - Panic হলেও execute হয়

### Named Return vs Regular Return:

**Named Return:**
```go
func test() (result int) {
    defer func() { result++ }()
    return 5
}
// Returns: 6 (defer modify করতে পারে)
```

**Regular Return:**
```go
func test() int {
    x := 5
    defer func() { x++ }()
    return x
}
// Returns: 5 (defer modify করতে পারে না)
```

### কখন ব্যবহার করবেন:
- File close
- Database connection close
- Mutex unlock
- Panic recovery

### সতর্কতা:
- Loop এর মধ্যে defer এড়িয়ে চলুন
- Nil pointer check করুন
- Return value handle করুন

**মনে রাখুন:**
```
defer = "এই কাজটা function শেষে করো"
```

---

## 🔀 17. Variadic Function

- **Variadic** = যত খুশি parameter নেওয়া যায়
- **Syntax:** `func name(nums ...int)`
- Function এর ভিতরে slice হিসেবে কাজ করে

**উদাহরণ:**
```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3, 4, 5)  // যত খুশি argument
```

**Slice expand করা:**
```go
nums := []int{1, 2, 3}
sum(nums...)  // ... দিয়ে expand
```

---

## 🎲 18. Vogus Datatypes (Interface{})

- **Empty Interface** → `interface{}`
- যেকোনো type hold করতে পারে
- Type assertion দিয়ে actual type বের করা

**উদাহরণ:**
```go
var x interface{}
x = 10
x = "hello"
x = true  // সব possible
```

---

## 📝 সারাংশ - মনে রাখার টিপস

### 🔑 Core Concepts:
1. **Variables** → মান পরিবর্তনশীল
2. **Constants** → মান অপরিবর্তনীয়
3. **Scope** → কোথায় accessible
4. **Functions** → কোড reuse
5. **Closure** → Environment মনে রাখা

### 🏗️ Data Structures:
1. **Array** → Fixed size
2. **Slice** → Dynamic size
3. **Struct** → Custom type
4. **Pointer** → Memory address

### ⚙️ Advanced:
1. **Goroutines** → Concurrency
2. **Defer** → Cleanup
3. **Receiver** → Methods
4. **Variadic** → Flexible parameters

### 💡 Best Practices:
- Local variables বেশি ব্যবহার করুন
- Pointer সঠিক জায়গায় ব্যবহার করুন
- Defer দিয়ে resource cleanup করুন
- Slice এর append() সবসময় assign করুন
- Goroutines দিয়ে concurrency handle করুন

---

## 🎯 Interview এর জন্য মনে রাখুন:

1. **Slice vs Array** → Dynamic vs Fixed
2. **Value vs Pointer** → Copy vs Reference
3. **Concurrency vs Parallelism** → Dealing vs Doing
4. **Defer execution order** → LIFO
5. **Goroutine vs Thread** → Lightweight vs Heavy
6. **Closure** → Function + Environment
7. **Named return** → Defer modify করতে পারে

---

**শুভকামনা! Go তে expert হয়ে যান! 🚀**
