# 🧠 Go Struct + Receiver Function 

---

## 📌 সম্পূর্ণ কোড

```go
package main

import "fmt"

type Person struct {
	FullName string
	Age      int
	Gender   string
}

// receiver function
func (person Person) printDetails() {
	var data string
	if person.Gender == "Male" {
		data = "His"
	} else {
		data = "Her"
	}

	fmt.Println("The name of the person is: ", person.FullName)
	fmt.Println(data, "Age", person.Age, " Years Old")
}

func main() {
	person1 := Person{
		FullName: "Tanvir Hossen Bappy",
		Age:      28,
		Gender:   "Male",
	}

	person2 := Person{
		FullName: "Ruhana Al Rumi",
		Age:      27,
		Gender:   "Female",
	}

	person1.printDetails()
	person2.printDetails()
}

func init() {
	fmt.Println("Welcome to struct and receiver func learning")
}
```

---

## 🔰 Struct কী? (সংক্ষেপে)

**Struct হলো Go-এর এমন একটি data type যেটা দিয়ে বাস্তব জীবনের কোনো object কে কোডে প্রকাশ করা হয়।**

এখানে আমরা `Person` নামের একটি Struct বানিয়েছি।

---

## 🧩 Person Struct বিশ্লেষণ

```go
type Person struct {
	FullName string
	Age      int
	Gender   string
}
```

### বাস্তব জীবনের সাথে মিল

একজন Person-এর যা যা থাকে:

* পূর্ণ নাম
* বয়স
* লিঙ্গ

👉 এগুলো একসাথে রাখার জন্যই Struct ব্যবহার করা হয়েছে।

---

## 🏗️ Struct থেকে Value তৈরি

```go
person1 := Person{
	FullName: "Tanvir Hossen Bappy",
	Age:      28,
	Gender:   "Male",
}
```

এখানে:

* `person1` হলো `Person` struct-এর একটি value
* Field নাম দিয়ে value বসানো হয়েছে

---

## 🔁 Receiver Function কী? (সবচেয়ে সহজ ভাষায়)

> **Receiver function হলো এমন একটি function যেটা কোনো Struct-এর সাথে যুক্ত থাকে।**

মানে:

* Function টা Struct-এর data ব্যবহার করতে পারে
* Struct-এর behaviour বোঝায়

👉 একে অন্য ভাষায় বলা যায়: **Struct-এর নিজের function**

---

## 🧠 Receiver Function Syntax

```go
func (person Person) printDetails() {
	// code
}
```

### এখানে কী হচ্ছে?

* `person` → receiver variable (Struct-এর একটি copy)
* `Person` → কোন Struct-এর সাথে function যুক্ত
* `printDetails` → function নাম

📌 মনে রাখো:

```
Struct + Function = Receiver Function
```

---

## 🧪 printDetails() Function ভেতরে কী হচ্ছে?

```go
if person.Gender == "Male" {
	data = "His"
} else {
	data = "Her"
}
```

👉 Gender অনুযায়ী লেখা পরিবর্তন করা হচ্ছে

```go
fmt.Println("The name of the person is: ", person.FullName)
fmt.Println(data, "Age", person.Age, " Years Old")
```

👉 Struct-এর field ব্যবহার করে সুন্দর output দেখানো হচ্ছে

---

## ▶️ Receiver Function কল করা

```go
person1.printDetails()
person2.printDetails()
```

🧠 মনে রাখার ট্রিক:

> **Receiver function সবসময় Struct value দিয়ে কল করা হয়**

Dot (`.`) মানে:

> "এই Struct-এর এই function চালাও"

---

## 🚀 init() Function কেন আছে?

```go
func init() {
	fmt.Println("Welcome to struct and receiver func learning")
}
```

### init() সম্পর্কে গুরুত্বপূর্ণ কথা

* `main()` এর আগে চলে
* Automatically call হয়
* Program শুরুতেই message দেখাতে কাজে লাগে

Execution flow:

```
Program start → init() → main()
```

---

## 🧠 Receiver Function কেন দরকার?

✔️ Struct-এর data সুন্দরভাবে ব্যবহার করতে
✔️ Code clean ও readable করতে
✔️ Real-world behaviour দেখাতে
✔️ Go-এর object-like design অনুসরণ করতে

---

## 💎 এক লাইনে Receiver Function

> **Receiver function হলো Struct-এর সাথে যুক্ত এমন function যেটা Struct-এর data ব্যবহার করে কাজ করে।**

---
 