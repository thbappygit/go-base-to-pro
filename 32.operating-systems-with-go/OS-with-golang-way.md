# Operating System Concepts - Golang Way 

## Table of Contents
- [Process - প্রসেস](#process---প্রসেস)
- [Thread - থ্রেড](#thread---থ্রেড)
- [Goroutines - গোরুটিন](#goroutines---গোরুটিন)
- [PCB (Process Control Block)](#pcb-process-control-block)
- [Context Switching - কনটেক্সট স্যুইচিং](#context-switching---কনটেক্সট-স্যুইচিং)
- [Concurrency - কনকারেন্সি](#concurrency---কনকারেন্সি)
- [Parallelism - প্যারালেলিজম](#parallelism---প্যারালেলিজম)
- [Interview Questions](#interview-questions)

---

## Parallelism - প্যারালেলিজম

### Parallelism কী?
একাধিক tasks literally একই সময়ে execute হয়। Go তে `GOMAXPROCS` set করে multiple CPU cores use করা যায়।

### Basic Parallelism Example

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func cpuIntensiveTask(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    // CPU-bound কাজ
    sum := 0
    for i := 0; i < 100000000; i++ {
        sum += i
    }
    
    fmt.Printf("Task %d completed on CPU core\n", id)
}

func main() {
    fmt.Println("=== Parallelism Demo ===\n")
    
    // Available CPU cores
    numCPU := runtime.NumCPU()
    fmt.Printf("Available CPU Cores: %d\n", numCPU)
    
    // Test 1: Single Core (No Parallelism)
    fmt.Println("\nTest 1: Single Core (Sequential-like)")
    runtime.GOMAXPROCS(1)
    start := time.Now()
    var wg1 sync.WaitGroup
    for i := 1; i <= 4; i++ {
        wg1.Add(1)
        go cpuIntensiveTask(i, &wg1)
    }
    wg1.Wait()
    fmt.Printf("Time with 1 core: %v\n", time.Since(start))
    
    // Test 2: Multiple Cores (True Parallelism)
    fmt.Println("\nTest 2: Multiple Cores (Parallel)")
    runtime.GOMAXPROCS(numCPU)
    start = time.Now()
    var wg2 sync.WaitGroup
    for i := 1; i <= 4; i++ {
        wg2.Add(1)
        go cpuIntensiveTask(i, &wg2)
    }
    wg2.Wait()
    fmt.Printf("Time with %d cores: %v (দ্রুত!)\n", numCPU, time.Since(start))
}
```

### Data Parallelism

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

// Parallel array processing
func processChunk(data []int, start, end int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    sum := 0
    for i := start; i < end; i++ {
        sum += data[i] * 2 // Data processing
    }
    
    results <- sum
}

func main() {
    fmt.Println("=== Data Parallelism ===\n")
    
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    // বড় array
    data := make([]int, 1000000)
    for i := 0; i < len(data); i++ {
        data[i] = i + 1
    }
    
    // Array কে chunks এ ভাগ করা (parallel processing এর জন্য)
    numWorkers := 4
    chunkSize := len(data) / numWorkers
    
    results := make(chan int, numWorkers)
    var wg sync.WaitGroup
    
    fmt.Printf("Processing with %d parallel workers...\n", numWorkers)
    
    // প্রতিটি worker একটা chunk process করবে (parallel এ)
    for i := 0; i < numWorkers; i++ {
        start := i * chunkSize
        end := start + chunkSize
        if i == numWorkers-1 {
            end = len(data) // শেষ chunk
        }
        
        wg.Add(1)
        go processChunk(data, start, end, results, &wg)
    }
    
    // Wait এবং results collect করা
    go func() {
        wg.Wait()
        close(results)
    }()
    
    totalSum := 0
    for sum := range results {
        totalSum += sum
    }
    
    fmt.Printf("Total Sum: %d\n", totalSum)
}
```

### Task Parallelism

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Different tasks যা parallel এ run করবে
func taskImageProcessing(wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Task 1: Image Processing শুরু...")
    time.Sleep(2 * time.Second)
    fmt.Println("Task 1: Image Processing শেষ!")
}

func taskVideoEncoding(wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Task 2: Video Encoding শুরু...")
    time.Sleep(3 * time.Second)
    fmt.Println("Task 2: Video Encoding শেষ!")
}

func taskDataAnalysis(wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Task 3: Data Analysis শুরু...")
    time.Sleep(2 * time.Second)
    fmt.Println("Task 3: Data Analysis শেষ!")
}

func taskAudioProcessing(wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Task 4: Audio Processing শুরু...")
    time.Sleep(1 * time.Second)
    fmt.Println("Task 4: Audio Processing শেষ!")
}

func main() {
    fmt.Println("=== Task Parallelism ===\n")
    
    runtime.GOMAXPROCS(runtime.NumCPU())
    fmt.Printf("Using %d CPU cores\n\n", runtime.NumCPU())
    
    start := time.Now()
    var wg sync.WaitGroup
    
    // চারটা different tasks parallel এ run করবে
    wg.Add(4)
    go taskImageProcessing(&wg)
    go taskVideoEncoding(&wg)
    go taskDataAnalysis(&wg)
    go taskAudioProcessing(&wg)
    
    wg.Wait()
    
    fmt.Printf("\nসব tasks শেষ! Total time: %v\n", time.Since(start))
    fmt.Println("(Sequential এ চললে ~8 seconds লাগতো)")
}
```

### Worker Pool Pattern (Parallel Processing)

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

type Job struct {
    ID    int
    Value int
}

type Result struct {
    JobID  int
    Result int
}

// Worker function যা jobs process করবে
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for job := range jobs {
        fmt.Printf("Worker %d: Processing job %d\n", id, job.ID)
        
        // CPU-intensive কাজ simulate করা
        time.Sleep(500 * time.Millisecond)
        result := job.Value * job.Value
        
        results <- Result{JobID: job.ID, Result: result}
    }
}

func main() {
    fmt.Println("=== Worker Pool Pattern ===\n")
    
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    numWorkers := 4
    numJobs := 20
    
    jobs := make(chan Job, numJobs)
    results := make(chan Result, numJobs)
    
    // Worker pool তৈরি করা (parallel workers)
    var wg sync.WaitGroup
    fmt.Printf("Starting %d parallel workers...\n\n", numWorkers)
    
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }
    
    // Jobs পাঠানো
    go func() {
        for j := 1; j <= numJobs; j++ {
            jobs <- Job{ID: j, Value: j * 10}
        }
        close(jobs)
    }()
    
    // Wait এবং results close করা
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Results collect করা
    fmt.Println("\nResults:")
    for result := range results {
        fmt.Printf("Job %d: Result = %d\n", result.JobID, result.Result)
    }
}
```

---

## Concurrency vs Parallelism - তুলনা

### Visual Comparison in Code

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func task(name string, duration time.Duration) {
    fmt.Printf("[%s] শুরু হলো\n", name)
    time.Sleep(duration)
    fmt.Printf("[%s] শেষ হলো\n", name)
}

func main() {
    fmt.Println("=== Concurrency vs Parallelism ===\n")
    
    // Test 1: Concurrency (Single Core)
    fmt.Println("1. CONCURRENCY (Single Core - Time Sharing)")
    fmt.Println("   Tasks interleaved হবে কিন্তু একসাথে চলবে না\n")
    
    runtime.GOMAXPROCS(1) // Force single core
    start := time.Now()
    var wg1 sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg1.Add(1)
        go func(id int) {
            defer wg1.Done()
            task(fmt.Sprintf("Task-%d", id), 1*time.Second)
        }(i)
    }
    wg1.Wait()
    fmt.Printf("   Time: %v (সব tasks সময় নিয়েছে)\n\n", time.Since(start))
    
    // Test 2: Parallelism (Multiple Cores)
    fmt.Println("2. PARALLELISM (Multiple Cores - True Parallel)")
    fmt.Println("   Tasks literally একই সময়ে চলবে\n")
    
    runtime.GOMAXPROCS(runtime.NumCPU()) // Use all cores
    start = time.Now()
    var wg2 sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg2.Add(1)
        go func(id int) {
            defer wg2.Done()
            // CPU-bound task
            sum := 0
            for j := 0; j < 100000000; j++ {
                sum += j
            }
            fmt.Printf("[Task-%d] শেষ হলো (Sum: %d)\n", id, sum)
        }(i)
    }
    wg2.Wait()
    fmt.Printf("   Time: %v (দ্রুত! parallel execution)\n", time.Since(start))
}
```

### Detailed Comparison

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func demonstrateDifference() {
    fmt.Println("=== বিস্তারিত তুলনা ===\n")
    
    // Concurrency Example
    fmt.Println("CONCURRENCY উদাহরণ:")
    fmt.Println("একজন শেফ তিনটা dish রান্না করছে:")
    
    dishes := []string{"Pasta", "Salad", "Soup"}
    
    // Simulate একজন শেফ (single core)
    runtime.GOMAXPROCS(1)
    
    for _, dish := range dishes {
        go func(d string) {
            for i := 1; i <= 3; i++ {
                fmt.Printf("  %s - Step %d\n", d, i)
                time.Sleep(200 * time.Millisecond)
                runtime.Gosched() // অন্য task কে সুযোগ দেওয়া
            }
        }(dish)
    }
    
    time.Sleep(2 * time.Second)
    fmt.Println("  → সব dish একটু একটু করে এগিয়েছে (interleaved)\n")
    
    // Parallelism Example
    fmt.Println("PARALLELISM উদাহরণ:")
    fmt.Println("তিনজন শেফ একসাথে কাজ করছে:")
    
    runtime.GOMAXPROCS(3) // 3 cores
    
    for i, dish := range dishes {
        go func(id int, d string) {
            fmt.Printf("  Chef-%d: %s তৈরি করছি...\n", id+1, d)
            time.Sleep(1 * time.Second)
            fmt.Printf("  Chef-%d: %s তৈরি হয়ে গেছে!\n", id+1, d)
        }(i, dish)
    }
    
    time.Sleep(2 * time.Second)
    fmt.Println("  → সব dish একসাথে তৈরি হয়েছে (parallel)\n")
}

func main() {
    demonstrateDifference()
    
    // Summary
    fmt.Println("\n=== মূল পার্থক্য ===")
    fmt.Println("┌─────────────────┬──────────────────────────────────────┐")
    fmt.Println("│ Concurrency     │ Dealing with multiple things at once │")
    fmt.Println("│ Parallelism     │ Doing multiple things at once        │")
    fmt.Println("├─────────────────┼──────────────────────────────────────┤")
    fmt.Println("│ Single core OK  │ Multiple cores লাগবে                │")
    fmt.Println("│ Interleaved     │ Simultaneous                         │")
    fmt.Println("│ Time-sharing    │ True parallel execution              │")
    fmt.Println("└─────────────────┴──────────────────────────────────────┘")
}
```

### When to Use Which? (Practical Example)

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// I/O Bound Task (Concurrency এর জন্য ভালো)
func fetchURL(url string, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Fetching %s...\n", url)
    time.Sleep(1 * time.Second) // Network I/O simulate
    fmt.Printf("Fetched %s\n", url)
}

// CPU Bound Task (Parallelism এর জন্য ভালো)
func heavyComputation(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    sum := 0
    for i := 0; i < 100000000; i++ {
        sum += i
    }
    fmt.Printf("Computation %d done: %d\n", id, sum)
}

func main() {
    fmt.Println("=== কখন কী ব্যবহার করবেন? ===\n")
    
    // Use Case 1: I/O Bound (Concurrency)
    fmt.Println("1. I/O Bound Tasks (Concurrency):")
    fmt.Println("   Network requests - Single core এই যথেষ্ট\n")
    
    runtime.GOMAXPROCS(1)
    start := time.Now()
    var wg1 sync.WaitGroup
    
    urls := []string{"api.com/1", "api.com/2", "api.com/3"}
    for _, url := range urls {
        wg1.Add(1)
        go fetchURL(url, &wg1)
    }
    wg1.Wait()
    fmt.Printf("   Time: %v (concurrent I/O)\n\n", time.Since(start))
    
    // Use Case 2: CPU Bound (Parallelism)
    fmt.Println("2. CPU Bound Tasks (Parallelism):")
    fmt.Println("   Heavy computations - Multiple cores দরকার\n")
    
    runtime.GOMAXPROCS(runtime.NumCPU())
    start = time.Now()
    var wg2 sync.WaitGroup
    
    for i := 1; i <= 4; i++ {
        wg2.Add(1)
        go heavyComputation(i, &wg2)
    }
    wg2.Wait()
    fmt.Printf("   Time: %v (parallel computation)\n", time.Since(start))
}
```

---

## Interview Questions

### Q1: Process এবং Thread এর পার্থক্য? Go তে কিভাবে কাজ করে?

```go
package main

import "fmt"

func explainProcessVsThread() {
    fmt.Println("=== Process vs Thread ===\n")
    
    fmt.Println("PROCESS:")
    fmt.Println("- Separate memory space")
    fmt.Println("- Heavy weight (বেশি resource)")
    fmt.Println("- Inter-process communication complex")
    fmt.Println("- Example: প্রতিটি Go program একটা process")
    fmt.Println()
    
    fmt.Println("THREAD:")
    fmt.Println("- Shared memory space")
    fmt.Println("- Light weight (কম resource)")
    fmt.Println("- Communication easy (shared memory)")
    fmt.Println("- Example: Go তে goroutines (এরা threads থেকেও lighter)")
    fmt.Println()
    
    fmt.Println("GO তে:")
    fmt.Println("- Goroutines OS threads এর উপর multiplex হয়")
    fmt.Println("- অনেক goroutines কিন্তু কম OS threads")
    fmt.Println("- Go runtime scheduler manage করে")
}

func main() {
    explainProcessVsThread()
}
```

**Answer:**
- **Process**: Isolated memory, heavy, আলাদা program instance
- **Thread**: Shared memory, lighter, same process এর মধ্যে
- **Go**: Goroutines use করে যা threads থেকেও lightweight এবং Go runtime দ্বারা efficiently managed

### Q2: Context Switching কেন overhead তৈরি করে?

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateOverhead() {
    fmt.Println("=== Context Switching Overhead ===\n")
    
    fmt.Println("Context Switch এ যা হয়:")
    fmt.Println("1. Current state save (registers, PC, SP)")
    fmt.Println("2. PCB update করা")
    fmt.Println("3. Memory mapping change")
    fmt.Println("4. Cache invalidation (cold cache)")
    fmt.Println("5. TLB flush")
    fmt.Println("6. New state restore")
    fmt.Println()
    
    fmt.Println("এই পুরো সময় কোন useful কাজ হয় না!")
    fmt.Println("তাই বেশি context switching = performance কমে যায়")
}

func main() {
    demonstrateOverhead()
}
```

**Answer:**
Context switching overhead কারণ:
- State save/restore করতে সময়
- Cache miss (cold cache)
- TLB flush
- Memory mapping change
- এই সময় productive work হয় না

### Q3: Goroutine এবং OS Thread এর পার্থক্য?

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("=== Goroutine vs OS Thread ===\n")
    
    fmt.Printf("OS Threads: %d\n", runtime.GOMAXPROCS(0))
    fmt.Printf("Goroutines: %d\n\n", runtime.NumGoroutine())
    
    // 1000 goroutines তৈরি করা
    for i := 0; i < 1000; i++ {
        go func() {
            select {} // চিরকাল block
        }()
    }
    
    fmt.Printf("After creating 1000 goroutines:\n")
    fmt.Printf("OS Threads: %d (প্রায় same!)\n", runtime.GOMAXPROCS(0))
    fmt.Printf("Goroutines: %d\n\n", runtime.NumGoroutine())
    
    fmt.Println("পার্থক্য:")
    fmt.Println("┌────────────┬──────────────┬─────────────┐")
    fmt.Println("│            │ Goroutine    │ OS Thread   │")
    fmt.Println("├────────────┼──────────────┼─────────────┤")
    fmt.Println("│ Stack Size │ ~2KB (grow)  │ ~1-2MB      │")
    fmt.Println("│ Creation   │ Fast (~µs)   │ Slow (~ms)  │")
    fmt.Println("│ Switching  │ Cheap        │ Expensive   │")
    fmt.Println("│ Managed by │ Go runtime   │ OS kernel   │")
    fmt.Println("└────────────┴──────────────┴─────────────┘")
}
```

**Answer:**
- **Goroutine**: Lightweight (2KB stack), fast creation, Go runtime manages
- **OS Thread**: Heavy (1-2MB stack), slow creation, OS kernel manages
- Go তে M:N scheduling: M goroutines, N OS threads এ run করে

### Q4: Race Condition কী এবং কিভাবে solve করবেন?

```go
package main

import (
    "fmt"
    "sync"
)

var counter int
var mutex sync.Mutex

// Race condition (problem)
func incrementRacy(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        counter++ // UNSAFE!
    }
}

// Solution: Mutex
func incrementSafe(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        mutex.Lock()
        counter++ // SAFE
        mutex.Unlock()
    }
}

func main() {
    fmt.Println("=== Race Condition Demo ===\n")
    
    // Test: Race condition
    counter = 0
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go incrementRacy(&wg)
    }
    wg.Wait()
    fmt.Printf("Race condition: Expected 10000, Got %d ❌\n\n", counter)
    
    // Solution
    fmt.Println("Solutions:")
    fmt.Println("1. Mutex (sync.Mutex)")
    fmt.Println("2. Channels (Go way)")
    fmt.Println("3. Atomic operations (sync/atomic)")
}
```

**Answer:**
- **Race Condition**: একাধিক goroutines shared data access করলে এবং result execution order এর উপর depend করলে
- **Solutions**: Mutex, Channels, Atomic operations

### Q5: Concurrency এবং Parallelism কখন ব্যবহার করবেন?

```go
package main

import "fmt"

func main() {
    fmt.Println("=== Concurrency vs Parallelism - কখন কী? ===\n")
    
    fmt.Println("CONCURRENCY ব্যবহার করুন যখন:")
    fmt.Println("✓ I/O bound tasks (network, disk, database)")
    fmt.Println("✓ Responsive UI চাই")
    fmt.Println("✓ Multiple tasks manage করতে চাই")
    fmt.Println("✓ Example: Web server (multiple requests)")
    fmt.Println()
    
    fmt.Println("PARALLELISM ব্যবহার করুন যখন:")
    fmt.Println("✓ CPU bound tasks (computation, processing)")
    fmt.Println("✓ বড় data process করতে হবে")
    fmt.Println("✓ Multiple cores available")
    fmt.Println("✓ Example: Video encoding, image processing")
    fmt.Println()
    
    fmt.Println("GO তে:")
    fmt.Println("- Concurrency: goroutines + channels")
    fmt.Println("- Parallelism: runtime.GOMAXPROCS()")
    fmt.Println("- সাধারণত দুটোই একসাথে use হয়!")
}
```

---

## Quick Reference

### Go Concurrency Patterns

```go
// 1. Basic Goroutine
go func() {
    // কাজ করুন
}()

// 2. WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // কাজ
}()
wg.Wait()

// 3. Channels
ch := make(chan int)
go func() { ch <- 42 }()
value := <-ch

// 4. Select
select {
case v := <-ch1:
    // ch1 থেকে data
case v := <-ch2:
    // ch2 থেকে data
}

// 5. Mutex
var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()
```

### Important Functions

```go
// Runtime control
runtime.GOMAXPROCS(n)      // OS threads set করা
runtime.NumGoroutine()     // Current goroutines সংখ্যা
runtime.Gosched()          // CPU ছেড়ে দেওয়া
runtime.NumCPU()           // CPU cores সংখ্যা

// Process info
os.Getpid()                // Current process PID
os.Getppid()               // Parent process PID
```

---

## সারাংশ

### মনে রাখার মূল বিষয়:

1. **Process**: Program in execution, isolated memory
2. **Thread**: Lightweight, shared memory
3. **Goroutine**: Go এর ultra-lightweight threads
4. **PCB**: OS এর data structure, process info রাখে
5. **Context Switching**: Process/goroutine switch করা (overhead আছে)
6. **Concurrency**: একসাথে progress (dealing with)
7. **Parallelism**: একসাথে execution (doing)

### Interview Tip:
- Real code example দিতে পারলে impression ভালো হয়
- Go এর goroutines এবং channels এর সুবিধা বলুন
- Race condition এবং synchronization বুঝুন
- Practical use cases জানুন (I/O vs CPU bound)

**Best of luck! 🚀** Process - প্রসেস

### প্রসেস কী?
একটি **running program** কে process বলে। যখন আপনি একটি Go program compile করে run করেন, OS একটি process তৈরি করে।

### Golang Example

```go
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    // Current process এর PID দেখা যাক
    pid := os.Getpid()
    fmt.Printf("Process ID: %d\n", pid)
    
    // Parent process এর PID
    ppid := os.Getppid()
    fmt.Printf("Parent Process ID: %d\n", ppid)
    
    fmt.Println("Process চলছে...")
    time.Sleep(5 * time.Second)
    fmt.Println("Process শেষ!")
}
```

**Output:**
```
Process ID: 12345
Parent Process ID: 6789
Process চলছে...
Process শেষ!
```

### Process Memory Layout

```go
package main

import "fmt"

// Data Section - Global variables
var globalVar int = 100

func main() {
    // Stack - Local variables
    localVar := 200
    
    // Heap - Dynamic allocation
    heapVar := new(int)
    *heapVar = 300
    
    fmt.Printf("Global (Data): %p\n", &globalVar)
    fmt.Printf("Local (Stack): %p\n", &localVar)
    fmt.Printf("Heap: %p\n", heapVar)
}
```

**ব্যাখ্যা:**
- `globalVar` → Data section এ থাকে
- `localVar` → Stack এ থাকে
- `heapVar` → Heap এ allocate হয়

### Process States Example

```go
package main

import (
    "fmt"
    "time"
)

func simulateProcessStates() {
    // NEW state - Process শুরু হচ্ছে
    fmt.Println("NEW: Process তৈরি হচ্ছে")
    
    // READY state - CPU পাওয়ার জন্য ready
    fmt.Println("READY: CPU এর জন্য অপেক্ষা করছি")
    
    // RUNNING state - Execute হচ্ছে
    fmt.Println("RUNNING: Code execute হচ্ছে")
    for i := 0; i < 3; i++ {
        fmt.Printf("কাজ করছি... %d\n", i)
        time.Sleep(1 * time.Second)
    }
    
    // WAITING state - I/O এর জন্য অপেক্ষা
    fmt.Println("WAITING: I/O operation এর জন্য blocked")
    time.Sleep(2 * time.Second) // I/O simulate করছি
    
    // RUNNING state - আবার চলছে
    fmt.Println("RUNNING: আবার execute হচ্ছে")
    
    // TERMINATED state - শেষ
    fmt.Println("TERMINATED: Process শেষ!")
}

func main() {
    simulateProcessStates()
}
```

### Child Process তৈরি করা

```go
package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    fmt.Printf("Parent Process PID: %d\n", os.Getpid())
    
    // Child process তৈরি করা
    cmd := exec.Command("go", "version")
    
    // Child process run করা
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Printf("Child process output: %s\n", output)
    fmt.Println("Parent process চলছে...")
}
```

---

## Thread - থ্রেড

### Thread কী?
Process এর মধ্যে execution এর smallest unit। একটি process এর মধ্যে একাধিক threads থাকতে পারে যারা same memory share করে।

### Traditional Threading (OS Level)

Go তে direct OS threads control করা যায় না, কিন্তু concept বোঝার জন্য:

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// এটা OS thread এর মতো কাজ করে না, 
// কিন্তু concept বোঝার জন্য উদাহরণ

var counter int
var mutex sync.Mutex

func threadLikeFunction(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 5; i++ {
        mutex.Lock()
        counter++
        fmt.Printf("Thread %d: Counter = %d\n", id, counter)
        mutex.Unlock()
        
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // OS threads এর সংখ্যা set করা
    runtime.GOMAXPROCS(4) // 4টি OS threads use করবে
    
    var wg sync.WaitGroup
    
    // 3টি "thread-like" behavior
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go threadLikeFunction(i, &wg)
    }
    
    wg.Wait()
    fmt.Printf("Final Counter: %d\n", counter)
}
```

### Thread Memory Sharing

```go
package main

import (
    "fmt"
    "sync"
)

// Shared memory
var sharedData = []int{1, 2, 3, 4, 5}
var mutex sync.Mutex

func readerThread(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    mutex.Lock()
    fmt.Printf("Reader %d: %v\n", id, sharedData)
    mutex.Unlock()
}

func writerThread(id int, value int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    mutex.Lock()
    sharedData = append(sharedData, value)
    fmt.Printf("Writer %d: Added %d\n", id, value)
    mutex.Unlock()
}

func main() {
    var wg sync.WaitGroup
    
    // সবাই same memory (sharedData) access করছে
    wg.Add(4)
    go readerThread(1, &wg)
    go writerThread(1, 100, &wg)
    go readerThread(2, &wg)
    go writerThread(2, 200, &wg)
    
    wg.Wait()
}
```

---

## Goroutines - গোরুটিন

### Goroutine কী?
Go এর lightweight threads। OS threads থেকে অনেক হালকা এবং efficient। Go runtime নিজেই goroutines manage করে।

### Goroutine vs OS Thread

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func goroutineExample(id int) {
    fmt.Printf("Goroutine %d শুরু হলো\n", id)
    time.Sleep(1 * time.Second)
    fmt.Printf("Goroutine %d শেষ হলো\n", id)
}

func main() {
    // কতগুলো OS threads use হচ্ছে দেখা যাক
    fmt.Printf("OS Threads (GOMAXPROCS): %d\n", runtime.GOMAXPROCS(0))
    fmt.Printf("Current Goroutines: %d\n", runtime.NumGoroutine())
    
    // 1000 goroutines তৈরি করা (কিন্তু OS threads কম!)
    for i := 1; i <= 1000; i++ {
        go goroutineExample(i)
    }
    
    time.Sleep(100 * time.Millisecond)
    fmt.Printf("Goroutines চলছে: %d\n", runtime.NumGoroutine())
    
    time.Sleep(2 * time.Second)
    fmt.Printf("শেষে Goroutines: %d\n", runtime.NumGoroutine())
}
```

**Key Point:**
- 1000 goroutines তৈরি হলো
- কিন্তু হয়তো মাত্র 4-8টা OS threads use হচ্ছে
- Go runtime efficiently goroutines schedule করে

### Goroutine Scheduling

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    // CPU-bound কাজ
    sum := 0
    for i := 0; i < 1000000; i++ {
        sum += i
    }
    
    fmt.Printf("Worker %d: Sum = %d, Thread = %d\n", 
        id, sum, runtime.NumCPU())
}

func main() {
    runtime.GOMAXPROCS(2) // 2টা OS threads use করবে
    
    var wg sync.WaitGroup
    
    // 10টা goroutines কিন্তু 2টা threads এ চলবে
    for i := 1; i <= 10; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    
    wg.Wait()
    fmt.Println("সব workers শেষ!")
}
```

**ব্যাখ্যা:**
- 10টা goroutines তৈরি হলো
- কিন্তু মাত্র 2টা OS threads এ চলছে
- Go scheduler automatically goroutines distribute করে

### Goroutine Communication (Channels)

```go
package main

import (
    "fmt"
    "time"
)

func producer(ch chan int) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Producing: %d\n", i)
        ch <- i // Data পাঠানো
        time.Sleep(500 * time.Millisecond)
    }
    close(ch)
}

func consumer(ch chan int) {
    for num := range ch {
        fmt.Printf("Consuming: %d\n", num)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    ch := make(chan int, 2) // Buffered channel
    
    go producer(ch)
    go consumer(ch)
    
    time.Sleep(10 * time.Second)
}
```

---

## PCB (Process Control Block)

### PCB কী?
OS এর data structure যা প্রতিটি process এর information রাখে। Go তে directly PCB access করা যায় না, কিন্তু concept বোঝা যায়।

### PCB Simulation in Go

```go
package main

import (
    "fmt"
    "time"
)

// PCB এর structure simulate করা
type ProcessControlBlock struct {
    PID              int
    State            string
    ProgramCounter   int
    Registers        map[string]int
    Priority         int
    CPUTime          time.Duration
    MemoryInfo       MemoryInfo
    OpenFiles        []string
}

type MemoryInfo struct {
    BaseAddress  uintptr
    LimitAddress uintptr
    PageTable    []int
}

func main() {
    // একটি PCB তৈরি করা
    pcb := ProcessControlBlock{
        PID:            12345,
        State:          "RUNNING",
        ProgramCounter: 1024,
        Registers: map[string]int{
            "R0": 100,
            "R1": 200,
            "R2": 300,
        },
        Priority: 5,
        CPUTime:  time.Second * 10,
        MemoryInfo: MemoryInfo{
            BaseAddress:  0x1000,
            LimitAddress: 0x5000,
            PageTable:    []int{0, 1, 2, 3},
        },
        OpenFiles: []string{"file1.txt", "file2.txt"},
    }
    
    // PCB information display করা
    fmt.Println("=== Process Control Block ===")
    fmt.Printf("PID: %d\n", pcb.PID)
    fmt.Printf("State: %s\n", pcb.State)
    fmt.Printf("Program Counter: %d\n", pcb.ProgramCounter)
    fmt.Printf("Priority: %d\n", pcb.Priority)
    fmt.Printf("CPU Time: %v\n", pcb.CPUTime)
    fmt.Printf("Registers: %v\n", pcb.Registers)
    fmt.Printf("Memory Base: 0x%X\n", pcb.MemoryInfo.BaseAddress)
    fmt.Printf("Open Files: %v\n", pcb.OpenFiles)
}
```

### Context Save/Restore Simulation

```go
package main

import "fmt"

type Context struct {
    ProgramCounter int
    StackPointer   int
    Registers      [4]int
}

type Process struct {
    PID     int
    Name    string
    Context Context
}

func saveContext(p *Process) Context {
    fmt.Printf("Saving context for %s (PID: %d)\n", p.Name, p.PID)
    return p.Context
}

func restoreContext(p *Process, ctx Context) {
    fmt.Printf("Restoring context for %s (PID: %d)\n", p.Name, p.PID)
    p.Context = ctx
}

func main() {
    // Process 1
    p1 := Process{
        PID:  1,
        Name: "Process-A",
        Context: Context{
            ProgramCounter: 100,
            StackPointer:   500,
            Registers:      [4]int{10, 20, 30, 40},
        },
    }
    
    // Process 2
    p2 := Process{
        PID:  2,
        Name: "Process-B",
        Context: Context{
            ProgramCounter: 200,
            StackPointer:   600,
            Registers:      [4]int{50, 60, 70, 80},
        },
    }
    
    fmt.Println("=== Context Switching Simulation ===")
    
    // P1 চলছে
    fmt.Printf("\n%s চলছে...\n", p1.Name)
    fmt.Printf("PC: %d, SP: %d\n", p1.Context.ProgramCounter, p1.Context.StackPointer)
    
    // Context switch: P1 -> P2
    fmt.Println("\n--- Context Switch: P1 -> P2 ---")
    savedCtx := saveContext(&p1)
    restoreContext(&p2, p2.Context)
    
    fmt.Printf("\n%s চলছে...\n", p2.Name)
    fmt.Printf("PC: %d, SP: %d\n", p2.Context.ProgramCounter, p2.Context.StackPointer)
    
    // Context switch: P2 -> P1
    fmt.Println("\n--- Context Switch: P2 -> P1 ---")
    saveContext(&p2)
    restoreContext(&p1, savedCtx)
    
    fmt.Printf("\n%s আবার চলছে...\n", p1.Name)
    fmt.Printf("PC: %d, SP: %d\n", p1.Context.ProgramCounter, p1.Context.StackPointer)
}
```

---

## Context Switching - কনটেক্সট স্যুইচিং

### Context Switching কী?
CPU একটি process/goroutine থেকে আরেকটিতে switch করা। Go runtime automatically goroutines এর মধ্যে context switch করে।

### Goroutine Context Switching

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func task(name string, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 1; i <= 3; i++ {
        fmt.Printf("[%s] Step %d - Goroutine %d\n", 
            name, i, runtime.NumGoroutine())
        
        // অন্য goroutines কে CPU দেওয়ার সুযোগ
        runtime.Gosched() // Manual context switch
        
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    var wg sync.WaitGroup
    
    fmt.Println("=== Context Switching Demo ===")
    fmt.Printf("Starting Goroutines: %d\n\n", runtime.NumGoroutine())
    
    // 3টা goroutines যারা context switch করবে
    tasks := []string{"Task-A", "Task-B", "Task-C"}
    
    for _, t := range tasks {
        wg.Add(1)
        go task(t, &wg)
    }
    
    wg.Wait()
    fmt.Printf("\nEnding Goroutines: %d\n", runtime.NumGoroutine())
}
```

### Context Switching Overhead Example

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func heavyTask(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    sum := 0
    for i := 0; i < 100000; i++ {
        sum += i
        // বেশি context switching = বেশি overhead
        if i%1000 == 0 {
            runtime.Gosched()
        }
    }
}

func main() {
    runtime.GOMAXPROCS(1) // Single thread
    
    // Test 1: বেশি context switching
    fmt.Println("Test 1: বেশি context switching")
    start := time.Now()
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go heavyTask(i, &wg)
    }
    wg.Wait()
    fmt.Printf("Time: %v\n\n", time.Since(start))
    
    // Test 2: কম context switching
    fmt.Println("Test 2: কম context switching (sequential)")
    start = time.Now()
    for i := 0; i < 100; i++ {
        sum := 0
        for j := 0; j < 100000; j++ {
            sum += j
        }
    }
    fmt.Printf("Time: %v\n", time.Since(start))
}
```

### Context Switch Triggers

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func demonstrateContextSwitches() {
    fmt.Println("=== Context Switch Triggers ===\n")
    
    // 1. Channel operation এ blocking
    fmt.Println("1. Channel Blocking:")
    ch := make(chan int)
    go func() {
        fmt.Println("   Goroutine: Channel থেকে receive করার চেষ্টা...")
        <-ch // Block হবে, context switch হবে
        fmt.Println("   Goroutine: Data received!")
    }()
    time.Sleep(100 * time.Millisecond)
    fmt.Println("   Main: Data পাঠাচ্ছি...")
    ch <- 42
    time.Sleep(100 * time.Millisecond)
    
    // 2. Sleep এ blocking
    fmt.Println("\n2. Sleep (I/O like operation):")
    go func() {
        fmt.Println("   Goroutine: Sleep করছি...")
        time.Sleep(1 * time.Second) // Context switch হবে
        fmt.Println("   Goroutine: জেগে উঠলাম!")
    }()
    fmt.Println("   Main: অন্য কাজ করতে পারি")
    time.Sleep(2 * time.Second)
    
    // 3. Manual context switch (Gosched)
    fmt.Println("\n3. Manual Context Switch (Gosched):")
    go func() {
        for i := 0; i < 3; i++ {
            fmt.Printf("   Goroutine: Step %d\n", i)
            runtime.Gosched() // CPU ছেড়ে দিচ্ছি
        }
    }()
    for i := 0; i < 3; i++ {
        fmt.Printf("   Main: Step %d\n", i)
        runtime.Gosched()
    }
    time.Sleep(100 * time.Millisecond)
    
    // 4. System call
    fmt.Println("\n4. System Call:")
    go func() {
        fmt.Println("   Goroutine: File operation করছি...")
        // File operation system call করবে, context switch হবে
        time.Sleep(500 * time.Millisecond)
        fmt.Println("   Goroutine: File operation শেষ!")
    }()
    time.Sleep(1 * time.Second)
}

func main() {
    demonstrateContextSwitches()
}
```

---

## Concurrency - কনকারেন্সি

### Concurrency কী?
একাধিক tasks একই সময়ে progress করছে মনে হয়। Go তে channels এবং goroutines দিয়ে concurrency implement করা হয়।

### Basic Concurrency Example

```go
package main

import (
    "fmt"
    "time"
)

func task1() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Task-1: Step %d\n", i)
        time.Sleep(300 * time.Millisecond)
    }
}

func task2() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Task-2: Step %d\n", i)
        time.Sleep(400 * time.Millisecond)
    }
}

func task3() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Task-3: Step %d\n", i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    fmt.Println("=== Concurrent Execution ===\n")
    
    // তিনটি tasks concurrently চলবে
    go task1()
    go task2()
    go task3()
    
    // সব tasks শেষ হওয়ার জন্য অপেক্ষা
    time.Sleep(3 * time.Second)
    fmt.Println("\nসব tasks শেষ!")
}
```

**Output দেখবেন:**
```
Task-1: Step 1
Task-2: Step 1
Task-3: Step 1
Task-1: Step 2
Task-2: Step 2
Task-1: Step 3
... (interleaved output)
```

### Producer-Consumer Pattern (Classic Concurrency)

```go
package main

import (
    "fmt"
    "time"
)

func producer(name string, ch chan<- int) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("[Producer %s] তৈরি করছি: %d\n", name, i)
        ch <- i
        time.Sleep(500 * time.Millisecond)
    }
}

func consumer(name string, ch <-chan int) {
    for num := range ch {
        fmt.Printf("[Consumer %s] ব্যবহার করছি: %d\n", name, num)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    ch := make(chan int, 3) // Buffer size 3
    
    fmt.Println("=== Producer-Consumer Pattern ===\n")
    
    // 2 producers, 1 consumer (concurrent)
    go producer("P1", ch)
    go producer("P2", ch)
    go consumer("C1", ch)
    
    time.Sleep(8 * time.Second)
    close(ch)
    time.Sleep(1 * time.Second)
}
```

### Fan-Out Fan-In Pattern

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Fan-Out: একটা input থেকে multiple workers
func fanOut(input <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    
    for i := 0; i < workers; i++ {
        outputs[i] = worker(i+1, input)
    }
    
    return outputs
}

func worker(id int, input <-chan int) <-chan int {
    output := make(chan int)
    
    go func() {
        defer close(output)
        for num := range input {
            fmt.Printf("Worker-%d: Processing %d\n", id, num)
            time.Sleep(500 * time.Millisecond)
            output <- num * 2 // Processing: double করা
        }
    }()
    
    return output
}

// Fan-In: Multiple inputs থেকে একটা output
func fanIn(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    output := make(chan int)
    
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for num := range c {
                output <- num
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(output)
    }()
    
    return output
}

func main() {
    fmt.Println("=== Fan-Out Fan-In Pattern ===\n")
    
    // Input channel
    input := make(chan int)
    
    // Input data পাঠানো
    go func() {
        for i := 1; i <= 10; i++ {
            input <- i
        }
        close(input)
    }()
    
    // Fan-Out: 3 workers
    workers := fanOut(input, 3)
    
    // Fan-In: সব workers থেকে একটা output
    result := fanIn(workers...)
    
    // Results collect করা
    fmt.Println("\nResults:")
    for r := range result {
        fmt.Printf("Result: %d\n", r)
    }
}
```

### Race Condition এবং Mutex

```go
package main

import (
    "fmt"
    "sync"
)

var counter int

// Race condition (সমস্যা)
func incrementWithoutMutex(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        counter++ // Race condition!
    }
}

// Solution: Mutex দিয়ে
var counterSafe int
var mutex sync.Mutex

func incrementWithMutex(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        mutex.Lock()
        counterSafe++
        mutex.Unlock()
    }
}

func main() {
    // Test 1: Without Mutex (Race Condition)
    fmt.Println("=== Without Mutex (Race Condition) ===")
    counter = 0
    var wg1 sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg1.Add(1)
        go incrementWithoutMutex(&wg1)
    }
    wg1.Wait()
    fmt.Printf("Expected: 10000, Got: %d (ভুল!)\n\n", counter)
    
    // Test 2: With Mutex (Safe)
    fmt.Println("=== With Mutex (Thread-Safe) ===")
    counterSafe = 0
    var wg2 sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg2.Add(1)
        go incrementWithMutex(&wg2)
    }
    wg2.Wait()
    fmt.Printf("Expected: 10000, Got: %d (সঠিক!)\n", counterSafe)
}
```

### Select Statement (Multiple Channels)

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("=== Select Statement Demo ===\n")
    
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    // Channel 1: প্রতি 1 সেকেন্ডে data
    go func() {
        for i := 1; ; i++ {
            time.Sleep(1 * time.Second)
            ch1 <- fmt.Sprintf("Channel-1: Message %d", i)
        }
    }()
    
    // Channel 2: প্রতি 2 সেকেন্ডে data
    go func() {
        for i := 1; ; i++ {
            time.Sleep(2 * time.Second)
            ch2 <- fmt.Sprintf("Channel-2: Message %d", i)
        }
    }()
    
    // Select: যে channel ready সেই থেকে receive করবে
    for i := 0; i < 10; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        case <-time.After(3 * time.Second):
            fmt.Println("Timeout! কোন channel ready না")
        }
    }
}
```

---

##