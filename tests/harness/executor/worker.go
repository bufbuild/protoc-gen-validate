package main

import (
	"bytes"
	"context"
	"log"
	"reflect"
	"strings"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	harness "github.com/envoyproxy/protoc-gen-validate/tests/harness/go"
)

type SupportedLanguages = []string
type FeatureSet struct {
	errorRules SupportedLanguages
}

// a map from features to harness names for features that may not yet have support across all languages
var featureSupport = &FeatureSet{
	errorRules: []string{"go"}, // error_rules indicates support for a machine-readable `rule` in addition to the human-readable `reason`
}

func Work(wg *sync.WaitGroup, in <-chan TestCase, out chan<- TestResult, harnesses []Harness) {
	for tc := range in {
		execTestCase(tc, harnesses, out)
	}
	wg.Done()
}

func execTestCase(tc TestCase, harnesses []Harness, out chan<- TestResult) {
	tcMessage, err := anypb.New(tc.Message)
	if err != nil {
		log.Printf("unable to convert test case %q to Any - %v", tc.Name, err)
		out <- TestResult{false, false}
		return
	}

	b, err := proto.Marshal(&harness.TestCase{Message: tcMessage})
	if err != nil {
		log.Printf("unable to marshal test case %q - %v", tc.Name, err)
		out <- TestResult{false, false}
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
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
				log.Printf("[%s] (%s harness) internal harness error: %s", tc.Name, h.Name, res.Reasons)
				out <- TestResult{false, false}
				return
			}

			// Backwards compatibility for languages with no multi-error
			// feature: check results of validation in "fail fast" mode only
			if !res.CheckMultipleErrors {
				tcValid := tc.Failures == 0
				if res.Valid != tcValid {
					if res.AllowFailure {
						log.Printf("[%s] (%s harness) ignoring test failure: %v", tc.Name, h.Name, res.Reasons)
						out <- TestResult{false, true}
					} else if tcValid {
						log.Printf("[%s] (%s harness) expected valid, got invalid: %v", tc.Name, h.Name, res.Reasons)
						out <- TestResult{false, false}
					} else {
						log.Printf("[%s] (%s harness) expected invalid, got valid: %v", tc.Name, h.Name, res.Reasons)
						out <- TestResult{false, false}
					}
				} else {
					out <- TestResult{true, false}
				}
				return
			}

			// Check results of validation in "extensive" mode
			if len(res.Reasons) != tc.Failures {
				if res.AllowFailure {
					log.Printf("[%s] (%s harness) ignoring bad number of failures: %v", tc.Name, h.Name, res.Reasons)
					out <- TestResult{false, true}
				} else {
					log.Printf("[%s] (%s harness) expected %d failures, got %d:\n %v", tc.Name, h.Name, tc.Failures, len(res.Reasons), strings.Join(res.Reasons, "\n "))
					out <- TestResult{false, false}
				}
				return
			}

			// test for failing rules being accurate
			for _, lang := range featureSupport.errorRules {
				if lang != h.Name {
					continue
				}

				if !reflect.DeepEqual(res.Rules, tc.ExpectedRules) {
					log.Printf("[%s] (%s harness) expected rules to be %v, got %v", tc.Name, h.Name, tc.ExpectedRules, res.Rules)
					out <- TestResult{false, false}
				}
			}

			out <- TestResult{true, false}
		}()
	}

	wg.Wait()
	return
}
