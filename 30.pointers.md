# Golang Pointers - সম্পূর্ণ গাইড

---

## 📌 Pointer কি?

Pointer হলো একটি বিশেষ ধরনের variable যা অন্য একটি variable এর **memory address** সংরক্ষণ করে। সহজ ভাষায়, pointer মেমোরিতে কোন ডেটা কোথায় আছে তা নির্দেশ করে।

---

---

## 🔧 Pointer এর মূল উপাদান

### 1. `&` অপারেটর (Address Operator)
কোন variable এর memory address পেতে `&` অপারেটর ব্যবহার করা হয়।

```go
examMark := 100
addressOfExamMark := &examMark  // examMark এর address নিচ্ছি
```

### 2. `*` অপারেটর (Dereference Operator)
Pointer এর মাধ্যমে actual value access করতে `*` অপারেটর ব্যবহার করা হয়।

```go
valueOfTheAddress := *addressOfExamMark  // address থেকে value নিচ্ছি
```

---

## 💻 কোড বিশ্লেষণ

### উদাহরণ ১: Basic Pointer

```go
examMark := 100
addressOfExamMark := &examMark
valueOfTheAddress := *addressOfExamMark

fmt.Println("Address of examMark : ", addressOfExamMark)    // Output: 0xc0000b4008 (example)
fmt.Println("value of examMark : ", valueOfTheAddress)       // Output: 100
```

**ব্যাখ্যা:**
- `examMark` variable এ ১০০ সংরক্ষিত
- `&examMark` দিয়ে এর memory address পাচ্ছি
- `*addressOfExamMark` দিয়ে আবার original value পাচ্ছি

### উদাহরণ ২: Array Pointer

```go
func printF(data *[5]int) {
    fmt.Println(data)
}

nArr := [5]int{1, 2, 3, 4, 5}
printF(&nArr)  // array এর pointer পাঠাচ্ছি
```

### উদাহরণ ৩: Struct Pointer

```go
func developerInfo(test *Developer) {
    fmt.Println(test.Stack)
}

myDev := Developer{Name: "Tanvir", Stack: "Go"}
developerInfo(&myDev)  // struct এর pointer পাঠাচ্ছি
```

---

## 🔄 Pass by Value vs Pass by Reference

### Pass by Value (মান দ্বারা পাস)

Go-তে সাধারণভাবে সব কিছু **pass by value** হয়। মানে function এ parameter পাঠালে তার একটি **copy** তৈরি হয়।

```go
func changeValue(x int) {
    x = 200  // শুধু copy টি পরিবর্তন হবে
    fmt.Println("Inside function:", x)  // Output: 200
}

func main() {
    num := 100
    changeValue(num)
    fmt.Println("Outside function:", num)  // Output: 100 (unchanged)
}
```

**সমস্যা:** Original value পরিবর্তন হয় না কারণ function শুধু copy নিয়ে কাজ করছে।

### Pass by Reference (রেফারেন্স দ্বারা পাস)

Pointer ব্যবহার করে আমরা **reference** পাঠাতে পারি। এতে original value পরিবর্তন করা সম্ভব।

```go
func changeValueByReference(x *int) {
    *x = 200  // original value পরিবর্তন হবে
    fmt.Println("Inside function:", *x)  // Output: 200
}

func main() {
    num := 100
    changeValueByReference(&num)  // address পাঠাচ্ছি
    fmt.Println("Outside function:", num)  // Output: 200 (changed!)
}
```

**সুবিধা:** Original variable এ সরাসরি পরিবর্তন করা যায়।

## বিস্তারিত উদাহরণ

### Struct Modify করা

```go
// Pass by Value - Original struct পরিবর্তন হবে না
func updateSalaryByValue(dev Developer) {
    dev.Salary = 150000
    fmt.Println("Inside function:", dev.Salary)  // 150000
}

// Pass by Reference - Original struct পরিবর্তন হবে
func updateSalaryByReference(dev *Developer) {
    dev.Salary = 150000
    fmt.Println("Inside function:", dev.Salary)  // 150000
}

func main() {
    myDev := Developer{Name: "Tanvir", Salary: 100000}
    
    updateSalaryByValue(myDev)
    fmt.Println("After by value:", myDev.Salary)  // 100000 (unchanged)
    
    updateSalaryByReference(&myDev)
    fmt.Println("After by reference:", myDev.Salary)  // 150000 (changed!)
}
```

### Multiple Values Modify

```go
func swapNumbers(a *int, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    x := 10
    y := 20
    
    fmt.Println("Before swap: x =", x, "y =", y)  // x = 10, y = 20
    swapNumbers(&x, &y)
    fmt.Println("After swap: x =", x, "y =", y)   // x = 20, y = 10
}
```

### Slice Modify (বিশেষ কেস)

```go
func addSkill(skills []string) {
    skills = append(skills, "JavaScript")  // শুধু local slice পরিবর্তন হয়
}

func addSkillByReference(skills *[]string) {
    *skills = append(*skills, "JavaScript")  // original slice পরিবর্তন হয়
}

func main() {
    mySkills := []string{"Go", "Python"}
    
    addSkill(mySkills)
    fmt.Println("After by value:", mySkills)  // ["Go", "Python"]
    
    addSkillByReference(&mySkills)
    fmt.Println("After by reference:", mySkills)  // ["Go", "Python", "JavaScript"]
}
```

## কখন Pointer ব্যবহার করবেন?

### ✅ Pointer ব্যবহার করুন:
1. **বড় struct/data** পাঠানোর সময় (memory efficient)
2. Original value **modify** করতে চাইলে
3. Function থেকে **multiple values return** করতে চাইলে
4. **Performance** উন্নতির জন্য (copy না করে)

### ❌ Pointer ব্যবহার না করলেও চলে:
1. ছোট data type (int, bool, string) এর জন্য
2. Read-only operation এর জন্য
3. Immutable behavior চাইলে

## Pointer এর সুবিধা

1. **Memory Efficient:** বড় data structure এর copy না করে reference পাঠানো যায়
2. **Direct Modification:** Original value সরাসরি পরিবর্তন করা যায়
3. **Performance:** Copy operation এড়ানো যায়

## সতর্কতা

1. **Nil Pointer:** সবসময় check করুন pointer nil কিনা
```go
func safePrint(data *Developer) {
    if data == nil {
        fmt.Println("Pointer is nil!")
        return
    }
    fmt.Println(data.Name)
}
```

2. **Memory Leak:** Unnecessary pointer ব্যবহার এড়িয়ে চলুন

## সম্পূর্ণ উদাহরণ কোড

```go
package main

import "fmt"

type Developer struct {
    Name      string
    Age       int
    Stack     string
    Salary    int
    Country   string
    City      string
    Skills    []string
    Projects  []string
}

func printArray(data *[5]int) {
    fmt.Println("Array:", data)
}

func developerInfo(test *Developer) {
    fmt.Println("Developer Stack:", test.Stack)
}

func updateDeveloper(dev *Developer, newSalary int, newSkill string) {
    dev.Salary = newSalary
    dev.Skills = append(dev.Skills, newSkill)
}

func main() {
    // Basic pointer example
    examMark := 100
    addressOfExamMark := &examMark
    valueOfTheAddress := *addressOfExamMark

    fmt.Println("Address:", addressOfExamMark)
    fmt.Println("Value:", valueOfTheAddress)

    // Array pointer
    nArr := [5]int{1, 2, 3, 4, 5}
    printArray(&nArr)

    // Struct pointer
    myDev := Developer{
        Name:      "Tanvir",
        Age:       28,
        Stack:     "Go",
        Salary:    100000,
        Country:   "Bangladesh",
        City:      "Dhaka",
        Skills:    []string{"Go", "Python", "C++"},
        Projects:  []string{"Go-lang", "Python-lang"},
    }

    developerInfo(&myDev)
    
    fmt.Println("Before update:", myDev.Salary, myDev.Skills)
    updateDeveloper(&myDev, 150000, "Rust")
    fmt.Println("After update:", myDev.Salary, myDev.Skills)
}
```

## সারসংক্ষেপ

- **Pointer** হলো memory address এর ধারক
- **`&`** দিয়ে address পাওয়া যায়
- **`*`** দিয়ে value access করা যায়
- **Pass by Value:** Copy তৈরি হয়, original unchanged
- **Pass by Reference:** Pointer দিয়ে original modify করা যায়

Golang এ pointer mastering করলে আপনার code আরো efficient এবং powerful হবে! 🚀