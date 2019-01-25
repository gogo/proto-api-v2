// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gogo.proto

package gogoproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/v2/types/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_GoprotoEnumPrefix = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62001,
	Name:          "gogoproto.goproto_enum_prefix",
	Tag:           "varint,62001,opt,name=goproto_enum_prefix",
	Filename:      "gogo.proto",
}

var E_GoprotoEnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62021,
	Name:          "gogoproto.goproto_enum_stringer",
	Tag:           "varint,62021,opt,name=goproto_enum_stringer",
	Filename:      "gogo.proto",
}

var E_EnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62022,
	Name:          "gogoproto.enum_stringer",
	Tag:           "varint,62022,opt,name=enum_stringer",
	Filename:      "gogo.proto",
}

var E_EnumCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         62023,
	Name:          "gogoproto.enum_customname",
	Tag:           "bytes,62023,opt,name=enum_customname",
	Filename:      "gogo.proto",
}

var E_Enumdecl = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62024,
	Name:          "gogoproto.enumdecl",
	Tag:           "varint,62024,opt,name=enumdecl",
	Filename:      "gogo.proto",
}

var E_EnumvalueCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         66001,
	Name:          "gogoproto.enumvalue_customname",
	Tag:           "bytes,66001,opt,name=enumvalue_customname",
	Filename:      "gogo.proto",
}

var E_GoprotoGettersAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63001,
	Name:          "gogoproto.goproto_getters_all",
	Tag:           "varint,63001,opt,name=goproto_getters_all",
	Filename:      "gogo.proto",
}

var E_GoprotoEnumPrefixAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63002,
	Name:          "gogoproto.goproto_enum_prefix_all",
	Tag:           "varint,63002,opt,name=goproto_enum_prefix_all",
	Filename:      "gogo.proto",
}

var E_GoprotoStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63003,
	Name:          "gogoproto.goproto_stringer_all",
	Tag:           "varint,63003,opt,name=goproto_stringer_all",
	Filename:      "gogo.proto",
}

var E_VerboseEqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63004,
	Name:          "gogoproto.verbose_equal_all",
	Tag:           "varint,63004,opt,name=verbose_equal_all",
	Filename:      "gogo.proto",
}

var E_FaceAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63005,
	Name:          "gogoproto.face_all",
	Tag:           "varint,63005,opt,name=face_all",
	Filename:      "gogo.proto",
}

var E_GostringAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63006,
	Name:          "gogoproto.gostring_all",
	Tag:           "varint,63006,opt,name=gostring_all",
	Filename:      "gogo.proto",
}

var E_PopulateAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63007,
	Name:          "gogoproto.populate_all",
	Tag:           "varint,63007,opt,name=populate_all",
	Filename:      "gogo.proto",
}

var E_StringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63008,
	Name:          "gogoproto.stringer_all",
	Tag:           "varint,63008,opt,name=stringer_all",
	Filename:      "gogo.proto",
}

var E_OnlyoneAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63009,
	Name:          "gogoproto.onlyone_all",
	Tag:           "varint,63009,opt,name=onlyone_all",
	Filename:      "gogo.proto",
}

var E_EqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63013,
	Name:          "gogoproto.equal_all",
	Tag:           "varint,63013,opt,name=equal_all",
	Filename:      "gogo.proto",
}

var E_DescriptionAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63014,
	Name:          "gogoproto.description_all",
	Tag:           "varint,63014,opt,name=description_all",
	Filename:      "gogo.proto",
}

var E_TestgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63015,
	Name:          "gogoproto.testgen_all",
	Tag:           "varint,63015,opt,name=testgen_all",
	Filename:      "gogo.proto",
}

var E_BenchgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63016,
	Name:          "gogoproto.benchgen_all",
	Tag:           "varint,63016,opt,name=benchgen_all",
	Filename:      "gogo.proto",
}

var E_MarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63017,
	Name:          "gogoproto.marshaler_all",
	Tag:           "varint,63017,opt,name=marshaler_all",
	Filename:      "gogo.proto",
}

var E_UnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63018,
	Name:          "gogoproto.unmarshaler_all",
	Tag:           "varint,63018,opt,name=unmarshaler_all",
	Filename:      "gogo.proto",
}

var E_StableMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63019,
	Name:          "gogoproto.stable_marshaler_all",
	Tag:           "varint,63019,opt,name=stable_marshaler_all",
	Filename:      "gogo.proto",
}

var E_SizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63020,
	Name:          "gogoproto.sizer_all",
	Tag:           "varint,63020,opt,name=sizer_all",
	Filename:      "gogo.proto",
}

var E_GoprotoEnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63021,
	Name:          "gogoproto.goproto_enum_stringer_all",
	Tag:           "varint,63021,opt,name=goproto_enum_stringer_all",
	Filename:      "gogo.proto",
}

var E_EnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63022,
	Name:          "gogoproto.enum_stringer_all",
	Tag:           "varint,63022,opt,name=enum_stringer_all",
	Filename:      "gogo.proto",
}

var E_UnsafeMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63023,
	Name:          "gogoproto.unsafe_marshaler_all",
	Tag:           "varint,63023,opt,name=unsafe_marshaler_all",
	Filename:      "gogo.proto",
}

var E_UnsafeUnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63024,
	Name:          "gogoproto.unsafe_unmarshaler_all",
	Tag:           "varint,63024,opt,name=unsafe_unmarshaler_all",
	Filename:      "gogo.proto",
}

var E_GoprotoExtensionsMapAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63025,
	Name:          "gogoproto.goproto_extensions_map_all",
	Tag:           "varint,63025,opt,name=goproto_extensions_map_all",
	Filename:      "gogo.proto",
}

var E_GoprotoUnrecognizedAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63026,
	Name:          "gogoproto.goproto_unrecognized_all",
	Tag:           "varint,63026,opt,name=goproto_unrecognized_all",
	Filename:      "gogo.proto",
}

var E_GogoprotoImport = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63027,
	Name:          "gogoproto.gogoproto_import",
	Tag:           "varint,63027,opt,name=gogoproto_import",
	Filename:      "gogo.proto",
}

var E_ProtosizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63028,
	Name:          "gogoproto.protosizer_all",
	Tag:           "varint,63028,opt,name=protosizer_all",
	Filename:      "gogo.proto",
}

var E_CompareAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63029,
	Name:          "gogoproto.compare_all",
	Tag:           "varint,63029,opt,name=compare_all",
	Filename:      "gogo.proto",
}

var E_TypedeclAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63030,
	Name:          "gogoproto.typedecl_all",
	Tag:           "varint,63030,opt,name=typedecl_all",
	Filename:      "gogo.proto",
}

var E_EnumdeclAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63031,
	Name:          "gogoproto.enumdecl_all",
	Tag:           "varint,63031,opt,name=enumdecl_all",
	Filename:      "gogo.proto",
}

var E_GoprotoRegistration = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63032,
	Name:          "gogoproto.goproto_registration",
	Tag:           "varint,63032,opt,name=goproto_registration",
	Filename:      "gogo.proto",
}

var E_MessagenameAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63033,
	Name:          "gogoproto.messagename_all",
	Tag:           "varint,63033,opt,name=messagename_all",
	Filename:      "gogo.proto",
}

var E_GoprotoSizecacheAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63034,
	Name:          "gogoproto.goproto_sizecache_all",
	Tag:           "varint,63034,opt,name=goproto_sizecache_all",
	Filename:      "gogo.proto",
}

var E_GoprotoUnkeyedAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63035,
	Name:          "gogoproto.goproto_unkeyed_all",
	Tag:           "varint,63035,opt,name=goproto_unkeyed_all",
	Filename:      "gogo.proto",
}

var E_GoprotoGetters = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64001,
	Name:          "gogoproto.goproto_getters",
	Tag:           "varint,64001,opt,name=goproto_getters",
	Filename:      "gogo.proto",
}

var E_GoprotoStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64003,
	Name:          "gogoproto.goproto_stringer",
	Tag:           "varint,64003,opt,name=goproto_stringer",
	Filename:      "gogo.proto",
}

var E_VerboseEqual = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64004,
	Name:          "gogoproto.verbose_equal",
	Tag:           "varint,64004,opt,name=verbose_equal",
	Filename:      "gogo.proto",
}

var E_Face = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64005,
	Name:          "gogoproto.face",
	Tag:           "varint,64005,opt,name=face",
	Filename:      "gogo.proto",
}

var E_Gostring = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64006,
	Name:          "gogoproto.gostring",
	Tag:           "varint,64006,opt,name=gostring",
	Filename:      "gogo.proto",
}

var E_Populate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64007,
	Name:          "gogoproto.populate",
	Tag:           "varint,64007,opt,name=populate",
	Filename:      "gogo.proto",
}

var E_Stringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         67008,
	Name:          "gogoproto.stringer",
	Tag:           "varint,67008,opt,name=stringer",
	Filename:      "gogo.proto",
}

var E_Onlyone = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64009,
	Name:          "gogoproto.onlyone",
	Tag:           "varint,64009,opt,name=onlyone",
	Filename:      "gogo.proto",
}

var E_Equal = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64013,
	Name:          "gogoproto.equal",
	Tag:           "varint,64013,opt,name=equal",
	Filename:      "gogo.proto",
}

var E_Description = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64014,
	Name:          "gogoproto.description",
	Tag:           "varint,64014,opt,name=description",
	Filename:      "gogo.proto",
}

var E_Testgen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64015,
	Name:          "gogoproto.testgen",
	Tag:           "varint,64015,opt,name=testgen",
	Filename:      "gogo.proto",
}

var E_Benchgen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64016,
	Name:          "gogoproto.benchgen",
	Tag:           "varint,64016,opt,name=benchgen",
	Filename:      "gogo.proto",
}

var E_Marshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64017,
	Name:          "gogoproto.marshaler",
	Tag:           "varint,64017,opt,name=marshaler",
	Filename:      "gogo.proto",
}

var E_Unmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64018,
	Name:          "gogoproto.unmarshaler",
	Tag:           "varint,64018,opt,name=unmarshaler",
	Filename:      "gogo.proto",
}

var E_StableMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64019,
	Name:          "gogoproto.stable_marshaler",
	Tag:           "varint,64019,opt,name=stable_marshaler",
	Filename:      "gogo.proto",
}

var E_Sizer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64020,
	Name:          "gogoproto.sizer",
	Tag:           "varint,64020,opt,name=sizer",
	Filename:      "gogo.proto",
}

var E_UnsafeMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64023,
	Name:          "gogoproto.unsafe_marshaler",
	Tag:           "varint,64023,opt,name=unsafe_marshaler",
	Filename:      "gogo.proto",
}

var E_UnsafeUnmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64024,
	Name:          "gogoproto.unsafe_unmarshaler",
	Tag:           "varint,64024,opt,name=unsafe_unmarshaler",
	Filename:      "gogo.proto",
}

var E_GoprotoExtensionsMap = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64025,
	Name:          "gogoproto.goproto_extensions_map",
	Tag:           "varint,64025,opt,name=goproto_extensions_map",
	Filename:      "gogo.proto",
}

var E_GoprotoUnrecognized = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64026,
	Name:          "gogoproto.goproto_unrecognized",
	Tag:           "varint,64026,opt,name=goproto_unrecognized",
	Filename:      "gogo.proto",
}

var E_Protosizer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64028,
	Name:          "gogoproto.protosizer",
	Tag:           "varint,64028,opt,name=protosizer",
	Filename:      "gogo.proto",
}

var E_Compare = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64029,
	Name:          "gogoproto.compare",
	Tag:           "varint,64029,opt,name=compare",
	Filename:      "gogo.proto",
}

var E_Typedecl = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64030,
	Name:          "gogoproto.typedecl",
	Tag:           "varint,64030,opt,name=typedecl",
	Filename:      "gogo.proto",
}

var E_Messagename = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64033,
	Name:          "gogoproto.messagename",
	Tag:           "varint,64033,opt,name=messagename",
	Filename:      "gogo.proto",
}

var E_GoprotoSizecache = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64034,
	Name:          "gogoproto.goproto_sizecache",
	Tag:           "varint,64034,opt,name=goproto_sizecache",
	Filename:      "gogo.proto",
}

var E_GoprotoUnkeyed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64035,
	Name:          "gogoproto.goproto_unkeyed",
	Tag:           "varint,64035,opt,name=goproto_unkeyed",
	Filename:      "gogo.proto",
}

var E_Nullable = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65001,
	Name:          "gogoproto.nullable",
	Tag:           "varint,65001,opt,name=nullable",
	Filename:      "gogo.proto",
}

var E_Embed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65002,
	Name:          "gogoproto.embed",
	Tag:           "varint,65002,opt,name=embed",
	Filename:      "gogo.proto",
}

var E_Customtype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65003,
	Name:          "gogoproto.customtype",
	Tag:           "bytes,65003,opt,name=customtype",
	Filename:      "gogo.proto",
}

var E_Customname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65004,
	Name:          "gogoproto.customname",
	Tag:           "bytes,65004,opt,name=customname",
	Filename:      "gogo.proto",
}

var E_Jsontag = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65005,
	Name:          "gogoproto.jsontag",
	Tag:           "bytes,65005,opt,name=jsontag",
	Filename:      "gogo.proto",
}

var E_Moretags = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65006,
	Name:          "gogoproto.moretags",
	Tag:           "bytes,65006,opt,name=moretags",
	Filename:      "gogo.proto",
}

var E_Casttype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65007,
	Name:          "gogoproto.casttype",
	Tag:           "bytes,65007,opt,name=casttype",
	Filename:      "gogo.proto",
}

var E_Castkey = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65008,
	Name:          "gogoproto.castkey",
	Tag:           "bytes,65008,opt,name=castkey",
	Filename:      "gogo.proto",
}

var E_Castvalue = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65009,
	Name:          "gogoproto.castvalue",
	Tag:           "bytes,65009,opt,name=castvalue",
	Filename:      "gogo.proto",
}

var E_Stdtime = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65010,
	Name:          "gogoproto.stdtime",
	Tag:           "varint,65010,opt,name=stdtime",
	Filename:      "gogo.proto",
}

var E_Stdduration = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65011,
	Name:          "gogoproto.stdduration",
	Tag:           "varint,65011,opt,name=stdduration",
	Filename:      "gogo.proto",
}

var E_Wktpointer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65012,
	Name:          "gogoproto.wktpointer",
	Tag:           "varint,65012,opt,name=wktpointer",
	Filename:      "gogo.proto",
}

func init() {
	proto.RegisterExtension(E_GoprotoEnumPrefix)
	proto.RegisterExtension(E_GoprotoEnumStringer)
	proto.RegisterExtension(E_EnumStringer)
	proto.RegisterExtension(E_EnumCustomname)
	proto.RegisterExtension(E_Enumdecl)
	proto.RegisterExtension(E_EnumvalueCustomname)
	proto.RegisterExtension(E_GoprotoGettersAll)
	proto.RegisterExtension(E_GoprotoEnumPrefixAll)
	proto.RegisterExtension(E_GoprotoStringerAll)
	proto.RegisterExtension(E_VerboseEqualAll)
	proto.RegisterExtension(E_FaceAll)
	proto.RegisterExtension(E_GostringAll)
	proto.RegisterExtension(E_PopulateAll)
	proto.RegisterExtension(E_StringerAll)
	proto.RegisterExtension(E_OnlyoneAll)
	proto.RegisterExtension(E_EqualAll)
	proto.RegisterExtension(E_DescriptionAll)
	proto.RegisterExtension(E_TestgenAll)
	proto.RegisterExtension(E_BenchgenAll)
	proto.RegisterExtension(E_MarshalerAll)
	proto.RegisterExtension(E_UnmarshalerAll)
	proto.RegisterExtension(E_StableMarshalerAll)
	proto.RegisterExtension(E_SizerAll)
	proto.RegisterExtension(E_GoprotoEnumStringerAll)
	proto.RegisterExtension(E_EnumStringerAll)
	proto.RegisterExtension(E_UnsafeMarshalerAll)
	proto.RegisterExtension(E_UnsafeUnmarshalerAll)
	proto.RegisterExtension(E_GoprotoExtensionsMapAll)
	proto.RegisterExtension(E_GoprotoUnrecognizedAll)
	proto.RegisterExtension(E_GogoprotoImport)
	proto.RegisterExtension(E_ProtosizerAll)
	proto.RegisterExtension(E_CompareAll)
	proto.RegisterExtension(E_TypedeclAll)
	proto.RegisterExtension(E_EnumdeclAll)
	proto.RegisterExtension(E_GoprotoRegistration)
	proto.RegisterExtension(E_MessagenameAll)
	proto.RegisterExtension(E_GoprotoSizecacheAll)
	proto.RegisterExtension(E_GoprotoUnkeyedAll)
	proto.RegisterExtension(E_GoprotoGetters)
	proto.RegisterExtension(E_GoprotoStringer)
	proto.RegisterExtension(E_VerboseEqual)
	proto.RegisterExtension(E_Face)
	proto.RegisterExtension(E_Gostring)
	proto.RegisterExtension(E_Populate)
	proto.RegisterExtension(E_Stringer)
	proto.RegisterExtension(E_Onlyone)
	proto.RegisterExtension(E_Equal)
	proto.RegisterExtension(E_Description)
	proto.RegisterExtension(E_Testgen)
	proto.RegisterExtension(E_Benchgen)
	proto.RegisterExtension(E_Marshaler)
	proto.RegisterExtension(E_Unmarshaler)
	proto.RegisterExtension(E_StableMarshaler)
	proto.RegisterExtension(E_Sizer)
	proto.RegisterExtension(E_UnsafeMarshaler)
	proto.RegisterExtension(E_UnsafeUnmarshaler)
	proto.RegisterExtension(E_GoprotoExtensionsMap)
	proto.RegisterExtension(E_GoprotoUnrecognized)
	proto.RegisterExtension(E_Protosizer)
	proto.RegisterExtension(E_Compare)
	proto.RegisterExtension(E_Typedecl)
	proto.RegisterExtension(E_Messagename)
	proto.RegisterExtension(E_GoprotoSizecache)
	proto.RegisterExtension(E_GoprotoUnkeyed)
	proto.RegisterExtension(E_Nullable)
	proto.RegisterExtension(E_Embed)
	proto.RegisterExtension(E_Customtype)
	proto.RegisterExtension(E_Customname)
	proto.RegisterExtension(E_Jsontag)
	proto.RegisterExtension(E_Moretags)
	proto.RegisterExtension(E_Casttype)
	proto.RegisterExtension(E_Castkey)
	proto.RegisterExtension(E_Castvalue)
	proto.RegisterExtension(E_Stdtime)
	proto.RegisterExtension(E_Stdduration)
	proto.RegisterExtension(E_Wktpointer)
}

func init() { proto.RegisterFile("gogo.proto", fileDescriptor_592445b5231bc2b9) }

var fileDescriptor_592445b5231bc2b9 = []byte{
	// 1324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x98, 0xc9, 0x6f, 0x1c, 0x45,
	0x17, 0xc0, 0xf5, 0xe9, 0x4b, 0x64, 0x4f, 0x79, 0x1f, 0x1b, 0x13, 0x22, 0x10, 0xcb, 0x89, 0x93,
	0x7d, 0x8a, 0x50, 0xca, 0x8a, 0x2c, 0xc7, 0x72, 0xac, 0x20, 0x1c, 0x8c, 0x13, 0x87, 0xed, 0x30,
	0xea, 0xe9, 0x29, 0xb7, 0x1b, 0x77, 0x77, 0x35, 0xdd, 0xd5, 0x21, 0xce, 0x0d, 0x85, 0x45, 0x08,
	0xb1, 0x23, 0x41, 0x42, 0x12, 0x08, 0x88, 0x7d, 0x0d, 0xfb, 0x72, 0xe1, 0xc2, 0x72, 0xe5, 0x7f,
	0xe0, 0x02, 0x98, 0xdd, 0x37, 0x5f, 0xd0, 0xeb, 0x7e, 0xaf, 0xa7, 0xa6, 0x3d, 0x52, 0xd5, 0xdc,
	0xda, 0x9e, 0xfa, 0xfd, 0xa6, 0xfa, 0xbd, 0xae, 0xf7, 0xde, 0x34, 0x63, 0x9e, 0xf4, 0xe4, 0x54,
	0x9c, 0x48, 0x25, 0xeb, 0x35, 0xb8, 0xce, 0x2f, 0xf7, 0x8f, 0xb6, 0x44, 0xea, 0x26, 0x7e, 0xac,
	0x64, 0x52, 0x7c, 0xc8, 0x8f, 0xb1, 0x71, 0xfc, 0xb0, 0x21, 0xa2, 0x2c, 0x6c, 0xc4, 0x89, 0x58,
	0xf3, 0x4f, 0xd7, 0xaf, 0x9d, 0xf2, 0xa4, 0xf4, 0x02, 0x51, 0xac, 0x6a, 0x66, 0x6b, 0x53, 0x0b,
	0x51, 0x16, 0xde, 0x1e, 0x2b, 0x5f, 0x46, 0xe9, 0xbe, 0x2b, 0x3f, 0xff, 0xff, 0x86, 0xff, 0xdd,
	0xdc, 0xbf, 0x32, 0x86, 0x28, 0x7c, 0xb6, 0x9c, 0x83, 0x7c, 0x85, 0x5d, 0xd5, 0xe1, 0x4b, 0x55,
	0xe2, 0x47, 0x9e, 0x48, 0x0c, 0xc6, 0xef, 0xd0, 0x38, 0xae, 0x19, 0x8f, 0x23, 0xca, 0xe7, 0xd9,
	0x50, 0x2f, 0xae, 0xef, 0xd1, 0x35, 0x28, 0x74, 0xc9, 0x22, 0x1b, 0xc9, 0x25, 0x6e, 0x96, 0x2a,
	0x19, 0x46, 0x4e, 0x28, 0x0c, 0x9a, 0x1f, 0x72, 0x4d, 0x6d, 0x65, 0x18, 0xb0, 0xf9, 0x92, 0xe2,
	0x9c, 0xf5, 0xc3, 0x7f, 0x5a, 0xc2, 0x0d, 0x0c, 0x86, 0x1f, 0x71, 0x23, 0xe5, 0x7a, 0x7e, 0x92,
	0x4d, 0xc0, 0xf5, 0x29, 0x27, 0xc8, 0x84, 0xbe, 0x93, 0x1b, 0xbb, 0x7a, 0x4e, 0xc2, 0x32, 0x92,
	0xfd, 0x74, 0x76, 0x4f, 0xbe, 0x9d, 0xf1, 0x52, 0xa0, 0xed, 0x49, 0xcb, 0xa2, 0x27, 0x94, 0x12,
	0x49, 0xda, 0x70, 0x82, 0x6e, 0xdb, 0x3b, 0xe2, 0x07, 0xa5, 0xf1, 0xdc, 0x56, 0x67, 0x16, 0x17,
	0x0b, 0x72, 0x2e, 0x08, 0xf8, 0x2a, 0xbb, 0xba, 0xcb, 0x53, 0x61, 0xe1, 0x3c, 0x8f, 0xce, 0x89,
	0x5d, 0x4f, 0x06, 0x68, 0x97, 0x19, 0xfd, 0xbf, 0xcc, 0xa5, 0x85, 0xf3, 0x25, 0x74, 0xd6, 0x91,
	0xa5, 0x94, 0x82, 0xf1, 0x56, 0x36, 0x76, 0x4a, 0x24, 0x4d, 0x99, 0x8a, 0x86, 0xb8, 0x3f, 0x73,
	0x02, 0x0b, 0xdd, 0x05, 0xd4, 0x8d, 0x20, 0xb8, 0x00, 0x1c, 0xb8, 0x0e, 0xb2, 0xfe, 0x35, 0xc7,
	0x15, 0x16, 0x8a, 0x8b, 0xa8, 0xe8, 0x83, 0xf5, 0x80, 0xce, 0xb1, 0x41, 0x4f, 0x16, 0xb7, 0x64,
	0x81, 0x5f, 0x42, 0x7c, 0x80, 0x18, 0x54, 0xc4, 0x32, 0xce, 0x02, 0x47, 0xd9, 0xec, 0xe0, 0x65,
	0x52, 0x10, 0x83, 0x8a, 0x1e, 0xc2, 0xfa, 0x0a, 0x29, 0x52, 0x2d, 0x9e, 0xb3, 0x6c, 0x40, 0x46,
	0xc1, 0xa6, 0x8c, 0x6c, 0x36, 0x71, 0x19, 0x0d, 0x0c, 0x11, 0x10, 0xcc, 0xb0, 0x9a, 0x6d, 0x22,
	0x5e, 0xdf, 0xa2, 0xe3, 0x41, 0x19, 0x58, 0x64, 0x23, 0x54, 0xa0, 0x7c, 0x19, 0x59, 0x28, 0xde,
	0x40, 0xc5, 0xb0, 0x86, 0xe1, 0x6d, 0x28, 0x91, 0x2a, 0x4f, 0xd8, 0x48, 0xde, 0xa4, 0xdb, 0x40,
	0x04, 0x43, 0xd9, 0x14, 0x91, 0xbb, 0x6e, 0x67, 0x78, 0x8b, 0x42, 0x49, 0x0c, 0x28, 0xe6, 0xd9,
	0x50, 0xe8, 0x24, 0xe9, 0xba, 0x13, 0x58, 0xa5, 0xe3, 0x6d, 0x74, 0x0c, 0x96, 0x10, 0x46, 0x24,
	0x8b, 0x7a, 0xd1, 0xbc, 0x43, 0x11, 0xd1, 0x30, 0x3c, 0x7a, 0xa9, 0x72, 0x9a, 0x81, 0x68, 0xf4,
	0x62, 0x7b, 0x97, 0x8e, 0x5e, 0xc1, 0x2e, 0xe9, 0xc6, 0x19, 0x56, 0x4b, 0xfd, 0x33, 0x56, 0x9a,
	0xf7, 0x28, 0xd3, 0x39, 0x00, 0xf0, 0xdd, 0xec, 0x9a, 0xae, 0x6d, 0xc2, 0x42, 0xf6, 0x3e, 0xca,
	0x26, 0xbb, 0xb4, 0x0a, 0x2c, 0x09, 0xbd, 0x2a, 0x3f, 0xa0, 0x92, 0x20, 0x2a, 0xae, 0x65, 0x36,
	0x91, 0x45, 0xa9, 0xb3, 0xd6, 0x5b, 0xd4, 0x3e, 0xa4, 0xa8, 0x15, 0x6c, 0x47, 0xd4, 0x4e, 0xb0,
	0x49, 0x34, 0xf6, 0x96, 0xd7, 0x8f, 0xa8, 0xb0, 0x16, 0xf4, 0x6a, 0x67, 0x76, 0xef, 0x65, 0xfb,
	0xcb, 0x70, 0x9e, 0x56, 0x22, 0x4a, 0x81, 0x69, 0x84, 0x4e, 0x6c, 0x61, 0xbe, 0x82, 0x66, 0xaa,
	0xf8, 0x0b, 0xa5, 0x60, 0xc9, 0x89, 0x41, 0x7e, 0x17, 0xdb, 0x47, 0xf2, 0x2c, 0x4a, 0x84, 0x2b,
	0xbd, 0xc8, 0x3f, 0x23, 0x5a, 0x16, 0xea, 0x8f, 0x2b, 0xa9, 0x5a, 0xd5, 0x70, 0x30, 0x1f, 0x65,
	0xa3, 0xe5, 0x6c, 0xd2, 0xf0, 0xc3, 0x58, 0x26, 0xca, 0x60, 0xfc, 0x84, 0x32, 0x55, 0x72, 0x47,
	0x73, 0x8c, 0x2f, 0xb0, 0xe1, 0xfc, 0x4f, 0xdb, 0x47, 0xf2, 0x53, 0x14, 0x0d, 0xb5, 0x29, 0x2c,
	0x1c, 0xae, 0x0c, 0x63, 0x27, 0xb1, 0xa9, 0x7f, 0x9f, 0x51, 0xe1, 0x40, 0x04, 0x0b, 0x87, 0xda,
	0x8c, 0x05, 0x74, 0x7b, 0x0b, 0xc3, 0xe7, 0x54, 0x38, 0x88, 0x41, 0x05, 0x0d, 0x0c, 0x16, 0x8a,
	0x2f, 0x48, 0x41, 0x0c, 0x28, 0xee, 0x68, 0x37, 0xda, 0x44, 0x78, 0x7e, 0xaa, 0x12, 0x07, 0x56,
	0x1b, 0x54, 0x5f, 0x6e, 0x75, 0x0e, 0x61, 0x2b, 0x1a, 0x0a, 0x95, 0x28, 0x14, 0x69, 0xea, 0x78,
	0x02, 0x26, 0x0e, 0x8b, 0x8d, 0x7d, 0x45, 0x95, 0x48, 0xc3, 0x60, 0x6f, 0xda, 0x84, 0x08, 0x61,
	0x77, 0x1d, 0x77, 0xdd, 0x46, 0xf7, 0x75, 0x65, 0x73, 0xc7, 0x89, 0x05, 0xa7, 0x36, 0xff, 0x64,
	0xd1, 0x86, 0xd8, 0xb4, 0x7a, 0x3a, 0xbf, 0xa9, 0xcc, 0x3f, 0xab, 0x05, 0x59, 0xd4, 0x90, 0x91,
	0xca, 0x3c, 0x55, 0xbf, 0x7e, 0x97, 0x6b, 0xa9, 0xb8, 0x2f, 0xd2, 0x3d, 0xb8, 0x8d, 0xf7, 0xdb,
	0x39, 0x4e, 0xf1, 0xdb, 0xe0, 0x21, 0xef, 0x1c, 0x7a, 0xcc, 0xb2, 0xb3, 0xdb, 0xe5, 0x73, 0xde,
	0x31, 0xf3, 0xf0, 0x23, 0x6c, 0xa8, 0x63, 0xe0, 0x31, 0xab, 0x1e, 0x42, 0xd5, 0xa0, 0x3e, 0xef,
	0xf0, 0x03, 0x6c, 0x0f, 0x0c, 0x2f, 0x66, 0xfc, 0x61, 0xc4, 0xf3, 0xe5, 0xfc, 0x10, 0xeb, 0xa7,
	0xa1, 0xc5, 0x8c, 0x3e, 0x82, 0x68, 0x89, 0x00, 0x4e, 0x03, 0x8b, 0x19, 0x7f, 0x94, 0x70, 0x42,
	0x00, 0xb7, 0x0f, 0xe1, 0xb7, 0x8f, 0xef, 0xc1, 0xa6, 0x43, 0xb1, 0x9b, 0x61, 0x7d, 0x38, 0xa9,
	0x98, 0xe9, 0xc7, 0xf0, 0xcb, 0x89, 0xe0, 0xb7, 0xb0, 0xbd, 0x96, 0x01, 0x7f, 0x02, 0xd1, 0x62,
	0x3d, 0x9f, 0x67, 0x03, 0xda, 0x74, 0x62, 0xc6, 0x9f, 0x44, 0x5c, 0xa7, 0x60, 0xeb, 0x38, 0x9d,
	0x98, 0x05, 0x4f, 0xd1, 0xd6, 0x91, 0x80, 0xb0, 0xd1, 0x60, 0x62, 0xa6, 0x9f, 0xa6, 0xa8, 0x13,
	0xc2, 0x67, 0x59, 0xad, 0x6c, 0x36, 0x66, 0xfe, 0x19, 0xe4, 0xdb, 0x0c, 0x44, 0x40, 0x6b, 0x76,
	0x66, 0xc5, 0xb3, 0x14, 0x01, 0x8d, 0x82, 0x63, 0x54, 0x1d, 0x60, 0xcc, 0xa6, 0xe7, 0xe8, 0x18,
	0x55, 0xe6, 0x17, 0xc8, 0x66, 0x5e, 0xf3, 0xcd, 0x8a, 0xe7, 0x29, 0x9b, 0xf9, 0x7a, 0xd8, 0x46,
	0x75, 0x22, 0x30, 0x3b, 0x5e, 0xa0, 0x6d, 0x54, 0x06, 0x02, 0xbe, 0xcc, 0xea, 0xbb, 0xa7, 0x01,
	0xb3, 0xef, 0x45, 0xf4, 0x8d, 0xed, 0x1a, 0x06, 0xf8, 0x9d, 0x6c, 0xb2, 0xfb, 0x24, 0x60, 0xb6,
	0x9e, 0xdb, 0xae, 0xfc, 0x76, 0xd3, 0x07, 0x01, 0x7e, 0xa2, 0xdd, 0x52, 0xf4, 0x29, 0xc0, 0xac,
	0x3d, 0xbf, 0xdd, 0x59, 0xb8, 0xf5, 0x21, 0x80, 0xcf, 0x31, 0xd6, 0x6e, 0xc0, 0x66, 0xd7, 0x05,
	0x74, 0x69, 0x10, 0x1c, 0x0d, 0xec, 0xbf, 0x66, 0xfe, 0x22, 0x1d, 0x0d, 0x24, 0xe0, 0x68, 0x50,
	0xeb, 0x35, 0xd3, 0x97, 0xe8, 0x68, 0x10, 0x02, 0x4f, 0xb6, 0xd6, 0xdd, 0xcc, 0x86, 0xcb, 0xf4,
	0x64, 0x6b, 0x14, 0x3f, 0xc6, 0xc6, 0x76, 0x35, 0x44, 0xb3, 0xea, 0x55, 0x54, 0x8d, 0x56, 0xfb,
	0xa1, 0xde, 0xbc, 0xb0, 0x19, 0x9a, 0x6d, 0xaf, 0x55, 0x9a, 0x17, 0xf6, 0x42, 0x3e, 0xc3, 0xfa,
	0xa3, 0x2c, 0x08, 0xe0, 0xf0, 0xd4, 0xaf, 0xeb, 0xd2, 0x4d, 0x45, 0xd0, 0x22, 0xc5, 0x2f, 0x3b,
	0x18, 0x1d, 0x02, 0xf8, 0x01, 0xb6, 0x57, 0x84, 0x4d, 0xd1, 0x32, 0x91, 0xbf, 0xee, 0x50, 0xc1,
	0x84, 0xd5, 0x7c, 0x96, 0xb1, 0xe2, 0xd5, 0x08, 0x84, 0xd9, 0xc4, 0xfe, 0xb6, 0x53, 0xbc, 0xa5,
	0xd1, 0x90, 0xb6, 0x20, 0x4f, 0x8a, 0x41, 0xb0, 0xd5, 0x29, 0xc8, 0x33, 0x72, 0x90, 0xf5, 0xdd,
	0x97, 0xca, 0x48, 0x39, 0x9e, 0x89, 0xfe, 0x1d, 0x69, 0x5a, 0x0f, 0x01, 0x0b, 0x65, 0x22, 0x94,
	0xe3, 0xa5, 0x26, 0xf6, 0x0f, 0x64, 0x4b, 0x00, 0x60, 0xd7, 0x49, 0x95, 0xcd, 0x7d, 0xff, 0x49,
	0x30, 0x01, 0xb0, 0x69, 0xb8, 0xde, 0x10, 0x9b, 0x26, 0xf6, 0x2f, 0xda, 0x34, 0xae, 0xe7, 0x87,
	0x58, 0x0d, 0x2e, 0xf3, 0xb7, 0x4a, 0x26, 0xf8, 0x6f, 0x84, 0xdb, 0x04, 0x7c, 0x73, 0xaa, 0x5a,
	0xca, 0x37, 0x07, 0xfb, 0x1f, 0xcc, 0x34, 0xad, 0xe7, 0x73, 0x6c, 0x20, 0x55, 0xad, 0x56, 0x86,
	0xf3, 0xa9, 0x01, 0xff, 0x77, 0xa7, 0x7c, 0x65, 0x51, 0x32, 0x90, 0xed, 0x07, 0x36, 0x54, 0x2c,
	0xfd, 0x48, 0x89, 0xc4, 0x64, 0xd8, 0x46, 0x83, 0x86, 0x1c, 0x5e, 0x60, 0xe3, 0xae, 0x0c, 0xab,
	0xdc, 0x61, 0xb6, 0x28, 0x17, 0xe5, 0x72, 0x5e, 0x67, 0xee, 0xb9, 0xc9, 0xf3, 0xd5, 0x7a, 0xd6,
	0x9c, 0x72, 0x65, 0x38, 0x0d, 0xbf, 0x3c, 0xa6, 0x69, 0xd5, 0x74, 0xf9, 0x3b, 0xe4, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa9, 0x28, 0xff, 0x8d, 0x73, 0x15, 0x00, 0x00,
}
