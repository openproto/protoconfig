// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: protoconfig/v1/extensions.proto

// Proto Config Extensions Format 1.0.
//
// This defines extensions to protocol-buffer-compatible language that is build on proto version 3 with protobuf custom options.
//
// ProtoConfig Extensions allows software authors to define metadata about required configuration like default values or field
// optionality etc.
//
// Extension can be used as natively described in https://developers.google.com/protocol-buffers/docs/proto#customoptions
//
// Below extensions are not mandatory to use while defining your configuration, but mandatory to be implemented by
// ProtoConfig 1.0 compatible parsers and generators.

package protoconfig

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// / Metadata is an Message option that when put on message indicates an entry point for certain configuration.
// / One `Configuration Proto Definition` can have many structs marked as this option.
// / TODO(bwplotka): Make it non pointers (in Go).
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// name represents name of annotated configuration entry point.
	/// It's recommended to use executable name here.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	/// version is the semantic version of the annotated configuration.
	Version     string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	/// delivery_mechanism optionally specifies the delivery that is implemented by configurable that consumes this
	/// configuration. This allows `Configurator` to discover how to pass the `Encoded Configuration Message`
	/// without extra information outside of definition.
	// TODO(bwplotka): This might be blocking reusability. Rethink?
	//
	// Types that are assignable to DeliveryMechanism:
	//	*Metadata_StdinDelivery
	//	*Metadata_FlagDelivery
	DeliveryMechanism isMetadata_DeliveryMechanism `protobuf_oneof:"delivery_mechanism"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoconfig_v1_extensions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_protoconfig_v1_extensions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_protoconfig_v1_extensions_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Metadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *Metadata) GetDeliveryMechanism() isMetadata_DeliveryMechanism {
	if m != nil {
		return m.DeliveryMechanism
	}
	return nil
}

func (x *Metadata) GetStdinDelivery() *StdinDelivery {
	if x, ok := x.GetDeliveryMechanism().(*Metadata_StdinDelivery); ok {
		return x.StdinDelivery
	}
	return nil
}

func (x *Metadata) GetFlagDelivery() *FlagDelivery {
	if x, ok := x.GetDeliveryMechanism().(*Metadata_FlagDelivery); ok {
		return x.FlagDelivery
	}
	return nil
}

type isMetadata_DeliveryMechanism interface {
	isMetadata_DeliveryMechanism()
}

type Metadata_StdinDelivery struct {
	StdinDelivery *StdinDelivery `protobuf:"bytes,101,opt,name=stdin_delivery,json=stdinDelivery,proto3,oneof"`
}

type Metadata_FlagDelivery struct {
	FlagDelivery *FlagDelivery `protobuf:"bytes,102,opt,name=flag_delivery,json=flagDelivery,proto3,oneof"`
}

func (*Metadata_StdinDelivery) isMetadata_DeliveryMechanism() {}

func (*Metadata_FlagDelivery) isMetadata_DeliveryMechanism() {}

type StdinDelivery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StdinDelivery) Reset() {
	*x = StdinDelivery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoconfig_v1_extensions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StdinDelivery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StdinDelivery) ProtoMessage() {}

func (x *StdinDelivery) ProtoReflect() protoreflect.Message {
	mi := &file_protoconfig_v1_extensions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StdinDelivery.ProtoReflect.Descriptor instead.
func (*StdinDelivery) Descriptor() ([]byte, []int) {
	return file_protoconfig_v1_extensions_proto_rawDescGZIP(), []int{1}
}

type FlagDelivery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// name represents custom flag name (including `-` if any) for a flag that consumes bytes of `Encoded Configuration Message`
	// ProtoConfig 1.0 recommends `--protoconfigv1` name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FlagDelivery) Reset() {
	*x = FlagDelivery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoconfig_v1_extensions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagDelivery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagDelivery) ProtoMessage() {}

func (x *FlagDelivery) ProtoReflect() protoreflect.Message {
	mi := &file_protoconfig_v1_extensions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagDelivery.ProtoReflect.Descriptor instead.
func (*FlagDelivery) Descriptor() ([]byte, []int) {
	return file_protoconfig_v1_extensions_proto_rawDescGZIP(), []int{2}
}

func (x *FlagDelivery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var file_protoconfig_v1_extensions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*Metadata)(nil),
		Field:         5000,
		Name:          "protoconfig.v1.metadata",
		Tag:           "bytes,5000,opt,name=metadata",
		Filename:      "protoconfig/v1/extensions.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         5000,
		Name:          "protoconfig.v1.default",
		Tag:           "bytes,5000,opt,name=default",
		Filename:      "protoconfig/v1/extensions.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         5001,
		Name:          "protoconfig.v1.hidden",
		Tag:           "varint,5001,opt,name=hidden",
		Filename:      "protoconfig/v1/extensions.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         5002,
		Name:          "protoconfig.v1.required",
		Tag:           "varint,5002,opt,name=required",
		Filename:      "protoconfig/v1/extensions.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         5003,
		Name:          "protoconfig.v1.experimental",
		Tag:           "varint,5003,opt,name=experimental",
		Filename:      "protoconfig/v1/extensions.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	/// metadata represents
	//
	// optional protoconfig.v1.Metadata metadata = 5000;
	E_Metadata = &file_protoconfig_v1_extensions_proto_extTypes[0]
)

// Extension fields to descriptor.FieldOptions.
var (
	/// default represents an option that sets a default value for the field.
	//
	// optional string default = 5000;
	E_Default = &file_protoconfig_v1_extensions_proto_extTypes[1]
	/// hidden represents an option that marks a field as hidden. What it actually causes is up to the Configurable.
	/// ProtoConfig 1.0 recommends hiding it from the documentation.
	//
	// optional bool hidden = 5001;
	E_Hidden = &file_protoconfig_v1_extensions_proto_extTypes[2]
	/// required represents an option that marks a field as mandatory and if empty, Configurable does not accept the whole configuration.
	//
	// optional bool required = 5002;
	E_Required = &file_protoconfig_v1_extensions_proto_extTypes[3]
	/// experimental represents an option that marks a field as experimental. What it actually causes is up to the Configurable.
	/// ProtoConfig 1.0 recommends warning in the documentation.
	//
	// optional bool experimental = 5003;
	E_Experimental = &file_protoconfig_v1_extensions_proto_extTypes[4]
)

var File_protoconfig_v1_extensions_proto protoreflect.FileDescriptor

var file_protoconfig_v1_extensions_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x46, 0x0a, 0x0e, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x64, 0x69, 0x6e, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x64, 0x69, 0x6e,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x43, 0x0a, 0x0d, 0x66, 0x6c, 0x61, 0x67,
	0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x6c, 0x61, 0x67, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52,
	0x0c, 0x66, 0x6c, 0x61, 0x67, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x42, 0x14, 0x0a,
	0x12, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x73, 0x6d, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x64, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x22, 0x22, 0x0a, 0x0c, 0x46, 0x6c, 0x61, 0x67, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x56, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x88, 0x27, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x3a, 0x38, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x88, 0x27, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x36, 0x0a, 0x06, 0x68, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x89, 0x27, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64,
	0x65, 0x6e, 0x3a, 0x3a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8a, 0x27,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x3a, 0x42,
	0x0a, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8b, 0x27,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoconfig_v1_extensions_proto_rawDescOnce sync.Once
	file_protoconfig_v1_extensions_proto_rawDescData = file_protoconfig_v1_extensions_proto_rawDesc
)

func file_protoconfig_v1_extensions_proto_rawDescGZIP() []byte {
	file_protoconfig_v1_extensions_proto_rawDescOnce.Do(func() {
		file_protoconfig_v1_extensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoconfig_v1_extensions_proto_rawDescData)
	})
	return file_protoconfig_v1_extensions_proto_rawDescData
}

var file_protoconfig_v1_extensions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protoconfig_v1_extensions_proto_goTypes = []interface{}{
	(*Metadata)(nil),                  // 0: protoconfig.v1.Metadata
	(*StdinDelivery)(nil),             // 1: protoconfig.v1.StdinDelivery
	(*FlagDelivery)(nil),              // 2: protoconfig.v1.FlagDelivery
	(*descriptor.MessageOptions)(nil), // 3: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
}
var file_protoconfig_v1_extensions_proto_depIdxs = []int32{
	1, // 0: protoconfig.v1.Metadata.stdin_delivery:type_name -> protoconfig.v1.StdinDelivery
	2, // 1: protoconfig.v1.Metadata.flag_delivery:type_name -> protoconfig.v1.FlagDelivery
	3, // 2: protoconfig.v1.metadata:extendee -> google.protobuf.MessageOptions
	4, // 3: protoconfig.v1.default:extendee -> google.protobuf.FieldOptions
	4, // 4: protoconfig.v1.hidden:extendee -> google.protobuf.FieldOptions
	4, // 5: protoconfig.v1.required:extendee -> google.protobuf.FieldOptions
	4, // 6: protoconfig.v1.experimental:extendee -> google.protobuf.FieldOptions
	0, // 7: protoconfig.v1.metadata:type_name -> protoconfig.v1.Metadata
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	7, // [7:8] is the sub-list for extension type_name
	2, // [2:7] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protoconfig_v1_extensions_proto_init() }
func file_protoconfig_v1_extensions_proto_init() {
	if File_protoconfig_v1_extensions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoconfig_v1_extensions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protoconfig_v1_extensions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StdinDelivery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protoconfig_v1_extensions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagDelivery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_protoconfig_v1_extensions_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Metadata_StdinDelivery)(nil),
		(*Metadata_FlagDelivery)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protoconfig_v1_extensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_protoconfig_v1_extensions_proto_goTypes,
		DependencyIndexes: file_protoconfig_v1_extensions_proto_depIdxs,
		MessageInfos:      file_protoconfig_v1_extensions_proto_msgTypes,
		ExtensionInfos:    file_protoconfig_v1_extensions_proto_extTypes,
	}.Build()
	File_protoconfig_v1_extensions_proto = out.File
	file_protoconfig_v1_extensions_proto_rawDesc = nil
	file_protoconfig_v1_extensions_proto_goTypes = nil
	file_protoconfig_v1_extensions_proto_depIdxs = nil
}
