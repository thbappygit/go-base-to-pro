# Operating System Concepts - Preparation

## Table of Contents
- [Process](#process)
- [Thread](#thread)
- [PCB (Process Control Block)](#pcb-process-control-block)
- [Context Switching](#context-switching)
- [Concurrency](#concurrency)
- [Parallelism](#parallelism)
- [Concurrency vs Parallelism](#concurrency-vs-parallelism)
- [Interview FAQs](#interview-faqs)

---

## Process

### Definition
- একটি program যখন execution এ থাকে তখন তাকে **process** বলে
- Program এর active instance
- প্রতিটি process এর নিজস্ব isolated memory space থাকে

### Process Components
```
┌─────────────────┐
│   Stack         │  ← Function calls, local variables
├─────────────────┤
│   Heap          │  ← Dynamic memory allocation
├─────────────────┤
│   Data          │  ← Global & static variables
├─────────────────┤
│   Text/Code     │  ← Program instructions
└─────────────────┘
```

### Process States
1. **New** → Process তৈরি হচ্ছে
2. **Ready** → CPU পাওয়ার জন্য waiting
3. **Running** → Instructions execute হচ্ছে
4. **Waiting/Blocked** → I/O বা event এর জন্য অপেক্ষা
5. **Terminated** → Execution complete

### State Diagram
```
New → Ready ⇄ Running → Terminated
         ↕       ↓
      Waiting ←─┘
```

---

## Thread

### Definition
- Process এর মধ্যে execution এর **smallest unit**
- **Lightweight process** বলা হয়
- একই process এর সব threads same memory space share করে

### Thread vs Process

| Aspect | Process | Thread |
|--------|---------|--------|
| Memory | Separate memory space | Shared memory space |
| Creation | Slow & resource-intensive | Fast & lightweight |
| Communication | Complex (IPC required) | Easy (shared memory) |
| Overhead | High | Low |
| Isolation | Independent | Dependent on process |

### Types of Threads
1. **User-level Threads**
    - User space এ manage হয়
    - Fast context switching
    - OS জানে না

2. **Kernel-level Threads**
    - OS দ্বারা manage হয়
    - True parallelism support
    - Slower context switching

### Multithreading Benefits
- Better responsiveness
- Resource sharing
- Economy (cheaper than processes)
- Scalability on multi-core systems

---

## PCB (Process Control Block)

### Definition
- OS এর data structure যা **প্রতিটি process এর information** store করে
- Process descriptor বলা হয়
- Process table এ থাকে

### PCB Contents
```
┌──────────────────────────────┐
│  Process ID (PID)            │
│  Process State               │
│  Program Counter             │
│  CPU Registers               │
│  CPU Scheduling Info         │
│  Memory Management Info      │
│  Accounting Information      │
│  I/O Status Information      │
└──────────────────────────────┘
```

### PCB Components Details

| Component | Description |
|-----------|-------------|
| **PID** | Unique process identifier |
| **Process State** | Current state (ready/running/waiting) |
| **Program Counter** | Next instruction এর address |
| **CPU Registers** | Accumulator, index registers, stack pointer |
| **Scheduling Info** | Priority, queue pointers |
| **Memory Info** | Page tables, segment tables, memory limits |
| **I/O Info** | Open files, I/O devices list |
| **Accounting** | CPU time used, time limits |

### PCB Importance
- Context switching এর জন্য essential
- Process management এর core
- Scheduler এর decision making এ use হয়

---

## Context Switching

### Definition
- CPU একটি process/thread থেকে আরেকটিতে **switch** করার mechanism
- Current state save এবং new state restore করা

### Context Switching Steps
```
1. Save current process state → PCB
2. Update process state (running → ready/waiting)
3. Select next process (Scheduler)
4. Restore new process state ← PCB
5. Update process state (ready → running)
6. Resume execution
```

### When Does Context Switching Occur?
- ⏰ **Time slice expiry** (time quantum শেষ)
- 🔔 **Interrupt** (hardware/software)
- 📞 **System call** (I/O request)
- ⬆️ **Higher priority process** আসলে
- 🛑 **Process blocking** (waiting for resource)

### Context Switching Overhead
- **Direct Costs:**
    - Save/restore register values
    - Update PCB
    - Switch memory address space

- **Indirect Costs:**
    - Cache invalidation (cold cache)
    - TLB (Translation Lookaside Buffer) flush
    - Pipeline flush

- **Impact:**
    - Context switching সময় কোন useful work হয় না
    - বেশি context switching = performance degradation

### Optimization Techniques
- Thread ব্যবহার (lighter weight)
- Efficient scheduling algorithms
- Minimize time quantum changes
- Hardware support (multiple register sets)

---

## Concurrency

### Definition
- একাধিক tasks **একই সময়ে progress** করছে মনে হয়
- Tasks overlapping time periods এ execute হয়
- **Single core** এও সম্ভব (time-sharing)

### How It Works
```
Time →
Core: [T1][T2][T1][T3][T2][T1][T3]
```
- CPU দ্রুত tasks এর মধ্যে switch করে
- প্রতিটি task একটু একটু করে এগিয়ে যায়
- Illusion of simultaneity

### Key Characteristics
- ✅ **Dealing with** multiple things at once
- ✅ Single core এ possible
- ✅ Interleaved execution
- ✅ Context switching use করে

### Examples
- Single core system এ multiple programs
- Web browser: page loading + music playing
- Text editor: typing + auto-save

### Concurrency Challenges
- **Race Conditions** → Shared resource access
- **Deadlock** → Circular waiting
- **Starvation** → Process never gets CPU
- **Synchronization** → Coordination needed

### Solutions
- Locks, Semaphores, Monitors
- Mutex (Mutual Exclusion)
- Atomic operations
- Message passing

---

## Parallelism

### Definition
- একাধিক tasks **literally একই সময়ে** execute হয়
- **True simultaneous execution**
- Multiple cores/processors দরকার

### How It Works
```
Time →
Core 1: [────T1────][────T1────]
Core 2: [────T2────][────T2────]
Core 3: [────T3────][────T3────]
Core 4: [────T4────][────T4────]
```

### Types of Parallelism

#### 1. Data Parallelism
- Same operation, different data তে
- Example: Image processing এ প্রতিটি pixel আলাদা core এ

```python
# Example: Array processing
for i in range(len(array)):
    array[i] = array[i] * 2  # Each element parallel এ process হতে পারে
```

#### 2. Task Parallelism
- Different operations simultaneously
- Example: Video encoding এ audio এবং video আলাদা cores এ

### Requirements
- ✅ Multiple CPU cores/processors
- ✅ Parallel programming (OpenMP, MPI, CUDA)
- ✅ Thread-safe code
- ✅ Proper synchronization

### Examples
- Multi-core processor এ different threads
- Video rendering: multiple frames একসাথে
- Scientific computing: matrix operations
- MapReduce in distributed systems

---

## Concurrency vs Parallelism

### Visual Comparison

**Concurrency (Single Core):**
```
Time →
─────[T1]──[T2]──[T1]──[T3]──[T2]──[T1]─────→
     Context switches
```

**Parallelism (Multi-Core):**
```
Time →
Core 1: ─────[T1]─────[T1]─────[T1]─────→
Core 2: ─────[T2]─────[T2]─────[T2]─────→
Core 3: ─────[T3]─────[T3]─────[T3]─────→
```

### Key Differences

| Aspect | Concurrency | Parallelism |
|--------|-------------|-------------|
| **Definition** | Dealing with multiple things | Doing multiple things |
| **Hardware** | Single core এ possible | Multiple cores লাগে |
| **Execution** | Interleaved (সময় ভাগ করে) | Simultaneous (একসাথে) |
| **Focus** | Structure & composition | Actual execution |
| **Mechanism** | Time-sharing, context switching | Multiple processors |
| **Goal** | Better responsiveness | Faster computation |

### Real-World Analogy

**Concurrency:**
```
একজন chef একাধিক dish রান্না করছে
- একটু pasta → একটু salad → একটু soup
- দ্রুত switch করে সবকিছু manage করছে
```

**Parallelism:**
```
তিনজন chef একসাথে কাজ করছে
- Chef 1: pasta তৈরি করছে
- Chef 2: salad তৈরি করছে  
- Chef 3: soup তৈরি করছে
```

### When to Use What?

**Use Concurrency When:**
- I/O bound tasks
- Responsive UI needed
- Resource sharing required
- Single core system

**Use Parallelism When:**
- CPU bound tasks
- Computationally intensive work
- Multiple cores available
- Independent tasks

**Combined Approach:**
- Modern systems use both
- Concurrent tasks running in parallel
- Example: Web server handling multiple requests across multiple cores

---

## Interview FAQs

### Q1: Process এবং Thread এর মধ্যে main পার্থক্য কী?

**Answer:**
- **Memory:** Process isolated, Thread shared
- **Creation:** Process heavyweight, Thread lightweight
- **Communication:** Process complex (IPC), Thread easy (shared memory)
- **Overhead:** Process বেশি, Thread কম
- **Crash:** একটি thread crash হলে পুরো process crash, কিন্তু process independent

### Q2: Context switching কেন overhead তৈরি করে?

**Answer:**
Context switching এ overhead তৈরি হয় কারণ:
- Register values save/restore করতে হয়
- PCB update করতে হয়
- Memory mapping change করতে হয়
- Cache invalidate হয়ে যায় (cold cache)
- TLB flush হতে পারে
- এই পুরো সময় কোন useful work হয় না

### Q3: PCB তে কী কী information থাকে?

**Answer:**
- Process ID (unique identifier)
- Process State (ready/running/waiting)
- Program Counter (next instruction)
- CPU Registers (সব register values)
- Scheduling information (priority, queue)
- Memory management (page tables)
- I/O status (open files, devices)
- Accounting (CPU time used)

### Q4: Concurrency এবং Parallelism এর পার্থক্য?

**Answer:**
- **Concurrency:** Multiple tasks progress করছে, কিন্তু একসাথে execute হচ্ছে না। Single core এও সম্ভব।
- **Parallelism:** Multiple tasks literally একই সময়ে execute হচ্ছে। Multiple cores লাগে।
- সহজ কথায়: Concurrency হলো "dealing with" আর Parallelism হলো "doing"

### Q5: User-level thread এবং Kernel-level thread এর পার্থক্য?

**Answer:**

| User-level | Kernel-level |
|------------|--------------|
| User space এ manage | OS দ্বারা manage |
| Fast context switching | Slow context switching |
| OS জানে না | OS জানে |
| True parallelism না | True parallelism support |
| Blocking call problem | Better I/O handling |

### Q6: কখন context switching ঘটে?

**Answer:**
- Time quantum শেষ হলে
- Process I/O request করলে
- Interrupt আসলে (hardware/software)
- Higher priority process ready হলে
- System call করলে

### Q7: Multithreading এর advantages কী?

**Answer:**
- **Responsiveness:** UI responsive থাকে
- **Resource sharing:** Memory এবং resources share করে
- **Economy:** Process থেকে সস্তা
- **Scalability:** Multi-core system এ better performance
- **Better throughput:** Multiple tasks parallel এ

### Q8: Race condition কী এবং কিভাবে solve করবেন?

**Answer:**
**Race Condition:** যখন একাধিক threads/processes shared resource access করে এবং execution order এর উপর result depend করে।

**Solutions:**
- Mutex locks
- Semaphores
- Monitors
- Atomic operations
- Critical section protection

### Q9: Deadlock কী এবং কিভাবে prevent করবেন?

**Answer:**
**Deadlock:** যখন processes একে অপরের resource এর জন্য অসীমভাবে wait করে।

**4টি necessary conditions:**
1. Mutual Exclusion
2. Hold and Wait
3. No Preemption
4. Circular Wait

**Prevention:** যেকোনো একটি condition break করলেই deadlock prevent হবে।

### Q10: Thread-safe code মানে কী?

**Answer:**
যে code একাধিক threads থেকে simultaneously call করা যায় কোন race condition বা data corruption ছাড়াই। এর জন্য:
- Shared data protect করতে হয়
- Proper synchronization লাগে
- Atomic operations use করতে হয়
- Immutable data structures ব্যবহার করা ভালো

---

## Quick Revision Points

### Process
✓ Program in execution  
✓ Isolated memory  
✓ 5 states: New, Ready, Running, Waiting, Terminated

### Thread
✓ Lightweight process  
✓ Shared memory within process  
✓ Faster creation and switching

### PCB
✓ Process information storage  
✓ Essential for context switching  
✓ Contains PID, state, registers, etc.

### Context Switching
✓ Save current → Load next  
✓ Overhead but necessary  
✓ Happens on interrupt, time slice, I/O

### Concurrency
✓ Dealing with multiple tasks  
✓ Single core এও possible  
✓ Interleaved execution

### Parallelism
✓ Doing multiple tasks simultaneously  
✓ Needs multiple cores  
✓ True parallel execution

---

## Additional Resources

- **Books:** Operating System Concepts (Silberschatz)
- **Practice:** LeetCode concurrency problems
- **Videos:** MIT OpenCourseWare - Operating Systems

 