// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package msgfmt implements a text marshaler combining the desirable features
// of both the JSON and proto text formats.
// It is optimized for human readability and has no associated deserializer.
package msgfmt

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/detectknown"
	"google.golang.org/protobuf/internal/detrand"
	"google.golang.org/protobuf/internal/mapsort"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Format returns a formatted string for the message.
func Format(m proto.Message) string {
	return string(appendMessage(nil, m.ProtoReflect()))
}

func appendValue(b []byte, v protoreflect.Value, fd protoreflect.FieldDescriptor) []byte {
	switch v := v.Interface().(type) {
	case bool, int32, int64, uint32, uint64, float32, float64:
		return append(b, fmt.Sprint(v)...)
	case string:
		return append(b, strconv.Quote(string(v))...)
	case []byte:
		return append(b, strconv.Quote(string(v))...)
	case protoreflect.EnumNumber:
		return appendEnum(b, v, fd.Enum())
	case protoreflect.Message:
		return appendMessage(b, v)
	case protoreflect.List:
		return appendList(b, v, fd)
	case protoreflect.Map:
		return appendMap(b, v, fd)
	default:
		panic(fmt.Sprintf("invalid type: %T", v))
	}
}

func appendEnum(b []byte, v protoreflect.EnumNumber, ed protoreflect.EnumDescriptor) []byte {
	if ev := ed.Values().ByNumber(v); ev != nil {
		return append(b, ev.Name()...)
	}
	return strconv.AppendInt(b, int64(v), 10)
}

func appendMessage(b []byte, m protoreflect.Message) []byte {
	if b2 := appendKnownMessage(b, m); b2 != nil {
		return b2
	}

	var fds []protoreflect.FieldDescriptor
	m.Range(func(fd protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
		fds = append(fds, fd)
		return true
	})
	sort.Slice(fds, func(i, j int) bool {
		fdi, fdj := fds[i], fds[j]
		switch {
		case !fdi.IsExtension() && !fdj.IsExtension():
			return fdi.Index() < fdj.Index()
		case fdi.IsExtension() && fdj.IsExtension():
			return fdi.FullName() < fdj.FullName()
		default:
			return !fdi.IsExtension() && fdj.IsExtension()
		}
	})

	b = append(b, '{')
	for _, fd := range fds {
		k := string(fd.Name())
		if fd.IsExtension() {
			k = string("[" + fd.FullName() + "]")
		}

		b = append(b, k...)
		b = append(b, ':')
		b = appendValue(b, m.Get(fd), fd)
		b = append(b, delim()...)
	}
	b = appendUnknown(b, m.GetUnknown())
	b = bytes.TrimRight(b, delim())
	b = append(b, '}')
	return b
}

var protocmpMessageType = reflect.TypeOf(map[string]interface{}(nil))

func appendKnownMessage(b []byte, m protoreflect.Message) []byte {
	md := m.Descriptor()
	fds := md.Fields()
	switch detectknown.Which(md.FullName()) {
	case detectknown.AnyProto:
		var msgVal protoreflect.Message
		url := m.Get(fds.ByName("type_url")).String()
		if v := reflect.ValueOf(m); v.Type().ConvertibleTo(protocmpMessageType) {
			// For protocmp.Message, directly obtain the sub-message value
			// which is stored in structured form, rather than as raw bytes.
			m2 := v.Convert(protocmpMessageType).Interface().(map[string]interface{})
			v, ok := m2["value"].(proto.Message)
			if !ok {
				return nil
			}
			msgVal = v.ProtoReflect()
		} else {
			val := m.Get(fds.ByName("value")).Bytes()
			mt, err := protoregistry.GlobalTypes.FindMessageByURL(url)
			if err != nil {
				return nil
			}
			msgVal = mt.New()
			err = proto.UnmarshalOptions{AllowPartial: true}.Unmarshal(val, msgVal.Interface())
			if err != nil {
				return nil
			}
		}

		b = append(b, '{')
		b = append(b, "["+url+"]"...)
		b = append(b, ':')
		b = appendMessage(b, msgVal)
		b = append(b, '}')
		return b

	case detectknown.TimestampProto:
		secs := m.Get(fds.ByName("seconds")).Int()
		nanos := m.Get(fds.ByName("nanos")).Int()
		if nanos < 0 || nanos >= 1e9 {
			return nil
		}
		t := time.Unix(secs, nanos).UTC()
		x := t.Format("2006-01-02T15:04:05.000000000") // RFC 3339
		x = strings.TrimSuffix(x, "000")
		x = strings.TrimSuffix(x, "000")
		x = strings.TrimSuffix(x, ".000")
		return append(b, x+"Z"...)

	case detectknown.DurationProto:
		secs := m.Get(fds.ByName("seconds")).Int()
		nanos := m.Get(fds.ByName("nanos")).Int()
		if nanos <= -1e9 || nanos >= 1e9 || (secs > 0 && nanos < 0) || (secs < 0 && nanos > 0) {
			return nil
		}
		x := fmt.Sprintf("%d.%09d", secs, int64(math.Abs(float64(nanos))))
		x = strings.TrimSuffix(x, "000")
		x = strings.TrimSuffix(x, "000")
		x = strings.TrimSuffix(x, ".000")
		return append(b, x+"s"...)

	case detectknown.WrappersProto:
		fd := fds.ByName("value")
		return appendValue(b, m.Get(fd), fd)
	}

	return nil
}

func appendUnknown(b []byte, raw protoreflect.RawFields) []byte {
	rs := make(map[protoreflect.FieldNumber][]protoreflect.RawFields)
	for len(raw) > 0 {
		num, _, n := protowire.ConsumeField(raw)
		rs[num] = append(rs[num], raw[:n])
		raw = raw[n:]
	}

	var ns []protoreflect.FieldNumber
	for n := range rs {
		ns = append(ns, n)
	}
	sort.Slice(ns, func(i, j int) bool { return ns[i] < ns[j] })

	for _, n := range ns {
		var leftBracket, rightBracket string
		if len(rs[n]) > 1 {
			leftBracket, rightBracket = "[", "]"
		}

		b = strconv.AppendInt(b, int64(n), 10)
		b = append(b, ':')
		b = append(b, leftBracket...)
		for _, r := range rs[n] {
			num, typ, n := protowire.ConsumeTag(r)
			r = r[n:]
			switch typ {
			case protowire.VarintType:
				v, _ := protowire.ConsumeVarint(r)
				b = strconv.AppendInt(b, int64(v), 10)
			case protowire.Fixed32Type:
				v, _ := protowire.ConsumeFixed32(r)
				b = append(b, fmt.Sprintf("0x%08x", v)...)
			case protowire.Fixed64Type:
				v, _ := protowire.ConsumeFixed64(r)
				b = append(b, fmt.Sprintf("0x%016x", v)...)
			case protowire.BytesType:
				v, _ := protowire.ConsumeBytes(r)
				b = strconv.AppendQuote(b, string(v))
			case protowire.StartGroupType:
				v, _ := protowire.ConsumeGroup(num, r)
				b = append(b, '{')
				b = appendUnknown(b, v)
				b = bytes.TrimRight(b, delim())
				b = append(b, '}')
			default:
				panic(fmt.Sprintf("invalid type: %v", typ))
			}
			b = append(b, delim()...)
		}
		b = bytes.TrimRight(b, delim())
		b = append(b, rightBracket...)
		b = append(b, delim()...)
	}
	return b
}

func appendList(b []byte, v protoreflect.List, fd protoreflect.FieldDescriptor) []byte {
	b = append(b, '[')
	for i := 0; i < v.Len(); i++ {
		b = appendValue(b, v.Get(i), fd)
		b = append(b, delim()...)
	}
	b = bytes.TrimRight(b, delim())
	b = append(b, ']')
	return b
}

func appendMap(b []byte, v protoreflect.Map, fd protoreflect.FieldDescriptor) []byte {
	var ks []protoreflect.MapKey
	mapsort.Range(v, fd.MapKey().Kind(), func(k protoreflect.MapKey, _ protoreflect.Value) bool {
		ks = append(ks, k)
		return true
	})

	b = append(b, '{')
	for _, k := range ks {
		b = appendValue(b, k.Value(), fd.MapKey())
		b = append(b, ':')
		b = appendValue(b, v.Get(k), fd.MapValue())
		b = append(b, delim()...)
	}
	b = bytes.TrimRight(b, delim())
	b = append(b, '}')
	return b
}

func delim() string {
	// Deliberately introduce instability into the message string to
	// discourage users from depending on it.
	if detrand.Bool() {
		return "  "
	}
	return ", "
}
