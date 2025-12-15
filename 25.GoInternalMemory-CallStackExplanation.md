# Go Internal Memory, Heap Escape & Call Stack (Diagram সহ ব্যাখ্যা)

এই ডকুমেন্টে তোমার দেওয়া Go কোড ব্যবহার করে নিচের ৩টি বিষয় **Markdown diagram সহ** ব্যাখ্যা করা হলো:

* ✅ Internal Memory Layout (Stack / Heap / Data Segment)
* ✅ Heap Escape Example
* ✅ Call Stack Animation (Step by Step)

---

## 🔹 Original Code (Reference)

```go
package main

import "fmt"

const age = 25

var p = 10
var r = 10

func callN(c, d int) {
	e := c + d
	fmt.Println(e)
}

func main() {
	add := func(a, b int) {
		fmt.Println(a + b)
	}

	add(age, r)
	callN(age, p)
}

func init() {
	fmt.Println("Welcome to Go Programming")
}
```

---

## 1️⃣ Go Internal Memory Layout (Diagram)

Go প্রোগ্রাম চলাকালীন মূলত ৩ ধরনের মেমরি ব্যবহার করে:

```
+---------------------------+
|        Heap Memory        |
|  (dynamic / escaped data)|
+---------------------------+

+---------------------------+
|        Stack Memory       |
|  main(), callN(), add()  |
|  local variables         |
+---------------------------+

+---------------------------+
|        Data Segment       |
|  global vars (p, r)      |
+---------------------------+

+---------------------------+
|     Compile Time Const    |
|        age = 25           |
+---------------------------+
```

### 🧠 Explanation

| Element  | Memory Location | Reason                       |
| -------- | --------------- | ---------------------------- |
| `age`    | Compile time    | constant, runtime memory নেই |
| `p, r`   | Data Segment    | global variable              |
| `main()` | Stack           | function call                |
| `add()`  | Stack           | anonymous func (no escape)   |
| `c,d,e`  | Stack           | local variables              |

---

## 2️⃣ Call Stack Animation (Step by Step)

### ▶️ Program Start

```
(empty)
```

---

### ▶️ init() Call

```
+-----------------+
| init()          |
+-----------------+
```

Output:

```
Welcome to Go Programming
```

init() শেষ → stack clear

---

### ▶️ main() Call

```
+-----------------+
| main()          |
+-----------------+
```

---

### ▶️ add(age, r) Call

```
+-----------------+
| add(a=25,b=10)  |
+-----------------+
| main()          |
+-----------------+
```

Output:

```
35
```

add() শেষ → pop

---

### ▶️ callN(age, p) Call

```
+-----------------+
| callN(c=25,d=10)|
| e=35            |
+-----------------+
| main()          |
+-----------------+
```

Output:

```
35
```

callN() শেষ → main() শেষ → program exit

---

## 3️⃣ Heap Escape Example (Extra Important 🔥)

### ❌ Normal Stack Allocation

```go
func sum() int {
	a := 10
	return a
}
```

📌 `a` শুধুই stack-এ থাকে

---

### ✅ Heap Escape Example

```go
func sumPtr() *int {
	a := 10
	return &a
}
```

### 🧠 কেন Heap এ গেল?

* `a` এর address return করা হয়েছে
* function শেষ হলেও value দরকার
* তাই Go compiler `a` কে **Heap-এ সরিয়ে দেয়**

```
Heap
 └── a = 10

Stack
 └── sumPtr() pointer
```

📌 একে বলে **Escape Analysis**

---

## 4️⃣ Anonymous Function + Heap Escape Example

```go
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
```

### Memory Diagram

```
Heap
 └── i = 0  ← shared

Stack
 └── returned function
```

📌 কারণ: inner function outer variable ব্যবহার করছে

---

## 5️⃣ Key Takeaways (Interview Ready ✅)

* ✔️ Constant → compile time
* ✔️ Global variable → data segment
* ✔️ Local variable → stack
* ✔️ Pointer return / closure → heap escape
* ✔️ Go নিজে থেকেই memory decide করে (GC friendly)

---

## ✅ Final Output of Given Code

```
Welcome to Go Programming
35
35
```

---


* 2 phases for the Go Program running:
    1. compilation phase
    2. execution phase

  --> go run main.go => compile--> create a binary file named main--> then automatically execute the  binary file

  --> go build main.go => compile--> create a binary file named main --> then we can execute the binary file by running ./main

  --> Go is a compiled language, so it compiles the code before executing it.
  -->In the compilation phase, all the constants and functions are allocated in the code segment(read-only). And all the global variables are allocated in the data segment.
  --> In the execution phase, the stack is used for function execution and local variables.
 
