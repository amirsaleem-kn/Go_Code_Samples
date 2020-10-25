## Go Routines

1. Goroutines are like a lightweight thread.
2. Many goroutines execute within a single OS thread.
3. Why goroutines are faster? Because all the goroutines execute withing a single main thread. So, for OS
there is only one thread to manage. So, there is no context switching or scheduling is done by the OS
4. The context switching between the Goroutines is handled by the Go Runtime Scheduler.
5. Go Runtime Scheduler is like a little OS inside a single OS thread. (SMART!)
6. So, Go Runtime scheduler gives us the concurrency model in which it is mapped to a single Logical Processor.
7. We can run multiple cores and each core can have its own logical processor to manage the concurrency within the thread.

### Memory Interleavings

The order in which the tasks are exeucted by OS is non-deterministic. When you have concurrent code executing at the machine level,
you don't know which instruction of the task is going to be executed at first and in which order. The programmer has to pay good amount of attention while writing the code to avoid possible issue because of memory interleavings.

### Race Condition

Due to the memory interleaving problems, race condition could occur. If two tasks are depedant upon each other's data, and data
is not given in proper order, the Race Condition occurs.

Possible interleavings 

Task 1: Instruction 1 : x = 1
---> Task 2 Instruction 1: print x // prints 1
Task 1: Instruction 2 : x = x + 1

Task 1: Instruction 1 : x = 1
Task 1: Instruction 2 : x = x + 1
---> Task 2 Instruction 1: print x // prints 2

The outcome of the task is depened upon the execution order, this might break the code so we have pushed ourselves in a race condition.

Race condition occurs due to communication

How to avoid race conditions?

1. Threads are largely independant but not completely independant.

## Go Routines

1. A goroutine exits when its code is complete
2. Main is also a goroutine
3. When the main goroutine is complete, all other goroutine exit
4. A goroutine may not complete its execution because main completes first

## Synchronization

1. Using global events whose execution is viewed by all threads, simultaneously.
2. We can achieve synchronization with the help of global events. These events are available for all threads to view

## Sync Wait Groups

1. Sync package contains functions to synchronize between goroutines
2. sync.WaitGroup forces a goroutine to wait for other goroutines
3. WaitGroups uses the logic of counter to maintain the execution order of goroutines
4. For every go routine, the WaitGroup add function increments the counter
5. WaitGroup wait function waits for the goroutine to finish
6. WaitGroup done function marks the completion of a goroutine so that the counter can be decremented
7. WaitGroup waits until the counter becomes 0

## Communcation between threads/goroutines

1. Communcation between goroutines is done with the help of channels
2. Channels are used to transfer the data between goroutines
3. Channels are typed (Data Type must be specified)
4. Use make to create a channel, c := make(chan int)
5. Send data on the channel c <- 3
6. Receive data from the channel x := <- c

## Blocking in channels

1. Unbuffered channels cannot hold data in transit
2. Default is unbuffered. c := make(chan int) is an unbueffered channel
3. Sending blocks until data is received
4. Receiving blocks until data is sent

## Buffered Channels

1. Channels can contain a limited number of objects
2. Capacity is the number of objects it can hold in transit
3. c := make(chan int, 3) // note the third argument to the make function
4. We can do 3 Send/Receives without blocking the channel, as our channel can hold 3 buffered data in transit.
5. Sending only blocks if buffer is full
6. Receiving only blocks if buffer is empty
7. We can continously receive data from channels by iterating using for loops

## Misc

1. Use the select statement to choose the data from one of the channels
2. Select can be used for both receive and send data on a channel
3. You can issue an abort case to handle the abort as well

for {
    select {
        case a = <- inputChannel:
            fmt.println(a)
        case b <- c:
            fmt.println(b)
        case <- abort:
            return // break the infinite loop
    }
}

# Mutex (Correct sharing of data b/w goroutines)

1. To avoid unexpected data update on a shared variable, mutex is helpful.
2. Mutex stands for Mutual Exclusion
3. Code segments in different goroutines which cannot execute concurrently
4. Writing to shared variables should be mutually exclusive
5. sync.mutex uses a binary semaphore
6. Flag up - shared variable is in use
7. Flag down - shared variable is available
8. sync.Lock() to put the flag up, meaning shared variable is in use
9. sync.Unlock() to put the flag down, meaning shared variable is available for use

# Synchronous Initialization

1. The objective is to initialize the variables before they can be available to use to other goroutines
2. The problem is where to initialize and how to initialize for the variables to be available to other goroutines for use
3. One solution is to initialize in the main goroutine which is going to call the other goroutines.
4. The other solution is to use the sync package, and use sync.Once method call, so the variable gets initialised only once
5. sync.once.Do(f)  is the syntax to use signle initialisation
6. you can put this once.Do(f) call in multiple goroutines, the sync package will ensure that this is executed only once

# Deadlock

1. 