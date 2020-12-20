package main

import (
	"math"

	"time"

	cases "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go"
	other_package "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/go"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type TestType int

const (
	TestTypeValidate  TestType = 0
	TestTypeAllErrors TestType = 1
)

type TestCase struct {
	Name    string
	TestType TestType
	Message proto.Message
	Valid   bool
	ErrorCount int
}

type TestResult struct {
	OK, Skipped bool
}

var TestCases []TestCase

func init() {
	sets := [][]TestCase{
		floatCases,
		doubleCases,
		int32Cases,
		int64Cases,
		uint32Cases,
		uint64Cases,
		sint32Cases,
		sint64Cases,
		fixed32Cases,
		fixed64Cases,
		sfixed32Cases,
		sfixed64Cases,
		boolCases,
		stringCases,
		bytesCases,
		enumCases,
		messageCases,
		repeatedCases,
		mapCases,
		oneofCases,
		wrapperCases,
		durationCases,
		timestampCases,
		anyCases,
		kitchenSink,
	}

	for _, set := range sets {
		TestCases = append(TestCases, set...)
	}
}

var floatCases = []TestCase{
	{"float - none - valid", TestTypeValidate, &cases.FloatNone{Val: -1.23456}, true, 0},

	{"float - const - valid", TestTypeValidate, &cases.FloatConst{Val: 1.23}, true, 0},
	{"float - const - invalid", TestTypeValidate, &cases.FloatConst{Val: 4.56}, false, 1},

	{"float - in - valid", TestTypeValidate, &cases.FloatIn{Val: 7.89}, true, 0},
	{"float - in - invalid", TestTypeValidate, &cases.FloatIn{Val: 10.11}, false, 1},

	{"float - not in - valid", TestTypeValidate, &cases.FloatNotIn{Val: 1}, true, 0},
	{"float - not in - invalid", TestTypeValidate, &cases.FloatNotIn{Val: 0}, false, 1},

	{"float - lt - valid", TestTypeValidate, &cases.FloatLT{Val: -1}, true, 0},
	{"float - lt - invalid (equal)", TestTypeValidate, &cases.FloatLT{Val: 0}, false, 1},
	{"float - lt - invalid", TestTypeValidate, &cases.FloatLT{Val: 1}, false, 1},

	{"float - lte - valid", TestTypeValidate, &cases.FloatLTE{Val: 63}, true, 0},
	{"float - lte - valid (equal)", TestTypeValidate, &cases.FloatLTE{Val: 64}, true, 0},
	{"float - lte - invalid", TestTypeValidate, &cases.FloatLTE{Val: 65}, false, 1},

	{"float - gt - valid", TestTypeValidate, &cases.FloatGT{Val: 17}, true, 0},
	{"float - gt - invalid (equal)", TestTypeValidate, &cases.FloatGT{Val: 16}, false, 1},
	{"float - gt - invalid", TestTypeValidate, &cases.FloatGT{Val: 15}, false, 1},

	{"float - gte - valid", TestTypeValidate, &cases.FloatGTE{Val: 9}, true, 0},
	{"float - gte - valid (equal)", TestTypeValidate, &cases.FloatGTE{Val: 8}, true, 0},
	{"float - gte - invalid", TestTypeValidate, &cases.FloatGTE{Val: 7}, false, 1},

	{"float - gt & lt - valid", TestTypeValidate, &cases.FloatGTLT{Val: 5}, true, 0},
	{"float - gt & lt - invalid (above)", TestTypeValidate, &cases.FloatGTLT{Val: 11}, false, 1},
	{"float - gt & lt - invalid (below)", TestTypeValidate, &cases.FloatGTLT{Val: -1}, false, 1},
	{"float - gt & lt - invalid (max)", TestTypeValidate, &cases.FloatGTLT{Val: 10}, false, 1},
	{"float - gt & lt - invalid (min)", TestTypeValidate, &cases.FloatGTLT{Val: 0}, false, 1},

	{"float - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.FloatExLTGT{Val: 11}, true, 0},
	{"float - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.FloatExLTGT{Val: -1}, true, 0},
	{"float - exclusive gt & lt - invalid", TestTypeValidate, &cases.FloatExLTGT{Val: 5}, false, 1},
	{"float - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.FloatExLTGT{Val: 10}, false, 1},
	{"float - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.FloatExLTGT{Val: 0}, false, 1},

	{"float - gte & lte - valid", TestTypeValidate, &cases.FloatGTELTE{Val: 200}, true, 0},
	{"float - gte & lte - valid (max)", TestTypeValidate, &cases.FloatGTELTE{Val: 256}, true, 0},
	{"float - gte & lte - valid (min)", TestTypeValidate, &cases.FloatGTELTE{Val: 128}, true, 0},
	{"float - gte & lte - invalid (above)", TestTypeValidate, &cases.FloatGTELTE{Val: 300}, false, 1},
	{"float - gte & lte - invalid (below)", TestTypeValidate, &cases.FloatGTELTE{Val: 100}, false, 1},

	{"float - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.FloatExGTELTE{Val: 300}, true, 0},
	{"float - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.FloatExGTELTE{Val: 100}, true, 0},
	{"float - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.FloatExGTELTE{Val: 256}, true, 0},
	{"float - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.FloatExGTELTE{Val: 128}, true, 0},
	{"float - exclusive gte & lte - invalid", TestTypeValidate, &cases.FloatExGTELTE{Val: 200}, false, 1},
}

var doubleCases = []TestCase{
	{"double - none - valid", TestTypeValidate, &cases.DoubleNone{Val: -1.23456}, true, 0},

	{"double - const - valid", TestTypeValidate, &cases.DoubleConst{Val: 1.23}, true, 0},
	{"double - const - invalid", TestTypeValidate, &cases.DoubleConst{Val: 4.56}, false, 1},

	{"double - in - valid", TestTypeValidate, &cases.DoubleIn{Val: 7.89}, true, 0},
	{"double - in - invalid", TestTypeValidate, &cases.DoubleIn{Val: 10.11}, false, 1},

	{"double - not in - valid", TestTypeValidate, &cases.DoubleNotIn{Val: 1}, true, 0},
	{"double - not in - invalid", TestTypeValidate, &cases.DoubleNotIn{Val: 0}, false, 1},

	{"double - lt - valid", TestTypeValidate, &cases.DoubleLT{Val: -1}, true, 0},
	{"double - lt - invalid (equal)", TestTypeValidate, &cases.DoubleLT{Val: 0}, false, 1},
	{"double - lt - invalid", TestTypeValidate, &cases.DoubleLT{Val: 1}, false, 1},

	{"double - lte - valid", TestTypeValidate, &cases.DoubleLTE{Val: 63}, true, 0},
	{"double - lte - valid (equal)", TestTypeValidate, &cases.DoubleLTE{Val: 64}, true, 0},
	{"double - lte - invalid", TestTypeValidate, &cases.DoubleLTE{Val: 65}, false, 1},

	{"double - gt - valid", TestTypeValidate, &cases.DoubleGT{Val: 17}, true, 0},
	{"double - gt - invalid (equal)", TestTypeValidate, &cases.DoubleGT{Val: 16}, false, 1},
	{"double - gt - invalid", TestTypeValidate, &cases.DoubleGT{Val: 15}, false, 1},

	{"double - gte - valid", TestTypeValidate, &cases.DoubleGTE{Val: 9}, true, 0},
	{"double - gte - valid (equal)", TestTypeValidate, &cases.DoubleGTE{Val: 8}, true, 0},
	{"double - gte - invalid", TestTypeValidate, &cases.DoubleGTE{Val: 7}, false, 1},

	{"double - gt & lt - valid", TestTypeValidate, &cases.DoubleGTLT{Val: 5}, true, 0},
	{"double - gt & lt - invalid (above)", TestTypeValidate, &cases.DoubleGTLT{Val: 11}, false, 1},
	{"double - gt & lt - invalid (below)", TestTypeValidate, &cases.DoubleGTLT{Val: -1}, false, 1},
	{"double - gt & lt - invalid (max)", TestTypeValidate, &cases.DoubleGTLT{Val: 10}, false, 1},
	{"double - gt & lt - invalid (min)", TestTypeValidate, &cases.DoubleGTLT{Val: 0}, false, 1},

	{"double - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.DoubleExLTGT{Val: 11}, true, 0},
	{"double - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.DoubleExLTGT{Val: -1}, true, 0},
	{"double - exclusive gt & lt - invalid", TestTypeValidate, &cases.DoubleExLTGT{Val: 5}, false, 1},
	{"double - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.DoubleExLTGT{Val: 10}, false, 1},
	{"double - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.DoubleExLTGT{Val: 0}, false, 1},

	{"double - gte & lte - valid", TestTypeValidate, &cases.DoubleGTELTE{Val: 200}, true, 0},
	{"double - gte & lte - valid (max)", TestTypeValidate, &cases.DoubleGTELTE{Val: 256}, true, 0},
	{"double - gte & lte - valid (min)", TestTypeValidate, &cases.DoubleGTELTE{Val: 128}, true, 0},
	{"double - gte & lte - invalid (above)", TestTypeValidate, &cases.DoubleGTELTE{Val: 300}, false, 1},
	{"double - gte & lte - invalid (below)", TestTypeValidate, &cases.DoubleGTELTE{Val: 100}, false, 1},

	{"double - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.DoubleExGTELTE{Val: 300}, true, 0},
	{"double - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.DoubleExGTELTE{Val: 100}, true, 0},
	{"double - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.DoubleExGTELTE{Val: 256}, true, 0},
	{"double - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.DoubleExGTELTE{Val: 128}, true, 0},
	{"double - exclusive gte & lte - invalid", TestTypeValidate, &cases.DoubleExGTELTE{Val: 200}, false, 1},
}

var int32Cases = []TestCase{
	{"int32 - none - valid", TestTypeValidate, &cases.Int32None{Val: 123}, true, 0},

	{"int32 - const - valid", TestTypeValidate, &cases.Int32Const{Val: 1}, true, 0},
	{"int32 - const - invalid", TestTypeValidate, &cases.Int32Const{Val: 2}, false, 1},

	{"int32 - in - valid", TestTypeValidate, &cases.Int32In{Val: 3}, true, 0},
	{"int32 - in - invalid", TestTypeValidate, &cases.Int32In{Val: 5}, false, 1},

	{"int32 - not in - valid", TestTypeValidate, &cases.Int32NotIn{Val: 1}, true, 0},
	{"int32 - not in - invalid", TestTypeValidate, &cases.Int32NotIn{Val: 0}, false, 1},

	{"int32 - lt - valid", TestTypeValidate, &cases.Int32LT{Val: -1}, true, 0},
	{"int32 - lt - invalid (equal)", TestTypeValidate, &cases.Int32LT{Val: 0}, false, 1},
	{"int32 - lt - invalid", TestTypeValidate, &cases.Int32LT{Val: 1}, false, 1},

	{"int32 - lte - valid", TestTypeValidate, &cases.Int32LTE{Val: 63}, true, 0},
	{"int32 - lte - valid (equal)", TestTypeValidate, &cases.Int32LTE{Val: 64}, true, 0},
	{"int32 - lte - invalid", TestTypeValidate, &cases.Int32LTE{Val: 65}, false, 1},

	{"int32 - gt - valid", TestTypeValidate, &cases.Int32GT{Val: 17}, true, 0},
	{"int32 - gt - invalid (equal)", TestTypeValidate, &cases.Int32GT{Val: 16}, false, 1},
	{"int32 - gt - invalid", TestTypeValidate, &cases.Int32GT{Val: 15}, false, 1},

	{"int32 - gte - valid", TestTypeValidate, &cases.Int32GTE{Val: 9}, true, 0},
	{"int32 - gte - valid (equal)", TestTypeValidate, &cases.Int32GTE{Val: 8}, true, 0},
	{"int32 - gte - invalid", TestTypeValidate, &cases.Int32GTE{Val: 7}, false, 1},

	{"int32 - gt & lt - valid", TestTypeValidate, &cases.Int32GTLT{Val: 5}, true, 0},
	{"int32 - gt & lt - invalid (above)", TestTypeValidate, &cases.Int32GTLT{Val: 11}, false, 1},
	{"int32 - gt & lt - invalid (below)", TestTypeValidate, &cases.Int32GTLT{Val: -1}, false, 1},
	{"int32 - gt & lt - invalid (max)", TestTypeValidate, &cases.Int32GTLT{Val: 10}, false, 1},
	{"int32 - gt & lt - invalid (min)", TestTypeValidate, &cases.Int32GTLT{Val: 0}, false, 1},

	{"int32 - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.Int32ExLTGT{Val: 11}, true, 0},
	{"int32 - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.Int32ExLTGT{Val: -1}, true, 0},
	{"int32 - exclusive gt & lt - invalid", TestTypeValidate, &cases.Int32ExLTGT{Val: 5}, false, 1},
	{"int32 - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.Int32ExLTGT{Val: 10}, false, 1},
	{"int32 - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.Int32ExLTGT{Val: 0}, false, 1},

	{"int32 - gte & lte - valid", TestTypeValidate, &cases.Int32GTELTE{Val: 200}, true, 0},
	{"int32 - gte & lte - valid (max)", TestTypeValidate, &cases.Int32GTELTE{Val: 256}, true, 0},
	{"int32 - gte & lte - valid (min)", TestTypeValidate, &cases.Int32GTELTE{Val: 128}, true, 0},
	{"int32 - gte & lte - invalid (above)", TestTypeValidate, &cases.Int32GTELTE{Val: 300}, false, 1},
	{"int32 - gte & lte - invalid (below)", TestTypeValidate, &cases.Int32GTELTE{Val: 100}, false, 1},

	{"int32 - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.Int32ExGTELTE{Val: 300}, true, 0},
	{"int32 - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.Int32ExGTELTE{Val: 100}, true, 0},
	{"int32 - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.Int32ExGTELTE{Val: 256}, true, 0},
	{"int32 - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.Int32ExGTELTE{Val: 128}, true, 0},
	{"int32 - exclusive gte & lte - invalid", TestTypeValidate, &cases.Int32ExGTELTE{Val: 200}, false, 1},
}

var int64Cases = []TestCase{
	{"int64 - none - valid", TestTypeValidate, &cases.Int64None{Val: 123}, true, 0},

	{"int64 - const - valid", TestTypeValidate, &cases.Int64Const{Val: 1}, true, 0},
	{"int64 - const - invalid", TestTypeValidate, &cases.Int64Const{Val: 2}, false, 1},

	{"int64 - in - valid", TestTypeValidate, &cases.Int64In{Val: 3}, true, 0},
	{"int64 - in - invalid", TestTypeValidate, &cases.Int64In{Val: 5}, false, 1},

	{"int64 - not in - valid", TestTypeValidate, &cases.Int64NotIn{Val: 1}, true, 0},
	{"int64 - not in - invalid", TestTypeValidate, &cases.Int64NotIn{Val: 0}, false, 1},

	{"int64 - lt - valid", TestTypeValidate, &cases.Int64LT{Val: -1}, true, 0},
	{"int64 - lt - invalid (equal)", TestTypeValidate, &cases.Int64LT{Val: 0}, false, 1},
	{"int64 - lt - invalid", TestTypeValidate, &cases.Int64LT{Val: 1}, false, 1},

	{"int64 - lte - valid", TestTypeValidate, &cases.Int64LTE{Val: 63}, true, 0},
	{"int64 - lte - valid (equal)", TestTypeValidate, &cases.Int64LTE{Val: 64}, true, 0},
	{"int64 - lte - invalid", TestTypeValidate, &cases.Int64LTE{Val: 65}, false, 1},

	{"int64 - gt - valid", TestTypeValidate, &cases.Int64GT{Val: 17}, true, 0},
	{"int64 - gt - invalid (equal)", TestTypeValidate, &cases.Int64GT{Val: 16}, false, 1},
	{"int64 - gt - invalid", TestTypeValidate, &cases.Int64GT{Val: 15}, false, 1},

	{"int64 - gte - valid", TestTypeValidate, &cases.Int64GTE{Val: 9}, true, 0},
	{"int64 - gte - valid (equal)", TestTypeValidate, &cases.Int64GTE{Val: 8}, true, 0},
	{"int64 - gte - invalid", TestTypeValidate, &cases.Int64GTE{Val: 7}, false, 1},

	{"int64 - gt & lt - valid", TestTypeValidate, &cases.Int64GTLT{Val: 5}, true, 0},
	{"int64 - gt & lt - invalid (above)", TestTypeValidate, &cases.Int64GTLT{Val: 11}, false, 1},
	{"int64 - gt & lt - invalid (below)", TestTypeValidate, &cases.Int64GTLT{Val: -1}, false, 1},
	{"int64 - gt & lt - invalid (max)", TestTypeValidate, &cases.Int64GTLT{Val: 10}, false, 1},
	{"int64 - gt & lt - invalid (min)", TestTypeValidate, &cases.Int64GTLT{Val: 0}, false, 1},

	{"int64 - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.Int64ExLTGT{Val: 11}, true, 0},
	{"int64 - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.Int64ExLTGT{Val: -1}, true, 0},
	{"int64 - exclusive gt & lt - invalid", TestTypeValidate, &cases.Int64ExLTGT{Val: 5}, false, 1},
	{"int64 - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.Int64ExLTGT{Val: 10}, false, 1},
	{"int64 - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.Int64ExLTGT{Val: 0}, false, 1},

	{"int64 - gte & lte - valid", TestTypeValidate, &cases.Int64GTELTE{Val: 200}, true, 0},
	{"int64 - gte & lte - valid (max)", TestTypeValidate, &cases.Int64GTELTE{Val: 256}, true, 0},
	{"int64 - gte & lte - valid (min)", TestTypeValidate, &cases.Int64GTELTE{Val: 128}, true, 0},
	{"int64 - gte & lte - invalid (above)", TestTypeValidate, &cases.Int64GTELTE{Val: 300}, false, 1},
	{"int64 - gte & lte - invalid (below)", TestTypeValidate, &cases.Int64GTELTE{Val: 100}, false, 1},

	{"int64 - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.Int64ExGTELTE{Val: 300}, true, 0},
	{"int64 - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.Int64ExGTELTE{Val: 100}, true, 0},
	{"int64 - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.Int64ExGTELTE{Val: 256}, true, 0},
	{"int64 - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.Int64ExGTELTE{Val: 128}, true, 0},
	{"int64 - exclusive gte & lte - invalid", TestTypeValidate, &cases.Int64ExGTELTE{Val: 200}, false, 1},
}

var uint32Cases = []TestCase{
	{"uint32 - none - valid", TestTypeValidate, &cases.UInt32None{Val: 123}, true, 0},

	{"uint32 - const - valid", TestTypeValidate, &cases.UInt32Const{Val: 1}, true, 0},
	{"uint32 - const - invalid", TestTypeValidate, &cases.UInt32Const{Val: 2}, false, 1},

	{"uint32 - in - valid", TestTypeValidate, &cases.UInt32In{Val: 3}, true, 0},
	{"uint32 - in - invalid", TestTypeValidate, &cases.UInt32In{Val: 5}, false, 1},

	{"uint32 - not in - valid", TestTypeValidate, &cases.UInt32NotIn{Val: 1}, true, 0},
	{"uint32 - not in - invalid", TestTypeValidate, &cases.UInt32NotIn{Val: 0}, false, 1},

	{"uint32 - lt - valid", TestTypeValidate, &cases.UInt32LT{Val: 4}, true, 0},
	{"uint32 - lt - invalid (equal)", TestTypeValidate, &cases.UInt32LT{Val: 5}, false, 1},
	{"uint32 - lt - invalid", TestTypeValidate, &cases.UInt32LT{Val: 6}, false, 1},

	{"uint32 - lte - valid", TestTypeValidate, &cases.UInt32LTE{Val: 63}, true, 0},
	{"uint32 - lte - valid (equal)", TestTypeValidate, &cases.UInt32LTE{Val: 64}, true, 0},
	{"uint32 - lte - invalid", TestTypeValidate, &cases.UInt32LTE{Val: 65}, false, 1},

	{"uint32 - gt - valid", TestTypeValidate, &cases.UInt32GT{Val: 17}, true, 0},
	{"uint32 - gt - invalid (equal)", TestTypeValidate, &cases.UInt32GT{Val: 16}, false, 1},
	{"uint32 - gt - invalid", TestTypeValidate, &cases.UInt32GT{Val: 15}, false, 1},

	{"uint32 - gte - valid", TestTypeValidate, &cases.UInt32GTE{Val: 9}, true, 0},
	{"uint32 - gte - valid (equal)", TestTypeValidate, &cases.UInt32GTE{Val: 8}, true, 0},
	{"uint32 - gte - invalid", TestTypeValidate, &cases.UInt32GTE{Val: 7}, false, 1},

	{"uint32 - gt & lt - valid", TestTypeValidate, &cases.UInt32GTLT{Val: 7}, true, 0},
	{"uint32 - gt & lt - invalid (above)", TestTypeValidate, &cases.UInt32GTLT{Val: 11}, false, 1},
	{"uint32 - gt & lt - invalid (below)", TestTypeValidate, &cases.UInt32GTLT{Val: 1}, false, 1},
	{"uint32 - gt & lt - invalid (max)", TestTypeValidate, &cases.UInt32GTLT{Val: 10}, false, 1},
	{"uint32 - gt & lt - invalid (min)", TestTypeValidate, &cases.UInt32GTLT{Val: 5}, false, 1},

	{"uint32 - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.UInt32ExLTGT{Val: 11}, true, 0},
	{"uint32 - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.UInt32ExLTGT{Val: 4}, true, 0},
	{"uint32 - exclusive gt & lt - invalid", TestTypeValidate, &cases.UInt32ExLTGT{Val: 7}, false, 1},
	{"uint32 - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.UInt32ExLTGT{Val: 10}, false, 1},
	{"uint32 - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.UInt32ExLTGT{Val: 5}, false, 1},

	{"uint32 - gte & lte - valid", TestTypeValidate, &cases.UInt32GTELTE{Val: 200}, true, 0},
	{"uint32 - gte & lte - valid (max)", TestTypeValidate, &cases.UInt32GTELTE{Val: 256}, true, 0},
	{"uint32 - gte & lte - valid (min)", TestTypeValidate, &cases.UInt32GTELTE{Val: 128}, true, 0},
	{"uint32 - gte & lte - invalid (above)", TestTypeValidate, &cases.UInt32GTELTE{Val: 300}, false, 1},
	{"uint32 - gte & lte - invalid (below)", TestTypeValidate, &cases.UInt32GTELTE{Val: 100}, false, 1},

	{"uint32 - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.UInt32ExGTELTE{Val: 300}, true, 0},
	{"uint32 - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.UInt32ExGTELTE{Val: 100}, true, 0},
	{"uint32 - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.UInt32ExGTELTE{Val: 256}, true, 0},
	{"uint32 - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.UInt32ExGTELTE{Val: 128}, true, 0},
	{"uint32 - exclusive gte & lte - invalid", TestTypeValidate, &cases.UInt32ExGTELTE{Val: 200}, false, 1},
}

var uint64Cases = []TestCase{
	{"uint64 - none - valid", TestTypeValidate, &cases.UInt64None{Val: 123}, true, 0},

	{"uint64 - const - valid", TestTypeValidate, &cases.UInt64Const{Val: 1}, true, 0},
	{"uint64 - const - invalid", TestTypeValidate, &cases.UInt64Const{Val: 2}, false, 1},

	{"uint64 - in - valid", TestTypeValidate, &cases.UInt64In{Val: 3}, true, 0},
	{"uint64 - in - invalid", TestTypeValidate, &cases.UInt64In{Val: 5}, false, 1},

	{"uint64 - not in - valid", TestTypeValidate, &cases.UInt64NotIn{Val: 1}, true, 0},
	{"uint64 - not in - invalid", TestTypeValidate, &cases.UInt64NotIn{Val: 0}, false, 1},

	{"uint64 - lt - valid", TestTypeValidate, &cases.UInt64LT{Val: 4}, true, 0},
	{"uint64 - lt - invalid (equal)", TestTypeValidate, &cases.UInt64LT{Val: 5}, false, 1},
	{"uint64 - lt - invalid", TestTypeValidate, &cases.UInt64LT{Val: 6}, false, 1},

	{"uint64 - lte - valid", TestTypeValidate, &cases.UInt64LTE{Val: 63}, true, 0},
	{"uint64 - lte - valid (equal)", TestTypeValidate, &cases.UInt64LTE{Val: 64}, true, 0},
	{"uint64 - lte - invalid", TestTypeValidate, &cases.UInt64LTE{Val: 65}, false, 1},

	{"uint64 - gt - valid", TestTypeValidate, &cases.UInt64GT{Val: 17}, true, 0},
	{"uint64 - gt - invalid (equal)", TestTypeValidate, &cases.UInt64GT{Val: 16}, false, 1},
	{"uint64 - gt - invalid", TestTypeValidate, &cases.UInt64GT{Val: 15}, false, 1},

	{"uint64 - gte - valid", TestTypeValidate, &cases.UInt64GTE{Val: 9}, true, 0},
	{"uint64 - gte - valid (equal)", TestTypeValidate, &cases.UInt64GTE{Val: 8}, true, 0},
	{"uint64 - gte - invalid", TestTypeValidate, &cases.UInt64GTE{Val: 7}, false, 1},

	{"uint64 - gt & lt - valid", TestTypeValidate, &cases.UInt64GTLT{Val: 7}, true, 0},
	{"uint64 - gt & lt - invalid (above)", TestTypeValidate, &cases.UInt64GTLT{Val: 11}, false, 1},
	{"uint64 - gt & lt - invalid (below)", TestTypeValidate, &cases.UInt64GTLT{Val: 1}, false, 1},
	{"uint64 - gt & lt - invalid (max)", TestTypeValidate, &cases.UInt64GTLT{Val: 10}, false, 1},
	{"uint64 - gt & lt - invalid (min)", TestTypeValidate, &cases.UInt64GTLT{Val: 5}, false, 1},

	{"uint64 - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.UInt64ExLTGT{Val: 11}, true, 0},
	{"uint64 - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.UInt64ExLTGT{Val: 4}, true, 0},
	{"uint64 - exclusive gt & lt - invalid", TestTypeValidate, &cases.UInt64ExLTGT{Val: 7}, false, 1},
	{"uint64 - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.UInt64ExLTGT{Val: 10}, false, 1},
	{"uint64 - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.UInt64ExLTGT{Val: 5}, false, 1},

	{"uint64 - gte & lte - valid", TestTypeValidate, &cases.UInt64GTELTE{Val: 200}, true, 0},
	{"uint64 - gte & lte - valid (max)", TestTypeValidate, &cases.UInt64GTELTE{Val: 256}, true, 0},
	{"uint64 - gte & lte - valid (min)", TestTypeValidate, &cases.UInt64GTELTE{Val: 128}, true, 0},
	{"uint64 - gte & lte - invalid (above)", TestTypeValidate, &cases.UInt64GTELTE{Val: 300}, false, 1},
	{"uint64 - gte & lte - invalid (below)", TestTypeValidate, &cases.UInt64GTELTE{Val: 100}, false, 1},

	{"uint64 - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.UInt64ExGTELTE{Val: 300}, true, 0},
	{"uint64 - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.UInt64ExGTELTE{Val: 100}, true, 0},
	{"uint64 - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.UInt64ExGTELTE{Val: 256}, true, 0},
	{"uint64 - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.UInt64ExGTELTE{Val: 128}, true, 0},
	{"uint64 - exclusive gte & lte - invalid", TestTypeValidate, &cases.UInt64ExGTELTE{Val: 200}, false, 1},
}

var sint32Cases = []TestCase{
	{"sint32 - none - valid", TestTypeValidate, &cases.SInt32None{Val: 123}, true, 0},

	{"sint32 - const - valid", TestTypeValidate, &cases.SInt32Const{Val: 1}, true, 0},
	{"sint32 - const - invalid", TestTypeValidate, &cases.SInt32Const{Val: 2}, false, 1},

	{"sint32 - in - valid", TestTypeValidate, &cases.SInt32In{Val: 3}, true, 0},
	{"sint32 - in - invalid", TestTypeValidate, &cases.SInt32In{Val: 5}, false, 1},

	{"sint32 - not in - valid", TestTypeValidate, &cases.SInt32NotIn{Val: 1}, true, 0},
	{"sint32 - not in - invalid", TestTypeValidate, &cases.SInt32NotIn{Val: 0}, false, 1},

	{"sint32 - lt - valid", TestTypeValidate, &cases.SInt32LT{Val: -1}, true, 0},
	{"sint32 - lt - invalid (equal)", TestTypeValidate, &cases.SInt32LT{Val: 0}, false, 1},
	{"sint32 - lt - invalid", TestTypeValidate, &cases.SInt32LT{Val: 1}, false, 1},

	{"sint32 - lte - valid", TestTypeValidate, &cases.SInt32LTE{Val: 63}, true, 0},
	{"sint32 - lte - valid (equal)", TestTypeValidate, &cases.SInt32LTE{Val: 64}, true, 0},
	{"sint32 - lte - invalid", TestTypeValidate, &cases.SInt32LTE{Val: 65}, false, 1},

	{"sint32 - gt - valid", TestTypeValidate, &cases.SInt32GT{Val: 17}, true, 0},
	{"sint32 - gt - invalid (equal)", TestTypeValidate, &cases.SInt32GT{Val: 16}, false, 1},
	{"sint32 - gt - invalid", TestTypeValidate, &cases.SInt32GT{Val: 15}, false, 1},

	{"sint32 - gte - valid", TestTypeValidate, &cases.SInt32GTE{Val: 9}, true, 0},
	{"sint32 - gte - valid (equal)", TestTypeValidate, &cases.SInt32GTE{Val: 8}, true, 0},
	{"sint32 - gte - invalid", TestTypeValidate, &cases.SInt32GTE{Val: 7}, false, 1},

	{"sint32 - gt & lt - valid", TestTypeValidate, &cases.SInt32GTLT{Val: 5}, true, 0},
	{"sint32 - gt & lt - invalid (above)", TestTypeValidate, &cases.SInt32GTLT{Val: 11}, false, 1},
	{"sint32 - gt & lt - invalid (below)", TestTypeValidate, &cases.SInt32GTLT{Val: -1}, false, 1},
	{"sint32 - gt & lt - invalid (max)", TestTypeValidate, &cases.SInt32GTLT{Val: 10}, false, 1},
	{"sint32 - gt & lt - invalid (min)", TestTypeValidate, &cases.SInt32GTLT{Val: 0}, false, 1},

	{"sint32 - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.SInt32ExLTGT{Val: 11}, true, 0},
	{"sint32 - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.SInt32ExLTGT{Val: -1}, true, 0},
	{"sint32 - exclusive gt & lt - invalid", TestTypeValidate, &cases.SInt32ExLTGT{Val: 5}, false, 1},
	{"sint32 - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.SInt32ExLTGT{Val: 10}, false, 1},
	{"sint32 - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.SInt32ExLTGT{Val: 0}, false, 1},

	{"sint32 - gte & lte - valid", TestTypeValidate, &cases.SInt32GTELTE{Val: 200}, true, 0},
	{"sint32 - gte & lte - valid (max)", TestTypeValidate, &cases.SInt32GTELTE{Val: 256}, true, 0},
	{"sint32 - gte & lte - valid (min)", TestTypeValidate, &cases.SInt32GTELTE{Val: 128}, true, 0},
	{"sint32 - gte & lte - invalid (above)", TestTypeValidate, &cases.SInt32GTELTE{Val: 300}, false, 1},
	{"sint32 - gte & lte - invalid (below)", TestTypeValidate, &cases.SInt32GTELTE{Val: 100}, false, 1},

	{"sint32 - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.SInt32ExGTELTE{Val: 300}, true, 0},
	{"sint32 - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.SInt32ExGTELTE{Val: 100}, true, 0},
	{"sint32 - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.SInt32ExGTELTE{Val: 256}, true, 0},
	{"sint32 - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.SInt32ExGTELTE{Val: 128}, true, 0},
	{"sint32 - exclusive gte & lte - invalid", TestTypeValidate, &cases.SInt32ExGTELTE{Val: 200}, false, 1},
}

var sint64Cases = []TestCase{
	{"sint64 - none - valid", TestTypeValidate, &cases.SInt64None{Val: 123}, true, 0},

	{"sint64 - const - valid", TestTypeValidate, &cases.SInt64Const{Val: 1}, true, 0},
	{"sint64 - const - invalid", TestTypeValidate, &cases.SInt64Const{Val: 2}, false, 1},

	{"sint64 - in - valid", TestTypeValidate, &cases.SInt64In{Val: 3}, true, 0},
	{"sint64 - in - invalid", TestTypeValidate, &cases.SInt64In{Val: 5}, false, 1},

	{"sint64 - not in - valid", TestTypeValidate, &cases.SInt64NotIn{Val: 1}, true, 0},
	{"sint64 - not in - invalid", TestTypeValidate, &cases.SInt64NotIn{Val: 0}, false, 1},

	{"sint64 - lt - valid", TestTypeValidate, &cases.SInt64LT{Val: -1}, true, 0},
	{"sint64 - lt - invalid (equal)", TestTypeValidate, &cases.SInt64LT{Val: 0}, false, 1},
	{"sint64 - lt - invalid", TestTypeValidate, &cases.SInt64LT{Val: 1}, false, 1},

	{"sint64 - lte - valid", TestTypeValidate, &cases.SInt64LTE{Val: 63}, true, 0},
	{"sint64 - lte - valid (equal)", TestTypeValidate, &cases.SInt64LTE{Val: 64}, true, 0},
	{"sint64 - lte - invalid", TestTypeValidate, &cases.SInt64LTE{Val: 65}, false, 1},

	{"sint64 - gt - valid", TestTypeValidate, &cases.SInt64GT{Val: 17}, true, 0},
	{"sint64 - gt - invalid (equal)", TestTypeValidate, &cases.SInt64GT{Val: 16}, false, 1},
	{"sint64 - gt - invalid", TestTypeValidate, &cases.SInt64GT{Val: 15}, false, 1},

	{"sint64 - gte - valid", TestTypeValidate, &cases.SInt64GTE{Val: 9}, true, 0},
	{"sint64 - gte - valid (equal)", TestTypeValidate, &cases.SInt64GTE{Val: 8}, true, 0},
	{"sint64 - gte - invalid", TestTypeValidate, &cases.SInt64GTE{Val: 7}, false, 1},

	{"sint64 - gt & lt - valid", TestTypeValidate, &cases.SInt64GTLT{Val: 5}, true, 0},
	{"sint64 - gt & lt - invalid (above)", TestTypeValidate, &cases.SInt64GTLT{Val: 11}, false, 1},
	{"sint64 - gt & lt - invalid (below)", TestTypeValidate, &cases.SInt64GTLT{Val: -1}, false, 1},
	{"sint64 - gt & lt - invalid (max)", TestTypeValidate, &cases.SInt64GTLT{Val: 10}, false, 1},
	{"sint64 - gt & lt - invalid (min)", TestTypeValidate, &cases.SInt64GTLT{Val: 0}, false, 1},

	{"sint64 - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.SInt64ExLTGT{Val: 11}, true, 0},
	{"sint64 - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.SInt64ExLTGT{Val: -1}, true, 0},
	{"sint64 - exclusive gt & lt - invalid", TestTypeValidate, &cases.SInt64ExLTGT{Val: 5}, false, 1},
	{"sint64 - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.SInt64ExLTGT{Val: 10}, false, 1},
	{"sint64 - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.SInt64ExLTGT{Val: 0}, false, 1},

	{"sint64 - gte & lte - valid", TestTypeValidate, &cases.SInt64GTELTE{Val: 200}, true, 0},
	{"sint64 - gte & lte - valid (max)", TestTypeValidate, &cases.SInt64GTELTE{Val: 256}, true, 0},
	{"sint64 - gte & lte - valid (min)", TestTypeValidate, &cases.SInt64GTELTE{Val: 128}, true, 0},
	{"sint64 - gte & lte - invalid (above)", TestTypeValidate, &cases.SInt64GTELTE{Val: 300}, false, 1},
	{"sint64 - gte & lte - invalid (below)", TestTypeValidate, &cases.SInt64GTELTE{Val: 100}, false, 1},

	{"sint64 - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.SInt64ExGTELTE{Val: 300}, true, 0},
	{"sint64 - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.SInt64ExGTELTE{Val: 100}, true, 0},
	{"sint64 - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.SInt64ExGTELTE{Val: 256}, true, 0},
	{"sint64 - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.SInt64ExGTELTE{Val: 128}, true, 0},
	{"sint64 - exclusive gte & lte - invalid", TestTypeValidate, &cases.SInt64ExGTELTE{Val: 200}, false, 1},
}

var fixed32Cases = []TestCase{
	{"fixed32 - none - valid", TestTypeValidate, &cases.Fixed32None{Val: 123}, true, 0},

	{"fixed32 - const - valid", TestTypeValidate, &cases.Fixed32Const{Val: 1}, true, 0},
	{"fixed32 - const - invalid", TestTypeValidate, &cases.Fixed32Const{Val: 2}, false, 1},

	{"fixed32 - in - valid", TestTypeValidate, &cases.Fixed32In{Val: 3}, true, 0},
	{"fixed32 - in - invalid", TestTypeValidate, &cases.Fixed32In{Val: 5}, false, 1},

	{"fixed32 - not in - valid", TestTypeValidate, &cases.Fixed32NotIn{Val: 1}, true, 0},
	{"fixed32 - not in - invalid", TestTypeValidate, &cases.Fixed32NotIn{Val: 0}, false, 1},

	{"fixed32 - lt - valid", TestTypeValidate, &cases.Fixed32LT{Val: 4}, true, 0},
	{"fixed32 - lt - invalid (equal)", TestTypeValidate, &cases.Fixed32LT{Val: 5}, false, 1},
	{"fixed32 - lt - invalid", TestTypeValidate, &cases.Fixed32LT{Val: 6}, false, 1},

	{"fixed32 - lte - valid", TestTypeValidate, &cases.Fixed32LTE{Val: 63}, true, 0},
	{"fixed32 - lte - valid (equal)", TestTypeValidate, &cases.Fixed32LTE{Val: 64}, true, 0},
	{"fixed32 - lte - invalid", TestTypeValidate, &cases.Fixed32LTE{Val: 65}, false, 1},

	{"fixed32 - gt - valid", TestTypeValidate, &cases.Fixed32GT{Val: 17}, true, 0},
	{"fixed32 - gt - invalid (equal)", TestTypeValidate, &cases.Fixed32GT{Val: 16}, false, 1},
	{"fixed32 - gt - invalid", TestTypeValidate, &cases.Fixed32GT{Val: 15}, false, 1},

	{"fixed32 - gte - valid", TestTypeValidate, &cases.Fixed32GTE{Val: 9}, true, 0},
	{"fixed32 - gte - valid (equal)", TestTypeValidate, &cases.Fixed32GTE{Val: 8}, true, 0},
	{"fixed32 - gte - invalid", TestTypeValidate, &cases.Fixed32GTE{Val: 7}, false, 1},

	{"fixed32 - gt & lt - valid", TestTypeValidate, &cases.Fixed32GTLT{Val: 7}, true, 0},
	{"fixed32 - gt & lt - invalid (above)", TestTypeValidate, &cases.Fixed32GTLT{Val: 11}, false, 1},
	{"fixed32 - gt & lt - invalid (below)", TestTypeValidate, &cases.Fixed32GTLT{Val: 1}, false, 1},
	{"fixed32 - gt & lt - invalid (max)", TestTypeValidate, &cases.Fixed32GTLT{Val: 10}, false, 1},
	{"fixed32 - gt & lt - invalid (min)", TestTypeValidate, &cases.Fixed32GTLT{Val: 5}, false, 1},

	{"fixed32 - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.Fixed32ExLTGT{Val: 11}, true, 0},
	{"fixed32 - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.Fixed32ExLTGT{Val: 4}, true, 0},
	{"fixed32 - exclusive gt & lt - invalid", TestTypeValidate, &cases.Fixed32ExLTGT{Val: 7}, false, 1},
	{"fixed32 - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.Fixed32ExLTGT{Val: 10}, false, 1},
	{"fixed32 - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.Fixed32ExLTGT{Val: 5}, false, 1},

	{"fixed32 - gte & lte - valid", TestTypeValidate, &cases.Fixed32GTELTE{Val: 200}, true, 0},
	{"fixed32 - gte & lte - valid (max)", TestTypeValidate, &cases.Fixed32GTELTE{Val: 256}, true, 0},
	{"fixed32 - gte & lte - valid (min)", TestTypeValidate, &cases.Fixed32GTELTE{Val: 128}, true, 0},
	{"fixed32 - gte & lte - invalid (above)", TestTypeValidate, &cases.Fixed32GTELTE{Val: 300}, false, 1},
	{"fixed32 - gte & lte - invalid (below)", TestTypeValidate, &cases.Fixed32GTELTE{Val: 100}, false, 1},

	{"fixed32 - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.Fixed32ExGTELTE{Val: 300}, true, 0},
	{"fixed32 - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.Fixed32ExGTELTE{Val: 100}, true, 0},
	{"fixed32 - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.Fixed32ExGTELTE{Val: 256}, true, 0},
	{"fixed32 - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.Fixed32ExGTELTE{Val: 128}, true, 0},
	{"fixed32 - exclusive gte & lte - invalid", TestTypeValidate, &cases.Fixed32ExGTELTE{Val: 200}, false, 1},
}

var fixed64Cases = []TestCase{
	{"fixed64 - none - valid", TestTypeValidate, &cases.Fixed64None{Val: 123}, true, 0},

	{"fixed64 - const - valid", TestTypeValidate, &cases.Fixed64Const{Val: 1}, true, 0},
	{"fixed64 - const - invalid", TestTypeValidate, &cases.Fixed64Const{Val: 2}, false, 1},

	{"fixed64 - in - valid", TestTypeValidate, &cases.Fixed64In{Val: 3}, true, 0},
	{"fixed64 - in - invalid", TestTypeValidate, &cases.Fixed64In{Val: 5}, false, 1},

	{"fixed64 - not in - valid", TestTypeValidate, &cases.Fixed64NotIn{Val: 1}, true, 0},
	{"fixed64 - not in - invalid", TestTypeValidate, &cases.Fixed64NotIn{Val: 0}, false, 1},

	{"fixed64 - lt - valid", TestTypeValidate, &cases.Fixed64LT{Val: 4}, true, 0},
	{"fixed64 - lt - invalid (equal)", TestTypeValidate, &cases.Fixed64LT{Val: 5}, false, 1},
	{"fixed64 - lt - invalid", TestTypeValidate, &cases.Fixed64LT{Val: 6}, false, 1},

	{"fixed64 - lte - valid", TestTypeValidate, &cases.Fixed64LTE{Val: 63}, true, 0},
	{"fixed64 - lte - valid (equal)", TestTypeValidate, &cases.Fixed64LTE{Val: 64}, true, 0},
	{"fixed64 - lte - invalid", TestTypeValidate, &cases.Fixed64LTE{Val: 65}, false, 1},

	{"fixed64 - gt - valid", TestTypeValidate, &cases.Fixed64GT{Val: 17}, true, 0},
	{"fixed64 - gt - invalid (equal)", TestTypeValidate, &cases.Fixed64GT{Val: 16}, false, 1},
	{"fixed64 - gt - invalid", TestTypeValidate, &cases.Fixed64GT{Val: 15}, false, 1},

	{"fixed64 - gte - valid", TestTypeValidate, &cases.Fixed64GTE{Val: 9}, true, 0},
	{"fixed64 - gte - valid (equal)", TestTypeValidate, &cases.Fixed64GTE{Val: 8}, true, 0},
	{"fixed64 - gte - invalid", TestTypeValidate, &cases.Fixed64GTE{Val: 7}, false, 1},

	{"fixed64 - gt & lt - valid", TestTypeValidate, &cases.Fixed64GTLT{Val: 7}, true, 0},
	{"fixed64 - gt & lt - invalid (above)", TestTypeValidate, &cases.Fixed64GTLT{Val: 11}, false, 1},
	{"fixed64 - gt & lt - invalid (below)", TestTypeValidate, &cases.Fixed64GTLT{Val: 1}, false, 1},
	{"fixed64 - gt & lt - invalid (max)", TestTypeValidate, &cases.Fixed64GTLT{Val: 10}, false, 1},
	{"fixed64 - gt & lt - invalid (min)", TestTypeValidate, &cases.Fixed64GTLT{Val: 5}, false, 1},

	{"fixed64 - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.Fixed64ExLTGT{Val: 11}, true, 0},
	{"fixed64 - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.Fixed64ExLTGT{Val: 4}, true, 0},
	{"fixed64 - exclusive gt & lt - invalid", TestTypeValidate, &cases.Fixed64ExLTGT{Val: 7}, false, 1},
	{"fixed64 - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.Fixed64ExLTGT{Val: 10}, false, 1},
	{"fixed64 - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.Fixed64ExLTGT{Val: 5}, false, 1},

	{"fixed64 - gte & lte - valid", TestTypeValidate, &cases.Fixed64GTELTE{Val: 200}, true, 0},
	{"fixed64 - gte & lte - valid (max)", TestTypeValidate, &cases.Fixed64GTELTE{Val: 256}, true, 0},
	{"fixed64 - gte & lte - valid (min)", TestTypeValidate, &cases.Fixed64GTELTE{Val: 128}, true, 0},
	{"fixed64 - gte & lte - invalid (above)", TestTypeValidate, &cases.Fixed64GTELTE{Val: 300}, false, 1},
	{"fixed64 - gte & lte - invalid (below)", TestTypeValidate, &cases.Fixed64GTELTE{Val: 100}, false, 1},

	{"fixed64 - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.Fixed64ExGTELTE{Val: 300}, true, 0},
	{"fixed64 - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.Fixed64ExGTELTE{Val: 100}, true, 0},
	{"fixed64 - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.Fixed64ExGTELTE{Val: 256}, true, 0},
	{"fixed64 - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.Fixed64ExGTELTE{Val: 128}, true, 0},
	{"fixed64 - exclusive gte & lte - invalid", TestTypeValidate, &cases.Fixed64ExGTELTE{Val: 200}, false, 1},
}

var sfixed32Cases = []TestCase{
	{"sfixed32 - none - valid", TestTypeValidate, &cases.SFixed32None{Val: 123}, true, 0},

	{"sfixed32 - const - valid", TestTypeValidate, &cases.SFixed32Const{Val: 1}, true, 0},
	{"sfixed32 - const - invalid", TestTypeValidate, &cases.SFixed32Const{Val: 2}, false, 1},

	{"sfixed32 - in - valid", TestTypeValidate, &cases.SFixed32In{Val: 3}, true, 0},
	{"sfixed32 - in - invalid", TestTypeValidate, &cases.SFixed32In{Val: 5}, false, 1},

	{"sfixed32 - not in - valid", TestTypeValidate, &cases.SFixed32NotIn{Val: 1}, true, 0},
	{"sfixed32 - not in - invalid", TestTypeValidate, &cases.SFixed32NotIn{Val: 0}, false, 1},

	{"sfixed32 - lt - valid", TestTypeValidate, &cases.SFixed32LT{Val: -1}, true, 0},
	{"sfixed32 - lt - invalid (equal)", TestTypeValidate, &cases.SFixed32LT{Val: 0}, false, 1},
	{"sfixed32 - lt - invalid", TestTypeValidate, &cases.SFixed32LT{Val: 1}, false, 1},

	{"sfixed32 - lte - valid", TestTypeValidate, &cases.SFixed32LTE{Val: 63}, true, 0},
	{"sfixed32 - lte - valid (equal)", TestTypeValidate, &cases.SFixed32LTE{Val: 64}, true, 0},
	{"sfixed32 - lte - invalid", TestTypeValidate, &cases.SFixed32LTE{Val: 65}, false, 1},

	{"sfixed32 - gt - valid", TestTypeValidate, &cases.SFixed32GT{Val: 17}, true, 0},
	{"sfixed32 - gt - invalid (equal)", TestTypeValidate, &cases.SFixed32GT{Val: 16}, false, 1},
	{"sfixed32 - gt - invalid", TestTypeValidate, &cases.SFixed32GT{Val: 15}, false, 1},

	{"sfixed32 - gte - valid", TestTypeValidate, &cases.SFixed32GTE{Val: 9}, true, 0},
	{"sfixed32 - gte - valid (equal)", TestTypeValidate, &cases.SFixed32GTE{Val: 8}, true, 0},
	{"sfixed32 - gte - invalid", TestTypeValidate, &cases.SFixed32GTE{Val: 7}, false, 1},

	{"sfixed32 - gt & lt - valid", TestTypeValidate, &cases.SFixed32GTLT{Val: 5}, true, 0},
	{"sfixed32 - gt & lt - invalid (above)", TestTypeValidate, &cases.SFixed32GTLT{Val: 11}, false, 1},
	{"sfixed32 - gt & lt - invalid (below)", TestTypeValidate, &cases.SFixed32GTLT{Val: -1}, false, 1},
	{"sfixed32 - gt & lt - invalid (max)", TestTypeValidate, &cases.SFixed32GTLT{Val: 10}, false, 1},
	{"sfixed32 - gt & lt - invalid (min)", TestTypeValidate, &cases.SFixed32GTLT{Val: 0}, false, 1},

	{"sfixed32 - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.SFixed32ExLTGT{Val: 11}, true, 0},
	{"sfixed32 - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.SFixed32ExLTGT{Val: -1}, true, 0},
	{"sfixed32 - exclusive gt & lt - invalid", TestTypeValidate, &cases.SFixed32ExLTGT{Val: 5}, false, 1},
	{"sfixed32 - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.SFixed32ExLTGT{Val: 10}, false, 1},
	{"sfixed32 - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.SFixed32ExLTGT{Val: 0}, false, 1},

	{"sfixed32 - gte & lte - valid", TestTypeValidate, &cases.SFixed32GTELTE{Val: 200}, true, 0},
	{"sfixed32 - gte & lte - valid (max)", TestTypeValidate, &cases.SFixed32GTELTE{Val: 256}, true, 0},
	{"sfixed32 - gte & lte - valid (min)", TestTypeValidate, &cases.SFixed32GTELTE{Val: 128}, true, 0},
	{"sfixed32 - gte & lte - invalid (above)", TestTypeValidate, &cases.SFixed32GTELTE{Val: 300}, false, 1},
	{"sfixed32 - gte & lte - invalid (below)", TestTypeValidate, &cases.SFixed32GTELTE{Val: 100}, false, 1},

	{"sfixed32 - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.SFixed32ExGTELTE{Val: 300}, true, 0},
	{"sfixed32 - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.SFixed32ExGTELTE{Val: 100}, true, 0},
	{"sfixed32 - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.SFixed32ExGTELTE{Val: 256}, true, 0},
	{"sfixed32 - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.SFixed32ExGTELTE{Val: 128}, true, 0},
	{"sfixed32 - exclusive gte & lte - invalid", TestTypeValidate, &cases.SFixed32ExGTELTE{Val: 200}, false, 1},
}

var sfixed64Cases = []TestCase{
	{"sfixed64 - none - valid", TestTypeValidate, &cases.SFixed64None{Val: 123}, true, 0},

	{"sfixed64 - const - valid", TestTypeValidate, &cases.SFixed64Const{Val: 1}, true, 0},
	{"sfixed64 - const - invalid", TestTypeValidate, &cases.SFixed64Const{Val: 2}, false, 1},

	{"sfixed64 - in - valid", TestTypeValidate, &cases.SFixed64In{Val: 3}, true, 0},
	{"sfixed64 - in - invalid", TestTypeValidate, &cases.SFixed64In{Val: 5}, false, 1},

	{"sfixed64 - not in - valid", TestTypeValidate, &cases.SFixed64NotIn{Val: 1}, true, 0},
	{"sfixed64 - not in - invalid", TestTypeValidate, &cases.SFixed64NotIn{Val: 0}, false, 1},

	{"sfixed64 - lt - valid", TestTypeValidate, &cases.SFixed64LT{Val: -1}, true, 0},
	{"sfixed64 - lt - invalid (equal)", TestTypeValidate, &cases.SFixed64LT{Val: 0}, false, 1},
	{"sfixed64 - lt - invalid", TestTypeValidate, &cases.SFixed64LT{Val: 1}, false, 1},

	{"sfixed64 - lte - valid", TestTypeValidate, &cases.SFixed64LTE{Val: 63}, true, 0},
	{"sfixed64 - lte - valid (equal)", TestTypeValidate, &cases.SFixed64LTE{Val: 64}, true, 0},
	{"sfixed64 - lte - invalid", TestTypeValidate, &cases.SFixed64LTE{Val: 65}, false, 1},

	{"sfixed64 - gt - valid", TestTypeValidate, &cases.SFixed64GT{Val: 17}, true, 0},
	{"sfixed64 - gt - invalid (equal)", TestTypeValidate, &cases.SFixed64GT{Val: 16}, false, 1},
	{"sfixed64 - gt - invalid", TestTypeValidate, &cases.SFixed64GT{Val: 15}, false, 1},

	{"sfixed64 - gte - valid", TestTypeValidate, &cases.SFixed64GTE{Val: 9}, true, 0},
	{"sfixed64 - gte - valid (equal)", TestTypeValidate, &cases.SFixed64GTE{Val: 8}, true, 0},
	{"sfixed64 - gte - invalid", TestTypeValidate, &cases.SFixed64GTE{Val: 7}, false, 1},

	{"sfixed64 - gt & lt - valid", TestTypeValidate, &cases.SFixed64GTLT{Val: 5}, true, 0},
	{"sfixed64 - gt & lt - invalid (above)", TestTypeValidate, &cases.SFixed64GTLT{Val: 11}, false, 1},
	{"sfixed64 - gt & lt - invalid (below)", TestTypeValidate, &cases.SFixed64GTLT{Val: -1}, false, 1},
	{"sfixed64 - gt & lt - invalid (max)", TestTypeValidate, &cases.SFixed64GTLT{Val: 10}, false, 1},
	{"sfixed64 - gt & lt - invalid (min)", TestTypeValidate, &cases.SFixed64GTLT{Val: 0}, false, 1},

	{"sfixed64 - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.SFixed64ExLTGT{Val: 11}, true, 0},
	{"sfixed64 - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.SFixed64ExLTGT{Val: -1}, true, 0},
	{"sfixed64 - exclusive gt & lt - invalid", TestTypeValidate, &cases.SFixed64ExLTGT{Val: 5}, false, 1},
	{"sfixed64 - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.SFixed64ExLTGT{Val: 10}, false, 1},
	{"sfixed64 - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.SFixed64ExLTGT{Val: 0}, false, 1},

	{"sfixed64 - gte & lte - valid", TestTypeValidate, &cases.SFixed64GTELTE{Val: 200}, true, 0},
	{"sfixed64 - gte & lte - valid (max)", TestTypeValidate, &cases.SFixed64GTELTE{Val: 256}, true, 0},
	{"sfixed64 - gte & lte - valid (min)", TestTypeValidate, &cases.SFixed64GTELTE{Val: 128}, true, 0},
	{"sfixed64 - gte & lte - invalid (above)", TestTypeValidate, &cases.SFixed64GTELTE{Val: 300}, false, 1},
	{"sfixed64 - gte & lte - invalid (below)", TestTypeValidate, &cases.SFixed64GTELTE{Val: 100}, false, 1},

	{"sfixed64 - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.SFixed64ExGTELTE{Val: 300}, true, 0},
	{"sfixed64 - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.SFixed64ExGTELTE{Val: 100}, true, 0},
	{"sfixed64 - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.SFixed64ExGTELTE{Val: 256}, true, 0},
	{"sfixed64 - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.SFixed64ExGTELTE{Val: 128}, true, 0},
	{"sfixed64 - exclusive gte & lte - invalid", TestTypeValidate, &cases.SFixed64ExGTELTE{Val: 200}, false, 1},
}

var boolCases = []TestCase{
	{"bool - none - valid", TestTypeValidate, &cases.BoolNone{Val: true}, true, 0},
	{"bool - const (true) - valid", TestTypeValidate, &cases.BoolConstTrue{Val: true}, true, 0},
	{"bool - const (true) - invalid", TestTypeValidate, &cases.BoolConstTrue{Val: false}, false, 1},
	{"bool - const (false) - valid", TestTypeValidate, &cases.BoolConstFalse{Val: false}, true, 0},
	{"bool - const (false) - invalid", TestTypeValidate, &cases.BoolConstFalse{Val: true}, false, 1},
}

var stringCases = []TestCase{
	{"string - none - valid", TestTypeValidate, &cases.StringNone{Val: "quux"}, true, 0},

	{"string - const - valid", TestTypeValidate, &cases.StringConst{Val: "foo"}, true, 0},
	{"string - const - invalid", TestTypeValidate, &cases.StringConst{Val: "bar"}, false, 1},

	{"string - in - valid", TestTypeValidate, &cases.StringIn{Val: "bar"}, true, 0},
	{"string - in - invalid", TestTypeValidate, &cases.StringIn{Val: "quux"}, false, 1},
	{"string - not in - valid", TestTypeValidate, &cases.StringNotIn{Val: "quux"}, true, 0},
	{"string - not in - invalid", TestTypeValidate, &cases.StringNotIn{Val: "fizz"}, false, 1},

	{"string - len - valid", TestTypeValidate, &cases.StringLen{Val: "baz"}, true, 0},
	{"string - len - valid (multibyte)", TestTypeValidate, &cases.StringLen{Val: "你好吖"}, true, 0},
	{"string - len - invalid (lt)", TestTypeValidate, &cases.StringLen{Val: "go"}, false, 1},
	{"string - len - invalid (gt)", TestTypeValidate, &cases.StringLen{Val: "fizz"}, false, 1},
	{"string - len - invalid (multibyte)", TestTypeValidate, &cases.StringLen{Val: "你好"}, false, 1},

	{"string - min len - valid", TestTypeValidate, &cases.StringMinLen{Val: "protoc"}, true, 0},
	{"string - min len - valid (min)", TestTypeValidate, &cases.StringMinLen{Val: "baz"}, true, 0},
	{"string - min len - invalid", TestTypeValidate, &cases.StringMinLen{Val: "go"}, false, 1},
	{"string - min len - invalid (multibyte)", TestTypeValidate, &cases.StringMinLen{Val: "你好"}, false, 1},

	{"string - max len - valid", TestTypeValidate, &cases.StringMaxLen{Val: "foo"}, true, 0},
	{"string - max len - valid (max)", TestTypeValidate, &cases.StringMaxLen{Val: "proto"}, true, 0},
	{"string - max len - valid (multibyte)", TestTypeValidate, &cases.StringMaxLen{Val: "你好你好"}, true, 0},
	{"string - max len - invalid", TestTypeValidate, &cases.StringMaxLen{Val: "1234567890"}, false, 1},

	{"string - min/max len - valid", TestTypeValidate, &cases.StringMinMaxLen{Val: "quux"}, true, 0},
	{"string - min/max len - valid (min)", TestTypeValidate, &cases.StringMinMaxLen{Val: "foo"}, true, 0},
	{"string - min/max len - valid (max)", TestTypeValidate, &cases.StringMinMaxLen{Val: "proto"}, true, 0},
	{"string - min/max len - valid (multibyte)", TestTypeValidate, &cases.StringMinMaxLen{Val: "你好你好"}, true, 0},
	{"string - min/max len - invalid (below)", TestTypeValidate, &cases.StringMinMaxLen{Val: "go"}, false, 1},
	{"string - min/max len - invalid (above)", TestTypeValidate, &cases.StringMinMaxLen{Val: "validate"}, false, 1},

	{"string - equal min/max len - valid", TestTypeValidate, &cases.StringEqualMinMaxLen{Val: "proto"}, true, 0},
	{"string - equal min/max len - invalid", TestTypeValidate, &cases.StringEqualMinMaxLen{Val: "validate"}, false, 1},

	{"string - len bytes - valid", TestTypeValidate, &cases.StringLenBytes{Val: "pace"}, true, 0},
	{"string - len bytes - invalid (lt)", TestTypeValidate, &cases.StringLenBytes{Val: "val"}, false, 1},
	{"string - len bytes - invalid (gt)", TestTypeValidate, &cases.StringLenBytes{Val: "world"}, false, 1},
	{"string - len bytes - invalid (multibyte)", TestTypeValidate, &cases.StringLenBytes{Val: "世界和平"}, false, 1},

	{"string - min bytes - valid", TestTypeValidate, &cases.StringMinBytes{Val: "proto"}, true, 0},
	{"string - min bytes - valid (min)", TestTypeValidate, &cases.StringMinBytes{Val: "quux"}, true, 0},
	{"string - min bytes - valid (multibyte)", TestTypeValidate, &cases.StringMinBytes{Val: "你好"}, true, 0},
	{"string - min bytes - invalid", TestTypeValidate, &cases.StringMinBytes{Val: ""}, false, 1},

	{"string - max bytes - valid", TestTypeValidate, &cases.StringMaxBytes{Val: "foo"}, true, 0},
	{"string - max bytes - valid (max)", TestTypeValidate, &cases.StringMaxBytes{Val: "12345678"}, true, 0},
	{"string - max bytes - invalid", TestTypeValidate, &cases.StringMaxBytes{Val: "123456789"}, false, 1},
	{"string - max bytes - invalid (multibyte)", TestTypeValidate, &cases.StringMaxBytes{Val: "你好你好你好"}, false, 1},

	{"string - min/max bytes - valid", TestTypeValidate, &cases.StringMinMaxBytes{Val: "protoc"}, true, 0},
	{"string - min/max bytes - valid (min)", TestTypeValidate, &cases.StringMinMaxBytes{Val: "quux"}, true, 0},
	{"string - min/max bytes - valid (max)", TestTypeValidate, &cases.StringMinMaxBytes{Val: "fizzbuzz"}, true, 0},
	{"string - min/max bytes - valid (multibyte)", TestTypeValidate, &cases.StringMinMaxBytes{Val: "你好"}, true, 0},
	{"string - min/max bytes - invalid (below)", TestTypeValidate, &cases.StringMinMaxBytes{Val: "foo"}, false, 1},
	{"string - min/max bytes - invalid (above)", TestTypeValidate, &cases.StringMinMaxBytes{Val: "你好你好你"}, false, 1},

	{"string - equal min/max bytes - valid", TestTypeValidate, &cases.StringEqualMinMaxBytes{Val: "protoc"}, true, 0},
	{"string - equal min/max bytes - invalid", TestTypeValidate, &cases.StringEqualMinMaxBytes{Val: "foo"}, false, 1},

	{"string - pattern - valid", TestTypeValidate, &cases.StringPattern{Val: "Foo123"}, true, 0},
	{"string - pattern - invalid", TestTypeValidate, &cases.StringPattern{Val: "!@#$%^&*()"}, false, 1},
	{"string - pattern - invalid (empty)", TestTypeValidate, &cases.StringPattern{Val: ""}, false, 1},
	{"string - pattern - invalid (null)", TestTypeValidate, &cases.StringPattern{Val: "a\000"}, false, 1},

	{"string - pattern (escapes) - valid", TestTypeValidate, &cases.StringPatternEscapes{Val: "* \\ x"}, true, 0},
	{"string - pattern (escapes) - invalid", TestTypeValidate, &cases.StringPatternEscapes{Val: "invalid"}, false, 1},
	{"string - pattern (escapes) - invalid (empty)", TestTypeValidate, &cases.StringPatternEscapes{Val: ""}, false, 1},

	{"string - prefix - valid", TestTypeValidate, &cases.StringPrefix{Val: "foobar"}, true, 0},
	{"string - prefix - valid (only)", TestTypeValidate, &cases.StringPrefix{Val: "foo"}, true, 0},
	{"string - prefix - invalid", TestTypeValidate, &cases.StringPrefix{Val: "bar"}, false, 1},
	{"string - prefix - invalid (case-sensitive)", TestTypeValidate, &cases.StringPrefix{Val: "Foobar"}, false, 1},

	{"string - contains - valid", TestTypeValidate, &cases.StringContains{Val: "candy bars"}, true, 0},
	{"string - contains - valid (only)", TestTypeValidate, &cases.StringContains{Val: "bar"}, true, 0},
	{"string - contains - invalid", TestTypeValidate, &cases.StringContains{Val: "candy bazs"}, false, 1},
	{"string - contains - invalid (case-sensitive)", TestTypeValidate, &cases.StringContains{Val: "Candy Bars"}, false, 1},

	{"string - not contains - valid", TestTypeValidate, &cases.StringNotContains{Val: "candy bazs"}, true, 0},
	{"string - not contains - valid (case-sensitive)", TestTypeValidate, &cases.StringNotContains{Val: "Candy Bars"}, true, 0},
	{"string - not contains - invalid", TestTypeValidate, &cases.StringNotContains{Val: "candy bars"}, false, 1},
	{"string - not contains - invalid (equal)", TestTypeValidate, &cases.StringNotContains{Val: "bar"}, false, 1},

	{"string - suffix - valid", TestTypeValidate, &cases.StringSuffix{Val: "foobaz"}, true, 0},
	{"string - suffix - valid (only)", TestTypeValidate, &cases.StringSuffix{Val: "baz"}, true, 0},
	{"string - suffix - invalid", TestTypeValidate, &cases.StringSuffix{Val: "foobar"}, false, 1},
	{"string - suffix - invalid (case-sensitive)", TestTypeValidate, &cases.StringSuffix{Val: "FooBaz"}, false, 1},

	{"string - email - valid", TestTypeValidate, &cases.StringEmail{Val: "foo@bar.com"}, true, 0},
	{"string - email - valid (name)", TestTypeValidate, &cases.StringEmail{Val: "John Smith <foo@bar.com>"}, true, 0},
	{"string - email - invalid", TestTypeValidate, &cases.StringEmail{Val: "foobar"}, false, 1},
	{"string - email - invalid (local segment too long)", TestTypeValidate, &cases.StringEmail{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789@example.com"}, false, 1},
	{"string - email - invalid (hostname too long)", TestTypeValidate, &cases.StringEmail{Val: "foo@x0123456789012345678901234567890123456789012345678901234567890123456789.com"}, false, 1},
	{"string - email - invalid (bad hostname)", TestTypeValidate, &cases.StringEmail{Val: "foo@-bar.com"}, false, 1},
	{"string - email - empty", TestTypeValidate, &cases.StringEmail{Val: ""}, false, 1},

	{"string - address - valid hostname", TestTypeValidate, &cases.StringAddress{Val: "example.com"}, true, 0},
	{"string - address - valid hostname (uppercase)", TestTypeValidate, &cases.StringAddress{Val: "ASD.example.com"}, true, 0},
	{"string - address - valid hostname (hyphens)", TestTypeValidate, &cases.StringAddress{Val: "foo-bar.com"}, true, 0},
	{"string - address - valid hostname (trailing dot)", TestTypeValidate, &cases.StringAddress{Val: "example.com."}, true, 0},
	{"string - address - invalid hostname", TestTypeValidate, &cases.StringAddress{Val: "!@#$%^&"}, false, 1},
	{"string - address - invalid hostname (underscore)", TestTypeValidate, &cases.StringAddress{Val: "foo_bar.com"}, false, 1},
	{"string - address - invalid hostname (too long)", TestTypeValidate, &cases.StringAddress{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789.com"}, false, 1},
	{"string - address - invalid hostname (trailing hyphens)", TestTypeValidate, &cases.StringAddress{Val: "foo-bar-.com"}, false, 1},
	{"string - address - invalid hostname (leading hyphens)", TestTypeValidate, &cases.StringAddress{Val: "foo-bar.-com"}, false, 1},
	{"string - address - invalid hostname (empty)", TestTypeValidate, &cases.StringAddress{Val: "asd..asd.com"}, false, 1},
	{"string - address - invalid hostname (IDNs)", TestTypeValidate, &cases.StringAddress{Val: "你好.com"}, false, 1},
	{"string - address - valid ip (v4)", TestTypeValidate, &cases.StringAddress{Val: "192.168.0.1"}, true, 0},
	{"string - address - valid ip (v6)", TestTypeValidate, &cases.StringAddress{Val: "3e::99"}, true, 0},
	{"string - address - invalid ip", TestTypeValidate, &cases.StringAddress{Val: "ff::fff::0b"}, false, 1},

	{"string - hostname - valid", TestTypeValidate, &cases.StringHostname{Val: "example.com"}, true, 0},
	{"string - hostname - valid (uppercase)", TestTypeValidate, &cases.StringHostname{Val: "ASD.example.com"}, true, 0},
	{"string - hostname - valid (hyphens)", TestTypeValidate, &cases.StringHostname{Val: "foo-bar.com"}, true, 0},
	{"string - hostname - valid (trailing dot)", TestTypeValidate, &cases.StringHostname{Val: "example.com."}, true, 0},
	{"string - hostname - invalid", TestTypeValidate, &cases.StringHostname{Val: "!@#$%^&"}, false, 1},
	{"string - hostname - invalid (underscore)", TestTypeValidate, &cases.StringHostname{Val: "foo_bar.com"}, false, 1},
	{"string - hostname - invalid (too long)", TestTypeValidate, &cases.StringHostname{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789.com"}, false, 1},
	{"string - hostname - invalid (trailing hyphens)", TestTypeValidate, &cases.StringHostname{Val: "foo-bar-.com"}, false, 1},
	{"string - hostname - invalid (leading hyphens)", TestTypeValidate, &cases.StringHostname{Val: "foo-bar.-com"}, false, 1},
	{"string - hostname - invalid (empty)", TestTypeValidate, &cases.StringHostname{Val: "asd..asd.com"}, false, 1},
	{"string - hostname - invalid (IDNs)", TestTypeValidate, &cases.StringHostname{Val: "你好.com"}, false, 1},

	{"string - IP - valid (v4)", TestTypeValidate, &cases.StringIP{Val: "192.168.0.1"}, true, 0},
	{"string - IP - valid (v6)", TestTypeValidate, &cases.StringIP{Val: "3e::99"}, true, 0},
	{"string - IP - invalid", TestTypeValidate, &cases.StringIP{Val: "foobar"}, false, 1},

	{"string - IPv4 - valid", TestTypeValidate, &cases.StringIPv4{Val: "192.168.0.1"}, true, 0},
	{"string - IPv4 - invalid", TestTypeValidate, &cases.StringIPv4{Val: "foobar"}, false, 1},
	{"string - IPv4 - invalid (erroneous)", TestTypeValidate, &cases.StringIPv4{Val: "256.0.0.0"}, false, 1},
	{"string - IPv4 - invalid (v6)", TestTypeValidate, &cases.StringIPv4{Val: "3e::99"}, false, 1},

	{"string - IPv6 - valid", TestTypeValidate, &cases.StringIPv6{Val: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, true, 0},
	{"string - IPv6 - valid (collapsed)", TestTypeValidate, &cases.StringIPv6{Val: "2001:db8:85a3::8a2e:370:7334"}, true, 0},
	{"string - IPv6 - invalid", TestTypeValidate, &cases.StringIPv6{Val: "foobar"}, false, 1},
	{"string - IPv6 - invalid (v4)", TestTypeValidate, &cases.StringIPv6{Val: "192.168.0.1"}, false, 1},
	{"string - IPv6 - invalid (erroneous)", TestTypeValidate, &cases.StringIPv6{Val: "ff::fff::0b"}, false, 1},

	{"string - URI - valid", TestTypeValidate, &cases.StringURI{Val: "http://example.com/foo/bar?baz=quux"}, true, 0},
	{"string - URI - invalid", TestTypeValidate, &cases.StringURI{Val: "!@#$%^&*%$#"}, false, 1},
	{"string - URI - invalid (relative)", TestTypeValidate, &cases.StringURI{Val: "/foo/bar?baz=quux"}, false, 1},

	{"string - URI - valid", TestTypeValidate, &cases.StringURIRef{Val: "http://example.com/foo/bar?baz=quux"}, true, 0},
	{"string - URI - valid (relative)", TestTypeValidate, &cases.StringURIRef{Val: "/foo/bar?baz=quux"}, true, 0},
	{"string - URI - invalid", TestTypeValidate, &cases.StringURIRef{Val: "!@#$%^&*%$#"}, false, 1},

	{"string - UUID - valid (nil)", TestTypeValidate, &cases.StringUUID{Val: "00000000-0000-0000-0000-000000000000"}, true, 0},
	{"string - UUID - valid (v1)", TestTypeValidate, &cases.StringUUID{Val: "b45c0c80-8880-11e9-a5b1-000000000000"}, true, 0},
	{"string - UUID - valid (v1 - case-insensitive)", TestTypeValidate, &cases.StringUUID{Val: "B45C0C80-8880-11E9-A5B1-000000000000"}, true, 0},
	{"string - UUID - valid (v2)", TestTypeValidate, &cases.StringUUID{Val: "b45c0c80-8880-21e9-a5b1-000000000000"}, true, 0},
	{"string - UUID - valid (v2 - case-insensitive)", TestTypeValidate, &cases.StringUUID{Val: "B45C0C80-8880-21E9-A5B1-000000000000"}, true, 0},
	{"string - UUID - valid (v3)", TestTypeValidate, &cases.StringUUID{Val: "a3bb189e-8bf9-3888-9912-ace4e6543002"}, true, 0},
	{"string - UUID - valid (v3 - case-insensitive)", TestTypeValidate, &cases.StringUUID{Val: "A3BB189E-8BF9-3888-9912-ACE4E6543002"}, true, 0},
	{"string - UUID - valid (v4)", TestTypeValidate, &cases.StringUUID{Val: "8b208305-00e8-4460-a440-5e0dcd83bb0a"}, true, 0},
	{"string - UUID - valid (v4 - case-insensitive)", TestTypeValidate, &cases.StringUUID{Val: "8B208305-00E8-4460-A440-5E0DCD83BB0A"}, true, 0},
	{"string - UUID - valid (v5)", TestTypeValidate, &cases.StringUUID{Val: "a6edc906-2f9f-5fb2-a373-efac406f0ef2"}, true, 0},
	{"string - UUID - valid (v5 - case-insensitive)", TestTypeValidate, &cases.StringUUID{Val: "A6EDC906-2F9F-5FB2-A373-EFAC406F0EF2"}, true, 0},
	{"string - UUID - invalid", TestTypeValidate, &cases.StringUUID{Val: "foobar"}, false, 1},
	{"string - UUID - invalid (bad UUID)", TestTypeValidate, &cases.StringUUID{Val: "ffffffff-ffff-ffff-ffff-fffffffffffff"}, false, 1},

	{"string - http header name - valid", TestTypeValidate, &cases.StringHttpHeaderName{Val: "clustername"}, true, 0},
	{"string - http header name - valid", TestTypeValidate, &cases.StringHttpHeaderName{Val: ":path"}, true, 0},
	{"string - http header name - valid (nums)", TestTypeValidate, &cases.StringHttpHeaderName{Val: "cluster-123"}, true, 0},
	{"string - http header name - valid (special token)", TestTypeValidate, &cases.StringHttpHeaderName{Val: "!+#&.%"}, true, 0},
	{"string - http header name - valid (period)", TestTypeValidate, &cases.StringHttpHeaderName{Val: "CLUSTER.NAME"}, true, 0},
	{"string - http header name - invalid", TestTypeValidate, &cases.StringHttpHeaderName{Val: ":"}, false, 1},
	{"string - http header name - invalid", TestTypeValidate, &cases.StringHttpHeaderName{Val: ":path:"}, false, 1},
	{"string - http header name - invalid (space)", TestTypeValidate, &cases.StringHttpHeaderName{Val: "cluster name"}, false, 1},
	{"string - http header name - invalid (return)", TestTypeValidate, &cases.StringHttpHeaderName{Val: "example\r"}, false, 1},
	{"string - http header name - invalid (tab)", TestTypeValidate, &cases.StringHttpHeaderName{Val: "example\t"}, false, 1},
	{"string - http header name - invalid (slash)", TestTypeValidate, &cases.StringHttpHeaderName{Val: "/test/long/url"}, false, 1},

	{"string - http header value - valid", TestTypeValidate, &cases.StringHttpHeaderValue{Val: "cluster.name.123"}, true, 0},
	{"string - http header value - valid (uppercase)", TestTypeValidate, &cases.StringHttpHeaderValue{Val: "/TEST/LONG/URL"}, true, 0},
	{"string - http header value - valid (spaces)", TestTypeValidate, &cases.StringHttpHeaderValue{Val: "cluster name"}, true, 0},
	{"string - http header value - valid (tab)", TestTypeValidate, &cases.StringHttpHeaderValue{Val: "example\t"}, true, 0},
	{"string - http header value - valid (special token)", TestTypeValidate, &cases.StringHttpHeaderValue{Val: "!#%&./+"}, true, 0},
	{"string - http header value - invalid (NUL)", TestTypeValidate, &cases.StringHttpHeaderValue{Val: "foo\u0000bar"}, false, 1},
	{"string - http header value - invalid (DEL)", TestTypeValidate, &cases.StringHttpHeaderValue{Val: "\u007f"}, false, 1},
	{"string - http header value - invalid", TestTypeValidate, &cases.StringHttpHeaderValue{Val: "example\r"}, false, 1},

	{"string - non-strict valid header - valid", TestTypeValidate, &cases.StringValidHeader{Val: "cluster.name.123"}, true, 0},
	{"string - non-strict valid header - valid (uppercase)", TestTypeValidate, &cases.StringValidHeader{Val: "/TEST/LONG/URL"}, true, 0},
	{"string - non-strict valid header - valid (spaces)", TestTypeValidate, &cases.StringValidHeader{Val: "cluster name"}, true, 0},
	{"string - non-strict valid header - valid (tab)", TestTypeValidate, &cases.StringValidHeader{Val: "example\t"}, true, 0},
	{"string - non-strict valid header - valid (DEL)", TestTypeValidate, &cases.StringValidHeader{Val: "\u007f"}, true, 0},
	{"string - non-strict valid header - invalid (NUL)", TestTypeValidate, &cases.StringValidHeader{Val: "foo\u0000bar"}, false, 1},
	{"string - non-strict valid header - invalid (CR)", TestTypeValidate, &cases.StringValidHeader{Val: "example\r"}, false, 1},
	{"string - non-strict valid header - invalid (NL)", TestTypeValidate, &cases.StringValidHeader{Val: "exa\u000Ample"}, false, 1},
}

var bytesCases = []TestCase{
	{"bytes - none - valid", TestTypeValidate, &cases.BytesNone{Val: []byte("quux")}, true, 0},

	{"bytes - const - valid", TestTypeValidate, &cases.BytesConst{Val: []byte("foo")}, true, 0},
	{"bytes - const - invalid", TestTypeValidate, &cases.BytesConst{Val: []byte("bar")}, false, 1},

	{"bytes - in - valid", TestTypeValidate, &cases.BytesIn{Val: []byte("bar")}, true, 0},
	{"bytes - in - invalid", TestTypeValidate, &cases.BytesIn{Val: []byte("quux")}, false, 1},
	{"bytes - not in - valid", TestTypeValidate, &cases.BytesNotIn{Val: []byte("quux")}, true, 0},
	{"bytes - not in - invalid", TestTypeValidate, &cases.BytesNotIn{Val: []byte("fizz")}, false, 1},

	{"bytes - len - valid", TestTypeValidate, &cases.BytesLen{Val: []byte("baz")}, true, 0},
	{"bytes - len - invalid (lt)", TestTypeValidate, &cases.BytesLen{Val: []byte("go")}, false, 1},
	{"bytes - len - invalid (gt)", TestTypeValidate, &cases.BytesLen{Val: []byte("fizz")}, false, 1},

	{"bytes - min len - valid", TestTypeValidate, &cases.BytesMinLen{Val: []byte("fizz")}, true, 0},
	{"bytes - min len - valid (min)", TestTypeValidate, &cases.BytesMinLen{Val: []byte("baz")}, true, 0},
	{"bytes - min len - invalid", TestTypeValidate, &cases.BytesMinLen{Val: []byte("go")}, false, 1},

	{"bytes - max len - valid", TestTypeValidate, &cases.BytesMaxLen{Val: []byte("foo")}, true, 0},
	{"bytes - max len - valid (max)", TestTypeValidate, &cases.BytesMaxLen{Val: []byte("proto")}, true, 0},
	{"bytes - max len - invalid", TestTypeValidate, &cases.BytesMaxLen{Val: []byte("1234567890")}, false, 1},

	{"bytes - min/max len - valid", TestTypeValidate, &cases.BytesMinMaxLen{Val: []byte("quux")}, true, 0},
	{"bytes - min/max len - valid (min)", TestTypeValidate, &cases.BytesMinMaxLen{Val: []byte("foo")}, true, 0},
	{"bytes - min/max len - valid (max)", TestTypeValidate, &cases.BytesMinMaxLen{Val: []byte("proto")}, true, 0},
	{"bytes - min/max len - invalid (below)", TestTypeValidate, &cases.BytesMinMaxLen{Val: []byte("go")}, false, 1},
	{"bytes - min/max len - invalid (above)", TestTypeValidate, &cases.BytesMinMaxLen{Val: []byte("validate")}, false, 1},

	{"bytes - equal min/max len - valid", TestTypeValidate, &cases.BytesEqualMinMaxLen{Val: []byte("proto")}, true, 0},
	{"bytes - equal min/max len - invalid", TestTypeValidate, &cases.BytesEqualMinMaxLen{Val: []byte("validate")}, false, 1},

	{"bytes - pattern - valid", TestTypeValidate, &cases.BytesPattern{Val: []byte("Foo123")}, true, 0},
	{"bytes - pattern - invalid", TestTypeValidate, &cases.BytesPattern{Val: []byte("你好你好")}, false, 1},
	{"bytes - pattern - invalid (empty)", TestTypeValidate, &cases.BytesPattern{Val: []byte("")}, false, 1},

	{"bytes - prefix - valid", TestTypeValidate, &cases.BytesPrefix{Val: []byte{0x99, 0x9f, 0x08}}, true, 0},
	{"bytes - prefix - valid (only)", TestTypeValidate, &cases.BytesPrefix{Val: []byte{0x99}}, true, 0},
	{"bytes - prefix - invalid", TestTypeValidate, &cases.BytesPrefix{Val: []byte("bar")}, false, 1},

	{"bytes - contains - valid", TestTypeValidate, &cases.BytesContains{Val: []byte("candy bars")}, true, 0},
	{"bytes - contains - valid (only)", TestTypeValidate, &cases.BytesContains{Val: []byte("bar")}, true, 0},
	{"bytes - contains - invalid", TestTypeValidate, &cases.BytesContains{Val: []byte("candy bazs")}, false, 1},

	{"bytes - suffix - valid", TestTypeValidate, &cases.BytesSuffix{Val: []byte{0x62, 0x75, 0x7A, 0x7A}}, true, 0},
	{"bytes - suffix - valid (only)", TestTypeValidate, &cases.BytesSuffix{Val: []byte("\x62\x75\x7A\x7A")}, true, 0},
	{"bytes - suffix - invalid", TestTypeValidate, &cases.BytesSuffix{Val: []byte("foobar")}, false, 1},
	{"bytes - suffix - invalid (case-sensitive)", TestTypeValidate, &cases.BytesSuffix{Val: []byte("FooBaz")}, false, 1},

	{"bytes - IP - valid (v4)", TestTypeValidate, &cases.BytesIP{Val: []byte{0xC0, 0xA8, 0x00, 0x01}}, true, 0},
	{"bytes - IP - valid (v6)", TestTypeValidate, &cases.BytesIP{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")}, true, 0},
	{"bytes - IP - invalid", TestTypeValidate, &cases.BytesIP{Val: []byte("foobar")}, false, 1},

	{"bytes - IPv4 - valid", TestTypeValidate, &cases.BytesIPv4{Val: []byte{0xC0, 0xA8, 0x00, 0x01}}, true, 0},
	{"bytes - IPv4 - invalid", TestTypeValidate, &cases.BytesIPv4{Val: []byte("foobar")}, false, 1},
	{"bytes - IPv4 - invalid (v6)", TestTypeValidate, &cases.BytesIPv4{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")}, false, 1},

	{"bytes - IPv6 - valid", TestTypeValidate, &cases.BytesIPv6{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")}, true, 0},
	{"bytes - IPv6 - invalid", TestTypeValidate, &cases.BytesIPv6{Val: []byte("fooar")}, false, 1},
	{"bytes - IPv6 - invalid (v4)", TestTypeValidate, &cases.BytesIPv6{Val: []byte{0xC0, 0xA8, 0x00, 0x01}}, false, 1},
}

var enumCases = []TestCase{
	{"enum - none - valid", TestTypeValidate, &cases.EnumNone{Val: cases.TestEnum_ONE}, true, 0},

	{"enum - const - valid", TestTypeValidate, &cases.EnumConst{Val: cases.TestEnum_TWO}, true, 0},
	{"enum - const - invalid", TestTypeValidate, &cases.EnumConst{Val: cases.TestEnum_ONE}, false, 1},
	{"enum alias - const - valid", TestTypeValidate, &cases.EnumAliasConst{Val: cases.TestEnumAlias_C}, true, 0},
	{"enum alias - const - valid (alias)", TestTypeValidate, &cases.EnumAliasConst{Val: cases.TestEnumAlias_GAMMA}, true, 0},
	{"enum alias - const - invalid", TestTypeValidate, &cases.EnumAliasConst{Val: cases.TestEnumAlias_ALPHA}, false, 1},

	{"enum - defined_only - valid", TestTypeValidate, &cases.EnumDefined{Val: 0}, true, 0},
	{"enum - defined_only - invalid", TestTypeValidate, &cases.EnumDefined{Val: math.MaxInt32}, false, 1},
	{"enum alias - defined_only - valid", TestTypeValidate, &cases.EnumAliasDefined{Val: 1}, true, 0},
	{"enum alias - defined_only - invalid", TestTypeValidate, &cases.EnumAliasDefined{Val: math.MaxInt32}, false, 1},

	{"enum - in - valid", TestTypeValidate, &cases.EnumIn{Val: cases.TestEnum_TWO}, true, 0},
	{"enum - in - invalid", TestTypeValidate, &cases.EnumIn{Val: cases.TestEnum_ONE}, false, 1},
	{"enum alias - in - valid", TestTypeValidate, &cases.EnumAliasIn{Val: cases.TestEnumAlias_A}, true, 0},
	{"enum alias - in - valid (alias)", TestTypeValidate, &cases.EnumAliasIn{Val: cases.TestEnumAlias_ALPHA}, true, 0},
	{"enum alias - in - invalid", TestTypeValidate, &cases.EnumAliasIn{Val: cases.TestEnumAlias_BETA}, false, 1},

	{"enum - not in - valid", TestTypeValidate, &cases.EnumNotIn{Val: cases.TestEnum_ZERO}, true, 0},
	{"enum - not in - valid (undefined)", TestTypeValidate, &cases.EnumNotIn{Val: math.MaxInt32}, true, 0},
	{"enum - not in - invalid", TestTypeValidate, &cases.EnumNotIn{Val: cases.TestEnum_ONE}, false, 1},
	{"enum alias - not in - valid", TestTypeValidate, &cases.EnumAliasNotIn{Val: cases.TestEnumAlias_ALPHA}, true, 0},
	{"enum alias - not in - invalid", TestTypeValidate, &cases.EnumAliasNotIn{Val: cases.TestEnumAlias_B}, false, 1},
	{"enum alias - not in - invalid (alias)", TestTypeValidate, &cases.EnumAliasNotIn{Val: cases.TestEnumAlias_BETA}, false, 1},

	{"enum external - defined_only - valid", TestTypeValidate, &cases.EnumExternal{Val: other_package.Embed_VALUE}, true, 0},
	{"enum external - defined_only - invalid", TestTypeValidate, &cases.EnumExternal{Val: math.MaxInt32}, false, 1},

	{"enum repeated - defined_only - valid", TestTypeValidate, &cases.RepeatedEnumDefined{Val: []cases.TestEnum{cases.TestEnum_ONE, cases.TestEnum_TWO}}, true, 0},
	{"enum repeated - defined_only - invalid", TestTypeValidate, &cases.RepeatedEnumDefined{Val: []cases.TestEnum{cases.TestEnum_ONE, math.MaxInt32}}, false, 1},

	{"enum repeated (external) - defined_only - valid", TestTypeValidate, &cases.RepeatedExternalEnumDefined{Val: []other_package.Embed_Enumerated{other_package.Embed_VALUE}}, true, 0},
	{"enum repeated (external) - defined_only - invalid", TestTypeValidate, &cases.RepeatedExternalEnumDefined{Val: []other_package.Embed_Enumerated{math.MaxInt32}}, false, 1},

	{"enum map - defined_only - valid", TestTypeValidate, &cases.MapEnumDefined{Val: map[string]cases.TestEnum{"foo": cases.TestEnum_TWO}}, true, 0},
	{"enum map - defined_only - invalid", TestTypeValidate, &cases.MapEnumDefined{Val: map[string]cases.TestEnum{"foo": math.MaxInt32}}, false, 1},

	{"enum map (external) - defined_only - valid", TestTypeValidate, &cases.MapExternalEnumDefined{Val: map[string]other_package.Embed_Enumerated{"foo": other_package.Embed_VALUE}}, true, 0},
	{"enum map (external) - defined_only - invalid", TestTypeValidate, &cases.MapExternalEnumDefined{Val: map[string]other_package.Embed_Enumerated{"foo": math.MaxInt32}}, false, 1},
}

var messageCases = []TestCase{
	{"message - none - valid", TestTypeValidate, &cases.MessageNone{Val: &cases.MessageNone_NoneMsg{}}, true, 0},
	{"message - none - valid (unset)", TestTypeValidate, &cases.MessageNone{}, true, 0},

	{"message - disabled - valid", TestTypeValidate, &cases.MessageDisabled{Val: 456}, true, 0},
	{"message - disabled - valid (invalid field)", TestTypeValidate, &cases.MessageDisabled{Val: 0}, true, 0},

	{"message - ignored - valid", TestTypeValidate, &cases.MessageIgnored{Val: 456}, true, 0},
	{"message - ignored - valid (invalid field)", TestTypeValidate, &cases.MessageIgnored{Val: 0}, true, 0},

	{"message - field - valid", TestTypeValidate, &cases.Message{Val: &cases.TestMsg{Const: "foo"}}, true, 0},
	{"message - field - valid (unset)", TestTypeValidate, &cases.Message{}, true, 0},
	{"message - field - invalid", TestTypeValidate, &cases.Message{Val: &cases.TestMsg{}}, false, 1},
	{"message - field - invalid (transitive)", TestTypeValidate, &cases.Message{Val: &cases.TestMsg{Const: "foo", Nested: &cases.TestMsg{}}}, false, 1},

	{"message - skip - valid", TestTypeValidate, &cases.MessageSkip{Val: &cases.TestMsg{}}, true, 0},

	{"message - required - valid", TestTypeValidate, &cases.MessageRequired{Val: &cases.TestMsg{Const: "foo"}}, true, 0},
	{"message - required - invalid", TestTypeValidate, &cases.MessageRequired{}, false, 1},

	{"message - cross-package embed none - valid", TestTypeValidate, &cases.MessageCrossPackage{Val: &other_package.Embed{Val: 1}}, true, 0},
	{"message - cross-package embed none - valid (nil)", TestTypeValidate, &cases.MessageCrossPackage{}, true, 0},
	{"message - cross-package embed none - valid (empty)", TestTypeValidate, &cases.MessageCrossPackage{Val: &other_package.Embed{}}, false, 1},
	{"message - cross-package embed none - invalid", TestTypeValidate, &cases.MessageCrossPackage{Val: &other_package.Embed{Val: -1}}, false, 1},
}

var repeatedCases = []TestCase{
	{"repeated - none - valid", TestTypeValidate, &cases.RepeatedNone{Val: []int64{1, 2, 3}}, true, 0},

	{"repeated - embed none - valid", TestTypeValidate, &cases.RepeatedEmbedNone{Val: []*cases.Embed{{Val: 1}}}, true, 0},
	{"repeated - embed none - valid (nil)", TestTypeValidate, &cases.RepeatedEmbedNone{}, true, 0},
	{"repeated - embed none - valid (empty)", TestTypeValidate, &cases.RepeatedEmbedNone{Val: []*cases.Embed{}}, true, 0},
	{"repeated - embed none - invalid", TestTypeValidate, &cases.RepeatedEmbedNone{Val: []*cases.Embed{{Val: -1}}}, false, 1},

	{"repeated - cross-package embed none - valid", TestTypeValidate, &cases.RepeatedEmbedCrossPackageNone{Val: []*other_package.Embed{{Val: 1}}}, true, 0},
	{"repeated - cross-package embed none - valid (nil)", TestTypeValidate, &cases.RepeatedEmbedCrossPackageNone{}, true, 0},
	{"repeated - cross-package embed none - valid (empty)", TestTypeValidate, &cases.RepeatedEmbedCrossPackageNone{Val: []*other_package.Embed{}}, true, 0},
	{"repeated - cross-package embed none - invalid", TestTypeValidate, &cases.RepeatedEmbedCrossPackageNone{Val: []*other_package.Embed{{Val: -1}}}, false, 1},

	{"repeated - min - valid", TestTypeValidate, &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}, {Val: 2}, {Val: 3}}}, true, 0},
	{"repeated - min - valid (equal)", TestTypeValidate, &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}, {Val: 2}}}, true, 0},
	{"repeated - min - invalid", TestTypeValidate, &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}}}, false, 1},
	{"repeated - min - invalid (element)", TestTypeValidate, &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}, {Val: -1}}}, false, 1},

	{"repeated - max - valid", TestTypeValidate, &cases.RepeatedMax{Val: []float64{1, 2}}, true, 0},
	{"repeated - max - valid (equal)", TestTypeValidate, &cases.RepeatedMax{Val: []float64{1, 2, 3}}, true, 0},
	{"repeated - max - invalid", TestTypeValidate, &cases.RepeatedMax{Val: []float64{1, 2, 3, 4}}, false, 1},

	{"repeated - min/max - valid", TestTypeValidate, &cases.RepeatedMinMax{Val: []int32{1, 2, 3}}, true, 0},
	{"repeated - min/max - valid (min)", TestTypeValidate, &cases.RepeatedMinMax{Val: []int32{1, 2}}, true, 0},
	{"repeated - min/max - valid (max)", TestTypeValidate, &cases.RepeatedMinMax{Val: []int32{1, 2, 3, 4}}, true, 0},
	{"repeated - min/max - invalid (below)", TestTypeValidate, &cases.RepeatedMinMax{Val: []int32{}}, false, 1},
	{"repeated - min/max - invalid (above)", TestTypeValidate, &cases.RepeatedMinMax{Val: []int32{1, 2, 3, 4, 5}}, false, 1},

	{"repeated - exact - valid", TestTypeValidate, &cases.RepeatedExact{Val: []uint32{1, 2, 3}}, true, 0},
	{"repeated - exact - invalid (below)", TestTypeValidate, &cases.RepeatedExact{Val: []uint32{1, 2}}, false, 1},
	{"repeated - exact - invalid (above)", TestTypeValidate, &cases.RepeatedExact{Val: []uint32{1, 2, 3, 4}}, false, 1},

	{"repeated - unique - valid", TestTypeValidate, &cases.RepeatedUnique{Val: []string{"foo", "bar", "baz"}}, true, 0},
	{"repeated - unique - valid (empty)", TestTypeValidate, &cases.RepeatedUnique{}, true, 0},
	{"repeated - unique - valid (case sensitivity)", TestTypeValidate, &cases.RepeatedUnique{Val: []string{"foo", "Foo"}}, true, 0},
	{"repeated - unique - invalid", TestTypeValidate, &cases.RepeatedUnique{Val: []string{"foo", "bar", "foo", "baz"}}, false, 1},

	{"repeated - items - valid", TestTypeValidate, &cases.RepeatedItemRule{Val: []float32{1, 2, 3}}, true, 0},
	{"repeated - items - valid (empty)", TestTypeValidate, &cases.RepeatedItemRule{Val: []float32{}}, true, 0},
	{"repeated - items - valid (pattern)", TestTypeValidate, &cases.RepeatedItemPattern{Val: []string{"Alpha", "Beta123"}}, true, 0},
	{"repeated - items - invalid", TestTypeValidate, &cases.RepeatedItemRule{Val: []float32{1, -2, 3}}, false, 1},
	{"repeated - items - invalid (pattern)", TestTypeValidate, &cases.RepeatedItemPattern{Val: []string{"Alpha", "!@#$%^&*()"}}, false, 1},
	{"repeated - items - invalid (in)", TestTypeValidate, &cases.RepeatedItemIn{Val: []string{"baz"}}, false, 1},
	{"repeated - items - valid (in)", TestTypeValidate, &cases.RepeatedItemIn{Val: []string{"foo"}}, true, 0},
	{"repeated - items - invalid (not_in)", TestTypeValidate, &cases.RepeatedItemNotIn{Val: []string{"foo"}}, false, 1},
	{"repeated - items - valid (not_in)", TestTypeValidate, &cases.RepeatedItemNotIn{Val: []string{"baz"}}, true, 0},

	{"repeated - items - invalid (enum in)", TestTypeValidate, &cases.RepeatedEnumIn{Val: []cases.AnEnum{1}}, false, 1},
	{"repeated - items - valid (enum in)", TestTypeValidate, &cases.RepeatedEnumIn{Val: []cases.AnEnum{0}}, true, 0},
	{"repeated - items - invalid (enum not_in)", TestTypeValidate, &cases.RepeatedEnumNotIn{Val: []cases.AnEnum{0}}, false, 1},
	{"repeated - items - valid (enum not_in)", TestTypeValidate, &cases.RepeatedEnumNotIn{Val: []cases.AnEnum{1}}, true, 0},
	{"repeated - items - invalid (embedded enum in)", TestTypeValidate, &cases.RepeatedEmbeddedEnumIn{Val: []cases.RepeatedEmbeddedEnumIn_AnotherInEnum{1}}, false, 1},
	{"repeated - items - valid (embedded enum in)", TestTypeValidate, &cases.RepeatedEmbeddedEnumIn{Val: []cases.RepeatedEmbeddedEnumIn_AnotherInEnum{0}}, true, 0},
	{"repeated - items - invalid (embedded enum not_in)", TestTypeValidate, &cases.RepeatedEmbeddedEnumNotIn{Val: []cases.RepeatedEmbeddedEnumNotIn_AnotherNotInEnum{0}}, false, 1},
	{"repeated - items - valid (embedded enum not_in)", TestTypeValidate, &cases.RepeatedEmbeddedEnumNotIn{Val: []cases.RepeatedEmbeddedEnumNotIn_AnotherNotInEnum{1}}, true, 0},

	{"repeated - embed skip - valid", TestTypeValidate, &cases.RepeatedEmbedSkip{Val: []*cases.Embed{{Val: 1}}}, true, 0},
	{"repeated - embed skip - valid (invalid element)", TestTypeValidate, &cases.RepeatedEmbedSkip{Val: []*cases.Embed{{Val: -1}}}, true, 0},
	{"repeated - min and items len - valid", TestTypeValidate, &cases.RepeatedMinAndItemLen{Val: []string{"aaa", "bbb"}}, true, 0},
	{"repeated - min and items len - invalid (min)", TestTypeValidate, &cases.RepeatedMinAndItemLen{Val: []string{}}, false, 1},
	{"repeated - min and items len - invalid (len)", TestTypeValidate, &cases.RepeatedMinAndItemLen{Val: []string{"x"}}, false, 1},
	{"repeated - min and max items len - valid", TestTypeValidate, &cases.RepeatedMinAndMaxItemLen{Val: []string{"aaa", "bbb"}}, true, 0},
	{"repeated - min and max items len - invalid (min_len)", TestTypeValidate, &cases.RepeatedMinAndMaxItemLen{}, false, 1},
	{"repeated - min and max items len - invalid (max_len)", TestTypeValidate, &cases.RepeatedMinAndMaxItemLen{Val: []string{"aaa", "bbb", "ccc", "ddd"}}, false, 1},

	{"repeated - duration - gte - valid", TestTypeValidate, &cases.RepeatedDuration{Val: []*duration.Duration{{Seconds: 3}}}, true, 0},
	{"repeated - duration - gte - valid (empty)", TestTypeValidate, &cases.RepeatedDuration{}, true, 0},
	{"repeated - duration - gte - valid (equal)", TestTypeValidate, &cases.RepeatedDuration{Val: []*duration.Duration{{Nanos: 1000000}}}, true, 0},
	{"repeated - duration - gte - invalid", TestTypeValidate, &cases.RepeatedDuration{Val: []*duration.Duration{{Seconds: -1}}}, false, 1},
}

var mapCases = []TestCase{
	{"map - none - valid", TestTypeValidate, &cases.MapNone{Val: map[uint32]bool{123: true, 456: false}}, true, 0},

	{"map - min pairs - valid", TestTypeValidate, &cases.MapMin{Val: map[int32]float32{1: 2, 3: 4, 5: 6}}, true, 0},
	{"map - min pairs - valid (equal)", TestTypeValidate, &cases.MapMin{Val: map[int32]float32{1: 2, 3: 4}}, true, 0},
	{"map - min pairs - invalid", TestTypeValidate, &cases.MapMin{Val: map[int32]float32{1: 2}}, false, 1},

	{"map - max pairs - valid", TestTypeValidate, &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4}}, true, 0},
	{"map - max pairs - valid (equal)", TestTypeValidate, &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4, 5: 6}}, true, 0},
	{"map - max pairs - invalid", TestTypeValidate, &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4, 5: 6, 7: 8}}, false, 1},

	{"map - min/max - valid", TestTypeValidate, &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true}}, true, 0},
	{"map - min/max - valid (min)", TestTypeValidate, &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false}}, true, 0},
	{"map - min/max - valid (max)", TestTypeValidate, &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true, "d": false}}, true, 0},
	{"map - min/max - invalid (below)", TestTypeValidate, &cases.MapMinMax{Val: map[string]bool{}}, false, 1},
	{"map - min/max - invalid (above)", TestTypeValidate, &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true, "d": false, "e": true}}, false, 1},

	{"map - exact - valid", TestTypeValidate, &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b", 3: "c"}}, true, 0},
	{"map - exact - invalid (below)", TestTypeValidate, &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b"}}, false, 1},
	{"map - exact - invalid (above)", TestTypeValidate, &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b", 3: "c", 4: "d"}}, false, 1},

	{"map - no sparse - valid", TestTypeValidate, &cases.MapNoSparse{Val: map[uint32]*cases.MapNoSparse_Msg{1: {}, 2: {}}}, true, 0},
	{"map - no sparse - valid (empty)", TestTypeValidate, &cases.MapNoSparse{Val: map[uint32]*cases.MapNoSparse_Msg{}}, true, 0},
	// sparse maps are no longer supported, so this case is no longer possible
	//{"map - no sparse - invalid", &cases.MapNoSparse{Val: map[uint32]*cases.MapNoSparse_Msg{1: {}, 2: nil}}, false, 1},

	{"map - keys - valid", TestTypeValidate, &cases.MapKeys{Val: map[int64]string{-1: "a", -2: "b"}}, true, 0},
	{"map - keys - valid (empty)", TestTypeValidate, &cases.MapKeys{Val: map[int64]string{}}, true, 0},
	{"map - keys - valid (pattern)", TestTypeValidate, &cases.MapKeysPattern{Val: map[string]string{"A": "a"}}, true, 0},
	{"map - keys - invalid", TestTypeValidate, &cases.MapKeys{Val: map[int64]string{1: "a"}}, false, 1},
	{"map - keys - invalid (pattern)", TestTypeValidate, &cases.MapKeysPattern{Val: map[string]string{"A": "a", "!@#$%^&*()": "b"}}, false, 1},

	{"map - values - valid", TestTypeValidate, &cases.MapValues{Val: map[string]string{"a": "Alpha", "b": "Beta"}}, true, 0},
	{"map - values - valid (empty)", TestTypeValidate, &cases.MapValues{Val: map[string]string{}}, true, 0},
	{"map - values - valid (pattern)", TestTypeValidate, &cases.MapValuesPattern{Val: map[string]string{"a": "A"}}, true, 0},
	{"map - values - invalid", TestTypeValidate, &cases.MapValues{Val: map[string]string{"a": "A", "b": "B"}}, false, 1},
	{"map - values - invalid (pattern)", TestTypeValidate, &cases.MapValuesPattern{Val: map[string]string{"a": "A", "b": "!@#$%^&*()"}}, false, 1},

	{"map - recursive - valid", TestTypeValidate, &cases.MapRecursive{Val: map[uint32]*cases.MapRecursive_Msg{1: {Val: "abc"}}}, true, 0},
	{"map - recursive - invalid", TestTypeValidate, &cases.MapRecursive{Val: map[uint32]*cases.MapRecursive_Msg{1: {}}}, false, 1},
}

var oneofCases = []TestCase{
	{"oneof - none - valid", TestTypeValidate, &cases.OneOfNone{O: &cases.OneOfNone_X{X: "foo"}}, true, 0},
	{"oneof - none - valid (empty)", TestTypeValidate, &cases.OneOfNone{}, true, 0},

	{"oneof - field - valid (X)", TestTypeValidate, &cases.OneOf{O: &cases.OneOf_X{X: "foobar"}}, true, 0},
	{"oneof - field - valid (Y)", TestTypeValidate, &cases.OneOf{O: &cases.OneOf_Y{Y: 123}}, true, 0},
	{"oneof - field - valid (Z)", TestTypeValidate, &cases.OneOf{O: &cases.OneOf_Z{Z: &cases.TestOneOfMsg{Val: true}}}, true, 0},
	{"oneof - field - valid (empty)", TestTypeValidate, &cases.OneOf{}, true, 0},
	{"oneof - field - invalid (X)", TestTypeValidate, &cases.OneOf{O: &cases.OneOf_X{X: "fizzbuzz"}}, false, 1},
	{"oneof - field - invalid (Y)", TestTypeValidate, &cases.OneOf{O: &cases.OneOf_Y{Y: -1}}, false, 1},
	{"oneof - filed - invalid (Z)", TestTypeValidate, &cases.OneOf{O: &cases.OneOf_Z{Z: &cases.TestOneOfMsg{}}}, false, 1},

	{"oneof - required - valid", TestTypeValidate, &cases.OneOfRequired{O: &cases.OneOfRequired_X{X: ""}}, true, 0},
	{"oneof - require - invalid", TestTypeValidate, &cases.OneOfRequired{}, false, 1},
}

var wrapperCases = []TestCase{
	{"wrapper - none - valid", TestTypeValidate, &cases.WrapperNone{Val: &wrappers.Int32Value{Value: 123}}, true, 0},
	{"wrapper - none - valid (empty)", TestTypeValidate, &cases.WrapperNone{Val: nil}, true, 0},

	{"wrapper - float - valid", TestTypeValidate, &cases.WrapperFloat{Val: &wrappers.FloatValue{Value: 1}}, true, 0},
	{"wrapper - float - valid (empty)", TestTypeValidate, &cases.WrapperFloat{Val: nil}, true, 0},
	{"wrapper - float - invalid", TestTypeValidate, &cases.WrapperFloat{Val: &wrappers.FloatValue{Value: 0}}, false, 1},

	{"wrapper - double - valid", TestTypeValidate, &cases.WrapperDouble{Val: &wrappers.DoubleValue{Value: 1}}, true, 0},
	{"wrapper - double - valid (empty)", TestTypeValidate, &cases.WrapperDouble{Val: nil}, true, 0},
	{"wrapper - double - invalid", TestTypeValidate, &cases.WrapperDouble{Val: &wrappers.DoubleValue{Value: 0}}, false, 1},

	{"wrapper - int64 - valid", TestTypeValidate, &cases.WrapperInt64{Val: &wrappers.Int64Value{Value: 1}}, true, 0},
	{"wrapper - int64 - valid (empty)", TestTypeValidate, &cases.WrapperInt64{Val: nil}, true, 0},
	{"wrapper - int64 - invalid", TestTypeValidate, &cases.WrapperInt64{Val: &wrappers.Int64Value{Value: 0}}, false, 1},

	{"wrapper - int32 - valid", TestTypeValidate, &cases.WrapperInt32{Val: &wrappers.Int32Value{Value: 1}}, true, 0},
	{"wrapper - int32 - valid (empty)", TestTypeValidate, &cases.WrapperInt32{Val: nil}, true, 0},
	{"wrapper - int32 - invalid", TestTypeValidate, &cases.WrapperInt32{Val: &wrappers.Int32Value{Value: 0}}, false, 1},

	{"wrapper - uint64 - valid", TestTypeValidate, &cases.WrapperUInt64{Val: &wrappers.UInt64Value{Value: 1}}, true, 0},
	{"wrapper - uint64 - valid (empty)", TestTypeValidate, &cases.WrapperUInt64{Val: nil}, true, 0},
	{"wrapper - uint64 - invalid", TestTypeValidate, &cases.WrapperUInt64{Val: &wrappers.UInt64Value{Value: 0}}, false, 1},

	{"wrapper - uint32 - valid", TestTypeValidate, &cases.WrapperUInt32{Val: &wrappers.UInt32Value{Value: 1}}, true, 0},
	{"wrapper - uint32 - valid (empty)", TestTypeValidate, &cases.WrapperUInt32{Val: nil}, true, 0},
	{"wrapper - uint32 - invalid", TestTypeValidate, &cases.WrapperUInt32{Val: &wrappers.UInt32Value{Value: 0}}, false, 1},

	{"wrapper - bool - valid", TestTypeValidate, &cases.WrapperBool{Val: &wrappers.BoolValue{Value: true}}, true, 0},
	{"wrapper - bool - valid (empty)", TestTypeValidate, &cases.WrapperBool{Val: nil}, true, 0},
	{"wrapper - bool - invalid", TestTypeValidate, &cases.WrapperBool{Val: &wrappers.BoolValue{Value: false}}, false, 1},

	{"wrapper - string - valid", TestTypeValidate, &cases.WrapperString{Val: &wrappers.StringValue{Value: "foobar"}}, true, 0},
	{"wrapper - string - valid (empty)", TestTypeValidate, &cases.WrapperString{Val: nil}, true, 0},
	{"wrapper - string - invalid", TestTypeValidate, &cases.WrapperString{Val: &wrappers.StringValue{Value: "fizzbuzz"}}, false, 1},

	{"wrapper - bytes - valid", TestTypeValidate, &cases.WrapperBytes{Val: &wrappers.BytesValue{Value: []byte("foo")}}, true, 0},
	{"wrapper - bytes - valid (empty)", TestTypeValidate, &cases.WrapperBytes{Val: nil}, true, 0},
	{"wrapper - bytes - invalid", TestTypeValidate, &cases.WrapperBytes{Val: &wrappers.BytesValue{Value: []byte("x")}}, false, 1},

	{"wrapper - required - string - valid", TestTypeValidate, &cases.WrapperRequiredString{Val: &wrappers.StringValue{Value: "bar"}}, true, 0},
	{"wrapper - required - string - invalid", TestTypeValidate, &cases.WrapperRequiredString{Val: &wrappers.StringValue{Value: "foo"}}, false, 1},
	{"wrapper - required - string - invalid (empty)", TestTypeValidate, &cases.WrapperRequiredString{}, false, 1},

	{"wrapper - required - string (empty) - valid", TestTypeValidate, &cases.WrapperRequiredEmptyString{Val: &wrappers.StringValue{Value: ""}}, true, 0},
	{"wrapper - required - string (empty) - invalid", TestTypeValidate, &cases.WrapperRequiredEmptyString{Val: &wrappers.StringValue{Value: "foo"}}, false, 1},
	{"wrapper - required - string (empty) - invalid (empty)", TestTypeValidate, &cases.WrapperRequiredEmptyString{}, false, 1},

	{"wrapper - optional - string (uuid) - valid", TestTypeValidate, &cases.WrapperOptionalUuidString{Val: &wrappers.StringValue{Value: "8b72987b-024a-43b3-b4cf-647a1f925c5d"}}, true, 0},
	{"wrapper - optional - string (uuid) - valid (empty)", TestTypeValidate, &cases.WrapperOptionalUuidString{}, true, 0},
	{"wrapper - optional - string (uuid) - invalid", TestTypeValidate, &cases.WrapperOptionalUuidString{Val: &wrappers.StringValue{Value: "foo"}}, false, 1},

	{"wrapper - required - float - valid", TestTypeValidate, &cases.WrapperRequiredFloat{Val: &wrappers.FloatValue{Value: 1}}, true, 0},
	{"wrapper - required - float - invalid", TestTypeValidate, &cases.WrapperRequiredFloat{Val: &wrappers.FloatValue{Value: -5}}, false, 1},
	{"wrapper - required - float - invalid (empty)", TestTypeValidate, &cases.WrapperRequiredFloat{}, false, 1},
}

var durationCases = []TestCase{
	{"duration - none - valid", TestTypeValidate, &cases.DurationNone{Val: &duration.Duration{Seconds: 123}}, true, 0},

	{"duration - required - valid", TestTypeValidate, &cases.DurationRequired{Val: &duration.Duration{}}, true, 0},
	{"duration - required - invalid", TestTypeValidate, &cases.DurationRequired{Val: nil}, false, 1},

	{"duration - const - valid", TestTypeValidate, &cases.DurationConst{Val: &duration.Duration{Seconds: 3}}, true, 0},
	{"duration - const - valid (empty)", TestTypeValidate, &cases.DurationConst{}, true, 0},
	{"duration - const - invalid", TestTypeValidate, &cases.DurationConst{Val: &duration.Duration{Nanos: 3}}, false, 1},

	{"duration - in - valid", TestTypeValidate, &cases.DurationIn{Val: &duration.Duration{Seconds: 1}}, true, 0},
	{"duration - in - valid (empty)", TestTypeValidate, &cases.DurationIn{}, true, 0},
	{"duration - in - invalid", TestTypeValidate, &cases.DurationIn{Val: &duration.Duration{}}, false, 1},

	{"duration - not in - valid", TestTypeValidate, &cases.DurationNotIn{Val: &duration.Duration{Nanos: 1}}, true, 0},
	{"duration - not in - valid (empty)", TestTypeValidate, &cases.DurationNotIn{}, true, 0},
	{"duration - not in - invalid", TestTypeValidate, &cases.DurationNotIn{Val: &duration.Duration{}}, false, 1},

	{"duration - lt - valid", TestTypeValidate, &cases.DurationLT{Val: &duration.Duration{Nanos: -1}}, true, 0},
	{"duration - lt - valid (empty)", TestTypeValidate, &cases.DurationLT{}, true, 0},
	{"duration - lt - invalid (equal)", TestTypeValidate, &cases.DurationLT{Val: &duration.Duration{}}, false, 1},
	{"duration - lt - invalid", TestTypeValidate, &cases.DurationLT{Val: &duration.Duration{Seconds: 1}}, false, 1},

	{"duration - lte - valid", TestTypeValidate, &cases.DurationLTE{Val: &duration.Duration{}}, true, 0},
	{"duration - lte - valid (empty)", TestTypeValidate, &cases.DurationLTE{}, true, 0},
	{"duration - lte - valid (equal)", TestTypeValidate, &cases.DurationLTE{Val: &duration.Duration{Seconds: 1}}, true, 0},
	{"duration - lte - invalid", TestTypeValidate, &cases.DurationLTE{Val: &duration.Duration{Seconds: 1, Nanos: 1}}, false, 1},

	{"duration - gt - valid", TestTypeValidate, &cases.DurationGT{Val: &duration.Duration{Seconds: 1}}, true, 0},
	{"duration - gt - valid (empty)", TestTypeValidate, &cases.DurationGT{}, true, 0},
	{"duration - gt - invalid (equal)", TestTypeValidate, &cases.DurationGT{Val: &duration.Duration{Nanos: 1000}}, false, 1},
	{"duration - gt - invalid", TestTypeValidate, &cases.DurationGT{Val: &duration.Duration{}}, false, 1},

	{"duration - gte - valid", TestTypeValidate, &cases.DurationGTE{Val: &duration.Duration{Seconds: 3}}, true, 0},
	{"duration - gte - valid (empty)", TestTypeValidate, &cases.DurationGTE{}, true, 0},
	{"duration - gte - valid (equal)", TestTypeValidate, &cases.DurationGTE{Val: &duration.Duration{Nanos: 1000000}}, true, 0},
	{"duration - gte - invalid", TestTypeValidate, &cases.DurationGTE{Val: &duration.Duration{Seconds: -1}}, false, 1},

	{"duration - gt & lt - valid", TestTypeValidate, &cases.DurationGTLT{Val: &duration.Duration{Nanos: 1000}}, true, 0},
	{"duration - gt & lt - valid (empty)", TestTypeValidate, &cases.DurationGTLT{}, true, 0},
	{"duration - gt & lt - invalid (above)", TestTypeValidate, &cases.DurationGTLT{Val: &duration.Duration{Seconds: 1000}}, false, 1},
	{"duration - gt & lt - invalid (below)", TestTypeValidate, &cases.DurationGTLT{Val: &duration.Duration{Nanos: -1000}}, false, 1},
	{"duration - gt & lt - invalid (max)", TestTypeValidate, &cases.DurationGTLT{Val: &duration.Duration{Seconds: 1}}, false, 1},
	{"duration - gt & lt - invalid (min)", TestTypeValidate, &cases.DurationGTLT{Val: &duration.Duration{}}, false, 1},

	{"duration - exclusive gt & lt - valid (empty)", TestTypeValidate, &cases.DurationExLTGT{}, true, 0},
	{"duration - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.DurationExLTGT{Val: &duration.Duration{Seconds: 2}}, true, 0},
	{"duration - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.DurationExLTGT{Val: &duration.Duration{Nanos: -1}}, true, 0},
	{"duration - exclusive gt & lt - invalid", TestTypeValidate, &cases.DurationExLTGT{Val: &duration.Duration{Nanos: 1000}}, false, 1},
	{"duration - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.DurationExLTGT{Val: &duration.Duration{Seconds: 1}}, false, 1},
	{"duration - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.DurationExLTGT{Val: &duration.Duration{}}, false, 1},

	{"duration - gte & lte - valid", TestTypeValidate, &cases.DurationGTELTE{Val: &duration.Duration{Seconds: 60, Nanos: 1}}, true, 0},
	{"duration - gte & lte - valid (empty)", TestTypeValidate, &cases.DurationGTELTE{}, true, 0},
	{"duration - gte & lte - valid (max)", TestTypeValidate, &cases.DurationGTELTE{Val: &duration.Duration{Seconds: 3600}}, true, 0},
	{"duration - gte & lte - valid (min)", TestTypeValidate, &cases.DurationGTELTE{Val: &duration.Duration{Seconds: 60}}, true, 0},
	{"duration - gte & lte - invalid (above)", TestTypeValidate, &cases.DurationGTELTE{Val: &duration.Duration{Seconds: 3600, Nanos: 1}}, false, 1},
	{"duration - gte & lte - invalid (below)", TestTypeValidate, &cases.DurationGTELTE{Val: &duration.Duration{Seconds: 59}}, false, 1},

	{"duration - gte & lte - valid (empty)", TestTypeValidate, &cases.DurationExGTELTE{}, true, 0},
	{"duration - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.DurationExGTELTE{Val: &duration.Duration{Seconds: 3601}}, true, 0},
	{"duration - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.DurationExGTELTE{Val: &duration.Duration{}}, true, 0},
	{"duration - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.DurationExGTELTE{Val: &duration.Duration{Seconds: 3600}}, true, 0},
	{"duration - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.DurationExGTELTE{Val: &duration.Duration{Seconds: 60}}, true, 0},
	{"duration - exclusive gte & lte - invalid", TestTypeValidate, &cases.DurationExGTELTE{Val: &duration.Duration{Seconds: 61}}, false, 1},
	{"duration - fields with other fields - invalid other field", TestTypeValidate, &cases.DurationFieldWithOtherFields{DurationVal: nil, IntVal: 12}, false, 1},
}

var timestampCases = []TestCase{
	{"timestamp - none - valid", TestTypeValidate, &cases.TimestampNone{Val: &timestamp.Timestamp{Seconds: 123}}, true, 0},

	{"timestamp - required - valid", TestTypeValidate, &cases.TimestampRequired{Val: &timestamp.Timestamp{}}, true, 0},
	{"timestamp - required - invalid", TestTypeValidate, &cases.TimestampRequired{Val: nil}, false, 1},

	{"timestamp - const - valid", TestTypeValidate, &cases.TimestampConst{Val: &timestamp.Timestamp{Seconds: 3}}, true, 0},
	{"timestamp - const - valid (empty)", TestTypeValidate, &cases.TimestampConst{}, true, 0},
	{"timestamp - const - invalid", TestTypeValidate, &cases.TimestampConst{Val: &timestamp.Timestamp{Nanos: 3}}, false, 1},

	{"timestamp - lt - valid", TestTypeValidate, &cases.TimestampLT{Val: &timestamp.Timestamp{Seconds: -1}}, true, 0},
	{"timestamp - lt - valid (empty)", TestTypeValidate, &cases.TimestampLT{}, true, 0},
	{"timestamp - lt - invalid (equal)", TestTypeValidate, &cases.TimestampLT{Val: &timestamp.Timestamp{}}, false, 1},
	{"timestamp - lt - invalid", TestTypeValidate, &cases.TimestampLT{Val: &timestamp.Timestamp{Seconds: 1}}, false, 1},

	{"timestamp - lte - valid", TestTypeValidate, &cases.TimestampLTE{Val: &timestamp.Timestamp{}}, true, 0},
	{"timestamp - lte - valid (empty)", TestTypeValidate, &cases.TimestampLTE{}, true, 0},
	{"timestamp - lte - valid (equal)", TestTypeValidate, &cases.TimestampLTE{Val: &timestamp.Timestamp{Seconds: 1}}, true, 0},
	{"timestamp - lte - invalid", TestTypeValidate, &cases.TimestampLTE{Val: &timestamp.Timestamp{Seconds: 1, Nanos: 1}}, false, 1},

	{"timestamp - gt - valid", TestTypeValidate, &cases.TimestampGT{Val: &timestamp.Timestamp{Seconds: 1}}, true, 0},
	{"timestamp - gt - valid (empty)", TestTypeValidate, &cases.TimestampGT{}, true, 0},
	{"timestamp - gt - invalid (equal)", TestTypeValidate, &cases.TimestampGT{Val: &timestamp.Timestamp{Nanos: 1000}}, false, 1},
	{"timestamp - gt - invalid", TestTypeValidate, &cases.TimestampGT{Val: &timestamp.Timestamp{}}, false, 1},

	{"timestamp - gte - valid", TestTypeValidate, &cases.TimestampGTE{Val: &timestamp.Timestamp{Seconds: 3}}, true, 0},
	{"timestamp - gte - valid (empty)", TestTypeValidate, &cases.TimestampGTE{}, true, 0},
	{"timestamp - gte - valid (equal)", TestTypeValidate, &cases.TimestampGTE{Val: &timestamp.Timestamp{Nanos: 1000000}}, true, 0},
	{"timestamp - gte - invalid", TestTypeValidate, &cases.TimestampGTE{Val: &timestamp.Timestamp{Seconds: -1}}, false, 1},

	{"timestamp - gt & lt - valid", TestTypeValidate, &cases.TimestampGTLT{Val: &timestamp.Timestamp{Nanos: 1000}}, true, 0},
	{"timestamp - gt & lt - valid (empty)", TestTypeValidate, &cases.TimestampGTLT{}, true, 0},
	{"timestamp - gt & lt - invalid (above)", TestTypeValidate, &cases.TimestampGTLT{Val: &timestamp.Timestamp{Seconds: 1000}}, false, 1},
	{"timestamp - gt & lt - invalid (below)", TestTypeValidate, &cases.TimestampGTLT{Val: &timestamp.Timestamp{Seconds: -1000}}, false, 1},
	{"timestamp - gt & lt - invalid (max)", TestTypeValidate, &cases.TimestampGTLT{Val: &timestamp.Timestamp{Seconds: 1}}, false, 1},
	{"timestamp - gt & lt - invalid (min)", TestTypeValidate, &cases.TimestampGTLT{Val: &timestamp.Timestamp{}}, false, 1},

	{"timestamp - exclusive gt & lt - valid (empty)", TestTypeValidate, &cases.TimestampExLTGT{}, true, 0},
	{"timestamp - exclusive gt & lt - valid (above)", TestTypeValidate, &cases.TimestampExLTGT{Val: &timestamp.Timestamp{Seconds: 2}}, true, 0},
	{"timestamp - exclusive gt & lt - valid (below)", TestTypeValidate, &cases.TimestampExLTGT{Val: &timestamp.Timestamp{Seconds: -1}}, true, 0},
	{"timestamp - exclusive gt & lt - invalid", TestTypeValidate, &cases.TimestampExLTGT{Val: &timestamp.Timestamp{Nanos: 1000}}, false, 1},
	{"timestamp - exclusive gt & lt - invalid (max)", TestTypeValidate, &cases.TimestampExLTGT{Val: &timestamp.Timestamp{Seconds: 1}}, false, 1},
	{"timestamp - exclusive gt & lt - invalid (min)", TestTypeValidate, &cases.TimestampExLTGT{Val: &timestamp.Timestamp{}}, false, 1},

	{"timestamp - gte & lte - valid", TestTypeValidate, &cases.TimestampGTELTE{Val: &timestamp.Timestamp{Seconds: 60, Nanos: 1}}, true, 0},
	{"timestamp - gte & lte - valid (empty)", TestTypeValidate, &cases.TimestampGTELTE{}, true, 0},
	{"timestamp - gte & lte - valid (max)", TestTypeValidate, &cases.TimestampGTELTE{Val: &timestamp.Timestamp{Seconds: 3600}}, true, 0},
	{"timestamp - gte & lte - valid (min)", TestTypeValidate, &cases.TimestampGTELTE{Val: &timestamp.Timestamp{Seconds: 60}}, true, 0},
	{"timestamp - gte & lte - invalid (above)", TestTypeValidate, &cases.TimestampGTELTE{Val: &timestamp.Timestamp{Seconds: 3600, Nanos: 1}}, false, 1},
	{"timestamp - gte & lte - invalid (below)", TestTypeValidate, &cases.TimestampGTELTE{Val: &timestamp.Timestamp{Seconds: 59}}, false, 1},

	{"timestamp - gte & lte - valid (empty)", TestTypeValidate, &cases.TimestampExGTELTE{}, true, 0},
	{"timestamp - exclusive gte & lte - valid (above)", TestTypeValidate, &cases.TimestampExGTELTE{Val: &timestamp.Timestamp{Seconds: 3601}}, true, 0},
	{"timestamp - exclusive gte & lte - valid (below)", TestTypeValidate, &cases.TimestampExGTELTE{Val: &timestamp.Timestamp{}}, true, 0},
	{"timestamp - exclusive gte & lte - valid (max)", TestTypeValidate, &cases.TimestampExGTELTE{Val: &timestamp.Timestamp{Seconds: 3600}}, true, 0},
	{"timestamp - exclusive gte & lte - valid (min)", TestTypeValidate, &cases.TimestampExGTELTE{Val: &timestamp.Timestamp{Seconds: 60}}, true, 0},
	{"timestamp - exclusive gte & lte - invalid", TestTypeValidate, &cases.TimestampExGTELTE{Val: &timestamp.Timestamp{Seconds: 61}}, false, 1},

	{"timestamp - lt now - valid", TestTypeValidate, &cases.TimestampLTNow{Val: &timestamp.Timestamp{}}, true, 0},
	{"timestamp - lt now - valid (empty)", TestTypeValidate, &cases.TimestampLTNow{}, true, 0},
	{"timestamp - lt now - invalid", TestTypeValidate, &cases.TimestampLTNow{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 7200}}, false, 1},

	{"timestamp - gt now - valid", TestTypeValidate, &cases.TimestampGTNow{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 7200}}, true, 0},
	{"timestamp - gt now - valid (empty)", TestTypeValidate, &cases.TimestampGTNow{}, true, 0},
	{"timestamp - gt now - invalid", TestTypeValidate, &cases.TimestampGTNow{Val: &timestamp.Timestamp{}}, false, 1},

	{"timestamp - within - valid", TestTypeValidate, &cases.TimestampWithin{Val: ptypes.TimestampNow()}, true, 0},
	{"timestamp - within - valid (empty)", TestTypeValidate, &cases.TimestampWithin{}, true, 0},
	{"timestamp - within - invalid (below)", TestTypeValidate, &cases.TimestampWithin{Val: &timestamp.Timestamp{}}, false, 1},
	{"timestamp - within - invalid (above)", TestTypeValidate, &cases.TimestampWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 7200}}, false, 1},

	{"timestamp - lt now within - valid", TestTypeValidate, &cases.TimestampLTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() - 1800}}, true, 0},
	{"timestamp - lt now within - valid (empty)", TestTypeValidate, &cases.TimestampLTNowWithin{}, true, 0},
	{"timestamp - lt now within - invalid (lt)", TestTypeValidate, &cases.TimestampLTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 1800}}, false, 1},
	{"timestamp - lt now within - invalid (within)", TestTypeValidate, &cases.TimestampLTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() - 7200}}, false, 1},

	{"timestamp - gt now within - valid", TestTypeValidate, &cases.TimestampGTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 1800}}, true, 0},
	{"timestamp - gt now within - valid (empty)", TestTypeValidate, &cases.TimestampGTNowWithin{}, true, 0},
	{"timestamp - gt now within - invalid (gt)", TestTypeValidate, &cases.TimestampGTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() - 1800}}, false, 1},
	{"timestamp - gt now within - invalid (within)", TestTypeValidate, &cases.TimestampGTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 7200}}, false, 1},
}

var anyCases = []TestCase{
	{"any - none - valid", TestTypeValidate, &cases.AnyNone{Val: &any.Any{}}, true, 0},

	{"any - required - valid", TestTypeValidate, &cases.AnyRequired{Val: &any.Any{}}, true, 0},
	{"any - required - invalid", TestTypeValidate, &cases.AnyRequired{Val: nil}, false, 1},

	{"any - in - valid", TestTypeValidate, &cases.AnyIn{Val: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}}, true, 0},
	{"any - in - valid (empty)", TestTypeValidate, &cases.AnyIn{}, true, 0},
	{"any - in - invalid", TestTypeValidate, &cases.AnyIn{Val: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Timestamp"}}, false, 1},

	{"any - not in - valid", TestTypeValidate, &cases.AnyNotIn{Val: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}}, true, 0},
	{"any - not in - valid (empty)", TestTypeValidate, &cases.AnyNotIn{}, true, 0},
	{"any - not in - invalid", TestTypeValidate, &cases.AnyNotIn{Val: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Timestamp"}}, false, 1},
}

var kitchenSink = []TestCase{
	{"kitchensink - field - valid", TestTypeValidate, &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Const: "abcd", IntConst: 5, BoolConst: false, FloatVal: &wrappers.FloatValue{Value: 1}, DurVal: &duration.Duration{Seconds: 3}, TsVal: &timestamp.Timestamp{Seconds: 17}, FloatConst: 7, DoubleIn: 123, EnumConst: cases.ComplexTestEnum_ComplexTWO, AnyVal: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}, RepTsVal: []*timestamp.Timestamp{{Seconds: 3}}, MapVal: map[int32]string{-1: "a", -2: "b"}, BytesVal: []byte("\x00\x99"), O: &cases.ComplexTestMsg_X{X: "foobar"}}}, true, 0},
	{"kitchensink - valid (unset)", TestTypeValidate, &cases.KitchenSinkMessage{}, true, 0},
	{"kitchensink - field - invalid", TestTypeValidate, &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{}}, false, 1},
	{"kitchensink - field - embedded - invalid", TestTypeValidate, &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Another: &cases.ComplexTestMsg{}}}, false, 1},
	{"kitchensink - field - invalid (transitive)", TestTypeValidate, &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Const: "abcd", BoolConst: true, Nested: &cases.ComplexTestMsg{}}}, false, 1},
	{"kitchensink - field - all errors valid", TestTypeAllErrors, &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Const: "abcd", IntConst: 5, BoolConst: false, FloatVal: &wrappers.FloatValue{Value: 1}, DurVal: &duration.Duration{Seconds: 3}, TsVal: &timestamp.Timestamp{Seconds: 17}, FloatConst: 7, DoubleIn: 123, EnumConst: cases.ComplexTestEnum_ComplexTWO, AnyVal: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}, RepTsVal: []*timestamp.Timestamp{{Seconds: 3}}, MapVal: map[int32]string{-1: "a", -2: "b"}, BytesVal: []byte("\x00\x99"), O: &cases.ComplexTestMsg_X{X: "foobar"}}}, true, 0},
	{"kitchensink - field - all errors invalid", TestTypeAllErrors, &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Const: "abcde", IntConst: 6, BoolConst: true, FloatVal: &wrappers.FloatValue{Value: 0}, DurVal: &duration.Duration{Seconds: 18}, TsVal: &timestamp.Timestamp{Seconds: 6}, FloatConst: 9, DoubleIn: 122, EnumConst: cases.ComplexTestEnum_ComplexONE, AnyVal: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration2"}, RepTsVal: []*timestamp.Timestamp{{Seconds: 0}}, MapVal: map[int32]string{0: "a", 1: "b"}, BytesVal: []byte("\x00\x00"), O: nil}}, false, 1},
	{"kitchensink - complex message - field - all errors invalid", TestTypeAllErrors, &cases.ComplexTestMsg{Const: "abcde", IntConst: 6, BoolConst: true, FloatVal: &wrappers.FloatValue{Value: 0}, DurVal: &duration.Duration{Seconds: 18}, TsVal: &timestamp.Timestamp{Seconds: 6}, FloatConst: 9, DoubleIn: 122, EnumConst: cases.ComplexTestEnum_ComplexONE, AnyVal: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration2"}, RepTsVal: []*timestamp.Timestamp{{Seconds: 0}}, MapVal: map[int32]string{0: "a", 1: "b"}, BytesVal: []byte("\x00\x00"), O: nil}, false, 15},
	{"kitchensink - complex message - field - partial errors invalid", TestTypeAllErrors, &cases.ComplexTestMsg{Const: "abcd", IntConst: 6, BoolConst: false, FloatVal: &wrappers.FloatValue{Value: 0}, DurVal: &duration.Duration{Seconds: 18}, TsVal: &timestamp.Timestamp{Seconds: 6}, FloatConst: 9, DoubleIn: 122, EnumConst: cases.ComplexTestEnum_ComplexONE, AnyVal: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}, RepTsVal: []*timestamp.Timestamp{{Seconds: 0}}, MapVal: map[int32]string{0: "a", 1: "b"}, BytesVal: []byte("\x00\x00"), O: nil}, false, 12},
}