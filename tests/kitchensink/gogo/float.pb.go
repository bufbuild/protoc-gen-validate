// Code generated by protoc-gen-gogo.
// source: float.proto
// DO NOT EDIT!

package tests_kitchensink

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Float struct {
	None      float32 `protobuf:"fixed32,1,opt,name=none,proto3" json:"none,omitempty"`
	Lt        float32 `protobuf:"fixed32,2,opt,name=lt,proto3" json:"lt,omitempty"`
	Lte       float32 `protobuf:"fixed32,3,opt,name=lte,proto3" json:"lte,omitempty"`
	Gt        float32 `protobuf:"fixed32,4,opt,name=gt,proto3" json:"gt,omitempty"`
	Gte       float32 `protobuf:"fixed32,5,opt,name=gte,proto3" json:"gte,omitempty"`
	LtGt      float32 `protobuf:"fixed32,6,opt,name=lt_gt,json=ltGt,proto3" json:"lt_gt,omitempty"`
	LtGte     float32 `protobuf:"fixed32,7,opt,name=lt_gte,json=ltGte,proto3" json:"lt_gte,omitempty"`
	LteGt     float32 `protobuf:"fixed32,8,opt,name=lte_gt,json=lteGt,proto3" json:"lte_gt,omitempty"`
	LteGte    float32 `protobuf:"fixed32,9,opt,name=lte_gte,json=lteGte,proto3" json:"lte_gte,omitempty"`
	LtGtInv   float32 `protobuf:"fixed32,10,opt,name=lt_gt_inv,json=ltGtInv,proto3" json:"lt_gt_inv,omitempty"`
	LtGteInv  float32 `protobuf:"fixed32,11,opt,name=lt_gte_inv,json=ltGteInv,proto3" json:"lt_gte_inv,omitempty"`
	LteGtInv  float32 `protobuf:"fixed32,12,opt,name=lte_gt_inv,json=lteGtInv,proto3" json:"lte_gt_inv,omitempty"`
	LteGteInv float32 `protobuf:"fixed32,13,opt,name=lte_gte_inv,json=lteGteInv,proto3" json:"lte_gte_inv,omitempty"`
	In        float32 `protobuf:"fixed32,14,opt,name=in,proto3" json:"in,omitempty"`
	NotIn     float32 `protobuf:"fixed32,15,opt,name=not_in,json=notIn,proto3" json:"not_in,omitempty"`
	Const     float32 `protobuf:"fixed32,16,opt,name=const,proto3" json:"const,omitempty"`
}

func (m *Float) Reset()                    { *m = Float{} }
func (m *Float) String() string            { return proto.CompactTextString(m) }
func (*Float) ProtoMessage()               {}
func (*Float) Descriptor() ([]byte, []int) { return fileDescriptorFloat, []int{0} }

func (m *Float) GetNone() float32 {
	if m != nil {
		return m.None
	}
	return 0
}

func (m *Float) GetLt() float32 {
	if m != nil {
		return m.Lt
	}
	return 0
}

func (m *Float) GetLte() float32 {
	if m != nil {
		return m.Lte
	}
	return 0
}

func (m *Float) GetGt() float32 {
	if m != nil {
		return m.Gt
	}
	return 0
}

func (m *Float) GetGte() float32 {
	if m != nil {
		return m.Gte
	}
	return 0
}

func (m *Float) GetLtGt() float32 {
	if m != nil {
		return m.LtGt
	}
	return 0
}

func (m *Float) GetLtGte() float32 {
	if m != nil {
		return m.LtGte
	}
	return 0
}

func (m *Float) GetLteGt() float32 {
	if m != nil {
		return m.LteGt
	}
	return 0
}

func (m *Float) GetLteGte() float32 {
	if m != nil {
		return m.LteGte
	}
	return 0
}

func (m *Float) GetLtGtInv() float32 {
	if m != nil {
		return m.LtGtInv
	}
	return 0
}

func (m *Float) GetLtGteInv() float32 {
	if m != nil {
		return m.LtGteInv
	}
	return 0
}

func (m *Float) GetLteGtInv() float32 {
	if m != nil {
		return m.LteGtInv
	}
	return 0
}

func (m *Float) GetLteGteInv() float32 {
	if m != nil {
		return m.LteGteInv
	}
	return 0
}

func (m *Float) GetIn() float32 {
	if m != nil {
		return m.In
	}
	return 0
}

func (m *Float) GetNotIn() float32 {
	if m != nil {
		return m.NotIn
	}
	return 0
}

func (m *Float) GetConst() float32 {
	if m != nil {
		return m.Const
	}
	return 0
}

func init() {
	proto.RegisterType((*Float)(nil), "tests.kitchensink.Float")
}
func (m *Float) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Float) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.None != 0 {
		dAtA[i] = 0xd
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.None))))
	}
	if m.Lt != 0 {
		dAtA[i] = 0x15
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.Lt))))
	}
	if m.Lte != 0 {
		dAtA[i] = 0x1d
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.Lte))))
	}
	if m.Gt != 0 {
		dAtA[i] = 0x25
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.Gt))))
	}
	if m.Gte != 0 {
		dAtA[i] = 0x2d
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.Gte))))
	}
	if m.LtGt != 0 {
		dAtA[i] = 0x35
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.LtGt))))
	}
	if m.LtGte != 0 {
		dAtA[i] = 0x3d
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.LtGte))))
	}
	if m.LteGt != 0 {
		dAtA[i] = 0x45
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.LteGt))))
	}
	if m.LteGte != 0 {
		dAtA[i] = 0x4d
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.LteGte))))
	}
	if m.LtGtInv != 0 {
		dAtA[i] = 0x55
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.LtGtInv))))
	}
	if m.LtGteInv != 0 {
		dAtA[i] = 0x5d
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.LtGteInv))))
	}
	if m.LteGtInv != 0 {
		dAtA[i] = 0x65
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.LteGtInv))))
	}
	if m.LteGteInv != 0 {
		dAtA[i] = 0x6d
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.LteGteInv))))
	}
	if m.In != 0 {
		dAtA[i] = 0x75
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.In))))
	}
	if m.NotIn != 0 {
		dAtA[i] = 0x7d
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.NotIn))))
	}
	if m.Const != 0 {
		dAtA[i] = 0x85
		i++
		dAtA[i] = 0x1
		i++
		i = encodeFixed32Float(dAtA, i, uint32(math.Float32bits(float32(m.Const))))
	}
	return i, nil
}

func encodeFixed64Float(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Float(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFloat(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Float) Size() (n int) {
	var l int
	_ = l
	if m.None != 0 {
		n += 5
	}
	if m.Lt != 0 {
		n += 5
	}
	if m.Lte != 0 {
		n += 5
	}
	if m.Gt != 0 {
		n += 5
	}
	if m.Gte != 0 {
		n += 5
	}
	if m.LtGt != 0 {
		n += 5
	}
	if m.LtGte != 0 {
		n += 5
	}
	if m.LteGt != 0 {
		n += 5
	}
	if m.LteGte != 0 {
		n += 5
	}
	if m.LtGtInv != 0 {
		n += 5
	}
	if m.LtGteInv != 0 {
		n += 5
	}
	if m.LteGtInv != 0 {
		n += 5
	}
	if m.LteGteInv != 0 {
		n += 5
	}
	if m.In != 0 {
		n += 5
	}
	if m.NotIn != 0 {
		n += 5
	}
	if m.Const != 0 {
		n += 6
	}
	return n
}

func sovFloat(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFloat(x uint64) (n int) {
	return sovFloat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Float) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFloat
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
			return fmt.Errorf("proto: Float: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Float: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field None", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.None = float32(math.Float32frombits(v))
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lt", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.Lt = float32(math.Float32frombits(v))
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lte", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.Lte = float32(math.Float32frombits(v))
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gt", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.Gt = float32(math.Float32frombits(v))
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gte", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.Gte = float32(math.Float32frombits(v))
		case 6:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field LtGt", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.LtGt = float32(math.Float32frombits(v))
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field LtGte", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.LtGte = float32(math.Float32frombits(v))
		case 8:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field LteGt", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.LteGt = float32(math.Float32frombits(v))
		case 9:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field LteGte", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.LteGte = float32(math.Float32frombits(v))
		case 10:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field LtGtInv", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.LtGtInv = float32(math.Float32frombits(v))
		case 11:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field LtGteInv", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.LtGteInv = float32(math.Float32frombits(v))
		case 12:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field LteGtInv", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.LteGtInv = float32(math.Float32frombits(v))
		case 13:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field LteGteInv", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.LteGteInv = float32(math.Float32frombits(v))
		case 14:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field In", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.In = float32(math.Float32frombits(v))
		case 15:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotIn", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.NotIn = float32(math.Float32frombits(v))
		case 16:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Const", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.Const = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipFloat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFloat
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
func skipFloat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFloat
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
					return 0, ErrIntOverflowFloat
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
					return 0, ErrIntOverflowFloat
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
				return 0, ErrInvalidLengthFloat
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFloat
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
				next, err := skipFloat(dAtA[start:])
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
	ErrInvalidLengthFloat = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFloat   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("float.proto", fileDescriptorFloat) }

var fileDescriptorFloat = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x27, 0x69, 0x26, 0x69, 0x4f, 0x73, 0xd5, 0x3b, 0x50, 0x1c, 0xc4, 0x86, 0x5a, 0xa8,
	0x14, 0x21, 0x2d, 0x52, 0xba, 0x2c, 0x0d, 0x5d, 0x58, 0xba, 0xcd, 0x0b, 0xd4, 0x5a, 0xa7, 0x31,
	0x74, 0x98, 0x88, 0x39, 0xf4, 0x39, 0xdc, 0xea, 0x56, 0x10, 0x7c, 0x05, 0x57, 0x5d, 0xba, 0xf4,
	0x11, 0xa4, 0x3b, 0xf7, 0x3e, 0x80, 0xcc, 0x8c, 0xb5, 0x0c, 0xd9, 0x0d, 0xf9, 0xbf, 0xef, 0x9c,
	0x3f, 0xc9, 0x40, 0xf7, 0x20, 0xab, 0x1d, 0x4e, 0xde, 0x7f, 0xa8, 0xb0, 0x62, 0xf7, 0x28, 0x6a,
	0xac, 0x27, 0xc7, 0x12, 0xf7, 0xef, 0x84, 0xaa, 0x4b, 0x75, 0x7c, 0xf2, 0xf8, 0xb4, 0x93, 0xe5,
	0xdb, 0x1d, 0x8a, 0xe9, 0xf5, 0x60, 0xd9, 0xe1, 0x9f, 0x00, 0xe8, 0x2b, 0xed, 0x32, 0x06, 0x81,
	0xaa, 0x94, 0xe0, 0xde, 0xc0, 0x1b, 0xfb, 0xb9, 0x39, 0xb3, 0xa7, 0xe0, 0x4b, 0xe4, 0xbe, 0x7e,
	0xb2, 0x8a, 0xbf, 0xff, 0x3e, 0xb7, 0x22, 0xa0, 0x3d, 0x42, 0xc8, 0x32, 0xf7, 0x25, 0xb2, 0x04,
	0x5a, 0x12, 0x05, 0x6f, 0x39, 0x71, 0xdf, 0xc4, 0x3a, 0xd0, 0x76, 0x81, 0x3c, 0x70, 0xe2, 0x91,
	0xb5, 0x0b, 0x63, 0x17, 0x28, 0x38, 0x75, 0xe2, 0xd4, 0xda, 0x05, 0x0a, 0xf6, 0x1c, 0xa8, 0xc4,
	0x6d, 0x81, 0x3c, 0x34, 0xc4, 0xbd, 0x26, 0x62, 0x80, 0xde, 0x6c, 0xf6, 0x25, 0x1b, 0x11, 0x72,
	0x5e, 0xe6, 0x81, 0xc4, 0x35, 0xb2, 0x31, 0x84, 0x86, 0x13, 0x3c, 0x6a, 0x82, 0xa9, 0x01, 0xa9,
	0x06, 0x85, 0x25, 0x85, 0x1e, 0xd9, 0x76, 0xc8, 0xfe, 0x6d, 0x24, 0x95, 0x28, 0xd6, 0xc8, 0x5e,
	0x40, 0x64, 0x49, 0xc1, 0x3b, 0x4d, 0xd4, 0x0e, 0x0d, 0x0d, 0x2a, 0x58, 0x0a, 0x1d, 0xb3, 0x7f,
	0x5b, 0xaa, 0x13, 0x07, 0xb7, 0x02, 0x21, 0x83, 0x6c, 0x74, 0x38, 0x7c, 0xcd, 0xf2, 0x48, 0x57,
	0xd8, 0xa8, 0x13, 0x9b, 0x02, 0xd8, 0xba, 0x86, 0xef, 0x36, 0xf9, 0xd4, 0xf0, 0x6d, 0x53, 0xf9,
	0xbf, 0x20, 0xae, 0x0b, 0x62, 0xb7, 0xce, 0x6d, 0x41, 0xdb, 0xd4, 0xd1, 0xc2, 0x4b, 0xe8, 0xfe,
	0x2b, 0x6f, 0x8c, 0xbb, 0xa6, 0x61, 0x57, 0x74, 0xec, 0x0b, 0x68, 0xe5, 0x19, 0xf8, 0xa5, 0xe2,
	0x0f, 0x1c, 0x72, 0x4e, 0xc8, 0xeb, 0x6c, 0xfe, 0xf9, 0xd3, 0xb7, 0x2c, 0xf7, 0x4b, 0xa5, 0x3f,
	0x9e, 0xaa, 0x74, 0x07, 0xfe, 0xd0, 0xc1, 0x16, 0x1a, 0x5b, 0x18, 0x8c, 0xaa, 0x0a, 0x37, 0x8a,
	0x0d, 0x81, 0xee, 0x2b, 0x55, 0x23, 0x7f, 0xe4, 0xfc, 0xda, 0x3b, 0x42, 0x3e, 0x66, 0xb9, 0x8d,
	0x56, 0xf1, 0x8f, 0x4b, 0xe2, 0xfd, 0xbc, 0x24, 0xde, 0xaf, 0x4b, 0xe2, 0xbd, 0x09, 0xcd, 0x5d,
	0x9c, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x2e, 0x2e, 0x6a, 0xc6, 0x02, 0x00, 0x00,
}
