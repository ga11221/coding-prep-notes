/*
Go vs JVM concurrency:

|                       | Go                                  | JVM (Java)                                                 |
|-----------------------|-------------------------------------|------------------------------------------------------------|
| Primitive             | Goroutines (user-space, ~4KB stack) | OS threads (kernel-space, ~1MB stack)                      |
| M:N scheduling        | Yes — M goroutines on N OS threads  | No — 1:1 OS threads (Loom adds virtual threads in Java 21) |
| Communication         | Channels (CSP)                      | Shared memory, locks, java.util.concurrent                  |
| Stack growth          | Dynamic, starts tiny                | Fixed (OS thread stack, typically 1MB+)                    |
| Startup cost          | ~4KB, ~1µs                          | ~1MB, heavy syscall                                        |
| Concurrency model     | Fork-join, pipelines, fan-in/out    | Thread pools, Futures, CompletableFuture, ForkJoinPool     |
| Project Loom          | Built-in from start                 | Virtual threads (Java 21+), same concept as goroutines     |
| Shared memory         | Discouraged — channels preferred    | Standard — synchronized, volatile, Atomic*, ReentrantLock  |

Key difference: Go multiplexes goroutines onto a small thread pool transparently.
Java traditionally relied on thread pools to avoid OS thread exhaustion.
Loom (Java 21+) closes the gap with virtual threads.
*/

package main

import (
	"fmt"
	"sync"
)

// Fork-Join: spawn goroutines ("fork"), wait for all to finish ("join").

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker", id)
}

func forkJoinExample() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait() // join point — blocks until all 10 Done() calls
}

// Pipeline: stage produces, stage transforms, stage consumes.
// Each stage is a goroutine connected by channels. Lazy streaming — O(1) memory per stage.

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out) // signals next stage no more data
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in { // blocks until data or channel closed
			out <- n * n
		}
		close(out)
	}()
	return out
}

func pipelineExample() {
	// gen -> sq -> sq -> main
	for n := range sq(sq(gen(2, 3, 4))) {
		fmt.Println(n) // 4, 9, 16 (squared twice)
	}
}

// Fan-out: multiple goroutines read from the same channel, distributing work.

func fanOut(in <-chan int, workers int) {
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for n := range in {
				fmt.Printf("worker %d: %d\n", id, n)
			}
		}(i)
	}
	wg.Wait()
}

// Fan-in: multiple goroutines write to the same channel, merging streams.

func fanIn(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, c := range chs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(c)
	}
	go func() {
		wg.Wait()  // all merges done
		close(out) // signal consumer
	}()
	return out
}

func fanInOutExample() {
	c1 := gen(1, 2, 3)
	c2 := gen(4, 5, 6)
	merged := fanIn(c1, c2)
	for n := range merged {
		fmt.Println(n)
	}
}
