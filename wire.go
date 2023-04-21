// Contains utilities for working with the proto wire format.

package main

import "google.golang.org/protobuf/types/descriptorpb"

const (
	// note: group wire types (3, 4) are unsupported

	varintWireType int32 = 0
	i64WireType    int32 = 1
	lenWireType    int32 = 2
	i32WireType    int32 = 5

	// wireTypeMask extracts the bottom 3 bits of a proto binary message record,
	// returning the wire type.
	// See https://protobuf.dev/programming-guides/encoding/#structure
	wireTypeMask = 0b111
)

func wireType(f *descriptorpb.FieldDescriptorProto) int32 {
	switch f.GetType() {
	case
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64,
		descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		return i64WireType
	case descriptorpb.FieldDescriptorProto_TYPE_STRING,
		descriptorpb.FieldDescriptorProto_TYPE_BYTES,
		descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		return lenWireType
	case descriptorpb.FieldDescriptorProto_TYPE_GROUP:
		fatalf("groups are currently unsupported")
		return -1 // unreachable
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED32,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32,
		descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		return i32WireType
	default:
		return varintWireType
	}
}

func isPackedField(f *descriptorpb.FieldDescriptorProto) bool {
	// If the "packed" option is set explicitly, return it.
	if f.Options != nil && f.Options.Packed != nil {
		return *f.Options.Packed
	}
	return isPackableField(f)
}

func isPackableField(f *descriptorpb.FieldDescriptorProto) bool {
	// Only repeated fields can be packed
	if f.GetLabel() != descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		return false
	}
	// Only varint, i32, and i64 wire types can be packed
	switch wireType(f) {
	case varintWireType, i32WireType, i64WireType:
		return true
	default:
		return false
	}
}
