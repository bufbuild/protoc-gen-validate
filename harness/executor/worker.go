package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/lyft/protoc-gen-validate/harness"
)

func Work(wg *sync.WaitGroup, in <-chan TestCase, out chan<- bool) {
	for tc := range in {
		out <- execTestCase(tc)
	}
	wg.Done()
}

func execTestCase(tc TestCase) (ok bool) {
	any, err := ptypes.MarshalAny(tc.Message)
	if err != nil {
		log.Printf("unable to convert test case %q to Any - %v", tc.Name, err)
		return false
	}

	b, err := proto.Marshal(&harness.TestCase{Message: any})
	if err != nil {
		log.Printf("unable to marshal test case %q - %v", tc.Name, err)
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	wg := new(sync.WaitGroup)
	wg.Add(len(Harnesses))

	errs := make(chan error, len(Harnesses))

	for _, h := range Harnesses {
		fn := h
		go func() {
			defer wg.Done()

			res, err := fn(ctx, bytes.NewReader(b))
			if err != nil {
				errs <- err
				return
			}

			if res.Error {
				errs <- fmt.Errorf("internal harness error: %s", res.Reason)
			} else if res.Valid != tc.Valid {
				if tc.Valid {
					errs <- fmt.Errorf("expected valid, got: %s", res.Reason)
				} else {
					errs <- errors.New("expected invalid, but got valid")
				}
			}
		}()
	}

	wg.Wait()
	close(errs)

	ok = true

	for err := range errs {
		log.Printf("[%s] %v", tc.Name, err)
		ok = false
	}

	return ok
}
