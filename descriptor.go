// Contains utilities for working with protobuf descriptors.

package main

import "google.golang.org/protobuf/types/descriptorpb"

// Tag number constants for the descriptor proto. These numbers are used in
// "path" lists for comments.
// See https://github.com/protocolbuffers/protobuf/blob/983dac24c34f07a278d45fdfd30348775d95d12a/src/google/protobuf/descriptor.proto#L892-L915
const (
	fileDescriptorMessageTypeTagNumber int32 = 4
	fileDescriptorEnumTypeTagNumber    int32 = 5
	fileDescriptorServiceTagNumber     int32 = 6

	descriptorFieldTagNumber      int32 = 2
	descriptorNestedTypeTagNumber int32 = 3
	descriptorEnumTypeTagNumber   int32 = 4

	enumDescriptorValueTagNumber int32 = 2

	serviceMethodTagNumber int32 = 2
)

// excludeMapEntries filters auto-generated map entry messages from the given
// message type list.
//
// Map entry messages only exist to support older code generators that don't
// support map fields.
func excludeMapEntries(messageTypes []*descriptorpb.DescriptorProto) []*descriptorpb.DescriptorProto {
	out := make([]*descriptorpb.DescriptorProto, 0, len(messageTypes))
	for _, m := range messageTypes {
		if m.GetOptions().GetMapEntry() {
			continue
		}
		out = append(out, m)
	}
	return out
}

// parentOneof returns the OneOfDescriptorProto for the oneof that the given
// field is a part of, given the field's parent message.
func parentOneof(message *descriptorpb.DescriptorProto, field *descriptorpb.FieldDescriptorProto) *descriptorpb.OneofDescriptorProto {
	return message.OneofDecl[field.GetOneofIndex()]
}
