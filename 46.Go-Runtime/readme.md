# Go Runtime সম্পূর্ণ গাইড  

## সূচিপত্র
1. [Go Runtime কী?](#go-runtime-কী)
2. [M:P:G Model](#mpg-model)
3. [Go Scheduler](#go-scheduler)
4. [Queue System](#queue-system)
5. [Epoll System](#epoll-system)
6. [Network I/O](#network-io)
7. [Memory Management](#memory-management)
8. [Interview Questions](#interview-questions)

---

## Go Runtime কী?

**সহজ উত্তর:**
Go Runtime হলো একটা built-in program যেটা আপনার Go code execute করে। এটা automatically goroutines manage করে, memory allocate করে এবং garbage collection করে।

**Key Features:**
- Lightweight concurrency (Goroutines)
- Automatic memory management (GC)
- Efficient scheduling
- Built-in network poller

---

## M:P:G Model

### Component গুলো:

#### 1. **M (Machine/OS Thread)**
```
- Operating System এর actual thread
- CPU তে actual কাজ করে
- Expensive to create
- Limited সংখ্যায় থাকে
```

#### 2. **P (Logical Processor)**
```
- Context/Resource holder
- Default: GOMAXPROCS = CPU cores
- প্রতিটা P এর নিজস্ব Local Run Queue আছে
- M কে resources provide করে
```

#### 3. **G (Goroutine)**
```
- Lightweight thread (2KB stack)
- হাজার হাজার goroutine একসাথে run করতে পারে
- User-space এ manage হয়
- Cheap to create and destroy
```

### Model Structure:
```
M:P:G = 4:4:16
- 4টা OS Threads (M)
- 4টা Logical Processors (P)
- 16টা Goroutines (G)

┌─────┐     ┌─────┐     ┌─────┐     ┌─────┐
│ M₁  │────│ P₁  │     │ P₂  │────│ M₂  │
└─────┘     └─────┘     └─────┘     └─────┘
             │   │       │   │
             G₁  G₂      G₃  G₄
```

---

## Go Scheduler

### Initialization Process:

```go
1. initialize go scheduler
2. go runtime → sys call → kernel → epoll_create
3. set up GC (Garbage Collector)
```

### Scheduling States:

**Goroutine States:**
```
1. Running   - actively executing
2. Runnable  - waiting in queue
3. Waiting   - blocked (I/O, channel, lock)
4. Dead      - execution completed
```

### Work Stealing Algorithm:

```
যদি P এর Local Queue খালি:
1. Global Queue থেকে G নেওয়ার চেষ্টা করে
2. অন্য P এর Local Queue থেকে অর্ধেক G চুরি করে
3. Network poller check করে
4. কোনো কাজ না পেলে M sleep হয়
```

**Example:**
```go
// P1 এর queue: [G1, G2, G3, G4]
// P2 এর queue: []

// Work stealing এর পর:
// P1 এর queue: [G1, G2]
// P2 এর queue: [G3, G4]  ← চুরি করেছে
```

---

## Queue System

### 1. Local Run Queue (LRQ)

```
Properties:
- প্রতিটা P এর নিজস্ব queue
- Maximum 256 goroutines ধরতে পারে
- Lock-free access (fast)
- LIFO order (Last In First Out)

Example:
P1 → [G1, G2, G3, G4] ← Local Queue
```

### 2. Global Run Queue (GRQ)

```
Properties:
- সব P share করে
- Unlimited capacity
- Mutex protected (slow)
- FIFO order (First In First Out)

কখন ব্যবহার হয়:
- Local queue full হলে
- New goroutine create করার সময়
- System calls থেকে return করার পর
```

### Queue Priority:
```
P Local Queue (highest priority)
    ↓ (if empty)
Global Queue (medium priority)
    ↓ (if empty)
Network Poller (lowest priority)
```

---

## Epoll System

### কী এবং কেন?

**Traditional Blocking I/O Problem:**
```go
// এভাবে করলে entire thread block হয়
data := conn.Read()  // Thread wait করে
```

**Epoll Solution:**
```go
// Non-blocking I/O
// Thread কাজ করতে পারে অন্য কিছুতে
```

### Epoll Components:

#### 1. **epoll_create()**
```
Purpose: Epoll instance তৈরি করে
Returns: File descriptor

Example:
epfd = epoll_create(1)
```

#### 2. **epoll_ctl()**
```
Purpose: File descriptors manage করে

Operations:
- EPOLL_CTL_ADD: নতুন fd add করে
- EPOLL_CTL_MOD: existing fd modify করে
- EPOLL_CTL_DEL: fd remove করে

Example:
epoll_ctl(epfd, EPOLL_CTL_ADD, socket_fd, &event)
```

#### 3. **epoll_wait()**
```
Purpose: I/O events এর জন্য wait করে

Behavior:
- Ready events থাকলে immediately return করে
- না থাকলে block করে
- Timeout set করা যায়

Example:
n = epoll_wait(epfd, events, maxevents, timeout)
```

### Epoll Flow Diagram:
```
Goroutine → Network Call
    ↓
Register fd with epoll_ctl
    ↓
Park goroutine (not blocking thread!)
    ↓
M continues with other goroutines
    ↓
epoll_wait signals when ready
    ↓
Scheduler wakes up parked goroutine
    ↓
Goroutine continues execution
```

---

## Network I/O

### HTTP Server Example Analysis:

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, "Hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "About page")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)
    mux.HandleFunc("/about", aboutHandler)
    
    fmt.Println("Server running on :8080")
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        fmt.Println("Error starting server", err)
    }
}
```

### Execution Flow:

```
1. main() goroutine starts
   └─> Main thread এ run হয়

2. http.ListenAndServe() calls
   └─> Socket create করে
   └─> bind() and listen() করে
   └─> Loop এ accept() করতে থাকে

3. প্রতিটা connection এর জন্য:
   ├─> l.Accept() → blocks
   ├─> File descriptor (fd) পায়
   ├─> fd কে netpoll এ register করে (epoll_ctl)
   ├─> নতুন goroutine spawn করে: go c.serve(connCtx)
   └─> এই goroutine connection handle করে

4. Goroutine parked হয় যখন:
   ├─> conn.Read() waiting
   ├─> conn.Write() waiting
   └─> epoll_wait signal দেয় যখন ready

5. Request complete হলে:
   └─> Goroutine terminates
   └─> Connection close হয়
```

### Network Polling Internals:

```
Socket Operations:
┌──────────────────────────────────────┐
│  Application Layer                   │
│  go func() { conn.Read() }          │
└──────────────┬───────────────────────┘
               ↓
┌──────────────────────────────────────┐
│  Go Runtime (netpoll)                │
│  - epoll_ctl(ADD, fd, events)       │
│  - Park goroutine                    │
│  - epoll_wait() in background       │
└──────────────┬───────────────────────┘
               ↓
┌──────────────────────────────────────┐
│  Kernel (epoll)                      │
│  - Monitor file descriptors          │
│  - Signal when data ready            │
└──────────────────────────────────────┘
```

---

## Memory Management

### Memory Layout:

```
High Address
┌──────────────────────────────────────┐
│          Stack (grows down)          │  ← Function calls, local vars
│              ↓                       │
│                                      │
│          Free Space                  │
│                                      │
│              ↑                       │
│          Heap (grows up)             │  ← Dynamic allocation
├──────────────────────────────────────┤
│          Data Segment                │  ← Global variables
├──────────────────────────────────────┤
│          Code Segment                │  ← Program instructions
└──────────────────────────────────────┘
Low Address
```

### Goroutine Stack:

```
Initial: 2KB (very small!)
Maximum: 1GB (can grow)

Growing Process:
1. Stack overflow detect করে
2. নতুন বড় stack allocate করে
3. Data copy করে
4. পুরনো stack free করে
```

**Example:**
```go
func recursive(n int) {
    if n == 0 {
        return
    }
    // Stack grows as recursion deepens
    recursive(n - 1)
}

func main() {
    recursive(10000) // Stack automatically grows
}
```

### Heap Allocation:

```go
// Stack allocation (fast)
func foo() {
    x := 10  // Stack এ allocate
}

// Heap allocation (slower but flexible)
func bar() *int {
    x := 10
    return &x  // Escapes to heap
}

func main() {
    s := make([]int, 1000)  // Heap এ allocate
    m := make(map[string]int)  // Heap এ allocate
}
```

---

## Interview Questions

### Basic Level:

**Q1: Goroutine এবং Thread এর মধ্যে পার্থক্য কী?**

```
Answer:
┌─────────────────┬──────────────────┬──────────────────┐
│   Feature       │   Thread (OS)    │   Goroutine      │
├─────────────────┼──────────────────┼──────────────────┤
│ Stack Size      │ 1-2 MB           │ 2 KB (grows)     │
│ Creation Time   │ Slow (~1-2 ms)   │ Fast (~few µs)   │
│ Switching Cost  │ High             │ Very Low         │
│ Managed By      │ OS Kernel        │ Go Runtime       │
│ Communication   │ Shared Memory    │ Channels         │
│ Scalability     │ 1000s            │ Millions         │
└─────────────────┴──────────────────┴──────────────────┘

Example:
// Thread এ 10,000 create করা heavy
// Goroutine এ 1,000,000 create করা easy!

for i := 0; i < 1000000; i++ {
    go func() {
        // Do something
    }()
}
```

**Q2: GOMAXPROCS কী এবং এটা কীভাবে কাজ করে?**

```
Answer:
GOMAXPROCS = একসাথে কতগুলো OS thread goroutines run করতে পারবে

Default: runtime.NumCPU() (CPU cores সংখ্যা)

Example:
runtime.GOMAXPROCS(4)  // Max 4টা thread use করবে

Effect:
- বেশি হলে: Better parallelism, বেশি context switching
- কম হলে: Less overhead, কম parallelism

Best Practice: Default value ই best (CPU cores)
```

**Q3: Goroutine কখন block হয়?**

```
Answer:
Blocking Operations:
1. Channel operations (send/receive when not ready)
2. Network I/O (read/write)
3. System calls (file operations)
4. Mutex locks (sync.Mutex)
5. Sleep (time.Sleep)

Non-blocking (Scheduler switches):
1. CPU-intensive loops (cooperative)
2. Function calls
3. Memory allocations

Example:
ch := make(chan int)

go func() {
    v := <-ch  // Blocks until data available
    fmt.Println(v)
}()

ch <- 10  // Unblocks the goroutine
```

---

### Intermediate Level:

**Q4: Work Stealing Algorithm ব্যাখ্যা করুন।**

```
Answer:
Work Stealing = একটা P যখন idle থাকে, তখন অন্য busy P থেকে কাজ চুরি করে

Algorithm Steps:
1. Check Local Queue (নিজের)
2. Check Global Queue (সবার shared)
3. Check Network Poller
4. Steal from other P (অর্ধেক নেয়)
5. If nothing found, M goes to sleep

Example Scenario:
P1: [G1, G2, G3, G4, G5, G6]  ← Busy
P2: []                         ← Idle

After Work Stealing:
P1: [G1, G2, G3]
P2: [G4, G5, G6]  ← Stole half from P1

Benefits:
- Load balancing
- Better CPU utilization
- Prevents thread starvation
```

**Q5: Netpoller কীভাবে কাজ করে বিস্তারিত বলুন।**

```
Answer:
Netpoller = Go এর internal mechanism যেটা network I/O efficiently handle করে

Architecture:
┌────────────────────────────────────┐
│   Goroutine                        │
│   conn.Read() / conn.Write()       │
└──────────────┬─────────────────────┘
               ↓
┌────────────────────────────────────┐
│   Runtime (netpoll.go)             │
│   - pollDesc structure             │
│   - Park goroutine                 │
│   - Register with epoll            │
└──────────────┬─────────────────────┘
               ↓
┌────────────────────────────────────┐
│   OS (epoll/kqueue/iocp)          │
│   - Monitor file descriptors       │
│   - Non-blocking I/O               │
└────────────────────────────────────┘

Flow:
1. conn.Read() called
2. Check if data immediately available
   - Yes: Return immediately
   - No: Continue below
3. Create pollDesc for this fd
4. epoll_ctl(EPOLL_CTL_ADD, fd)
5. gopark() - Park the goroutine
6. M continues with other goroutines
7. Background thread runs epoll_wait()
8. When data ready: epoll_wait returns
9. goready() - Wake up goroutine
10. Goroutine continues execution

Code Example:
func (fd *netFD) Read(p []byte) (n int, err error) {
    // Try immediate read
    n, err = syscall.Read(fd.Sysfd, p)
    if err == syscall.EAGAIN {
        // Would block - use netpoller
        if err := fd.pd.waitRead(); err != nil {
            return 0, err
        }
        n, err = syscall.Read(fd.Sysfd, p)
    }
    return n, err
}
```

**Q6: Goroutine leak কী এবং কীভাবে prevent করবেন?**

```
Answer:
Goroutine Leak = যখন goroutine কখনো terminate হয় না এবং memory তে থেকে যায়

Common Causes:

1. Channel without receiver:
❌ Bad:
ch := make(chan int)
go func() {
    ch <- 10  // Forever blocked! Leak!
}()

✅ Good:
ch := make(chan int, 1)  // Buffered
go func() {
    ch <- 10
}()

2. Waiting on channel forever:
❌ Bad:
go func() {
    <-ch  // If nothing sent, forever waiting
}()

✅ Good:
go func() {
    select {
    case <-ch:
        // Process
    case <-time.After(5 * time.Second):
        return  // Timeout
    }
}()

3. Context not used:
❌ Bad:
go func() {
    for {
        // Infinite loop, no way to stop
        doWork()
    }
}()

✅ Good:
go func(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return  // Graceful exit
        default:
            doWork()
        }
    }
}(ctx)

Detection:
- Use pprof: go tool pprof http://localhost:6060/debug/pprof/goroutine
- Monitor: runtime.NumGoroutine()
```

---

### Advanced Level:

**Q7: M:P:G model এ P কেন দরকার? শুধু M আর G থাকলে চলে না?**

```
Answer:
P (Logical Processor) থাকার কারণ:

Without P (M:G only):
M1 ← G1, G2, G3
M2 ← G4, G5
Problem:
- M কে directly schedule করা expensive
- M create/destroy করা slow
- OS scheduling overhead বেশি

With P (M:P:G):
M1 ← P1 ← G1, G2, G3
M2 ← P2 ← G4, G5

Benefits:

1. Decoupling:
   - M and G decoupled হয়
   - P context হিসেবে কাজ করে
   - M can be reused

2. Local Scheduling:
   - P তে Local Run Queue থাকে
   - Lock-free scheduling
   - Better cache locality

3. Work Stealing:
   - P থেকে P তে কাজ move করা easy
   - M এর সাথে directly না করে P দিয়ে করা fast

4. System Calls:
   - Blocking syscall এ M detach হয়
   - P অন্য M এর সাথে attach হতে পারে
   - Goroutines বন্ধ হয় না

Example:
G1 running on M1:P1
G1 does blocking syscall (e.g., file read)
→ M1 blocks with G1
→ P1 detaches from M1
→ P1 attaches to M2 (or creates new M)
→ Other goroutines (G2, G3) continue on M2:P1
```

**Q8: Scheduler preemption কীভাবে কাজ করে?**

```
Answer:
Preemption = Scheduler forcefully একটা goroutine থামিয়ে অন্যটা run করানো

Types:

1. Cooperative Preemption (Old):
- Function calls এ check
- Goroutine নিজে yield করে
- Problem: Tight loops preempt হয় না

❌ Problem Code:
func compute() {
    for i := 0; i < 1000000000; i++ {
        // No function calls
        // No preemption point!
    }
}

2. Asynchronous Preemption (Go 1.14+):
- Signal-based (SIGURG)
- Forcefully preempt করতে পারে
- Tight loops ও preempt হয়

How it works:
┌─────────────────────────────────┐
│  Sysmon thread (monitors)       │
│  - Runs every 10ms              │
│  - Checks goroutine runtime     │
└──────────────┬──────────────────┘
               ↓
        Running > 10ms?
               ↓ Yes
┌──────────────────────────────────┐
│  Send preemption signal          │
│  SIGURG → goroutine              │
└──────────────┬───────────────────┘
               ↓
┌──────────────────────────────────┐
│  Signal handler                  │
│  - Save state                    │
│  - Switch to another G           │
└──────────────────────────────────┘

Preemption Points:
1. Function calls
2. Channel operations
3. Memory allocation
4. System calls
5. Timer signals (async)

✅ Now Works:
func compute() {
    for i := 0; i < 1000000000; i++ {
        // Will be preempted after 10ms
    }
}
```

**Q9: Sysmon thread এর role কী?**

```
Answer:
Sysmon = System Monitor, একটা special thread যেটা background এ monitoring করে

Characteristics:
- OS thread (M without P)
- Never blocks
- Runs independently
- Interval: starts 20µs, max 10ms

Responsibilities:

1. Network Poller:
   - epoll_wait() periodically check করে
   - Ready connections এর goroutines wake করে

2. Preemption:
   - Long-running goroutines detect করে (>10ms)
   - Async preemption signal পাঠায়

3. Garbage Collection:
   - Force GC trigger করে (2 minutes idle)
   - scavenge heap memory

4. Deadlock Detection:
   - সব goroutines blocked detect করে
   - Panic raise করে

5. Timer Management:
   - Expired timers handle করে
   - time.Sleep() এর goroutines wake করে

6. Retake P from Syscall:
   - Syscall blocked P detect করে (>10ms)
   - P detach করে অন্য M দিয়ে দেয়

Code Flow:
func sysmon() {
    for {
        // Check and handle each responsibility
        
        // 1. Netpoll
        netpoll(false)
        
        // 2. Retake P from long syscall
        retake(now)
        
        // 3. Force GC if needed
        if shouldTriggerGC() {
            forcegc()
        }
        
        // Sleep with exponential backoff
        usleep(delay)
    }
}
```

**Q10: Race condition এবং Go এ কীভাবে handle করবেন?**

```
Answer:
Race Condition = Multiple goroutines একই memory একসাথে access করে

Example of Race:
❌ Bad Code:
var counter int

func increment() {
    counter++  // Not atomic! Race condition!
}

func main() {
    for i := 0; i < 1000; i++ {
        go increment()
    }
    time.Sleep(time.Second)
    fmt.Println(counter)  // Random value!
}

Solutions:

1. Mutex (sync.Mutex):
✅ Solution 1:
var (
    counter int
    mu      sync.Mutex
)

func increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}

2. Atomic Operations:
✅ Solution 2:
var counter int64

func increment() {
    atomic.AddInt64(&counter, 1)
}

3. Channels:
✅ Solution 3:
func main() {
    ch := make(chan int)
    counter := 0
    
    // Single goroutine owns counter
    go func() {
        for range ch {
            counter++
        }
    }()
    
    for i := 0; i < 1000; i++ {
        ch <- 1
    }
}

4. sync.RWMutex (Read-heavy):
✅ Solution 4:
var (
    cache = make(map[string]string)
    mu    sync.RWMutex
)

func read(key string) string {
    mu.RLock()         // Multiple readers OK
    defer mu.RUnlock()
    return cache[key]
}

func write(key, value string) {
    mu.Lock()          // Only one writer
    defer mu.Unlock()
    cache[key] = value
}

Detection:
go run -race main.go
// Shows race conditions during execution

Best Practices:
1. Share memory by communicating (channels)
2. Don't communicate by sharing memory
3. Use sync package primitives when needed
4. Always test with -race flag
```

---

## Quick Reference

### Key Concepts Summary:

```
┌────────────────────────────────────────────────────────┐
│ Component    │ Description                             │
├──────────────┼─────────────────────────────────────────┤
│ M            │ OS Thread (Machine)                     │
│ P            │ Logical Processor (Context)             │
│ G            │ Goroutine (Lightweight thread)          │
│ LRQ          │ Local Run Queue (per P)                 │
│ GRQ          │ Global Run Queue (shared)               │
│ Netpoller    │ I/O multiplexing (epoll/kqueue)        │
│ Sysmon       │ Background monitoring thread            │
│ Preemption   │ Forceful goroutine switching           │
│ Work Stealing│ Load balancing between Ps              │
└──────────────┴─────────────────────────────────────────┘
```

### Common Interview Topics:

```
✅ Must Know:
- Goroutine vs Thread difference
- M:P:G model basics
- Channel operations
- Mutex usage
- Race conditions

✅ Should Know:
- Scheduler algorithm
- Work stealing
- Netpoller basics
- Context usage
- WaitGroup

✅ Advanced:
- Preemption mechanism
- Sysmon responsibilities
- Memory model
- GC interaction
- Performance tuning
```

### Code Patterns:

```go
// 1. Basic Goroutine
go func() {
    // Do work
}()

// 2. Goroutine with WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Do work
}()
wg.Wait()

// 3. Context with Timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

go func() {
    select {
    case <-ctx.Done():
        return
    case <-doWork():
        // Handle result
    }
}()

// 4. Channel Pattern
ch := make(chan int, 10)
go producer(ch)
go consumer(ch)

// 5. Worker Pool
jobs := make(chan Job, 100)
for i := 0; i < numWorkers; i++ {
    go worker(jobs)
}
```

---

## Performance Tips

### 1. Goroutine Management:
```go
// ❌ Avoid: Creating too many goroutines unnecessarily
for i := 0; i < 1000000; i++ {
    go doSmallWork()  // Overhead!
}

// ✅ Better: Use worker pool
jobs := make(chan Work, 1000)
for i := 0; i < 100; i++ {  // Limited workers
    go worker(jobs)
}
```

### 2. Channel Buffering:
```go
// ❌ Avoid: Unbuffered for high throughput
ch := make(chan int)  // Synchronous, slower

// ✅ Better: Buffered for async
ch := make(chan int, 100)  // Asynchronous, faster
```

### 3. Mutex Granularity:
```go
// ❌ Avoid: Coarse-grained locking
mu.Lock()
processData()      // Long operation
sendData()         // Long operation
mu.Unlock()

// ✅ Better: Fine-grained locking
mu.Lock()
data := copyData()
mu.Unlock()

processData(data)  // Outside lock
sendData(data)     // Outside lock
```

### 4. Avoid Goroutine Leaks:
```go
// ❌ Avoid: No timeout
func fetchData() {
    ch := make(chan Data)
    go func() {
        data := expensiveOperation()
        ch <- data  // May block forever!
    }()
    return <-ch
}

// ✅ Better: With timeout
func fetchData() (Data, error) {
    ch := make(chan Data, 1)
    go func() {
        data := expensiveOperation()
        ch <- data
    }()
    
    select {
    case data := <-ch:
        return data, nil
    case <-time.After(5 * time.Second):
        return Data{}, errors.New("timeout")
    }
}
```

---

## Resources

### Official Documentation:
- Go Scheduler Design Doc
- Effective Go
- Go Memory Model

### Tools:
```bash
# Race detector
go run -race main.go

# Profiling
go tool pprof http://localhost:6060/debug/pprof/goroutine

# Trace
go tool trace trace.out

# Check goroutine count
runtime.NumGoroutine()
```

---

## Final Tips for Interview:

1. **শুরুতে সহজ উত্তর দিন**, তারপর detail এ যান
2. **Example দিয়ে ব্যাখ্যা করুন**
3. **Diagram আঁকতে পারলে আঁকুন**
4. **Trade-offs discuss করুন**
5. **Real-world scenarios relate করুন**

**মনে রাখবেন:**
- Goroutines lightweight
- Channels for communication
- Scheduler cooperative + preemptive
- Netpoller non-blocking I/O
- M:P:G model efficient concurrency

---

 

*Created: 2026-01-31*
 