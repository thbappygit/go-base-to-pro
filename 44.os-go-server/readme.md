# Go HTTP Server – Request Journey (Kernel → Socket → Go Runtime)

এই README ফাইলটি ব্যাখ্যা করে **একটি HTTP request কীভাবে Client থেকে শুরু করে Go server-এর handler পর্যন্ত পৌঁছায়**, এবং এর মাঝে **NIC, Kernel, Socket, File Descriptor, Go Runtime, Goroutine, ServeMux**—সবকিছুর ভূমিকা কী।

---

## 📌 High-Level Architecture

```
Client
  ↓
Router / Internet
  ↓
Server NIC (Network Interface Card)
  ↓
RAM (Write Buffer)
  ↓
NIC Interrupt
  ↓
Kernel Network Stack
  ↓
Socket (Port 3000)
  ↓
Socket Receive Buffer
  ↓
File Descriptor (FD)
  ↓
Go Runtime (netpoller)
  ↓
Goroutine
  ↓
HTTP ServeMux
  ↓
Handler Function
```

---

## 🧠 Core Concepts Explained

### 1️⃣ NIC (Network Interface Card)

NIC হলো সার্ভারের নেটওয়ার্ক হার্ডওয়্যার।

* বাইরের দুনিয়া থেকে আসা সব binary packet প্রথমে NIC গ্রহণ করে
* NIC সরাসরি Go প্রোগ্রামের সাথে কথা বলতে পারে না
* ডাটা RAM-এর একটি buffer-এ লেখে (DMA ব্যবহার করে)

---

### 2️⃣ Write Buffer

Write Buffer হলো:

* RAM-এর একটি জায়গা যেখানে NIC প্রথমে incoming data লেখে
* এখনো কোনো process বা socket এই ডাটার মালিক নয়

সহজভাবে: **“ডাটা এসেছে, এখন Kernel ঠিক করবে কে পাবে”**

---

### 3️⃣ NIC Interrupt

ডাটা আসার পর NIC CPU-কে interrupt দেয়:

* CPU বর্তমান কাজ থামায়
* Kernel-এর network subsystem চালু হয়

---

### 4️⃣ Kernel Network Stack

Kernel তখন:

* Packet inspect করে (TCP/UDP)
* Destination IP ও Port দেখে
* Port 3000-এর জন্য কোনো listening socket আছে কিনা খোঁজে

---

### 5️⃣ Socket

Socket হলো network communication-এর endpoint।
একটি socket চিহ্নিত হয়:

```
(IP, Port, Protocol)
```

উদাহরণ:

```
0.0.0.0 : 3000 : TCP
```

---

### 6️⃣ Socket Receive Buffer

* Kernel-এর memory space-এ থাকে
* Socket-এর জন্য আসা ডাটা এখানে জমা হয়
* Go process সরাসরি NIC বা RAM touch করে না
* Go কেবল এই buffer থেকেই ডাটা পড়ে

---

### 7️⃣ File Descriptor (FD)

Linux-এ সবকিছুই file হিসেবে দেখা হয়:

* File
* Socket
* Pipe

File Descriptor হলো:

* একটি integer number (যেমন: 3, 4, 5)
* Kernel এই number দিয়ে process-কে socket access করতে দেয়

---

### 8️⃣ Ring Buffer

Ring Buffer বা Circular Buffer হলো:

* Fixed size buffer
* Read pointer এবং Write pointer থাকে
* High-performance network ও kernel design-এ ব্যবহৃত

ব্যবহার হয়:

* NIC buffer
* Socket buffer
* Event queue

---

## 🚀 Go Runtime & Kernel Interaction

Go প্রোগ্রাম চলে **user space**-এ।
Kernel থাকে **kernel space**-এ।

Go runtime kernel-এর সাথে syscall করে:

* `socket()`
* `bind()`
* `listen()`
* `accept()`
* `read()`
* `write()`

---

## 🔊 What happens in `ListenAndServe`

```go
http.ListenAndServe(":3000", mux)
```

ভেতরে যা হয়:

1. Kernel socket তৈরি হয়
2. Port 3000 bind হয়
3. Socket listen mode-এ যায়
4. Infinite accept loop শুরু হয়

```go
for {
    conn, err := l.Accept()
    go c.serve(conn)
}
```

---

### `Accept()` Behaviour

* কোনো request না থাকলে goroutine sleep করে
* Kernel নতুন connection দিলে Accept wake হয়
* প্রতিটা client-এর জন্য নতুন socket + FD তৈরি হয়

---

## 🧵 Goroutine কেন?

* প্রতিটা client-এর জন্য আলাদা goroutine
* হাজার হাজার concurrent request handle করা যায়
* Go scheduler OS thread efficiently ব্যবহার করে

---

## 🌐 HTTP Request Handling Flow in Go

1. Kernel socket buffer থেকে data আসে
2. Go netpoller (epoll/kqueue) detect করে
3. HTTP parser request বানায়
4. ServeMux path match করে
5. সংশ্লিষ্ট handler call হয়

---

## 🗺️ ServeMux Routing Example

```go
mux.HandleFunc("/hello", helloHandler)
mux.HandleFunc("/about", aboutHandler)
mux.HandleFunc("/details", detailHandler)
```

Path অনুযায়ী handler execute হয়:

* `/hello` → helloHandler
* `/about` → aboutHandler
* `/details` → detailHandler

---

## 📤 Response Path (Reverse Flow)

```
Handler
 → Go Runtime
 → write() syscall
 → Kernel Send Buffer
 → NIC
 → Router
 → Client
```

---

## ✅ Summary (One-Line Flow)

```
Client → NIC → RAM Buffer → Kernel → Socket → FD → Go Runtime → Goroutine → Mux → Handler
```

---

## 📚 Ideal For

* Go backend developers
* OS & Networking learners
* High-performance server architecture understanding

---

Happy Coding 🚀
