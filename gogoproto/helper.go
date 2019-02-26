// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package gogoproto

import (
	"reflect"
	descriptorpb "github.com/golang/protobuf/v2/types/descriptor"
	proto "github.com/golang/protobuf/proto"
)

func IsNullable(fieldOptions proto.Message) bool {
	return GetBoolExtension(fieldOptions, E_Nullable, true)
}

func IsMarshaler(file *descriptorpb.FileDescriptorProto, messageOptions proto.Message) bool {
	return GetBoolExtension(messageOptions, E_Marshaler, GetBoolExtension(file.Options, E_MarshalerAll, false))
}

func IsUnmarshaler(file *descriptorpb.FileDescriptorProto, messageOptions proto.Message) bool {
	return GetBoolExtension(messageOptions, E_Marshaler, GetBoolExtension(file.Options, E_UnmarshalerAll, false))
}

func IsCastType(fieldOptions proto.Message) bool {
	typ := GetCastType(fieldOptions)
	if len(typ) > 0 {
		return true
	}
	return false
}

func GetBoolExtension(pb proto.Message, extension *proto.ExtensionDesc, ifnotset bool) bool {
	if reflect.ValueOf(pb).IsNil() {
		return ifnotset
	}
	value, err := proto.GetExtension(pb, extension)
	if err != nil {
		return ifnotset
	}
	if value == nil {
		return ifnotset
	}
	if value.(*bool) == nil {
		return ifnotset
	}
	return *(value.(*bool))
}

func GetCastType(fieldOptions proto.Message) string {
	if fieldOptions != nil {
		v, err := proto.GetExtension(fieldOptions, E_Casttype)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

