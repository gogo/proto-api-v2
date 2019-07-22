// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tag marshals and unmarshals the legacy struct tags as generated
// by historical versions of protoc-gen-go.
package tag

import (
	"reflect"
	"strconv"
	"strings"

	defval "github.com/gogo/protobuf/internal/encoding/defval"
	pref "google.golang.org/protobuf/reflect/protoreflect"
)

var byteType = reflect.TypeOf(byte(0))


// Marshal encodes the protoreflect.FieldDescriptor as a tag.
//
// The enumName must be provided if the kind is an enum.
// Historically, the formulation of the enum "name" was the proto package
// dot-concatenated with the generated Go identifier for the enum type.
// Depending on the context on how Marshal is called, there are different ways
// through which that information is determined. As such it is the caller's
// responsibility to provide a function to obtain that information.
func Marshal(fd pref.FieldDescriptor, enumName string) string {
	var tag []string
	switch fd.Kind() {
	case pref.BoolKind, pref.EnumKind, pref.Int32Kind, pref.Uint32Kind, pref.Int64Kind, pref.Uint64Kind:
		tag = append(tag, "varint")
	case pref.Sint32Kind:
		tag = append(tag, "zigzag32")
	case pref.Sint64Kind:
		tag = append(tag, "zigzag64")
	case pref.Sfixed32Kind, pref.Fixed32Kind, pref.FloatKind:
		tag = append(tag, "fixed32")
	case pref.Sfixed64Kind, pref.Fixed64Kind, pref.DoubleKind:
		tag = append(tag, "fixed64")
	case pref.StringKind, pref.BytesKind, pref.MessageKind:
		tag = append(tag, "bytes")
	case pref.GroupKind:
		tag = append(tag, "group")
	}
	tag = append(tag, strconv.Itoa(int(fd.Number())))
	switch fd.Cardinality() {
	case pref.Optional:
		tag = append(tag, "opt")
	case pref.Required:
		tag = append(tag, "req")
	case pref.Repeated:
		tag = append(tag, "rep")
	}
	if fd.IsPacked() {
		tag = append(tag, "packed")
	}
	// TODO: Weak fields?
	name := string(fd.Name())
	if fd.Kind() == pref.GroupKind {
		// The name of the FieldDescriptor for a group field is
		// lowercased. To find the original capitalization, we
		// look in the field's MessageType.
		name = string(fd.Message().Name())
	}
	tag = append(tag, "name="+name)
	if jsonName := fd.JSONName(); jsonName != "" && jsonName != name && !fd.IsExtension() {
		// TODO: The jsonName != name condition looks wrong.
		tag = append(tag, "json="+jsonName)
	}
	// The previous implementation does not tag extension fields as proto3,
	// even when the field is defined in a proto3 file. Match that behavior
	// for consistency.
	if fd.Syntax() == pref.Proto3 && !fd.IsExtension() {
		tag = append(tag, "proto3")
	}
	if fd.Kind() == pref.EnumKind && enumName != "" {
		tag = append(tag, "enum="+enumName)
	}
	if fd.ContainingOneof() != nil {
		tag = append(tag, "oneof")
	}
	// This must appear last in the tag, since commas in strings aren't escaped.
	if fd.HasDefault() {
		def, _ := defval.Marshal(fd.Default(), fd.DefaultEnumValue(), fd.Kind(), defval.GoTag)
		tag = append(tag, "def="+def)
	}
	return strings.Join(tag, ",")
}
