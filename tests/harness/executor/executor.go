package main

import (
	"log"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func main() {
	start := time.Now()
	log.SetFlags(0)

	successes, failures := uint64(0), uint64(0)
	wg := new(sync.WaitGroup)
	wg.Add(runtime.NumCPU())

	in := make(chan TestCase)
	out := make(chan bool)
	done := make(chan struct{})

	for i := 0; i < runtime.NumCPU(); i++ {
		go Work(wg, in, out)
	}

	go func() {
		for success := range out {
			if success {
				atomic.AddUint64(&successes, 1)
			} else {
				atomic.AddUint64(&failures, 1)
			}
		}
		close(done)
	}()

	log.Println("loading test cases")

	for _, test := range TestCases {
		in <- test
	}

	close(in)
	wg.Wait()
	close(out)
	<-done

	log.Printf("Successes: %d | Failures: %d (%v)", successes, failures, time.Since(start))

	if failures > 0 {
		os.Exit(1)
	}
}

func TestEverything(test *testing.T) {
	main()
}
