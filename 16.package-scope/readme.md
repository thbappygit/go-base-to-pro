# 📦 Go (Golang) — Package Scope সম্পূর্ণ গাইড (বাংলা)

Go তে **Package Scope** বলতে বোঝায় একটি প্যাকেজের মধ্যে ভেরিয়েবল, ফাংশন, কনস্ট্যান্ট এবং টাইপ এর অ্যাক্সেসিবিলিটি। এটি Go এর একটি গুরুত্বপূর্ণ ধারণা যা কোড organization এবং encapsulation নিশ্চিত করে।

---

## 📚 Package Scope কি?

Package Scope হলো এমন একটি mechanism যেখানে:
- একই প্যাকেজের সকল ফাইল থেকে variable/function access করা যায়
- **Uppercase** দিয়ে শুরু হওয়া identifiers অন্য package থেকেও access করা যায় (**Exported**)
- **Lowercase** দিয়ে শুরু হওয়া identifiers শুধু নিজের package এ থাকে (**Unexported**)

---

## 🔤 Exported vs Unexported

### Exported (বড় হাতের অক্ষর দিয়ে শুরু)
```go
var GlobalVariable = "আমি exported"        // অন্য package থেকে access করা যাবে
const GlobalConstant = "আমিও exported"     // অন্য package থেকে access করা যাবে
func PublicFunction() {}                    // অন্য package থেকে call করা যাবে
type PublicStruct struct {}                 // অন্য package থেকে use করা যাবে
```

### Unexported (ছোট হাতের অক্ষর দিয়ে শুরু)
```go
var privateVariable = "আমি unexported"     // শুধুমাত্র এই package এ access করা যাবে
const privateConstant = "আমিও unexported"  // শুধুমাত্র এই package এ access করা যাবে
func privateFunction() {}                   // শুধুমাত্র এই package এ call করা যাবে
type privateStruct struct {}                // শুধুমাত্র এই package এ use করা যাবে
```

---

## 📁 আমাদের প্রজেক্ট স্ট্রাকচার

```
16.package-scope/
├── go.mod                    # মডিউল ডিক্লারেশন ফাইল
├── main.go                   # main প্যাকেজ (প্রধান প্রোগ্রাম)
└── mathlib/
    └── summation.go          # mathlib প্যাকেজ (সেকেন্ডারি প্যাকেজ)
```

---

## 🔧 Module Setup - `go mod init mypack`

### ধাপ ১: Module Initialize করা

আপনার প্রজেক্ট ডিরেক্টরিতে terminal খুলুন এবং এই কমান্ড চালান:

```bash
go mod init mypack
```

**আউটপুট:**
```
go: creating new go.mod: module mypack
```

এই কমান্ড একটি `go.mod` ফাইল তৈরি করে যা আপনার মডিউলের তথ্য রাখে।

### ধাপ ২: go.mod ফাইলের বিষয়বস্তু

```mod
module mypack

go 1.25.4
```

**ব্যাখ্যা:**
- `module mypack` - আপনার মডিউলের নাম
- `go 1.25.4` - আপনি যে Go ভার্সন ব্যবহার করছেন

---

## 💻 আমাদের কোড

### ফাইল ১: `main.go` (main package)

```go
// filepath: main.go
package main

import (
	"mypack/mathlib"
)

func main() {
	num1 := 200
	num2 := 300
	mathlib.Summation(num1, num2)
}
```

**ব্যাখ্যা:**
- `package main` - এটি main package যেখানে প্রোগ্রাম execution শুরু হয়
- `import ("mypack/mathlib")` - অন্য package import করছি (module path)
- `mathlib.Summation()` - exported function call করছি (বড় অক্ষর দিয়ে শুরু)

### ফাইল ২: `mathlib/summation.go` (mathlib package)

```go
// filepath: mathlib/summation.go
package mathlib

import "fmt"

func Summation(n1 int, n2 int) {
	res := n1 + n2
	fmt.Println(res)
}
```

**ব্যাখ্যা:**
- `package mathlib` - এটি mathlib package
- `func Summation()` - exported function (বড় অক্ষর দিয়ে শুরু) তাই অন্য package থেকে access করা যায়
- `fmt.Println(res)` - ফলাফল প্রিন্ট করছি

---

## 🚀 প্রোগ্রাম চালানো

### কমান্ড ১: Program রান করা

```bash
go run main.go
```

**আউটপুট:**
```
500
```

**কিভাবে কাজ করে:**
1. `main.go` execute হয়
2. `main()` function call হয়
3. `num1 = 200`, `num2 = 300`
4. `mathlib.Summation(200, 300)` call হয়
5. `summation.go` এর `Summation` function execute হয়
6. `res = 200 + 300 = 500`
7. `500` print হয়

### কমান্ড ২: সম্পূর্ণ প্রজেক্ট রান করা

```bash
go run ./...
```

এই কমান্ড সমস্ত `.go` ফাইল compile এবং run করে।

---

## 🏗️ Executable তৈরি করা (Build করা)

### কমান্ড ১: Build করা

```bash
go build
```
 

---

## 📊 Package Scope এর বৈশিষ্ট্য

| বৈশিষ্ট্য | উদাহরণ | ফলাফল |
|----------|---------|---------|
| **Package Level Access** | same package এর সব file থেকে | ✅ Access করা যায় |
| **Exported (Capital)** | `func Summation()` | ✅ অন্য package থেকেও access |
| **Unexported (Small)** | `func calculate()` | ❌ শুধু নিজের package এ |
| **Import Path** | `mypack/mathlib` | ✅ মডিউল path ব্যবহার করে |
| **Module Name** | `module mypack` | ✅ go.mod এ ডিফাইন করা |

---

## 🔍 Package Scope নিয়ম

### নিয়ম ১: একই Package এর সব File Access করতে পারে

```go
// file1.go
package mypackage
var x = 10

// file2.go
package mypackage
func printX() {
    fmt.Println(x) // ✅ Access করা যায়
}
```

### নিয়ম ২: Exported হতে হলে Capital Letter দিয়ে শুরু

```go
// mypackage/math.go
package mypackage

var privatVar = 10        // ❌ অন্য package এ access না
func privateFunc() {}     // ❌ অন্য package এ access না

var PublicVar = 20        // ✅ অন্য package এ access করা যায়
func PublicFunc() {}       // ✅ অন্য package এ access করা যায়
```

### নিয়ম ৩: Import Path সঠিক হতে হবে

```go
// ✅ সঠিক
import "mypack/mathlib"
mathlib.Summation(10, 20)

// ❌ ভুল (mypack define করা হয়নি)
import "mathlib"
mathlib.Summation(10, 20)
```

---

## 🎯 Practical উদাহরণ

### উদাহরণ ১: Multiple Functions একই Package এ

```go
// mathlib/summation.go
package mathlib

import "fmt"

// Exported Function - অন্য package থেকে call করা যায়
func Summation(n1 int, n2 int) {
	fmt.Println(calculate(n1, n2))
}

// Unexported Function - শুধু mathlib package এ use হয়
func calculate(a int, b int) int {
	return a + b
}
```

### উদাহরণ ২: একাধিক Package File

```
mathlib/
├── summation.go     # Summation function
├── subtraction.go   # Subtraction function
└── helper.go        # Helper functions
```

সব ফাইল `package mathlib` দিয়ে শুরু হয় এবং একে অপরের private function access করতে পারে।

---

## ⚠️ সাধারণ ভুল এবং সমাধান

### ভুল ১: Unexported Function Import করা

```go
// ❌ ভুল - summation.go তে privateFunc আছে
import "mypack/mathlib"
mathlib.privateFunc() // Compile Error!

// ✅ সঠিক - শুধু Exported Function use করুন
mathlib.Summation(10, 20)
```

### ভুল ২: Wrong Import Path

```go
// ❌ ভুল
import "mathlib"
mathlib.Summation(10, 20) // mypack prefix ছাড়া

// ✅ সঠিক
import "mypack/mathlib"
mathlib.Summation(10, 20) // mypack prefix সহ
```

### ভুল ৩: go.mod ছাড়া Project

```bash
# ❌ ভুল - go.mod নেই
go run main.go   // Import error!

# ✅ সঠিক - প্রথমে initialize করুন
go mod init mypack
go run main.go
```

---

## 📝 চেকলিস্ট

আপনার প্রজেক্ট setup এর জন্য:

- [x] `go mod init mypack` চালিয়ে `go.mod` তৈরি করা
- [x] Package structure তৈরি করা (main, mathlib)
- [x] Exported functions Capital letter দিয়ে শুরু করা
- [x] Import path সঠিক করা (`mypack/mathlib`)
- [x] `go run main.go` দিয়ে test করা
- [x] `go build` দিয়ে executable তৈরি করা

---

## 📖 সংক্ষিপ্ত সারাংশ

**Package Scope** Go এ code organization এর জন্য গুরুত্বপূর্ণ:

1. **Module এর প্রয়োজনীয়তা**: `go mod init mypack` দিয়ে শুরু করুন
2. **Exported vs Unexported**: Capital letter = exported, small letter = unexported
3. **Import Path**: `mypack/mathlib` এর মতো সঠিক path ব্যবহার করুন
4. **Same Package Access**: একই package এর সব file একে অপরের সব identifier access করতে পারে
5. **Encapsulation**: Unexported functions দিয়ে internal implementation hide করুন

---

## 🔗 দরকারি কমান্ড

| কমান্ড | কাজ |
|--------|------|
| `go mod init <name>` | নতুন module initialize |
| `go run main.go` | প্রোগ্রাম চালান |
| `go build` | Executable তৈরি করুন |
| `go build -o <name>` | Custom নামে build করুন |
| `go mod tidy` | Unused dependencies সরান |
| `go get <package>` | External package ডাউনলোড করুন |

---

✅ **এখন আপনি Package Scope সম্পূর্ণভাবে বুঝেছেন এবং আপনার প্রজেক্ট সঠিকভাবে setup করতে পারবেন!**

