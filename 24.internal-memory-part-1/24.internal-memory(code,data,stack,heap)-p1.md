# Go (Golang) Internal Memory – সহজ ও বিস্তারিত ব্যাখ্যা

Go প্রোগ্রাম যখন রান করে, তখন অপারেটিং সিস্টেম প্রোগ্রামের জন্য মেমোরি বরাদ্দ করে।  
এই মেমোরি ভেতরে ভেতরে কয়েকটি ভাগে বিভক্ত থাকে, যাকে আমরা **Go Internal Memory Layout** বলি।

---

## Go Internal Memory এর প্রধান অংশগুলো

Go মেমোরি মূলত ৪টি গুরুত্বপূর্ণ অংশে বিভক্ত:

1. Code Segment
2. Data Segment
3. Stack
4. Heap

---

## 1. Code Segment

### কী থাকে এখানে?
- Go কোড কম্পাইল হওয়ার পর যে **machine instructions** তৈরি হয়
- ফাংশনের লজিক
- প্রোগ্রাম ফ্লো

### বৈশিষ্ট্য
- Read-only (পরিবর্তন করা যায় না)
- প্রোগ্রাম চলার পুরো সময় থাকে
- খুবই নিরাপদ

### উদাহরণ
```go
func add(a, b int) int {
    return a + b
}
```

এই ফাংশনের লজিক Code Segment এ সংরক্ষিত থাকে।

---

## 2. Data Segment

### কী থাকে এখানে?
- Global variables
- Package-level variables
- Static data

### বৈশিষ্ট্য
- প্রোগ্রাম শুরু থেকে শেষ পর্যন্ত থাকে
- একবারই মেমোরি বরাদ্দ হয়

### উদাহরণ
```go
var appName = "MyApp"
var count = 10
```

`appName` এবং `count` → Data Segment এ থাকে।

---

## 3. Stack

### কী থাকে এখানে?
- Function এর local variables
- Function parameters
- Return address

### বৈশিষ্ট্য
- খুব দ্রুত কাজ করে
- LIFO (Last In First Out)
- Function শেষ হলে মেমোরি নিজে থেকেই মুক্ত হয়
- আকার সীমিত

### উদাহরণ
```go
func sum() {
    a := 5
    b := 10
    fmt.Println(a + b)
}
```

`a` এবং `b` → Stack এ থাকে।

### বাস্তব উদাহরণ
Stack কে ভাবতে পারো:  
**প্লেটের স্তূপ** – শেষ প্লেটটাই আগে ওঠে।

---

## 4. Heap

### কী থাকে এখানে?
- Dynamic memory
- Pointer দ্বারা রেফার করা ডেটা
- বড় struct বা object

### বৈশিষ্ট্য
- Stack এর তুলনায় ধীর
- আকার বড়
- Go এর **Garbage Collector** নিজে নিজে পরিষ্কার করে

### উদাহরণ
```go
func main() {
    x := new(int)
    *x = 100
}
```

`x` Heap এ সংরক্ষিত হয়।

---

## সংক্ষেপে তুলনা

| অংশ | কী কাজ |
|----|-------|
| Code Segment | কী করতে হবে |
| Data Segment | স্থায়ী ডেটা |
| Stack | অস্থায়ী ডেটা |
| Heap | ডাইনামিক ডেটা |

---

## গুরুত্বপূর্ণ নোট
- Go নিজে থেকেই Stack vs Heap ম্যানেজ করে
- Developer কে ম্যানুয়ালি free করতে হয় না
- Garbage Collector Heap পরিষ্কার করে

---

**Go Memory Management বোঝা মানে Go ভালোভাবে বোঝা।**