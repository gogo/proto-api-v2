// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package defval marshals and unmarshals textual forms of default values.
//
// This package handles both the form historically used in Go struct field tags
// and also the form used by google.protobuf.FieldDescriptorProto.default_value
// since they differ in superficial ways.
package defval

import (
	"fmt"
	"math"
	"strconv"

	"errors"
	pref "google.golang.org/protobuf/reflect/protoreflect"
)

// Format is the serialization format used to represent the default value.
type Format int

const (
	_ Format = iota

	// Descriptor uses the serialization format that protoc uses with the
	// google.protobuf.FieldDescriptorProto.default_value field.
	Descriptor

	// GoTag uses the historical serialization format in Go struct field tags.
	GoTag)


// Marshal serializes v as the default string according to the given kind k.
// When specifying the Descriptor format for an enum kind, the associated
// enum value descriptor must be provided.
func Marshal(v pref.Value, ev pref.EnumValueDescriptor, k pref.Kind, f Format) (string, error) {
	switch k {
	case pref.BoolKind:
		if f == GoTag {
			if v.Bool() {
				return "1", nil
			} else {
				return "0", nil
			}
		} else {
			if v.Bool() {
				return "true", nil
			} else {
				return "false", nil
			}
		}
	case pref.EnumKind:
		if f == GoTag {
			return strconv.FormatInt(int64(v.Enum()), 10), nil
		} else {
			return string(ev.Name()), nil
		}
	case pref.Int32Kind, pref.Sint32Kind, pref.Sfixed32Kind, pref.Int64Kind, pref.Sint64Kind, pref.Sfixed64Kind:
		return strconv.FormatInt(v.Int(), 10), nil
	case pref.Uint32Kind, pref.Fixed32Kind, pref.Uint64Kind, pref.Fixed64Kind:
		return strconv.FormatUint(v.Uint(), 10), nil
	case pref.FloatKind, pref.DoubleKind:
		f := v.Float()
		switch {
		case math.IsInf(f, -1):
			return "-inf", nil
		case math.IsInf(f, +1):
			return "inf", nil
		case math.IsNaN(f):
			return "nan", nil
		default:
			if k == pref.FloatKind {
				return strconv.FormatFloat(f, 'g', -1, 32), nil
			} else {
				return strconv.FormatFloat(f, 'g', -1, 64), nil
			}
		}
	case pref.StringKind:
		// String values are serialized as is without any escaping.
		return v.String(), nil
	case pref.BytesKind:
		if s, ok := marshalBytes(v.Bytes()); ok {
			return s, nil
		}
	}
	return "", errors.New(fmt.Sprintf("could not format value for %v: %v", k, v))
}

// marshalBytes serializes bytes by using C escaping.
// To match the exact output of protoc, this is identical to the
// CEscape function in strutil.cc of the protoc source code.
func marshalBytes(b []byte) (string, bool) {
	var s []byte
	for _, c := range b {
		switch c {
		case '\n':
			s = append(s, `\n`...)
		case '\r':
			s = append(s, `\r`...)
		case '\t':
			s = append(s, `\t`...)
		case '"':
			s = append(s, `\"`...)
		case '\'':
			s = append(s, `\'`...)
		case '\\':
			s = append(s, `\\`...)
		default:
			if printableASCII := c >= 0x20 && c <= 0x7e; printableASCII {
				s = append(s, c)
			} else {
				s = append(s, fmt.Sprintf(`\%03o`, c)...)
			}
		}
	}
	return string(s), true
}
