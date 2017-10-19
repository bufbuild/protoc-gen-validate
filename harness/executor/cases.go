package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/lyft/protoc-gen-validate/harness/cases/go"
)

type TestCase struct {
	Name    string
	Message proto.Message
	Valid   bool
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
	}

	for _, set := range sets {
		TestCases = append(TestCases, set...)
	}
}

var floatCases = []TestCase{
	{"float - none - valid", &cases.FloatNone{Val: -1.23456}, true},

	{"float - const - valid", &cases.FloatConst{Val: 1.23}, true},
	{"float - const - invalid", &cases.FloatConst{Val: 4.56}, false},

	{"float - in - valid", &cases.FloatIn{Val: 7.89}, true},
	{"float - in - invalid", &cases.FloatIn{Val: 10.11}, false},

	{"float - not in - valid", &cases.FloatNotIn{Val: 1}, true},
	{"float - not in - invalid", &cases.FloatNotIn{Val: 0}, false},

	{"float - lt - valid", &cases.FloatLT{Val: -1}, true},
	{"float - lt - invalid (equal)", &cases.FloatLT{Val: 0}, false},
	{"float - lt - invalid", &cases.FloatLT{Val: 1}, false},

	{"float - lte - valid", &cases.FloatLTE{Val: 63}, true},
	{"float - lte - valid (equal)", &cases.FloatLTE{Val: 64}, true},
	{"float - lte - invalid", &cases.FloatLTE{Val: 65}, false},

	{"float - gt - valid", &cases.FloatGT{Val: 17}, true},
	{"float - gt - invalid (equal)", &cases.FloatGT{Val: 16}, false},
	{"float - gt - invalid", &cases.FloatGT{Val: 15}, false},

	{"float - gte - valid", &cases.FloatGTE{Val: 9}, true},
	{"float - gte - valid (equal)", &cases.FloatGTE{Val: 8}, true},
	{"float - gte - invalid", &cases.FloatGTE{Val: 7}, false},

	{"float - gt & lt - valid", &cases.FloatGTLT{Val: 5}, true},
	{"float - gt & lt - invalid (above)", &cases.FloatGTLT{Val: 11}, false},
	{"float - gt & lt - invalid (below)", &cases.FloatGTLT{Val: -1}, false},
	{"float - gt & lt - invalid (max)", &cases.FloatGTLT{Val: 10}, false},
	{"float - gt & lt - invalid (min)", &cases.FloatGTLT{Val: 0}, false},

	{"float - exclusive gt & lt - valid (above)", &cases.FloatExLTGT{Val: 11}, true},
	{"float - exclusive gt & lt - valid (below)", &cases.FloatExLTGT{Val: -1}, true},
	{"float - exclusive gt & lt - invalid", &cases.FloatExLTGT{Val: 5}, false},
	{"float - exclusive gt & lt - invalid (max)", &cases.FloatExLTGT{Val: 10}, false},
	{"float - exclusive gt & lt - invalid (min)", &cases.FloatExLTGT{Val: 0}, false},

	{"float - gte & lte - valid", &cases.FloatGTELTE{Val: 200}, true},
	{"float - gte & lte - valid (max)", &cases.FloatGTELTE{Val: 256}, true},
	{"float - gte & lte - valid (min)", &cases.FloatGTELTE{Val: 128}, true},
	{"float - gte & lte - invalid (above)", &cases.FloatGTELTE{Val: 300}, false},
	{"float - gte & lte - invalid (below)", &cases.FloatGTELTE{Val: 100}, false},

	{"float - exclusive gte & lte - valid (above)", &cases.FloatExGTELTE{Val: 300}, true},
	{"float - exclusive gte & lte - valid (below)", &cases.FloatExGTELTE{Val: 100}, true},
	{"float - exclusive gte & lte - valid (max)", &cases.FloatExGTELTE{Val: 256}, true},
	{"float - exclusive gte & lte - valid (min)", &cases.FloatExGTELTE{Val: 128}, true},
	{"float - exclusive gte & lte - invalid", &cases.FloatExGTELTE{Val: 200}, false},
}

var doubleCases = []TestCase{
	{"double - none - valid", &cases.DoubleNone{Val: -1.23456}, true},

	{"double - const - valid", &cases.DoubleConst{Val: 1.23}, true},
	{"double - const - invalid", &cases.DoubleConst{Val: 4.56}, false},

	{"double - in - valid", &cases.DoubleIn{Val: 7.89}, true},
	{"double - in - invalid", &cases.DoubleIn{Val: 10.11}, false},

	{"double - not in - valid", &cases.DoubleNotIn{Val: 1}, true},
	{"double - not in - invalid", &cases.DoubleNotIn{Val: 0}, false},

	{"double - lt - valid", &cases.DoubleLT{Val: -1}, true},
	{"double - lt - invalid (equal)", &cases.DoubleLT{Val: 0}, false},
	{"double - lt - invalid", &cases.DoubleLT{Val: 1}, false},

	{"double - lte - valid", &cases.DoubleLTE{Val: 63}, true},
	{"double - lte - valid (equal)", &cases.DoubleLTE{Val: 64}, true},
	{"double - lte - invalid", &cases.DoubleLTE{Val: 65}, false},

	{"double - gt - valid", &cases.DoubleGT{Val: 17}, true},
	{"double - gt - invalid (equal)", &cases.DoubleGT{Val: 16}, false},
	{"double - gt - invalid", &cases.DoubleGT{Val: 15}, false},

	{"double - gte - valid", &cases.DoubleGTE{Val: 9}, true},
	{"double - gte - valid (equal)", &cases.DoubleGTE{Val: 8}, true},
	{"double - gte - invalid", &cases.DoubleGTE{Val: 7}, false},

	{"double - gt & lt - valid", &cases.DoubleGTLT{Val: 5}, true},
	{"double - gt & lt - invalid (above)", &cases.DoubleGTLT{Val: 11}, false},
	{"double - gt & lt - invalid (below)", &cases.DoubleGTLT{Val: -1}, false},
	{"double - gt & lt - invalid (max)", &cases.DoubleGTLT{Val: 10}, false},
	{"double - gt & lt - invalid (min)", &cases.DoubleGTLT{Val: 0}, false},

	{"double - exclusive gt & lt - valid (above)", &cases.DoubleExLTGT{Val: 11}, true},
	{"double - exclusive gt & lt - valid (below)", &cases.DoubleExLTGT{Val: -1}, true},
	{"double - exclusive gt & lt - invalid", &cases.DoubleExLTGT{Val: 5}, false},
	{"double - exclusive gt & lt - invalid (max)", &cases.DoubleExLTGT{Val: 10}, false},
	{"double - exclusive gt & lt - invalid (min)", &cases.DoubleExLTGT{Val: 0}, false},

	{"double - gte & lte - valid", &cases.DoubleGTELTE{Val: 200}, true},
	{"double - gte & lte - valid (max)", &cases.DoubleGTELTE{Val: 256}, true},
	{"double - gte & lte - valid (min)", &cases.DoubleGTELTE{Val: 128}, true},
	{"double - gte & lte - invalid (above)", &cases.DoubleGTELTE{Val: 300}, false},
	{"double - gte & lte - invalid (below)", &cases.DoubleGTELTE{Val: 100}, false},

	{"double - exclusive gte & lte - valid (above)", &cases.DoubleExGTELTE{Val: 300}, true},
	{"double - exclusive gte & lte - valid (below)", &cases.DoubleExGTELTE{Val: 100}, true},
	{"double - exclusive gte & lte - valid (max)", &cases.DoubleExGTELTE{Val: 256}, true},
	{"double - exclusive gte & lte - valid (min)", &cases.DoubleExGTELTE{Val: 128}, true},
	{"double - exclusive gte & lte - invalid", &cases.DoubleExGTELTE{Val: 200}, false},
}

var int32Cases = []TestCase{
	{"int32 - none - valid", &cases.Int32None{Val: 123}, true},

	{"int32 - const - valid", &cases.Int32Const{Val: 1}, true},
	{"int32 - const - invalid", &cases.Int32Const{Val: 2}, false},

	{"int32 - in - valid", &cases.Int32In{Val: 3}, true},
	{"int32 - in - invalid", &cases.Int32In{Val: 5}, false},

	{"int32 - not in - valid", &cases.Int32NotIn{Val: 1}, true},
	{"int32 - not in - invalid", &cases.Int32NotIn{Val: 0}, false},

	{"int32 - lt - valid", &cases.Int32LT{Val: -1}, true},
	{"int32 - lt - invalid (equal)", &cases.Int32LT{Val: 0}, false},
	{"int32 - lt - invalid", &cases.Int32LT{Val: 1}, false},

	{"int32 - lte - valid", &cases.Int32LTE{Val: 63}, true},
	{"int32 - lte - valid (equal)", &cases.Int32LTE{Val: 64}, true},
	{"int32 - lte - invalid", &cases.Int32LTE{Val: 65}, false},

	{"int32 - gt - valid", &cases.Int32GT{Val: 17}, true},
	{"int32 - gt - invalid (equal)", &cases.Int32GT{Val: 16}, false},
	{"int32 - gt - invalid", &cases.Int32GT{Val: 15}, false},

	{"int32 - gte - valid", &cases.Int32GTE{Val: 9}, true},
	{"int32 - gte - valid (equal)", &cases.Int32GTE{Val: 8}, true},
	{"int32 - gte - invalid", &cases.Int32GTE{Val: 7}, false},

	{"int32 - gt & lt - valid", &cases.Int32GTLT{Val: 5}, true},
	{"int32 - gt & lt - invalid (above)", &cases.Int32GTLT{Val: 11}, false},
	{"int32 - gt & lt - invalid (below)", &cases.Int32GTLT{Val: -1}, false},
	{"int32 - gt & lt - invalid (max)", &cases.Int32GTLT{Val: 10}, false},
	{"int32 - gt & lt - invalid (min)", &cases.Int32GTLT{Val: 0}, false},

	{"int32 - exclusive gt & lt - valid (above)", &cases.Int32ExLTGT{Val: 11}, true},
	{"int32 - exclusive gt & lt - valid (below)", &cases.Int32ExLTGT{Val: -1}, true},
	{"int32 - exclusive gt & lt - invalid", &cases.Int32ExLTGT{Val: 5}, false},
	{"int32 - exclusive gt & lt - invalid (max)", &cases.Int32ExLTGT{Val: 10}, false},
	{"int32 - exclusive gt & lt - invalid (min)", &cases.Int32ExLTGT{Val: 0}, false},

	{"int32 - gte & lte - valid", &cases.Int32GTELTE{Val: 200}, true},
	{"int32 - gte & lte - valid (max)", &cases.Int32GTELTE{Val: 256}, true},
	{"int32 - gte & lte - valid (min)", &cases.Int32GTELTE{Val: 128}, true},
	{"int32 - gte & lte - invalid (above)", &cases.Int32GTELTE{Val: 300}, false},
	{"int32 - gte & lte - invalid (below)", &cases.Int32GTELTE{Val: 100}, false},

	{"int32 - exclusive gte & lte - valid (above)", &cases.Int32ExGTELTE{Val: 300}, true},
	{"int32 - exclusive gte & lte - valid (below)", &cases.Int32ExGTELTE{Val: 100}, true},
	{"int32 - exclusive gte & lte - valid (max)", &cases.Int32ExGTELTE{Val: 256}, true},
	{"int32 - exclusive gte & lte - valid (min)", &cases.Int32ExGTELTE{Val: 128}, true},
	{"int32 - exclusive gte & lte - invalid", &cases.Int32ExGTELTE{Val: 200}, false},
}

var int64Cases = []TestCase{
	{"int64 - none - valid", &cases.Int64None{Val: 123}, true},

	{"int64 - const - valid", &cases.Int64Const{Val: 1}, true},
	{"int64 - const - invalid", &cases.Int64Const{Val: 2}, false},

	{"int64 - in - valid", &cases.Int64In{Val: 3}, true},
	{"int64 - in - invalid", &cases.Int64In{Val: 5}, false},

	{"int64 - not in - valid", &cases.Int64NotIn{Val: 1}, true},
	{"int64 - not in - invalid", &cases.Int64NotIn{Val: 0}, false},

	{"int64 - lt - valid", &cases.Int64LT{Val: -1}, true},
	{"int64 - lt - invalid (equal)", &cases.Int64LT{Val: 0}, false},
	{"int64 - lt - invalid", &cases.Int64LT{Val: 1}, false},

	{"int64 - lte - valid", &cases.Int64LTE{Val: 63}, true},
	{"int64 - lte - valid (equal)", &cases.Int64LTE{Val: 64}, true},
	{"int64 - lte - invalid", &cases.Int64LTE{Val: 65}, false},

	{"int64 - gt - valid", &cases.Int64GT{Val: 17}, true},
	{"int64 - gt - invalid (equal)", &cases.Int64GT{Val: 16}, false},
	{"int64 - gt - invalid", &cases.Int64GT{Val: 15}, false},

	{"int64 - gte - valid", &cases.Int64GTE{Val: 9}, true},
	{"int64 - gte - valid (equal)", &cases.Int64GTE{Val: 8}, true},
	{"int64 - gte - invalid", &cases.Int64GTE{Val: 7}, false},

	{"int64 - gt & lt - valid", &cases.Int64GTLT{Val: 5}, true},
	{"int64 - gt & lt - invalid (above)", &cases.Int64GTLT{Val: 11}, false},
	{"int64 - gt & lt - invalid (below)", &cases.Int64GTLT{Val: -1}, false},
	{"int64 - gt & lt - invalid (max)", &cases.Int64GTLT{Val: 10}, false},
	{"int64 - gt & lt - invalid (min)", &cases.Int64GTLT{Val: 0}, false},

	{"int64 - exclusive gt & lt - valid (above)", &cases.Int64ExLTGT{Val: 11}, true},
	{"int64 - exclusive gt & lt - valid (below)", &cases.Int64ExLTGT{Val: -1}, true},
	{"int64 - exclusive gt & lt - invalid", &cases.Int64ExLTGT{Val: 5}, false},
	{"int64 - exclusive gt & lt - invalid (max)", &cases.Int64ExLTGT{Val: 10}, false},
	{"int64 - exclusive gt & lt - invalid (min)", &cases.Int64ExLTGT{Val: 0}, false},

	{"int64 - gte & lte - valid", &cases.Int64GTELTE{Val: 200}, true},
	{"int64 - gte & lte - valid (max)", &cases.Int64GTELTE{Val: 256}, true},
	{"int64 - gte & lte - valid (min)", &cases.Int64GTELTE{Val: 128}, true},
	{"int64 - gte & lte - invalid (above)", &cases.Int64GTELTE{Val: 300}, false},
	{"int64 - gte & lte - invalid (below)", &cases.Int64GTELTE{Val: 100}, false},

	{"int64 - exclusive gte & lte - valid (above)", &cases.Int64ExGTELTE{Val: 300}, true},
	{"int64 - exclusive gte & lte - valid (below)", &cases.Int64ExGTELTE{Val: 100}, true},
	{"int64 - exclusive gte & lte - valid (max)", &cases.Int64ExGTELTE{Val: 256}, true},
	{"int64 - exclusive gte & lte - valid (min)", &cases.Int64ExGTELTE{Val: 128}, true},
	{"int64 - exclusive gte & lte - invalid", &cases.Int64ExGTELTE{Val: 200}, false},
}

var uint32Cases = []TestCase{
	{"uint32 - none - valid", &cases.UInt32None{Val: 123}, true},

	{"uint32 - const - valid", &cases.UInt32Const{Val: 1}, true},
	{"uint32 - const - invalid", &cases.UInt32Const{Val: 2}, false},

	{"uint32 - in - valid", &cases.UInt32In{Val: 3}, true},
	{"uint32 - in - invalid", &cases.UInt32In{Val: 5}, false},

	{"uint32 - not in - valid", &cases.UInt32NotIn{Val: 1}, true},
	{"uint32 - not in - invalid", &cases.UInt32NotIn{Val: 0}, false},

	{"uint32 - lt - valid", &cases.UInt32LT{Val: 4}, true},
	{"uint32 - lt - invalid (equal)", &cases.UInt32LT{Val: 5}, false},
	{"uint32 - lt - invalid", &cases.UInt32LT{Val: 6}, false},

	{"uint32 - lte - valid", &cases.UInt32LTE{Val: 63}, true},
	{"uint32 - lte - valid (equal)", &cases.UInt32LTE{Val: 64}, true},
	{"uint32 - lte - invalid", &cases.UInt32LTE{Val: 65}, false},

	{"uint32 - gt - valid", &cases.UInt32GT{Val: 17}, true},
	{"uint32 - gt - invalid (equal)", &cases.UInt32GT{Val: 16}, false},
	{"uint32 - gt - invalid", &cases.UInt32GT{Val: 15}, false},

	{"uint32 - gte - valid", &cases.UInt32GTE{Val: 9}, true},
	{"uint32 - gte - valid (equal)", &cases.UInt32GTE{Val: 8}, true},
	{"uint32 - gte - invalid", &cases.UInt32GTE{Val: 7}, false},

	{"uint32 - gt & lt - valid", &cases.UInt32GTLT{Val: 7}, true},
	{"uint32 - gt & lt - invalid (above)", &cases.UInt32GTLT{Val: 11}, false},
	{"uint32 - gt & lt - invalid (below)", &cases.UInt32GTLT{Val: 1}, false},
	{"uint32 - gt & lt - invalid (max)", &cases.UInt32GTLT{Val: 10}, false},
	{"uint32 - gt & lt - invalid (min)", &cases.UInt32GTLT{Val: 5}, false},

	{"uint32 - exclusive gt & lt - valid (above)", &cases.UInt32ExLTGT{Val: 11}, true},
	{"uint32 - exclusive gt & lt - valid (below)", &cases.UInt32ExLTGT{Val: 4}, true},
	{"uint32 - exclusive gt & lt - invalid", &cases.UInt32ExLTGT{Val: 7}, false},
	{"uint32 - exclusive gt & lt - invalid (max)", &cases.UInt32ExLTGT{Val: 10}, false},
	{"uint32 - exclusive gt & lt - invalid (min)", &cases.UInt32ExLTGT{Val: 5}, false},

	{"uint32 - gte & lte - valid", &cases.UInt32GTELTE{Val: 200}, true},
	{"uint32 - gte & lte - valid (max)", &cases.UInt32GTELTE{Val: 256}, true},
	{"uint32 - gte & lte - valid (min)", &cases.UInt32GTELTE{Val: 128}, true},
	{"uint32 - gte & lte - invalid (above)", &cases.UInt32GTELTE{Val: 300}, false},
	{"uint32 - gte & lte - invalid (below)", &cases.UInt32GTELTE{Val: 100}, false},

	{"uint32 - exclusive gte & lte - valid (above)", &cases.UInt32ExGTELTE{Val: 300}, true},
	{"uint32 - exclusive gte & lte - valid (below)", &cases.UInt32ExGTELTE{Val: 100}, true},
	{"uint32 - exclusive gte & lte - valid (max)", &cases.UInt32ExGTELTE{Val: 256}, true},
	{"uint32 - exclusive gte & lte - valid (min)", &cases.UInt32ExGTELTE{Val: 128}, true},
	{"uint32 - exclusive gte & lte - invalid", &cases.UInt32ExGTELTE{Val: 200}, false},
}

var uint64Cases = []TestCase{
	{"uint64 - none - valid", &cases.UInt64None{Val: 123}, true},

	{"uint64 - const - valid", &cases.UInt64Const{Val: 1}, true},
	{"uint64 - const - invalid", &cases.UInt64Const{Val: 2}, false},

	{"uint64 - in - valid", &cases.UInt64In{Val: 3}, true},
	{"uint64 - in - invalid", &cases.UInt64In{Val: 5}, false},

	{"uint64 - not in - valid", &cases.UInt64NotIn{Val: 1}, true},
	{"uint64 - not in - invalid", &cases.UInt64NotIn{Val: 0}, false},

	{"uint64 - lt - valid", &cases.UInt64LT{Val: 4}, true},
	{"uint64 - lt - invalid (equal)", &cases.UInt64LT{Val: 5}, false},
	{"uint64 - lt - invalid", &cases.UInt64LT{Val: 6}, false},

	{"uint64 - lte - valid", &cases.UInt64LTE{Val: 63}, true},
	{"uint64 - lte - valid (equal)", &cases.UInt64LTE{Val: 64}, true},
	{"uint64 - lte - invalid", &cases.UInt64LTE{Val: 65}, false},

	{"uint64 - gt - valid", &cases.UInt64GT{Val: 17}, true},
	{"uint64 - gt - invalid (equal)", &cases.UInt64GT{Val: 16}, false},
	{"uint64 - gt - invalid", &cases.UInt64GT{Val: 15}, false},

	{"uint64 - gte - valid", &cases.UInt64GTE{Val: 9}, true},
	{"uint64 - gte - valid (equal)", &cases.UInt64GTE{Val: 8}, true},
	{"uint64 - gte - invalid", &cases.UInt64GTE{Val: 7}, false},

	{"uint64 - gt & lt - valid", &cases.UInt64GTLT{Val: 7}, true},
	{"uint64 - gt & lt - invalid (above)", &cases.UInt64GTLT{Val: 11}, false},
	{"uint64 - gt & lt - invalid (below)", &cases.UInt64GTLT{Val: 1}, false},
	{"uint64 - gt & lt - invalid (max)", &cases.UInt64GTLT{Val: 10}, false},
	{"uint64 - gt & lt - invalid (min)", &cases.UInt64GTLT{Val: 5}, false},

	{"uint64 - exclusive gt & lt - valid (above)", &cases.UInt64ExLTGT{Val: 11}, true},
	{"uint64 - exclusive gt & lt - valid (below)", &cases.UInt64ExLTGT{Val: 4}, true},
	{"uint64 - exclusive gt & lt - invalid", &cases.UInt64ExLTGT{Val: 7}, false},
	{"uint64 - exclusive gt & lt - invalid (max)", &cases.UInt64ExLTGT{Val: 10}, false},
	{"uint64 - exclusive gt & lt - invalid (min)", &cases.UInt64ExLTGT{Val: 5}, false},

	{"uint64 - gte & lte - valid", &cases.UInt64GTELTE{Val: 200}, true},
	{"uint64 - gte & lte - valid (max)", &cases.UInt64GTELTE{Val: 256}, true},
	{"uint64 - gte & lte - valid (min)", &cases.UInt64GTELTE{Val: 128}, true},
	{"uint64 - gte & lte - invalid (above)", &cases.UInt64GTELTE{Val: 300}, false},
	{"uint64 - gte & lte - invalid (below)", &cases.UInt64GTELTE{Val: 100}, false},

	{"uint64 - exclusive gte & lte - valid (above)", &cases.UInt64ExGTELTE{Val: 300}, true},
	{"uint64 - exclusive gte & lte - valid (below)", &cases.UInt64ExGTELTE{Val: 100}, true},
	{"uint64 - exclusive gte & lte - valid (max)", &cases.UInt64ExGTELTE{Val: 256}, true},
	{"uint64 - exclusive gte & lte - valid (min)", &cases.UInt64ExGTELTE{Val: 128}, true},
	{"uint64 - exclusive gte & lte - invalid", &cases.UInt64ExGTELTE{Val: 200}, false},
}

var sint32Cases = []TestCase{
	{"sint32 - none - valid", &cases.SInt32None{Val: 123}, true},

	{"sint32 - const - valid", &cases.SInt32Const{Val: 1}, true},
	{"sint32 - const - invalid", &cases.SInt32Const{Val: 2}, false},

	{"sint32 - in - valid", &cases.SInt32In{Val: 3}, true},
	{"sint32 - in - invalid", &cases.SInt32In{Val: 5}, false},

	{"sint32 - not in - valid", &cases.SInt32NotIn{Val: 1}, true},
	{"sint32 - not in - invalid", &cases.SInt32NotIn{Val: 0}, false},

	{"sint32 - lt - valid", &cases.SInt32LT{Val: -1}, true},
	{"sint32 - lt - invalid (equal)", &cases.SInt32LT{Val: 0}, false},
	{"sint32 - lt - invalid", &cases.SInt32LT{Val: 1}, false},

	{"sint32 - lte - valid", &cases.SInt32LTE{Val: 63}, true},
	{"sint32 - lte - valid (equal)", &cases.SInt32LTE{Val: 64}, true},
	{"sint32 - lte - invalid", &cases.SInt32LTE{Val: 65}, false},

	{"sint32 - gt - valid", &cases.SInt32GT{Val: 17}, true},
	{"sint32 - gt - invalid (equal)", &cases.SInt32GT{Val: 16}, false},
	{"sint32 - gt - invalid", &cases.SInt32GT{Val: 15}, false},

	{"sint32 - gte - valid", &cases.SInt32GTE{Val: 9}, true},
	{"sint32 - gte - valid (equal)", &cases.SInt32GTE{Val: 8}, true},
	{"sint32 - gte - invalid", &cases.SInt32GTE{Val: 7}, false},

	{"sint32 - gt & lt - valid", &cases.SInt32GTLT{Val: 5}, true},
	{"sint32 - gt & lt - invalid (above)", &cases.SInt32GTLT{Val: 11}, false},
	{"sint32 - gt & lt - invalid (below)", &cases.SInt32GTLT{Val: -1}, false},
	{"sint32 - gt & lt - invalid (max)", &cases.SInt32GTLT{Val: 10}, false},
	{"sint32 - gt & lt - invalid (min)", &cases.SInt32GTLT{Val: 0}, false},

	{"sint32 - exclusive gt & lt - valid (above)", &cases.SInt32ExLTGT{Val: 11}, true},
	{"sint32 - exclusive gt & lt - valid (below)", &cases.SInt32ExLTGT{Val: -1}, true},
	{"sint32 - exclusive gt & lt - invalid", &cases.SInt32ExLTGT{Val: 5}, false},
	{"sint32 - exclusive gt & lt - invalid (max)", &cases.SInt32ExLTGT{Val: 10}, false},
	{"sint32 - exclusive gt & lt - invalid (min)", &cases.SInt32ExLTGT{Val: 0}, false},

	{"sint32 - gte & lte - valid", &cases.SInt32GTELTE{Val: 200}, true},
	{"sint32 - gte & lte - valid (max)", &cases.SInt32GTELTE{Val: 256}, true},
	{"sint32 - gte & lte - valid (min)", &cases.SInt32GTELTE{Val: 128}, true},
	{"sint32 - gte & lte - invalid (above)", &cases.SInt32GTELTE{Val: 300}, false},
	{"sint32 - gte & lte - invalid (below)", &cases.SInt32GTELTE{Val: 100}, false},

	{"sint32 - exclusive gte & lte - valid (above)", &cases.SInt32ExGTELTE{Val: 300}, true},
	{"sint32 - exclusive gte & lte - valid (below)", &cases.SInt32ExGTELTE{Val: 100}, true},
	{"sint32 - exclusive gte & lte - valid (max)", &cases.SInt32ExGTELTE{Val: 256}, true},
	{"sint32 - exclusive gte & lte - valid (min)", &cases.SInt32ExGTELTE{Val: 128}, true},
	{"sint32 - exclusive gte & lte - invalid", &cases.SInt32ExGTELTE{Val: 200}, false},
}

var sint64Cases = []TestCase{
	{"sint64 - none - valid", &cases.SInt64None{Val: 123}, true},

	{"sint64 - const - valid", &cases.SInt64Const{Val: 1}, true},
	{"sint64 - const - invalid", &cases.SInt64Const{Val: 2}, false},

	{"sint64 - in - valid", &cases.SInt64In{Val: 3}, true},
	{"sint64 - in - invalid", &cases.SInt64In{Val: 5}, false},

	{"sint64 - not in - valid", &cases.SInt64NotIn{Val: 1}, true},
	{"sint64 - not in - invalid", &cases.SInt64NotIn{Val: 0}, false},

	{"sint64 - lt - valid", &cases.SInt64LT{Val: -1}, true},
	{"sint64 - lt - invalid (equal)", &cases.SInt64LT{Val: 0}, false},
	{"sint64 - lt - invalid", &cases.SInt64LT{Val: 1}, false},

	{"sint64 - lte - valid", &cases.SInt64LTE{Val: 63}, true},
	{"sint64 - lte - valid (equal)", &cases.SInt64LTE{Val: 64}, true},
	{"sint64 - lte - invalid", &cases.SInt64LTE{Val: 65}, false},

	{"sint64 - gt - valid", &cases.SInt64GT{Val: 17}, true},
	{"sint64 - gt - invalid (equal)", &cases.SInt64GT{Val: 16}, false},
	{"sint64 - gt - invalid", &cases.SInt64GT{Val: 15}, false},

	{"sint64 - gte - valid", &cases.SInt64GTE{Val: 9}, true},
	{"sint64 - gte - valid (equal)", &cases.SInt64GTE{Val: 8}, true},
	{"sint64 - gte - invalid", &cases.SInt64GTE{Val: 7}, false},

	{"sint64 - gt & lt - valid", &cases.SInt64GTLT{Val: 5}, true},
	{"sint64 - gt & lt - invalid (above)", &cases.SInt64GTLT{Val: 11}, false},
	{"sint64 - gt & lt - invalid (below)", &cases.SInt64GTLT{Val: -1}, false},
	{"sint64 - gt & lt - invalid (max)", &cases.SInt64GTLT{Val: 10}, false},
	{"sint64 - gt & lt - invalid (min)", &cases.SInt64GTLT{Val: 0}, false},

	{"sint64 - exclusive gt & lt - valid (above)", &cases.SInt64ExLTGT{Val: 11}, true},
	{"sint64 - exclusive gt & lt - valid (below)", &cases.SInt64ExLTGT{Val: -1}, true},
	{"sint64 - exclusive gt & lt - invalid", &cases.SInt64ExLTGT{Val: 5}, false},
	{"sint64 - exclusive gt & lt - invalid (max)", &cases.SInt64ExLTGT{Val: 10}, false},
	{"sint64 - exclusive gt & lt - invalid (min)", &cases.SInt64ExLTGT{Val: 0}, false},

	{"sint64 - gte & lte - valid", &cases.SInt64GTELTE{Val: 200}, true},
	{"sint64 - gte & lte - valid (max)", &cases.SInt64GTELTE{Val: 256}, true},
	{"sint64 - gte & lte - valid (min)", &cases.SInt64GTELTE{Val: 128}, true},
	{"sint64 - gte & lte - invalid (above)", &cases.SInt64GTELTE{Val: 300}, false},
	{"sint64 - gte & lte - invalid (below)", &cases.SInt64GTELTE{Val: 100}, false},

	{"sint64 - exclusive gte & lte - valid (above)", &cases.SInt64ExGTELTE{Val: 300}, true},
	{"sint64 - exclusive gte & lte - valid (below)", &cases.SInt64ExGTELTE{Val: 100}, true},
	{"sint64 - exclusive gte & lte - valid (max)", &cases.SInt64ExGTELTE{Val: 256}, true},
	{"sint64 - exclusive gte & lte - valid (min)", &cases.SInt64ExGTELTE{Val: 128}, true},
	{"sint64 - exclusive gte & lte - invalid", &cases.SInt64ExGTELTE{Val: 200}, false},
}

var fixed32Cases = []TestCase{
	{"fixed32 - none - valid", &cases.Fixed32None{Val: 123}, true},

	{"fixed32 - const - valid", &cases.Fixed32Const{Val: 1}, true},
	{"fixed32 - const - invalid", &cases.Fixed32Const{Val: 2}, false},

	{"fixed32 - in - valid", &cases.Fixed32In{Val: 3}, true},
	{"fixed32 - in - invalid", &cases.Fixed32In{Val: 5}, false},

	{"fixed32 - not in - valid", &cases.Fixed32NotIn{Val: 1}, true},
	{"fixed32 - not in - invalid", &cases.Fixed32NotIn{Val: 0}, false},

	{"fixed32 - lt - valid", &cases.Fixed32LT{Val: 4}, true},
	{"fixed32 - lt - invalid (equal)", &cases.Fixed32LT{Val: 5}, false},
	{"fixed32 - lt - invalid", &cases.Fixed32LT{Val: 6}, false},

	{"fixed32 - lte - valid", &cases.Fixed32LTE{Val: 63}, true},
	{"fixed32 - lte - valid (equal)", &cases.Fixed32LTE{Val: 64}, true},
	{"fixed32 - lte - invalid", &cases.Fixed32LTE{Val: 65}, false},

	{"fixed32 - gt - valid", &cases.Fixed32GT{Val: 17}, true},
	{"fixed32 - gt - invalid (equal)", &cases.Fixed32GT{Val: 16}, false},
	{"fixed32 - gt - invalid", &cases.Fixed32GT{Val: 15}, false},

	{"fixed32 - gte - valid", &cases.Fixed32GTE{Val: 9}, true},
	{"fixed32 - gte - valid (equal)", &cases.Fixed32GTE{Val: 8}, true},
	{"fixed32 - gte - invalid", &cases.Fixed32GTE{Val: 7}, false},

	{"fixed32 - gt & lt - valid", &cases.Fixed32GTLT{Val: 7}, true},
	{"fixed32 - gt & lt - invalid (above)", &cases.Fixed32GTLT{Val: 11}, false},
	{"fixed32 - gt & lt - invalid (below)", &cases.Fixed32GTLT{Val: 1}, false},
	{"fixed32 - gt & lt - invalid (max)", &cases.Fixed32GTLT{Val: 10}, false},
	{"fixed32 - gt & lt - invalid (min)", &cases.Fixed32GTLT{Val: 5}, false},

	{"fixed32 - exclusive gt & lt - valid (above)", &cases.Fixed32ExLTGT{Val: 11}, true},
	{"fixed32 - exclusive gt & lt - valid (below)", &cases.Fixed32ExLTGT{Val: 4}, true},
	{"fixed32 - exclusive gt & lt - invalid", &cases.Fixed32ExLTGT{Val: 7}, false},
	{"fixed32 - exclusive gt & lt - invalid (max)", &cases.Fixed32ExLTGT{Val: 10}, false},
	{"fixed32 - exclusive gt & lt - invalid (min)", &cases.Fixed32ExLTGT{Val: 5}, false},

	{"fixed32 - gte & lte - valid", &cases.Fixed32GTELTE{Val: 200}, true},
	{"fixed32 - gte & lte - valid (max)", &cases.Fixed32GTELTE{Val: 256}, true},
	{"fixed32 - gte & lte - valid (min)", &cases.Fixed32GTELTE{Val: 128}, true},
	{"fixed32 - gte & lte - invalid (above)", &cases.Fixed32GTELTE{Val: 300}, false},
	{"fixed32 - gte & lte - invalid (below)", &cases.Fixed32GTELTE{Val: 100}, false},

	{"fixed32 - exclusive gte & lte - valid (above)", &cases.Fixed32ExGTELTE{Val: 300}, true},
	{"fixed32 - exclusive gte & lte - valid (below)", &cases.Fixed32ExGTELTE{Val: 100}, true},
	{"fixed32 - exclusive gte & lte - valid (max)", &cases.Fixed32ExGTELTE{Val: 256}, true},
	{"fixed32 - exclusive gte & lte - valid (min)", &cases.Fixed32ExGTELTE{Val: 128}, true},
	{"fixed32 - exclusive gte & lte - invalid", &cases.Fixed32ExGTELTE{Val: 200}, false},
}

var fixed64Cases = []TestCase{
	{"fixed64 - none - valid", &cases.Fixed64None{Val: 123}, true},

	{"fixed64 - const - valid", &cases.Fixed64Const{Val: 1}, true},
	{"fixed64 - const - invalid", &cases.Fixed64Const{Val: 2}, false},

	{"fixed64 - in - valid", &cases.Fixed64In{Val: 3}, true},
	{"fixed64 - in - invalid", &cases.Fixed64In{Val: 5}, false},

	{"fixed64 - not in - valid", &cases.Fixed64NotIn{Val: 1}, true},
	{"fixed64 - not in - invalid", &cases.Fixed64NotIn{Val: 0}, false},

	{"fixed64 - lt - valid", &cases.Fixed64LT{Val: 4}, true},
	{"fixed64 - lt - invalid (equal)", &cases.Fixed64LT{Val: 5}, false},
	{"fixed64 - lt - invalid", &cases.Fixed64LT{Val: 6}, false},

	{"fixed64 - lte - valid", &cases.Fixed64LTE{Val: 63}, true},
	{"fixed64 - lte - valid (equal)", &cases.Fixed64LTE{Val: 64}, true},
	{"fixed64 - lte - invalid", &cases.Fixed64LTE{Val: 65}, false},

	{"fixed64 - gt - valid", &cases.Fixed64GT{Val: 17}, true},
	{"fixed64 - gt - invalid (equal)", &cases.Fixed64GT{Val: 16}, false},
	{"fixed64 - gt - invalid", &cases.Fixed64GT{Val: 15}, false},

	{"fixed64 - gte - valid", &cases.Fixed64GTE{Val: 9}, true},
	{"fixed64 - gte - valid (equal)", &cases.Fixed64GTE{Val: 8}, true},
	{"fixed64 - gte - invalid", &cases.Fixed64GTE{Val: 7}, false},

	{"fixed64 - gt & lt - valid", &cases.Fixed64GTLT{Val: 7}, true},
	{"fixed64 - gt & lt - invalid (above)", &cases.Fixed64GTLT{Val: 11}, false},
	{"fixed64 - gt & lt - invalid (below)", &cases.Fixed64GTLT{Val: 1}, false},
	{"fixed64 - gt & lt - invalid (max)", &cases.Fixed64GTLT{Val: 10}, false},
	{"fixed64 - gt & lt - invalid (min)", &cases.Fixed64GTLT{Val: 5}, false},

	{"fixed64 - exclusive gt & lt - valid (above)", &cases.Fixed64ExLTGT{Val: 11}, true},
	{"fixed64 - exclusive gt & lt - valid (below)", &cases.Fixed64ExLTGT{Val: 4}, true},
	{"fixed64 - exclusive gt & lt - invalid", &cases.Fixed64ExLTGT{Val: 7}, false},
	{"fixed64 - exclusive gt & lt - invalid (max)", &cases.Fixed64ExLTGT{Val: 10}, false},
	{"fixed64 - exclusive gt & lt - invalid (min)", &cases.Fixed64ExLTGT{Val: 5}, false},

	{"fixed64 - gte & lte - valid", &cases.Fixed64GTELTE{Val: 200}, true},
	{"fixed64 - gte & lte - valid (max)", &cases.Fixed64GTELTE{Val: 256}, true},
	{"fixed64 - gte & lte - valid (min)", &cases.Fixed64GTELTE{Val: 128}, true},
	{"fixed64 - gte & lte - invalid (above)", &cases.Fixed64GTELTE{Val: 300}, false},
	{"fixed64 - gte & lte - invalid (below)", &cases.Fixed64GTELTE{Val: 100}, false},

	{"fixed64 - exclusive gte & lte - valid (above)", &cases.Fixed64ExGTELTE{Val: 300}, true},
	{"fixed64 - exclusive gte & lte - valid (below)", &cases.Fixed64ExGTELTE{Val: 100}, true},
	{"fixed64 - exclusive gte & lte - valid (max)", &cases.Fixed64ExGTELTE{Val: 256}, true},
	{"fixed64 - exclusive gte & lte - valid (min)", &cases.Fixed64ExGTELTE{Val: 128}, true},
	{"fixed64 - exclusive gte & lte - invalid", &cases.Fixed64ExGTELTE{Val: 200}, false},
}

var sfixed32Cases = []TestCase{
	{"sfixed32 - none - valid", &cases.SFixed32None{Val: 123}, true},

	{"sfixed32 - const - valid", &cases.SFixed32Const{Val: 1}, true},
	{"sfixed32 - const - invalid", &cases.SFixed32Const{Val: 2}, false},

	{"sfixed32 - in - valid", &cases.SFixed32In{Val: 3}, true},
	{"sfixed32 - in - invalid", &cases.SFixed32In{Val: 5}, false},

	{"sfixed32 - not in - valid", &cases.SFixed32NotIn{Val: 1}, true},
	{"sfixed32 - not in - invalid", &cases.SFixed32NotIn{Val: 0}, false},

	{"sfixed32 - lt - valid", &cases.SFixed32LT{Val: -1}, true},
	{"sfixed32 - lt - invalid (equal)", &cases.SFixed32LT{Val: 0}, false},
	{"sfixed32 - lt - invalid", &cases.SFixed32LT{Val: 1}, false},

	{"sfixed32 - lte - valid", &cases.SFixed32LTE{Val: 63}, true},
	{"sfixed32 - lte - valid (equal)", &cases.SFixed32LTE{Val: 64}, true},
	{"sfixed32 - lte - invalid", &cases.SFixed32LTE{Val: 65}, false},

	{"sfixed32 - gt - valid", &cases.SFixed32GT{Val: 17}, true},
	{"sfixed32 - gt - invalid (equal)", &cases.SFixed32GT{Val: 16}, false},
	{"sfixed32 - gt - invalid", &cases.SFixed32GT{Val: 15}, false},

	{"sfixed32 - gte - valid", &cases.SFixed32GTE{Val: 9}, true},
	{"sfixed32 - gte - valid (equal)", &cases.SFixed32GTE{Val: 8}, true},
	{"sfixed32 - gte - invalid", &cases.SFixed32GTE{Val: 7}, false},

	{"sfixed32 - gt & lt - valid", &cases.SFixed32GTLT{Val: 5}, true},
	{"sfixed32 - gt & lt - invalid (above)", &cases.SFixed32GTLT{Val: 11}, false},
	{"sfixed32 - gt & lt - invalid (below)", &cases.SFixed32GTLT{Val: -1}, false},
	{"sfixed32 - gt & lt - invalid (max)", &cases.SFixed32GTLT{Val: 10}, false},
	{"sfixed32 - gt & lt - invalid (min)", &cases.SFixed32GTLT{Val: 0}, false},

	{"sfixed32 - exclusive gt & lt - valid (above)", &cases.SFixed32ExLTGT{Val: 11}, true},
	{"sfixed32 - exclusive gt & lt - valid (below)", &cases.SFixed32ExLTGT{Val: -1}, true},
	{"sfixed32 - exclusive gt & lt - invalid", &cases.SFixed32ExLTGT{Val: 5}, false},
	{"sfixed32 - exclusive gt & lt - invalid (max)", &cases.SFixed32ExLTGT{Val: 10}, false},
	{"sfixed32 - exclusive gt & lt - invalid (min)", &cases.SFixed32ExLTGT{Val: 0}, false},

	{"sfixed32 - gte & lte - valid", &cases.SFixed32GTELTE{Val: 200}, true},
	{"sfixed32 - gte & lte - valid (max)", &cases.SFixed32GTELTE{Val: 256}, true},
	{"sfixed32 - gte & lte - valid (min)", &cases.SFixed32GTELTE{Val: 128}, true},
	{"sfixed32 - gte & lte - invalid (above)", &cases.SFixed32GTELTE{Val: 300}, false},
	{"sfixed32 - gte & lte - invalid (below)", &cases.SFixed32GTELTE{Val: 100}, false},

	{"sfixed32 - exclusive gte & lte - valid (above)", &cases.SFixed32ExGTELTE{Val: 300}, true},
	{"sfixed32 - exclusive gte & lte - valid (below)", &cases.SFixed32ExGTELTE{Val: 100}, true},
	{"sfixed32 - exclusive gte & lte - valid (max)", &cases.SFixed32ExGTELTE{Val: 256}, true},
	{"sfixed32 - exclusive gte & lte - valid (min)", &cases.SFixed32ExGTELTE{Val: 128}, true},
	{"sfixed32 - exclusive gte & lte - invalid", &cases.SFixed32ExGTELTE{Val: 200}, false},
}

var sfixed64Cases = []TestCase{
	{"sfixed64 - none - valid", &cases.SFixed64None{Val: 123}, true},

	{"sfixed64 - const - valid", &cases.SFixed64Const{Val: 1}, true},
	{"sfixed64 - const - invalid", &cases.SFixed64Const{Val: 2}, false},

	{"sfixed64 - in - valid", &cases.SFixed64In{Val: 3}, true},
	{"sfixed64 - in - invalid", &cases.SFixed64In{Val: 5}, false},

	{"sfixed64 - not in - valid", &cases.SFixed64NotIn{Val: 1}, true},
	{"sfixed64 - not in - invalid", &cases.SFixed64NotIn{Val: 0}, false},

	{"sfixed64 - lt - valid", &cases.SFixed64LT{Val: -1}, true},
	{"sfixed64 - lt - invalid (equal)", &cases.SFixed64LT{Val: 0}, false},
	{"sfixed64 - lt - invalid", &cases.SFixed64LT{Val: 1}, false},

	{"sfixed64 - lte - valid", &cases.SFixed64LTE{Val: 63}, true},
	{"sfixed64 - lte - valid (equal)", &cases.SFixed64LTE{Val: 64}, true},
	{"sfixed64 - lte - invalid", &cases.SFixed64LTE{Val: 65}, false},

	{"sfixed64 - gt - valid", &cases.SFixed64GT{Val: 17}, true},
	{"sfixed64 - gt - invalid (equal)", &cases.SFixed64GT{Val: 16}, false},
	{"sfixed64 - gt - invalid", &cases.SFixed64GT{Val: 15}, false},

	{"sfixed64 - gte - valid", &cases.SFixed64GTE{Val: 9}, true},
	{"sfixed64 - gte - valid (equal)", &cases.SFixed64GTE{Val: 8}, true},
	{"sfixed64 - gte - invalid", &cases.SFixed64GTE{Val: 7}, false},

	{"sfixed64 - gt & lt - valid", &cases.SFixed64GTLT{Val: 5}, true},
	{"sfixed64 - gt & lt - invalid (above)", &cases.SFixed64GTLT{Val: 11}, false},
	{"sfixed64 - gt & lt - invalid (below)", &cases.SFixed64GTLT{Val: -1}, false},
	{"sfixed64 - gt & lt - invalid (max)", &cases.SFixed64GTLT{Val: 10}, false},
	{"sfixed64 - gt & lt - invalid (min)", &cases.SFixed64GTLT{Val: 0}, false},

	{"sfixed64 - exclusive gt & lt - valid (above)", &cases.SFixed64ExLTGT{Val: 11}, true},
	{"sfixed64 - exclusive gt & lt - valid (below)", &cases.SFixed64ExLTGT{Val: -1}, true},
	{"sfixed64 - exclusive gt & lt - invalid", &cases.SFixed64ExLTGT{Val: 5}, false},
	{"sfixed64 - exclusive gt & lt - invalid (max)", &cases.SFixed64ExLTGT{Val: 10}, false},
	{"sfixed64 - exclusive gt & lt - invalid (min)", &cases.SFixed64ExLTGT{Val: 0}, false},

	{"sfixed64 - gte & lte - valid", &cases.SFixed64GTELTE{Val: 200}, true},
	{"sfixed64 - gte & lte - valid (max)", &cases.SFixed64GTELTE{Val: 256}, true},
	{"sfixed64 - gte & lte - valid (min)", &cases.SFixed64GTELTE{Val: 128}, true},
	{"sfixed64 - gte & lte - invalid (above)", &cases.SFixed64GTELTE{Val: 300}, false},
	{"sfixed64 - gte & lte - invalid (below)", &cases.SFixed64GTELTE{Val: 100}, false},

	{"sfixed64 - exclusive gte & lte - valid (above)", &cases.SFixed64ExGTELTE{Val: 300}, true},
	{"sfixed64 - exclusive gte & lte - valid (below)", &cases.SFixed64ExGTELTE{Val: 100}, true},
	{"sfixed64 - exclusive gte & lte - valid (max)", &cases.SFixed64ExGTELTE{Val: 256}, true},
	{"sfixed64 - exclusive gte & lte - valid (min)", &cases.SFixed64ExGTELTE{Val: 128}, true},
	{"sfixed64 - exclusive gte & lte - invalid", &cases.SFixed64ExGTELTE{Val: 200}, false},
}

var boolCases = []TestCase{
	{"bool - none - valid", &cases.BoolNone{Val: true}, true},
	{"bool - const (true) - valid", &cases.BoolConstTrue{Val: true}, true},
	{"bool - const (true) - invalid", &cases.BoolConstTrue{Val: false}, false},
	{"bool - const (false) - valid", &cases.BoolConstFalse{Val: false}, true},
	{"bool - const (false) - invalid", &cases.BoolConstFalse{Val: true}, false},
}

var stringCases = []TestCase{
	{"string - none - valid", &cases.StringNone{Val: "quux"}, true},

	{"string - const - valid", &cases.StringConst{Val: "foo"}, true},
	{"string - const - invalid", &cases.StringConst{Val: "bar"}, false},

	{"string - in - valid", &cases.StringIn{Val: "bar"}, true},
	{"string - in - invalid", &cases.StringIn{Val: "quux"}, false},
	{"string - not in - valid", &cases.StringNotIn{Val: "quux"}, true},
	{"string - not in - invalid", &cases.StringNotIn{"fizz"}, false},

	{"string - min len - valid", &cases.StringMinLen{Val: "protoc"}, true},
	{"string - min len - valid (min)", &cases.StringMinLen{Val: "baz"}, true},
	{"string - min len - invalid", &cases.StringMinLen{Val: "go"}, false},
	{"string - min len - invalid (multibyte)", &cases.StringMinLen{Val: "你好"}, false},

	{"string - max len - valid", &cases.StringMaxLen{Val: "foo"}, true},
	{"string - max len - valid (max)", &cases.StringMaxLen{Val: "proto"}, true},
	{"string - max len - valid (multibyte)", &cases.StringMaxLen{Val: "你好你好"}, true},
	{"string - max len - invalid", &cases.StringMaxLen{Val: "1234567890"}, false},

	{"string - min/max len - valid", &cases.StringMinMaxLen{Val: "quux"}, true},
	{"string - min/max len - valid (min)", &cases.StringMinMaxLen{Val: "foo"}, true},
	{"string - min/max len - valid (max)", &cases.StringMinMaxLen{Val: "proto"}, true},
	{"string - min/max len - valid (multibyte)", &cases.StringMinMaxLen{Val: "你好你好"}, true},
	{"string - min/max len - invalid (below)", &cases.StringMinMaxLen{Val: "go"}, false},
	{"string - min/max len - invalid (above)", &cases.StringMinMaxLen{Val: "validate"}, false},

	{"string - min bytes - valid", &cases.StringMinBytes{Val: "proto"}, true},
	{"string - min bytes - valid (min)", &cases.StringMinBytes{Val: "quux"}, true},
	{"string - min bytes - valid (multibyte)", &cases.StringMinBytes{Val: "你好"}, true},
	{"string - min bytes - invalid", &cases.StringMinBytes{Val: ""}, false},

	{"string - max bytes - valid", &cases.StringMaxBytes{Val: "foo"}, true},
	{"string - max bytes - valid (max)", &cases.StringMaxBytes{Val: "12345678"}, true},
	{"string - max bytes - invalid", &cases.StringMaxBytes{Val: "123456789"}, false},
	{"string - max bytes - invalid (multibyte)", &cases.StringMaxBytes{Val: "你好你好你好"}, false},

	{"string - min/max bytes - valid", &cases.StringMinMaxBytes{Val: "protoc"}, true},
	{"string - min/max bytes - valid (min)", &cases.StringMinMaxBytes{Val: "quux"}, true},
	{"string - min/max bytes - valid (max)", &cases.StringMinMaxBytes{Val: "fizzbuzz"}, true},
	{"string - min/max bytes - valid (multibyte)", &cases.StringMinMaxBytes{Val: "你好"}, true},
	{"string - min/max bytes - invalid (below)", &cases.StringMinMaxBytes{Val: "foo"}, false},
	{"string - min/max bytes - invalid (above)", &cases.StringMinMaxBytes{Val: "你好你好你"}, false},

	{"string - pattern - valid", &cases.StringPattern{Val: "Foo123"}, true},
	{"string - pattern - invalid", &cases.StringPattern{Val: "!@#$%^&*()"}, false},
	{"string - pattern - invalid (empty)", &cases.StringPattern{Val: ""}, false},

	{"string - prefix - valid", &cases.StringPrefix{Val: "foobar"}, true},
	{"string - prefix - valid (only)", &cases.StringPrefix{Val: "foo"}, true},
	{"string - prefix - invalid", &cases.StringPrefix{Val: "bar"}, false},
	{"string - prefix - invalid (case-sensitive)", &cases.StringPrefix{Val: "Foobar"}, false},

	{"string - contains - valid", &cases.StringContains{Val: "candy bars"}, true},
	{"string - contains - valid (only)", &cases.StringContains{Val: "bar"}, true},
	{"string - contains - invalid", &cases.StringContains{Val: "candy bazs"}, false},
	{"string - contains - invalid (case-sensitive)", &cases.StringContains{Val: "Candy Bars"}, false},

	{"string - suffix - valid", &cases.StringSuffix{Val: "foobaz"}, true},
	{"string - suffix - valid (only)", &cases.StringSuffix{Val: "baz"}, true},
	{"string - suffix - invalid", &cases.StringSuffix{Val: "foobar"}, false},
	{"string - suffix - invalid (case-sensitive)", &cases.StringSuffix{Val: "FooBaz"}, false},

	{"string - email - valid", &cases.StringEmail{Val: "foo@bar.com"}, true},
	{"string - email - valid (name)", &cases.StringEmail{Val: "John Smith <foo@bar.com>"}, true},
	{"string - email - invalid", &cases.StringEmail{"foobar"}, false},
	{"string - email - invalid (local segment too long)", &cases.StringEmail{"x0123456789012345678901234567890123456789012345678901234567890123456789@example.com"}, false},
	{"string - email - invalid (hostname too long)", &cases.StringEmail{"foo@x0123456789012345678901234567890123456789012345678901234567890123456789.com"}, false},
	{"string - email - invalid (bad hostname)", &cases.StringEmail{"foo@-bar.com"}, false},

	{"string - hostname - valid", &cases.StringHostname{Val: "example.com"}, true},
	{"string - hostname - invalid", &cases.StringHostname{Val: "!@#$%^&"}, false},
	{"string - hostname - invalid (underscore)", &cases.StringHostname{Val: "foo_bar.com"}, false},
	{"string - hostname - invalid (too long)", &cases.StringHostname{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789.com"}, false},
	{"string - hostname - invalid (IDNs)", &cases.StringHostname{Val: "你好.com"}, false},

	{"string - IP - valid (v4)", &cases.StringIP{Val: "192.168.0.1"}, true},
	{"string - IP - valid (v6)", &cases.StringIP{Val: "3e::99"}, true},
	{"string - IP - invalid", &cases.StringIP{Val: "foobar"}, false},

	{"string - IPv4 - valid", &cases.StringIPv4{Val: "192.168.0.1"}, true},
	{"string - IPv4 - invalid", &cases.StringIPv4{Val: "foobar"}, false},
	{"string - IPv4 - invalid (erroneous)", &cases.StringIPv4{Val: "256.0.0.0"}, false},
	{"string - IPv4 - invalid (v6)", &cases.StringIPv4{Val: "3e::99"}, false},

	{"string - IPv6 - valid", &cases.StringIPv6{"2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, true},
	{"string - IPv6 - valid (collapsed)", &cases.StringIPv6{"2001:db8:85a3::8a2e:370:7334"}, true},
	{"string - IPv6 - invalid", &cases.StringIPv6{"foobar"}, false},
	{"string - IPv6 - invalid (v4)", &cases.StringIPv6{"192.168.0.1"}, false},
	{"string - IPv6 - invalid (erroneous)", &cases.StringIPv6{"efgh::0b"}, false},
}
