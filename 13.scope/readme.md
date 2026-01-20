# 📘 Go (Golang) — Scope সম্পূর্ণ গাইড (বাংলা)

Go তে **Scope** বলতে বোঝায় কোন ভেরিয়েবল, কনস্ট্যান্ট, টাইপ, ফাংশন বা প্যাকেজ আইডেন্টিফায়ার কোথা থেকে অ্যাক্সেসযোগ্য হবে।

Scope বুঝা গুরুত্বপূর্ণ কারণ:
- ভুল variable access থেকে রক্ষা করে
- মেমরি ব্যবস্থাপনা সহজ করে
- কোড রিডেবিলিটি বৃদ্ধি করে

---

# 🧭 Scope এর ধরন

Go তে প্রধানত ৫ ধরনের Scope রয়েছে:

1. **Package Scope (Global)**
2. **File Scope**
3. **Block Scope (Local)**
    - Function Scope
    - Loop Scope
    - If/Else Scope
    - Switch Scope
4. **Lexical Scope (Static Scope)**
5. **Shadowing**

---
 নিচে প্রতিটি Scope সম্পর্কে বিস্তারিত আলোচনা করা হলো:

---

## ১. Package Scope (Global Scope)

Package Scope হলো যখন কোনো variable বা function একটি package এর যেকোনো জায়গা থেকে access করা যায়। Package level এ declare করা variable গুলো সকল files এ accessible হয়।

**বৈশিষ্ট্য:**
- Package এর সব file থেকে access করা যায়
- Uppercase দিয়ে শুরু হলে অন্য package থেকেও access করা যায় (Exported)
- Lowercase দিয়ে শুরু হলে শুধু নিজের package এ থাকে (Unexported)
```go
package main

import "fmt"

// Package scope - সব জায়গা থেকে access করা যাবে
var globalVar = "আমি global variable"
const GlobalConst = "আমি exported constant" // Uppercase - অন্য package থেকে access করা যাবে

func main() {
    fmt.Println(globalVar)
    printGlobal()
}

func printGlobal() {
    fmt.Println(globalVar) // এখানেও access করা যাচ্ছে
}
```

---

## ২. File Scope

File Scope শুধুমাত্র সেই specific file এর জন্য। এটি মূলত `import` statements এবং `package` declaration এর ক্ষেত্রে প্রযোজ্য।
```go
package main

import (
    "fmt"
    f "fmt" // alias - শুধু এই file এ কাজ করবে
)

func main() {
    fmt.Println("Regular import")
    f.Println("Alias import - শুধু এই file এ কাজ করবে")
}
```

---

## ৩. Block Scope (Local Scope)

Block Scope হলো curly braces `{}` এর মধ্যে declare করা variable গুলো। এগুলো শুধুমাত্র ঐ block এর মধ্যেই accessible।

### ৩.১ Function Scope

Function এর মধ্যে declare করা variable শুধু ঐ function এর মধ্যে accessible।
```go
package main

import "fmt"

func main() {
    message := "আমি main function এর ভিতরে"
    fmt.Println(message)
    
    anotherFunction()
    // fmt.Println(innerMessage) // Error: innerMessage এখানে accessible না
}

func anotherFunction() {
    innerMessage := "আমি অন্য function এর ভিতরে"
    fmt.Println(innerMessage)
    // fmt.Println(message) // Error: message এখানে accessible না
}
```

### ৩.২ Loop Scope

Loop এর মধ্যে declare করা variable শুধু loop এর মধ্যেই accessible।
```go
package main

import "fmt"

func main() {
    // for loop scope
    for i := 0; i < 3; i++ {
        loopVar := "আমি loop এর ভিতরে"
        fmt.Printf("i = %d, %s\n", i, loopVar)
    }
    // fmt.Println(i) // Error: i এখানে accessible না
    // fmt.Println(loopVar) // Error: loopVar এখানে accessible না
    
    // range loop scope
    numbers := []int{1, 2, 3}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
    // fmt.Println(index) // Error: index এখানে accessible না
}
```

### ৩.৩ If/Else Scope

If/Else block এর মধ্যে declare করা variable শুধু ঐ block এর মধ্যেই accessible।
```go
package main

import "fmt"

func main() {
    x := 10
    
    // If statement এ variable declare করা
    if y := 20; x < y {
        z := 30 // if block এর ভিতরে
        fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)
    } else {
        w := 40 // else block এর ভিতরে
        fmt.Printf("x=%d, y=%d, w=%d\n", x, y, w)
        // fmt.Println(z) // Error: z এখানে accessible না
    }
    
    // fmt.Println(y) // Error: y এখানে accessible না
    // fmt.Println(z) // Error: z এখানে accessible না
}
```

### ৩.৪ Switch Scope

Switch block এর মধ্যে declare করা variable শুধু ঐ block এর মধ্যেই accessible।
```go
package main

import "fmt"

func main() {
    // Switch statement এ variable declare করা
    switch day := "শনিবার"; day {
    case "শনিবার":
        weekend := true
        fmt.Printf("%s - সপ্তাহান্ত: %v\n", day, weekend)
    case "রবিবার":
        weekend := true
        fmt.Printf("%s - সপ্তাহান্ত: %v\n", day, weekend)
    default:
        weekend := false
        fmt.Printf("%s - কর্মদিবস: %v\n", day, weekend)
    }
    
    // fmt.Println(day) // Error: day এখানে accessible না
    // fmt.Println(weekend) // Error: weekend এখানে accessible না
}
```

---

## ৪. Lexical Scope (Static Scope)

Lexical Scope মানে হলো inner function তার outer function এর variable access করতে পারে। এটি closure এর ভিত্তি।
```go
package main

import "fmt"

func outerFunction() func() {
    outerVar := "আমি outer function এর variable"
    counter := 0
    
    // Inner function (closure) outer variable access করতে পারে
    innerFunction := func() {
        counter++ // outer variable modify করছি
        innerVar := "আমি inner function এর variable"
        fmt.Printf("%s\n", outerVar)     // Outer variable access
        fmt.Printf("%s\n", innerVar)      // Inner variable access
        fmt.Printf("Counter: %d\n", counter)
    }
    
    return innerFunction
}

func main() {
    myFunc := outerFunction()
    myFunc() // Counter: 1
    myFunc() // Counter: 2
    myFunc() // Counter: 3
    
    // Lexical scope এর আরেকটি উদাহরণ
    x := 10
    
    func() {
        y := 20
        fmt.Printf("Outer x: %d\n", x) // Outer variable access
        
        func() {
            z := 30
            fmt.Printf("x: %d, y: %d, z: %d\n", x, y, z) // সব level এর variable access
        }()
    }()
}
```

---

## ৫. Shadowing (Variable Shadowing)

Shadowing হলো যখন inner scope এ একই নামের variable declare করা হয় যা outer scope এর variable কে "hide" করে দেয়।
```go
package main

import "fmt"

var x = "global x"

func main() {
    fmt.Println(x) // Output: global x
    
    x := "main function এর x" // Shadowing global x
    fmt.Println(x) // Output: main function এর x
    
    {
        x := "block এর x" // Shadowing main function এর x
        fmt.Println(x) // Output: block এর x
        
        {
            x := "inner block এর x" // আরো একবার shadowing
            fmt.Println(x) // Output: inner block এর x
        }
        
        fmt.Println(x) // Output: block এর x (inner block এর বাইরে)
    }
    
    fmt.Println(x) // Output: main function এর x
}

// Shadowing এর আরেকটি উদাহরণ
func shadowExample() {
    name := "বাইরের নাম"
    
    if true {
        name := "ভিতরের নাম" // Shadowing
        fmt.Println("If block:", name) // Output: ভিতরের নাম
    }
    
    fmt.Println("Function:", name) // Output: বাইরের নাম
}
```

### Shadowing এর সতর্কতা
```go
package main

import "fmt"

func main() {
    count := 10
    fmt.Println("প্রথমে count:", count) // 10
    
    if true {
        count := 20 // নতুন variable (shadowing) - সাবধান!
        fmt.Println("If এ count:", count) // 20
    }
    
    fmt.Println("পরে count:", count) // 10 (পরিবর্তিত হয়নি!)
    
    // সঠিক উপায় - shadowing এড়ানো
    value := 10
    fmt.Println("প্রথমে value:", value) // 10
    
    if true {
        value = 20 // একই variable update হচ্ছে
        fmt.Println("If এ value:", value) // 20
    }
    
    fmt.Println("পরে value:", value) // 20 (পরিবর্তিত হয়েছে)
}
```

---

## সারাংশ

| Scope Type | Accessibility | Example |
|-----------|---------------|---------|
| **Package Scope** | পুরো package | `var globalVar = 10` |
| **File Scope** | শুধু ঐ file | `import` aliases |
| **Block Scope** | শুধু ঐ block | `if`, `for`, `switch`, function |
| **Lexical Scope** | Outer থেকে inner | Closures |
| **Shadowing** | Inner scope outer কে hide করে | একই নামের variable |

---

## মনে রাখার টিপস

1. **Global variable** সব জায়গায় accessible কিন্তু অতিরিক্ত ব্যবহার করা উচিত নয়
2. **Local variable** যতটা সম্ভব ব্যবহার করুন - code আরো clean হয়
3. **Shadowing** এর ব্যাপারে সতর্ক থাকুন - bug এর কারণ হতে পারে
4. **Lexical scope** closure তৈরিতে অত্যন্ত useful
5. **Block scope** ব্যবহার করে memory efficient code লেখা যায়