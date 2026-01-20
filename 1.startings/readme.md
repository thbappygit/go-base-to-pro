# 🧩 Go `fmt` Package - Print Methods Explained (বাংলায়)

এই ডকুমেন্টে আমরা Go-এর `fmt` প্যাকেজের সব ধরনের Print ফাংশন একদম সহজ বাংলায় বুঝব।

---

## 📦 Overview

`fmt` প্যাকেজ Go-এর **formatting powerhouse** —  
Laravel এর `echo`, `print_r`, `sprintf`, `var_dump` একসাথে যা করে, Go তে সেগুলো সব `fmt` দিয়েই করা যায়!

---

## 🟢 1. `fmt.Print()`

**কাজ:** প্রিন্ট করে কিন্তু শেষে **নতুন লাইন (\n)** যোগ করে না।

```go
fmt.Print("Hello")
fmt.Print("World")
```

**Output:**
```
HelloWorld
```

✅ **Tip:** নিজে হাতে space বা newline দিন:
```go
fmt.Print("Hello ", "World\n")
```

---

## 🟡 2. `fmt.Println()`

**কাজ:** প্রিন্ট করে এবং শেষে **newline (\n)** যোগ করে।

```go
fmt.Println("Hello")
fmt.Println("World")
```

**Output:**
```
Hello
World
```

✅ একাধিক ভ্যালুর মাঝে Go স্বয়ংক্রিয়ভাবে **space** দেয়:
```go
fmt.Println("Name:", "Habib", "Age:", 25)
```

**Output:**
```
Name: Habib Age: 25
```

---

## 🔵 3. `fmt.Printf()`

**কাজ:** ফরম্যাট করা আউটপুট প্রিন্ট করে।  
C ভাষার `printf()`-এর মতো।

```go
name := "Habib"
age := 25
fmt.Printf("My name is %s and I am %d years old.\n", name, age)
```

**Output:**
```
My name is Habib and I am 25 years old.
```

### 🧮 Common Format Specifiers

| Format | Description |
|--------|--------------|
| `%s` | string |
| `%d` | integer |
| `%f` | float |
| `%t` | boolean |
| `%v` | any value |
| `%#v` | Go syntax format |
| `%T` | data type |
| `%%` | percent sign |

---

## 🟣 4. `fmt.Sprint()`

**কাজ:** প্রিন্ট না করে **একটা string রিটার্ন করে।**

```go
msg := fmt.Sprint("Hello", " ", "World")
fmt.Print(msg)
```

**Output:**
```
Hello World
```

---

## 🔴 5. `fmt.Sprintln()`

**কাজ:** `Sprint` এর মতোই, তবে শেষে **newline (\n)** যোগ করে।

```go
msg := fmt.Sprintln("Hello", "World")
fmt.Print(msg)
```

**Output:**
```
Hello World
```

---

## 🟠 6. `fmt.Sprintf()`

**কাজ:** `Printf` এর মতোই, কিন্তু প্রিন্ট না করে **string রিটার্ন করে।**

```go
name := "Habib"
msg := fmt.Sprintf("Welcome, %s!", name)
fmt.Print(msg)
```

**Output:**
```
Welcome, Habib!
```

---

## ⚪ 7. `fmt.Fprint()`

**কাজ:** নির্দিষ্ট **writer** (যেমন file) এ আউটপুট পাঠায়।

```go
f, _ := os.Create("output.txt")
fmt.Fprint(f, "Hello File!")
```

**ফাইলের ভিতরে লেখা হবে:**
```
Hello File!
```

---

## ⚫ 8. `fmt.Fprintf()`

**কাজ:** `Printf` এর মতো format করে, কিন্তু **writer**-এ লেখে।

```go
f, _ := os.Create("log.txt")
name := "Habib"
fmt.Fprintf(f, "User: %s\n", name)
```

**ফাইলের ভিতরে লেখা হবে:**
```
User: Habib
```

---

## 🟤 9. `fmt.Fprintln()`

**কাজ:** `Fprint` এর মতোই, কিন্তু শেষে **newline** যোগ করে।

```go
f, _ := os.Create("out.txt")
fmt.Fprintln(f, "Hello", "World")
```

**ফাইলের ভিতরে লেখা হবে:**
```
Hello World
```

---

## 🧠 Summary Table

| Function | Prints to Console | Returns String | Adds Newline | Target |
|-----------|------------------|----------------|---------------|---------|
| `Print` | ✅ | ❌ | ❌ | Console |
| `Println` | ✅ | ❌ | ✅ | Console |
| `Printf` | ✅ | ❌ | ⚙️ Depends | Console |
| `Sprint` | ❌ | ✅ | ❌ | String |
| `Sprintln` | ❌ | ✅ | ✅ | String |
| `Sprintf` | ❌ | ✅ | ⚙️ Depends | String |
| `Fprint` | ✅ | ❌ | ❌ | Writer/File |
| `Fprintln` | ✅ | ❌ | ✅ | Writer/File |
| `Fprintf` | ✅ | ❌ | ⚙️ Depends | Writer/File |

---

## 💡 Full Example

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Print("1. Print")
    fmt.Println(" 2. Println")
    fmt.Printf("3. Printf -> %d\n", 3)

    s := fmt.Sprint("4. Sprint")
    s2 := fmt.Sprintln(" 5. Sprintln")
    s3 := fmt.Sprintf("6. Sprintf -> %s", "text")
    fmt.Print(s, s2, s3, "\n")

    f, _ := os.Create("test.txt")
    fmt.Fprint(f, "7. Fprint\n")
    fmt.Fprintln(f, "8. Fprintln")
    fmt.Fprintf(f, "9. Fprintf -> %d\n", 9)
}
```

---

## 🎯 Conclusion

`fmt` প্যাকেজ হলো Go ভাষার **printing, formatting, এবং string-building toolkit**।  
আপনি কখনো কনসোলে, কখনো ফাইলে, আবার কখনো string বানানোর জন্য — সবকিছু `fmt` দিয়েই করতে পারবেন।

