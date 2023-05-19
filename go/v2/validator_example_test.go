// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protovalidate

import (
	"errors"
	"fmt"
	"log"

	pb "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/tests/example/v1"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func Example() {
	validator, err := New()
	if err != nil {
		log.Fatal(err)
	}

	person := &pb.Person{
		Id:    1234,
		Email: "protovalidate@buf.build",
		Name:  "Buf Build",
		Home: &pb.Coordinates{
			Lat: 27.380583333333334,
			Lng: 33.631838888888886,
		},
	}

	err = validator.Validate(person)
	fmt.Println("valid:", err)

	person.Email = "not an email"
	err = validator.Validate(person)
	fmt.Println("invalid:", err)

	// output:
	// valid: <nil>
	// invalid: validation error:
	//  - email: value must be a valid email address [string.email]
}

func ExampleWithFailFast() {
	loc := &pb.Coordinates{Lat: 999.999, Lng: -999.999}

	validator, err := New()
	if err != nil {
		log.Fatal(err)
	}
	err = validator.Validate(loc)
	fmt.Println("default:", err)

	validator, err = New(WithFailFast(true))
	if err != nil {
		log.Fatal(err)
	}
	err = validator.Validate(loc)
	fmt.Println("fail fast:", err)

	// output:
	// default: validation error:
	//  - lat: value must be greater than or equal to -90 and less than or equal to 90 [double.gte_lte]
	//  - lng: value must be greater than or equal to -180 and less than or equal to 180 [double.gte_lte]
	// fail fast: validation error:
	//  - lat: value must be greater than or equal to -90 and less than or equal to 90 [double.gte_lte]
}

func ExampleWithMessages() {
	validator, err := New(
		WithMessages(&pb.Person{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	person := &pb.Person{
		Id:    1234,
		Email: "protovalidate@buf.build",
		Name:  "Protocol Buffer",
	}
	err = validator.Validate(person)
	fmt.Println(err)

	// output: <nil>
}

func ExampleWithDescriptors() {
	pbType, err := protoregistry.GlobalTypes.FindMessageByName("tests.example.v1.Person")
	if err != nil {
		log.Fatal(err)
	}

	validator, err := New(
		WithDescriptors(
			pbType.Descriptor(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	person := &pb.Person{
		Id:    1234,
		Email: "protovalidate@buf.build",
		Name:  "Protocol Buffer",
	}
	err = validator.Validate(person)
	fmt.Println(err)

	// output: <nil>
}

func ExampleWithDisableLazy() {
	person := &pb.Person{
		Id:    1234,
		Email: "protovalidate@buf.build",
		Name:  "Buf Build",
		Home: &pb.Coordinates{
			Lat: 27.380583333333334,
			Lng: 33.631838888888886,
		},
	}

	validator, err := New(
		WithMessages(&pb.Coordinates{}),
		WithDisableLazy(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = validator.Validate(person.Home)
	fmt.Println("person.Home:", err)
	err = validator.Validate(person)
	fmt.Println("person:", err)

	// output:
	// person.Home: <nil>
	// person: compilation error: no evaluator available for tests.example.v1.Person
}

func ExampleValidationError() {
	validator, err := New()
	if err != nil {
		log.Fatal(err)
	}

	loc := &pb.Coordinates{Lat: 999.999}
	err = validator.Validate(loc)
	var valErr *ValidationError
	if ok := errors.As(err, &valErr); ok {
		msg := valErr.ToProto()
		fmt.Println(msg.Violations[0].FieldPath, msg.Violations[0].ConstraintId)
	}

	// output: lat double.gte_lte
}
