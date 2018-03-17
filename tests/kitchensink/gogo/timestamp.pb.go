// Code generated by protoc-gen-gogo.
// source: timestamp.proto
// DO NOT EDIT!

package tests_kitchensink

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Timestamp struct {
	None        *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=none" json:"none,omitempty"`
	Lt          *google_protobuf2.Timestamp `protobuf:"bytes,2,opt,name=lt" json:"lt,omitempty"`
	Lte         *google_protobuf2.Timestamp `protobuf:"bytes,3,opt,name=lte" json:"lte,omitempty"`
	Gt          *google_protobuf2.Timestamp `protobuf:"bytes,4,opt,name=gt" json:"gt,omitempty"`
	Gte         *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=gte" json:"gte,omitempty"`
	LtGt        *google_protobuf2.Timestamp `protobuf:"bytes,6,opt,name=lt_gt,json=ltGt" json:"lt_gt,omitempty"`
	LtGte       *google_protobuf2.Timestamp `protobuf:"bytes,7,opt,name=lt_gte,json=ltGte" json:"lt_gte,omitempty"`
	LteGt       *google_protobuf2.Timestamp `protobuf:"bytes,8,opt,name=lte_gt,json=lteGt" json:"lte_gt,omitempty"`
	LteGte      *google_protobuf2.Timestamp `protobuf:"bytes,9,opt,name=lte_gte,json=lteGte" json:"lte_gte,omitempty"`
	LtGtInv     *google_protobuf2.Timestamp `protobuf:"bytes,10,opt,name=lt_gt_inv,json=ltGtInv" json:"lt_gt_inv,omitempty"`
	LtGteInv    *google_protobuf2.Timestamp `protobuf:"bytes,11,opt,name=lt_gte_inv,json=ltGteInv" json:"lt_gte_inv,omitempty"`
	LteGtInv    *google_protobuf2.Timestamp `protobuf:"bytes,12,opt,name=lte_gt_inv,json=lteGtInv" json:"lte_gt_inv,omitempty"`
	LteGteInv   *google_protobuf2.Timestamp `protobuf:"bytes,13,opt,name=lte_gte_inv,json=lteGteInv" json:"lte_gte_inv,omitempty"`
	Required    *google_protobuf2.Timestamp `protobuf:"bytes,14,opt,name=required" json:"required,omitempty"`
	LtNow       *google_protobuf2.Timestamp `protobuf:"bytes,15,opt,name=lt_now,json=ltNow" json:"lt_now,omitempty"`
	GtNow       *google_protobuf2.Timestamp `protobuf:"bytes,16,opt,name=gt_now,json=gtNow" json:"gt_now,omitempty"`
	LtNowWithin *google_protobuf2.Timestamp `protobuf:"bytes,17,opt,name=lt_now_within,json=ltNowWithin" json:"lt_now_within,omitempty"`
	GtNowWithin *google_protobuf2.Timestamp `protobuf:"bytes,18,opt,name=gt_now_within,json=gtNowWithin" json:"gt_now_within,omitempty"`
	Within      *google_protobuf2.Timestamp `protobuf:"bytes,19,opt,name=within" json:"within,omitempty"`
	Gogo1       *time.Time                  `protobuf:"bytes,20,opt,name=gogo1,stdtime" json:"gogo1,omitempty"`
	Gogo2       google_protobuf2.Timestamp  `protobuf:"bytes,21,opt,name=gogo2" json:"gogo2"`
	Gogo3       time.Time                   `protobuf:"bytes,22,opt,name=gogo3,stdtime" json:"gogo3"`
	Gogo4       google_protobuf2.Timestamp  `protobuf:"bytes,23,opt,name=gogo4" json:"gogo4"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptorTimestamp, []int{0} }

func (m *Timestamp) GetNone() *google_protobuf2.Timestamp {
	if m != nil {
		return m.None
	}
	return nil
}

func (m *Timestamp) GetLt() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Lt
	}
	return nil
}

func (m *Timestamp) GetLte() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Lte
	}
	return nil
}

func (m *Timestamp) GetGt() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Gt
	}
	return nil
}

func (m *Timestamp) GetGte() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Gte
	}
	return nil
}

func (m *Timestamp) GetLtGt() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LtGt
	}
	return nil
}

func (m *Timestamp) GetLtGte() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LtGte
	}
	return nil
}

func (m *Timestamp) GetLteGt() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LteGt
	}
	return nil
}

func (m *Timestamp) GetLteGte() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LteGte
	}
	return nil
}

func (m *Timestamp) GetLtGtInv() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LtGtInv
	}
	return nil
}

func (m *Timestamp) GetLtGteInv() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LtGteInv
	}
	return nil
}

func (m *Timestamp) GetLteGtInv() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LteGtInv
	}
	return nil
}

func (m *Timestamp) GetLteGteInv() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LteGteInv
	}
	return nil
}

func (m *Timestamp) GetRequired() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Required
	}
	return nil
}

func (m *Timestamp) GetLtNow() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LtNow
	}
	return nil
}

func (m *Timestamp) GetGtNow() *google_protobuf2.Timestamp {
	if m != nil {
		return m.GtNow
	}
	return nil
}

func (m *Timestamp) GetLtNowWithin() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LtNowWithin
	}
	return nil
}

func (m *Timestamp) GetGtNowWithin() *google_protobuf2.Timestamp {
	if m != nil {
		return m.GtNowWithin
	}
	return nil
}

func (m *Timestamp) GetWithin() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Within
	}
	return nil
}

func (m *Timestamp) GetGogo1() *time.Time {
	if m != nil {
		return m.Gogo1
	}
	return nil
}

func (m *Timestamp) GetGogo2() google_protobuf2.Timestamp {
	if m != nil {
		return m.Gogo2
	}
	return google_protobuf2.Timestamp{}
}

func (m *Timestamp) GetGogo3() time.Time {
	if m != nil {
		return m.Gogo3
	}
	return time.Time{}
}

func (m *Timestamp) GetGogo4() google_protobuf2.Timestamp {
	if m != nil {
		return m.Gogo4
	}
	return google_protobuf2.Timestamp{}
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "tests.kitchensink.Timestamp")
}
func (m *Timestamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Timestamp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.None != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.None.Size()))
		n1, err := m.None.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Lt != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.Lt.Size()))
		n2, err := m.Lt.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Lte != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.Lte.Size()))
		n3, err := m.Lte.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Gt != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.Gt.Size()))
		n4, err := m.Gt.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Gte != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.Gte.Size()))
		n5, err := m.Gte.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.LtGt != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.LtGt.Size()))
		n6, err := m.LtGt.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.LtGte != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.LtGte.Size()))
		n7, err := m.LtGte.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.LteGt != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.LteGt.Size()))
		n8, err := m.LteGt.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.LteGte != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.LteGte.Size()))
		n9, err := m.LteGte.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.LtGtInv != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.LtGtInv.Size()))
		n10, err := m.LtGtInv.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	if m.LtGteInv != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.LtGteInv.Size()))
		n11, err := m.LtGteInv.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n11
	}
	if m.LteGtInv != nil {
		dAtA[i] = 0x62
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.LteGtInv.Size()))
		n12, err := m.LteGtInv.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n12
	}
	if m.LteGteInv != nil {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.LteGteInv.Size()))
		n13, err := m.LteGteInv.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n13
	}
	if m.Required != nil {
		dAtA[i] = 0x72
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.Required.Size()))
		n14, err := m.Required.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n14
	}
	if m.LtNow != nil {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.LtNow.Size()))
		n15, err := m.LtNow.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n15
	}
	if m.GtNow != nil {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.GtNow.Size()))
		n16, err := m.GtNow.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n16
	}
	if m.LtNowWithin != nil {
		dAtA[i] = 0x8a
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.LtNowWithin.Size()))
		n17, err := m.LtNowWithin.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n17
	}
	if m.GtNowWithin != nil {
		dAtA[i] = 0x92
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.GtNowWithin.Size()))
		n18, err := m.GtNowWithin.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n18
	}
	if m.Within != nil {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(m.Within.Size()))
		n19, err := m.Within.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n19
	}
	if m.Gogo1 != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintTimestamp(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.Gogo1)))
		n20, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.Gogo1, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n20
	}
	dAtA[i] = 0xaa
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintTimestamp(dAtA, i, uint64(m.Gogo2.Size()))
	n21, err := m.Gogo2.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n21
	dAtA[i] = 0xb2
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintTimestamp(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.Gogo3)))
	n22, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Gogo3, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n22
	dAtA[i] = 0xba
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintTimestamp(dAtA, i, uint64(m.Gogo4.Size()))
	n23, err := m.Gogo4.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n23
	return i, nil
}

func encodeFixed64Timestamp(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Timestamp(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTimestamp(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Timestamp) Size() (n int) {
	var l int
	_ = l
	if m.None != nil {
		l = m.None.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.Lt != nil {
		l = m.Lt.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.Lte != nil {
		l = m.Lte.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.Gt != nil {
		l = m.Gt.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.Gte != nil {
		l = m.Gte.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.LtGt != nil {
		l = m.LtGt.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.LtGte != nil {
		l = m.LtGte.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.LteGt != nil {
		l = m.LteGt.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.LteGte != nil {
		l = m.LteGte.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.LtGtInv != nil {
		l = m.LtGtInv.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.LtGteInv != nil {
		l = m.LtGteInv.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.LteGtInv != nil {
		l = m.LteGtInv.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.LteGteInv != nil {
		l = m.LteGteInv.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.Required != nil {
		l = m.Required.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.LtNow != nil {
		l = m.LtNow.Size()
		n += 1 + l + sovTimestamp(uint64(l))
	}
	if m.GtNow != nil {
		l = m.GtNow.Size()
		n += 2 + l + sovTimestamp(uint64(l))
	}
	if m.LtNowWithin != nil {
		l = m.LtNowWithin.Size()
		n += 2 + l + sovTimestamp(uint64(l))
	}
	if m.GtNowWithin != nil {
		l = m.GtNowWithin.Size()
		n += 2 + l + sovTimestamp(uint64(l))
	}
	if m.Within != nil {
		l = m.Within.Size()
		n += 2 + l + sovTimestamp(uint64(l))
	}
	if m.Gogo1 != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.Gogo1)
		n += 2 + l + sovTimestamp(uint64(l))
	}
	l = m.Gogo2.Size()
	n += 2 + l + sovTimestamp(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Gogo3)
	n += 2 + l + sovTimestamp(uint64(l))
	l = m.Gogo4.Size()
	n += 2 + l + sovTimestamp(uint64(l))
	return n
}

func sovTimestamp(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTimestamp(x uint64) (n int) {
	return sovTimestamp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Timestamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimestamp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field None", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.None == nil {
				m.None = &google_protobuf2.Timestamp{}
			}
			if err := m.None.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lt == nil {
				m.Lt = &google_protobuf2.Timestamp{}
			}
			if err := m.Lt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lte", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lte == nil {
				m.Lte = &google_protobuf2.Timestamp{}
			}
			if err := m.Lte.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Gt == nil {
				m.Gt = &google_protobuf2.Timestamp{}
			}
			if err := m.Gt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gte", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Gte == nil {
				m.Gte = &google_protobuf2.Timestamp{}
			}
			if err := m.Gte.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LtGt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LtGt == nil {
				m.LtGt = &google_protobuf2.Timestamp{}
			}
			if err := m.LtGt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LtGte", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LtGte == nil {
				m.LtGte = &google_protobuf2.Timestamp{}
			}
			if err := m.LtGte.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LteGt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LteGt == nil {
				m.LteGt = &google_protobuf2.Timestamp{}
			}
			if err := m.LteGt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LteGte", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LteGte == nil {
				m.LteGte = &google_protobuf2.Timestamp{}
			}
			if err := m.LteGte.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LtGtInv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LtGtInv == nil {
				m.LtGtInv = &google_protobuf2.Timestamp{}
			}
			if err := m.LtGtInv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LtGteInv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LtGteInv == nil {
				m.LtGteInv = &google_protobuf2.Timestamp{}
			}
			if err := m.LtGteInv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LteGtInv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LteGtInv == nil {
				m.LteGtInv = &google_protobuf2.Timestamp{}
			}
			if err := m.LteGtInv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LteGteInv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LteGteInv == nil {
				m.LteGteInv = &google_protobuf2.Timestamp{}
			}
			if err := m.LteGteInv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Required", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Required == nil {
				m.Required = &google_protobuf2.Timestamp{}
			}
			if err := m.Required.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LtNow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LtNow == nil {
				m.LtNow = &google_protobuf2.Timestamp{}
			}
			if err := m.LtNow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GtNow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GtNow == nil {
				m.GtNow = &google_protobuf2.Timestamp{}
			}
			if err := m.GtNow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LtNowWithin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LtNowWithin == nil {
				m.LtNowWithin = &google_protobuf2.Timestamp{}
			}
			if err := m.LtNowWithin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GtNowWithin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GtNowWithin == nil {
				m.GtNowWithin = &google_protobuf2.Timestamp{}
			}
			if err := m.GtNowWithin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Within", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Within == nil {
				m.Within = &google_protobuf2.Timestamp{}
			}
			if err := m.Within.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gogo1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Gogo1 == nil {
				m.Gogo1 = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.Gogo1, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gogo2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Gogo2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gogo3", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Gogo3, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 23:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gogo4", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Gogo4.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimestamp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimestamp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTimestamp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimestamp
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTimestamp
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTimestamp
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTimestamp(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTimestamp = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimestamp   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("timestamp.proto", fileDescriptorTimestamp) }

var fileDescriptorTimestamp = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd4, 0xcf, 0x8e, 0xd2, 0x40,
	0x1c, 0x07, 0x70, 0xa7, 0xd0, 0x32, 0x4c, 0xd9, 0x65, 0x77, 0x44, 0x19, 0x39, 0x80, 0xe1, 0x64,
	0x38, 0x94, 0x58, 0x3c, 0x10, 0x13, 0x13, 0xac, 0xd1, 0x55, 0x62, 0x34, 0x69, 0x4c, 0x3c, 0x12,
	0xd6, 0x1d, 0x67, 0x9b, 0xed, 0xb6, 0x2b, 0xcc, 0xc2, 0x6b, 0xf8, 0x12, 0xde, 0x3c, 0x98, 0x9e,
	0x8c, 0x27, 0x8f, 0x7b, 0xf4, 0x09, 0xd4, 0x70, 0xd3, 0xa7, 0x30, 0xbf, 0x99, 0x96, 0xb8, 0xba,
	0x49, 0xdb, 0xdb, 0x04, 0xe6, 0xfb, 0xf9, 0xfd, 0x69, 0x53, 0xd2, 0x94, 0xc1, 0x29, 0x5f, 0xca,
	0xf9, 0xe9, 0x99, 0x73, 0xb6, 0x88, 0x65, 0x4c, 0xf7, 0x25, 0x5f, 0xca, 0xa5, 0x73, 0x12, 0xc8,
	0x37, 0xc7, 0x3c, 0x5a, 0x06, 0xd1, 0x49, 0xa7, 0x27, 0xe2, 0x58, 0x84, 0x7c, 0xa8, 0x2e, 0x1c,
	0x9e, 0xbf, 0x1d, 0xfe, 0x93, 0xe9, 0xb4, 0x57, 0xf3, 0x30, 0x38, 0x9a, 0x4b, 0x3e, 0xcc, 0x0e,
	0xe9, 0x1f, 0x2d, 0x11, 0x8b, 0x58, 0x1d, 0x87, 0x70, 0xd2, 0xbf, 0xf6, 0x3f, 0x34, 0x48, 0xfd,
	0x55, 0x46, 0x50, 0x87, 0x54, 0xa3, 0x38, 0xe2, 0x0c, 0xdd, 0x46, 0x77, 0x6c, 0xb7, 0xe3, 0xe8,
	0x62, 0x4e, 0x56, 0xcc, 0xd9, 0xde, 0xf4, 0xd5, 0x3d, 0x3a, 0x26, 0x46, 0x28, 0x99, 0x91, 0x77,
	0xdb, 0x6b, 0x7c, 0xf9, 0xf5, 0xb5, 0x52, 0x4b, 0x50, 0xb5, 0x63, 0x60, 0xd3, 0x37, 0x42, 0x49,
	0xef, 0x93, 0x4a, 0x28, 0x39, 0xab, 0x14, 0x8f, 0xf6, 0x21, 0x0a, 0x21, 0xa8, 0x2a, 0x24, 0xab,
	0x16, 0x8f, 0x0e, 0x54, 0x55, 0xa1, 0xaa, 0x0a, 0xc9, 0x99, 0x59, 0x3c, 0xea, 0xaa, 0xaa, 0x42,
	0x72, 0xfa, 0x90, 0x98, 0xa1, 0x9c, 0x09, 0xc9, 0xac, 0xdc, 0xf4, 0x1e, 0xa4, 0xed, 0x04, 0xe1,
	0x8e, 0x81, 0x9b, 0x03, 0x03, 0x13, 0xbf, 0x1a, 0xca, 0x03, 0x49, 0x1f, 0x11, 0x4b, 0x11, 0x9c,
	0xd5, 0xca, 0x19, 0x2e, 0x18, 0x26, 0x18, 0x5c, 0x23, 0x1c, 0x1a, 0xc1, 0xc5, 0x91, 0x7e, 0xd6,
	0x88, 0x19, 0x4a, 0x7e, 0x20, 0xe9, 0x63, 0x52, 0xd3, 0x08, 0x67, 0xf5, 0x72, 0x8a, 0x6a, 0xc5,
	0x52, 0x0a, 0xa7, 0x4f, 0x49, 0x5d, 0x0d, 0x34, 0x0b, 0xa2, 0x15, 0x23, 0xa5, 0x66, 0x6a, 0x0d,
	0x0c, 0x7c, 0xcb, 0xaf, 0xc1, 0x4c, 0xcf, 0xa2, 0x15, 0x9d, 0x12, 0xa2, 0x57, 0xa3, 0x28, 0xbb,
	0x1c, 0xe5, 0x02, 0x85, 0xd5, 0x7a, 0xb6, 0x16, 0xcf, 0xda, 0x6a, 0x94, 0x9a, 0x4f, 0xb7, 0x85,
	0xd5, 0x7c, 0x60, 0x3d, 0x27, 0x76, 0xba, 0x28, 0x85, 0xed, 0x94, 0xc3, 0x54, 0x63, 0x75, 0xbd,
	0x2c, 0xd0, 0x3c, 0x82, 0x17, 0xfc, 0xdd, 0x79, 0xb0, 0xe0, 0x47, 0x6c, 0x37, 0x97, 0x22, 0x40,
	0x99, 0x09, 0x32, 0x30, 0xf2, 0xb7, 0x39, 0xfa, 0x40, 0xbd, 0x44, 0x51, 0xbc, 0x66, 0xcd, 0xe2,
	0xc2, 0x18, 0xc1, 0x93, 0x7f, 0x11, 0xaf, 0x21, 0x2e, 0x74, 0x7c, 0xaf, 0x78, 0x7c, 0x82, 0x7c,
	0x53, 0xa8, 0xf8, 0x4b, 0xb2, 0xa3, 0xab, 0xcf, 0xd6, 0x81, 0x3c, 0x0e, 0x22, 0xb6, 0x9f, 0xab,
	0x34, 0x41, 0x21, 0x09, 0xaa, 0x8d, 0xd1, 0xb4, 0x82, 0x3f, 0x1a, 0xbe, 0xad, 0x3a, 0x79, 0xad,
	0xf2, 0x00, 0x8a, 0x4b, 0x20, 0x2d, 0x0e, 0x4e, 0x00, 0xfc, 0xdc, 0xf6, 0x6d, 0xf1, 0x17, 0xe8,
	0x11, 0x2b, 0x95, 0xae, 0xe7, 0x4a, 0xbb, 0x20, 0xd5, 0x13, 0x64, 0x4d, 0xab, 0xf8, 0xd3, 0x6f,
	0xe4, 0xa7, 0x49, 0xea, 0x11, 0x13, 0xbe, 0x91, 0x77, 0x59, 0x2b, 0xff, 0x79, 0xbf, 0xff, 0xd1,
	0x43, 0x97, 0x3e, 0x6f, 0x3a, 0x9a, 0x19, 0x2e, 0xbb, 0x91, 0x6f, 0x5c, 0x7c, 0xef, 0x5d, 0xfb,
	0xdf, 0x70, 0xe9, 0x13, 0x6d, 0x8c, 0xd8, 0xcd, 0x5c, 0xa3, 0x05, 0xc6, 0xd5, 0xbd, 0x8c, 0xe8,
	0x44, 0x3b, 0xf7, 0x58, 0x3b, 0x7f, 0x25, 0x59, 0x2f, 0xe9, 0x8b, 0xa7, 0x83, 0x5e, 0xe3, 0x62,
	0xd3, 0x45, 0xdf, 0x36, 0x5d, 0xf4, 0x73, 0xd3, 0x45, 0x87, 0x96, 0x0a, 0x8e, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xd5, 0x9a, 0x96, 0xf5, 0xb2, 0x06, 0x00, 0x00,
}
