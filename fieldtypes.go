// Contains utilities for working with protobuf field types.

package main

import "google.golang.org/protobuf/types/descriptorpb"

// isLong returns whether the given type is a 64-bit integer type.
func isLong(t descriptorpb.FieldDescriptorProto_Type) bool {
	switch t {
	case descriptorpb.FieldDescriptorProto_TYPE_INT64,
		descriptorpb.FieldDescriptorProto_TYPE_UINT64,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64,
		descriptorpb.FieldDescriptorProto_TYPE_SINT64:
		return true
	default:
		return false
	}
}

// isFloatingPoint returns whether the given type is float or double.
func isFloatingPoint(t descriptorpb.FieldDescriptorProto_Type) bool {
	return t == descriptorpb.FieldDescriptorProto_TYPE_DOUBLE || t == descriptorpb.FieldDescriptorProto_TYPE_FLOAT
}

// isUint returns whether the given type is an unsigned int type.
func isUint(t descriptorpb.FieldDescriptorProto_Type) bool {
	switch t {
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64,
		descriptorpb.FieldDescriptorProto_TYPE_UINT32:
		return true
	default:
		return false
	}
}
