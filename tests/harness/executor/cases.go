package main

import (
	"math"
	"time"

	cases "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go"
	other_package "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/go"
	yet_another_package "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/yet_another_package/go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type TestCase struct {
	Name          string
	Message       proto.Message
	Failures      int // expected number of failed validation errors
	ExpectedRules []string
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
		nestedCases,
	}

	for _, set := range sets {
		TestCases = append(TestCases, set...)
	}
}

var floatCases = []TestCase{
	{"float - none - valid", &cases.FloatNone{Val: -1.23456}, 0, nil},

	{"float - const - valid", &cases.FloatConst{Val: 1.23}, 0, nil},
	{"float - const - invalid", &cases.FloatConst{Val: 4.56}, 1, []string{"float.const"}},

	{"float - in - valid", &cases.FloatIn{Val: 7.89}, 0, nil},
	{"float - in - invalid", &cases.FloatIn{Val: 10.11}, 1, []string{"float.in"}},

	{"float - not in - valid", &cases.FloatNotIn{Val: 1}, 0, nil},
	{"float - not in - invalid", &cases.FloatNotIn{Val: 0}, 1, []string{"float.not_in"}},

	{"float - lt - valid", &cases.FloatLT{Val: -1}, 0, nil},
	{"float - lt - invalid (equal)", &cases.FloatLT{Val: 0}, 1, []string{"float.lt"}},
	{"float - lt - invalid", &cases.FloatLT{Val: 1}, 1, []string{"float.lt"}},

	{"float - lte - valid", &cases.FloatLTE{Val: 63}, 0, nil},
	{"float - lte - valid (equal)", &cases.FloatLTE{Val: 64}, 0, nil},
	{"float - lte - invalid", &cases.FloatLTE{Val: 65}, 1, []string{"float.lte"}},

	{"float - gt - valid", &cases.FloatGT{Val: 17}, 0, nil},
	{"float - gt - invalid (equal)", &cases.FloatGT{Val: 16}, 1, []string{"float.gt"}},
	{"float - gt - invalid", &cases.FloatGT{Val: 15}, 1, []string{"float.gt"}},

	{"float - gte - valid", &cases.FloatGTE{Val: 9}, 0, nil},
	{"float - gte - valid (equal)", &cases.FloatGTE{Val: 8}, 0, nil},
	{"float - gte - invalid", &cases.FloatGTE{Val: 7}, 1, []string{"float.gte"}},

	{"float - gt & lt - valid", &cases.FloatGTLT{Val: 5}, 0, nil},
	{"float - gt & lt - invalid (above)", &cases.FloatGTLT{Val: 11}, 1, []string{"float.in_range_exclusive"}},
	{"float - gt & lt - invalid (below)", &cases.FloatGTLT{Val: -1}, 1, []string{"float.in_range_exclusive"}},
	{"float - gt & lt - invalid (max)", &cases.FloatGTLT{Val: 10}, 1, []string{"float.in_range_exclusive"}},
	{"float - gt & lt - invalid (min)", &cases.FloatGTLT{Val: 0}, 1, []string{"float.in_range_exclusive"}},

	{"float - exclusive gt & lt - valid (above)", &cases.FloatExLTGT{Val: 11}, 0, nil},
	{"float - exclusive gt & lt - valid (below)", &cases.FloatExLTGT{Val: -1}, 0, nil},
	{"float - exclusive gt & lt - invalid", &cases.FloatExLTGT{Val: 5}, 1, []string{"float.out_of_range"}},
	{"float - exclusive gt & lt - invalid (max)", &cases.FloatExLTGT{Val: 10}, 1, []string{"float.out_of_range"}},
	{"float - exclusive gt & lt - invalid (min)", &cases.FloatExLTGT{Val: 0}, 1, []string{"float.out_of_range"}},

	{"float - gte & lte - valid", &cases.FloatGTELTE{Val: 200}, 0, nil},
	{"float - gte & lte - valid (max)", &cases.FloatGTELTE{Val: 256}, 0, nil},
	{"float - gte & lte - valid (min)", &cases.FloatGTELTE{Val: 128}, 0, nil},
	{"float - gte & lte - invalid (above)", &cases.FloatGTELTE{Val: 300}, 1, []string{"float.in_range"}},
	{"float - gte & lte - invalid (below)", &cases.FloatGTELTE{Val: 100}, 1, []string{"float.in_range"}},

	{"float - exclusive gte & lte - valid (above)", &cases.FloatExGTELTE{Val: 300}, 0, nil},
	{"float - exclusive gte & lte - valid (below)", &cases.FloatExGTELTE{Val: 100}, 0, nil},
	{"float - exclusive gte & lte - valid (max)", &cases.FloatExGTELTE{Val: 256}, 0, nil},
	{"float - exclusive gte & lte - valid (min)", &cases.FloatExGTELTE{Val: 128}, 0, nil},
	{"float - exclusive gte & lte - invalid", &cases.FloatExGTELTE{Val: 200}, 1, []string{"float.out_of_range_inclusive"}},

	{"float - ignore_empty gte & lte - valid", &cases.FloatIgnore{Val: 0}, 0, nil},
}

var doubleCases = []TestCase{
	{"double - none - valid", &cases.DoubleNone{Val: -1.23456}, 0, nil},

	{"double - const - valid", &cases.DoubleConst{Val: 1.23}, 0, nil},
	{"double - const - invalid", &cases.DoubleConst{Val: 4.56}, 1, []string{"double.const"}},

	{"double - in - valid", &cases.DoubleIn{Val: 7.89}, 0, nil},
	{"double - in - invalid", &cases.DoubleIn{Val: 10.11}, 1, []string{"double.in"}},

	{"double - not in - valid", &cases.DoubleNotIn{Val: 1}, 0, nil},
	{"double - not in - invalid", &cases.DoubleNotIn{Val: 0}, 1, []string{"double.not_in"}},

	{"double - lt - valid", &cases.DoubleLT{Val: -1}, 0, nil},
	{"double - lt - invalid (equal)", &cases.DoubleLT{Val: 0}, 1, []string{"double.lt"}},
	{"double - lt - invalid", &cases.DoubleLT{Val: 1}, 1, []string{"doublt.lt"}},

	{"double - lte - valid", &cases.DoubleLTE{Val: 63}, 0, nil},
	{"double - lte - valid (equal)", &cases.DoubleLTE{Val: 64}, 0, nil},
	{"double - lte - invalid", &cases.DoubleLTE{Val: 65}, 1, []string{"double.lte"}},

	{"double - gt - valid", &cases.DoubleGT{Val: 17}, 0, nil},
	{"double - gt - invalid (equal)", &cases.DoubleGT{Val: 16}, 1, []string{"double.gt"}},
	{"double - gt - invalid", &cases.DoubleGT{Val: 15}, 1, []string{"double.gt"}},

	{"double - gte - valid", &cases.DoubleGTE{Val: 9}, 0, nil},
	{"double - gte - valid (equal)", &cases.DoubleGTE{Val: 8}, 0, nil},
	{"double - gte - invalid", &cases.DoubleGTE{Val: 7}, 1, []string{"double.gte"}},

	{"double - gt & lt - valid", &cases.DoubleGTLT{Val: 5}, 0, nil},
	{"double - gt & lt - invalid (above)", &cases.DoubleGTLT{Val: 11}, 1, []string{"double.in_range_exclusive"}},
	{"double - gt & lt - invalid (below)", &cases.DoubleGTLT{Val: -1}, 1, []string{"double.in_range_exclusive"}},
	{"double - gt & lt - invalid (max)", &cases.DoubleGTLT{Val: 10}, 1, []string{"double.in_range_exclusive"}},
	{"double - gt & lt - invalid (min)", &cases.DoubleGTLT{Val: 0}, 1, []string{"double.in_range_exclusive"}},

	{"double - exclusive gt & lt - valid (above)", &cases.DoubleExLTGT{Val: 11}, 0, nil},
	{"double - exclusive gt & lt - valid (below)", &cases.DoubleExLTGT{Val: -1}, 0, nil},
	{"double - exclusive gt & lt - invalid", &cases.DoubleExLTGT{Val: 5}, 1, []string{"double.out_of_range"}},
	{"double - exclusive gt & lt - invalid (max)", &cases.DoubleExLTGT{Val: 10}, 1, []string{"double.out_of_range"}},
	{"double - exclusive gt & lt - invalid (min)", &cases.DoubleExLTGT{Val: 0}, 1, []string{"double.out_of_range"}},

	{"double - gte & lte - valid", &cases.DoubleGTELTE{Val: 200}, 0, nil},
	{"double - gte & lte - valid (max)", &cases.DoubleGTELTE{Val: 256}, 0, nil},
	{"double - gte & lte - valid (min)", &cases.DoubleGTELTE{Val: 128}, 0, nil},
	{"double - gte & lte - invalid (above)", &cases.DoubleGTELTE{Val: 300}, 1, []string{"double.in_range"}},
	{"double - gte & lte - invalid (below)", &cases.DoubleGTELTE{Val: 100}, 1, []string{"double.in_range"}},

	{"double - exclusive gte & lte - valid (above)", &cases.DoubleExGTELTE{Val: 300}, 0, nil},
	{"double - exclusive gte & lte - valid (below)", &cases.DoubleExGTELTE{Val: 100}, 0, nil},
	{"double - exclusive gte & lte - valid (max)", &cases.DoubleExGTELTE{Val: 256}, 0, nil},
	{"double - exclusive gte & lte - valid (min)", &cases.DoubleExGTELTE{Val: 128}, 0, nil},
	{"double - exclusive gte & lte - invalid", &cases.DoubleExGTELTE{Val: 200}, 1, []string{"double.out_of_range_inclusive"}},

	{"double - ignore_empty gte & lte - valid", &cases.DoubleIgnore{Val: 0}, 0, nil},
}

var int32Cases = []TestCase{
	{"int32 - none - valid", &cases.Int32None{Val: 123}, 0, nil},

	{"int32 - const - valid", &cases.Int32Const{Val: 1}, 0, nil},
	{"int32 - const - invalid", &cases.Int32Const{Val: 2}, 1, nil},

	{"int32 - in - valid", &cases.Int32In{Val: 3}, 0, nil},
	{"int32 - in - invalid", &cases.Int32In{Val: 5}, 1, nil},

	{"int32 - not in - valid", &cases.Int32NotIn{Val: 1}, 0, nil},
	{"int32 - not in - invalid", &cases.Int32NotIn{Val: 0}, 1, nil},

	{"int32 - lt - valid", &cases.Int32LT{Val: -1}, 0, nil},
	{"int32 - lt - invalid (equal)", &cases.Int32LT{Val: 0}, 1, nil},
	{"int32 - lt - invalid", &cases.Int32LT{Val: 1}, 1, nil},

	{"int32 - lte - valid", &cases.Int32LTE{Val: 63}, 0, nil},
	{"int32 - lte - valid (equal)", &cases.Int32LTE{Val: 64}, 0, nil},
	{"int32 - lte - invalid", &cases.Int32LTE{Val: 65}, 1, nil},

	{"int32 - gt - valid", &cases.Int32GT{Val: 17}, 0, nil},
	{"int32 - gt - invalid (equal)", &cases.Int32GT{Val: 16}, 1, nil},
	{"int32 - gt - invalid", &cases.Int32GT{Val: 15}, 1, nil},

	{"int32 - gte - valid", &cases.Int32GTE{Val: 9}, 0, nil},
	{"int32 - gte - valid (equal)", &cases.Int32GTE{Val: 8}, 0, nil},
	{"int32 - gte - invalid", &cases.Int32GTE{Val: 7}, 1, nil},

	{"int32 - gt & lt - valid", &cases.Int32GTLT{Val: 5}, 0, nil},
	{"int32 - gt & lt - invalid (above)", &cases.Int32GTLT{Val: 11}, 1, nil},
	{"int32 - gt & lt - invalid (below)", &cases.Int32GTLT{Val: -1}, 1, nil},
	{"int32 - gt & lt - invalid (max)", &cases.Int32GTLT{Val: 10}, 1, nil},
	{"int32 - gt & lt - invalid (min)", &cases.Int32GTLT{Val: 0}, 1, nil},

	{"int32 - exclusive gt & lt - valid (above)", &cases.Int32ExLTGT{Val: 11}, 0, nil},
	{"int32 - exclusive gt & lt - valid (below)", &cases.Int32ExLTGT{Val: -1}, 0, nil},
	{"int32 - exclusive gt & lt - invalid", &cases.Int32ExLTGT{Val: 5}, 1, nil},
	{"int32 - exclusive gt & lt - invalid (max)", &cases.Int32ExLTGT{Val: 10}, 1, nil},
	{"int32 - exclusive gt & lt - invalid (min)", &cases.Int32ExLTGT{Val: 0}, 1, nil},

	{"int32 - gte & lte - valid", &cases.Int32GTELTE{Val: 200}, 0, nil},
	{"int32 - gte & lte - valid (max)", &cases.Int32GTELTE{Val: 256}, 0, nil},
	{"int32 - gte & lte - valid (min)", &cases.Int32GTELTE{Val: 128}, 0, nil},
	{"int32 - gte & lte - invalid (above)", &cases.Int32GTELTE{Val: 300}, 1, nil},
	{"int32 - gte & lte - invalid (below)", &cases.Int32GTELTE{Val: 100}, 1, nil},

	{"int32 - exclusive gte & lte - valid (above)", &cases.Int32ExGTELTE{Val: 300}, 0, nil},
	{"int32 - exclusive gte & lte - valid (below)", &cases.Int32ExGTELTE{Val: 100}, 0, nil},
	{"int32 - exclusive gte & lte - valid (max)", &cases.Int32ExGTELTE{Val: 256}, 0, nil},
	{"int32 - exclusive gte & lte - valid (min)", &cases.Int32ExGTELTE{Val: 128}, 0, nil},
	{"int32 - exclusive gte & lte - invalid", &cases.Int32ExGTELTE{Val: 200}, 1, nil},

	{"int32 - ignore_empty gte & lte - valid", &cases.Int32Ignore{Val: 0}, 0, nil},
}

var int64Cases = []TestCase{
	{"int64 - none - valid", &cases.Int64None{Val: 123}, 0, nil},

	{"int64 - const - valid", &cases.Int64Const{Val: 1}, 0, nil},
	{"int64 - const - invalid", &cases.Int64Const{Val: 2}, 1, nil},

	{"int64 - in - valid", &cases.Int64In{Val: 3}, 0, nil},
	{"int64 - in - invalid", &cases.Int64In{Val: 5}, 1, nil},

	{"int64 - not in - valid", &cases.Int64NotIn{Val: 1}, 0, nil},
	{"int64 - not in - invalid", &cases.Int64NotIn{Val: 0}, 1, nil},

	{"int64 - lt - valid", &cases.Int64LT{Val: -1}, 0, nil},
	{"int64 - lt - invalid (equal)", &cases.Int64LT{Val: 0}, 1, nil},
	{"int64 - lt - invalid", &cases.Int64LT{Val: 1}, 1, nil},

	{"int64 - lte - valid", &cases.Int64LTE{Val: 63}, 0, nil},
	{"int64 - lte - valid (equal)", &cases.Int64LTE{Val: 64}, 0, nil},
	{"int64 - lte - invalid", &cases.Int64LTE{Val: 65}, 1, nil},

	{"int64 - gt - valid", &cases.Int64GT{Val: 17}, 0, nil},
	{"int64 - gt - invalid (equal)", &cases.Int64GT{Val: 16}, 1, nil},
	{"int64 - gt - invalid", &cases.Int64GT{Val: 15}, 1, nil},

	{"int64 - gte - valid", &cases.Int64GTE{Val: 9}, 0, nil},
	{"int64 - gte - valid (equal)", &cases.Int64GTE{Val: 8}, 0, nil},
	{"int64 - gte - invalid", &cases.Int64GTE{Val: 7}, 1, nil},

	{"int64 - gt & lt - valid", &cases.Int64GTLT{Val: 5}, 0, nil},
	{"int64 - gt & lt - invalid (above)", &cases.Int64GTLT{Val: 11}, 1, nil},
	{"int64 - gt & lt - invalid (below)", &cases.Int64GTLT{Val: -1}, 1, nil},
	{"int64 - gt & lt - invalid (max)", &cases.Int64GTLT{Val: 10}, 1, nil},
	{"int64 - gt & lt - invalid (min)", &cases.Int64GTLT{Val: 0}, 1, nil},

	{"int64 - exclusive gt & lt - valid (above)", &cases.Int64ExLTGT{Val: 11}, 0, nil},
	{"int64 - exclusive gt & lt - valid (below)", &cases.Int64ExLTGT{Val: -1}, 0, nil},
	{"int64 - exclusive gt & lt - invalid", &cases.Int64ExLTGT{Val: 5}, 1, nil},
	{"int64 - exclusive gt & lt - invalid (max)", &cases.Int64ExLTGT{Val: 10}, 1, nil},
	{"int64 - exclusive gt & lt - invalid (min)", &cases.Int64ExLTGT{Val: 0}, 1, nil},

	{"int64 - gte & lte - valid", &cases.Int64GTELTE{Val: 200}, 0, nil},
	{"int64 - gte & lte - valid (max)", &cases.Int64GTELTE{Val: 256}, 0, nil},
	{"int64 - gte & lte - valid (min)", &cases.Int64GTELTE{Val: 128}, 0, nil},
	{"int64 - gte & lte - invalid (above)", &cases.Int64GTELTE{Val: 300}, 1, nil},
	{"int64 - gte & lte - invalid (below)", &cases.Int64GTELTE{Val: 100}, 1, nil},

	{"int64 - exclusive gte & lte - valid (above)", &cases.Int64ExGTELTE{Val: 300}, 0, nil},
	{"int64 - exclusive gte & lte - valid (below)", &cases.Int64ExGTELTE{Val: 100}, 0, nil},
	{"int64 - exclusive gte & lte - valid (max)", &cases.Int64ExGTELTE{Val: 256}, 0, nil},
	{"int64 - exclusive gte & lte - valid (min)", &cases.Int64ExGTELTE{Val: 128}, 0, nil},
	{"int64 - exclusive gte & lte - invalid", &cases.Int64ExGTELTE{Val: 200}, 1, nil},

	{"int64 - ignore_empty gte & lte - valid", &cases.Int64Ignore{Val: 0}, 0, nil},

	{"int64 optional - lte - valid", &cases.Int64LTEOptional{Val: &wrapperspb.Int64(63).Value}, 0, nil},
	{"int64 optional - lte - valid (equal)", &cases.Int64LTEOptional{Val: &wrapperspb.Int64(64).Value}, 0, nil},
	{"int64 optional - lte - valid (unset)", &cases.Int64LTEOptional{}, 0, nil},
}

var uint32Cases = []TestCase{
	{"uint32 - none - valid", &cases.UInt32None{Val: 123}, 0, nil},

	{"uint32 - const - valid", &cases.UInt32Const{Val: 1}, 0, nil},
	{"uint32 - const - invalid", &cases.UInt32Const{Val: 2}, 1, nil},

	{"uint32 - in - valid", &cases.UInt32In{Val: 3}, 0, nil},
	{"uint32 - in - invalid", &cases.UInt32In{Val: 5}, 1, nil},

	{"uint32 - not in - valid", &cases.UInt32NotIn{Val: 1}, 0, nil},
	{"uint32 - not in - invalid", &cases.UInt32NotIn{Val: 0}, 1, nil},

	{"uint32 - lt - valid", &cases.UInt32LT{Val: 4}, 0, nil},
	{"uint32 - lt - invalid (equal)", &cases.UInt32LT{Val: 5}, 1, nil},
	{"uint32 - lt - invalid", &cases.UInt32LT{Val: 6}, 1, nil},

	{"uint32 - lte - valid", &cases.UInt32LTE{Val: 63}, 0, nil},
	{"uint32 - lte - valid (equal)", &cases.UInt32LTE{Val: 64}, 0, nil},
	{"uint32 - lte - invalid", &cases.UInt32LTE{Val: 65}, 1, nil},

	{"uint32 - gt - valid", &cases.UInt32GT{Val: 17}, 0, nil},
	{"uint32 - gt - invalid (equal)", &cases.UInt32GT{Val: 16}, 1, nil},
	{"uint32 - gt - invalid", &cases.UInt32GT{Val: 15}, 1, nil},

	{"uint32 - gte - valid", &cases.UInt32GTE{Val: 9}, 0, nil},
	{"uint32 - gte - valid (equal)", &cases.UInt32GTE{Val: 8}, 0, nil},
	{"uint32 - gte - invalid", &cases.UInt32GTE{Val: 7}, 1, nil},

	{"uint32 - gt & lt - valid", &cases.UInt32GTLT{Val: 7}, 0, nil},
	{"uint32 - gt & lt - invalid (above)", &cases.UInt32GTLT{Val: 11}, 1, nil},
	{"uint32 - gt & lt - invalid (below)", &cases.UInt32GTLT{Val: 1}, 1, nil},
	{"uint32 - gt & lt - invalid (max)", &cases.UInt32GTLT{Val: 10}, 1, nil},
	{"uint32 - gt & lt - invalid (min)", &cases.UInt32GTLT{Val: 5}, 1, nil},

	{"uint32 - exclusive gt & lt - valid (above)", &cases.UInt32ExLTGT{Val: 11}, 0, nil},
	{"uint32 - exclusive gt & lt - valid (below)", &cases.UInt32ExLTGT{Val: 4}, 0, nil},
	{"uint32 - exclusive gt & lt - invalid", &cases.UInt32ExLTGT{Val: 7}, 1, nil},
	{"uint32 - exclusive gt & lt - invalid (max)", &cases.UInt32ExLTGT{Val: 10}, 1, nil},
	{"uint32 - exclusive gt & lt - invalid (min)", &cases.UInt32ExLTGT{Val: 5}, 1, nil},

	{"uint32 - gte & lte - valid", &cases.UInt32GTELTE{Val: 200}, 0, nil},
	{"uint32 - gte & lte - valid (max)", &cases.UInt32GTELTE{Val: 256}, 0, nil},
	{"uint32 - gte & lte - valid (min)", &cases.UInt32GTELTE{Val: 128}, 0, nil},
	{"uint32 - gte & lte - invalid (above)", &cases.UInt32GTELTE{Val: 300}, 1, nil},
	{"uint32 - gte & lte - invalid (below)", &cases.UInt32GTELTE{Val: 100}, 1, nil},

	{"uint32 - exclusive gte & lte - valid (above)", &cases.UInt32ExGTELTE{Val: 300}, 0, nil},
	{"uint32 - exclusive gte & lte - valid (below)", &cases.UInt32ExGTELTE{Val: 100}, 0, nil},
	{"uint32 - exclusive gte & lte - valid (max)", &cases.UInt32ExGTELTE{Val: 256}, 0, nil},
	{"uint32 - exclusive gte & lte - valid (min)", &cases.UInt32ExGTELTE{Val: 128}, 0, nil},
	{"uint32 - exclusive gte & lte - invalid", &cases.UInt32ExGTELTE{Val: 200}, 1, nil},

	{"uint32 - ignore_empty gte & lte - valid", &cases.UInt32Ignore{Val: 0}, 0, nil},
}

var uint64Cases = []TestCase{
	{"uint64 - none - valid", &cases.UInt64None{Val: 123}, 0, nil},

	{"uint64 - const - valid", &cases.UInt64Const{Val: 1}, 0, nil},
	{"uint64 - const - invalid", &cases.UInt64Const{Val: 2}, 1, nil},

	{"uint64 - in - valid", &cases.UInt64In{Val: 3}, 0, nil},
	{"uint64 - in - invalid", &cases.UInt64In{Val: 5}, 1, nil},

	{"uint64 - not in - valid", &cases.UInt64NotIn{Val: 1}, 0, nil},
	{"uint64 - not in - invalid", &cases.UInt64NotIn{Val: 0}, 1, nil},

	{"uint64 - lt - valid", &cases.UInt64LT{Val: 4}, 0, nil},
	{"uint64 - lt - invalid (equal)", &cases.UInt64LT{Val: 5}, 1, nil},
	{"uint64 - lt - invalid", &cases.UInt64LT{Val: 6}, 1, nil},

	{"uint64 - lte - valid", &cases.UInt64LTE{Val: 63}, 0, nil},
	{"uint64 - lte - valid (equal)", &cases.UInt64LTE{Val: 64}, 0, nil},
	{"uint64 - lte - invalid", &cases.UInt64LTE{Val: 65}, 1, nil},

	{"uint64 - gt - valid", &cases.UInt64GT{Val: 17}, 0, nil},
	{"uint64 - gt - invalid (equal)", &cases.UInt64GT{Val: 16}, 1, nil},
	{"uint64 - gt - invalid", &cases.UInt64GT{Val: 15}, 1, nil},

	{"uint64 - gte - valid", &cases.UInt64GTE{Val: 9}, 0, nil},
	{"uint64 - gte - valid (equal)", &cases.UInt64GTE{Val: 8}, 0, nil},
	{"uint64 - gte - invalid", &cases.UInt64GTE{Val: 7}, 1, nil},

	{"uint64 - gt & lt - valid", &cases.UInt64GTLT{Val: 7}, 0, nil},
	{"uint64 - gt & lt - invalid (above)", &cases.UInt64GTLT{Val: 11}, 1, nil},
	{"uint64 - gt & lt - invalid (below)", &cases.UInt64GTLT{Val: 1}, 1, nil},
	{"uint64 - gt & lt - invalid (max)", &cases.UInt64GTLT{Val: 10}, 1, nil},
	{"uint64 - gt & lt - invalid (min)", &cases.UInt64GTLT{Val: 5}, 1, nil},

	{"uint64 - exclusive gt & lt - valid (above)", &cases.UInt64ExLTGT{Val: 11}, 0, nil},
	{"uint64 - exclusive gt & lt - valid (below)", &cases.UInt64ExLTGT{Val: 4}, 0, nil},
	{"uint64 - exclusive gt & lt - invalid", &cases.UInt64ExLTGT{Val: 7}, 1, nil},
	{"uint64 - exclusive gt & lt - invalid (max)", &cases.UInt64ExLTGT{Val: 10}, 1, nil},
	{"uint64 - exclusive gt & lt - invalid (min)", &cases.UInt64ExLTGT{Val: 5}, 1, nil},

	{"uint64 - gte & lte - valid", &cases.UInt64GTELTE{Val: 200}, 0, nil},
	{"uint64 - gte & lte - valid (max)", &cases.UInt64GTELTE{Val: 256}, 0, nil},
	{"uint64 - gte & lte - valid (min)", &cases.UInt64GTELTE{Val: 128}, 0, nil},
	{"uint64 - gte & lte - invalid (above)", &cases.UInt64GTELTE{Val: 300}, 1, nil},
	{"uint64 - gte & lte - invalid (below)", &cases.UInt64GTELTE{Val: 100}, 1, nil},

	{"uint64 - exclusive gte & lte - valid (above)", &cases.UInt64ExGTELTE{Val: 300}, 0, nil},
	{"uint64 - exclusive gte & lte - valid (below)", &cases.UInt64ExGTELTE{Val: 100}, 0, nil},
	{"uint64 - exclusive gte & lte - valid (max)", &cases.UInt64ExGTELTE{Val: 256}, 0, nil},
	{"uint64 - exclusive gte & lte - valid (min)", &cases.UInt64ExGTELTE{Val: 128}, 0, nil},
	{"uint64 - exclusive gte & lte - invalid", &cases.UInt64ExGTELTE{Val: 200}, 1, nil},

	{"uint64 - ignore_empty gte & lte - valid", &cases.UInt64Ignore{Val: 0}, 0, nil},
}

var sint32Cases = []TestCase{
	{"sint32 - none - valid", &cases.SInt32None{Val: 123}, 0, nil},

	{"sint32 - const - valid", &cases.SInt32Const{Val: 1}, 0, nil},
	{"sint32 - const - invalid", &cases.SInt32Const{Val: 2}, 1, nil},

	{"sint32 - in - valid", &cases.SInt32In{Val: 3}, 0, nil},
	{"sint32 - in - invalid", &cases.SInt32In{Val: 5}, 1, nil},

	{"sint32 - not in - valid", &cases.SInt32NotIn{Val: 1}, 0, nil},
	{"sint32 - not in - invalid", &cases.SInt32NotIn{Val: 0}, 1, nil},

	{"sint32 - lt - valid", &cases.SInt32LT{Val: -1}, 0, nil},
	{"sint32 - lt - invalid (equal)", &cases.SInt32LT{Val: 0}, 1, nil},
	{"sint32 - lt - invalid", &cases.SInt32LT{Val: 1}, 1, nil},

	{"sint32 - lte - valid", &cases.SInt32LTE{Val: 63}, 0, nil},
	{"sint32 - lte - valid (equal)", &cases.SInt32LTE{Val: 64}, 0, nil},
	{"sint32 - lte - invalid", &cases.SInt32LTE{Val: 65}, 1, nil},

	{"sint32 - gt - valid", &cases.SInt32GT{Val: 17}, 0, nil},
	{"sint32 - gt - invalid (equal)", &cases.SInt32GT{Val: 16}, 1, nil},
	{"sint32 - gt - invalid", &cases.SInt32GT{Val: 15}, 1, nil},

	{"sint32 - gte - valid", &cases.SInt32GTE{Val: 9}, 0, nil},
	{"sint32 - gte - valid (equal)", &cases.SInt32GTE{Val: 8}, 0, nil},
	{"sint32 - gte - invalid", &cases.SInt32GTE{Val: 7}, 1, nil},

	{"sint32 - gt & lt - valid", &cases.SInt32GTLT{Val: 5}, 0, nil},
	{"sint32 - gt & lt - invalid (above)", &cases.SInt32GTLT{Val: 11}, 1, nil},
	{"sint32 - gt & lt - invalid (below)", &cases.SInt32GTLT{Val: -1}, 1, nil},
	{"sint32 - gt & lt - invalid (max)", &cases.SInt32GTLT{Val: 10}, 1, nil},
	{"sint32 - gt & lt - invalid (min)", &cases.SInt32GTLT{Val: 0}, 1, nil},

	{"sint32 - exclusive gt & lt - valid (above)", &cases.SInt32ExLTGT{Val: 11}, 0, nil},
	{"sint32 - exclusive gt & lt - valid (below)", &cases.SInt32ExLTGT{Val: -1}, 0, nil},
	{"sint32 - exclusive gt & lt - invalid", &cases.SInt32ExLTGT{Val: 5}, 1, nil},
	{"sint32 - exclusive gt & lt - invalid (max)", &cases.SInt32ExLTGT{Val: 10}, 1, nil},
	{"sint32 - exclusive gt & lt - invalid (min)", &cases.SInt32ExLTGT{Val: 0}, 1, nil},

	{"sint32 - gte & lte - valid", &cases.SInt32GTELTE{Val: 200}, 0, nil},
	{"sint32 - gte & lte - valid (max)", &cases.SInt32GTELTE{Val: 256}, 0, nil},
	{"sint32 - gte & lte - valid (min)", &cases.SInt32GTELTE{Val: 128}, 0, nil},
	{"sint32 - gte & lte - invalid (above)", &cases.SInt32GTELTE{Val: 300}, 1, nil},
	{"sint32 - gte & lte - invalid (below)", &cases.SInt32GTELTE{Val: 100}, 1, nil},

	{"sint32 - exclusive gte & lte - valid (above)", &cases.SInt32ExGTELTE{Val: 300}, 0, nil},
	{"sint32 - exclusive gte & lte - valid (below)", &cases.SInt32ExGTELTE{Val: 100}, 0, nil},
	{"sint32 - exclusive gte & lte - valid (max)", &cases.SInt32ExGTELTE{Val: 256}, 0, nil},
	{"sint32 - exclusive gte & lte - valid (min)", &cases.SInt32ExGTELTE{Val: 128}, 0, nil},
	{"sint32 - exclusive gte & lte - invalid", &cases.SInt32ExGTELTE{Val: 200}, 1, nil},

	{"sint32 - ignore_empty gte & lte - valid", &cases.SInt32Ignore{Val: 0}, 0, nil},
}

var sint64Cases = []TestCase{
	{"sint64 - none - valid", &cases.SInt64None{Val: 123}, 0, nil},

	{"sint64 - const - valid", &cases.SInt64Const{Val: 1}, 0, nil},
	{"sint64 - const - invalid", &cases.SInt64Const{Val: 2}, 1, nil},

	{"sint64 - in - valid", &cases.SInt64In{Val: 3}, 0, nil},
	{"sint64 - in - invalid", &cases.SInt64In{Val: 5}, 1, nil},

	{"sint64 - not in - valid", &cases.SInt64NotIn{Val: 1}, 0, nil},
	{"sint64 - not in - invalid", &cases.SInt64NotIn{Val: 0}, 1, nil},

	{"sint64 - lt - valid", &cases.SInt64LT{Val: -1}, 0, nil},
	{"sint64 - lt - invalid (equal)", &cases.SInt64LT{Val: 0}, 1, nil},
	{"sint64 - lt - invalid", &cases.SInt64LT{Val: 1}, 1, nil},

	{"sint64 - lte - valid", &cases.SInt64LTE{Val: 63}, 0, nil},
	{"sint64 - lte - valid (equal)", &cases.SInt64LTE{Val: 64}, 0, nil},
	{"sint64 - lte - invalid", &cases.SInt64LTE{Val: 65}, 1, nil},

	{"sint64 - gt - valid", &cases.SInt64GT{Val: 17}, 0, nil},
	{"sint64 - gt - invalid (equal)", &cases.SInt64GT{Val: 16}, 1, nil},
	{"sint64 - gt - invalid", &cases.SInt64GT{Val: 15}, 1, nil},

	{"sint64 - gte - valid", &cases.SInt64GTE{Val: 9}, 0, nil},
	{"sint64 - gte - valid (equal)", &cases.SInt64GTE{Val: 8}, 0, nil},
	{"sint64 - gte - invalid", &cases.SInt64GTE{Val: 7}, 1, nil},

	{"sint64 - gt & lt - valid", &cases.SInt64GTLT{Val: 5}, 0, nil},
	{"sint64 - gt & lt - invalid (above)", &cases.SInt64GTLT{Val: 11}, 1, nil},
	{"sint64 - gt & lt - invalid (below)", &cases.SInt64GTLT{Val: -1}, 1, nil},
	{"sint64 - gt & lt - invalid (max)", &cases.SInt64GTLT{Val: 10}, 1, nil},
	{"sint64 - gt & lt - invalid (min)", &cases.SInt64GTLT{Val: 0}, 1, nil},

	{"sint64 - exclusive gt & lt - valid (above)", &cases.SInt64ExLTGT{Val: 11}, 0, nil},
	{"sint64 - exclusive gt & lt - valid (below)", &cases.SInt64ExLTGT{Val: -1}, 0, nil},
	{"sint64 - exclusive gt & lt - invalid", &cases.SInt64ExLTGT{Val: 5}, 1, nil},
	{"sint64 - exclusive gt & lt - invalid (max)", &cases.SInt64ExLTGT{Val: 10}, 1, nil},
	{"sint64 - exclusive gt & lt - invalid (min)", &cases.SInt64ExLTGT{Val: 0}, 1, nil},

	{"sint64 - gte & lte - valid", &cases.SInt64GTELTE{Val: 200}, 0, nil},
	{"sint64 - gte & lte - valid (max)", &cases.SInt64GTELTE{Val: 256}, 0, nil},
	{"sint64 - gte & lte - valid (min)", &cases.SInt64GTELTE{Val: 128}, 0, nil},
	{"sint64 - gte & lte - invalid (above)", &cases.SInt64GTELTE{Val: 300}, 1, nil},
	{"sint64 - gte & lte - invalid (below)", &cases.SInt64GTELTE{Val: 100}, 1, nil},

	{"sint64 - exclusive gte & lte - valid (above)", &cases.SInt64ExGTELTE{Val: 300}, 0, nil},
	{"sint64 - exclusive gte & lte - valid (below)", &cases.SInt64ExGTELTE{Val: 100}, 0, nil},
	{"sint64 - exclusive gte & lte - valid (max)", &cases.SInt64ExGTELTE{Val: 256}, 0, nil},
	{"sint64 - exclusive gte & lte - valid (min)", &cases.SInt64ExGTELTE{Val: 128}, 0, nil},
	{"sint64 - exclusive gte & lte - invalid", &cases.SInt64ExGTELTE{Val: 200}, 1, nil},

	{"sint64 - ignore_empty gte & lte - valid", &cases.SInt64Ignore{Val: 0}, 0, nil},
}

var fixed32Cases = []TestCase{
	{"fixed32 - none - valid", &cases.Fixed32None{Val: 123}, 0, nil},

	{"fixed32 - const - valid", &cases.Fixed32Const{Val: 1}, 0, nil},
	{"fixed32 - const - invalid", &cases.Fixed32Const{Val: 2}, 1, nil},

	{"fixed32 - in - valid", &cases.Fixed32In{Val: 3}, 0, nil},
	{"fixed32 - in - invalid", &cases.Fixed32In{Val: 5}, 1, nil},

	{"fixed32 - not in - valid", &cases.Fixed32NotIn{Val: 1}, 0, nil},
	{"fixed32 - not in - invalid", &cases.Fixed32NotIn{Val: 0}, 1, nil},

	{"fixed32 - lt - valid", &cases.Fixed32LT{Val: 4}, 0, nil},
	{"fixed32 - lt - invalid (equal)", &cases.Fixed32LT{Val: 5}, 1, nil},
	{"fixed32 - lt - invalid", &cases.Fixed32LT{Val: 6}, 1, nil},

	{"fixed32 - lte - valid", &cases.Fixed32LTE{Val: 63}, 0, nil},
	{"fixed32 - lte - valid (equal)", &cases.Fixed32LTE{Val: 64}, 0, nil},
	{"fixed32 - lte - invalid", &cases.Fixed32LTE{Val: 65}, 1, nil},

	{"fixed32 - gt - valid", &cases.Fixed32GT{Val: 17}, 0, nil},
	{"fixed32 - gt - invalid (equal)", &cases.Fixed32GT{Val: 16}, 1, nil},
	{"fixed32 - gt - invalid", &cases.Fixed32GT{Val: 15}, 1, nil},

	{"fixed32 - gte - valid", &cases.Fixed32GTE{Val: 9}, 0, nil},
	{"fixed32 - gte - valid (equal)", &cases.Fixed32GTE{Val: 8}, 0, nil},
	{"fixed32 - gte - invalid", &cases.Fixed32GTE{Val: 7}, 1, nil},

	{"fixed32 - gt & lt - valid", &cases.Fixed32GTLT{Val: 7}, 0, nil},
	{"fixed32 - gt & lt - invalid (above)", &cases.Fixed32GTLT{Val: 11}, 1, nil},
	{"fixed32 - gt & lt - invalid (below)", &cases.Fixed32GTLT{Val: 1}, 1, nil},
	{"fixed32 - gt & lt - invalid (max)", &cases.Fixed32GTLT{Val: 10}, 1, nil},
	{"fixed32 - gt & lt - invalid (min)", &cases.Fixed32GTLT{Val: 5}, 1, nil},

	{"fixed32 - exclusive gt & lt - valid (above)", &cases.Fixed32ExLTGT{Val: 11}, 0, nil},
	{"fixed32 - exclusive gt & lt - valid (below)", &cases.Fixed32ExLTGT{Val: 4}, 0, nil},
	{"fixed32 - exclusive gt & lt - invalid", &cases.Fixed32ExLTGT{Val: 7}, 1, nil},
	{"fixed32 - exclusive gt & lt - invalid (max)", &cases.Fixed32ExLTGT{Val: 10}, 1, nil},
	{"fixed32 - exclusive gt & lt - invalid (min)", &cases.Fixed32ExLTGT{Val: 5}, 1, nil},

	{"fixed32 - gte & lte - valid", &cases.Fixed32GTELTE{Val: 200}, 0, nil},
	{"fixed32 - gte & lte - valid (max)", &cases.Fixed32GTELTE{Val: 256}, 0, nil},
	{"fixed32 - gte & lte - valid (min)", &cases.Fixed32GTELTE{Val: 128}, 0, nil},
	{"fixed32 - gte & lte - invalid (above)", &cases.Fixed32GTELTE{Val: 300}, 1, nil},
	{"fixed32 - gte & lte - invalid (below)", &cases.Fixed32GTELTE{Val: 100}, 1, nil},

	{"fixed32 - exclusive gte & lte - valid (above)", &cases.Fixed32ExGTELTE{Val: 300}, 0, nil},
	{"fixed32 - exclusive gte & lte - valid (below)", &cases.Fixed32ExGTELTE{Val: 100}, 0, nil},
	{"fixed32 - exclusive gte & lte - valid (max)", &cases.Fixed32ExGTELTE{Val: 256}, 0, nil},
	{"fixed32 - exclusive gte & lte - valid (min)", &cases.Fixed32ExGTELTE{Val: 128}, 0, nil},
	{"fixed32 - exclusive gte & lte - invalid", &cases.Fixed32ExGTELTE{Val: 200}, 1, nil},

	{"fixed32 - ignore_empty gte & lte - valid", &cases.Fixed32Ignore{Val: 0}, 0, nil},
}

var fixed64Cases = []TestCase{
	{"fixed64 - none - valid", &cases.Fixed64None{Val: 123}, 0, nil},

	{"fixed64 - const - valid", &cases.Fixed64Const{Val: 1}, 0, nil},
	{"fixed64 - const - invalid", &cases.Fixed64Const{Val: 2}, 1, nil},

	{"fixed64 - in - valid", &cases.Fixed64In{Val: 3}, 0, nil},
	{"fixed64 - in - invalid", &cases.Fixed64In{Val: 5}, 1, nil},

	{"fixed64 - not in - valid", &cases.Fixed64NotIn{Val: 1}, 0, nil},
	{"fixed64 - not in - invalid", &cases.Fixed64NotIn{Val: 0}, 1, nil},

	{"fixed64 - lt - valid", &cases.Fixed64LT{Val: 4}, 0, nil},
	{"fixed64 - lt - invalid (equal)", &cases.Fixed64LT{Val: 5}, 1, nil},
	{"fixed64 - lt - invalid", &cases.Fixed64LT{Val: 6}, 1, nil},

	{"fixed64 - lte - valid", &cases.Fixed64LTE{Val: 63}, 0, nil},
	{"fixed64 - lte - valid (equal)", &cases.Fixed64LTE{Val: 64}, 0, nil},
	{"fixed64 - lte - invalid", &cases.Fixed64LTE{Val: 65}, 1, nil},

	{"fixed64 - gt - valid", &cases.Fixed64GT{Val: 17}, 0, nil},
	{"fixed64 - gt - invalid (equal)", &cases.Fixed64GT{Val: 16}, 1, nil},
	{"fixed64 - gt - invalid", &cases.Fixed64GT{Val: 15}, 1, nil},

	{"fixed64 - gte - valid", &cases.Fixed64GTE{Val: 9}, 0, nil},
	{"fixed64 - gte - valid (equal)", &cases.Fixed64GTE{Val: 8}, 0, nil},
	{"fixed64 - gte - invalid", &cases.Fixed64GTE{Val: 7}, 1, nil},

	{"fixed64 - gt & lt - valid", &cases.Fixed64GTLT{Val: 7}, 0, nil},
	{"fixed64 - gt & lt - invalid (above)", &cases.Fixed64GTLT{Val: 11}, 1, nil},
	{"fixed64 - gt & lt - invalid (below)", &cases.Fixed64GTLT{Val: 1}, 1, nil},
	{"fixed64 - gt & lt - invalid (max)", &cases.Fixed64GTLT{Val: 10}, 1, nil},
	{"fixed64 - gt & lt - invalid (min)", &cases.Fixed64GTLT{Val: 5}, 1, nil},

	{"fixed64 - exclusive gt & lt - valid (above)", &cases.Fixed64ExLTGT{Val: 11}, 0, nil},
	{"fixed64 - exclusive gt & lt - valid (below)", &cases.Fixed64ExLTGT{Val: 4}, 0, nil},
	{"fixed64 - exclusive gt & lt - invalid", &cases.Fixed64ExLTGT{Val: 7}, 1, nil},
	{"fixed64 - exclusive gt & lt - invalid (max)", &cases.Fixed64ExLTGT{Val: 10}, 1, nil},
	{"fixed64 - exclusive gt & lt - invalid (min)", &cases.Fixed64ExLTGT{Val: 5}, 1, nil},

	{"fixed64 - gte & lte - valid", &cases.Fixed64GTELTE{Val: 200}, 0, nil},
	{"fixed64 - gte & lte - valid (max)", &cases.Fixed64GTELTE{Val: 256}, 0, nil},
	{"fixed64 - gte & lte - valid (min)", &cases.Fixed64GTELTE{Val: 128}, 0, nil},
	{"fixed64 - gte & lte - invalid (above)", &cases.Fixed64GTELTE{Val: 300}, 1, nil},
	{"fixed64 - gte & lte - invalid (below)", &cases.Fixed64GTELTE{Val: 100}, 1, nil},

	{"fixed64 - exclusive gte & lte - valid (above)", &cases.Fixed64ExGTELTE{Val: 300}, 0, nil},
	{"fixed64 - exclusive gte & lte - valid (below)", &cases.Fixed64ExGTELTE{Val: 100}, 0, nil},
	{"fixed64 - exclusive gte & lte - valid (max)", &cases.Fixed64ExGTELTE{Val: 256}, 0, nil},
	{"fixed64 - exclusive gte & lte - valid (min)", &cases.Fixed64ExGTELTE{Val: 128}, 0, nil},
	{"fixed64 - exclusive gte & lte - invalid", &cases.Fixed64ExGTELTE{Val: 200}, 1, nil},

	{"fixed64 - ignore_empty gte & lte - valid", &cases.Fixed64Ignore{Val: 0}, 0, nil},
}

var sfixed32Cases = []TestCase{
	{"sfixed32 - none - valid", &cases.SFixed32None{Val: 123}, 0, nil},

	{"sfixed32 - const - valid", &cases.SFixed32Const{Val: 1}, 0, nil},
	{"sfixed32 - const - invalid", &cases.SFixed32Const{Val: 2}, 1, nil},

	{"sfixed32 - in - valid", &cases.SFixed32In{Val: 3}, 0, nil},
	{"sfixed32 - in - invalid", &cases.SFixed32In{Val: 5}, 1, nil},

	{"sfixed32 - not in - valid", &cases.SFixed32NotIn{Val: 1}, 0, nil},
	{"sfixed32 - not in - invalid", &cases.SFixed32NotIn{Val: 0}, 1, nil},

	{"sfixed32 - lt - valid", &cases.SFixed32LT{Val: -1}, 0, nil},
	{"sfixed32 - lt - invalid (equal)", &cases.SFixed32LT{Val: 0}, 1, nil},
	{"sfixed32 - lt - invalid", &cases.SFixed32LT{Val: 1}, 1, nil},

	{"sfixed32 - lte - valid", &cases.SFixed32LTE{Val: 63}, 0, nil},
	{"sfixed32 - lte - valid (equal)", &cases.SFixed32LTE{Val: 64}, 0, nil},
	{"sfixed32 - lte - invalid", &cases.SFixed32LTE{Val: 65}, 1, nil},

	{"sfixed32 - gt - valid", &cases.SFixed32GT{Val: 17}, 0, nil},
	{"sfixed32 - gt - invalid (equal)", &cases.SFixed32GT{Val: 16}, 1, nil},
	{"sfixed32 - gt - invalid", &cases.SFixed32GT{Val: 15}, 1, nil},

	{"sfixed32 - gte - valid", &cases.SFixed32GTE{Val: 9}, 0, nil},
	{"sfixed32 - gte - valid (equal)", &cases.SFixed32GTE{Val: 8}, 0, nil},
	{"sfixed32 - gte - invalid", &cases.SFixed32GTE{Val: 7}, 1, nil},

	{"sfixed32 - gt & lt - valid", &cases.SFixed32GTLT{Val: 5}, 0, nil},
	{"sfixed32 - gt & lt - invalid (above)", &cases.SFixed32GTLT{Val: 11}, 1, nil},
	{"sfixed32 - gt & lt - invalid (below)", &cases.SFixed32GTLT{Val: -1}, 1, nil},
	{"sfixed32 - gt & lt - invalid (max)", &cases.SFixed32GTLT{Val: 10}, 1, nil},
	{"sfixed32 - gt & lt - invalid (min)", &cases.SFixed32GTLT{Val: 0}, 1, nil},

	{"sfixed32 - exclusive gt & lt - valid (above)", &cases.SFixed32ExLTGT{Val: 11}, 0, nil},
	{"sfixed32 - exclusive gt & lt - valid (below)", &cases.SFixed32ExLTGT{Val: -1}, 0, nil},
	{"sfixed32 - exclusive gt & lt - invalid", &cases.SFixed32ExLTGT{Val: 5}, 1, nil},
	{"sfixed32 - exclusive gt & lt - invalid (max)", &cases.SFixed32ExLTGT{Val: 10}, 1, nil},
	{"sfixed32 - exclusive gt & lt - invalid (min)", &cases.SFixed32ExLTGT{Val: 0}, 1, nil},

	{"sfixed32 - gte & lte - valid", &cases.SFixed32GTELTE{Val: 200}, 0, nil},
	{"sfixed32 - gte & lte - valid (max)", &cases.SFixed32GTELTE{Val: 256}, 0, nil},
	{"sfixed32 - gte & lte - valid (min)", &cases.SFixed32GTELTE{Val: 128}, 0, nil},
	{"sfixed32 - gte & lte - invalid (above)", &cases.SFixed32GTELTE{Val: 300}, 1, nil},
	{"sfixed32 - gte & lte - invalid (below)", &cases.SFixed32GTELTE{Val: 100}, 1, nil},

	{"sfixed32 - exclusive gte & lte - valid (above)", &cases.SFixed32ExGTELTE{Val: 300}, 0, nil},
	{"sfixed32 - exclusive gte & lte - valid (below)", &cases.SFixed32ExGTELTE{Val: 100}, 0, nil},
	{"sfixed32 - exclusive gte & lte - valid (max)", &cases.SFixed32ExGTELTE{Val: 256}, 0, nil},
	{"sfixed32 - exclusive gte & lte - valid (min)", &cases.SFixed32ExGTELTE{Val: 128}, 0, nil},
	{"sfixed32 - exclusive gte & lte - invalid", &cases.SFixed32ExGTELTE{Val: 200}, 1, nil},

	{"sfixed32 - ignore_empty gte & lte - valid", &cases.SFixed32Ignore{Val: 0}, 0, nil},
}

var sfixed64Cases = []TestCase{
	{"sfixed64 - none - valid", &cases.SFixed64None{Val: 123}, 0, nil},

	{"sfixed64 - const - valid", &cases.SFixed64Const{Val: 1}, 0, nil},
	{"sfixed64 - const - invalid", &cases.SFixed64Const{Val: 2}, 1, nil},

	{"sfixed64 - in - valid", &cases.SFixed64In{Val: 3}, 0, nil},
	{"sfixed64 - in - invalid", &cases.SFixed64In{Val: 5}, 1, nil},

	{"sfixed64 - not in - valid", &cases.SFixed64NotIn{Val: 1}, 0, nil},
	{"sfixed64 - not in - invalid", &cases.SFixed64NotIn{Val: 0}, 1, nil},

	{"sfixed64 - lt - valid", &cases.SFixed64LT{Val: -1}, 0, nil},
	{"sfixed64 - lt - invalid (equal)", &cases.SFixed64LT{Val: 0}, 1, nil},
	{"sfixed64 - lt - invalid", &cases.SFixed64LT{Val: 1}, 1, nil},

	{"sfixed64 - lte - valid", &cases.SFixed64LTE{Val: 63}, 0, nil},
	{"sfixed64 - lte - valid (equal)", &cases.SFixed64LTE{Val: 64}, 0, nil},
	{"sfixed64 - lte - invalid", &cases.SFixed64LTE{Val: 65}, 1, nil},

	{"sfixed64 - gt - valid", &cases.SFixed64GT{Val: 17}, 0, nil},
	{"sfixed64 - gt - invalid (equal)", &cases.SFixed64GT{Val: 16}, 1, nil},
	{"sfixed64 - gt - invalid", &cases.SFixed64GT{Val: 15}, 1, nil},

	{"sfixed64 - gte - valid", &cases.SFixed64GTE{Val: 9}, 0, nil},
	{"sfixed64 - gte - valid (equal)", &cases.SFixed64GTE{Val: 8}, 0, nil},
	{"sfixed64 - gte - invalid", &cases.SFixed64GTE{Val: 7}, 1, nil},

	{"sfixed64 - gt & lt - valid", &cases.SFixed64GTLT{Val: 5}, 0, nil},
	{"sfixed64 - gt & lt - invalid (above)", &cases.SFixed64GTLT{Val: 11}, 1, nil},
	{"sfixed64 - gt & lt - invalid (below)", &cases.SFixed64GTLT{Val: -1}, 1, nil},
	{"sfixed64 - gt & lt - invalid (max)", &cases.SFixed64GTLT{Val: 10}, 1, nil},
	{"sfixed64 - gt & lt - invalid (min)", &cases.SFixed64GTLT{Val: 0}, 1, nil},

	{"sfixed64 - exclusive gt & lt - valid (above)", &cases.SFixed64ExLTGT{Val: 11}, 0, nil},
	{"sfixed64 - exclusive gt & lt - valid (below)", &cases.SFixed64ExLTGT{Val: -1}, 0, nil},
	{"sfixed64 - exclusive gt & lt - invalid", &cases.SFixed64ExLTGT{Val: 5}, 1, nil},
	{"sfixed64 - exclusive gt & lt - invalid (max)", &cases.SFixed64ExLTGT{Val: 10}, 1, nil},
	{"sfixed64 - exclusive gt & lt - invalid (min)", &cases.SFixed64ExLTGT{Val: 0}, 1, nil},

	{"sfixed64 - gte & lte - valid", &cases.SFixed64GTELTE{Val: 200}, 0, nil},
	{"sfixed64 - gte & lte - valid (max)", &cases.SFixed64GTELTE{Val: 256}, 0, nil},
	{"sfixed64 - gte & lte - valid (min)", &cases.SFixed64GTELTE{Val: 128}, 0, nil},
	{"sfixed64 - gte & lte - invalid (above)", &cases.SFixed64GTELTE{Val: 300}, 1, nil},
	{"sfixed64 - gte & lte - invalid (below)", &cases.SFixed64GTELTE{Val: 100}, 1, nil},

	{"sfixed64 - exclusive gte & lte - valid (above)", &cases.SFixed64ExGTELTE{Val: 300}, 0, nil},
	{"sfixed64 - exclusive gte & lte - valid (below)", &cases.SFixed64ExGTELTE{Val: 100}, 0, nil},
	{"sfixed64 - exclusive gte & lte - valid (max)", &cases.SFixed64ExGTELTE{Val: 256}, 0, nil},
	{"sfixed64 - exclusive gte & lte - valid (min)", &cases.SFixed64ExGTELTE{Val: 128}, 0, nil},
	{"sfixed64 - exclusive gte & lte - invalid", &cases.SFixed64ExGTELTE{Val: 200}, 1, nil},

	{"sfixed64 - ignore_empty gte & lte - valid", &cases.SFixed64Ignore{Val: 0}, 0, nil},
}

var boolCases = []TestCase{
	{"bool - none - valid", &cases.BoolNone{Val: true}, 0, nil},
	{"bool - const (true) - valid", &cases.BoolConstTrue{Val: true}, 0, nil},
	{"bool - const (true) - invalid", &cases.BoolConstTrue{Val: false}, 1, nil},
	{"bool - const (false) - valid", &cases.BoolConstFalse{Val: false}, 0, nil},
	{"bool - const (false) - invalid", &cases.BoolConstFalse{Val: true}, 1, nil},
}

var stringCases = []TestCase{
	{"string - none - valid", &cases.StringNone{Val: "quux"}, 0, nil},

	{"string - const - valid", &cases.StringConst{Val: "foo"}, 0, nil},
	{"string - const - invalid", &cases.StringConst{Val: "bar"}, 1, nil},

	{"string - in - valid", &cases.StringIn{Val: "bar"}, 0, nil},
	{"string - in - invalid", &cases.StringIn{Val: "quux"}, 1, nil},
	{"string - not in - valid", &cases.StringNotIn{Val: "quux"}, 0, nil},
	{"string - not in - invalid", &cases.StringNotIn{Val: "fizz"}, 1, nil},

	{"string - len - valid", &cases.StringLen{Val: "baz"}, 0, nil},
	{"string - len - valid (multibyte)", &cases.StringLen{Val: "你好吖"}, 0, nil},
	{"string - len - invalid (lt)", &cases.StringLen{Val: "go"}, 1, nil},
	{"string - len - invalid (gt)", &cases.StringLen{Val: "fizz"}, 1, nil},
	{"string - len - invalid (multibyte)", &cases.StringLen{Val: "你好"}, 1, nil},

	{"string - min len - valid", &cases.StringMinLen{Val: "protoc"}, 0, nil},
	{"string - min len - valid (min)", &cases.StringMinLen{Val: "baz"}, 0, nil},
	{"string - min len - invalid", &cases.StringMinLen{Val: "go"}, 1, nil},
	{"string - min len - invalid (multibyte)", &cases.StringMinLen{Val: "你好"}, 1, nil},

	{"string - max len - valid", &cases.StringMaxLen{Val: "foo"}, 0, nil},
	{"string - max len - valid (max)", &cases.StringMaxLen{Val: "proto"}, 0, nil},
	{"string - max len - valid (multibyte)", &cases.StringMaxLen{Val: "你好你好"}, 0, nil},
	{"string - max len - invalid", &cases.StringMaxLen{Val: "1234567890"}, 1, nil},

	{"string - min/max len - valid", &cases.StringMinMaxLen{Val: "quux"}, 0, nil},
	{"string - min/max len - valid (min)", &cases.StringMinMaxLen{Val: "foo"}, 0, nil},
	{"string - min/max len - valid (max)", &cases.StringMinMaxLen{Val: "proto"}, 0, nil},
	{"string - min/max len - valid (multibyte)", &cases.StringMinMaxLen{Val: "你好你好"}, 0, nil},
	{"string - min/max len - invalid (below)", &cases.StringMinMaxLen{Val: "go"}, 1, nil},
	{"string - min/max len - invalid (above)", &cases.StringMinMaxLen{Val: "validate"}, 1, nil},

	{"string - equal min/max len - valid", &cases.StringEqualMinMaxLen{Val: "proto"}, 0, nil},
	{"string - equal min/max len - invalid", &cases.StringEqualMinMaxLen{Val: "validate"}, 1, nil},

	{"string - len bytes - valid", &cases.StringLenBytes{Val: "pace"}, 0, nil},
	{"string - len bytes - invalid (lt)", &cases.StringLenBytes{Val: "val"}, 1, nil},
	{"string - len bytes - invalid (gt)", &cases.StringLenBytes{Val: "world"}, 1, nil},
	{"string - len bytes - invalid (multibyte)", &cases.StringLenBytes{Val: "世界和平"}, 1, nil},

	{"string - min bytes - valid", &cases.StringMinBytes{Val: "proto"}, 0, nil},
	{"string - min bytes - valid (min)", &cases.StringMinBytes{Val: "quux"}, 0, nil},
	{"string - min bytes - valid (multibyte)", &cases.StringMinBytes{Val: "你好"}, 0, nil},
	{"string - min bytes - invalid", &cases.StringMinBytes{Val: ""}, 1, nil},

	{"string - max bytes - valid", &cases.StringMaxBytes{Val: "foo"}, 0, nil},
	{"string - max bytes - valid (max)", &cases.StringMaxBytes{Val: "12345678"}, 0, nil},
	{"string - max bytes - invalid", &cases.StringMaxBytes{Val: "123456789"}, 1, nil},
	{"string - max bytes - invalid (multibyte)", &cases.StringMaxBytes{Val: "你好你好你好"}, 1, nil},

	{"string - min/max bytes - valid", &cases.StringMinMaxBytes{Val: "protoc"}, 0, nil},
	{"string - min/max bytes - valid (min)", &cases.StringMinMaxBytes{Val: "quux"}, 0, nil},
	{"string - min/max bytes - valid (max)", &cases.StringMinMaxBytes{Val: "fizzbuzz"}, 0, nil},
	{"string - min/max bytes - valid (multibyte)", &cases.StringMinMaxBytes{Val: "你好"}, 0, nil},
	{"string - min/max bytes - invalid (below)", &cases.StringMinMaxBytes{Val: "foo"}, 1, nil},
	{"string - min/max bytes - invalid (above)", &cases.StringMinMaxBytes{Val: "你好你好你"}, 1, nil},

	{"string - equal min/max bytes - valid", &cases.StringEqualMinMaxBytes{Val: "protoc"}, 0, nil},
	{"string - equal min/max bytes - invalid", &cases.StringEqualMinMaxBytes{Val: "foo"}, 1, nil},

	{"string - pattern - valid", &cases.StringPattern{Val: "Foo123"}, 0, nil},
	{"string - pattern - invalid", &cases.StringPattern{Val: "!@#$%^&*()"}, 1, nil},
	{"string - pattern - invalid (empty)", &cases.StringPattern{Val: ""}, 1, nil},
	{"string - pattern - invalid (null)", &cases.StringPattern{Val: "a\000"}, 1, nil},

	{"string - pattern (escapes) - valid", &cases.StringPatternEscapes{Val: "* \\ x"}, 0, nil},
	{"string - pattern (escapes) - invalid", &cases.StringPatternEscapes{Val: "invalid"}, 1, nil},
	{"string - pattern (escapes) - invalid (empty)", &cases.StringPatternEscapes{Val: ""}, 1, nil},

	{"string - prefix - valid", &cases.StringPrefix{Val: "foobar"}, 0, nil},
	{"string - prefix - valid (only)", &cases.StringPrefix{Val: "foo"}, 0, nil},
	{"string - prefix - invalid", &cases.StringPrefix{Val: "bar"}, 1, nil},
	{"string - prefix - invalid (case-sensitive)", &cases.StringPrefix{Val: "Foobar"}, 1, nil},

	{"string - contains - valid", &cases.StringContains{Val: "candy bars"}, 0, nil},
	{"string - contains - valid (only)", &cases.StringContains{Val: "bar"}, 0, nil},
	{"string - contains - invalid", &cases.StringContains{Val: "candy bazs"}, 1, nil},
	{"string - contains - invalid (case-sensitive)", &cases.StringContains{Val: "Candy Bars"}, 1, nil},

	{"string - not contains - valid", &cases.StringNotContains{Val: "candy bazs"}, 0, nil},
	{"string - not contains - valid (case-sensitive)", &cases.StringNotContains{Val: "Candy Bars"}, 0, nil},
	{"string - not contains - invalid", &cases.StringNotContains{Val: "candy bars"}, 1, nil},
	{"string - not contains - invalid (equal)", &cases.StringNotContains{Val: "bar"}, 1, nil},

	{"string - suffix - valid", &cases.StringSuffix{Val: "foobaz"}, 0, nil},
	{"string - suffix - valid (only)", &cases.StringSuffix{Val: "baz"}, 0, nil},
	{"string - suffix - invalid", &cases.StringSuffix{Val: "foobar"}, 1, nil},
	{"string - suffix - invalid (case-sensitive)", &cases.StringSuffix{Val: "FooBaz"}, 1, nil},

	{"string - email - valid", &cases.StringEmail{Val: "foo@bar.com"}, 0, nil},
	{"string - email - valid (name)", &cases.StringEmail{Val: "John Smith <foo@bar.com>"}, 0, nil},
	{"string - email - invalid", &cases.StringEmail{Val: "foobar"}, 1, nil},
	{"string - email - invalid (local segment too long)", &cases.StringEmail{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789@example.com"}, 1, nil},
	{"string - email - invalid (hostname too long)", &cases.StringEmail{Val: "foo@x0123456789012345678901234567890123456789012345678901234567890123456789.com"}, 1, nil},
	{"string - email - invalid (bad hostname)", &cases.StringEmail{Val: "foo@-bar.com"}, 1, nil},
	{"string - email - empty", &cases.StringEmail{Val: ""}, 1, nil},

	{"string - address - valid hostname", &cases.StringAddress{Val: "example.com"}, 0, nil},
	{"string - address - valid hostname (uppercase)", &cases.StringAddress{Val: "ASD.example.com"}, 0, nil},
	{"string - address - valid hostname (hyphens)", &cases.StringAddress{Val: "foo-bar.com"}, 0, nil},
	{"string - address - valid hostname (trailing dot)", &cases.StringAddress{Val: "example.com."}, 0, nil},
	{"string - address - invalid hostname", &cases.StringAddress{Val: "!@#$%^&"}, 1, nil},
	{"string - address - invalid hostname (underscore)", &cases.StringAddress{Val: "foo_bar.com"}, 1, nil},
	{"string - address - invalid hostname (too long)", &cases.StringAddress{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789.com"}, 1, nil},
	{"string - address - invalid hostname (trailing hyphens)", &cases.StringAddress{Val: "foo-bar-.com"}, 1, nil},
	{"string - address - invalid hostname (leading hyphens)", &cases.StringAddress{Val: "foo-bar.-com"}, 1, nil},
	{"string - address - invalid hostname (empty)", &cases.StringAddress{Val: "asd..asd.com"}, 1, nil},
	{"string - address - invalid hostname (IDNs)", &cases.StringAddress{Val: "你好.com"}, 1, nil},
	{"string - address - valid ip (v4)", &cases.StringAddress{Val: "192.168.0.1"}, 0, nil},
	{"string - address - valid ip (v6)", &cases.StringAddress{Val: "3e::99"}, 0, nil},
	{"string - address - invalid ip", &cases.StringAddress{Val: "ff::fff::0b"}, 1, nil},

	{"string - hostname - valid", &cases.StringHostname{Val: "example.com"}, 0, nil},
	{"string - hostname - valid (uppercase)", &cases.StringHostname{Val: "ASD.example.com"}, 0, nil},
	{"string - hostname - valid (hyphens)", &cases.StringHostname{Val: "foo-bar.com"}, 0, nil},
	{"string - hostname - valid (trailing dot)", &cases.StringHostname{Val: "example.com."}, 0, nil},
	{"string - hostname - invalid", &cases.StringHostname{Val: "!@#$%^&"}, 1, nil},
	{"string - hostname - invalid (underscore)", &cases.StringHostname{Val: "foo_bar.com"}, 1, nil},
	{"string - hostname - invalid (too long)", &cases.StringHostname{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789.com"}, 1, nil},
	{"string - hostname - invalid (trailing hyphens)", &cases.StringHostname{Val: "foo-bar-.com"}, 1, nil},
	{"string - hostname - invalid (leading hyphens)", &cases.StringHostname{Val: "foo-bar.-com"}, 1, nil},
	{"string - hostname - invalid (empty)", &cases.StringHostname{Val: "asd..asd.com"}, 1, nil},
	{"string - hostname - invalid (IDNs)", &cases.StringHostname{Val: "你好.com"}, 1, nil},

	{"string - IP - valid (v4)", &cases.StringIP{Val: "192.168.0.1"}, 0, nil},
	{"string - IP - valid (v6)", &cases.StringIP{Val: "3e::99"}, 0, nil},
	{"string - IP - invalid", &cases.StringIP{Val: "foobar"}, 1, nil},

	{"string - IPv4 - valid", &cases.StringIPv4{Val: "192.168.0.1"}, 0, nil},
	{"string - IPv4 - invalid", &cases.StringIPv4{Val: "foobar"}, 1, nil},
	{"string - IPv4 - invalid (erroneous)", &cases.StringIPv4{Val: "256.0.0.0"}, 1, nil},
	{"string - IPv4 - invalid (v6)", &cases.StringIPv4{Val: "3e::99"}, 1, nil},

	{"string - IPv6 - valid", &cases.StringIPv6{Val: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, 0, nil},
	{"string - IPv6 - valid (collapsed)", &cases.StringIPv6{Val: "2001:db8:85a3::8a2e:370:7334"}, 0, nil},
	{"string - IPv6 - invalid", &cases.StringIPv6{Val: "foobar"}, 1, nil},
	{"string - IPv6 - invalid (v4)", &cases.StringIPv6{Val: "192.168.0.1"}, 1, nil},
	{"string - IPv6 - invalid (erroneous)", &cases.StringIPv6{Val: "ff::fff::0b"}, 1, nil},

	{"string - URI - valid", &cases.StringURI{Val: "http://example.com/foo/bar?baz=quux"}, 0, nil},
	{"string - URI - invalid", &cases.StringURI{Val: "!@#$%^&*%$#"}, 1, nil},
	{"string - URI - invalid (relative)", &cases.StringURI{Val: "/foo/bar?baz=quux"}, 1, nil},

	{"string - URI - valid", &cases.StringURIRef{Val: "http://example.com/foo/bar?baz=quux"}, 0, nil},
	{"string - URI - valid (relative)", &cases.StringURIRef{Val: "/foo/bar?baz=quux"}, 0, nil},
	{"string - URI - invalid", &cases.StringURIRef{Val: "!@#$%^&*%$#"}, 1, nil},

	{"string - UUID - valid (nil)", &cases.StringUUID{Val: "00000000-0000-0000-0000-000000000000"}, 0, nil},
	{"string - UUID - valid (v1)", &cases.StringUUID{Val: "b45c0c80-8880-11e9-a5b1-000000000000"}, 0, nil},
	{"string - UUID - valid (v1 - case-insensitive)", &cases.StringUUID{Val: "B45C0C80-8880-11E9-A5B1-000000000000"}, 0, nil},
	{"string - UUID - valid (v2)", &cases.StringUUID{Val: "b45c0c80-8880-21e9-a5b1-000000000000"}, 0, nil},
	{"string - UUID - valid (v2 - case-insensitive)", &cases.StringUUID{Val: "B45C0C80-8880-21E9-A5B1-000000000000"}, 0, nil},
	{"string - UUID - valid (v3)", &cases.StringUUID{Val: "a3bb189e-8bf9-3888-9912-ace4e6543002"}, 0, nil},
	{"string - UUID - valid (v3 - case-insensitive)", &cases.StringUUID{Val: "A3BB189E-8BF9-3888-9912-ACE4E6543002"}, 0, nil},
	{"string - UUID - valid (v4)", &cases.StringUUID{Val: "8b208305-00e8-4460-a440-5e0dcd83bb0a"}, 0, nil},
	{"string - UUID - valid (v4 - case-insensitive)", &cases.StringUUID{Val: "8B208305-00E8-4460-A440-5E0DCD83BB0A"}, 0, nil},
	{"string - UUID - valid (v5)", &cases.StringUUID{Val: "a6edc906-2f9f-5fb2-a373-efac406f0ef2"}, 0, nil},
	{"string - UUID - valid (v5 - case-insensitive)", &cases.StringUUID{Val: "A6EDC906-2F9F-5FB2-A373-EFAC406F0EF2"}, 0, nil},
	{"string - UUID - invalid", &cases.StringUUID{Val: "foobar"}, 1, nil},
	{"string - UUID - invalid (bad UUID)", &cases.StringUUID{Val: "ffffffff-ffff-ffff-ffff-fffffffffffff"}, 1, nil},
	{"string - UUID - valid (ignore_empty)", &cases.StringUUIDIgnore{Val: ""}, 0, nil},

	{"string - http header name - valid", &cases.StringHttpHeaderName{Val: "clustername"}, 0, nil},
	{"string - http header name - valid", &cases.StringHttpHeaderName{Val: ":path"}, 0, nil},
	{"string - http header name - valid (nums)", &cases.StringHttpHeaderName{Val: "cluster-123"}, 0, nil},
	{"string - http header name - valid (special token)", &cases.StringHttpHeaderName{Val: "!+#&.%"}, 0, nil},
	{"string - http header name - valid (period)", &cases.StringHttpHeaderName{Val: "CLUSTER.NAME"}, 0, nil},
	{"string - http header name - invalid", &cases.StringHttpHeaderName{Val: ":"}, 1, nil},
	{"string - http header name - invalid", &cases.StringHttpHeaderName{Val: ":path:"}, 1, nil},
	{"string - http header name - invalid (space)", &cases.StringHttpHeaderName{Val: "cluster name"}, 1, nil},
	{"string - http header name - invalid (return)", &cases.StringHttpHeaderName{Val: "example\r"}, 1, nil},
	{"string - http header name - invalid (tab)", &cases.StringHttpHeaderName{Val: "example\t"}, 1, nil},
	{"string - http header name - invalid (slash)", &cases.StringHttpHeaderName{Val: "/test/long/url"}, 1, nil},

	{"string - http header value - valid", &cases.StringHttpHeaderValue{Val: "cluster.name.123"}, 0, nil},
	{"string - http header value - valid (uppercase)", &cases.StringHttpHeaderValue{Val: "/TEST/LONG/URL"}, 0, nil},
	{"string - http header value - valid (spaces)", &cases.StringHttpHeaderValue{Val: "cluster name"}, 0, nil},
	{"string - http header value - valid (tab)", &cases.StringHttpHeaderValue{Val: "example\t"}, 0, nil},
	{"string - http header value - valid (special token)", &cases.StringHttpHeaderValue{Val: "!#%&./+"}, 0, nil},
	{"string - http header value - invalid (NUL)", &cases.StringHttpHeaderValue{Val: "foo\u0000bar"}, 1, nil},
	{"string - http header value - invalid (DEL)", &cases.StringHttpHeaderValue{Val: "\u007f"}, 1, nil},
	{"string - http header value - invalid", &cases.StringHttpHeaderValue{Val: "example\r"}, 1, nil},

	{"string - non-strict valid header - valid", &cases.StringValidHeader{Val: "cluster.name.123"}, 0, nil},
	{"string - non-strict valid header - valid (uppercase)", &cases.StringValidHeader{Val: "/TEST/LONG/URL"}, 0, nil},
	{"string - non-strict valid header - valid (spaces)", &cases.StringValidHeader{Val: "cluster name"}, 0, nil},
	{"string - non-strict valid header - valid (tab)", &cases.StringValidHeader{Val: "example\t"}, 0, nil},
	{"string - non-strict valid header - valid (DEL)", &cases.StringValidHeader{Val: "\u007f"}, 0, nil},
	{"string - non-strict valid header - invalid (NUL)", &cases.StringValidHeader{Val: "foo\u0000bar"}, 1, nil},
	{"string - non-strict valid header - invalid (CR)", &cases.StringValidHeader{Val: "example\r"}, 1, nil},
	{"string - non-strict valid header - invalid (NL)", &cases.StringValidHeader{Val: "exa\u000Ample"}, 1, nil},
}

var bytesCases = []TestCase{
	{"bytes - none - valid", &cases.BytesNone{Val: []byte("quux")}, 0, nil},

	{"bytes - const - valid", &cases.BytesConst{Val: []byte("foo")}, 0, nil},
	{"bytes - const - invalid", &cases.BytesConst{Val: []byte("bar")}, 1, nil},

	{"bytes - in - valid", &cases.BytesIn{Val: []byte("bar")}, 0, nil},
	{"bytes - in - invalid", &cases.BytesIn{Val: []byte("quux")}, 1, nil},
	{"bytes - not in - valid", &cases.BytesNotIn{Val: []byte("quux")}, 0, nil},
	{"bytes - not in - invalid", &cases.BytesNotIn{Val: []byte("fizz")}, 1, nil},

	{"bytes - len - valid", &cases.BytesLen{Val: []byte("baz")}, 0, nil},
	{"bytes - len - invalid (lt)", &cases.BytesLen{Val: []byte("go")}, 1, nil},
	{"bytes - len - invalid (gt)", &cases.BytesLen{Val: []byte("fizz")}, 1, nil},

	{"bytes - min len - valid", &cases.BytesMinLen{Val: []byte("fizz")}, 0, nil},
	{"bytes - min len - valid (min)", &cases.BytesMinLen{Val: []byte("baz")}, 0, nil},
	{"bytes - min len - invalid", &cases.BytesMinLen{Val: []byte("go")}, 1, nil},

	{"bytes - max len - valid", &cases.BytesMaxLen{Val: []byte("foo")}, 0, nil},
	{"bytes - max len - valid (max)", &cases.BytesMaxLen{Val: []byte("proto")}, 0, nil},
	{"bytes - max len - invalid", &cases.BytesMaxLen{Val: []byte("1234567890")}, 1, nil},

	{"bytes - min/max len - valid", &cases.BytesMinMaxLen{Val: []byte("quux")}, 0, nil},
	{"bytes - min/max len - valid (min)", &cases.BytesMinMaxLen{Val: []byte("foo")}, 0, nil},
	{"bytes - min/max len - valid (max)", &cases.BytesMinMaxLen{Val: []byte("proto")}, 0, nil},
	{"bytes - min/max len - invalid (below)", &cases.BytesMinMaxLen{Val: []byte("go")}, 1, nil},
	{"bytes - min/max len - invalid (above)", &cases.BytesMinMaxLen{Val: []byte("validate")}, 1, nil},

	{"bytes - equal min/max len - valid", &cases.BytesEqualMinMaxLen{Val: []byte("proto")}, 0, nil},
	{"bytes - equal min/max len - invalid", &cases.BytesEqualMinMaxLen{Val: []byte("validate")}, 1, nil},

	{"bytes - pattern - valid", &cases.BytesPattern{Val: []byte("Foo123")}, 0, nil},
	{"bytes - pattern - invalid", &cases.BytesPattern{Val: []byte("你好你好")}, 1, nil},
	{"bytes - pattern - invalid (empty)", &cases.BytesPattern{Val: []byte("")}, 1, nil},

	{"bytes - prefix - valid", &cases.BytesPrefix{Val: []byte{0x99, 0x9f, 0x08}}, 0, nil},
	{"bytes - prefix - valid (only)", &cases.BytesPrefix{Val: []byte{0x99}}, 0, nil},
	{"bytes - prefix - invalid", &cases.BytesPrefix{Val: []byte("bar")}, 1, nil},

	{"bytes - contains - valid", &cases.BytesContains{Val: []byte("candy bars")}, 0, nil},
	{"bytes - contains - valid (only)", &cases.BytesContains{Val: []byte("bar")}, 0, nil},
	{"bytes - contains - invalid", &cases.BytesContains{Val: []byte("candy bazs")}, 1, nil},

	{"bytes - suffix - valid", &cases.BytesSuffix{Val: []byte{0x62, 0x75, 0x7A, 0x7A}}, 0, nil},
	{"bytes - suffix - valid (only)", &cases.BytesSuffix{Val: []byte("\x62\x75\x7A\x7A")}, 0, nil},
	{"bytes - suffix - invalid", &cases.BytesSuffix{Val: []byte("foobar")}, 1, nil},
	{"bytes - suffix - invalid (case-sensitive)", &cases.BytesSuffix{Val: []byte("FooBaz")}, 1, nil},

	{"bytes - IP - valid (v4)", &cases.BytesIP{Val: []byte{0xC0, 0xA8, 0x00, 0x01}}, 0, nil},
	{"bytes - IP - valid (v6)", &cases.BytesIP{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")}, 0, nil},
	{"bytes - IP - invalid", &cases.BytesIP{Val: []byte("foobar")}, 1, nil},

	{"bytes - IPv4 - valid", &cases.BytesIPv4{Val: []byte{0xC0, 0xA8, 0x00, 0x01}}, 0, nil},
	{"bytes - IPv4 - invalid", &cases.BytesIPv4{Val: []byte("foobar")}, 1, nil},
	{"bytes - IPv4 - invalid (v6)", &cases.BytesIPv4{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")}, 1, nil},

	{"bytes - IPv6 - valid", &cases.BytesIPv6{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")}, 0, nil},
	{"bytes - IPv6 - invalid", &cases.BytesIPv6{Val: []byte("fooar")}, 1, nil},
	{"bytes - IPv6 - invalid (v4)", &cases.BytesIPv6{Val: []byte{0xC0, 0xA8, 0x00, 0x01}}, 1, nil},

	{"bytes - IPv6 - valid (ignore_empty)", &cases.BytesIPv6Ignore{Val: nil}, 0, nil},
}

var enumCases = []TestCase{
	{"enum - none - valid", &cases.EnumNone{Val: cases.TestEnum_ONE}, 0, nil},

	{"enum - const - valid", &cases.EnumConst{Val: cases.TestEnum_TWO}, 0, nil},
	{"enum - const - invalid", &cases.EnumConst{Val: cases.TestEnum_ONE}, 1, nil},
	{"enum alias - const - valid", &cases.EnumAliasConst{Val: cases.TestEnumAlias_C}, 0, nil},
	{"enum alias - const - valid (alias)", &cases.EnumAliasConst{Val: cases.TestEnumAlias_GAMMA}, 0, nil},
	{"enum alias - const - invalid", &cases.EnumAliasConst{Val: cases.TestEnumAlias_ALPHA}, 1, nil},

	{"enum - defined_only - valid", &cases.EnumDefined{Val: 0}, 0, nil},
	{"enum - defined_only - invalid", &cases.EnumDefined{Val: math.MaxInt32}, 1, nil},
	{"enum alias - defined_only - valid", &cases.EnumAliasDefined{Val: 1}, 0, nil},
	{"enum alias - defined_only - invalid", &cases.EnumAliasDefined{Val: math.MaxInt32}, 1, nil},

	{"enum - in - valid", &cases.EnumIn{Val: cases.TestEnum_TWO}, 0, nil},
	{"enum - in - invalid", &cases.EnumIn{Val: cases.TestEnum_ONE}, 1, nil},
	{"enum alias - in - valid", &cases.EnumAliasIn{Val: cases.TestEnumAlias_A}, 0, nil},
	{"enum alias - in - valid (alias)", &cases.EnumAliasIn{Val: cases.TestEnumAlias_ALPHA}, 0, nil},
	{"enum alias - in - invalid", &cases.EnumAliasIn{Val: cases.TestEnumAlias_BETA}, 1, nil},

	{"enum - not in - valid", &cases.EnumNotIn{Val: cases.TestEnum_ZERO}, 0, nil},
	{"enum - not in - valid (undefined)", &cases.EnumNotIn{Val: math.MaxInt32}, 0, nil},
	{"enum - not in - invalid", &cases.EnumNotIn{Val: cases.TestEnum_ONE}, 1, nil},
	{"enum alias - not in - valid", &cases.EnumAliasNotIn{Val: cases.TestEnumAlias_ALPHA}, 0, nil},
	{"enum alias - not in - invalid", &cases.EnumAliasNotIn{Val: cases.TestEnumAlias_B}, 1, nil},
	{"enum alias - not in - invalid (alias)", &cases.EnumAliasNotIn{Val: cases.TestEnumAlias_BETA}, 1, nil},

	{"enum external - defined_only - valid", &cases.EnumExternal{Val: other_package.Embed_VALUE}, 0, nil},
	{"enum external - defined_only - invalid", &cases.EnumExternal{Val: math.MaxInt32}, 1, nil},

	{"enum repeated - defined_only - valid", &cases.RepeatedEnumDefined{Val: []cases.TestEnum{cases.TestEnum_ONE, cases.TestEnum_TWO}}, 0, nil},
	{"enum repeated - defined_only - invalid", &cases.RepeatedEnumDefined{Val: []cases.TestEnum{cases.TestEnum_ONE, math.MaxInt32}}, 1, nil},

	{"enum repeated (external) - defined_only - valid", &cases.RepeatedExternalEnumDefined{Val: []other_package.Embed_Enumerated{other_package.Embed_VALUE}}, 0, nil},
	{"enum repeated (external) - defined_only - invalid", &cases.RepeatedExternalEnumDefined{Val: []other_package.Embed_Enumerated{math.MaxInt32}}, 1, nil},

	{"enum repeated (another external) - defined_only - valid", &cases.RepeatedYetAnotherExternalEnumDefined{Val: []yet_another_package.Embed_Enumerated{yet_another_package.Embed_VALUE}}, 0, nil},

	{"enum map - defined_only - valid", &cases.MapEnumDefined{Val: map[string]cases.TestEnum{"foo": cases.TestEnum_TWO}}, 0, nil},
	{"enum map - defined_only - invalid", &cases.MapEnumDefined{Val: map[string]cases.TestEnum{"foo": math.MaxInt32}}, 1, nil},

	{"enum map (external) - defined_only - valid", &cases.MapExternalEnumDefined{Val: map[string]other_package.Embed_Enumerated{"foo": other_package.Embed_VALUE}}, 0, nil},
	{"enum map (external) - defined_only - invalid", &cases.MapExternalEnumDefined{Val: map[string]other_package.Embed_Enumerated{"foo": math.MaxInt32}}, 1, nil},
}

var messageCases = []TestCase{
	{"message - none - valid", &cases.MessageNone{Val: &cases.MessageNone_NoneMsg{}}, 0, nil},
	{"message - none - valid (unset)", &cases.MessageNone{}, 0, nil},

	{"message - disabled - valid", &cases.MessageDisabled{Val: 456}, 0, nil},
	{"message - disabled - valid (invalid field)", &cases.MessageDisabled{Val: 0}, 0, nil},

	{"message - ignored - valid", &cases.MessageIgnored{Val: 456}, 0, nil},
	{"message - ignored - valid (invalid field)", &cases.MessageIgnored{Val: 0}, 0, nil},

	{"message - field - valid", &cases.Message{Val: &cases.TestMsg{Const: "foo"}}, 0, nil},
	{"message - field - valid (unset)", &cases.Message{}, 0, nil},
	{"message - field - invalid", &cases.Message{Val: &cases.TestMsg{}}, 1, nil},
	{"message - field - invalid (transitive)", &cases.Message{Val: &cases.TestMsg{Const: "foo", Nested: &cases.TestMsg{}}}, 1, nil},

	{"message - skip - valid", &cases.MessageSkip{Val: &cases.TestMsg{}}, 0, nil},

	{"message - required - valid", &cases.MessageRequired{Val: &cases.TestMsg{Const: "foo"}}, 0, nil},
	{"message - required - valid (oneof)", &cases.MessageRequiredOneof{One: &cases.MessageRequiredOneof_Val{&cases.TestMsg{Const: "foo"}}}, 0, nil},
	{"message - required - invalid", &cases.MessageRequired{}, 1, nil},
	{"message - required - invalid (oneof)", &cases.MessageRequiredOneof{}, 1, nil},

	{"message - cross-package embed none - valid", &cases.MessageCrossPackage{Val: &other_package.Embed{Val: 1}}, 0, nil},
	{"message - cross-package embed none - valid (nil)", &cases.MessageCrossPackage{}, 0, nil},
	{"message - cross-package embed none - valid (empty)", &cases.MessageCrossPackage{Val: &other_package.Embed{}}, 1, nil},
	{"message - cross-package embed none - invalid", &cases.MessageCrossPackage{Val: &other_package.Embed{Val: -1}}, 1, nil},

	{"message - required - valid", &cases.MessageRequiredButOptional{Val: &cases.TestMsg{Const: "foo"}}, 0, nil},
	{"message - required - valid (unset)", &cases.MessageRequiredButOptional{}, 0, nil},
}

var repeatedCases = []TestCase{
	{"repeated - none - valid", &cases.RepeatedNone{Val: []int64{1, 2, 3}}, 0, nil},

	{"repeated - embed none - valid", &cases.RepeatedEmbedNone{Val: []*cases.Embed{{Val: 1}}}, 0, nil},
	{"repeated - embed none - valid (nil)", &cases.RepeatedEmbedNone{}, 0, nil},
	{"repeated - embed none - valid (empty)", &cases.RepeatedEmbedNone{Val: []*cases.Embed{}}, 0, nil},
	{"repeated - embed none - invalid", &cases.RepeatedEmbedNone{Val: []*cases.Embed{{Val: -1}}}, 1, nil},

	{"repeated - cross-package embed none - valid", &cases.RepeatedEmbedCrossPackageNone{Val: []*other_package.Embed{{Val: 1}}}, 0, nil},
	{"repeated - cross-package embed none - valid (nil)", &cases.RepeatedEmbedCrossPackageNone{}, 0, nil},
	{"repeated - cross-package embed none - valid (empty)", &cases.RepeatedEmbedCrossPackageNone{Val: []*other_package.Embed{}}, 0, nil},
	{"repeated - cross-package embed none - invalid", &cases.RepeatedEmbedCrossPackageNone{Val: []*other_package.Embed{{Val: -1}}}, 1, nil},

	{"repeated - min - valid", &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}, {Val: 2}, {Val: 3}}}, 0, nil},
	{"repeated - min - valid (equal)", &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}, {Val: 2}}}, 0, nil},
	{"repeated - min - invalid", &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}}}, 1, nil},
	{"repeated - min - invalid (element)", &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}, {Val: -1}}}, 1, nil},

	{"repeated - max - valid", &cases.RepeatedMax{Val: []float64{1, 2}}, 0, nil},
	{"repeated - max - valid (equal)", &cases.RepeatedMax{Val: []float64{1, 2, 3}}, 0, nil},
	{"repeated - max - invalid", &cases.RepeatedMax{Val: []float64{1, 2, 3, 4}}, 1, nil},

	{"repeated - min/max - valid", &cases.RepeatedMinMax{Val: []int32{1, 2, 3}}, 0, nil},
	{"repeated - min/max - valid (min)", &cases.RepeatedMinMax{Val: []int32{1, 2}}, 0, nil},
	{"repeated - min/max - valid (max)", &cases.RepeatedMinMax{Val: []int32{1, 2, 3, 4}}, 0, nil},
	{"repeated - min/max - invalid (below)", &cases.RepeatedMinMax{Val: []int32{}}, 1, nil},
	{"repeated - min/max - invalid (above)", &cases.RepeatedMinMax{Val: []int32{1, 2, 3, 4, 5}}, 1, nil},

	{"repeated - exact - valid", &cases.RepeatedExact{Val: []uint32{1, 2, 3}}, 0, nil},
	{"repeated - exact - invalid (below)", &cases.RepeatedExact{Val: []uint32{1, 2}}, 1, nil},
	{"repeated - exact - invalid (above)", &cases.RepeatedExact{Val: []uint32{1, 2, 3, 4}}, 1, nil},

	{"repeated - unique - valid", &cases.RepeatedUnique{Val: []string{"foo", "bar", "baz"}}, 0, nil},
	{"repeated - unique - valid (empty)", &cases.RepeatedUnique{}, 0, nil},
	{"repeated - unique - valid (case sensitivity)", &cases.RepeatedUnique{Val: []string{"foo", "Foo"}}, 0, nil},
	{"repeated - unique - invalid", &cases.RepeatedUnique{Val: []string{"foo", "bar", "foo", "baz"}}, 1, nil},

	{"repeated - items - valid", &cases.RepeatedItemRule{Val: []float32{1, 2, 3}}, 0, nil},
	{"repeated - items - valid (empty)", &cases.RepeatedItemRule{Val: []float32{}}, 0, nil},
	{"repeated - items - valid (pattern)", &cases.RepeatedItemPattern{Val: []string{"Alpha", "Beta123"}}, 0, nil},
	{"repeated - items - invalid", &cases.RepeatedItemRule{Val: []float32{1, -2, 3}}, 1, nil},
	{"repeated - items - invalid (pattern)", &cases.RepeatedItemPattern{Val: []string{"Alpha", "!@#$%^&*()"}}, 1, nil},
	{"repeated - items - invalid (in)", &cases.RepeatedItemIn{Val: []string{"baz"}}, 1, nil},
	{"repeated - items - valid (in)", &cases.RepeatedItemIn{Val: []string{"foo"}}, 0, nil},
	{"repeated - items - invalid (not_in)", &cases.RepeatedItemNotIn{Val: []string{"foo"}}, 1, nil},
	{"repeated - items - valid (not_in)", &cases.RepeatedItemNotIn{Val: []string{"baz"}}, 0, nil},

	{"repeated - items - invalid (enum in)", &cases.RepeatedEnumIn{Val: []cases.AnEnum{1}}, 1, nil},
	{"repeated - items - valid (enum in)", &cases.RepeatedEnumIn{Val: []cases.AnEnum{0}}, 0, nil},
	{"repeated - items - invalid (enum not_in)", &cases.RepeatedEnumNotIn{Val: []cases.AnEnum{0}}, 1, nil},
	{"repeated - items - valid (enum not_in)", &cases.RepeatedEnumNotIn{Val: []cases.AnEnum{1}}, 0, nil},
	{"repeated - items - invalid (embedded enum in)", &cases.RepeatedEmbeddedEnumIn{Val: []cases.RepeatedEmbeddedEnumIn_AnotherInEnum{1}}, 1, nil},
	{"repeated - items - valid (embedded enum in)", &cases.RepeatedEmbeddedEnumIn{Val: []cases.RepeatedEmbeddedEnumIn_AnotherInEnum{0}}, 0, nil},
	{"repeated - items - invalid (embedded enum not_in)", &cases.RepeatedEmbeddedEnumNotIn{Val: []cases.RepeatedEmbeddedEnumNotIn_AnotherNotInEnum{0}}, 1, nil},
	{"repeated - items - valid (embedded enum not_in)", &cases.RepeatedEmbeddedEnumNotIn{Val: []cases.RepeatedEmbeddedEnumNotIn_AnotherNotInEnum{1}}, 0, nil},

	{"repeated - items - invalid (any in)", &cases.RepeatedAnyIn{Val: []*anypb.Any{{TypeUrl: "type.googleapis.com/google.protobuf.Timestamp"}}}, 1, nil},
	{"repeated - items - valid (any in)", &cases.RepeatedAnyIn{Val: []*anypb.Any{{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}}}, 0, nil},
	{"repeated - items - invalid (any not_in)", &cases.RepeatedAnyNotIn{Val: []*anypb.Any{{TypeUrl: "type.googleapis.com/google.protobuf.Timestamp"}}}, 1, nil},
	{"repeated - items - valid (any not_in)", &cases.RepeatedAnyNotIn{Val: []*anypb.Any{{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}}}, 0, nil},

	{"repeated - embed skip - valid", &cases.RepeatedEmbedSkip{Val: []*cases.Embed{{Val: 1}}}, 0, nil},
	{"repeated - embed skip - valid (invalid element)", &cases.RepeatedEmbedSkip{Val: []*cases.Embed{{Val: -1}}}, 0, nil},
	{"repeated - min and items len - valid", &cases.RepeatedMinAndItemLen{Val: []string{"aaa", "bbb"}}, 0, nil},
	{"repeated - min and items len - invalid (min)", &cases.RepeatedMinAndItemLen{Val: []string{}}, 1, nil},
	{"repeated - min and items len - invalid (len)", &cases.RepeatedMinAndItemLen{Val: []string{"x"}}, 1, nil},
	{"repeated - min and max items len - valid", &cases.RepeatedMinAndMaxItemLen{Val: []string{"aaa", "bbb"}}, 0, nil},
	{"repeated - min and max items len - invalid (min_len)", &cases.RepeatedMinAndMaxItemLen{}, 1, nil},
	{"repeated - min and max items len - invalid (max_len)", &cases.RepeatedMinAndMaxItemLen{Val: []string{"aaa", "bbb", "ccc", "ddd"}}, 1, nil},

	{"repeated - duration - gte - valid", &cases.RepeatedDuration{Val: []*durationpb.Duration{{Seconds: 3}}}, 0, nil},
	{"repeated - duration - gte - valid (empty)", &cases.RepeatedDuration{}, 0, nil},
	{"repeated - duration - gte - valid (equal)", &cases.RepeatedDuration{Val: []*durationpb.Duration{{Nanos: 1000000}}}, 0, nil},
	{"repeated - duration - gte - invalid", &cases.RepeatedDuration{Val: []*durationpb.Duration{{Seconds: -1}}}, 1, nil},

	{"repeated - exact - valid (ignore_empty)", &cases.RepeatedExactIgnore{Val: nil}, 0, nil},
}

var mapCases = []TestCase{
	{"map - none - valid", &cases.MapNone{Val: map[uint32]bool{123: true, 456: false}}, 0, nil},

	{"map - min pairs - valid", &cases.MapMin{Val: map[int32]float32{1: 2, 3: 4, 5: 6}}, 0, nil},
	{"map - min pairs - valid (equal)", &cases.MapMin{Val: map[int32]float32{1: 2, 3: 4}}, 0, nil},
	{"map - min pairs - invalid", &cases.MapMin{Val: map[int32]float32{1: 2}}, 1, nil},

	{"map - max pairs - valid", &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4}}, 0, nil},
	{"map - max pairs - valid (equal)", &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4, 5: 6}}, 0, nil},
	{"map - max pairs - invalid", &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4, 5: 6, 7: 8}}, 1, nil},

	{"map - min/max - valid", &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true}}, 0, nil},
	{"map - min/max - valid (min)", &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false}}, 0, nil},
	{"map - min/max - valid (max)", &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true, "d": false}}, 0, nil},
	{"map - min/max - invalid (below)", &cases.MapMinMax{Val: map[string]bool{}}, 1, nil},
	{"map - min/max - invalid (above)", &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true, "d": false, "e": true}}, 1, nil},

	{"map - exact - valid", &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b", 3: "c"}}, 0, nil},
	{"map - exact - invalid (below)", &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b"}}, 1, nil},
	{"map - exact - invalid (above)", &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b", 3: "c", 4: "d"}}, 1, nil},

	{"map - no sparse - valid", &cases.MapNoSparse{Val: map[uint32]*cases.MapNoSparse_Msg{1: {}, 2: {}}}, 0, nil},
	{"map - no sparse - valid (empty)", &cases.MapNoSparse{Val: map[uint32]*cases.MapNoSparse_Msg{}}, 0, nil},
	// sparse maps are no longer supported, so this case is no longer possible
	//{"map - no sparse - invalid", &cases.MapNoSparse{Val: map[uint32]*cases.MapNoSparse_Msg{1: {}, 2: nil}}, 1, nil},

	{"map - keys - valid", &cases.MapKeys{Val: map[int64]string{-1: "a", -2: "b"}}, 0, nil},
	{"map - keys - valid (empty)", &cases.MapKeys{Val: map[int64]string{}}, 0, nil},
	{"map - keys - valid (pattern)", &cases.MapKeysPattern{Val: map[string]string{"A": "a"}}, 0, nil},
	{"map - keys - invalid", &cases.MapKeys{Val: map[int64]string{1: "a"}}, 1, nil},
	{"map - keys - invalid (pattern)", &cases.MapKeysPattern{Val: map[string]string{"A": "a", "!@#$%^&*()": "b"}}, 1, nil},

	{"map - values - valid", &cases.MapValues{Val: map[string]string{"a": "Alpha", "b": "Beta"}}, 0, nil},
	{"map - values - valid (empty)", &cases.MapValues{Val: map[string]string{}}, 0, nil},
	{"map - values - valid (pattern)", &cases.MapValuesPattern{Val: map[string]string{"a": "A"}}, 0, nil},
	{"map - values - invalid", &cases.MapValues{Val: map[string]string{"a": "A", "b": "B"}}, 2, nil},
	{"map - values - invalid (pattern)", &cases.MapValuesPattern{Val: map[string]string{"a": "A", "b": "!@#$%^&*()"}}, 1, nil},

	{"map - recursive - valid", &cases.MapRecursive{Val: map[uint32]*cases.MapRecursive_Msg{1: {Val: "abc"}}}, 0, nil},
	{"map - recursive - invalid", &cases.MapRecursive{Val: map[uint32]*cases.MapRecursive_Msg{1: {}}}, 1, nil},
	{"map - exact - valid (ignore_empty)", &cases.MapExactIgnore{Val: nil}, 0, nil},
	{"map - multiple - valid", &cases.MultipleMaps{First: map[uint32]string{1: "a", 2: "b"}, Second: map[int32]bool{-1: true, -2: false}}, 0, nil},
}

var oneofCases = []TestCase{
	{"oneof - none - valid", &cases.OneOfNone{O: &cases.OneOfNone_X{X: "foo"}}, 0, nil},
	{"oneof - none - valid (empty)", &cases.OneOfNone{}, 0, nil},

	{"oneof - field - valid (X)", &cases.OneOf{O: &cases.OneOf_X{X: "foobar"}}, 0, nil},
	{"oneof - field - valid (Y)", &cases.OneOf{O: &cases.OneOf_Y{Y: 123}}, 0, nil},
	{"oneof - field - valid (Z)", &cases.OneOf{O: &cases.OneOf_Z{Z: &cases.TestOneOfMsg{Val: true}}}, 0, nil},
	{"oneof - field - valid (empty)", &cases.OneOf{}, 0, nil},
	{"oneof - field - invalid (X)", &cases.OneOf{O: &cases.OneOf_X{X: "fizzbuzz"}}, 1, nil},
	{"oneof - field - invalid (Y)", &cases.OneOf{O: &cases.OneOf_Y{Y: -1}}, 1, nil},
	{"oneof - filed - invalid (Z)", &cases.OneOf{O: &cases.OneOf_Z{Z: &cases.TestOneOfMsg{}}}, 1, nil},

	{"oneof - required - valid", &cases.OneOfRequired{O: &cases.OneOfRequired_X{X: ""}}, 0, nil},
	{"oneof - require - invalid", &cases.OneOfRequired{}, 1, nil},

	{"oneof - ignore_empty - valid (X)", &cases.OneOfIgnoreEmpty{O: &cases.OneOfIgnoreEmpty_X{X: ""}}, 0, nil},
	{"oneof - ignore_empty - valid (Y)", &cases.OneOfIgnoreEmpty{O: &cases.OneOfIgnoreEmpty_Y{Y: []byte("")}}, 0, nil},
	{"oneof - ignore_empty - valid (Z)", &cases.OneOfIgnoreEmpty{O: &cases.OneOfIgnoreEmpty_Z{Z: 0}}, 0, nil},
}

var wrapperCases = []TestCase{
	{"wrapper - none - valid", &cases.WrapperNone{Val: &wrapperspb.Int32Value{Value: 123}}, 0, nil},
	{"wrapper - none - valid (empty)", &cases.WrapperNone{Val: nil}, 0, nil},

	{"wrapper - float - valid", &cases.WrapperFloat{Val: &wrapperspb.FloatValue{Value: 1}}, 0, nil},
	{"wrapper - float - valid (empty)", &cases.WrapperFloat{Val: nil}, 0, nil},
	{"wrapper - float - invalid", &cases.WrapperFloat{Val: &wrapperspb.FloatValue{Value: 0}}, 1, nil},

	{"wrapper - double - valid", &cases.WrapperDouble{Val: &wrapperspb.DoubleValue{Value: 1}}, 0, nil},
	{"wrapper - double - valid (empty)", &cases.WrapperDouble{Val: nil}, 0, nil},
	{"wrapper - double - invalid", &cases.WrapperDouble{Val: &wrapperspb.DoubleValue{Value: 0}}, 1, nil},

	{"wrapper - int64 - valid", &cases.WrapperInt64{Val: &wrapperspb.Int64Value{Value: 1}}, 0, nil},
	{"wrapper - int64 - valid (empty)", &cases.WrapperInt64{Val: nil}, 0, nil},
	{"wrapper - int64 - invalid", &cases.WrapperInt64{Val: &wrapperspb.Int64Value{Value: 0}}, 1, nil},

	{"wrapper - int32 - valid", &cases.WrapperInt32{Val: &wrapperspb.Int32Value{Value: 1}}, 0, nil},
	{"wrapper - int32 - valid (empty)", &cases.WrapperInt32{Val: nil}, 0, nil},
	{"wrapper - int32 - invalid", &cases.WrapperInt32{Val: &wrapperspb.Int32Value{Value: 0}}, 1, nil},

	{"wrapper - uint64 - valid", &cases.WrapperUInt64{Val: &wrapperspb.UInt64Value{Value: 1}}, 0, nil},
	{"wrapper - uint64 - valid (empty)", &cases.WrapperUInt64{Val: nil}, 0, nil},
	{"wrapper - uint64 - invalid", &cases.WrapperUInt64{Val: &wrapperspb.UInt64Value{Value: 0}}, 1, nil},

	{"wrapper - uint32 - valid", &cases.WrapperUInt32{Val: &wrapperspb.UInt32Value{Value: 1}}, 0, nil},
	{"wrapper - uint32 - valid (empty)", &cases.WrapperUInt32{Val: nil}, 0, nil},
	{"wrapper - uint32 - invalid", &cases.WrapperUInt32{Val: &wrapperspb.UInt32Value{Value: 0}}, 1, nil},

	{"wrapper - bool - valid", &cases.WrapperBool{Val: &wrapperspb.BoolValue{Value: true}}, 0, nil},
	{"wrapper - bool - valid (empty)", &cases.WrapperBool{Val: nil}, 0, nil},
	{"wrapper - bool - invalid", &cases.WrapperBool{Val: &wrapperspb.BoolValue{Value: false}}, 1, nil},

	{"wrapper - string - valid", &cases.WrapperString{Val: &wrapperspb.StringValue{Value: "foobar"}}, 0, nil},
	{"wrapper - string - valid (empty)", &cases.WrapperString{Val: nil}, 0, nil},
	{"wrapper - string - invalid", &cases.WrapperString{Val: &wrapperspb.StringValue{Value: "fizzbuzz"}}, 1, nil},

	{"wrapper - bytes - valid", &cases.WrapperBytes{Val: &wrapperspb.BytesValue{Value: []byte("foo")}}, 0, nil},
	{"wrapper - bytes - valid (empty)", &cases.WrapperBytes{Val: nil}, 0, nil},
	{"wrapper - bytes - invalid", &cases.WrapperBytes{Val: &wrapperspb.BytesValue{Value: []byte("x")}}, 1, nil},

	{"wrapper - required - string - valid", &cases.WrapperRequiredString{Val: &wrapperspb.StringValue{Value: "bar"}}, 0, nil},
	{"wrapper - required - string - invalid", &cases.WrapperRequiredString{Val: &wrapperspb.StringValue{Value: "foo"}}, 1, nil},
	{"wrapper - required - string - invalid (empty)", &cases.WrapperRequiredString{}, 1, nil},

	{"wrapper - required - string (empty) - valid", &cases.WrapperRequiredEmptyString{Val: &wrapperspb.StringValue{Value: ""}}, 0, nil},
	{"wrapper - required - string (empty) - invalid", &cases.WrapperRequiredEmptyString{Val: &wrapperspb.StringValue{Value: "foo"}}, 1, nil},
	{"wrapper - required - string (empty) - invalid (empty)", &cases.WrapperRequiredEmptyString{}, 1, nil},

	{"wrapper - optional - string (uuid) - valid", &cases.WrapperOptionalUuidString{Val: &wrapperspb.StringValue{Value: "8b72987b-024a-43b3-b4cf-647a1f925c5d"}}, 0, nil},
	{"wrapper - optional - string (uuid) - valid (empty)", &cases.WrapperOptionalUuidString{}, 0, nil},
	{"wrapper - optional - string (uuid) - invalid", &cases.WrapperOptionalUuidString{Val: &wrapperspb.StringValue{Value: "foo"}}, 1, nil},

	{"wrapper - required - float - valid", &cases.WrapperRequiredFloat{Val: &wrapperspb.FloatValue{Value: 1}}, 0, nil},
	{"wrapper - required - float - invalid", &cases.WrapperRequiredFloat{Val: &wrapperspb.FloatValue{Value: -5}}, 1, nil},
	{"wrapper - required - float - invalid (empty)", &cases.WrapperRequiredFloat{}, 1, nil},
}

var durationCases = []TestCase{
	{"duration - none - valid", &cases.DurationNone{Val: &durationpb.Duration{Seconds: 123}}, 0, nil},

	{"duration - required - valid", &cases.DurationRequired{Val: &durationpb.Duration{}}, 0, nil},
	{"duration - required - invalid", &cases.DurationRequired{Val: nil}, 1, nil},

	{"duration - const - valid", &cases.DurationConst{Val: &durationpb.Duration{Seconds: 3}}, 0, nil},
	{"duration - const - valid (empty)", &cases.DurationConst{}, 0, nil},
	{"duration - const - invalid", &cases.DurationConst{Val: &durationpb.Duration{Nanos: 3}}, 1, nil},

	{"duration - in - valid", &cases.DurationIn{Val: &durationpb.Duration{Seconds: 1}}, 0, nil},
	{"duration - in - valid (empty)", &cases.DurationIn{}, 0, nil},
	{"duration - in - invalid", &cases.DurationIn{Val: &durationpb.Duration{}}, 1, nil},

	{"duration - not in - valid", &cases.DurationNotIn{Val: &durationpb.Duration{Nanos: 1}}, 0, nil},
	{"duration - not in - valid (empty)", &cases.DurationNotIn{}, 0, nil},
	{"duration - not in - invalid", &cases.DurationNotIn{Val: &durationpb.Duration{}}, 1, nil},

	{"duration - lt - valid", &cases.DurationLT{Val: &durationpb.Duration{Nanos: -1}}, 0, nil},
	{"duration - lt - valid (empty)", &cases.DurationLT{}, 0, nil},
	{"duration - lt - invalid (equal)", &cases.DurationLT{Val: &durationpb.Duration{}}, 1, nil},
	{"duration - lt - invalid", &cases.DurationLT{Val: &durationpb.Duration{Seconds: 1}}, 1, nil},

	{"duration - lte - valid", &cases.DurationLTE{Val: &durationpb.Duration{}}, 0, nil},
	{"duration - lte - valid (empty)", &cases.DurationLTE{}, 0, nil},
	{"duration - lte - valid (equal)", &cases.DurationLTE{Val: &durationpb.Duration{Seconds: 1}}, 0, nil},
	{"duration - lte - invalid", &cases.DurationLTE{Val: &durationpb.Duration{Seconds: 1, Nanos: 1}}, 1, nil},

	{"duration - gt - valid", &cases.DurationGT{Val: &durationpb.Duration{Seconds: 1}}, 0, nil},
	{"duration - gt - valid (empty)", &cases.DurationGT{}, 0, nil},
	{"duration - gt - invalid (equal)", &cases.DurationGT{Val: &durationpb.Duration{Nanos: 1000}}, 1, nil},
	{"duration - gt - invalid", &cases.DurationGT{Val: &durationpb.Duration{}}, 1, nil},

	{"duration - gte - valid", &cases.DurationGTE{Val: &durationpb.Duration{Seconds: 3}}, 0, nil},
	{"duration - gte - valid (empty)", &cases.DurationGTE{}, 0, nil},
	{"duration - gte - valid (equal)", &cases.DurationGTE{Val: &durationpb.Duration{Nanos: 1000000}}, 0, nil},
	{"duration - gte - invalid", &cases.DurationGTE{Val: &durationpb.Duration{Seconds: -1}}, 1, nil},

	{"duration - gt & lt - valid", &cases.DurationGTLT{Val: &durationpb.Duration{Nanos: 1000}}, 0, nil},
	{"duration - gt & lt - valid (empty)", &cases.DurationGTLT{}, 0, nil},
	{"duration - gt & lt - invalid (above)", &cases.DurationGTLT{Val: &durationpb.Duration{Seconds: 1000}}, 1, nil},
	{"duration - gt & lt - invalid (below)", &cases.DurationGTLT{Val: &durationpb.Duration{Nanos: -1000}}, 1, nil},
	{"duration - gt & lt - invalid (max)", &cases.DurationGTLT{Val: &durationpb.Duration{Seconds: 1}}, 1, nil},
	{"duration - gt & lt - invalid (min)", &cases.DurationGTLT{Val: &durationpb.Duration{}}, 1, nil},

	{"duration - exclusive gt & lt - valid (empty)", &cases.DurationExLTGT{}, 0, nil},
	{"duration - exclusive gt & lt - valid (above)", &cases.DurationExLTGT{Val: &durationpb.Duration{Seconds: 2}}, 0, nil},
	{"duration - exclusive gt & lt - valid (below)", &cases.DurationExLTGT{Val: &durationpb.Duration{Nanos: -1}}, 0, nil},
	{"duration - exclusive gt & lt - invalid", &cases.DurationExLTGT{Val: &durationpb.Duration{Nanos: 1000}}, 1, nil},
	{"duration - exclusive gt & lt - invalid (max)", &cases.DurationExLTGT{Val: &durationpb.Duration{Seconds: 1}}, 1, nil},
	{"duration - exclusive gt & lt - invalid (min)", &cases.DurationExLTGT{Val: &durationpb.Duration{}}, 1, nil},

	{"duration - gte & lte - valid", &cases.DurationGTELTE{Val: &durationpb.Duration{Seconds: 60, Nanos: 1}}, 0, nil},
	{"duration - gte & lte - valid (empty)", &cases.DurationGTELTE{}, 0, nil},
	{"duration - gte & lte - valid (max)", &cases.DurationGTELTE{Val: &durationpb.Duration{Seconds: 3600}}, 0, nil},
	{"duration - gte & lte - valid (min)", &cases.DurationGTELTE{Val: &durationpb.Duration{Seconds: 60}}, 0, nil},
	{"duration - gte & lte - invalid (above)", &cases.DurationGTELTE{Val: &durationpb.Duration{Seconds: 3600, Nanos: 1}}, 1, nil},
	{"duration - gte & lte - invalid (below)", &cases.DurationGTELTE{Val: &durationpb.Duration{Seconds: 59}}, 1, nil},

	{"duration - gte & lte - valid (empty)", &cases.DurationExGTELTE{}, 0, nil},
	{"duration - exclusive gte & lte - valid (above)", &cases.DurationExGTELTE{Val: &durationpb.Duration{Seconds: 3601}}, 0, nil},
	{"duration - exclusive gte & lte - valid (below)", &cases.DurationExGTELTE{Val: &durationpb.Duration{}}, 0, nil},
	{"duration - exclusive gte & lte - valid (max)", &cases.DurationExGTELTE{Val: &durationpb.Duration{Seconds: 3600}}, 0, nil},
	{"duration - exclusive gte & lte - valid (min)", &cases.DurationExGTELTE{Val: &durationpb.Duration{Seconds: 60}}, 0, nil},
	{"duration - exclusive gte & lte - invalid", &cases.DurationExGTELTE{Val: &durationpb.Duration{Seconds: 61}}, 1, nil},
	{"duration - fields with other fields - invalid other field", &cases.DurationFieldWithOtherFields{DurationVal: nil, IntVal: 12}, 1, nil},
}

var timestampCases = []TestCase{
	{"timestamp - none - valid", &cases.TimestampNone{Val: &timestamppb.Timestamp{Seconds: 123}}, 0, nil},

	{"timestamp - required - valid", &cases.TimestampRequired{Val: &timestamppb.Timestamp{}}, 0, nil},
	{"timestamp - required - invalid", &cases.TimestampRequired{Val: nil}, 1, nil},

	{"timestamp - const - valid", &cases.TimestampConst{Val: &timestamppb.Timestamp{Seconds: 3}}, 0, nil},
	{"timestamp - const - valid (empty)", &cases.TimestampConst{}, 0, nil},
	{"timestamp - const - invalid", &cases.TimestampConst{Val: &timestamppb.Timestamp{Nanos: 3}}, 1, nil},

	{"timestamp - lt - valid", &cases.TimestampLT{Val: &timestamppb.Timestamp{Seconds: -1}}, 0, nil},
	{"timestamp - lt - valid (empty)", &cases.TimestampLT{}, 0, nil},
	{"timestamp - lt - invalid (equal)", &cases.TimestampLT{Val: &timestamppb.Timestamp{}}, 1, nil},
	{"timestamp - lt - invalid", &cases.TimestampLT{Val: &timestamppb.Timestamp{Seconds: 1}}, 1, nil},

	{"timestamp - lte - valid", &cases.TimestampLTE{Val: &timestamppb.Timestamp{}}, 0, nil},
	{"timestamp - lte - valid (empty)", &cases.TimestampLTE{}, 0, nil},
	{"timestamp - lte - valid (equal)", &cases.TimestampLTE{Val: &timestamppb.Timestamp{Seconds: 1}}, 0, nil},
	{"timestamp - lte - invalid", &cases.TimestampLTE{Val: &timestamppb.Timestamp{Seconds: 1, Nanos: 1}}, 1, nil},

	{"timestamp - gt - valid", &cases.TimestampGT{Val: &timestamppb.Timestamp{Seconds: 1}}, 0, nil},
	{"timestamp - gt - valid (empty)", &cases.TimestampGT{}, 0, nil},
	{"timestamp - gt - invalid (equal)", &cases.TimestampGT{Val: &timestamppb.Timestamp{Nanos: 1000}}, 1, nil},
	{"timestamp - gt - invalid", &cases.TimestampGT{Val: &timestamppb.Timestamp{}}, 1, nil},

	{"timestamp - gte - valid", &cases.TimestampGTE{Val: &timestamppb.Timestamp{Seconds: 3}}, 0, nil},
	{"timestamp - gte - valid (empty)", &cases.TimestampGTE{}, 0, nil},
	{"timestamp - gte - valid (equal)", &cases.TimestampGTE{Val: &timestamppb.Timestamp{Nanos: 1000000}}, 0, nil},
	{"timestamp - gte - invalid", &cases.TimestampGTE{Val: &timestamppb.Timestamp{Seconds: -1}}, 1, nil},

	{"timestamp - gt & lt - valid", &cases.TimestampGTLT{Val: &timestamppb.Timestamp{Nanos: 1000}}, 0, nil},
	{"timestamp - gt & lt - valid (empty)", &cases.TimestampGTLT{}, 0, nil},
	{"timestamp - gt & lt - invalid (above)", &cases.TimestampGTLT{Val: &timestamppb.Timestamp{Seconds: 1000}}, 1, nil},
	{"timestamp - gt & lt - invalid (below)", &cases.TimestampGTLT{Val: &timestamppb.Timestamp{Seconds: -1000}}, 1, nil},
	{"timestamp - gt & lt - invalid (max)", &cases.TimestampGTLT{Val: &timestamppb.Timestamp{Seconds: 1}}, 1, nil},
	{"timestamp - gt & lt - invalid (min)", &cases.TimestampGTLT{Val: &timestamppb.Timestamp{}}, 1, nil},

	{"timestamp - exclusive gt & lt - valid (empty)", &cases.TimestampExLTGT{}, 0, nil},
	{"timestamp - exclusive gt & lt - valid (above)", &cases.TimestampExLTGT{Val: &timestamppb.Timestamp{Seconds: 2}}, 0, nil},
	{"timestamp - exclusive gt & lt - valid (below)", &cases.TimestampExLTGT{Val: &timestamppb.Timestamp{Seconds: -1}}, 0, nil},
	{"timestamp - exclusive gt & lt - invalid", &cases.TimestampExLTGT{Val: &timestamppb.Timestamp{Nanos: 1000}}, 1, nil},
	{"timestamp - exclusive gt & lt - invalid (max)", &cases.TimestampExLTGT{Val: &timestamppb.Timestamp{Seconds: 1}}, 1, nil},
	{"timestamp - exclusive gt & lt - invalid (min)", &cases.TimestampExLTGT{Val: &timestamppb.Timestamp{}}, 1, nil},

	{"timestamp - gte & lte - valid", &cases.TimestampGTELTE{Val: &timestamppb.Timestamp{Seconds: 60, Nanos: 1}}, 0, nil},
	{"timestamp - gte & lte - valid (empty)", &cases.TimestampGTELTE{}, 0, nil},
	{"timestamp - gte & lte - valid (max)", &cases.TimestampGTELTE{Val: &timestamppb.Timestamp{Seconds: 3600}}, 0, nil},
	{"timestamp - gte & lte - valid (min)", &cases.TimestampGTELTE{Val: &timestamppb.Timestamp{Seconds: 60}}, 0, nil},
	{"timestamp - gte & lte - invalid (above)", &cases.TimestampGTELTE{Val: &timestamppb.Timestamp{Seconds: 3600, Nanos: 1}}, 1, nil},
	{"timestamp - gte & lte - invalid (below)", &cases.TimestampGTELTE{Val: &timestamppb.Timestamp{Seconds: 59}}, 1, nil},

	{"timestamp - gte & lte - valid (empty)", &cases.TimestampExGTELTE{}, 0, nil},
	{"timestamp - exclusive gte & lte - valid (above)", &cases.TimestampExGTELTE{Val: &timestamppb.Timestamp{Seconds: 3601}}, 0, nil},
	{"timestamp - exclusive gte & lte - valid (below)", &cases.TimestampExGTELTE{Val: &timestamppb.Timestamp{}}, 0, nil},
	{"timestamp - exclusive gte & lte - valid (max)", &cases.TimestampExGTELTE{Val: &timestamppb.Timestamp{Seconds: 3600}}, 0, nil},
	{"timestamp - exclusive gte & lte - valid (min)", &cases.TimestampExGTELTE{Val: &timestamppb.Timestamp{Seconds: 60}}, 0, nil},
	{"timestamp - exclusive gte & lte - invalid", &cases.TimestampExGTELTE{Val: &timestamppb.Timestamp{Seconds: 61}}, 1, nil},

	{"timestamp - lt now - valid", &cases.TimestampLTNow{Val: &timestamppb.Timestamp{}}, 0, nil},
	{"timestamp - lt now - valid (empty)", &cases.TimestampLTNow{}, 0, nil},
	{"timestamp - lt now - invalid", &cases.TimestampLTNow{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 7200}}, 1, nil},

	{"timestamp - gt now - valid", &cases.TimestampGTNow{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 7200}}, 0, nil},
	{"timestamp - gt now - valid (empty)", &cases.TimestampGTNow{}, 0, nil},
	{"timestamp - gt now - invalid", &cases.TimestampGTNow{Val: &timestamppb.Timestamp{}}, 1, nil},

	{"timestamp - within - valid", &cases.TimestampWithin{Val: timestamppb.Now()}, 0, nil},
	{"timestamp - within - valid (empty)", &cases.TimestampWithin{}, 0, nil},
	{"timestamp - within - invalid (below)", &cases.TimestampWithin{Val: &timestamppb.Timestamp{}}, 1, nil},
	{"timestamp - within - invalid (above)", &cases.TimestampWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 7200}}, 1, nil},

	{"timestamp - lt now within - valid", &cases.TimestampLTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() - 1800}}, 0, nil},
	{"timestamp - lt now within - valid (empty)", &cases.TimestampLTNowWithin{}, 0, nil},
	{"timestamp - lt now within - invalid (lt)", &cases.TimestampLTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 1800}}, 1, nil},
	{"timestamp - lt now within - invalid (within)", &cases.TimestampLTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() - 7200}}, 1, nil},

	{"timestamp - gt now within - valid", &cases.TimestampGTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 1800}}, 0, nil},
	{"timestamp - gt now within - valid (empty)", &cases.TimestampGTNowWithin{}, 0, nil},
	{"timestamp - gt now within - invalid (gt)", &cases.TimestampGTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() - 1800}}, 1, nil},
	{"timestamp - gt now within - invalid (within)", &cases.TimestampGTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 7200}}, 1, nil},
}

var anyCases = []TestCase{
	{"any - none - valid", &cases.AnyNone{Val: &anypb.Any{}}, 0, nil},

	{"any - required - valid", &cases.AnyRequired{Val: &anypb.Any{}}, 0, nil},
	{"any - required - invalid", &cases.AnyRequired{Val: nil}, 1, nil},

	{"any - in - valid", &cases.AnyIn{Val: &anypb.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}}, 0, nil},
	{"any - in - valid (empty)", &cases.AnyIn{}, 0, nil},
	{"any - in - invalid", &cases.AnyIn{Val: &anypb.Any{TypeUrl: "type.googleapis.com/google.protobuf.Timestamp"}}, 1, nil},

	{"any - not in - valid", &cases.AnyNotIn{Val: &anypb.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}}, 0, nil},
	{"any - not in - valid (empty)", &cases.AnyNotIn{}, 0, nil},
	{"any - not in - invalid", &cases.AnyNotIn{Val: &anypb.Any{TypeUrl: "type.googleapis.com/google.protobuf.Timestamp"}}, 1, nil},
}

var kitchenSink = []TestCase{
	{"kitchensink - field - valid", &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Const: "abcd", IntConst: 5, BoolConst: false, FloatVal: &wrapperspb.FloatValue{Value: 1}, DurVal: &durationpb.Duration{Seconds: 3}, TsVal: &timestamppb.Timestamp{Seconds: 17}, FloatConst: 7, DoubleIn: 123, EnumConst: cases.ComplexTestEnum_ComplexTWO, AnyVal: &anypb.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}, RepTsVal: []*timestamppb.Timestamp{{Seconds: 3}}, MapVal: map[int32]string{-1: "a", -2: "b"}, BytesVal: []byte("\x00\x99"), O: &cases.ComplexTestMsg_X{X: "foobar"}}}, 0, nil},
	{"kitchensink - valid (unset)", &cases.KitchenSinkMessage{}, 0, nil},
	{"kitchensink - field - invalid", &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{}}, 7, nil},
	{"kitchensink - field - embedded - invalid", &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Another: &cases.ComplexTestMsg{}}}, 14, nil},
	{"kitchensink - field - invalid (transitive)", &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Const: "abcd", BoolConst: true, Nested: &cases.ComplexTestMsg{}}}, 14, nil},
	{"kitchensink - many - all non-message fields invalid", &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{BoolConst: true, FloatVal: &wrapperspb.FloatValue{}, TsVal: &timestamppb.Timestamp{}, FloatConst: 8, AnyVal: &anypb.Any{TypeUrl: "asdf"}, RepTsVal: []*timestamppb.Timestamp{{Nanos: 1}}}}, 13, nil},
}

var nestedCases = []TestCase{
	{"nested wkt uuid - field - valid", &cases.WktLevelOne{Two: &cases.WktLevelOne_WktLevelTwo{Three: &cases.WktLevelOne_WktLevelTwo_WktLevelThree{Uuid: "f81d16ef-40e2-40c6-bebc-89aaf5292f9a"}}}, 0, nil},
	{"nested wkt uuid - field - invalid", &cases.WktLevelOne{Two: &cases.WktLevelOne_WktLevelTwo{Three: &cases.WktLevelOne_WktLevelTwo_WktLevelThree{Uuid: "not-a-valid-uuid"}}}, 1, nil},
}
