package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================
// CONCURRENCY IN GO
//
// Go's concurrency model is built on two ideas:
//   - Goroutines: lightweight threads (thousands can run at once)
//   - Channels: safe pipes for passing data between goroutines
//
// Philosophy: "Don't communicate by sharing memory;
//              share memory by communicating."
// ============================================================

// A simple function we'll run concurrently
func fetchData(source string, delay time.Duration) string {
	time.Sleep(delay) // simulate network delay
	return fmt.Sprintf("data from %s", source)
}

// ============================================================
// PART 1: GOROUTINES
// Launch with "go funcCall()"
// ============================================================

func goroutineDemo() {
	fmt.Println("=== Goroutines ===")

	// Without goroutine — sequential
	start := time.Now()
	r1 := fetchData("DB", 100*time.Millisecond)
	r2 := fetchData("API", 100*time.Millisecond)
	fmt.Println("Sequential:", r1, "|", r2)
	fmt.Println("Time:", time.Since(start).Round(time.Millisecond))
}

// ============================================================
// PART 2: CHANNELS
// Make a channel: ch := make(chan Type)
// Send:          ch <- value
// Receive:       value := <-ch
// ============================================================

func channelDemo() {
	fmt.Println("\n=== Channels ===")

	ch := make(chan string)

	// Launch goroutine, it sends result into channel
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- "result from goroutine"
	}()

	// Main goroutine blocks here until something arrives
	msg := <-ch
	fmt.Println("Received:", msg)

	// --- Buffered channel ---
	// Doesn't block the sender until the buffer is full
	buf := make(chan int, 3)
	buf <- 1
	buf <- 2
	buf <- 3
	fmt.Println("Buffered recv:", <-buf, <-buf, <-buf)

	// --- Range over a channel ---
	jobs := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs) // MUST close to stop range
	for j := range jobs {
		fmt.Println("Job:", j)
	}
}

// ============================================================
// PART 3: WAITGROUP
// Coordinate multiple goroutines — wait for all to finish
// ============================================================

func waitGroupDemo() {
	fmt.Println("\n=== WaitGroup ===")

	var wg sync.WaitGroup
	results := make([]string, 3)

	sources := []string{"DB", "Cache", "API"}
	for i, source := range sources {
		wg.Add(1) // tell WaitGroup: one more goroutine launching
		go func(idx int, src string) {
			defer wg.Done() // signals this goroutine is finished
			time.Sleep(50 * time.Millisecond)
			results[idx] = fmt.Sprintf("data from %s", src)
		}(i, source)
	}

	wg.Wait() // block until all Done() calls
	fmt.Println("All results:", results)
}

// ============================================================
// PART 4: MUTEX — protecting shared data
// Use sync.Mutex when multiple goroutines write to the SAME variable
// ============================================================

func mutexDemo() {
	fmt.Println("\n=== Mutex (safe counter) ===")

	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() // only one goroutine at a time
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Counter (should be 1000):", counter)
}

// ============================================================
// PART 5: SELECT
// Like a switch, but for channels — waits for whichever fires first
// ============================================================

func selectDemo() {
	fmt.Println("\n=== Select ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { time.Sleep(30 * time.Millisecond); ch1 <- "one" }()
	go func() { time.Sleep(10 * time.Millisecond); ch2 <- "two" }()

	// Wait for whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received from ch1:", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch2:", msg)
		}
	}

	// --- Select with timeout ---
	fmt.Println("\n--- Timeout with select ---")
	slow := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		slow <- "slow result"
	}()

	select {
	case res := <-slow:
		fmt.Println("Got:", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timed out!")
	}
}

func main() {
	goroutineDemo()
	channelDemo()
	waitGroupDemo()
	mutexDemo()
	selectDemo()

	// ============================================================
	// QUICK REFERENCE:
	//
	//  go func()                    → launch goroutine
	//  ch := make(chan T)           → unbuffered channel
	//  ch := make(chan T, n)        → buffered channel
	//  ch <- val                    → send
	//  val := <-ch                  → receive
	//  close(ch)                    → close (for range loops)
	//  var wg sync.WaitGroup        → coordinate goroutines
	//  wg.Add(1) / wg.Done() / wg.Wait()
	//  var mu sync.Mutex            → protect shared state
	//  mu.Lock() / mu.Unlock()
	//  select { case <-ch: ... }    → multi-channel wait
	//  time.After(d)                → timeout channel
	// ============================================================
}
