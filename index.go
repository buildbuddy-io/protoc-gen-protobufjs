package main

import (
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Index struct {
	// Map of file FS path to file descriptor proto.
	Files map[string]*descriptorpb.FileDescriptorProto
	// Map of fully-qualified message type name to descriptor proto.
	MessageTypes map[string]*descriptorpb.DescriptorProto
	// Map of fully-qualified enum type name to enum descriptor proto.
	EnumTypes map[string]*descriptorpb.EnumDescriptorProto
	// Map of fully-qualified message or enum type name to file descriptor proto.
	FilesByType map[string]*descriptorpb.FileDescriptorProto
}

func BuildIndex(req *pluginpb.CodeGeneratorRequest) *Index {
	idx := &Index{
		Files:        map[string]*descriptorpb.FileDescriptorProto{},
		MessageTypes: map[string]*descriptorpb.DescriptorProto{},
		EnumTypes:    map[string]*descriptorpb.EnumDescriptorProto{},
		FilesByType:  map[string]*descriptorpb.FileDescriptorProto{},
	}
	idx.build(req)
	return idx
}

func (idx *Index) build(req *pluginpb.CodeGeneratorRequest) {
	for _, f := range req.GetProtoFile() {
		idx.visitFile(f)
	}
}

func (idx *Index) visitFile(f *descriptorpb.FileDescriptorProto) {
	idx.Files[f.GetName()] = f

	path := "." + f.GetPackage()
	for _, m := range f.GetMessageType() {
		idx.visitMessageType(f, path, m)
	}
	for _, e := range f.GetEnumType() {
		idx.visitEnumType(f, path, e)
	}
}

func (idx *Index) visitMessageType(f *descriptorpb.FileDescriptorProto, path string, m *descriptorpb.DescriptorProto) {
	path = path + "." + m.GetName()
	idx.FilesByType[path] = f
	idx.MessageTypes[path] = m
	for _, m := range m.GetNestedType() {
		idx.visitMessageType(f, path, m)
	}
	for _, e := range m.GetEnumType() {
		idx.visitEnumType(f, path, e)
	}
}

func (idx *Index) visitEnumType(f *descriptorpb.FileDescriptorProto, path string, e *descriptorpb.EnumDescriptorProto) {
	path = path + "." + e.GetName()
	idx.FilesByType[path] = f
	idx.EnumTypes[path] = e
}
