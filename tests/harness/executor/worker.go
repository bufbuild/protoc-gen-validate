package main

import (
	"bytes"
	"context"
	"log"
	"sync"
	"time"

	harness "github.com/envoyproxy/protoc-gen-validate/tests/harness/go"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func Work(wg *sync.WaitGroup, in <-chan TestCase, out chan<- TestResult, harnesses []Harness) {
	for tc := range in {
		execTestCase(tc, harnesses, out)
	}
	wg.Done()
}

func execTestCase(tc TestCase, harnesses []Harness, out chan<- TestResult) {
	any, err := ptypes.MarshalAny(tc.Message)
	if err != nil {
		log.Printf("unable to convert test case %q to Any - %v", tc.Name, err)
		out <- TestResult{false, false}
		return
	}

	b, err := proto.Marshal(&harness.TestCase{Message: any})
	if err != nil {
		log.Printf("unable to marshal test case %q - %v", tc.Name, err)
		out <- TestResult{false, false}
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	wg := new(sync.WaitGroup)
	wg.Add(len(harnesses))


	for _, h := range harnesses {
		h := h
		go func() {
			defer wg.Done()

			res, err := h.Exec(ctx, bytes.NewReader(b))
			if err != nil {
				log.Printf("[%s] (%s harness) executor error: %s", tc.Name, h.Name, err.Error())
				out <- TestResult{false, false}
				return
			}

			if res.Error {
				log.Printf("[%s] (%s harness) internal harness error: %s", tc.Name, h.Name, res.Reason)
				out <- TestResult{false, false}
			} else if res.Valid != tc.Valid {
				if res.AllowFailure {
					log.Printf("[%s] (%s harness) ignoring test failure: %s", tc.Name, h.Name, res.Reason)
					out <- TestResult{false, true}
				} else if tc.Valid {
					log.Printf("[%s] (%s harness) expected valid, got invalid: %s", tc.Name, h.Name, res.Reason)
					out <- TestResult{false, false}
				} else {
					log.Printf("[%s] (%s harness) expected invalid, got valid: %s", tc.Name, h.Name, res.Reason)
					out <- TestResult{false, false}
				}
			} else {
				out <- TestResult{true, false}
			}
		}()
	}

	wg.Wait()
	return
}
