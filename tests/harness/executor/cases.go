package main

import (
	"math"
	"time"

	cases "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go"
	other_package "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/go"
	"google.golang.org/protobuf/proto"
	any "google.golang.org/protobuf/types/known/anypb"
	duration "google.golang.org/protobuf/types/known/durationpb"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
	wrappers "google.golang.org/protobuf/types/known/wrapperspb"
)

type TestCase struct {
	Name     string
	Message  proto.Message
	Valid    bool
	ErrorMsg string
	Failures int // expected number of failed validation errors
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
	{"float - none - valid", &cases.FloatNone{Val: -1.23456}, true, "", 0},

	{"float - const - valid", &cases.FloatConst{Val: 1.23}, true, "", 0},
	{"float - const - invalid", &cases.FloatConst{Val: 4.56}, false, "invalid FloatConst.Val: value must equal 1.23", 1},

	{"float - in - valid", &cases.FloatIn{Val: 7.89}, true, "", 0},
	{"float - in - invalid", &cases.FloatIn{Val: 10.11}, false, "invalid FloatIn.Val: value must be in list [4.56 7.89]", 1},

	{"float - not in - valid", &cases.FloatNotIn{Val: 1}, true, "", 0},
	{"float - not in - invalid", &cases.FloatNotIn{Val: 0}, false, "invalid FloatNotIn.Val: value must not be in list [0]", 1},

	{"float - lt - valid", &cases.FloatLT{Val: -1}, true, "", 0},
	{"float - lt - invalid (equal)", &cases.FloatLT{Val: 0}, false, "invalid FloatLT.Val: value must be less than 0", 1},
	{"float - lt - invalid", &cases.FloatLT{Val: 1}, false, "invalid FloatLT.Val: value must be less than 0", 1},

	{"float - lte - valid", &cases.FloatLTE{Val: 63}, true, "", 0},
	{"float - lte - valid (equal)", &cases.FloatLTE{Val: 64}, true, "", 0},
	{"float - lte - invalid", &cases.FloatLTE{Val: 65}, false, "invalid FloatLTE.Val: value must be less than or equal to 64", 1},

	{"float - gt - valid", &cases.FloatGT{Val: 17}, true, "", 0},
	{"float - gt - invalid (equal)", &cases.FloatGT{Val: 16}, false, "invalid FloatGT.Val: value must be greater than 16", 1},
	{"float - gt - invalid", &cases.FloatGT{Val: 15}, false, "invalid FloatGT.Val: value must be greater than 16", 1},

	{"float - gte - valid", &cases.FloatGTE{Val: 9}, true, "", 0},
	{"float - gte - valid (equal)", &cases.FloatGTE{Val: 8}, true, "", 0},
	{"float - gte - invalid", &cases.FloatGTE{Val: 7}, false, "invalid FloatGTE.Val: value must be greater than or equal to 8", 1},

	{"float - gt & lt - valid", &cases.FloatGTLT{Val: 5}, true, "", 0},
	{"float - gt & lt - invalid (above)", &cases.FloatGTLT{Val: 11}, false, "invalid FloatGTLT.Val: value must be inside range (0, 10)", 1},
	{"float - gt & lt - invalid (below)", &cases.FloatGTLT{Val: -1}, false, "invalid FloatGTLT.Val: value must be inside range (0, 10)", 1},
	{"float - gt & lt - invalid (max)", &cases.FloatGTLT{Val: 10}, false, "invalid FloatGTLT.Val: value must be inside range (0, 10)", 1},
	{"float - gt & lt - invalid (min)", &cases.FloatGTLT{Val: 0}, false, "invalid FloatGTLT.Val: value must be inside range (0, 10)", 1},

	{"float - exclusive gt & lt - valid (above)", &cases.FloatExLTGT{Val: 11}, true, "", 0},
	{"float - exclusive gt & lt - valid (below)", &cases.FloatExLTGT{Val: -1}, true, "", 0},
	{"float - exclusive gt & lt - invalid", &cases.FloatExLTGT{Val: 5}, false, "invalid FloatExLTGT.Val: value must be outside range [0, 10]", 1},
	{"float - exclusive gt & lt - invalid (max)", &cases.FloatExLTGT{Val: 10}, false, "invalid FloatExLTGT.Val: value must be outside range [0, 10]", 1},
	{"float - exclusive gt & lt - invalid (min)", &cases.FloatExLTGT{Val: 0}, false, "invalid FloatExLTGT.Val: value must be outside range [0, 10]", 1},

	{"float - gte & lte - valid", &cases.FloatGTELTE{Val: 200}, true, "", 0},
	{"float - gte & lte - valid (max)", &cases.FloatGTELTE{Val: 256}, true, "", 0},
	{"float - gte & lte - valid (min)", &cases.FloatGTELTE{Val: 128}, true, "", 0},
	{"float - gte & lte - invalid (above)", &cases.FloatGTELTE{Val: 300}, false, "invalid FloatGTELTE.Val: value must be inside range [128, 256]", 1},
	{"float - gte & lte - invalid (below)", &cases.FloatGTELTE{Val: 100}, false, "invalid FloatGTELTE.Val: value must be inside range [128, 256]", 1},

	{"float - exclusive gte & lte - valid (above)", &cases.FloatExGTELTE{Val: 300}, true, "", 0},
	{"float - exclusive gte & lte - valid (below)", &cases.FloatExGTELTE{Val: 100}, true, "", 0},
	{"float - exclusive gte & lte - valid (max)", &cases.FloatExGTELTE{Val: 256}, true, "", 0},
	{"float - exclusive gte & lte - valid (min)", &cases.FloatExGTELTE{Val: 128}, true, "", 0},
	{"float - exclusive gte & lte - invalid", &cases.FloatExGTELTE{Val: 200}, false, "invalid FloatExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var doubleCases = []TestCase{
	{"double - none - valid", &cases.DoubleNone{Val: -1.23456}, true, "", 0},

	{"double - const - valid", &cases.DoubleConst{Val: 1.23}, true, "", 0},
	{"double - const - invalid", &cases.DoubleConst{Val: 4.56}, false, "invalid DoubleConst.Val: value must equal 1.23", 1},

	{"double - in - valid", &cases.DoubleIn{Val: 7.89}, true, "", 0},
	{"double - in - invalid", &cases.DoubleIn{Val: 10.11}, false, "invalid DoubleIn.Val: value must be in list [4.56 7.89]", 1},

	{"double - not in - valid", &cases.DoubleNotIn{Val: 1}, true, "", 0},
	{"double - not in - invalid", &cases.DoubleNotIn{Val: 0}, false, "invalid DoubleNotIn.Val: value must not be in list [0]", 1},

	{"double - lt - valid", &cases.DoubleLT{Val: -1}, true, "", 0},
	{"double - lt - invalid (equal)", &cases.DoubleLT{Val: 0}, false, "invalid DoubleLT.Val: value must be less than 0", 1},
	{"double - lt - invalid", &cases.DoubleLT{Val: 1}, false, "invalid DoubleLT.Val: value must be less than 0", 1},

	{"double - lte - valid", &cases.DoubleLTE{Val: 63}, true, "", 0},
	{"double - lte - valid (equal)", &cases.DoubleLTE{Val: 64}, true, "", 0},
	{"double - lte - invalid", &cases.DoubleLTE{Val: 65}, false, "invalid DoubleLTE.Val: value must be less than or equal to 64", 1},

	{"double - gt - valid", &cases.DoubleGT{Val: 17}, true, "", 0},
	{"double - gt - invalid (equal)", &cases.DoubleGT{Val: 16}, false, "invalid DoubleGT.Val: value must be greater than 16", 1},
	{"double - gt - invalid", &cases.DoubleGT{Val: 15}, false, "invalid DoubleGT.Val: value must be greater than 16", 1},

	{"double - gte - valid", &cases.DoubleGTE{Val: 9}, true, "", 0},
	{"double - gte - valid (equal)", &cases.DoubleGTE{Val: 8}, true, "", 0},
	{"double - gte - invalid", &cases.DoubleGTE{Val: 7}, false, "invalid DoubleGTE.Val: value must be greater than or equal to 8", 1},

	{"double - gt & lt - valid", &cases.DoubleGTLT{Val: 5}, true, "", 0},
	{"double - gt & lt - invalid (above)", &cases.DoubleGTLT{Val: 11}, false, "invalid DoubleGTLT.Val: value must be inside range (0, 10)", 1},
	{"double - gt & lt - invalid (below)", &cases.DoubleGTLT{Val: -1}, false, "invalid DoubleGTLT.Val: value must be inside range (0, 10)", 1},
	{"double - gt & lt - invalid (max)", &cases.DoubleGTLT{Val: 10}, false, "invalid DoubleGTLT.Val: value must be inside range (0, 10)", 1},
	{"double - gt & lt - invalid (min)", &cases.DoubleGTLT{Val: 0}, false, "invalid DoubleGTLT.Val: value must be inside range (0, 10)", 1},

	{"double - exclusive gt & lt - valid (above)", &cases.DoubleExLTGT{Val: 11}, true, "", 0},
	{"double - exclusive gt & lt - valid (below)", &cases.DoubleExLTGT{Val: -1}, true, "", 0},
	{"double - exclusive gt & lt - invalid", &cases.DoubleExLTGT{Val: 5}, false, "invalid DoubleExLTGT.Val: value must be outside range [0, 10]", 1},
	{"double - exclusive gt & lt - invalid (max)", &cases.DoubleExLTGT{Val: 10}, false, "invalid DoubleExLTGT.Val: value must be outside range [0, 10]", 1},
	{"double - exclusive gt & lt - invalid (min)", &cases.DoubleExLTGT{Val: 0}, false, "invalid DoubleExLTGT.Val: value must be outside range [0, 10]", 1},

	{"double - gte & lte - valid", &cases.DoubleGTELTE{Val: 200}, true, "", 0},
	{"double - gte & lte - valid (max)", &cases.DoubleGTELTE{Val: 256}, true, "", 0},
	{"double - gte & lte - valid (min)", &cases.DoubleGTELTE{Val: 128}, true, "", 0},
	{"double - gte & lte - invalid (above)", &cases.DoubleGTELTE{Val: 300}, false, "invalid DoubleGTELTE.Val: value must be inside range [128, 256]", 1},
	{"double - gte & lte - invalid (below)", &cases.DoubleGTELTE{Val: 100}, false, "invalid DoubleGTELTE.Val: value must be inside range [128, 256]", 1},

	{"double - exclusive gte & lte - valid (above)", &cases.DoubleExGTELTE{Val: 300}, true, "", 0},
	{"double - exclusive gte & lte - valid (below)", &cases.DoubleExGTELTE{Val: 100}, true, "", 0},
	{"double - exclusive gte & lte - valid (max)", &cases.DoubleExGTELTE{Val: 256}, true, "", 0},
	{"double - exclusive gte & lte - valid (min)", &cases.DoubleExGTELTE{Val: 128}, true, "", 0},
	{"double - exclusive gte & lte - invalid", &cases.DoubleExGTELTE{Val: 200}, false, "invalid DoubleExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var int32Cases = []TestCase{
	{"int32 - none - valid", &cases.Int32None{Val: 123}, true, "", 0},

	{"int32 - const - valid", &cases.Int32Const{Val: 1}, true, "", 0},
	{"int32 - const - invalid", &cases.Int32Const{Val: 2}, false, "invalid Int32Const.Val: value must equal 1", 1},

	{"int32 - in - valid", &cases.Int32In{Val: 3}, true, "", 0},
	{"int32 - in - invalid", &cases.Int32In{Val: 5}, false, "invalid Int32In.Val: value must be in list [2 3]", 1},

	{"int32 - not in - valid", &cases.Int32NotIn{Val: 1}, true, "", 0},
	{"int32 - not in - invalid", &cases.Int32NotIn{Val: 0}, false, "invalid Int32NotIn.Val: value must not be in list [0]", 1},

	{"int32 - lt - valid", &cases.Int32LT{Val: -1}, true, "", 0},
	{"int32 - lt - invalid (equal)", &cases.Int32LT{Val: 0}, false, "invalid Int32LT.Val: value must be less than 0", 1},
	{"int32 - lt - invalid", &cases.Int32LT{Val: 1}, false, "invalid Int32LT.Val: value must be less than 0", 1},

	{"int32 - lte - valid", &cases.Int32LTE{Val: 63}, true, "", 0},
	{"int32 - lte - valid (equal)", &cases.Int32LTE{Val: 64}, true, "", 0},
	{"int32 - lte - invalid", &cases.Int32LTE{Val: 65}, false, "invalid Int32LTE.Val: value must be less than or equal to 64", 1},

	{"int32 - gt - valid", &cases.Int32GT{Val: 17}, true, "", 0},
	{"int32 - gt - invalid (equal)", &cases.Int32GT{Val: 16}, false, "invalid Int32GT.Val: value must be greater than 16", 1},
	{"int32 - gt - invalid", &cases.Int32GT{Val: 15}, false, "invalid Int32GT.Val: value must be greater than 16", 1},

	{"int32 - gte - valid", &cases.Int32GTE{Val: 9}, true, "", 0},
	{"int32 - gte - valid (equal)", &cases.Int32GTE{Val: 8}, true, "", 0},
	{"int32 - gte - invalid", &cases.Int32GTE{Val: 7}, false, "invalid Int32GTE.Val: value must be greater than or equal to 8", 1},

	{"int32 - gt & lt - valid", &cases.Int32GTLT{Val: 5}, true, "", 0},
	{"int32 - gt & lt - invalid (above)", &cases.Int32GTLT{Val: 11}, false, "invalid Int32GTLT.Val: value must be inside range (0, 10)", 1},
	{"int32 - gt & lt - invalid (below)", &cases.Int32GTLT{Val: -1}, false, "invalid Int32GTLT.Val: value must be inside range (0, 10)", 1},
	{"int32 - gt & lt - invalid (max)", &cases.Int32GTLT{Val: 10}, false, "invalid Int32GTLT.Val: value must be inside range (0, 10)", 1},
	{"int32 - gt & lt - invalid (min)", &cases.Int32GTLT{Val: 0}, false, "invalid Int32GTLT.Val: value must be inside range (0, 10)", 1},

	{"int32 - exclusive gt & lt - valid (above)", &cases.Int32ExLTGT{Val: 11}, true, "", 0},
	{"int32 - exclusive gt & lt - valid (below)", &cases.Int32ExLTGT{Val: -1}, true, "", 0},
	{"int32 - exclusive gt & lt - invalid", &cases.Int32ExLTGT{Val: 5}, false, "invalid Int32ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"int32 - exclusive gt & lt - invalid (max)", &cases.Int32ExLTGT{Val: 10}, false, "invalid Int32ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"int32 - exclusive gt & lt - invalid (min)", &cases.Int32ExLTGT{Val: 0}, false, "invalid Int32ExLTGT.Val: value must be outside range [0, 10]", 1},

	{"int32 - gte & lte - valid", &cases.Int32GTELTE{Val: 200}, true, "", 0},
	{"int32 - gte & lte - valid (max)", &cases.Int32GTELTE{Val: 256}, true, "", 0},
	{"int32 - gte & lte - valid (min)", &cases.Int32GTELTE{Val: 128}, true, "", 0},
	{"int32 - gte & lte - invalid (above)", &cases.Int32GTELTE{Val: 300}, false, "invalid Int32GTELTE.Val: value must be inside range [128, 256]", 1},
	{"int32 - gte & lte - invalid (below)", &cases.Int32GTELTE{Val: 100}, false, "invalid Int32GTELTE.Val: value must be inside range [128, 256]", 1},

	{"int32 - exclusive gte & lte - valid (above)", &cases.Int32ExGTELTE{Val: 300}, true, "", 0},
	{"int32 - exclusive gte & lte - valid (below)", &cases.Int32ExGTELTE{Val: 100}, true, "", 0},
	{"int32 - exclusive gte & lte - valid (max)", &cases.Int32ExGTELTE{Val: 256}, true, "", 0},
	{"int32 - exclusive gte & lte - valid (min)", &cases.Int32ExGTELTE{Val: 128}, true, "", 0},
	{"int32 - exclusive gte & lte - invalid", &cases.Int32ExGTELTE{Val: 200}, false, "invalid Int32ExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var int64Cases = []TestCase{
	{"int64 - none - valid", &cases.Int64None{Val: 123}, true, "", 0},

	{"int64 - const - valid", &cases.Int64Const{Val: 1}, true, "", 0},
	{"int64 - const - invalid", &cases.Int64Const{Val: 2}, false, "invalid Int64Const.Val: value must equal 1", 1},

	{"int64 - in - valid", &cases.Int64In{Val: 3}, true, "", 0},
	{"int64 - in - invalid", &cases.Int64In{Val: 5}, false, "invalid Int64In.Val: value must be in list [2 3]", 1},

	{"int64 - not in - valid", &cases.Int64NotIn{Val: 1}, true, "", 0},
	{"int64 - not in - invalid", &cases.Int64NotIn{Val: 0}, false, "invalid Int64NotIn.Val: value must not be in list [0]", 1},

	{"int64 - lt - valid", &cases.Int64LT{Val: -1}, true, "", 0},
	{"int64 - lt - invalid (equal)", &cases.Int64LT{Val: 0}, false, "invalid Int64LT.Val: value must be less than 0", 1},
	{"int64 - lt - invalid", &cases.Int64LT{Val: 1}, false, "invalid Int64LT.Val: value must be less than 0", 1},

	{"int64 - lte - valid", &cases.Int64LTE{Val: 63}, true, "", 0},
	{"int64 - lte - valid (equal)", &cases.Int64LTE{Val: 64}, true, "", 0},
	{"int64 - lte - invalid", &cases.Int64LTE{Val: 65}, false, "invalid Int64LTE.Val: value must be less than or equal to 64", 1},

	{"int64 - gt - valid", &cases.Int64GT{Val: 17}, true, "", 0},
	{"int64 - gt - invalid (equal)", &cases.Int64GT{Val: 16}, false, "invalid Int64GT.Val: value must be greater than 16", 1},
	{"int64 - gt - invalid", &cases.Int64GT{Val: 15}, false, "invalid Int64GT.Val: value must be greater than 16", 1},

	{"int64 - gte - valid", &cases.Int64GTE{Val: 9}, true, "", 0},
	{"int64 - gte - valid (equal)", &cases.Int64GTE{Val: 8}, true, "", 0},
	{"int64 - gte - invalid", &cases.Int64GTE{Val: 7}, false, "invalid Int64GTE.Val: value must be greater than or equal to 8", 1},

	{"int64 - gt & lt - valid", &cases.Int64GTLT{Val: 5}, true, "", 0},
	{"int64 - gt & lt - invalid (above)", &cases.Int64GTLT{Val: 11}, false, "invalid Int64GTLT.Val: value must be inside range (0, 10)", 1},
	{"int64 - gt & lt - invalid (below)", &cases.Int64GTLT{Val: -1}, false, "invalid Int64GTLT.Val: value must be inside range (0, 10)", 1},
	{"int64 - gt & lt - invalid (max)", &cases.Int64GTLT{Val: 10}, false, "invalid Int64GTLT.Val: value must be inside range (0, 10)", 1},
	{"int64 - gt & lt - invalid (min)", &cases.Int64GTLT{Val: 0}, false, "invalid Int64GTLT.Val: value must be inside range (0, 10)", 1},

	{"int64 - exclusive gt & lt - valid (above)", &cases.Int64ExLTGT{Val: 11}, true, "", 0},
	{"int64 - exclusive gt & lt - valid (below)", &cases.Int64ExLTGT{Val: -1}, true, "", 0},
	{"int64 - exclusive gt & lt - invalid", &cases.Int64ExLTGT{Val: 5}, false, "invalid Int64ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"int64 - exclusive gt & lt - invalid (max)", &cases.Int64ExLTGT{Val: 10}, false, "invalid Int64ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"int64 - exclusive gt & lt - invalid (min)", &cases.Int64ExLTGT{Val: 0}, false, "invalid Int64ExLTGT.Val: value must be outside range [0, 10]", 1},

	{"int64 - gte & lte - valid", &cases.Int64GTELTE{Val: 200}, true, "", 0},
	{"int64 - gte & lte - valid (max)", &cases.Int64GTELTE{Val: 256}, true, "", 0},
	{"int64 - gte & lte - valid (min)", &cases.Int64GTELTE{Val: 128}, true, "", 0},
	{"int64 - gte & lte - invalid (above)", &cases.Int64GTELTE{Val: 300}, false, "invalid Int64GTELTE.Val: value must be inside range [128, 256]", 1},
	{"int64 - gte & lte - invalid (below)", &cases.Int64GTELTE{Val: 100}, false, "invalid Int64GTELTE.Val: value must be inside range [128, 256]", 1},

	{"int64 - exclusive gte & lte - valid (above)", &cases.Int64ExGTELTE{Val: 300}, true, "", 0},
	{"int64 - exclusive gte & lte - valid (below)", &cases.Int64ExGTELTE{Val: 100}, true, "", 0},
	{"int64 - exclusive gte & lte - valid (max)", &cases.Int64ExGTELTE{Val: 256}, true, "", 0},
	{"int64 - exclusive gte & lte - valid (min)", &cases.Int64ExGTELTE{Val: 128}, true, "", 0},
	{"int64 - exclusive gte & lte - invalid", &cases.Int64ExGTELTE{Val: 200}, false, "invalid Int64ExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var uint32Cases = []TestCase{
	{"uint32 - none - valid", &cases.UInt32None{Val: 123}, true, "", 0},

	{"uint32 - const - valid", &cases.UInt32Const{Val: 1}, true, "", 0},
	{"uint32 - const - invalid", &cases.UInt32Const{Val: 2}, false, "invalid UInt32Const.Val: value must equal 1", 1},

	{"uint32 - in - valid", &cases.UInt32In{Val: 3}, true, "", 0},
	{"uint32 - in - invalid", &cases.UInt32In{Val: 5}, false, "invalid UInt32In.Val: value must be in list [2 3]", 1},

	{"uint32 - not in - valid", &cases.UInt32NotIn{Val: 1}, true, "", 0},
	{"uint32 - not in - invalid", &cases.UInt32NotIn{Val: 0}, false, "invalid UInt32NotIn.Val: value must not be in list [0]", 1},

	{"uint32 - lt - valid", &cases.UInt32LT{Val: 4}, true, "", 0},
	{"uint32 - lt - invalid (equal)", &cases.UInt32LT{Val: 5}, false, "invalid UInt32LT.Val: value must be less than 5", 1},
	{"uint32 - lt - invalid", &cases.UInt32LT{Val: 6}, false, "invalid UInt32LT.Val: value must be less than 5", 1},

	{"uint32 - lte - valid", &cases.UInt32LTE{Val: 63}, true, "", 0},
	{"uint32 - lte - valid (equal)", &cases.UInt32LTE{Val: 64}, true, "", 0},
	{"uint32 - lte - invalid", &cases.UInt32LTE{Val: 65}, false, "invalid UInt32LTE.Val: value must be less than or equal to 64", 1},

	{"uint32 - gt - valid", &cases.UInt32GT{Val: 17}, true, "", 0},
	{"uint32 - gt - invalid (equal)", &cases.UInt32GT{Val: 16}, false, "invalid UInt32GT.Val: value must be greater than 16", 1},
	{"uint32 - gt - invalid", &cases.UInt32GT{Val: 15}, false, "invalid UInt32GT.Val: value must be greater than 16", 1},

	{"uint32 - gte - valid", &cases.UInt32GTE{Val: 9}, true, "", 0},
	{"uint32 - gte - valid (equal)", &cases.UInt32GTE{Val: 8}, true, "", 0},
	{"uint32 - gte - invalid", &cases.UInt32GTE{Val: 7}, false, "invalid UInt32GTE.Val: value must be greater than or equal to 8", 1},

	{"uint32 - gt & lt - valid", &cases.UInt32GTLT{Val: 7}, true, "", 0},
	{"uint32 - gt & lt - invalid (above)", &cases.UInt32GTLT{Val: 11}, false, "invalid UInt32GTLT.Val: value must be inside range (5, 10)", 1},
	{"uint32 - gt & lt - invalid (below)", &cases.UInt32GTLT{Val: 1}, false, "invalid UInt32GTLT.Val: value must be inside range (5, 10)", 1},
	{"uint32 - gt & lt - invalid (max)", &cases.UInt32GTLT{Val: 10}, false, "invalid UInt32GTLT.Val: value must be inside range (5, 10)", 1},
	{"uint32 - gt & lt - invalid (min)", &cases.UInt32GTLT{Val: 5}, false, "invalid UInt32GTLT.Val: value must be inside range (5, 10)", 1},

	{"uint32 - exclusive gt & lt - valid (above)", &cases.UInt32ExLTGT{Val: 11}, true, "", 0},
	{"uint32 - exclusive gt & lt - valid (below)", &cases.UInt32ExLTGT{Val: 4}, true, "", 0},
	{"uint32 - exclusive gt & lt - invalid", &cases.UInt32ExLTGT{Val: 7}, false, "invalid UInt32ExLTGT.Val: value must be outside range [5, 10]", 1},
	{"uint32 - exclusive gt & lt - invalid (max)", &cases.UInt32ExLTGT{Val: 10}, false, "invalid UInt32ExLTGT.Val: value must be outside range [5, 10]", 1},
	{"uint32 - exclusive gt & lt - invalid (min)", &cases.UInt32ExLTGT{Val: 5}, false, "invalid UInt32ExLTGT.Val: value must be outside range [5, 10]", 1},

	{"uint32 - gte & lte - valid", &cases.UInt32GTELTE{Val: 200}, true, "", 0},
	{"uint32 - gte & lte - valid (max)", &cases.UInt32GTELTE{Val: 256}, true, "", 0},
	{"uint32 - gte & lte - valid (min)", &cases.UInt32GTELTE{Val: 128}, true, "", 0},
	{"uint32 - gte & lte - invalid (above)", &cases.UInt32GTELTE{Val: 300}, false, "invalid UInt32GTELTE.Val: value must be inside range [128, 256]", 1},
	{"uint32 - gte & lte - invalid (below)", &cases.UInt32GTELTE{Val: 100}, false, "invalid UInt32GTELTE.Val: value must be inside range [128, 256]", 1},

	{"uint32 - exclusive gte & lte - valid (above)", &cases.UInt32ExGTELTE{Val: 300}, true, "", 0},
	{"uint32 - exclusive gte & lte - valid (below)", &cases.UInt32ExGTELTE{Val: 100}, true, "", 0},
	{"uint32 - exclusive gte & lte - valid (max)", &cases.UInt32ExGTELTE{Val: 256}, true, "", 0},
	{"uint32 - exclusive gte & lte - valid (min)", &cases.UInt32ExGTELTE{Val: 128}, true, "", 0},
	{"uint32 - exclusive gte & lte - invalid", &cases.UInt32ExGTELTE{Val: 200}, false, "invalid UInt32ExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var uint64Cases = []TestCase{
	{"uint64 - none - valid", &cases.UInt64None{Val: 123}, true, "", 0},

	{"uint64 - const - valid", &cases.UInt64Const{Val: 1}, true, "", 0},
	{"uint64 - const - invalid", &cases.UInt64Const{Val: 2}, false, "invalid UInt64Const.Val: value must equal 1", 1},

	{"uint64 - in - valid", &cases.UInt64In{Val: 3}, true, "", 0},
	{"uint64 - in - invalid", &cases.UInt64In{Val: 5}, false, "invalid UInt64In.Val: value must be in list [2 3]", 1},

	{"uint64 - not in - valid", &cases.UInt64NotIn{Val: 1}, true, "", 0},
	{"uint64 - not in - invalid", &cases.UInt64NotIn{Val: 0}, false, "invalid UInt64NotIn.Val: value must not be in list [0]", 1},

	{"uint64 - lt - valid", &cases.UInt64LT{Val: 4}, true, "", 0},
	{"uint64 - lt - invalid (equal)", &cases.UInt64LT{Val: 5}, false, "invalid UInt64LT.Val: value must be less than 5", 1},
	{"uint64 - lt - invalid", &cases.UInt64LT{Val: 6}, false, "invalid UInt64LT.Val: value must be less than 5", 1},

	{"uint64 - lte - valid", &cases.UInt64LTE{Val: 63}, true, "", 0},
	{"uint64 - lte - valid (equal)", &cases.UInt64LTE{Val: 64}, true, "", 0},
	{"uint64 - lte - invalid", &cases.UInt64LTE{Val: 65}, false, "invalid UInt64LTE.Val: value must be less than or equal to 64", 1},

	{"uint64 - gt - valid", &cases.UInt64GT{Val: 17}, true, "", 0},
	{"uint64 - gt - invalid (equal)", &cases.UInt64GT{Val: 16}, false, "invalid UInt64GT.Val: value must be greater than 16", 1},
	{"uint64 - gt - invalid", &cases.UInt64GT{Val: 15}, false, "invalid UInt64GT.Val: value must be greater than 16", 1},

	{"uint64 - gte - valid", &cases.UInt64GTE{Val: 9}, true, "", 0},
	{"uint64 - gte - valid (equal)", &cases.UInt64GTE{Val: 8}, true, "", 0},
	{"uint64 - gte - invalid", &cases.UInt64GTE{Val: 7}, false, "invalid UInt64GTE.Val: value must be greater than or equal to 8", 1},

	{"uint64 - gt & lt - valid", &cases.UInt64GTLT{Val: 7}, true, "", 0},
	{"uint64 - gt & lt - invalid (above)", &cases.UInt64GTLT{Val: 11}, false, "invalid UInt64GTLT.Val: value must be inside range (5, 10)", 1},
	{"uint64 - gt & lt - invalid (below)", &cases.UInt64GTLT{Val: 1}, false, "invalid UInt64GTLT.Val: value must be inside range (5, 10)", 1},
	{"uint64 - gt & lt - invalid (max)", &cases.UInt64GTLT{Val: 10}, false, "invalid UInt64GTLT.Val: value must be inside range (5, 10)", 1},
	{"uint64 - gt & lt - invalid (min)", &cases.UInt64GTLT{Val: 5}, false, "invalid UInt64GTLT.Val: value must be inside range (5, 10)", 1},

	{"uint64 - exclusive gt & lt - valid (above)", &cases.UInt64ExLTGT{Val: 11}, true, "", 0},
	{"uint64 - exclusive gt & lt - valid (below)", &cases.UInt64ExLTGT{Val: 4}, true, "", 0},
	{"uint64 - exclusive gt & lt - invalid", &cases.UInt64ExLTGT{Val: 7}, false, "invalid UInt64ExLTGT.Val: value must be outside range [5, 10]", 1},
	{"uint64 - exclusive gt & lt - invalid (max)", &cases.UInt64ExLTGT{Val: 10}, false, "invalid UInt64ExLTGT.Val: value must be outside range [5, 10]", 1},
	{"uint64 - exclusive gt & lt - invalid (min)", &cases.UInt64ExLTGT{Val: 5}, false, "invalid UInt64ExLTGT.Val: value must be outside range [5, 10]", 1},

	{"uint64 - gte & lte - valid", &cases.UInt64GTELTE{Val: 200}, true, "", 0},
	{"uint64 - gte & lte - valid (max)", &cases.UInt64GTELTE{Val: 256}, true, "", 0},
	{"uint64 - gte & lte - valid (min)", &cases.UInt64GTELTE{Val: 128}, true, "", 0},
	{"uint64 - gte & lte - invalid (above)", &cases.UInt64GTELTE{Val: 300}, false, "invalid UInt64GTELTE.Val: value must be inside range [128, 256]", 1},
	{"uint64 - gte & lte - invalid (below)", &cases.UInt64GTELTE{Val: 100}, false, "invalid UInt64GTELTE.Val: value must be inside range [128, 256]", 1},

	{"uint64 - exclusive gte & lte - valid (above)", &cases.UInt64ExGTELTE{Val: 300}, true, "", 0},
	{"uint64 - exclusive gte & lte - valid (below)", &cases.UInt64ExGTELTE{Val: 100}, true, "", 0},
	{"uint64 - exclusive gte & lte - valid (max)", &cases.UInt64ExGTELTE{Val: 256}, true, "", 0},
	{"uint64 - exclusive gte & lte - valid (min)", &cases.UInt64ExGTELTE{Val: 128}, true, "", 0},
	{"uint64 - exclusive gte & lte - invalid", &cases.UInt64ExGTELTE{Val: 200}, false, "invalid UInt64ExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var sint32Cases = []TestCase{
	{"sint32 - none - valid", &cases.SInt32None{Val: 123}, true, "", 0},

	{"sint32 - const - valid", &cases.SInt32Const{Val: 1}, true, "", 0},
	{"sint32 - const - invalid", &cases.SInt32Const{Val: 2}, false, "invalid SInt32Const.Val: value must equal 1", 1},

	{"sint32 - in - valid", &cases.SInt32In{Val: 3}, true, "", 0},
	{"sint32 - in - invalid", &cases.SInt32In{Val: 5}, false, "invalid SInt32In.Val: value must be in list [2 3]", 1},

	{"sint32 - not in - valid", &cases.SInt32NotIn{Val: 1}, true, "", 0},
	{"sint32 - not in - invalid", &cases.SInt32NotIn{Val: 0}, false, "invalid SInt32NotIn.Val: value must not be in list [0]", 1},

	{"sint32 - lt - valid", &cases.SInt32LT{Val: -1}, true, "", 0},
	{"sint32 - lt - invalid (equal)", &cases.SInt32LT{Val: 0}, false, "invalid SInt32LT.Val: value must be less than 0", 1},
	{"sint32 - lt - invalid", &cases.SInt32LT{Val: 1}, false, "invalid SInt32LT.Val: value must be less than 0", 1},

	{"sint32 - lte - valid", &cases.SInt32LTE{Val: 63}, true, "", 0},
	{"sint32 - lte - valid (equal)", &cases.SInt32LTE{Val: 64}, true, "", 0},
	{"sint32 - lte - invalid", &cases.SInt32LTE{Val: 65}, false, "invalid SInt32LTE.Val: value must be less than or equal to 64", 1},

	{"sint32 - gt - valid", &cases.SInt32GT{Val: 17}, true, "", 0},
	{"sint32 - gt - invalid (equal)", &cases.SInt32GT{Val: 16}, false, "invalid SInt32GT.Val: value must be greater than 16", 1},
	{"sint32 - gt - invalid", &cases.SInt32GT{Val: 15}, false, "invalid SInt32GT.Val: value must be greater than 16", 1},

	{"sint32 - gte - valid", &cases.SInt32GTE{Val: 9}, true, "", 0},
	{"sint32 - gte - valid (equal)", &cases.SInt32GTE{Val: 8}, true, "", 0},
	{"sint32 - gte - invalid", &cases.SInt32GTE{Val: 7}, false, "invalid SInt32GTE.Val: value must be greater than or equal to 8", 1},

	{"sint32 - gt & lt - valid", &cases.SInt32GTLT{Val: 5}, true, "", 0},
	{"sint32 - gt & lt - invalid (above)", &cases.SInt32GTLT{Val: 11}, false, "invalid SInt32GTLT.Val: value must be inside range (0, 10)", 1},
	{"sint32 - gt & lt - invalid (below)", &cases.SInt32GTLT{Val: -1}, false, "invalid SInt32GTLT.Val: value must be inside range (0, 10)", 1},
	{"sint32 - gt & lt - invalid (max)", &cases.SInt32GTLT{Val: 10}, false, "invalid SInt32GTLT.Val: value must be inside range (0, 10)", 1},
	{"sint32 - gt & lt - invalid (min)", &cases.SInt32GTLT{Val: 0}, false, "invalid SInt32GTLT.Val: value must be inside range (0, 10)", 1},

	{"sint32 - exclusive gt & lt - valid (above)", &cases.SInt32ExLTGT{Val: 11}, true, "", 0},
	{"sint32 - exclusive gt & lt - valid (below)", &cases.SInt32ExLTGT{Val: -1}, true, "", 0},
	{"sint32 - exclusive gt & lt - invalid", &cases.SInt32ExLTGT{Val: 5}, false, "invalid SInt32ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"sint32 - exclusive gt & lt - invalid (max)", &cases.SInt32ExLTGT{Val: 10}, false, "invalid SInt32ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"sint32 - exclusive gt & lt - invalid (min)", &cases.SInt32ExLTGT{Val: 0}, false, "invalid SInt32ExLTGT.Val: value must be outside range [0, 10]", 1},

	{"sint32 - gte & lte - valid", &cases.SInt32GTELTE{Val: 200}, true, "", 0},
	{"sint32 - gte & lte - valid (max)", &cases.SInt32GTELTE{Val: 256}, true, "", 0},
	{"sint32 - gte & lte - valid (min)", &cases.SInt32GTELTE{Val: 128}, true, "", 0},
	{"sint32 - gte & lte - invalid (above)", &cases.SInt32GTELTE{Val: 300}, false, "invalid SInt32GTELTE.Val: value must be inside range [128, 256]", 1},
	{"sint32 - gte & lte - invalid (below)", &cases.SInt32GTELTE{Val: 100}, false, "invalid SInt32GTELTE.Val: value must be inside range [128, 256]", 1},

	{"sint32 - exclusive gte & lte - valid (above)", &cases.SInt32ExGTELTE{Val: 300}, true, "", 0},
	{"sint32 - exclusive gte & lte - valid (below)", &cases.SInt32ExGTELTE{Val: 100}, true, "", 0},
	{"sint32 - exclusive gte & lte - valid (max)", &cases.SInt32ExGTELTE{Val: 256}, true, "", 0},
	{"sint32 - exclusive gte & lte - valid (min)", &cases.SInt32ExGTELTE{Val: 128}, true, "", 0},
	{"sint32 - exclusive gte & lte - invalid", &cases.SInt32ExGTELTE{Val: 200}, false, "invalid SInt32ExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var sint64Cases = []TestCase{
	{"sint64 - none - valid", &cases.SInt64None{Val: 123}, true, "", 0},

	{"sint64 - const - valid", &cases.SInt64Const{Val: 1}, true, "", 0},
	{"sint64 - const - invalid", &cases.SInt64Const{Val: 2}, false, "invalid SInt64Const.Val: value must equal 1", 1},

	{"sint64 - in - valid", &cases.SInt64In{Val: 3}, true, "", 0},
	{"sint64 - in - invalid", &cases.SInt64In{Val: 5}, false, "invalid SInt64In.Val: value must be in list [2 3]", 1},

	{"sint64 - not in - valid", &cases.SInt64NotIn{Val: 1}, true, "", 0},
	{"sint64 - not in - invalid", &cases.SInt64NotIn{Val: 0}, false, "invalid SInt64NotIn.Val: value must not be in list [0]", 1},

	{"sint64 - lt - valid", &cases.SInt64LT{Val: -1}, true, "", 0},
	{"sint64 - lt - invalid (equal)", &cases.SInt64LT{Val: 0}, false, "invalid SInt64LT.Val: value must be less than 0", 1},
	{"sint64 - lt - invalid", &cases.SInt64LT{Val: 1}, false, "invalid SInt64LT.Val: value must be less than 0", 1},

	{"sint64 - lte - valid", &cases.SInt64LTE{Val: 63}, true, "", 0},
	{"sint64 - lte - valid (equal)", &cases.SInt64LTE{Val: 64}, true, "", 0},
	{"sint64 - lte - invalid", &cases.SInt64LTE{Val: 65}, false, "invalid SInt64LTE.Val: value must be less than or equal to 64", 1},

	{"sint64 - gt - valid", &cases.SInt64GT{Val: 17}, true, "", 0},
	{"sint64 - gt - invalid (equal)", &cases.SInt64GT{Val: 16}, false, "invalid SInt64GT.Val: value must be greater than 16", 1},
	{"sint64 - gt - invalid", &cases.SInt64GT{Val: 15}, false, "invalid SInt64GT.Val: value must be greater than 16", 1},

	{"sint64 - gte - valid", &cases.SInt64GTE{Val: 9}, true, "", 0},
	{"sint64 - gte - valid (equal)", &cases.SInt64GTE{Val: 8}, true, "", 0},
	{"sint64 - gte - invalid", &cases.SInt64GTE{Val: 7}, false, "invalid SInt64GTE.Val: value must be greater than or equal to 8", 1},

	{"sint64 - gt & lt - valid", &cases.SInt64GTLT{Val: 5}, true, "", 0},
	{"sint64 - gt & lt - invalid (above)", &cases.SInt64GTLT{Val: 11}, false, "invalid SInt64GTLT.Val: value must be inside range (0, 10)", 1},
	{"sint64 - gt & lt - invalid (below)", &cases.SInt64GTLT{Val: -1}, false, "invalid SInt64GTLT.Val: value must be inside range (0, 10)", 1},
	{"sint64 - gt & lt - invalid (max)", &cases.SInt64GTLT{Val: 10}, false, "invalid SInt64GTLT.Val: value must be inside range (0, 10)", 1},
	{"sint64 - gt & lt - invalid (min)", &cases.SInt64GTLT{Val: 0}, false, "invalid SInt64GTLT.Val: value must be inside range (0, 10)", 1},

	{"sint64 - exclusive gt & lt - valid (above)", &cases.SInt64ExLTGT{Val: 11}, true, "", 0},
	{"sint64 - exclusive gt & lt - valid (below)", &cases.SInt64ExLTGT{Val: -1}, true, "", 0},
	{"sint64 - exclusive gt & lt - invalid", &cases.SInt64ExLTGT{Val: 5}, false, "invalid SInt64ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"sint64 - exclusive gt & lt - invalid (max)", &cases.SInt64ExLTGT{Val: 10}, false, "invalid SInt64ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"sint64 - exclusive gt & lt - invalid (min)", &cases.SInt64ExLTGT{Val: 0}, false, "invalid SInt64ExLTGT.Val: value must be outside range [0, 10]", 1},

	{"sint64 - gte & lte - valid", &cases.SInt64GTELTE{Val: 200}, true, "", 0},
	{"sint64 - gte & lte - valid (max)", &cases.SInt64GTELTE{Val: 256}, true, "", 0},
	{"sint64 - gte & lte - valid (min)", &cases.SInt64GTELTE{Val: 128}, true, "", 0},
	{"sint64 - gte & lte - invalid (above)", &cases.SInt64GTELTE{Val: 300}, false, "invalid SInt64GTELTE.Val: value must be inside range [128, 256]", 1},
	{"sint64 - gte & lte - invalid (below)", &cases.SInt64GTELTE{Val: 100}, false, "invalid SInt64GTELTE.Val: value must be inside range [128, 256]", 1},

	{"sint64 - exclusive gte & lte - valid (above)", &cases.SInt64ExGTELTE{Val: 300}, true, "", 0},
	{"sint64 - exclusive gte & lte - valid (below)", &cases.SInt64ExGTELTE{Val: 100}, true, "", 0},
	{"sint64 - exclusive gte & lte - valid (max)", &cases.SInt64ExGTELTE{Val: 256}, true, "", 0},
	{"sint64 - exclusive gte & lte - valid (min)", &cases.SInt64ExGTELTE{Val: 128}, true, "", 0},
	{"sint64 - exclusive gte & lte - invalid", &cases.SInt64ExGTELTE{Val: 200}, false, "invalid SInt64ExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var fixed32Cases = []TestCase{
	{"fixed32 - none - valid", &cases.Fixed32None{Val: 123}, true, "", 0},

	{"fixed32 - const - valid", &cases.Fixed32Const{Val: 1}, true, "", 0},
	{"fixed32 - const - invalid", &cases.Fixed32Const{Val: 2}, false, "invalid Fixed32Const.Val: value must equal 1", 1},

	{"fixed32 - in - valid", &cases.Fixed32In{Val: 3}, true, "", 0},
	{"fixed32 - in - invalid", &cases.Fixed32In{Val: 5}, false, "invalid Fixed32In.Val: value must be in list [2 3]", 1},

	{"fixed32 - not in - valid", &cases.Fixed32NotIn{Val: 1}, true, "", 0},
	{"fixed32 - not in - invalid", &cases.Fixed32NotIn{Val: 0}, false, "invalid Fixed32NotIn.Val: value must not be in list [0]", 1},

	{"fixed32 - lt - valid", &cases.Fixed32LT{Val: 4}, true, "", 0},
	{"fixed32 - lt - invalid (equal)", &cases.Fixed32LT{Val: 5}, false, "invalid Fixed32LT.Val: value must be less than 5", 1},
	{"fixed32 - lt - invalid", &cases.Fixed32LT{Val: 6}, false, "invalid Fixed32LT.Val: value must be less than 5", 1},

	{"fixed32 - lte - valid", &cases.Fixed32LTE{Val: 63}, true, "", 0},
	{"fixed32 - lte - valid (equal)", &cases.Fixed32LTE{Val: 64}, true, "", 0},
	{"fixed32 - lte - invalid", &cases.Fixed32LTE{Val: 65}, false, "invalid Fixed32LTE.Val: value must be less than or equal to 64", 1},

	{"fixed32 - gt - valid", &cases.Fixed32GT{Val: 17}, true, "", 0},
	{"fixed32 - gt - invalid (equal)", &cases.Fixed32GT{Val: 16}, false, "invalid Fixed32GT.Val: value must be greater than 16", 1},
	{"fixed32 - gt - invalid", &cases.Fixed32GT{Val: 15}, false, "invalid Fixed32GT.Val: value must be greater than 16", 1},

	{"fixed32 - gte - valid", &cases.Fixed32GTE{Val: 9}, true, "", 0},
	{"fixed32 - gte - valid (equal)", &cases.Fixed32GTE{Val: 8}, true, "", 0},
	{"fixed32 - gte - invalid", &cases.Fixed32GTE{Val: 7}, false, "invalid Fixed32GTE.Val: value must be greater than or equal to 8", 1},

	{"fixed32 - gt & lt - valid", &cases.Fixed32GTLT{Val: 7}, true, "", 0},
	{"fixed32 - gt & lt - invalid (above)", &cases.Fixed32GTLT{Val: 11}, false, "invalid Fixed32GTLT.Val: value must be inside range (5, 10)", 1},
	{"fixed32 - gt & lt - invalid (below)", &cases.Fixed32GTLT{Val: 1}, false, "invalid Fixed32GTLT.Val: value must be inside range (5, 10)", 1},
	{"fixed32 - gt & lt - invalid (max)", &cases.Fixed32GTLT{Val: 10}, false, "invalid Fixed32GTLT.Val: value must be inside range (5, 10)", 1},
	{"fixed32 - gt & lt - invalid (min)", &cases.Fixed32GTLT{Val: 5}, false, "invalid Fixed32GTLT.Val: value must be inside range (5, 10)", 1},

	{"fixed32 - exclusive gt & lt - valid (above)", &cases.Fixed32ExLTGT{Val: 11}, true, "", 0},
	{"fixed32 - exclusive gt & lt - valid (below)", &cases.Fixed32ExLTGT{Val: 4}, true, "", 0},
	{"fixed32 - exclusive gt & lt - invalid", &cases.Fixed32ExLTGT{Val: 7}, false, "invalid Fixed32ExLTGT.Val: value must be outside range [5, 10]", 1},
	{"fixed32 - exclusive gt & lt - invalid (max)", &cases.Fixed32ExLTGT{Val: 10}, false, "invalid Fixed32ExLTGT.Val: value must be outside range [5, 10]", 1},
	{"fixed32 - exclusive gt & lt - invalid (min)", &cases.Fixed32ExLTGT{Val: 5}, false, "invalid Fixed32ExLTGT.Val: value must be outside range [5, 10]", 1},

	{"fixed32 - gte & lte - valid", &cases.Fixed32GTELTE{Val: 200}, true, "", 0},
	{"fixed32 - gte & lte - valid (max)", &cases.Fixed32GTELTE{Val: 256}, true, "", 0},
	{"fixed32 - gte & lte - valid (min)", &cases.Fixed32GTELTE{Val: 128}, true, "", 0},
	{"fixed32 - gte & lte - invalid (above)", &cases.Fixed32GTELTE{Val: 300}, false, "invalid Fixed32GTELTE.Val: value must be inside range [128, 256]", 1},
	{"fixed32 - gte & lte - invalid (below)", &cases.Fixed32GTELTE{Val: 100}, false, "invalid Fixed32GTELTE.Val: value must be inside range [128, 256]", 1},

	{"fixed32 - exclusive gte & lte - valid (above)", &cases.Fixed32ExGTELTE{Val: 300}, true, "", 0},
	{"fixed32 - exclusive gte & lte - valid (below)", &cases.Fixed32ExGTELTE{Val: 100}, true, "", 0},
	{"fixed32 - exclusive gte & lte - valid (max)", &cases.Fixed32ExGTELTE{Val: 256}, true, "", 0},
	{"fixed32 - exclusive gte & lte - valid (min)", &cases.Fixed32ExGTELTE{Val: 128}, true, "", 0},
	{"fixed32 - exclusive gte & lte - invalid", &cases.Fixed32ExGTELTE{Val: 200}, false, "invalid Fixed32ExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var fixed64Cases = []TestCase{
	{"fixed64 - none - valid", &cases.Fixed64None{Val: 123}, true, "", 0},

	{"fixed64 - const - valid", &cases.Fixed64Const{Val: 1}, true, "", 0},
	{"fixed64 - const - invalid", &cases.Fixed64Const{Val: 2}, false, "invalid Fixed64Const.Val: value must equal 1", 1},

	{"fixed64 - in - valid", &cases.Fixed64In{Val: 3}, true, "", 0},
	{"fixed64 - in - invalid", &cases.Fixed64In{Val: 5}, false, "invalid Fixed64In.Val: value must be in list [2 3]", 1},

	{"fixed64 - not in - valid", &cases.Fixed64NotIn{Val: 1}, true, "", 0},
	{"fixed64 - not in - invalid", &cases.Fixed64NotIn{Val: 0}, false, "invalid Fixed64NotIn.Val: value must not be in list [0]", 1},

	{"fixed64 - lt - valid", &cases.Fixed64LT{Val: 4}, true, "", 0},
	{"fixed64 - lt - invalid (equal)", &cases.Fixed64LT{Val: 5}, false, "invalid Fixed64LT.Val: value must be less than 5", 1},
	{"fixed64 - lt - invalid", &cases.Fixed64LT{Val: 6}, false, "invalid Fixed64LT.Val: value must be less than 5", 1},

	{"fixed64 - lte - valid", &cases.Fixed64LTE{Val: 63}, true, "", 0},
	{"fixed64 - lte - valid (equal)", &cases.Fixed64LTE{Val: 64}, true, "", 0},
	{"fixed64 - lte - invalid", &cases.Fixed64LTE{Val: 65}, false, "invalid Fixed64LTE.Val: value must be less than or equal to 64", 1},

	{"fixed64 - gt - valid", &cases.Fixed64GT{Val: 17}, true, "", 0},
	{"fixed64 - gt - invalid (equal)", &cases.Fixed64GT{Val: 16}, false, "invalid Fixed64GT.Val: value must be greater than 16", 1},
	{"fixed64 - gt - invalid", &cases.Fixed64GT{Val: 15}, false, "invalid Fixed64GT.Val: value must be greater than 16", 1},

	{"fixed64 - gte - valid", &cases.Fixed64GTE{Val: 9}, true, "", 0},
	{"fixed64 - gte - valid (equal)", &cases.Fixed64GTE{Val: 8}, true, "", 0},
	{"fixed64 - gte - invalid", &cases.Fixed64GTE{Val: 7}, false, "invalid Fixed64GTE.Val: value must be greater than or equal to 8", 1},

	{"fixed64 - gt & lt - valid", &cases.Fixed64GTLT{Val: 7}, true, "", 0},
	{"fixed64 - gt & lt - invalid (above)", &cases.Fixed64GTLT{Val: 11}, false, "invalid Fixed64GTLT.Val: value must be inside range (5, 10)", 1},
	{"fixed64 - gt & lt - invalid (below)", &cases.Fixed64GTLT{Val: 1}, false, "invalid Fixed64GTLT.Val: value must be inside range (5, 10)", 1},
	{"fixed64 - gt & lt - invalid (max)", &cases.Fixed64GTLT{Val: 10}, false, "invalid Fixed64GTLT.Val: value must be inside range (5, 10)", 1},
	{"fixed64 - gt & lt - invalid (min)", &cases.Fixed64GTLT{Val: 5}, false, "invalid Fixed64GTLT.Val: value must be inside range (5, 10)", 1},

	{"fixed64 - exclusive gt & lt - valid (above)", &cases.Fixed64ExLTGT{Val: 11}, true, "", 0},
	{"fixed64 - exclusive gt & lt - valid (below)", &cases.Fixed64ExLTGT{Val: 4}, true, "", 0},
	{"fixed64 - exclusive gt & lt - invalid", &cases.Fixed64ExLTGT{Val: 7}, false, "invalid Fixed64ExLTGT.Val: value must be outside range [5, 10]", 1},
	{"fixed64 - exclusive gt & lt - invalid (max)", &cases.Fixed64ExLTGT{Val: 10}, false, "invalid Fixed64ExLTGT.Val: value must be outside range [5, 10]", 1},
	{"fixed64 - exclusive gt & lt - invalid (min)", &cases.Fixed64ExLTGT{Val: 5}, false, "invalid Fixed64ExLTGT.Val: value must be outside range [5, 10]", 1},

	{"fixed64 - gte & lte - valid", &cases.Fixed64GTELTE{Val: 200}, true, "", 0},
	{"fixed64 - gte & lte - valid (max)", &cases.Fixed64GTELTE{Val: 256}, true, "", 0},
	{"fixed64 - gte & lte - valid (min)", &cases.Fixed64GTELTE{Val: 128}, true, "", 0},
	{"fixed64 - gte & lte - invalid (above)", &cases.Fixed64GTELTE{Val: 300}, false, "invalid Fixed64GTELTE.Val: value must be inside range [128, 256]", 1},
	{"fixed64 - gte & lte - invalid (below)", &cases.Fixed64GTELTE{Val: 100}, false, "invalid Fixed64GTELTE.Val: value must be inside range [128, 256]", 1},

	{"fixed64 - exclusive gte & lte - valid (above)", &cases.Fixed64ExGTELTE{Val: 300}, true, "", 0},
	{"fixed64 - exclusive gte & lte - valid (below)", &cases.Fixed64ExGTELTE{Val: 100}, true, "", 0},
	{"fixed64 - exclusive gte & lte - valid (max)", &cases.Fixed64ExGTELTE{Val: 256}, true, "", 0},
	{"fixed64 - exclusive gte & lte - valid (min)", &cases.Fixed64ExGTELTE{Val: 128}, true, "", 0},
	{"fixed64 - exclusive gte & lte - invalid", &cases.Fixed64ExGTELTE{Val: 200}, false, "invalid Fixed64ExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var sfixed32Cases = []TestCase{
	{"sfixed32 - none - valid", &cases.SFixed32None{Val: 123}, true, "", 0},

	{"sfixed32 - const - valid", &cases.SFixed32Const{Val: 1}, true, "", 0},
	{"sfixed32 - const - invalid", &cases.SFixed32Const{Val: 2}, false, "invalid SFixed32Const.Val: value must equal 1", 1},

	{"sfixed32 - in - valid", &cases.SFixed32In{Val: 3}, true, "", 0},
	{"sfixed32 - in - invalid", &cases.SFixed32In{Val: 5}, false, "invalid SFixed32In.Val: value must be in list [2 3]", 1},

	{"sfixed32 - not in - valid", &cases.SFixed32NotIn{Val: 1}, true, "", 0},
	{"sfixed32 - not in - invalid", &cases.SFixed32NotIn{Val: 0}, false, "invalid SFixed32NotIn.Val: value must not be in list [0]", 1},

	{"sfixed32 - lt - valid", &cases.SFixed32LT{Val: -1}, true, "", 0},
	{"sfixed32 - lt - invalid (equal)", &cases.SFixed32LT{Val: 0}, false, "invalid SFixed32LT.Val: value must be less than 0", 1},
	{"sfixed32 - lt - invalid", &cases.SFixed32LT{Val: 1}, false, "invalid SFixed32LT.Val: value must be less than 0", 1},

	{"sfixed32 - lte - valid", &cases.SFixed32LTE{Val: 63}, true, "", 0},
	{"sfixed32 - lte - valid (equal)", &cases.SFixed32LTE{Val: 64}, true, "", 0},
	{"sfixed32 - lte - invalid", &cases.SFixed32LTE{Val: 65}, false, "invalid SFixed32LTE.Val: value must be less than or equal to 64", 1},

	{"sfixed32 - gt - valid", &cases.SFixed32GT{Val: 17}, true, "", 0},
	{"sfixed32 - gt - invalid (equal)", &cases.SFixed32GT{Val: 16}, false, "invalid SFixed32GT.Val: value must be greater than 16", 1},
	{"sfixed32 - gt - invalid", &cases.SFixed32GT{Val: 15}, false, "invalid SFixed32GT.Val: value must be greater than 16", 1},

	{"sfixed32 - gte - valid", &cases.SFixed32GTE{Val: 9}, true, "", 0},
	{"sfixed32 - gte - valid (equal)", &cases.SFixed32GTE{Val: 8}, true, "", 0},
	{"sfixed32 - gte - invalid", &cases.SFixed32GTE{Val: 7}, false, "invalid SFixed32GTE.Val: value must be greater than or equal to 8", 1},

	{"sfixed32 - gt & lt - valid", &cases.SFixed32GTLT{Val: 5}, true, "", 0},
	{"sfixed32 - gt & lt - invalid (above)", &cases.SFixed32GTLT{Val: 11}, false, "invalid SFixed32GTLT.Val: value must be inside range (0, 10)", 1},
	{"sfixed32 - gt & lt - invalid (below)", &cases.SFixed32GTLT{Val: -1}, false, "invalid SFixed32GTLT.Val: value must be inside range (0, 10)", 1},
	{"sfixed32 - gt & lt - invalid (max)", &cases.SFixed32GTLT{Val: 10}, false, "invalid SFixed32GTLT.Val: value must be inside range (0, 10)", 1},
	{"sfixed32 - gt & lt - invalid (min)", &cases.SFixed32GTLT{Val: 0}, false, "invalid SFixed32GTLT.Val: value must be inside range (0, 10)", 1},

	{"sfixed32 - exclusive gt & lt - valid (above)", &cases.SFixed32ExLTGT{Val: 11}, true, "", 0},
	{"sfixed32 - exclusive gt & lt - valid (below)", &cases.SFixed32ExLTGT{Val: -1}, true, "", 0},
	{"sfixed32 - exclusive gt & lt - invalid", &cases.SFixed32ExLTGT{Val: 5}, false, "invalid SFixed32ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"sfixed32 - exclusive gt & lt - invalid (max)", &cases.SFixed32ExLTGT{Val: 10}, false, "invalid SFixed32ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"sfixed32 - exclusive gt & lt - invalid (min)", &cases.SFixed32ExLTGT{Val: 0}, false, "invalid SFixed32ExLTGT.Val: value must be outside range [0, 10]", 1},

	{"sfixed32 - gte & lte - valid", &cases.SFixed32GTELTE{Val: 200}, true, "", 0},
	{"sfixed32 - gte & lte - valid (max)", &cases.SFixed32GTELTE{Val: 256}, true, "", 0},
	{"sfixed32 - gte & lte - valid (min)", &cases.SFixed32GTELTE{Val: 128}, true, "", 0},
	{"sfixed32 - gte & lte - invalid (above)", &cases.SFixed32GTELTE{Val: 300}, false, "invalid SFixed32GTELTE.Val: value must be inside range [128, 256]", 1},
	{"sfixed32 - gte & lte - invalid (below)", &cases.SFixed32GTELTE{Val: 100}, false, "invalid SFixed32GTELTE.Val: value must be inside range [128, 256]", 1},

	{"sfixed32 - exclusive gte & lte - valid (above)", &cases.SFixed32ExGTELTE{Val: 300}, true, "", 0},
	{"sfixed32 - exclusive gte & lte - valid (below)", &cases.SFixed32ExGTELTE{Val: 100}, true, "", 0},
	{"sfixed32 - exclusive gte & lte - valid (max)", &cases.SFixed32ExGTELTE{Val: 256}, true, "", 0},
	{"sfixed32 - exclusive gte & lte - valid (min)", &cases.SFixed32ExGTELTE{Val: 128}, true, "", 0},
	{"sfixed32 - exclusive gte & lte - invalid", &cases.SFixed32ExGTELTE{Val: 200}, false, "invalid SFixed32ExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var sfixed64Cases = []TestCase{
	{"sfixed64 - none - valid", &cases.SFixed64None{Val: 123}, true, "", 0},

	{"sfixed64 - const - valid", &cases.SFixed64Const{Val: 1}, true, "", 0},
	{"sfixed64 - const - invalid", &cases.SFixed64Const{Val: 2}, false, "invalid SFixed64Const.Val: value must equal 1", 1},

	{"sfixed64 - in - valid", &cases.SFixed64In{Val: 3}, true, "", 0},
	{"sfixed64 - in - invalid", &cases.SFixed64In{Val: 5}, false, "invalid SFixed64In.Val: value must be in list [2 3]", 1},

	{"sfixed64 - not in - valid", &cases.SFixed64NotIn{Val: 1}, true, "", 0},
	{"sfixed64 - not in - invalid", &cases.SFixed64NotIn{Val: 0}, false, "invalid SFixed64NotIn.Val: value must not be in list [0]", 1},

	{"sfixed64 - lt - valid", &cases.SFixed64LT{Val: -1}, true, "", 0},
	{"sfixed64 - lt - invalid (equal)", &cases.SFixed64LT{Val: 0}, false, "invalid SFixed64LT.Val: value must be less than 0", 1},
	{"sfixed64 - lt - invalid", &cases.SFixed64LT{Val: 1}, false, "invalid SFixed64LT.Val: value must be less than 0", 1},

	{"sfixed64 - lte - valid", &cases.SFixed64LTE{Val: 63}, true, "", 0},
	{"sfixed64 - lte - valid (equal)", &cases.SFixed64LTE{Val: 64}, true, "", 0},
	{"sfixed64 - lte - invalid", &cases.SFixed64LTE{Val: 65}, false, "invalid SFixed64LTE.Val: value must be less than or equal to 64", 1},

	{"sfixed64 - gt - valid", &cases.SFixed64GT{Val: 17}, true, "", 0},
	{"sfixed64 - gt - invalid (equal)", &cases.SFixed64GT{Val: 16}, false, "invalid SFixed64GT.Val: value must be greater than 16", 1},
	{"sfixed64 - gt - invalid", &cases.SFixed64GT{Val: 15}, false, "invalid SFixed64GT.Val: value must be greater than 16", 1},

	{"sfixed64 - gte - valid", &cases.SFixed64GTE{Val: 9}, true, "", 0},
	{"sfixed64 - gte - valid (equal)", &cases.SFixed64GTE{Val: 8}, true, "", 0},
	{"sfixed64 - gte - invalid", &cases.SFixed64GTE{Val: 7}, false, "invalid SFixed64GTE.Val: value must be greater than or equal to 8", 1},

	{"sfixed64 - gt & lt - valid", &cases.SFixed64GTLT{Val: 5}, true, "", 0},
	{"sfixed64 - gt & lt - invalid (above)", &cases.SFixed64GTLT{Val: 11}, false, "invalid SFixed64GTLT.Val: value must be inside range (0, 10)", 1},
	{"sfixed64 - gt & lt - invalid (below)", &cases.SFixed64GTLT{Val: -1}, false, "invalid SFixed64GTLT.Val: value must be inside range (0, 10)", 1},
	{"sfixed64 - gt & lt - invalid (max)", &cases.SFixed64GTLT{Val: 10}, false, "invalid SFixed64GTLT.Val: value must be inside range (0, 10)", 1},
	{"sfixed64 - gt & lt - invalid (min)", &cases.SFixed64GTLT{Val: 0}, false, "invalid SFixed64GTLT.Val: value must be inside range (0, 10)", 1},

	{"sfixed64 - exclusive gt & lt - valid (above)", &cases.SFixed64ExLTGT{Val: 11}, true, "", 0},
	{"sfixed64 - exclusive gt & lt - valid (below)", &cases.SFixed64ExLTGT{Val: -1}, true, "", 0},
	{"sfixed64 - exclusive gt & lt - invalid", &cases.SFixed64ExLTGT{Val: 5}, false, "invalid SFixed64ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"sfixed64 - exclusive gt & lt - invalid (max)", &cases.SFixed64ExLTGT{Val: 10}, false, "invalid SFixed64ExLTGT.Val: value must be outside range [0, 10]", 1},
	{"sfixed64 - exclusive gt & lt - invalid (min)", &cases.SFixed64ExLTGT{Val: 0}, false, "invalid SFixed64ExLTGT.Val: value must be outside range [0, 10]", 1},

	{"sfixed64 - gte & lte - valid", &cases.SFixed64GTELTE{Val: 200}, true, "", 0},
	{"sfixed64 - gte & lte - valid (max)", &cases.SFixed64GTELTE{Val: 256}, true, "", 0},
	{"sfixed64 - gte & lte - valid (min)", &cases.SFixed64GTELTE{Val: 128}, true, "", 0},
	{"sfixed64 - gte & lte - invalid (above)", &cases.SFixed64GTELTE{Val: 300}, false, "invalid SFixed64GTELTE.Val: value must be inside range [128, 256]", 1},
	{"sfixed64 - gte & lte - invalid (below)", &cases.SFixed64GTELTE{Val: 100}, false, "invalid SFixed64GTELTE.Val: value must be inside range [128, 256]", 1},

	{"sfixed64 - exclusive gte & lte - valid (above)", &cases.SFixed64ExGTELTE{Val: 300}, true, "", 0},
	{"sfixed64 - exclusive gte & lte - valid (below)", &cases.SFixed64ExGTELTE{Val: 100}, true, "", 0},
	{"sfixed64 - exclusive gte & lte - valid (max)", &cases.SFixed64ExGTELTE{Val: 256}, true, "", 0},
	{"sfixed64 - exclusive gte & lte - valid (min)", &cases.SFixed64ExGTELTE{Val: 128}, true, "", 0},
	{"sfixed64 - exclusive gte & lte - invalid", &cases.SFixed64ExGTELTE{Val: 200}, false, "invalid SFixed64ExGTELTE.Val: value must be outside range (128, 256)", 1},
}

var boolCases = []TestCase{
	{"bool - none - valid", &cases.BoolNone{Val: true}, true, "", 0},
	{"bool - const (true) - valid", &cases.BoolConstTrue{Val: true}, true, "", 0},
	{"bool - const (true) - invalid", &cases.BoolConstTrue{Val: false}, false, "invalid BoolConstTrue.Val: value must equal true", 1},
	{"bool - const (false) - valid", &cases.BoolConstFalse{Val: false}, true, "", 0},
	{"bool - const (false) - invalid", &cases.BoolConstFalse{Val: true}, false, "invalid BoolConstFalse.Val: value must equal false", 1},
}

var stringCases = []TestCase{
	{"string - none - valid", &cases.StringNone{Val: "quux"}, true, "", 0},

	{"string - const - valid", &cases.StringConst{Val: "foo"}, true, "", 0},
	{"string - const - invalid", &cases.StringConst{Val: "bar"}, false, "invalid StringConst.Val: value must equal foo", 1},

	{"string - in - valid", &cases.StringIn{Val: "bar"}, true, "", 0},
	{"string - in - invalid", &cases.StringIn{Val: "quux"}, false, "invalid StringIn.Val: value must be in list [bar baz]", 1},
	{"string - not in - valid", &cases.StringNotIn{Val: "quux"}, true, "", 0},
	{"string - not in - invalid", &cases.StringNotIn{Val: "fizz"}, false, "invalid StringNotIn.Val: value must not be in list [fizz buzz]", 1},

	{"string - len - valid", &cases.StringLen{Val: "baz"}, true, "", 0},
	{"string - len - valid (multibyte)", &cases.StringLen{Val: "你好吖"}, true, "", 0},
	{"string - len - invalid (lt)", &cases.StringLen{Val: "go"}, false, "invalid StringLen.Val: value length must be 3 runes", 1},
	{"string - len - invalid (gt)", &cases.StringLen{Val: "fizz"}, false, "invalid StringLen.Val: value length must be 3 runes", 1},
	{"string - len - invalid (multibyte)", &cases.StringLen{Val: "你好"}, false, "invalid StringLen.Val: value length must be 3 runes", 1},

	{"string - min len - valid", &cases.StringMinLen{Val: "protoc"}, true, "", 0},
	{"string - min len - valid (min)", &cases.StringMinLen{Val: "baz"}, true, "", 0},
	{"string - min len - invalid", &cases.StringMinLen{Val: "go"}, false, "invalid StringMinLen.Val: value length must be at least 3 runes", 1},
	{"string - min len - invalid (multibyte)", &cases.StringMinLen{Val: "你好"}, false, "invalid StringMinLen.Val: value length must be at least 3 runes", 1},

	{"string - max len - valid", &cases.StringMaxLen{Val: "foo"}, true, "", 0},
	{"string - max len - valid (max)", &cases.StringMaxLen{Val: "proto"}, true, "", 0},
	{"string - max len - valid (multibyte)", &cases.StringMaxLen{Val: "你好你好"}, true, "", 0},
	{"string - max len - invalid", &cases.StringMaxLen{Val: "1234567890"}, false, "invalid StringMaxLen.Val: value length must be at most 5 runes", 1},

	{"string - min/max len - valid", &cases.StringMinMaxLen{Val: "quux"}, true, "", 0},
	{"string - min/max len - valid (min)", &cases.StringMinMaxLen{Val: "foo"}, true, "", 0},
	{"string - min/max len - valid (max)", &cases.StringMinMaxLen{Val: "proto"}, true, "", 0},
	{"string - min/max len - valid (multibyte)", &cases.StringMinMaxLen{Val: "你好你好"}, true, "", 0},
	{"string - min/max len - invalid (below)", &cases.StringMinMaxLen{Val: "go"}, false, "invalid StringMinMaxLen.Val: value length must be between 3 and 5 runes, inclusive", 1},
	{"string - min/max len - invalid (above)", &cases.StringMinMaxLen{Val: "validate"}, false, "invalid StringMinMaxLen.Val: value length must be between 3 and 5 runes, inclusive", 1},

	{"string - equal min/max len - valid", &cases.StringEqualMinMaxLen{Val: "proto"}, true, "", 0},
	{"string - equal min/max len - invalid", &cases.StringEqualMinMaxLen{Val: "validate"}, false, "invalid StringEqualMinMaxLen.Val: value length must be 5 runes", 1},

	{"string - len bytes - valid", &cases.StringLenBytes{Val: "pace"}, true, "", 0},
	{"string - len bytes - invalid (lt)", &cases.StringLenBytes{Val: "val"}, false, "invalid StringLenBytes.Val: value length must be 4 bytes", 1},
	{"string - len bytes - invalid (gt)", &cases.StringLenBytes{Val: "world"}, false, "invalid StringLenBytes.Val: value length must be 4 bytes", 1},
	{"string - len bytes - invalid (multibyte)", &cases.StringLenBytes{Val: "世界和平"}, false, "invalid StringLenBytes.Val: value length must be 4 bytes", 1},

	{"string - min bytes - valid", &cases.StringMinBytes{Val: "proto"}, true, "", 0},
	{"string - min bytes - valid (min)", &cases.StringMinBytes{Val: "quux"}, true, "", 0},
	{"string - min bytes - valid (multibyte)", &cases.StringMinBytes{Val: "你好"}, true, "", 0},
	{"string - min bytes - invalid", &cases.StringMinBytes{Val: ""}, false, "invalid StringMinBytes.Val: value length must be at least 4 bytes", 1},

	{"string - max bytes - valid", &cases.StringMaxBytes{Val: "foo"}, true, "", 0},
	{"string - max bytes - valid (max)", &cases.StringMaxBytes{Val: "12345678"}, true, "", 0},
	{"string - max bytes - invalid", &cases.StringMaxBytes{Val: "123456789"}, false, "invalid StringMaxBytes.Val: value length must be at most 8 bytes", 1},
	{"string - max bytes - invalid (multibyte)", &cases.StringMaxBytes{Val: "你好你好你好"}, false, "invalid StringMaxBytes.Val: value length must be at most 8 bytes", 1},

	{"string - min/max bytes - valid", &cases.StringMinMaxBytes{Val: "protoc"}, true, "", 0},
	{"string - min/max bytes - valid (min)", &cases.StringMinMaxBytes{Val: "quux"}, true, "", 0},
	{"string - min/max bytes - valid (max)", &cases.StringMinMaxBytes{Val: "fizzbuzz"}, true, "", 0},
	{"string - min/max bytes - valid (multibyte)", &cases.StringMinMaxBytes{Val: "你好"}, true, "", 0},
	{"string - min/max bytes - invalid (below)", &cases.StringMinMaxBytes{Val: "foo"}, false, "invalid StringMinMaxBytes.Val: value length must be between 4 and 8 bytes, inclusive", 1},
	{"string - min/max bytes - invalid (above)", &cases.StringMinMaxBytes{Val: "你好你好你"}, false, "invalid StringMinMaxBytes.Val: value length must be between 4 and 8 bytes, inclusive", 1},

	{"string - equal min/max bytes - valid", &cases.StringEqualMinMaxBytes{Val: "protoc"}, true, "", 0},
	{"string - equal min/max bytes - invalid", &cases.StringEqualMinMaxBytes{Val: "foo"}, false, "invalid StringEqualMinMaxBytes.Val: value length must be between 4 and 8 bytes, inclusive", 1},

	{"string - pattern - valid", &cases.StringPattern{Val: "Foo123"}, true, "", 0},
	{"string - pattern - invalid", &cases.StringPattern{Val: "!@#$%^&*()"}, false, "invalid StringPattern.Val: value does not match regex pattern \"(?i)^[a-z0-9]+$\"", 1},
	{"string - pattern - invalid (empty)", &cases.StringPattern{Val: ""}, false, "invalid StringPattern.Val: value does not match regex pattern \"(?i)^[a-z0-9]+$\"", 1},
	{"string - pattern - invalid (null)", &cases.StringPattern{Val: "a\000"}, false, "invalid StringPattern.Val: value does not match regex pattern \"(?i)^[a-z0-9]+$\"", 1},

	{"string - pattern (escapes) - valid", &cases.StringPatternEscapes{Val: "* \\ x"}, true, "", 0},
	{"string - pattern (escapes) - invalid", &cases.StringPatternEscapes{Val: "invalid"}, false, "invalid StringPatternEscapes.Val: value does not match regex pattern \"\\\\* \\\\\\\\ \\\\w\"", 1},
	{"string - pattern (escapes) - invalid (empty)", &cases.StringPatternEscapes{Val: ""}, false, "invalid StringPatternEscapes.Val: value does not match regex pattern \"\\\\* \\\\\\\\ \\\\w\"", 1},

	{"string - prefix - valid", &cases.StringPrefix{Val: "foobar"}, true, "", 0},
	{"string - prefix - valid (only)", &cases.StringPrefix{Val: "foo"}, true, "", 0},
	{"string - prefix - invalid", &cases.StringPrefix{Val: "bar"}, false, "invalid StringPrefix.Val: value does not have prefix \"foo\"", 1},
	{"string - prefix - invalid (case-sensitive)", &cases.StringPrefix{Val: "Foobar"}, false, "invalid StringPrefix.Val: value does not have prefix \"foo\"", 1},

	{"string - contains - valid", &cases.StringContains{Val: "candy bars"}, true, "", 0},
	{"string - contains - valid (only)", &cases.StringContains{Val: "bar"}, true, "", 0},
	{"string - contains - invalid", &cases.StringContains{Val: "candy bazs"}, false, "invalid StringContains.Val: value does not contain substring \"bar\"", 1},
	{"string - contains - invalid (case-sensitive)", &cases.StringContains{Val: "Candy Bars"}, false, "invalid StringContains.Val: value does not contain substring \"bar\"", 1},

	{"string - not contains - valid", &cases.StringNotContains{Val: "candy bazs"}, true, "", 0},
	{"string - not contains - valid (case-sensitive)", &cases.StringNotContains{Val: "Candy Bars"}, true, "", 0},
	{"string - not contains - invalid", &cases.StringNotContains{Val: "candy bars"}, false, "invalid StringNotContains.Val: value contains substring \"bar\"", 1},
	{"string - not contains - invalid (equal)", &cases.StringNotContains{Val: "bar"}, false, "invalid StringNotContains.Val: value contains substring \"bar\"", 1},

	{"string - suffix - valid", &cases.StringSuffix{Val: "foobaz"}, true, "", 0},
	{"string - suffix - valid (only)", &cases.StringSuffix{Val: "baz"}, true, "", 0},
	{"string - suffix - invalid", &cases.StringSuffix{Val: "foobar"}, false, "invalid StringSuffix.Val: value does not have suffix \"baz\"", 1},
	{"string - suffix - invalid (case-sensitive)", &cases.StringSuffix{Val: "FooBaz"}, false, "invalid StringSuffix.Val: value does not have suffix \"baz\"", 1},

	{"string - email - valid", &cases.StringEmail{Val: "foo@bar.com"}, true, "", 0},
	{"string - email - valid (name)", &cases.StringEmail{Val: "John Smith <foo@bar.com>"}, true, "", 0},
	{"string - email - invalid", &cases.StringEmail{Val: "foobar"}, false, "mail: missing '@' or angle-addr", 1},
	{"string - email - invalid (local segment too long)", &cases.StringEmail{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789@example.com"}, false, "email address local phrase cannot exceed 64 characters", 1},
	{"string - email - invalid (hostname too long)", &cases.StringEmail{Val: "foo@x0123456789012345678901234567890123456789012345678901234567890123456789.com"}, false, "hostname part must be non-empty and cannot exceed 63 characters", 1},
	{"string - email - invalid (bad hostname)", &cases.StringEmail{Val: "foo@-bar.com"}, false, "hostname parts cannot begin with hyphens", 1},
	{"string - email - empty", &cases.StringEmail{Val: ""}, false, "mail: no address", 1},

	{"string - address - valid hostname", &cases.StringAddress{Val: "example.com"}, true, "", 0},
	{"string - address - valid hostname (uppercase)", &cases.StringAddress{Val: "ASD.example.com"}, true, "", 0},
	{"string - address - valid hostname (hyphens)", &cases.StringAddress{Val: "foo-bar.com"}, true, "", 0},
	{"string - address - valid hostname (trailing dot)", &cases.StringAddress{Val: "example.com."}, true, "", 0},
	{"string - address - invalid hostname", &cases.StringAddress{Val: "!@#$%^&"}, false, "invalid StringAddress.Val: value must be a valid hostname, or ip address", 1},
	{"string - address - invalid hostname (underscore)", &cases.StringAddress{Val: "foo_bar.com"}, false, "invalid StringAddress.Val: value must be a valid hostname, or ip address", 1},
	{"string - address - invalid hostname (too long)", &cases.StringAddress{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789.com"}, false, "invalid StringAddress.Val: value must be a valid hostname, or ip address", 1},
	{"string - address - invalid hostname (trailing hyphens)", &cases.StringAddress{Val: "foo-bar-.com"}, false, "invalid StringAddress.Val: value must be a valid hostname, or ip address", 1},
	{"string - address - invalid hostname (leading hyphens)", &cases.StringAddress{Val: "foo-bar.-com"}, false, "invalid StringAddress.Val: value must be a valid hostname, or ip address", 1},
	{"string - address - invalid hostname (empty)", &cases.StringAddress{Val: "asd..asd.com"}, false, "invalid StringAddress.Val: value must be a valid hostname, or ip address", 1},
	{"string - address - invalid hostname (IDNs)", &cases.StringAddress{Val: "你好.com"}, false, "invalid StringAddress.Val: value must be a valid hostname, or ip address", 1},
	{"string - address - valid ip (v4)", &cases.StringAddress{Val: "192.168.0.1"}, true, "", 0},
	{"string - address - valid ip (v6)", &cases.StringAddress{Val: "3e::99"}, true, "", 0},
	{"string - address - invalid ip", &cases.StringAddress{Val: "ff::fff::0b"}, false, "invalid StringAddress.Val: value must be a valid hostname, or ip address", 1},

	{"string - hostname - valid", &cases.StringHostname{Val: "example.com"}, true, "", 0},
	{"string - hostname - valid (uppercase)", &cases.StringHostname{Val: "ASD.example.com"}, true, "", 0},
	{"string - hostname - valid (hyphens)", &cases.StringHostname{Val: "foo-bar.com"}, true, "", 0},
	{"string - hostname - valid (trailing dot)", &cases.StringHostname{Val: "example.com."}, true, "", 0},
	{"string - hostname - invalid", &cases.StringHostname{Val: "!@#$%^&"}, false, "hostname parts can only contain alphanumeric characters or hyphens, got \"!\"", 1},
	{"string - hostname - invalid (underscore)", &cases.StringHostname{Val: "foo_bar.com"}, false, "hostname parts can only contain alphanumeric characters or hyphens, got \"_\"", 1},
	{"string - hostname - invalid (too long)", &cases.StringHostname{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789.com"}, false, "hostname part must be non-empty and cannot exceed 63 characters", 1},
	{"string - hostname - invalid (trailing hyphens)", &cases.StringHostname{Val: "foo-bar-.com"}, false, "hostname parts cannot end with hyphens", 1},
	{"string - hostname - invalid (leading hyphens)", &cases.StringHostname{Val: "foo-bar.-com"}, false, "hostname parts cannot begin with hyphens", 1},
	{"string - hostname - invalid (empty)", &cases.StringHostname{Val: "asd..asd.com"}, false, "hostname part must be non-empty and cannot exceed 63 characters", 1},
	{"string - hostname - invalid (IDNs)", &cases.StringHostname{Val: "你好.com"}, false, "hostname parts can only contain alphanumeric characters or hyphens, got \"你\"", 1},

	{"string - IP - valid (v4)", &cases.StringIP{Val: "192.168.0.1"}, true, "", 0},
	{"string - IP - valid (v6)", &cases.StringIP{Val: "3e::99"}, true, "", 0},
	{"string - IP - invalid", &cases.StringIP{Val: "foobar"}, false, "invalid StringIP.Val: value must be a valid IP address", 1},

	{"string - IPv4 - valid", &cases.StringIPv4{Val: "192.168.0.1"}, true, "", 0},
	{"string - IPv4 - invalid", &cases.StringIPv4{Val: "foobar"}, false, "invalid StringIPv4.Val: value must be a valid IPv4 address", 1},
	{"string - IPv4 - invalid (erroneous)", &cases.StringIPv4{Val: "256.0.0.0"}, false, "invalid StringIPv4.Val: value must be a valid IPv4 address", 1},
	{"string - IPv4 - invalid (v6)", &cases.StringIPv4{Val: "3e::99"}, false, "invalid StringIPv4.Val: value must be a valid IPv4 address", 1},

	{"string - IPv6 - valid", &cases.StringIPv6{Val: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, true, "", 0},
	{"string - IPv6 - valid (collapsed)", &cases.StringIPv6{Val: "2001:db8:85a3::8a2e:370:7334"}, true, "", 0},
	{"string - IPv6 - invalid", &cases.StringIPv6{Val: "foobar"}, false, "invalid StringIPv6.Val: value must be a valid IPv6 address", 1},
	{"string - IPv6 - invalid (v4)", &cases.StringIPv6{Val: "192.168.0.1"}, false, "invalid StringIPv6.Val: value must be a valid IPv6 address", 1},
	{"string - IPv6 - invalid (erroneous)", &cases.StringIPv6{Val: "ff::fff::0b"}, false, "invalid StringIPv6.Val: value must be a valid IPv6 address", 1},

	{"string - URI - valid", &cases.StringURI{Val: "http://example.com/foo/bar?baz=quux"}, true, "", 0},
	{"string - URI - invalid", &cases.StringURI{Val: "!@#$%^&*%$#"}, false, "parse \"!@#$%^&*%$#\": invalid URL escape \"%^&\"", 1},
	{"string - URI - invalid (relative)", &cases.StringURI{Val: "/foo/bar?baz=quux"}, false, "invalid StringURI.Val: value must be absolute", 1},

	{"string - URI Ref - valid", &cases.StringURIRef{Val: "http://example.com/foo/bar?baz=quux"}, true, "", 0},
	{"string - URI Ref - valid (relative)", &cases.StringURIRef{Val: "/foo/bar?baz=quux"}, true, "", 0},
	{"string - URI Ref - invalid", &cases.StringURIRef{Val: "!@#$%^&*%$#"}, false, "parse \"!@#$%^&*%$#\": invalid URL escape \"%^&\"", 1},

	{"string - UUID - valid (nil)", &cases.StringUUID{Val: "00000000-0000-0000-0000-000000000000"}, true, "", 0},
	{"string - UUID - valid (v1)", &cases.StringUUID{Val: "b45c0c80-8880-11e9-a5b1-000000000000"}, true, "", 0},
	{"string - UUID - valid (v1 - case-insensitive)", &cases.StringUUID{Val: "B45C0C80-8880-11E9-A5B1-000000000000"}, true, "", 0},
	{"string - UUID - valid (v2)", &cases.StringUUID{Val: "b45c0c80-8880-21e9-a5b1-000000000000"}, true, "", 0},
	{"string - UUID - valid (v2 - case-insensitive)", &cases.StringUUID{Val: "B45C0C80-8880-21E9-A5B1-000000000000"}, true, "", 0},
	{"string - UUID - valid (v3)", &cases.StringUUID{Val: "a3bb189e-8bf9-3888-9912-ace4e6543002"}, true, "", 0},
	{"string - UUID - valid (v3 - case-insensitive)", &cases.StringUUID{Val: "A3BB189E-8BF9-3888-9912-ACE4E6543002"}, true, "", 0},
	{"string - UUID - valid (v4)", &cases.StringUUID{Val: "8b208305-00e8-4460-a440-5e0dcd83bb0a"}, true, "", 0},
	{"string - UUID - valid (v4 - case-insensitive)", &cases.StringUUID{Val: "8B208305-00E8-4460-A440-5E0DCD83BB0A"}, true, "", 0},
	{"string - UUID - valid (v5)", &cases.StringUUID{Val: "a6edc906-2f9f-5fb2-a373-efac406f0ef2"}, true, "", 0},
	{"string - UUID - valid (v5 - case-insensitive)", &cases.StringUUID{Val: "A6EDC906-2F9F-5FB2-A373-EFAC406F0EF2"}, true, "", 0},
	{"string - UUID - invalid", &cases.StringUUID{Val: "foobar"}, false, "invalid uuid format", 1},
	{"string - UUID - invalid (bad UUID)", &cases.StringUUID{Val: "ffffffff-ffff-ffff-ffff-fffffffffffff"}, false, "invalid uuid format", 1},

	{"string - http header name - valid", &cases.StringHttpHeaderName{Val: "clustername"}, true, "", 0},
	{"string - http header name - valid", &cases.StringHttpHeaderName{Val: ":path"}, true, "", 0},
	{"string - http header name - valid (nums)", &cases.StringHttpHeaderName{Val: "cluster-123"}, true, "", 0},
	{"string - http header name - valid (special token)", &cases.StringHttpHeaderName{Val: "!+#&.%"}, true, "", 0},
	{"string - http header name - valid (period)", &cases.StringHttpHeaderName{Val: "CLUSTER.NAME"}, true, "", 0},
	{"string - http header name - invalid", &cases.StringHttpHeaderName{Val: ":"}, false, "invalid StringHttpHeaderName.Val: value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"", 1},
	{"string - http header name - invalid", &cases.StringHttpHeaderName{Val: ":path:"}, false, "invalid StringHttpHeaderName.Val: value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"", 1},
	{"string - http header name - invalid (space)", &cases.StringHttpHeaderName{Val: "cluster name"}, false, "invalid StringHttpHeaderName.Val: value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"", 1},
	{"string - http header name - invalid (return)", &cases.StringHttpHeaderName{Val: "example\r"}, false, "invalid StringHttpHeaderName.Val: value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"", 1},
	{"string - http header name - invalid (tab)", &cases.StringHttpHeaderName{Val: "example\t"}, false, "invalid StringHttpHeaderName.Val: value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"", 1},
	{"string - http header name - invalid (slash)", &cases.StringHttpHeaderName{Val: "/test/long/url"}, false, "invalid StringHttpHeaderName.Val: value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"", 1},

	{"string - http header value - valid", &cases.StringHttpHeaderValue{Val: "cluster.name.123"}, true, "", 0},
	{"string - http header value - valid (uppercase)", &cases.StringHttpHeaderValue{Val: "/TEST/LONG/URL"}, true, "", 0},
	{"string - http header value - valid (spaces)", &cases.StringHttpHeaderValue{Val: "cluster name"}, true, "", 0},
	{"string - http header value - valid (tab)", &cases.StringHttpHeaderValue{Val: "example\t"}, true, "", 0},
	{"string - http header value - valid (special token)", &cases.StringHttpHeaderValue{Val: "!#%&./+"}, true, "", 0},
	{"string - http header value - invalid (NUL)", &cases.StringHttpHeaderValue{Val: "foo\u0000bar"}, false, "invalid StringHttpHeaderValue.Val: value does not match regex pattern \"^[^\\x00-\\b\\n-\\x1f\\u007f]*$\"", 1},
	{"string - http header value - invalid (DEL)", &cases.StringHttpHeaderValue{Val: "\u007f"}, false, "invalid StringHttpHeaderValue.Val: value does not match regex pattern \"^[^\\x00-\\b\\n-\\x1f\\u007f]*$\"", 1},
	{"string - http header value - invalid", &cases.StringHttpHeaderValue{Val: "example\r"}, false, "invalid StringHttpHeaderValue.Val: value does not match regex pattern \"^[^\\x00-\\b\\n-\\x1f\\u007f]*$\"", 1},

	{"string - non-strict valid header - valid", &cases.StringValidHeader{Val: "cluster.name.123"}, true, "", 0},
	{"string - non-strict valid header - valid (uppercase)", &cases.StringValidHeader{Val: "/TEST/LONG/URL"}, true, "", 0},
	{"string - non-strict valid header - valid (spaces)", &cases.StringValidHeader{Val: "cluster name"}, true, "", 0},
	{"string - non-strict valid header - valid (tab)", &cases.StringValidHeader{Val: "example\t"}, true, "", 0},
	{"string - non-strict valid header - valid (DEL)", &cases.StringValidHeader{Val: "\u007f"}, true, "", 0},
	{"string - non-strict valid header - invalid (NUL)", &cases.StringValidHeader{Val: "foo\u0000bar"}, false, "invalid StringValidHeader.Val: value does not match regex pattern \"^[^\\x00\\n\\r]*$\"", 1},
	{"string - non-strict valid header - invalid (CR)", &cases.StringValidHeader{Val: "example\r"}, false, "invalid StringValidHeader.Val: value does not match regex pattern \"^[^\\x00\\n\\r]*$\"", 1},
	{"string - non-strict valid header - invalid (NL)", &cases.StringValidHeader{Val: "exa\u000Ample"}, false, "invalid StringValidHeader.Val: value does not match regex pattern \"^[^\\x00\\n\\r]*$\"", 1},
}

var bytesCases = []TestCase{
	{"bytes - none - valid", &cases.BytesNone{Val: []byte("quux")}, true, "", 0},

	{"bytes - const - valid", &cases.BytesConst{Val: []byte("foo")}, true, "", 0},
	{"bytes - const - invalid", &cases.BytesConst{Val: []byte("bar")}, false, "invalid BytesConst.Val: value must equal [102 111 111]", 1},

	{"bytes - in - valid", &cases.BytesIn{Val: []byte("bar")}, true, "", 0},
	{"bytes - in - invalid", &cases.BytesIn{Val: []byte("quux")}, false, "invalid BytesIn.Val: value must be in list [[98 97 114] [98 97 122]]", 1},
	{"bytes - not in - valid", &cases.BytesNotIn{Val: []byte("quux")}, true, "", 0},
	{"bytes - not in - invalid", &cases.BytesNotIn{Val: []byte("fizz")}, false, "invalid BytesNotIn.Val: value must not be in list [[102 105 122 122] [98 117 122 122]]", 1},

	{"bytes - len - valid", &cases.BytesLen{Val: []byte("baz")}, true, "", 0},
	{"bytes - len - invalid (lt)", &cases.BytesLen{Val: []byte("go")}, false, "invalid BytesLen.Val: value length must be 3 bytes", 1},
	{"bytes - len - invalid (gt)", &cases.BytesLen{Val: []byte("fizz")}, false, "invalid BytesLen.Val: value length must be 3 bytes", 1},

	{"bytes - min len - valid", &cases.BytesMinLen{Val: []byte("fizz")}, true, "", 0},
	{"bytes - min len - valid (min)", &cases.BytesMinLen{Val: []byte("baz")}, true, "", 0},
	{"bytes - min len - invalid", &cases.BytesMinLen{Val: []byte("go")}, false, "invalid BytesMinLen.Val: value length must be at least 3 bytes", 1},

	{"bytes - max len - valid", &cases.BytesMaxLen{Val: []byte("foo")}, true, "", 0},
	{"bytes - max len - valid (max)", &cases.BytesMaxLen{Val: []byte("proto")}, true, "", 0},
	{"bytes - max len - invalid", &cases.BytesMaxLen{Val: []byte("1234567890")}, false, "invalid BytesMaxLen.Val: value length must be at most 5 bytes", 1},

	{"bytes - min/max len - valid", &cases.BytesMinMaxLen{Val: []byte("quux")}, true, "", 0},
	{"bytes - min/max len - valid (min)", &cases.BytesMinMaxLen{Val: []byte("foo")}, true, "", 0},
	{"bytes - min/max len - valid (max)", &cases.BytesMinMaxLen{Val: []byte("proto")}, true, "", 0},
	{"bytes - min/max len - invalid (below)", &cases.BytesMinMaxLen{Val: []byte("go")}, false, "invalid BytesMinMaxLen.Val: value length must be between 3 and 5 bytes, inclusive", 1},
	{"bytes - min/max len - invalid (above)", &cases.BytesMinMaxLen{Val: []byte("validate")}, false, "invalid BytesMinMaxLen.Val: value length must be between 3 and 5 bytes, inclusive", 1},

	{"bytes - equal min/max len - valid", &cases.BytesEqualMinMaxLen{Val: []byte("proto")}, true, "", 0},
	{"bytes - equal min/max len - invalid", &cases.BytesEqualMinMaxLen{Val: []byte("validate")}, false, "invalid BytesEqualMinMaxLen.Val: value length must be 5 bytes", 1},

	{"bytes - pattern - valid", &cases.BytesPattern{Val: []byte("Foo123")}, true, "", 0},
	{"bytes - pattern - invalid", &cases.BytesPattern{Val: []byte("你好你好")}, false, "invalid BytesPattern.Val: value does not match regex pattern \"^[\\x00-\\u007f]+$\"", 1},
	{"bytes - pattern - invalid (empty)", &cases.BytesPattern{Val: []byte("")}, false, "invalid BytesPattern.Val: value does not match regex pattern \"^[\\x00-\\u007f]+$\"", 1},

	{"bytes - prefix - valid", &cases.BytesPrefix{Val: []byte{0x99, 0x9f, 0x08}}, true, "", 0},
	{"bytes - prefix - valid (only)", &cases.BytesPrefix{Val: []byte{0x99}}, true, "", 0},
	{"bytes - prefix - invalid", &cases.BytesPrefix{Val: []byte("bar")}, false, "invalid BytesPrefix.Val: value does not have prefix \"\\x99\"", 1},

	{"bytes - contains - valid", &cases.BytesContains{Val: []byte("candy bars")}, true, "", 0},
	{"bytes - contains - valid (only)", &cases.BytesContains{Val: []byte("bar")}, true, "", 0},
	{"bytes - contains - invalid", &cases.BytesContains{Val: []byte("candy bazs")}, false, "invalid BytesContains.Val: value does not contain \"\\x62\\x61\\x72\"", 1},

	{"bytes - suffix - valid", &cases.BytesSuffix{Val: []byte{0x62, 0x75, 0x7A, 0x7A}}, true, "", 0},
	{"bytes - suffix - valid (only)", &cases.BytesSuffix{Val: []byte("\x62\x75\x7A\x7A")}, true, "", 0},
	{"bytes - suffix - invalid", &cases.BytesSuffix{Val: []byte("foobar")}, false, "invalid BytesSuffix.Val: value does not have suffix \"\\x62\\x75\\x7A\\x7A\"", 1},
	{"bytes - suffix - invalid (case-sensitive)", &cases.BytesSuffix{Val: []byte("FooBaz")}, false, "invalid BytesSuffix.Val: value does not have suffix \"\\x62\\x75\\x7A\\x7A\"", 1},

	{"bytes - IP - valid (v4)", &cases.BytesIP{Val: []byte{0xC0, 0xA8, 0x00, 0x01}}, true, "", 0},
	{"bytes - IP - valid (v6)", &cases.BytesIP{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")}, true, "", 0},
	{"bytes - IP - invalid", &cases.BytesIP{Val: []byte("foobar")}, false, "invalid BytesIP.Val: value must be a valid IP address", 1},

	{"bytes - IPv4 - valid", &cases.BytesIPv4{Val: []byte{0xC0, 0xA8, 0x00, 0x01}}, true, "", 0},
	{"bytes - IPv4 - invalid", &cases.BytesIPv4{Val: []byte("foobar")}, false, "invalid BytesIPv4.Val: value must be a valid IPv4 address", 1},
	{"bytes - IPv4 - invalid (v6)", &cases.BytesIPv4{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")}, false, "invalid BytesIPv4.Val: value must be a valid IPv4 address", 1},

	{"bytes - IPv6 - valid", &cases.BytesIPv6{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")}, true, "", 0},
	{"bytes - IPv6 - invalid", &cases.BytesIPv6{Val: []byte("fooar")}, false, "invalid BytesIPv6.Val: value must be a valid IPv6 address", 1},
	{"bytes - IPv6 - invalid (v4)", &cases.BytesIPv6{Val: []byte{0xC0, 0xA8, 0x00, 0x01}}, false, "invalid BytesIPv6.Val: value must be a valid IPv6 address", 1},
}

var enumCases = []TestCase{
	{"enum - none - valid", &cases.EnumNone{Val: cases.TestEnum_ONE}, true, "", 0},

	{"enum - const - valid", &cases.EnumConst{Val: cases.TestEnum_TWO}, true, "", 0},
	{"enum - const - invalid", &cases.EnumConst{Val: cases.TestEnum_ONE}, false, "invalid EnumConst.Val: value must equal 2", 1},
	{"enum alias - const - valid", &cases.EnumAliasConst{Val: cases.TestEnumAlias_C}, true, "", 0},
	{"enum alias - const - valid (alias)", &cases.EnumAliasConst{Val: cases.TestEnumAlias_GAMMA}, true, "", 0},
	{"enum alias - const - invalid", &cases.EnumAliasConst{Val: cases.TestEnumAlias_ALPHA}, false, "invalid EnumAliasConst.Val: value must equal 2", 1},

	{"enum - defined_only - valid", &cases.EnumDefined{Val: 0}, true, "", 0},
	{"enum - defined_only - invalid", &cases.EnumDefined{Val: math.MaxInt32}, false, "invalid EnumDefined.Val: value must be one of the defined enum values", 1},
	{"enum alias - defined_only - valid", &cases.EnumAliasDefined{Val: 1}, true, "", 0},
	{"enum alias - defined_only - invalid", &cases.EnumAliasDefined{Val: math.MaxInt32}, false, "invalid EnumAliasDefined.Val: value must be one of the defined enum values", 1},

	{"enum - in - valid", &cases.EnumIn{Val: cases.TestEnum_TWO}, true, "", 0},
	{"enum - in - invalid", &cases.EnumIn{Val: cases.TestEnum_ONE}, false, "invalid EnumIn.Val: value must be in list [0 2]", 1},
	{"enum alias - in - valid", &cases.EnumAliasIn{Val: cases.TestEnumAlias_A}, true, "", 0},
	{"enum alias - in - valid (alias)", &cases.EnumAliasIn{Val: cases.TestEnumAlias_ALPHA}, true, "", 0},
	{"enum alias - in - invalid", &cases.EnumAliasIn{Val: cases.TestEnumAlias_BETA}, false, "invalid EnumAliasIn.Val: value must be in list [0 2]", 1},

	{"enum - not in - valid", &cases.EnumNotIn{Val: cases.TestEnum_ZERO}, true, "", 0},
	{"enum - not in - valid (undefined)", &cases.EnumNotIn{Val: math.MaxInt32}, true, "", 0},
	{"enum - not in - invalid", &cases.EnumNotIn{Val: cases.TestEnum_ONE}, false, "invalid EnumNotIn.Val: value must not be in list [1]", 1},
	{"enum alias - not in - valid", &cases.EnumAliasNotIn{Val: cases.TestEnumAlias_ALPHA}, true, "", 0},
	{"enum alias - not in - invalid", &cases.EnumAliasNotIn{Val: cases.TestEnumAlias_B}, false, "invalid EnumAliasNotIn.Val: value must not be in list [1]", 1},
	{"enum alias - not in - invalid (alias)", &cases.EnumAliasNotIn{Val: cases.TestEnumAlias_BETA}, false, "invalid EnumAliasNotIn.Val: value must not be in list [1]", 1},

	{"enum external - defined_only - valid", &cases.EnumExternal{Val: other_package.Embed_VALUE}, true, "", 0},
	{"enum external - defined_only - invalid", &cases.EnumExternal{Val: math.MaxInt32}, false, "invalid EnumExternal.Val: value must be one of the defined enum values", 1},

	{"enum repeated - defined_only - valid", &cases.RepeatedEnumDefined{Val: []cases.TestEnum{cases.TestEnum_ONE, cases.TestEnum_TWO}}, true, "", 0},
	{"enum repeated - defined_only - invalid", &cases.RepeatedEnumDefined{Val: []cases.TestEnum{cases.TestEnum_ONE, math.MaxInt32}}, false, "invalid RepeatedEnumDefined.Val[1]: value must be one of the defined enum values", 1},

	{"enum repeated (external) - defined_only - valid", &cases.RepeatedExternalEnumDefined{Val: []other_package.Embed_Enumerated{other_package.Embed_VALUE}}, true, "", 0},
	{"enum repeated (external) - defined_only - invalid", &cases.RepeatedExternalEnumDefined{Val: []other_package.Embed_Enumerated{math.MaxInt32}}, false, "invalid RepeatedExternalEnumDefined.Val[0]: value must be one of the defined enum values", 1},

	{"enum map - defined_only - valid", &cases.MapEnumDefined{Val: map[string]cases.TestEnum{"foo": cases.TestEnum_TWO}}, true, "", 0},
	{"enum map - defined_only - invalid", &cases.MapEnumDefined{Val: map[string]cases.TestEnum{"foo": math.MaxInt32}}, false, "invalid MapEnumDefined.Val[foo]: value must be one of the defined enum values", 1},

	{"enum map (external) - defined_only - valid", &cases.MapExternalEnumDefined{Val: map[string]other_package.Embed_Enumerated{"foo": other_package.Embed_VALUE}}, true, "", 0},
	{"enum map (external) - defined_only - invalid", &cases.MapExternalEnumDefined{Val: map[string]other_package.Embed_Enumerated{"foo": math.MaxInt32}}, false, "invalid MapExternalEnumDefined.Val[foo]: value must be one of the defined enum values", 1},
}

var messageCases = []TestCase{
	{"message - none - valid", &cases.MessageNone{Val: &cases.MessageNone_NoneMsg{}}, true, "", 0},
	{"message - none - valid (unset)", &cases.MessageNone{}, true, "", 0},

	{"message - disabled - valid", &cases.MessageDisabled{Val: 456}, true, "", 0},
	{"message - disabled - valid (invalid field)", &cases.MessageDisabled{Val: 0}, true, "", 0},

	{"message - ignored - valid", &cases.MessageIgnored{Val: 456}, true, "", 0},
	{"message - ignored - valid (invalid field)", &cases.MessageIgnored{Val: 0}, true, "", 0},

	{"message - field - valid", &cases.Message{Val: &cases.TestMsg{Const: "foo"}}, true, "", 0},
	{"message - field - valid (unset)", &cases.Message{}, true, "", 0},
	{"message - field - invalid", &cases.Message{Val: &cases.TestMsg{}}, false, "invalid TestMsg.Const: value must equal foo", 1},
	{"message - field - invalid (transitive)", &cases.Message{Val: &cases.TestMsg{Const: "foo", Nested: &cases.TestMsg{}}}, false, "invalid TestMsg.Const: value must equal foo", 1},

	{"message - skip - valid", &cases.MessageSkip{Val: &cases.TestMsg{}}, true, "", 0},

	{"message - required - valid", &cases.MessageRequired{Val: &cases.TestMsg{Const: "foo"}}, true, "", 0},
	{"message - required - invalid", &cases.MessageRequired{}, false, "invalid MessageRequired.Val: value is required", 1},

	{"message - cross-package embed none - valid", &cases.MessageCrossPackage{Val: &other_package.Embed{Val: 1}}, true, "", 0},
	{"message - cross-package embed none - valid (nil)", &cases.MessageCrossPackage{}, true, "", 0},
	{"message - cross-package embed none - invalid (empty)", &cases.MessageCrossPackage{Val: &other_package.Embed{}}, false, "invalid Embed.Val: value must be greater than 0", 1},
	{"message - cross-package embed none - invalid", &cases.MessageCrossPackage{Val: &other_package.Embed{Val: -1}}, false, "invalid Embed.Val: value must be greater than 0", 1},
}

var repeatedCases = []TestCase{
	{"repeated - none - valid", &cases.RepeatedNone{Val: []int64{1, 2, 3}}, true, "", 0},

	{"repeated - embed none - valid", &cases.RepeatedEmbedNone{Val: []*cases.Embed{{Val: 1}}}, true, "", 0},
	{"repeated - embed none - valid (nil)", &cases.RepeatedEmbedNone{}, true, "", 0},
	{"repeated - embed none - valid (empty)", &cases.RepeatedEmbedNone{Val: []*cases.Embed{}}, true, "", 0},
	{"repeated - embed none - invalid", &cases.RepeatedEmbedNone{Val: []*cases.Embed{{Val: -1}}}, false, "invalid Embed.Val: value must be greater than 0", 1},

	{"repeated - cross-package embed none - valid", &cases.RepeatedEmbedCrossPackageNone{Val: []*other_package.Embed{{Val: 1}}}, true, "", 0},
	{"repeated - cross-package embed none - valid (nil)", &cases.RepeatedEmbedCrossPackageNone{}, true, "", 0},
	{"repeated - cross-package embed none - valid (empty)", &cases.RepeatedEmbedCrossPackageNone{Val: []*other_package.Embed{}}, true, "", 0},
	{"repeated - cross-package embed none - invalid", &cases.RepeatedEmbedCrossPackageNone{Val: []*other_package.Embed{{Val: -1}}}, false, "invalid Embed.Val: value must be greater than 0", 1},

	{"repeated - min - valid", &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}, {Val: 2}, {Val: 3}}}, true, "", 0},
	{"repeated - min - valid (equal)", &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}, {Val: 2}}}, true, "", 0},
	{"repeated - min - invalid", &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}}}, false, "invalid RepeatedMin.Val: value must contain at least 2 item(s)", 1},
	{"repeated - min - invalid (element)", &cases.RepeatedMin{Val: []*cases.Embed{{Val: 1}, {Val: -1}}}, false, "invalid Embed.Val: value must be greater than 0", 1},

	{"repeated - max - valid", &cases.RepeatedMax{Val: []float64{1, 2}}, true, "", 0},
	{"repeated - max - valid (equal)", &cases.RepeatedMax{Val: []float64{1, 2, 3}}, true, "", 0},
	{"repeated - max - invalid", &cases.RepeatedMax{Val: []float64{1, 2, 3, 4}}, false, "invalid RepeatedMax.Val: value must contain no more than 3 item(s)", 1},

	{"repeated - min/max - valid", &cases.RepeatedMinMax{Val: []int32{1, 2, 3}}, true, "", 0},
	{"repeated - min/max - valid (min)", &cases.RepeatedMinMax{Val: []int32{1, 2}}, true, "", 0},
	{"repeated - min/max - valid (max)", &cases.RepeatedMinMax{Val: []int32{1, 2, 3, 4}}, true, "", 0},
	{"repeated - min/max - invalid (below)", &cases.RepeatedMinMax{Val: []int32{}}, false, "invalid RepeatedMinMax.Val: value must contain between 2 and 4 items, inclusive", 1},
	{"repeated - min/max - invalid (above)", &cases.RepeatedMinMax{Val: []int32{1, 2, 3, 4, 5}}, false, "invalid RepeatedMinMax.Val: value must contain between 2 and 4 items, inclusive", 1},

	{"repeated - exact - valid", &cases.RepeatedExact{Val: []uint32{1, 2, 3}}, true, "", 0},
	{"repeated - exact - invalid (below)", &cases.RepeatedExact{Val: []uint32{1, 2}}, false, "invalid RepeatedExact.Val: value must contain exactly 3 item(s)", 1},
	{"repeated - exact - invalid (above)", &cases.RepeatedExact{Val: []uint32{1, 2, 3, 4}}, false, "invalid RepeatedExact.Val: value must contain exactly 3 item(s)", 1},

	{"repeated - unique - valid", &cases.RepeatedUnique{Val: []string{"foo", "bar", "baz"}}, true, "", 0},
	{"repeated - unique - valid (empty)", &cases.RepeatedUnique{}, true, "", 0},
	{"repeated - unique - valid (case sensitivity)", &cases.RepeatedUnique{Val: []string{"foo", "Foo"}}, true, "", 0},
	{"repeated - unique - invalid", &cases.RepeatedUnique{Val: []string{"foo", "bar", "foo", "baz"}}, false, "invalid RepeatedUnique.Val[2]: repeated value must contain unique items", 1},

	{"repeated - items - valid", &cases.RepeatedItemRule{Val: []float32{1, 2, 3}}, true, "", 0},
	{"repeated - items - valid (empty)", &cases.RepeatedItemRule{Val: []float32{}}, true, "", 0},
	{"repeated - items - valid (pattern)", &cases.RepeatedItemPattern{Val: []string{"Alpha", "Beta123"}}, true, "", 0},
	{"repeated - items - invalid", &cases.RepeatedItemRule{Val: []float32{1, -2, 3}}, false, "invalid RepeatedItemRule.Val[1]: value must be greater than 0", 1},
	{"repeated - items - invalid (pattern)", &cases.RepeatedItemPattern{Val: []string{"Alpha", "!@#$%^&*()"}}, false, "invalid RepeatedItemPattern.Val[1]: value does not match regex pattern \"(?i)^[a-z0-9]+$\"", 1},
	{"repeated - items - invalid (in)", &cases.RepeatedItemIn{Val: []string{"baz"}}, false, "invalid RepeatedItemIn.Val[0]: value must be in list [foo bar]", 1},
	{"repeated - items - valid (in)", &cases.RepeatedItemIn{Val: []string{"foo"}}, true, "", 0},
	{"repeated - items - invalid (not_in)", &cases.RepeatedItemNotIn{Val: []string{"foo"}}, false, "invalid RepeatedItemNotIn.Val[0]: value must not be in list [foo bar]", 1},
	{"repeated - items - valid (not_in)", &cases.RepeatedItemNotIn{Val: []string{"baz"}}, true, "", 0},

	{"repeated - items - invalid (enum in)", &cases.RepeatedEnumIn{Val: []cases.AnEnum{1}}, false, "invalid RepeatedEnumIn.Val[0]: value must be in list [0]", 1},
	{"repeated - items - valid (enum in)", &cases.RepeatedEnumIn{Val: []cases.AnEnum{0}}, true, "", 0},
	{"repeated - items - invalid (enum not_in)", &cases.RepeatedEnumNotIn{Val: []cases.AnEnum{0}}, false, "invalid RepeatedEnumNotIn.Val[0]: value must not be in list [0]", 1},
	{"repeated - items - valid (enum not_in)", &cases.RepeatedEnumNotIn{Val: []cases.AnEnum{1}}, true, "", 0},
	{"repeated - items - invalid (embedded enum in)", &cases.RepeatedEmbeddedEnumIn{Val: []cases.RepeatedEmbeddedEnumIn_AnotherInEnum{1}}, false, "invalid RepeatedEmbeddedEnumIn.Val[0]: value must be in list [0]", 1},
	{"repeated - items - valid (embedded enum in)", &cases.RepeatedEmbeddedEnumIn{Val: []cases.RepeatedEmbeddedEnumIn_AnotherInEnum{0}}, true, "", 0},
	{"repeated - items - invalid (embedded enum not_in)", &cases.RepeatedEmbeddedEnumNotIn{Val: []cases.RepeatedEmbeddedEnumNotIn_AnotherNotInEnum{0}}, false, "invalid RepeatedEmbeddedEnumNotIn.Val[0]: value must not be in list [0]", 1},
	{"repeated - items - valid (embedded enum not_in)", &cases.RepeatedEmbeddedEnumNotIn{Val: []cases.RepeatedEmbeddedEnumNotIn_AnotherNotInEnum{1}}, true, "", 0},

	{"repeated - embed skip - valid", &cases.RepeatedEmbedSkip{Val: []*cases.Embed{{Val: 1}}}, true, "", 0},
	{"repeated - embed skip - valid (invalid element)", &cases.RepeatedEmbedSkip{Val: []*cases.Embed{{Val: -1}}}, true, "", 0},
	{"repeated - min and items len - valid", &cases.RepeatedMinAndItemLen{Val: []string{"aaa", "bbb"}}, true, "", 0},
	{"repeated - min and items len - invalid (min)", &cases.RepeatedMinAndItemLen{Val: []string{}}, false, "invalid RepeatedMinAndItemLen.Val: value must contain at least 1 item(s)", 1},
	{"repeated - min and items len - invalid (len)", &cases.RepeatedMinAndItemLen{Val: []string{"x"}}, false, "invalid RepeatedMinAndItemLen.Val[0]: value length must be 3 runes", 1},
	{"repeated - min and max items len - valid", &cases.RepeatedMinAndMaxItemLen{Val: []string{"aaa", "bbb"}}, true, "", 0},
	{"repeated - min and max items len - invalid (min_len)", &cases.RepeatedMinAndMaxItemLen{}, false, "invalid RepeatedMinAndMaxItemLen.Val: value must contain between 1 and 3 items, inclusive", 1},
	{"repeated - min and max items len - invalid (max_len)", &cases.RepeatedMinAndMaxItemLen{Val: []string{"aaa", "bbb", "ccc", "ddd"}}, false, "invalid RepeatedMinAndMaxItemLen.Val: value must contain between 1 and 3 items, inclusive", 1},

	{"repeated - duration - gte - valid", &cases.RepeatedDuration{Val: []*duration.Duration{{Seconds: 3}}}, true, "", 0},
	{"repeated - duration - gte - valid (empty)", &cases.RepeatedDuration{}, true, "", 0},
	{"repeated - duration - gte - valid (equal)", &cases.RepeatedDuration{Val: []*duration.Duration{{Nanos: 1000000}}}, true, "", 0},
	{"repeated - duration - gte - invalid", &cases.RepeatedDuration{Val: []*duration.Duration{{Seconds: -1}}}, false, "invalid RepeatedDuration.Val[0]: value must be greater than or equal to 1ms", 1},
}

var mapCases = []TestCase{
	{"map - none - valid", &cases.MapNone{Val: map[uint32]bool{123: true, 456: false}}, true, "", 0},

	{"map - min pairs - valid", &cases.MapMin{Val: map[int32]float32{1: 2, 3: 4, 5: 6}}, true, "", 0},
	{"map - min pairs - valid (equal)", &cases.MapMin{Val: map[int32]float32{1: 2, 3: 4}}, true, "", 0},
	{"map - min pairs - invalid", &cases.MapMin{Val: map[int32]float32{1: 2}}, false, "invalid MapMin.Val: value must contain at least 2 pair(s)", 1},

	{"map - max pairs - valid", &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4}}, true, "", 0},
	{"map - max pairs - valid (equal)", &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4, 5: 6}}, true, "", 0},
	{"map - max pairs - invalid", &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4, 5: 6, 7: 8}}, false, "invalid MapMax.Val: value must contain no more than 3 pair(s)", 1},

	{"map - min/max - valid", &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true}}, true, "", 0},
	{"map - min/max - valid (min)", &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false}}, true, "", 0},
	{"map - min/max - valid (max)", &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true, "d": false}}, true, "", 0},
	{"map - min/max - invalid (below)", &cases.MapMinMax{Val: map[string]bool{}}, false, "invalid MapMinMax.Val: value must contain between 2 and 4 pairs, inclusive", 1},
	{"map - min/max - invalid (above)", &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true, "d": false, "e": true}}, false, "invalid MapMinMax.Val: value must contain between 2 and 4 pairs, inclusive", 1},

	{"map - exact - valid", &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b", 3: "c"}}, true, "", 0},
	{"map - exact - invalid (below)", &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b"}}, false, "invalid MapExact.Val: value must contain exactly 3 pair(s)", 1},
	{"map - exact - invalid (above)", &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b", 3: "c", 4: "d"}}, false, "invalid MapExact.Val: value must contain exactly 3 pair(s)", 1},

	{"map - no sparse - valid", &cases.MapNoSparse{Val: map[uint32]*cases.MapNoSparse_Msg{1: {}, 2: {}}}, true, "", 0},
	{"map - no sparse - valid (empty)", &cases.MapNoSparse{Val: map[uint32]*cases.MapNoSparse_Msg{}}, true, "", 0},
	// sparse maps are no longer supported, so this case is no longer possible
	//{"map - no sparse - invalid", &cases.MapNoSparse{Val: map[uint32]*cases.MapNoSparse_Msg{1: {}, 2: nil}}, false, "", 1},

	{"map - keys - valid", &cases.MapKeys{Val: map[int64]string{-1: "a", -2: "b"}}, true, "", 0},
	{"map - keys - valid (empty)", &cases.MapKeys{Val: map[int64]string{}}, true, "", 0},
	{"map - keys - valid (pattern)", &cases.MapKeysPattern{Val: map[string]string{"A": "a"}}, true, "", 0},
	{"map - keys - invalid", &cases.MapKeys{Val: map[int64]string{1: "a"}}, false, "invalid MapKeys.Val[1]: value must be less than 0", 1},
	{"map - keys - invalid (pattern)", &cases.MapKeysPattern{Val: map[string]string{"A": "a", "!@#$%^&*()": "b"}}, false, "invalid MapKeysPattern.Val[!@#$%^&*()]: value does not match regex pattern \"(?i)^[a-z0-9]+$\"", 1},

	{"map - values - valid", &cases.MapValues{Val: map[string]string{"a": "Alpha", "b": "Beta"}}, true, "", 0},
	{"map - values - valid (empty)", &cases.MapValues{Val: map[string]string{}}, true, "", 0},
	{"map - values - valid (pattern)", &cases.MapValuesPattern{Val: map[string]string{"a": "A"}}, true, "", 0},
	{"map - values - invalid", &cases.MapValues{Val: map[string]string{"a": "A", "b": "BCD"}}, false, "invalid MapValues.Val[a]: value length must be at least 3 runes", 1},
	{"map - values - invalid (pattern)", &cases.MapValuesPattern{Val: map[string]string{"a": "A", "b": "!@#$%^&*()"}}, false, "invalid MapValuesPattern.Val[b]: value does not match regex pattern \"(?i)^[a-z0-9]+$\"", 1},

	{"map - recursive - valid", &cases.MapRecursive{Val: map[uint32]*cases.MapRecursive_Msg{1: {Val: "abc"}}}, true, "", 0},
	{"map - recursive - invalid", &cases.MapRecursive{Val: map[uint32]*cases.MapRecursive_Msg{1: {}}}, false, "invalid MapRecursive_Msg.Val: value length must be at least 3 runes", 1},
}

var oneofCases = []TestCase{
	{"oneof - none - valid", &cases.OneOfNone{O: &cases.OneOfNone_X{X: "foo"}}, true, "", 0},
	{"oneof - none - valid (empty)", &cases.OneOfNone{}, true, "", 0},

	{"oneof - field - valid (X)", &cases.OneOf{O: &cases.OneOf_X{X: "foobar"}}, true, "", 0},
	{"oneof - field - valid (Y)", &cases.OneOf{O: &cases.OneOf_Y{Y: 123}}, true, "", 0},
	{"oneof - field - valid (Z)", &cases.OneOf{O: &cases.OneOf_Z{Z: &cases.TestOneOfMsg{Val: true}}}, true, "", 0},
	{"oneof - field - valid (empty)", &cases.OneOf{}, true, "", 0},
	{"oneof - field - invalid (X)", &cases.OneOf{O: &cases.OneOf_X{X: "fizzbuzz"}}, false, "invalid OneOf.X: value does not have prefix \"foo\"", 1},
	{"oneof - field - invalid (Y)", &cases.OneOf{O: &cases.OneOf_Y{Y: -1}}, false, "invalid OneOf.Y: value must be greater than 0", 1},
	{"oneof - filed - invalid (Z)", &cases.OneOf{O: &cases.OneOf_Z{Z: &cases.TestOneOfMsg{}}}, false, "invalid TestOneOfMsg.Val: value must equal true", 1},

	{"oneof - required - valid", &cases.OneOfRequired{O: &cases.OneOfRequired_X{X: ""}}, true, "", 0},
	{"oneof - require - invalid", &cases.OneOfRequired{}, false, "invalid OneOfRequired.O: value is required", 1},
}

var wrapperCases = []TestCase{
	{"wrapper - none - valid", &cases.WrapperNone{Val: &wrappers.Int32Value{Value: 123}}, true, "", 0},
	{"wrapper - none - valid (empty)", &cases.WrapperNone{Val: nil}, true, "", 0},

	{"wrapper - float - valid", &cases.WrapperFloat{Val: &wrappers.FloatValue{Value: 1}}, true, "", 0},
	{"wrapper - float - valid (empty)", &cases.WrapperFloat{Val: nil}, true, "", 0},
	{"wrapper - float - invalid", &cases.WrapperFloat{Val: &wrappers.FloatValue{Value: 0}}, false, "invalid WrapperFloat.Val: value must be greater than 0", 1},

	{"wrapper - double - valid", &cases.WrapperDouble{Val: &wrappers.DoubleValue{Value: 1}}, true, "", 0},
	{"wrapper - double - valid (empty)", &cases.WrapperDouble{Val: nil}, true, "", 0},
	{"wrapper - double - invalid", &cases.WrapperDouble{Val: &wrappers.DoubleValue{Value: 0}}, false, "invalid WrapperDouble.Val: value must be greater than 0", 1},

	{"wrapper - int64 - valid", &cases.WrapperInt64{Val: &wrappers.Int64Value{Value: 1}}, true, "", 0},
	{"wrapper - int64 - valid (empty)", &cases.WrapperInt64{Val: nil}, true, "", 0},
	{"wrapper - int64 - invalid", &cases.WrapperInt64{Val: &wrappers.Int64Value{Value: 0}}, false, "invalid WrapperInt64.Val: value must be greater than 0", 1},

	{"wrapper - int32 - valid", &cases.WrapperInt32{Val: &wrappers.Int32Value{Value: 1}}, true, "", 0},
	{"wrapper - int32 - valid (empty)", &cases.WrapperInt32{Val: nil}, true, "", 0},
	{"wrapper - int32 - invalid", &cases.WrapperInt32{Val: &wrappers.Int32Value{Value: 0}}, false, "invalid WrapperInt32.Val: value must be greater than 0", 1},

	{"wrapper - uint64 - valid", &cases.WrapperUInt64{Val: &wrappers.UInt64Value{Value: 1}}, true, "", 0},
	{"wrapper - uint64 - valid (empty)", &cases.WrapperUInt64{Val: nil}, true, "", 0},
	{"wrapper - uint64 - invalid", &cases.WrapperUInt64{Val: &wrappers.UInt64Value{Value: 0}}, false, "invalid WrapperUInt64.Val: value must be greater than 0", 1},

	{"wrapper - uint32 - valid", &cases.WrapperUInt32{Val: &wrappers.UInt32Value{Value: 1}}, true, "", 0},
	{"wrapper - uint32 - valid (empty)", &cases.WrapperUInt32{Val: nil}, true, "", 0},
	{"wrapper - uint32 - invalid", &cases.WrapperUInt32{Val: &wrappers.UInt32Value{Value: 0}}, false, "invalid WrapperUInt32.Val: value must be greater than 0", 1},

	{"wrapper - bool - valid", &cases.WrapperBool{Val: &wrappers.BoolValue{Value: true}}, true, "", 0},
	{"wrapper - bool - valid (empty)", &cases.WrapperBool{Val: nil}, true, "", 0},
	{"wrapper - bool - invalid", &cases.WrapperBool{Val: &wrappers.BoolValue{Value: false}}, false, "invalid WrapperBool.Val: value must equal true", 1},

	{"wrapper - string - valid", &cases.WrapperString{Val: &wrappers.StringValue{Value: "foobar"}}, true, "", 0},
	{"wrapper - string - valid (empty)", &cases.WrapperString{Val: nil}, true, "", 0},
	{"wrapper - string - invalid", &cases.WrapperString{Val: &wrappers.StringValue{Value: "fizzbuzz"}}, false, "invalid WrapperString.Val: value does not have suffix \"bar\"", 1},

	{"wrapper - bytes - valid", &cases.WrapperBytes{Val: &wrappers.BytesValue{Value: []byte("foo")}}, true, "", 0},
	{"wrapper - bytes - valid (empty)", &cases.WrapperBytes{Val: nil}, true, "", 0},
	{"wrapper - bytes - invalid", &cases.WrapperBytes{Val: &wrappers.BytesValue{Value: []byte("x")}}, false, "invalid WrapperBytes.Val: value length must be at least 3 bytes", 1},

	{"wrapper - required - string - valid", &cases.WrapperRequiredString{Val: &wrappers.StringValue{Value: "bar"}}, true, "", 0},
	{"wrapper - required - string - invalid", &cases.WrapperRequiredString{Val: &wrappers.StringValue{Value: "foo"}}, false, "invalid WrapperRequiredString.Val: value must equal bar", 1},
	{"wrapper - required - string - invalid (empty)", &cases.WrapperRequiredString{}, false, "invalid WrapperRequiredString.Val: value is required and must not be nil.", 1},

	{"wrapper - required - string (empty) - valid", &cases.WrapperRequiredEmptyString{Val: &wrappers.StringValue{Value: ""}}, true, "", 0},
	{"wrapper - required - string (empty) - invalid", &cases.WrapperRequiredEmptyString{Val: &wrappers.StringValue{Value: "foo"}}, false, "invalid WrapperRequiredEmptyString.Val: value must equal ", 1},
	{"wrapper - required - string (empty) - invalid (empty)", &cases.WrapperRequiredEmptyString{}, false, "invalid WrapperRequiredEmptyString.Val: value is required and must not be nil.", 1},

	{"wrapper - optional - string (uuid) - valid", &cases.WrapperOptionalUuidString{Val: &wrappers.StringValue{Value: "8b72987b-024a-43b3-b4cf-647a1f925c5d"}}, true, "", 0},
	{"wrapper - optional - string (uuid) - valid (empty)", &cases.WrapperOptionalUuidString{}, true, "", 0},
	{"wrapper - optional - string (uuid) - invalid", &cases.WrapperOptionalUuidString{Val: &wrappers.StringValue{Value: "foo"}}, false, "invalid uuid format", 1},

	{"wrapper - required - float - valid", &cases.WrapperRequiredFloat{Val: &wrappers.FloatValue{Value: 1}}, true, "", 0},
	{"wrapper - required - float - invalid", &cases.WrapperRequiredFloat{Val: &wrappers.FloatValue{Value: -5}}, false, "invalid WrapperRequiredFloat.Val: value must be greater than 0", 1},
	{"wrapper - required - float - invalid (empty)", &cases.WrapperRequiredFloat{}, false, "invalid WrapperRequiredFloat.Val: value is required and must not be nil.", 1},
}

var durationCases = []TestCase{
	{"duration - none - valid", &cases.DurationNone{Val: &duration.Duration{Seconds: 123}}, true, "", 0},

	{"duration - required - valid", &cases.DurationRequired{Val: &duration.Duration{}}, true, "", 0},
	{"duration - required - invalid", &cases.DurationRequired{Val: nil}, false, "invalid DurationRequired.Val: value is required", 1},

	{"duration - const - valid", &cases.DurationConst{Val: &duration.Duration{Seconds: 3}}, true, "", 0},
	{"duration - const - valid (empty)", &cases.DurationConst{}, true, "", 0},
	{"duration - const - invalid", &cases.DurationConst{Val: &duration.Duration{Nanos: 3}}, false, "invalid DurationConst.Val: value must equal 3s", 1},

	{"duration - in - valid", &cases.DurationIn{Val: &duration.Duration{Seconds: 1}}, true, "", 0},
	{"duration - in - valid (empty)", &cases.DurationIn{}, true, "", 0},
	{"duration - in - invalid", &cases.DurationIn{Val: &duration.Duration{}}, false, "invalid DurationIn.Val: value must be in list [seconds:1 nanos:1000]", 1},

	{"duration - not in - valid", &cases.DurationNotIn{Val: &duration.Duration{Nanos: 1}}, true, "", 0},
	{"duration - not in - valid (empty)", &cases.DurationNotIn{}, true, "", 0},
	{"duration - not in - invalid", &cases.DurationNotIn{Val: &duration.Duration{}}, false, "invalid DurationNotIn.Val: value must not be in list []", 1},

	{"duration - lt - valid", &cases.DurationLT{Val: &duration.Duration{Nanos: -1}}, true, "", 0},
	{"duration - lt - valid (empty)", &cases.DurationLT{}, true, "", 0},
	{"duration - lt - invalid (equal)", &cases.DurationLT{Val: &duration.Duration{}}, false, "invalid DurationLT.Val: value must be less than 0s", 1},
	{"duration - lt - invalid", &cases.DurationLT{Val: &duration.Duration{Seconds: 1}}, false, "invalid DurationLT.Val: value must be less than 0s", 1},

	{"duration - lte - valid", &cases.DurationLTE{Val: &duration.Duration{}}, true, "", 0},
	{"duration - lte - valid (empty)", &cases.DurationLTE{}, true, "", 0},
	{"duration - lte - valid (equal)", &cases.DurationLTE{Val: &duration.Duration{Seconds: 1}}, true, "", 0},
	{"duration - lte - invalid", &cases.DurationLTE{Val: &duration.Duration{Seconds: 1, Nanos: 1}}, false, "invalid DurationLTE.Val: value must be less than or equal to 1s", 1},

	{"duration - gt - valid", &cases.DurationGT{Val: &duration.Duration{Seconds: 1}}, true, "", 0},
	{"duration - gt - valid (empty)", &cases.DurationGT{}, true, "", 0},
	{"duration - gt - invalid (equal)", &cases.DurationGT{Val: &duration.Duration{Nanos: 1000}}, false, "invalid DurationGT.Val: value must be greater than 1µs", 1},
	{"duration - gt - invalid", &cases.DurationGT{Val: &duration.Duration{}}, false, "invalid DurationGT.Val: value must be greater than 1µs", 1},

	{"duration - gte - valid", &cases.DurationGTE{Val: &duration.Duration{Seconds: 3}}, true, "", 0},
	{"duration - gte - valid (empty)", &cases.DurationGTE{}, true, "", 0},
	{"duration - gte - valid (equal)", &cases.DurationGTE{Val: &duration.Duration{Nanos: 1000000}}, true, "", 0},
	{"duration - gte - invalid", &cases.DurationGTE{Val: &duration.Duration{Seconds: -1}}, false, "invalid DurationGTE.Val: value must be greater than or equal to 1ms", 1},

	{"duration - gt & lt - valid", &cases.DurationGTLT{Val: &duration.Duration{Nanos: 1000}}, true, "", 0},
	{"duration - gt & lt - valid (empty)", &cases.DurationGTLT{}, true, "", 0},
	{"duration - gt & lt - invalid (above)", &cases.DurationGTLT{Val: &duration.Duration{Seconds: 1000}}, false, "invalid DurationGTLT.Val: value must be inside range (0s, 1s)", 1},
	{"duration - gt & lt - invalid (below)", &cases.DurationGTLT{Val: &duration.Duration{Nanos: -1000}}, false, "invalid DurationGTLT.Val: value must be inside range (0s, 1s)", 1},
	{"duration - gt & lt - invalid (max)", &cases.DurationGTLT{Val: &duration.Duration{Seconds: 1}}, false, "invalid DurationGTLT.Val: value must be inside range (0s, 1s)", 1},
	{"duration - gt & lt - invalid (min)", &cases.DurationGTLT{Val: &duration.Duration{}}, false, "invalid DurationGTLT.Val: value must be inside range (0s, 1s)", 1},

	{"duration - exclusive gt & lt - valid (empty)", &cases.DurationExLTGT{}, true, "", 0},
	{"duration - exclusive gt & lt - valid (above)", &cases.DurationExLTGT{Val: &duration.Duration{Seconds: 2}}, true, "", 0},
	{"duration - exclusive gt & lt - valid (below)", &cases.DurationExLTGT{Val: &duration.Duration{Nanos: -1}}, true, "", 0},
	{"duration - exclusive gt & lt - invalid", &cases.DurationExLTGT{Val: &duration.Duration{Nanos: 1000}}, false, "invalid DurationExLTGT.Val: value must be outside range [0s, 1s]", 1},
	{"duration - exclusive gt & lt - invalid (max)", &cases.DurationExLTGT{Val: &duration.Duration{Seconds: 1}}, false, "invalid DurationExLTGT.Val: value must be outside range [0s, 1s]", 1},
	{"duration - exclusive gt & lt - invalid (min)", &cases.DurationExLTGT{Val: &duration.Duration{}}, false, "invalid DurationExLTGT.Val: value must be outside range [0s, 1s]", 1},

	{"duration - gte & lte - valid", &cases.DurationGTELTE{Val: &duration.Duration{Seconds: 60, Nanos: 1}}, true, "", 0},
	{"duration - gte & lte - valid (empty)", &cases.DurationGTELTE{}, true, "", 0},
	{"duration - gte & lte - valid (max)", &cases.DurationGTELTE{Val: &duration.Duration{Seconds: 3600}}, true, "", 0},
	{"duration - gte & lte - valid (min)", &cases.DurationGTELTE{Val: &duration.Duration{Seconds: 60}}, true, "", 0},
	{"duration - gte & lte - invalid (above)", &cases.DurationGTELTE{Val: &duration.Duration{Seconds: 3600, Nanos: 1}}, false, "invalid DurationGTELTE.Val: value must be inside range [1m0s, 1h0m0s]", 1},
	{"duration - gte & lte - invalid (below)", &cases.DurationGTELTE{Val: &duration.Duration{Seconds: 59}}, false, "invalid DurationGTELTE.Val: value must be inside range [1m0s, 1h0m0s]", 1},

	{"duration - gte & lte - valid (empty)", &cases.DurationExGTELTE{}, true, "", 0},
	{"duration - exclusive gte & lte - valid (above)", &cases.DurationExGTELTE{Val: &duration.Duration{Seconds: 3601}}, true, "", 0},
	{"duration - exclusive gte & lte - valid (below)", &cases.DurationExGTELTE{Val: &duration.Duration{}}, true, "", 0},
	{"duration - exclusive gte & lte - valid (max)", &cases.DurationExGTELTE{Val: &duration.Duration{Seconds: 3600}}, true, "", 0},
	{"duration - exclusive gte & lte - valid (min)", &cases.DurationExGTELTE{Val: &duration.Duration{Seconds: 60}}, true, "", 0},
	{"duration - exclusive gte & lte - invalid", &cases.DurationExGTELTE{Val: &duration.Duration{Seconds: 61}}, false, "invalid DurationExGTELTE.Val: value must be outside range (1m0s, 1h0m0s)", 1},
	{"duration - fields with other fields - invalid other field", &cases.DurationFieldWithOtherFields{DurationVal: nil, IntVal: 12}, false, "invalid DurationFieldWithOtherFields.IntVal: value must be greater than 16", 1},
}

var timestampCases = []TestCase{
	{"timestamp - none - valid", &cases.TimestampNone{Val: &timestamp.Timestamp{Seconds: 123}}, true, "", 0},

	{"timestamp - required - valid", &cases.TimestampRequired{Val: &timestamp.Timestamp{}}, true, "", 0},
	{"timestamp - required - invalid", &cases.TimestampRequired{Val: nil}, false, "invalid TimestampRequired.Val: value is required", 1},

	{"timestamp - const - valid", &cases.TimestampConst{Val: &timestamp.Timestamp{Seconds: 3}}, true, "", 0},
	{"timestamp - const - valid (empty)", &cases.TimestampConst{}, true, "", 0},
	{"timestamp - const - invalid", &cases.TimestampConst{Val: &timestamp.Timestamp{Nanos: 3}}, false, "invalid TimestampConst.Val: value must equal 1970-01-01 00:00:03 +0000 UTC", 1},

	{"timestamp - lt - valid", &cases.TimestampLT{Val: &timestamp.Timestamp{Seconds: -1}}, true, "", 0},
	{"timestamp - lt - valid (empty)", &cases.TimestampLT{}, true, "", 0},
	{"timestamp - lt - invalid (equal)", &cases.TimestampLT{Val: &timestamp.Timestamp{}}, false, "invalid TimestampLT.Val: value must be less than 1970-01-01 00:00:00 +0000 UTC", 1},
	{"timestamp - lt - invalid", &cases.TimestampLT{Val: &timestamp.Timestamp{Seconds: 1}}, false, "invalid TimestampLT.Val: value must be less than 1970-01-01 00:00:00 +0000 UTC", 1},

	{"timestamp - lte - valid", &cases.TimestampLTE{Val: &timestamp.Timestamp{}}, true, "", 0},
	{"timestamp - lte - valid (empty)", &cases.TimestampLTE{}, true, "", 0},
	{"timestamp - lte - valid (equal)", &cases.TimestampLTE{Val: &timestamp.Timestamp{Seconds: 1}}, true, "", 0},
	{"timestamp - lte - invalid", &cases.TimestampLTE{Val: &timestamp.Timestamp{Seconds: 1, Nanos: 1}}, false, "invalid TimestampLTE.Val: value must be less than or equal to 1970-01-01 00:00:01 +0000 UTC", 1},

	{"timestamp - gt - valid", &cases.TimestampGT{Val: &timestamp.Timestamp{Seconds: 1}}, true, "", 0},
	{"timestamp - gt - valid (empty)", &cases.TimestampGT{}, true, "", 0},
	{"timestamp - gt - invalid (equal)", &cases.TimestampGT{Val: &timestamp.Timestamp{Nanos: 1000}}, false, "invalid TimestampGT.Val: value must be greater than 1970-01-01 00:00:00.000001 +0000 UTC", 1},
	{"timestamp - gt - invalid", &cases.TimestampGT{Val: &timestamp.Timestamp{}}, false, "invalid TimestampGT.Val: value must be greater than 1970-01-01 00:00:00.000001 +0000 UTC", 1},

	{"timestamp - gte - valid", &cases.TimestampGTE{Val: &timestamp.Timestamp{Seconds: 3}}, true, "", 0},
	{"timestamp - gte - valid (empty)", &cases.TimestampGTE{}, true, "", 0},
	{"timestamp - gte - valid (equal)", &cases.TimestampGTE{Val: &timestamp.Timestamp{Nanos: 1000000}}, true, "", 0},
	{"timestamp - gte - invalid", &cases.TimestampGTE{Val: &timestamp.Timestamp{Seconds: -1}}, false, "invalid TimestampGTE.Val: value must be greater than or equal to 1970-01-01 00:00:00.001 +0000 UTC", 1},

	{"timestamp - gt & lt - valid", &cases.TimestampGTLT{Val: &timestamp.Timestamp{Nanos: 1000}}, true, "", 0},
	{"timestamp - gt & lt - valid (empty)", &cases.TimestampGTLT{}, true, "", 0},
	{"timestamp - gt & lt - invalid (above)", &cases.TimestampGTLT{Val: &timestamp.Timestamp{Seconds: 1000}}, false, "invalid TimestampGTLT.Val: value must be inside range (1970-01-01 00:00:00 +0000 UTC, 1970-01-01 00:00:01 +0000 UTC)", 1},
	{"timestamp - gt & lt - invalid (below)", &cases.TimestampGTLT{Val: &timestamp.Timestamp{Seconds: -1000}}, false, "invalid TimestampGTLT.Val: value must be inside range (1970-01-01 00:00:00 +0000 UTC, 1970-01-01 00:00:01 +0000 UTC)", 1},
	{"timestamp - gt & lt - invalid (max)", &cases.TimestampGTLT{Val: &timestamp.Timestamp{Seconds: 1}}, false, "invalid TimestampGTLT.Val: value must be inside range (1970-01-01 00:00:00 +0000 UTC, 1970-01-01 00:00:01 +0000 UTC)", 1},
	{"timestamp - gt & lt - invalid (min)", &cases.TimestampGTLT{Val: &timestamp.Timestamp{}}, false, "invalid TimestampGTLT.Val: value must be inside range (1970-01-01 00:00:00 +0000 UTC, 1970-01-01 00:00:01 +0000 UTC)", 1},

	{"timestamp - exclusive gt & lt - valid (empty)", &cases.TimestampExLTGT{}, true, "", 0},
	{"timestamp - exclusive gt & lt - valid (above)", &cases.TimestampExLTGT{Val: &timestamp.Timestamp{Seconds: 2}}, true, "", 0},
	{"timestamp - exclusive gt & lt - valid (below)", &cases.TimestampExLTGT{Val: &timestamp.Timestamp{Seconds: -1}}, true, "", 0},
	{"timestamp - exclusive gt & lt - invalid", &cases.TimestampExLTGT{Val: &timestamp.Timestamp{Nanos: 1000}}, false, "invalid TimestampExLTGT.Val: value must be outside range [1970-01-01 00:00:00 +0000 UTC, 1970-01-01 00:00:01 +0000 UTC]", 1},
	{"timestamp - exclusive gt & lt - invalid (max)", &cases.TimestampExLTGT{Val: &timestamp.Timestamp{Seconds: 1}}, false, "invalid TimestampExLTGT.Val: value must be outside range [1970-01-01 00:00:00 +0000 UTC, 1970-01-01 00:00:01 +0000 UTC]", 1},
	{"timestamp - exclusive gt & lt - invalid (min)", &cases.TimestampExLTGT{Val: &timestamp.Timestamp{}}, false, "invalid TimestampExLTGT.Val: value must be outside range [1970-01-01 00:00:00 +0000 UTC, 1970-01-01 00:00:01 +0000 UTC]", 1},

	{"timestamp - gte & lte - valid", &cases.TimestampGTELTE{Val: &timestamp.Timestamp{Seconds: 60, Nanos: 1}}, true, "", 0},
	{"timestamp - gte & lte - valid (empty)", &cases.TimestampGTELTE{}, true, "", 0},
	{"timestamp - gte & lte - valid (max)", &cases.TimestampGTELTE{Val: &timestamp.Timestamp{Seconds: 3600}}, true, "", 0},
	{"timestamp - gte & lte - valid (min)", &cases.TimestampGTELTE{Val: &timestamp.Timestamp{Seconds: 60}}, true, "", 0},
	{"timestamp - gte & lte - invalid (above)", &cases.TimestampGTELTE{Val: &timestamp.Timestamp{Seconds: 3600, Nanos: 1}}, false, "invalid TimestampGTELTE.Val: value must be inside range [1970-01-01 00:01:00 +0000 UTC, 1970-01-01 01:00:00 +0000 UTC]", 1},
	{"timestamp - gte & lte - invalid (below)", &cases.TimestampGTELTE{Val: &timestamp.Timestamp{Seconds: 59}}, false, "invalid TimestampGTELTE.Val: value must be inside range [1970-01-01 00:01:00 +0000 UTC, 1970-01-01 01:00:00 +0000 UTC]", 1},

	{"timestamp - gte & lte - valid (empty)", &cases.TimestampExGTELTE{}, true, "", 0},
	{"timestamp - exclusive gte & lte - valid (above)", &cases.TimestampExGTELTE{Val: &timestamp.Timestamp{Seconds: 3601}}, true, "", 0},
	{"timestamp - exclusive gte & lte - valid (below)", &cases.TimestampExGTELTE{Val: &timestamp.Timestamp{}}, true, "", 0},
	{"timestamp - exclusive gte & lte - valid (max)", &cases.TimestampExGTELTE{Val: &timestamp.Timestamp{Seconds: 3600}}, true, "", 0},
	{"timestamp - exclusive gte & lte - valid (min)", &cases.TimestampExGTELTE{Val: &timestamp.Timestamp{Seconds: 60}}, true, "", 0},
	{"timestamp - exclusive gte & lte - invalid", &cases.TimestampExGTELTE{Val: &timestamp.Timestamp{Seconds: 61}}, false, "invalid TimestampExGTELTE.Val: value must be outside range (1970-01-01 00:01:00 +0000 UTC, 1970-01-01 01:00:00 +0000 UTC)", 1},

	{"timestamp - lt now - valid", &cases.TimestampLTNow{Val: &timestamp.Timestamp{}}, true, "", 0},
	{"timestamp - lt now - valid (empty)", &cases.TimestampLTNow{}, true, "", 0},
	{"timestamp - lt now - invalid", &cases.TimestampLTNow{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 7200}}, false, "invalid TimestampLTNow.Val: value must be less than now", 1},

	{"timestamp - gt now - valid", &cases.TimestampGTNow{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 7200}}, true, "", 0},
	{"timestamp - gt now - valid (empty)", &cases.TimestampGTNow{}, true, "", 0},
	{"timestamp - gt now - invalid", &cases.TimestampGTNow{Val: &timestamp.Timestamp{}}, false, "invalid TimestampGTNow.Val: value must be greater than now", 1},

	{"timestamp - within - valid", &cases.TimestampWithin{Val: timestamp.Now()}, true, "", 0},
	{"timestamp - within - valid (empty)", &cases.TimestampWithin{}, true, "", 0},
	{"timestamp - within - invalid (below)", &cases.TimestampWithin{Val: &timestamp.Timestamp{}}, false, "invalid TimestampWithin.Val: value must be within 1h0m0s of now", 1},
	{"timestamp - within - invalid (above)", &cases.TimestampWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 7200}}, false, "invalid TimestampWithin.Val: value must be within 1h0m0s of now", 1},

	{"timestamp - lt now within - valid", &cases.TimestampLTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() - 1800}}, true, "", 0},
	{"timestamp - lt now within - valid (empty)", &cases.TimestampLTNowWithin{}, true, "", 0},
	{"timestamp - lt now within - invalid (lt)", &cases.TimestampLTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 1800}}, false, "invalid TimestampLTNowWithin.Val: value must be less than now within 1h0m0s", 1},
	{"timestamp - lt now within - invalid (within)", &cases.TimestampLTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() - 7200}}, false, "invalid TimestampLTNowWithin.Val: value must be less than now within 1h0m0s", 1},

	{"timestamp - gt now within - valid", &cases.TimestampGTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 1800}}, true, "", 0},
	{"timestamp - gt now within - valid (empty)", &cases.TimestampGTNowWithin{}, true, "", 0},
	{"timestamp - gt now within - invalid (gt)", &cases.TimestampGTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() - 1800}}, false, "invalid TimestampGTNowWithin.Val: value must be greater than now within 1h0m0s", 1},
	{"timestamp - gt now within - invalid (within)", &cases.TimestampGTNowWithin{Val: &timestamp.Timestamp{Seconds: time.Now().Unix() + 7200}}, false, "invalid TimestampGTNowWithin.Val: value must be greater than now within 1h0m0s", 1},
}

var anyCases = []TestCase{
	{"any - none - valid", &cases.AnyNone{Val: &any.Any{}}, true, "", 0},

	{"any - required - valid", &cases.AnyRequired{Val: &any.Any{}}, true, "", 0},
	{"any - required - invalid", &cases.AnyRequired{Val: nil}, false, "invalid AnyRequired.Val: value is required", 1},

	{"any - in - valid", &cases.AnyIn{Val: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}}, true, "", 0},
	{"any - in - valid (empty)", &cases.AnyIn{}, true, "", 0},
	{"any - in - invalid", &cases.AnyIn{Val: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Timestamp"}}, false, "invalid AnyIn.Val: type URL must be in list [type.googleapis.com/google.protobuf.Duration]", 1},

	{"any - not in - valid", &cases.AnyNotIn{Val: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}}, true, "", 0},
	{"any - not in - valid (empty)", &cases.AnyNotIn{}, true, "", 0},
	{"any - not in - invalid", &cases.AnyNotIn{Val: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Timestamp"}}, false, "invalid AnyNotIn.Val: type URL must not be in list [type.googleapis.com/google.protobuf.Timestamp]", 1},
}

var kitchenSink = []TestCase{
	{"kitchensink - field - valid", &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Const: "abcd", IntConst: 5, BoolConst: false, FloatVal: &wrappers.FloatValue{Value: 1}, DurVal: &duration.Duration{Seconds: 3}, TsVal: &timestamp.Timestamp{Seconds: 17}, FloatConst: 7, DoubleIn: 123, EnumConst: cases.ComplexTestEnum_ComplexTWO, AnyVal: &any.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}, RepTsVal: []*timestamp.Timestamp{{Seconds: 3}}, MapVal: map[int32]string{-1: "a", -2: "b"}, BytesVal: []byte("\x00\x99"), O: &cases.ComplexTestMsg_X{X: "foobar"}}}, true, "", 0},
	{"kitchensink - valid (unset)", &cases.KitchenSinkMessage{}, true, "", 0},
	{"kitchensink - field - invalid", &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{}}, false, "invalid ComplexTestMsg.Const: value must equal abcd", 1},
	{"kitchensink - field - embedded - invalid", &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Another: &cases.ComplexTestMsg{}}}, false, "invalid ComplexTestMsg.Const: value must equal abcd", 1},
	{"kitchensink - field - invalid (transitive)", &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Const: "abcd", BoolConst: true, Nested: &cases.ComplexTestMsg{}}}, false, "invalid ComplexTestMsg.Const: value must equal abcd", 1},
}

var nestedCases = []TestCase{
	{"nested wkt uuid - field - valid", &cases.WktLevelOne{Two: &cases.WktLevelOne_WktLevelTwo{Three: &cases.WktLevelOne_WktLevelTwo_WktLevelThree{Uuid: "f81d16ef-40e2-40c6-bebc-89aaf5292f9a"}}}, true, "", 0},
	{"nested wkt uuid - field - invalid", &cases.WktLevelOne{Two: &cases.WktLevelOne_WktLevelTwo{Three: &cases.WktLevelOne_WktLevelTwo_WktLevelThree{Uuid: "not-a-valid-uuid"}}}, false, "invalid uuid format", 1},
}
