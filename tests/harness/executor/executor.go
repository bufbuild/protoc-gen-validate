package main

import (
	"log"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	log.SetFlags(0)
}

func main() {
	start := time.Now()
	successes, failures, skips := run()

	log.Printf("Successes: %d | Failures: %d | Skips: %d (%v)",
		successes, failures, skips, time.Since(start))

	if failures > 0 {
		os.Exit(1)
	}
}

func run() (successes, failures, skips uint64) {
	wg := new(sync.WaitGroup)
	wg.Add(runtime.NumCPU())

	in := make(chan TestCase)
	out := make(chan TestResult)
	done := make(chan struct{})

	for i := 0; i < runtime.NumCPU(); i++ {
		go Work(wg, in, out)
	}

	go func() {
		for res := range out {
			if res.Skipped {
				atomic.AddUint64(&skips, 1)
			} else if res.OK {
				atomic.AddUint64(&successes, 1)
			} else {
				atomic.AddUint64(&failures, 1)
			}
		}
		close(done)
	}()

	for _, test := range TestCases {
		in <- test
	}
	close(in)

	wg.Wait()
	close(out)
	<-done

	return
}
