package ecocreditv1alpha2

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_EventCreateClass          protoreflect.MessageDescriptor
	fd_EventCreateClass_class_id protoreflect.FieldDescriptor
	fd_EventCreateClass_admin    protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_events_proto_init()
	md_EventCreateClass = File_regen_ecocredit_v1alpha2_events_proto.Messages().ByName("EventCreateClass")
	fd_EventCreateClass_class_id = md_EventCreateClass.Fields().ByName("class_id")
	fd_EventCreateClass_admin = md_EventCreateClass.Fields().ByName("admin")
}

var _ protoreflect.Message = (*fastReflection_EventCreateClass)(nil)

type fastReflection_EventCreateClass EventCreateClass

func (x *EventCreateClass) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventCreateClass)(x)
}

func (x *EventCreateClass) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventCreateClass_messageType fastReflection_EventCreateClass_messageType
var _ protoreflect.MessageType = fastReflection_EventCreateClass_messageType{}

type fastReflection_EventCreateClass_messageType struct{}

func (x fastReflection_EventCreateClass_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventCreateClass)(nil)
}
func (x fastReflection_EventCreateClass_messageType) New() protoreflect.Message {
	return new(fastReflection_EventCreateClass)
}
func (x fastReflection_EventCreateClass_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventCreateClass
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventCreateClass) Descriptor() protoreflect.MessageDescriptor {
	return md_EventCreateClass
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventCreateClass) Type() protoreflect.MessageType {
	return _fastReflection_EventCreateClass_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventCreateClass) New() protoreflect.Message {
	return new(fastReflection_EventCreateClass)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventCreateClass) Interface() protoreflect.ProtoMessage {
	return (*EventCreateClass)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventCreateClass) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_EventCreateClass_class_id, value) {
			return
		}
	}
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_EventCreateClass_admin, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventCreateClass) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateClass.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha2.EventCreateClass.admin":
		return x.Admin != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateClass does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateClass) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateClass.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha2.EventCreateClass.admin":
		x.Admin = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateClass does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventCreateClass) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateClass.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCreateClass.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateClass does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateClass) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateClass.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCreateClass.admin":
		x.Admin = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateClass does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateClass) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateClass.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha2.EventCreateClass is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCreateClass.admin":
		panic(fmt.Errorf("field admin of message regen.ecocredit.v1alpha2.EventCreateClass is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateClass does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventCreateClass) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateClass.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCreateClass.admin":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateClass does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventCreateClass) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.EventCreateClass", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventCreateClass) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateClass) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventCreateClass) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventCreateClass) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventCreateClass)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventCreateClass)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventCreateClass)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventCreateClass: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventCreateClass: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ClassId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_EventCreateProject                  protoreflect.MessageDescriptor
	fd_EventCreateProject_project_id       protoreflect.FieldDescriptor
	fd_EventCreateProject_class_id         protoreflect.FieldDescriptor
	fd_EventCreateProject_issuer           protoreflect.FieldDescriptor
	fd_EventCreateProject_project_location protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_events_proto_init()
	md_EventCreateProject = File_regen_ecocredit_v1alpha2_events_proto.Messages().ByName("EventCreateProject")
	fd_EventCreateProject_project_id = md_EventCreateProject.Fields().ByName("project_id")
	fd_EventCreateProject_class_id = md_EventCreateProject.Fields().ByName("class_id")
	fd_EventCreateProject_issuer = md_EventCreateProject.Fields().ByName("issuer")
	fd_EventCreateProject_project_location = md_EventCreateProject.Fields().ByName("project_location")
}

var _ protoreflect.Message = (*fastReflection_EventCreateProject)(nil)

type fastReflection_EventCreateProject EventCreateProject

func (x *EventCreateProject) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventCreateProject)(x)
}

func (x *EventCreateProject) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventCreateProject_messageType fastReflection_EventCreateProject_messageType
var _ protoreflect.MessageType = fastReflection_EventCreateProject_messageType{}

type fastReflection_EventCreateProject_messageType struct{}

func (x fastReflection_EventCreateProject_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventCreateProject)(nil)
}
func (x fastReflection_EventCreateProject_messageType) New() protoreflect.Message {
	return new(fastReflection_EventCreateProject)
}
func (x fastReflection_EventCreateProject_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventCreateProject
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventCreateProject) Descriptor() protoreflect.MessageDescriptor {
	return md_EventCreateProject
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventCreateProject) Type() protoreflect.MessageType {
	return _fastReflection_EventCreateProject_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventCreateProject) New() protoreflect.Message {
	return new(fastReflection_EventCreateProject)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventCreateProject) Interface() protoreflect.ProtoMessage {
	return (*EventCreateProject)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventCreateProject) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProjectId != "" {
		value := protoreflect.ValueOfString(x.ProjectId)
		if !f(fd_EventCreateProject_project_id, value) {
			return
		}
	}
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_EventCreateProject_class_id, value) {
			return
		}
	}
	if x.Issuer != "" {
		value := protoreflect.ValueOfString(x.Issuer)
		if !f(fd_EventCreateProject_issuer, value) {
			return
		}
	}
	if x.ProjectLocation != "" {
		value := protoreflect.ValueOfString(x.ProjectLocation)
		if !f(fd_EventCreateProject_project_location, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventCreateProject) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_id":
		return x.ProjectId != ""
	case "regen.ecocredit.v1alpha2.EventCreateProject.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha2.EventCreateProject.issuer":
		return x.Issuer != ""
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_location":
		return x.ProjectLocation != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateProject does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateProject) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_id":
		x.ProjectId = ""
	case "regen.ecocredit.v1alpha2.EventCreateProject.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha2.EventCreateProject.issuer":
		x.Issuer = ""
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_location":
		x.ProjectLocation = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateProject does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventCreateProject) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_id":
		value := x.ProjectId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCreateProject.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCreateProject.issuer":
		value := x.Issuer
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_location":
		value := x.ProjectLocation
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateProject does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateProject) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_id":
		x.ProjectId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCreateProject.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCreateProject.issuer":
		x.Issuer = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_location":
		x.ProjectLocation = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateProject does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateProject) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.v1alpha2.EventCreateProject is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCreateProject.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha2.EventCreateProject is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCreateProject.issuer":
		panic(fmt.Errorf("field issuer of message regen.ecocredit.v1alpha2.EventCreateProject is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_location":
		panic(fmt.Errorf("field project_location of message regen.ecocredit.v1alpha2.EventCreateProject is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateProject does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventCreateProject) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCreateProject.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCreateProject.issuer":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCreateProject.project_location":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateProject"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateProject does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventCreateProject) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.EventCreateProject", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventCreateProject) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateProject) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventCreateProject) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventCreateProject) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventCreateProject)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ProjectId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Issuer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ProjectLocation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventCreateProject)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.ProjectLocation) > 0 {
			i -= len(x.ProjectLocation)
			copy(dAtA[i:], x.ProjectLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectLocation)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Issuer) > 0 {
			i -= len(x.Issuer)
			copy(dAtA[i:], x.Issuer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuer)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ProjectId) > 0 {
			i -= len(x.ProjectId)
			copy(dAtA[i:], x.ProjectId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectId)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventCreateProject)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventCreateProject: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventCreateProject: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ProjectId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ClassId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Issuer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectLocation", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ProjectLocation = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_EventCreateBatch                  protoreflect.MessageDescriptor
	fd_EventCreateBatch_class_id         protoreflect.FieldDescriptor
	fd_EventCreateBatch_batch_denom      protoreflect.FieldDescriptor
	fd_EventCreateBatch_issuer           protoreflect.FieldDescriptor
	fd_EventCreateBatch_total_amount     protoreflect.FieldDescriptor
	fd_EventCreateBatch_start_date       protoreflect.FieldDescriptor
	fd_EventCreateBatch_end_date         protoreflect.FieldDescriptor
	fd_EventCreateBatch_project_location protoreflect.FieldDescriptor
	fd_EventCreateBatch_project_id       protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_events_proto_init()
	md_EventCreateBatch = File_regen_ecocredit_v1alpha2_events_proto.Messages().ByName("EventCreateBatch")
	fd_EventCreateBatch_class_id = md_EventCreateBatch.Fields().ByName("class_id")
	fd_EventCreateBatch_batch_denom = md_EventCreateBatch.Fields().ByName("batch_denom")
	fd_EventCreateBatch_issuer = md_EventCreateBatch.Fields().ByName("issuer")
	fd_EventCreateBatch_total_amount = md_EventCreateBatch.Fields().ByName("total_amount")
	fd_EventCreateBatch_start_date = md_EventCreateBatch.Fields().ByName("start_date")
	fd_EventCreateBatch_end_date = md_EventCreateBatch.Fields().ByName("end_date")
	fd_EventCreateBatch_project_location = md_EventCreateBatch.Fields().ByName("project_location")
	fd_EventCreateBatch_project_id = md_EventCreateBatch.Fields().ByName("project_id")
}

var _ protoreflect.Message = (*fastReflection_EventCreateBatch)(nil)

type fastReflection_EventCreateBatch EventCreateBatch

func (x *EventCreateBatch) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventCreateBatch)(x)
}

func (x *EventCreateBatch) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventCreateBatch_messageType fastReflection_EventCreateBatch_messageType
var _ protoreflect.MessageType = fastReflection_EventCreateBatch_messageType{}

type fastReflection_EventCreateBatch_messageType struct{}

func (x fastReflection_EventCreateBatch_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventCreateBatch)(nil)
}
func (x fastReflection_EventCreateBatch_messageType) New() protoreflect.Message {
	return new(fastReflection_EventCreateBatch)
}
func (x fastReflection_EventCreateBatch_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventCreateBatch
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventCreateBatch) Descriptor() protoreflect.MessageDescriptor {
	return md_EventCreateBatch
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventCreateBatch) Type() protoreflect.MessageType {
	return _fastReflection_EventCreateBatch_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventCreateBatch) New() protoreflect.Message {
	return new(fastReflection_EventCreateBatch)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventCreateBatch) Interface() protoreflect.ProtoMessage {
	return (*EventCreateBatch)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventCreateBatch) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_EventCreateBatch_class_id, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_EventCreateBatch_batch_denom, value) {
			return
		}
	}
	if x.Issuer != "" {
		value := protoreflect.ValueOfString(x.Issuer)
		if !f(fd_EventCreateBatch_issuer, value) {
			return
		}
	}
	if x.TotalAmount != "" {
		value := protoreflect.ValueOfString(x.TotalAmount)
		if !f(fd_EventCreateBatch_total_amount, value) {
			return
		}
	}
	if x.StartDate != "" {
		value := protoreflect.ValueOfString(x.StartDate)
		if !f(fd_EventCreateBatch_start_date, value) {
			return
		}
	}
	if x.EndDate != "" {
		value := protoreflect.ValueOfString(x.EndDate)
		if !f(fd_EventCreateBatch_end_date, value) {
			return
		}
	}
	if x.ProjectLocation != "" {
		value := protoreflect.ValueOfString(x.ProjectLocation)
		if !f(fd_EventCreateBatch_project_location, value) {
			return
		}
	}
	if x.ProjectId != "" {
		value := protoreflect.ValueOfString(x.ProjectId)
		if !f(fd_EventCreateBatch_project_id, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventCreateBatch) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateBatch.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.issuer":
		return x.Issuer != ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.total_amount":
		return x.TotalAmount != ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.start_date":
		return x.StartDate != ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.end_date":
		return x.EndDate != ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_location":
		return x.ProjectLocation != ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_id":
		return x.ProjectId != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateBatch does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateBatch) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateBatch.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.issuer":
		x.Issuer = ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.total_amount":
		x.TotalAmount = ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.start_date":
		x.StartDate = ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.end_date":
		x.EndDate = ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_location":
		x.ProjectLocation = ""
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_id":
		x.ProjectId = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateBatch does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventCreateBatch) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateBatch.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.issuer":
		value := x.Issuer
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.total_amount":
		value := x.TotalAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.start_date":
		value := x.StartDate
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.end_date":
		value := x.EndDate
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_location":
		value := x.ProjectLocation
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_id":
		value := x.ProjectId
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateBatch does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateBatch) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateBatch.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.issuer":
		x.Issuer = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.total_amount":
		x.TotalAmount = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.start_date":
		x.StartDate = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.end_date":
		x.EndDate = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_location":
		x.ProjectLocation = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_id":
		x.ProjectId = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateBatch does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateBatch) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateBatch.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha2.EventCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCreateBatch.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.EventCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCreateBatch.issuer":
		panic(fmt.Errorf("field issuer of message regen.ecocredit.v1alpha2.EventCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCreateBatch.total_amount":
		panic(fmt.Errorf("field total_amount of message regen.ecocredit.v1alpha2.EventCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCreateBatch.start_date":
		panic(fmt.Errorf("field start_date of message regen.ecocredit.v1alpha2.EventCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCreateBatch.end_date":
		panic(fmt.Errorf("field end_date of message regen.ecocredit.v1alpha2.EventCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_location":
		panic(fmt.Errorf("field project_location of message regen.ecocredit.v1alpha2.EventCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.v1alpha2.EventCreateBatch is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateBatch does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventCreateBatch) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCreateBatch.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCreateBatch.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCreateBatch.issuer":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCreateBatch.total_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCreateBatch.start_date":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCreateBatch.end_date":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_location":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCreateBatch.project_id":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCreateBatch does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventCreateBatch) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.EventCreateBatch", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventCreateBatch) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCreateBatch) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventCreateBatch) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventCreateBatch) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventCreateBatch)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Issuer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TotalAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.StartDate)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.EndDate)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ProjectLocation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ProjectId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventCreateBatch)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.ProjectId) > 0 {
			i -= len(x.ProjectId)
			copy(dAtA[i:], x.ProjectId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectId)))
			i--
			dAtA[i] = 0x42
		}
		if len(x.ProjectLocation) > 0 {
			i -= len(x.ProjectLocation)
			copy(dAtA[i:], x.ProjectLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectLocation)))
			i--
			dAtA[i] = 0x3a
		}
		if len(x.EndDate) > 0 {
			i -= len(x.EndDate)
			copy(dAtA[i:], x.EndDate)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.EndDate)))
			i--
			dAtA[i] = 0x32
		}
		if len(x.StartDate) > 0 {
			i -= len(x.StartDate)
			copy(dAtA[i:], x.StartDate)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.StartDate)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.TotalAmount) > 0 {
			i -= len(x.TotalAmount)
			copy(dAtA[i:], x.TotalAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TotalAmount)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Issuer) > 0 {
			i -= len(x.Issuer)
			copy(dAtA[i:], x.Issuer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuer)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventCreateBatch)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventCreateBatch: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventCreateBatch: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ClassId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Issuer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalAmount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.TotalAmount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.StartDate = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.EndDate = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectLocation", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ProjectLocation = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ProjectId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_EventReceive                 protoreflect.MessageDescriptor
	fd_EventReceive_sender          protoreflect.FieldDescriptor
	fd_EventReceive_recipient       protoreflect.FieldDescriptor
	fd_EventReceive_batch_denom     protoreflect.FieldDescriptor
	fd_EventReceive_tradable_amount protoreflect.FieldDescriptor
	fd_EventReceive_retired_amount  protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_events_proto_init()
	md_EventReceive = File_regen_ecocredit_v1alpha2_events_proto.Messages().ByName("EventReceive")
	fd_EventReceive_sender = md_EventReceive.Fields().ByName("sender")
	fd_EventReceive_recipient = md_EventReceive.Fields().ByName("recipient")
	fd_EventReceive_batch_denom = md_EventReceive.Fields().ByName("batch_denom")
	fd_EventReceive_tradable_amount = md_EventReceive.Fields().ByName("tradable_amount")
	fd_EventReceive_retired_amount = md_EventReceive.Fields().ByName("retired_amount")
}

var _ protoreflect.Message = (*fastReflection_EventReceive)(nil)

type fastReflection_EventReceive EventReceive

func (x *EventReceive) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventReceive)(x)
}

func (x *EventReceive) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventReceive_messageType fastReflection_EventReceive_messageType
var _ protoreflect.MessageType = fastReflection_EventReceive_messageType{}

type fastReflection_EventReceive_messageType struct{}

func (x fastReflection_EventReceive_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventReceive)(nil)
}
func (x fastReflection_EventReceive_messageType) New() protoreflect.Message {
	return new(fastReflection_EventReceive)
}
func (x fastReflection_EventReceive_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventReceive
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventReceive) Descriptor() protoreflect.MessageDescriptor {
	return md_EventReceive
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventReceive) Type() protoreflect.MessageType {
	return _fastReflection_EventReceive_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventReceive) New() protoreflect.Message {
	return new(fastReflection_EventReceive)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventReceive) Interface() protoreflect.ProtoMessage {
	return (*EventReceive)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventReceive) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sender != "" {
		value := protoreflect.ValueOfString(x.Sender)
		if !f(fd_EventReceive_sender, value) {
			return
		}
	}
	if x.Recipient != "" {
		value := protoreflect.ValueOfString(x.Recipient)
		if !f(fd_EventReceive_recipient, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_EventReceive_batch_denom, value) {
			return
		}
	}
	if x.TradableAmount != "" {
		value := protoreflect.ValueOfString(x.TradableAmount)
		if !f(fd_EventReceive_tradable_amount, value) {
			return
		}
	}
	if x.RetiredAmount != "" {
		value := protoreflect.ValueOfString(x.RetiredAmount)
		if !f(fd_EventReceive_retired_amount, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventReceive) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventReceive.sender":
		return x.Sender != ""
	case "regen.ecocredit.v1alpha2.EventReceive.recipient":
		return x.Recipient != ""
	case "regen.ecocredit.v1alpha2.EventReceive.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.EventReceive.tradable_amount":
		return x.TradableAmount != ""
	case "regen.ecocredit.v1alpha2.EventReceive.retired_amount":
		return x.RetiredAmount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventReceive"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventReceive does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventReceive) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventReceive.sender":
		x.Sender = ""
	case "regen.ecocredit.v1alpha2.EventReceive.recipient":
		x.Recipient = ""
	case "regen.ecocredit.v1alpha2.EventReceive.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.EventReceive.tradable_amount":
		x.TradableAmount = ""
	case "regen.ecocredit.v1alpha2.EventReceive.retired_amount":
		x.RetiredAmount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventReceive"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventReceive does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventReceive) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.EventReceive.sender":
		value := x.Sender
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventReceive.recipient":
		value := x.Recipient
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventReceive.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventReceive.tradable_amount":
		value := x.TradableAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventReceive.retired_amount":
		value := x.RetiredAmount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventReceive"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventReceive does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventReceive) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventReceive.sender":
		x.Sender = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventReceive.recipient":
		x.Recipient = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventReceive.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventReceive.tradable_amount":
		x.TradableAmount = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventReceive.retired_amount":
		x.RetiredAmount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventReceive"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventReceive does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventReceive) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventReceive.sender":
		panic(fmt.Errorf("field sender of message regen.ecocredit.v1alpha2.EventReceive is not mutable"))
	case "regen.ecocredit.v1alpha2.EventReceive.recipient":
		panic(fmt.Errorf("field recipient of message regen.ecocredit.v1alpha2.EventReceive is not mutable"))
	case "regen.ecocredit.v1alpha2.EventReceive.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.EventReceive is not mutable"))
	case "regen.ecocredit.v1alpha2.EventReceive.tradable_amount":
		panic(fmt.Errorf("field tradable_amount of message regen.ecocredit.v1alpha2.EventReceive is not mutable"))
	case "regen.ecocredit.v1alpha2.EventReceive.retired_amount":
		panic(fmt.Errorf("field retired_amount of message regen.ecocredit.v1alpha2.EventReceive is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventReceive"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventReceive does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventReceive) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventReceive.sender":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventReceive.recipient":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventReceive.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventReceive.tradable_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventReceive.retired_amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventReceive"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventReceive does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventReceive) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.EventReceive", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventReceive) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventReceive) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventReceive) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventReceive) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventReceive)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Sender)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Recipient)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TradableAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RetiredAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventReceive)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.RetiredAmount) > 0 {
			i -= len(x.RetiredAmount)
			copy(dAtA[i:], x.RetiredAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetiredAmount)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.TradableAmount) > 0 {
			i -= len(x.TradableAmount)
			copy(dAtA[i:], x.TradableAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TradableAmount)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Recipient) > 0 {
			i -= len(x.Recipient)
			copy(dAtA[i:], x.Recipient)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Recipient)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Sender) > 0 {
			i -= len(x.Sender)
			copy(dAtA[i:], x.Sender)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Sender)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventReceive)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventReceive: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventReceive: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sender = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Recipient = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TradableAmount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.TradableAmount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetiredAmount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RetiredAmount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_EventRetire             protoreflect.MessageDescriptor
	fd_EventRetire_retirer     protoreflect.FieldDescriptor
	fd_EventRetire_batch_denom protoreflect.FieldDescriptor
	fd_EventRetire_amount      protoreflect.FieldDescriptor
	fd_EventRetire_location    protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_events_proto_init()
	md_EventRetire = File_regen_ecocredit_v1alpha2_events_proto.Messages().ByName("EventRetire")
	fd_EventRetire_retirer = md_EventRetire.Fields().ByName("retirer")
	fd_EventRetire_batch_denom = md_EventRetire.Fields().ByName("batch_denom")
	fd_EventRetire_amount = md_EventRetire.Fields().ByName("amount")
	fd_EventRetire_location = md_EventRetire.Fields().ByName("location")
}

var _ protoreflect.Message = (*fastReflection_EventRetire)(nil)

type fastReflection_EventRetire EventRetire

func (x *EventRetire) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventRetire)(x)
}

func (x *EventRetire) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventRetire_messageType fastReflection_EventRetire_messageType
var _ protoreflect.MessageType = fastReflection_EventRetire_messageType{}

type fastReflection_EventRetire_messageType struct{}

func (x fastReflection_EventRetire_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventRetire)(nil)
}
func (x fastReflection_EventRetire_messageType) New() protoreflect.Message {
	return new(fastReflection_EventRetire)
}
func (x fastReflection_EventRetire_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventRetire
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventRetire) Descriptor() protoreflect.MessageDescriptor {
	return md_EventRetire
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventRetire) Type() protoreflect.MessageType {
	return _fastReflection_EventRetire_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventRetire) New() protoreflect.Message {
	return new(fastReflection_EventRetire)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventRetire) Interface() protoreflect.ProtoMessage {
	return (*EventRetire)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventRetire) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Retirer != "" {
		value := protoreflect.ValueOfString(x.Retirer)
		if !f(fd_EventRetire_retirer, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_EventRetire_batch_denom, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_EventRetire_amount, value) {
			return
		}
	}
	if x.Location != "" {
		value := protoreflect.ValueOfString(x.Location)
		if !f(fd_EventRetire_location, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventRetire) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventRetire.retirer":
		return x.Retirer != ""
	case "regen.ecocredit.v1alpha2.EventRetire.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.EventRetire.amount":
		return x.Amount != ""
	case "regen.ecocredit.v1alpha2.EventRetire.location":
		return x.Location != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventRetire does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventRetire) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventRetire.retirer":
		x.Retirer = ""
	case "regen.ecocredit.v1alpha2.EventRetire.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.EventRetire.amount":
		x.Amount = ""
	case "regen.ecocredit.v1alpha2.EventRetire.location":
		x.Location = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventRetire does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventRetire) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.EventRetire.retirer":
		value := x.Retirer
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventRetire.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventRetire.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventRetire.location":
		value := x.Location
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventRetire does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventRetire) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventRetire.retirer":
		x.Retirer = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventRetire.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventRetire.amount":
		x.Amount = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventRetire.location":
		x.Location = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventRetire does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventRetire) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventRetire.retirer":
		panic(fmt.Errorf("field retirer of message regen.ecocredit.v1alpha2.EventRetire is not mutable"))
	case "regen.ecocredit.v1alpha2.EventRetire.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.EventRetire is not mutable"))
	case "regen.ecocredit.v1alpha2.EventRetire.amount":
		panic(fmt.Errorf("field amount of message regen.ecocredit.v1alpha2.EventRetire is not mutable"))
	case "regen.ecocredit.v1alpha2.EventRetire.location":
		panic(fmt.Errorf("field location of message regen.ecocredit.v1alpha2.EventRetire is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventRetire does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventRetire) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventRetire.retirer":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventRetire.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventRetire.amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventRetire.location":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventRetire does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventRetire) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.EventRetire", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventRetire) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventRetire) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventRetire) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventRetire) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventRetire)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Retirer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Amount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Location)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventRetire)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Location) > 0 {
			i -= len(x.Location)
			copy(dAtA[i:], x.Location)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Location)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Retirer) > 0 {
			i -= len(x.Retirer)
			copy(dAtA[i:], x.Retirer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Retirer)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventRetire)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventRetire: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventRetire: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Retirer", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Retirer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Amount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Location = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_EventCancel             protoreflect.MessageDescriptor
	fd_EventCancel_canceller   protoreflect.FieldDescriptor
	fd_EventCancel_batch_denom protoreflect.FieldDescriptor
	fd_EventCancel_amount      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_events_proto_init()
	md_EventCancel = File_regen_ecocredit_v1alpha2_events_proto.Messages().ByName("EventCancel")
	fd_EventCancel_canceller = md_EventCancel.Fields().ByName("canceller")
	fd_EventCancel_batch_denom = md_EventCancel.Fields().ByName("batch_denom")
	fd_EventCancel_amount = md_EventCancel.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_EventCancel)(nil)

type fastReflection_EventCancel EventCancel

func (x *EventCancel) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventCancel)(x)
}

func (x *EventCancel) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventCancel_messageType fastReflection_EventCancel_messageType
var _ protoreflect.MessageType = fastReflection_EventCancel_messageType{}

type fastReflection_EventCancel_messageType struct{}

func (x fastReflection_EventCancel_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventCancel)(nil)
}
func (x fastReflection_EventCancel_messageType) New() protoreflect.Message {
	return new(fastReflection_EventCancel)
}
func (x fastReflection_EventCancel_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventCancel
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventCancel) Descriptor() protoreflect.MessageDescriptor {
	return md_EventCancel
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventCancel) Type() protoreflect.MessageType {
	return _fastReflection_EventCancel_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventCancel) New() protoreflect.Message {
	return new(fastReflection_EventCancel)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventCancel) Interface() protoreflect.ProtoMessage {
	return (*EventCancel)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventCancel) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Canceller != "" {
		value := protoreflect.ValueOfString(x.Canceller)
		if !f(fd_EventCancel_canceller, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_EventCancel_batch_denom, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_EventCancel_amount, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventCancel) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCancel.canceller":
		return x.Canceller != ""
	case "regen.ecocredit.v1alpha2.EventCancel.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.EventCancel.amount":
		return x.Amount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCancel does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCancel) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCancel.canceller":
		x.Canceller = ""
	case "regen.ecocredit.v1alpha2.EventCancel.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.EventCancel.amount":
		x.Amount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCancel does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventCancel) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.EventCancel.canceller":
		value := x.Canceller
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCancel.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventCancel.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCancel does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCancel) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCancel.canceller":
		x.Canceller = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCancel.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventCancel.amount":
		x.Amount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCancel does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCancel) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCancel.canceller":
		panic(fmt.Errorf("field canceller of message regen.ecocredit.v1alpha2.EventCancel is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCancel.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.EventCancel is not mutable"))
	case "regen.ecocredit.v1alpha2.EventCancel.amount":
		panic(fmt.Errorf("field amount of message regen.ecocredit.v1alpha2.EventCancel is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCancel does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventCancel) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventCancel.canceller":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCancel.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventCancel.amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventCancel does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventCancel) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.EventCancel", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventCancel) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventCancel) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventCancel) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventCancel) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventCancel)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Canceller)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Amount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventCancel)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Canceller) > 0 {
			i -= len(x.Canceller)
			copy(dAtA[i:], x.Canceller)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Canceller)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventCancel)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventCancel: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventCancel: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Canceller", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Canceller = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Amount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_EventSell                     protoreflect.MessageDescriptor
	fd_EventSell_order_id            protoreflect.FieldDescriptor
	fd_EventSell_batch_denom         protoreflect.FieldDescriptor
	fd_EventSell_quantity            protoreflect.FieldDescriptor
	fd_EventSell_ask_price           protoreflect.FieldDescriptor
	fd_EventSell_disable_auto_retire protoreflect.FieldDescriptor
	fd_EventSell_expiration          protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_events_proto_init()
	md_EventSell = File_regen_ecocredit_v1alpha2_events_proto.Messages().ByName("EventSell")
	fd_EventSell_order_id = md_EventSell.Fields().ByName("order_id")
	fd_EventSell_batch_denom = md_EventSell.Fields().ByName("batch_denom")
	fd_EventSell_quantity = md_EventSell.Fields().ByName("quantity")
	fd_EventSell_ask_price = md_EventSell.Fields().ByName("ask_price")
	fd_EventSell_disable_auto_retire = md_EventSell.Fields().ByName("disable_auto_retire")
	fd_EventSell_expiration = md_EventSell.Fields().ByName("expiration")
}

var _ protoreflect.Message = (*fastReflection_EventSell)(nil)

type fastReflection_EventSell EventSell

func (x *EventSell) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventSell)(x)
}

func (x *EventSell) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventSell_messageType fastReflection_EventSell_messageType
var _ protoreflect.MessageType = fastReflection_EventSell_messageType{}

type fastReflection_EventSell_messageType struct{}

func (x fastReflection_EventSell_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventSell)(nil)
}
func (x fastReflection_EventSell_messageType) New() protoreflect.Message {
	return new(fastReflection_EventSell)
}
func (x fastReflection_EventSell_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventSell
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventSell) Descriptor() protoreflect.MessageDescriptor {
	return md_EventSell
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventSell) Type() protoreflect.MessageType {
	return _fastReflection_EventSell_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventSell) New() protoreflect.Message {
	return new(fastReflection_EventSell)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventSell) Interface() protoreflect.ProtoMessage {
	return (*EventSell)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventSell) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.OrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.OrderId)
		if !f(fd_EventSell_order_id, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_EventSell_batch_denom, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_EventSell_quantity, value) {
			return
		}
	}
	if x.AskPrice != nil {
		value := protoreflect.ValueOfMessage(x.AskPrice.ProtoReflect())
		if !f(fd_EventSell_ask_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_EventSell_disable_auto_retire, value) {
			return
		}
	}
	if x.Expiration != nil {
		value := protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
		if !f(fd_EventSell_expiration, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventSell) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventSell.order_id":
		return x.OrderId != uint64(0)
	case "regen.ecocredit.v1alpha2.EventSell.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.EventSell.quantity":
		return x.Quantity != ""
	case "regen.ecocredit.v1alpha2.EventSell.ask_price":
		return x.AskPrice != nil
	case "regen.ecocredit.v1alpha2.EventSell.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.v1alpha2.EventSell.expiration":
		return x.Expiration != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventSell does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventSell) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventSell.order_id":
		x.OrderId = uint64(0)
	case "regen.ecocredit.v1alpha2.EventSell.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.EventSell.quantity":
		x.Quantity = ""
	case "regen.ecocredit.v1alpha2.EventSell.ask_price":
		x.AskPrice = nil
	case "regen.ecocredit.v1alpha2.EventSell.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.v1alpha2.EventSell.expiration":
		x.Expiration = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventSell does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventSell) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.EventSell.order_id":
		value := x.OrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1alpha2.EventSell.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventSell.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventSell.ask_price":
		value := x.AskPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventSell.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1alpha2.EventSell.expiration":
		value := x.Expiration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventSell does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventSell) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventSell.order_id":
		x.OrderId = value.Uint()
	case "regen.ecocredit.v1alpha2.EventSell.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventSell.quantity":
		x.Quantity = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventSell.ask_price":
		x.AskPrice = value.Message().Interface().(*v1beta1.Coin)
	case "regen.ecocredit.v1alpha2.EventSell.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.v1alpha2.EventSell.expiration":
		x.Expiration = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventSell does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventSell) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventSell.ask_price":
		if x.AskPrice == nil {
			x.AskPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.AskPrice.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventSell.expiration":
		if x.Expiration == nil {
			x.Expiration = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventSell.order_id":
		panic(fmt.Errorf("field order_id of message regen.ecocredit.v1alpha2.EventSell is not mutable"))
	case "regen.ecocredit.v1alpha2.EventSell.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.EventSell is not mutable"))
	case "regen.ecocredit.v1alpha2.EventSell.quantity":
		panic(fmt.Errorf("field quantity of message regen.ecocredit.v1alpha2.EventSell is not mutable"))
	case "regen.ecocredit.v1alpha2.EventSell.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.v1alpha2.EventSell is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventSell does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventSell) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventSell.order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1alpha2.EventSell.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventSell.quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventSell.ask_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventSell.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1alpha2.EventSell.expiration":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventSell"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventSell does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventSell) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.EventSell", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventSell) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventSell) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventSell) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventSell) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventSell)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.OrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.OrderId))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Quantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AskPrice != nil {
			l = options.Size(x.AskPrice)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DisableAutoRetire {
			n += 2
		}
		if x.Expiration != nil {
			l = options.Size(x.Expiration)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventSell)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Expiration != nil {
			encoded, err := options.Marshal(x.Expiration)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x32
		}
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x28
		}
		if x.AskPrice != nil {
			encoded, err := options.Marshal(x.AskPrice)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x12
		}
		if x.OrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.OrderId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventSell)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventSell: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventSell: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
				}
				x.OrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.OrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Quantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AskPrice", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.AskPrice == nil {
					x.AskPrice = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.AskPrice); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisableAutoRetire = bool(v != 0)
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Expiration == nil {
					x.Expiration = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Expiration); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_EventUpdateSellOrder                     protoreflect.MessageDescriptor
	fd_EventUpdateSellOrder_owner               protoreflect.FieldDescriptor
	fd_EventUpdateSellOrder_sell_order_id       protoreflect.FieldDescriptor
	fd_EventUpdateSellOrder_batch_denom         protoreflect.FieldDescriptor
	fd_EventUpdateSellOrder_new_quantity        protoreflect.FieldDescriptor
	fd_EventUpdateSellOrder_new_ask_price       protoreflect.FieldDescriptor
	fd_EventUpdateSellOrder_disable_auto_retire protoreflect.FieldDescriptor
	fd_EventUpdateSellOrder_new_expiration      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_events_proto_init()
	md_EventUpdateSellOrder = File_regen_ecocredit_v1alpha2_events_proto.Messages().ByName("EventUpdateSellOrder")
	fd_EventUpdateSellOrder_owner = md_EventUpdateSellOrder.Fields().ByName("owner")
	fd_EventUpdateSellOrder_sell_order_id = md_EventUpdateSellOrder.Fields().ByName("sell_order_id")
	fd_EventUpdateSellOrder_batch_denom = md_EventUpdateSellOrder.Fields().ByName("batch_denom")
	fd_EventUpdateSellOrder_new_quantity = md_EventUpdateSellOrder.Fields().ByName("new_quantity")
	fd_EventUpdateSellOrder_new_ask_price = md_EventUpdateSellOrder.Fields().ByName("new_ask_price")
	fd_EventUpdateSellOrder_disable_auto_retire = md_EventUpdateSellOrder.Fields().ByName("disable_auto_retire")
	fd_EventUpdateSellOrder_new_expiration = md_EventUpdateSellOrder.Fields().ByName("new_expiration")
}

var _ protoreflect.Message = (*fastReflection_EventUpdateSellOrder)(nil)

type fastReflection_EventUpdateSellOrder EventUpdateSellOrder

func (x *EventUpdateSellOrder) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventUpdateSellOrder)(x)
}

func (x *EventUpdateSellOrder) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventUpdateSellOrder_messageType fastReflection_EventUpdateSellOrder_messageType
var _ protoreflect.MessageType = fastReflection_EventUpdateSellOrder_messageType{}

type fastReflection_EventUpdateSellOrder_messageType struct{}

func (x fastReflection_EventUpdateSellOrder_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventUpdateSellOrder)(nil)
}
func (x fastReflection_EventUpdateSellOrder_messageType) New() protoreflect.Message {
	return new(fastReflection_EventUpdateSellOrder)
}
func (x fastReflection_EventUpdateSellOrder_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventUpdateSellOrder
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventUpdateSellOrder) Descriptor() protoreflect.MessageDescriptor {
	return md_EventUpdateSellOrder
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventUpdateSellOrder) Type() protoreflect.MessageType {
	return _fastReflection_EventUpdateSellOrder_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventUpdateSellOrder) New() protoreflect.Message {
	return new(fastReflection_EventUpdateSellOrder)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventUpdateSellOrder) Interface() protoreflect.ProtoMessage {
	return (*EventUpdateSellOrder)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventUpdateSellOrder) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_EventUpdateSellOrder_owner, value) {
			return
		}
	}
	if x.SellOrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.SellOrderId)
		if !f(fd_EventUpdateSellOrder_sell_order_id, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_EventUpdateSellOrder_batch_denom, value) {
			return
		}
	}
	if x.NewQuantity != "" {
		value := protoreflect.ValueOfString(x.NewQuantity)
		if !f(fd_EventUpdateSellOrder_new_quantity, value) {
			return
		}
	}
	if x.NewAskPrice != nil {
		value := protoreflect.ValueOfMessage(x.NewAskPrice.ProtoReflect())
		if !f(fd_EventUpdateSellOrder_new_ask_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_EventUpdateSellOrder_disable_auto_retire, value) {
			return
		}
	}
	if x.NewExpiration != nil {
		value := protoreflect.ValueOfMessage(x.NewExpiration.ProtoReflect())
		if !f(fd_EventUpdateSellOrder_new_expiration, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventUpdateSellOrder) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.owner":
		return x.Owner != ""
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.sell_order_id":
		return x.SellOrderId != uint64(0)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_quantity":
		return x.NewQuantity != ""
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_ask_price":
		return x.NewAskPrice != nil
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_expiration":
		return x.NewExpiration != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventUpdateSellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventUpdateSellOrder does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventUpdateSellOrder) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.owner":
		x.Owner = ""
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.sell_order_id":
		x.SellOrderId = uint64(0)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_quantity":
		x.NewQuantity = ""
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_ask_price":
		x.NewAskPrice = nil
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_expiration":
		x.NewExpiration = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventUpdateSellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventUpdateSellOrder does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventUpdateSellOrder) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.sell_order_id":
		value := x.SellOrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_quantity":
		value := x.NewQuantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_ask_price":
		value := x.NewAskPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_expiration":
		value := x.NewExpiration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventUpdateSellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventUpdateSellOrder does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventUpdateSellOrder) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.owner":
		x.Owner = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.sell_order_id":
		x.SellOrderId = value.Uint()
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_quantity":
		x.NewQuantity = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_ask_price":
		x.NewAskPrice = value.Message().Interface().(*v1beta1.Coin)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_expiration":
		x.NewExpiration = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventUpdateSellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventUpdateSellOrder does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventUpdateSellOrder) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_ask_price":
		if x.NewAskPrice == nil {
			x.NewAskPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.NewAskPrice.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_expiration":
		if x.NewExpiration == nil {
			x.NewExpiration = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.NewExpiration.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.owner":
		panic(fmt.Errorf("field owner of message regen.ecocredit.v1alpha2.EventUpdateSellOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.sell_order_id":
		panic(fmt.Errorf("field sell_order_id of message regen.ecocredit.v1alpha2.EventUpdateSellOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.EventUpdateSellOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_quantity":
		panic(fmt.Errorf("field new_quantity of message regen.ecocredit.v1alpha2.EventUpdateSellOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.v1alpha2.EventUpdateSellOrder is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventUpdateSellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventUpdateSellOrder does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventUpdateSellOrder) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.owner":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.sell_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_ask_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_expiration":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventUpdateSellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventUpdateSellOrder does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventUpdateSellOrder) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.EventUpdateSellOrder", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventUpdateSellOrder) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventUpdateSellOrder) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventUpdateSellOrder) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventUpdateSellOrder) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventUpdateSellOrder)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.SellOrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.SellOrderId))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NewQuantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.NewAskPrice != nil {
			l = options.Size(x.NewAskPrice)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DisableAutoRetire {
			n += 2
		}
		if x.NewExpiration != nil {
			l = options.Size(x.NewExpiration)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventUpdateSellOrder)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.NewExpiration != nil {
			encoded, err := options.Marshal(x.NewExpiration)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x3a
		}
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x30
		}
		if x.NewAskPrice != nil {
			encoded, err := options.Marshal(x.NewAskPrice)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.NewQuantity) > 0 {
			i -= len(x.NewQuantity)
			copy(dAtA[i:], x.NewQuantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NewQuantity)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x1a
		}
		if x.SellOrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SellOrderId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventUpdateSellOrder)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventUpdateSellOrder: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventUpdateSellOrder: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Owner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SellOrderId", wireType)
				}
				x.SellOrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SellOrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewQuantity", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.NewQuantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewAskPrice", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.NewAskPrice == nil {
					x.NewAskPrice = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.NewAskPrice); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisableAutoRetire = bool(v != 0)
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewExpiration", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.NewExpiration == nil {
					x.NewExpiration = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.NewExpiration); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_EventBuyOrderCreated                      protoreflect.MessageDescriptor
	fd_EventBuyOrderCreated_buy_order_id         protoreflect.FieldDescriptor
	fd_EventBuyOrderCreated_sell_order_id        protoreflect.FieldDescriptor
	fd_EventBuyOrderCreated_quantity             protoreflect.FieldDescriptor
	fd_EventBuyOrderCreated_bid_price            protoreflect.FieldDescriptor
	fd_EventBuyOrderCreated_disable_auto_retire  protoreflect.FieldDescriptor
	fd_EventBuyOrderCreated_disable_partial_fill protoreflect.FieldDescriptor
	fd_EventBuyOrderCreated_retirement_location  protoreflect.FieldDescriptor
	fd_EventBuyOrderCreated_expiration           protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_events_proto_init()
	md_EventBuyOrderCreated = File_regen_ecocredit_v1alpha2_events_proto.Messages().ByName("EventBuyOrderCreated")
	fd_EventBuyOrderCreated_buy_order_id = md_EventBuyOrderCreated.Fields().ByName("buy_order_id")
	fd_EventBuyOrderCreated_sell_order_id = md_EventBuyOrderCreated.Fields().ByName("sell_order_id")
	fd_EventBuyOrderCreated_quantity = md_EventBuyOrderCreated.Fields().ByName("quantity")
	fd_EventBuyOrderCreated_bid_price = md_EventBuyOrderCreated.Fields().ByName("bid_price")
	fd_EventBuyOrderCreated_disable_auto_retire = md_EventBuyOrderCreated.Fields().ByName("disable_auto_retire")
	fd_EventBuyOrderCreated_disable_partial_fill = md_EventBuyOrderCreated.Fields().ByName("disable_partial_fill")
	fd_EventBuyOrderCreated_retirement_location = md_EventBuyOrderCreated.Fields().ByName("retirement_location")
	fd_EventBuyOrderCreated_expiration = md_EventBuyOrderCreated.Fields().ByName("expiration")
}

var _ protoreflect.Message = (*fastReflection_EventBuyOrderCreated)(nil)

type fastReflection_EventBuyOrderCreated EventBuyOrderCreated

func (x *EventBuyOrderCreated) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventBuyOrderCreated)(x)
}

func (x *EventBuyOrderCreated) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventBuyOrderCreated_messageType fastReflection_EventBuyOrderCreated_messageType
var _ protoreflect.MessageType = fastReflection_EventBuyOrderCreated_messageType{}

type fastReflection_EventBuyOrderCreated_messageType struct{}

func (x fastReflection_EventBuyOrderCreated_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventBuyOrderCreated)(nil)
}
func (x fastReflection_EventBuyOrderCreated_messageType) New() protoreflect.Message {
	return new(fastReflection_EventBuyOrderCreated)
}
func (x fastReflection_EventBuyOrderCreated_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventBuyOrderCreated
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventBuyOrderCreated) Descriptor() protoreflect.MessageDescriptor {
	return md_EventBuyOrderCreated
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventBuyOrderCreated) Type() protoreflect.MessageType {
	return _fastReflection_EventBuyOrderCreated_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventBuyOrderCreated) New() protoreflect.Message {
	return new(fastReflection_EventBuyOrderCreated)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventBuyOrderCreated) Interface() protoreflect.ProtoMessage {
	return (*EventBuyOrderCreated)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventBuyOrderCreated) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BuyOrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BuyOrderId)
		if !f(fd_EventBuyOrderCreated_buy_order_id, value) {
			return
		}
	}
	if x.SellOrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.SellOrderId)
		if !f(fd_EventBuyOrderCreated_sell_order_id, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_EventBuyOrderCreated_quantity, value) {
			return
		}
	}
	if x.BidPrice != nil {
		value := protoreflect.ValueOfMessage(x.BidPrice.ProtoReflect())
		if !f(fd_EventBuyOrderCreated_bid_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_EventBuyOrderCreated_disable_auto_retire, value) {
			return
		}
	}
	if x.DisablePartialFill != false {
		value := protoreflect.ValueOfBool(x.DisablePartialFill)
		if !f(fd_EventBuyOrderCreated_disable_partial_fill, value) {
			return
		}
	}
	if x.RetirementLocation != "" {
		value := protoreflect.ValueOfString(x.RetirementLocation)
		if !f(fd_EventBuyOrderCreated_retirement_location, value) {
			return
		}
	}
	if x.Expiration != nil {
		value := protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
		if !f(fd_EventBuyOrderCreated_expiration, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventBuyOrderCreated) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.buy_order_id":
		return x.BuyOrderId != uint64(0)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.sell_order_id":
		return x.SellOrderId != uint64(0)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.quantity":
		return x.Quantity != ""
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.bid_price":
		return x.BidPrice != nil
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_partial_fill":
		return x.DisablePartialFill != false
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.retirement_location":
		return x.RetirementLocation != ""
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.expiration":
		return x.Expiration != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderCreated"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderCreated does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventBuyOrderCreated) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.buy_order_id":
		x.BuyOrderId = uint64(0)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.sell_order_id":
		x.SellOrderId = uint64(0)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.quantity":
		x.Quantity = ""
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.bid_price":
		x.BidPrice = nil
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_partial_fill":
		x.DisablePartialFill = false
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.retirement_location":
		x.RetirementLocation = ""
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.expiration":
		x.Expiration = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderCreated"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderCreated does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventBuyOrderCreated) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.buy_order_id":
		value := x.BuyOrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.sell_order_id":
		value := x.SellOrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.bid_price":
		value := x.BidPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_partial_fill":
		value := x.DisablePartialFill
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.retirement_location":
		value := x.RetirementLocation
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.expiration":
		value := x.Expiration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderCreated"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderCreated does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventBuyOrderCreated) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.buy_order_id":
		x.BuyOrderId = value.Uint()
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.sell_order_id":
		x.SellOrderId = value.Uint()
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.quantity":
		x.Quantity = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.bid_price":
		x.BidPrice = value.Message().Interface().(*v1beta1.Coin)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_partial_fill":
		x.DisablePartialFill = value.Bool()
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.retirement_location":
		x.RetirementLocation = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.expiration":
		x.Expiration = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderCreated"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderCreated does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventBuyOrderCreated) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.bid_price":
		if x.BidPrice == nil {
			x.BidPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.BidPrice.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.expiration":
		if x.Expiration == nil {
			x.Expiration = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.buy_order_id":
		panic(fmt.Errorf("field buy_order_id of message regen.ecocredit.v1alpha2.EventBuyOrderCreated is not mutable"))
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.sell_order_id":
		panic(fmt.Errorf("field sell_order_id of message regen.ecocredit.v1alpha2.EventBuyOrderCreated is not mutable"))
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.quantity":
		panic(fmt.Errorf("field quantity of message regen.ecocredit.v1alpha2.EventBuyOrderCreated is not mutable"))
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.v1alpha2.EventBuyOrderCreated is not mutable"))
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_partial_fill":
		panic(fmt.Errorf("field disable_partial_fill of message regen.ecocredit.v1alpha2.EventBuyOrderCreated is not mutable"))
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.retirement_location":
		panic(fmt.Errorf("field retirement_location of message regen.ecocredit.v1alpha2.EventBuyOrderCreated is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderCreated"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderCreated does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventBuyOrderCreated) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.buy_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.sell_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.bid_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.disable_partial_fill":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.retirement_location":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventBuyOrderCreated.expiration":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderCreated"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderCreated does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventBuyOrderCreated) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.EventBuyOrderCreated", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventBuyOrderCreated) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventBuyOrderCreated) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventBuyOrderCreated) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventBuyOrderCreated) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventBuyOrderCreated)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BuyOrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.BuyOrderId))
		}
		if x.SellOrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.SellOrderId))
		}
		l = len(x.Quantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.BidPrice != nil {
			l = options.Size(x.BidPrice)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DisableAutoRetire {
			n += 2
		}
		if x.DisablePartialFill {
			n += 2
		}
		l = len(x.RetirementLocation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Expiration != nil {
			l = options.Size(x.Expiration)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventBuyOrderCreated)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Expiration != nil {
			encoded, err := options.Marshal(x.Expiration)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x42
		}
		if len(x.RetirementLocation) > 0 {
			i -= len(x.RetirementLocation)
			copy(dAtA[i:], x.RetirementLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetirementLocation)))
			i--
			dAtA[i] = 0x3a
		}
		if x.DisablePartialFill {
			i--
			if x.DisablePartialFill {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x30
		}
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x28
		}
		if x.BidPrice != nil {
			encoded, err := options.Marshal(x.BidPrice)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x1a
		}
		if x.SellOrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SellOrderId))
			i--
			dAtA[i] = 0x10
		}
		if x.BuyOrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BuyOrderId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventBuyOrderCreated)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventBuyOrderCreated: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventBuyOrderCreated: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BuyOrderId", wireType)
				}
				x.BuyOrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BuyOrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SellOrderId", wireType)
				}
				x.SellOrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SellOrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Quantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BidPrice", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.BidPrice == nil {
					x.BidPrice = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BidPrice); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisableAutoRetire = bool(v != 0)
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisablePartialFill", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisablePartialFill = bool(v != 0)
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetirementLocation", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RetirementLocation = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Expiration == nil {
					x.Expiration = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Expiration); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_EventBuyOrderFilled               protoreflect.MessageDescriptor
	fd_EventBuyOrderFilled_buy_order_id  protoreflect.FieldDescriptor
	fd_EventBuyOrderFilled_sell_order_id protoreflect.FieldDescriptor
	fd_EventBuyOrderFilled_batch_denom   protoreflect.FieldDescriptor
	fd_EventBuyOrderFilled_quantity      protoreflect.FieldDescriptor
	fd_EventBuyOrderFilled_total_price   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_events_proto_init()
	md_EventBuyOrderFilled = File_regen_ecocredit_v1alpha2_events_proto.Messages().ByName("EventBuyOrderFilled")
	fd_EventBuyOrderFilled_buy_order_id = md_EventBuyOrderFilled.Fields().ByName("buy_order_id")
	fd_EventBuyOrderFilled_sell_order_id = md_EventBuyOrderFilled.Fields().ByName("sell_order_id")
	fd_EventBuyOrderFilled_batch_denom = md_EventBuyOrderFilled.Fields().ByName("batch_denom")
	fd_EventBuyOrderFilled_quantity = md_EventBuyOrderFilled.Fields().ByName("quantity")
	fd_EventBuyOrderFilled_total_price = md_EventBuyOrderFilled.Fields().ByName("total_price")
}

var _ protoreflect.Message = (*fastReflection_EventBuyOrderFilled)(nil)

type fastReflection_EventBuyOrderFilled EventBuyOrderFilled

func (x *EventBuyOrderFilled) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventBuyOrderFilled)(x)
}

func (x *EventBuyOrderFilled) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventBuyOrderFilled_messageType fastReflection_EventBuyOrderFilled_messageType
var _ protoreflect.MessageType = fastReflection_EventBuyOrderFilled_messageType{}

type fastReflection_EventBuyOrderFilled_messageType struct{}

func (x fastReflection_EventBuyOrderFilled_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventBuyOrderFilled)(nil)
}
func (x fastReflection_EventBuyOrderFilled_messageType) New() protoreflect.Message {
	return new(fastReflection_EventBuyOrderFilled)
}
func (x fastReflection_EventBuyOrderFilled_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventBuyOrderFilled
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventBuyOrderFilled) Descriptor() protoreflect.MessageDescriptor {
	return md_EventBuyOrderFilled
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventBuyOrderFilled) Type() protoreflect.MessageType {
	return _fastReflection_EventBuyOrderFilled_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventBuyOrderFilled) New() protoreflect.Message {
	return new(fastReflection_EventBuyOrderFilled)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventBuyOrderFilled) Interface() protoreflect.ProtoMessage {
	return (*EventBuyOrderFilled)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventBuyOrderFilled) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BuyOrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BuyOrderId)
		if !f(fd_EventBuyOrderFilled_buy_order_id, value) {
			return
		}
	}
	if x.SellOrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.SellOrderId)
		if !f(fd_EventBuyOrderFilled_sell_order_id, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_EventBuyOrderFilled_batch_denom, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_EventBuyOrderFilled_quantity, value) {
			return
		}
	}
	if x.TotalPrice != nil {
		value := protoreflect.ValueOfMessage(x.TotalPrice.ProtoReflect())
		if !f(fd_EventBuyOrderFilled_total_price, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventBuyOrderFilled) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.buy_order_id":
		return x.BuyOrderId != uint64(0)
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.sell_order_id":
		return x.SellOrderId != uint64(0)
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.quantity":
		return x.Quantity != ""
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.total_price":
		return x.TotalPrice != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderFilled"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderFilled does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventBuyOrderFilled) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.buy_order_id":
		x.BuyOrderId = uint64(0)
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.sell_order_id":
		x.SellOrderId = uint64(0)
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.quantity":
		x.Quantity = ""
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.total_price":
		x.TotalPrice = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderFilled"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderFilled does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventBuyOrderFilled) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.buy_order_id":
		value := x.BuyOrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.sell_order_id":
		value := x.SellOrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.total_price":
		value := x.TotalPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderFilled"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderFilled does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventBuyOrderFilled) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.buy_order_id":
		x.BuyOrderId = value.Uint()
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.sell_order_id":
		x.SellOrderId = value.Uint()
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.quantity":
		x.Quantity = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.total_price":
		x.TotalPrice = value.Message().Interface().(*v1beta1.Coin)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderFilled"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderFilled does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventBuyOrderFilled) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.total_price":
		if x.TotalPrice == nil {
			x.TotalPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.TotalPrice.ProtoReflect())
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.buy_order_id":
		panic(fmt.Errorf("field buy_order_id of message regen.ecocredit.v1alpha2.EventBuyOrderFilled is not mutable"))
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.sell_order_id":
		panic(fmt.Errorf("field sell_order_id of message regen.ecocredit.v1alpha2.EventBuyOrderFilled is not mutable"))
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.EventBuyOrderFilled is not mutable"))
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.quantity":
		panic(fmt.Errorf("field quantity of message regen.ecocredit.v1alpha2.EventBuyOrderFilled is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderFilled"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderFilled does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventBuyOrderFilled) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.buy_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.sell_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventBuyOrderFilled.total_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventBuyOrderFilled"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventBuyOrderFilled does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventBuyOrderFilled) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.EventBuyOrderFilled", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventBuyOrderFilled) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventBuyOrderFilled) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventBuyOrderFilled) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventBuyOrderFilled) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventBuyOrderFilled)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BuyOrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.BuyOrderId))
		}
		if x.SellOrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.SellOrderId))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Quantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.TotalPrice != nil {
			l = options.Size(x.TotalPrice)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventBuyOrderFilled)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.TotalPrice != nil {
			encoded, err := options.Marshal(x.TotalPrice)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x1a
		}
		if x.SellOrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SellOrderId))
			i--
			dAtA[i] = 0x10
		}
		if x.BuyOrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BuyOrderId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventBuyOrderFilled)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventBuyOrderFilled: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventBuyOrderFilled: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BuyOrderId", wireType)
				}
				x.BuyOrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BuyOrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SellOrderId", wireType)
				}
				x.SellOrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SellOrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Quantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalPrice", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.TotalPrice == nil {
					x.TotalPrice = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.TotalPrice); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_EventAllowAskDenom               protoreflect.MessageDescriptor
	fd_EventAllowAskDenom_denom         protoreflect.FieldDescriptor
	fd_EventAllowAskDenom_display_denom protoreflect.FieldDescriptor
	fd_EventAllowAskDenom_exponent      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_events_proto_init()
	md_EventAllowAskDenom = File_regen_ecocredit_v1alpha2_events_proto.Messages().ByName("EventAllowAskDenom")
	fd_EventAllowAskDenom_denom = md_EventAllowAskDenom.Fields().ByName("denom")
	fd_EventAllowAskDenom_display_denom = md_EventAllowAskDenom.Fields().ByName("display_denom")
	fd_EventAllowAskDenom_exponent = md_EventAllowAskDenom.Fields().ByName("exponent")
}

var _ protoreflect.Message = (*fastReflection_EventAllowAskDenom)(nil)

type fastReflection_EventAllowAskDenom EventAllowAskDenom

func (x *EventAllowAskDenom) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EventAllowAskDenom)(x)
}

func (x *EventAllowAskDenom) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EventAllowAskDenom_messageType fastReflection_EventAllowAskDenom_messageType
var _ protoreflect.MessageType = fastReflection_EventAllowAskDenom_messageType{}

type fastReflection_EventAllowAskDenom_messageType struct{}

func (x fastReflection_EventAllowAskDenom_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EventAllowAskDenom)(nil)
}
func (x fastReflection_EventAllowAskDenom_messageType) New() protoreflect.Message {
	return new(fastReflection_EventAllowAskDenom)
}
func (x fastReflection_EventAllowAskDenom_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EventAllowAskDenom
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EventAllowAskDenom) Descriptor() protoreflect.MessageDescriptor {
	return md_EventAllowAskDenom
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EventAllowAskDenom) Type() protoreflect.MessageType {
	return _fastReflection_EventAllowAskDenom_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EventAllowAskDenom) New() protoreflect.Message {
	return new(fastReflection_EventAllowAskDenom)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EventAllowAskDenom) Interface() protoreflect.ProtoMessage {
	return (*EventAllowAskDenom)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EventAllowAskDenom) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Denom != "" {
		value := protoreflect.ValueOfString(x.Denom)
		if !f(fd_EventAllowAskDenom_denom, value) {
			return
		}
	}
	if x.DisplayDenom != "" {
		value := protoreflect.ValueOfString(x.DisplayDenom)
		if !f(fd_EventAllowAskDenom_display_denom, value) {
			return
		}
	}
	if x.Exponent != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Exponent)
		if !f(fd_EventAllowAskDenom_exponent, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_EventAllowAskDenom) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.denom":
		return x.Denom != ""
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.display_denom":
		return x.DisplayDenom != ""
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.exponent":
		return x.Exponent != uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventAllowAskDenom does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventAllowAskDenom) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.denom":
		x.Denom = ""
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.display_denom":
		x.DisplayDenom = ""
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.exponent":
		x.Exponent = uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventAllowAskDenom does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EventAllowAskDenom) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.denom":
		value := x.Denom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.display_denom":
		value := x.DisplayDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.exponent":
		value := x.Exponent
		return protoreflect.ValueOfUint32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventAllowAskDenom does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventAllowAskDenom) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.denom":
		x.Denom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.display_denom":
		x.DisplayDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.exponent":
		x.Exponent = uint32(value.Uint())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventAllowAskDenom does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventAllowAskDenom) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.denom":
		panic(fmt.Errorf("field denom of message regen.ecocredit.v1alpha2.EventAllowAskDenom is not mutable"))
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.display_denom":
		panic(fmt.Errorf("field display_denom of message regen.ecocredit.v1alpha2.EventAllowAskDenom is not mutable"))
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.exponent":
		panic(fmt.Errorf("field exponent of message regen.ecocredit.v1alpha2.EventAllowAskDenom is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventAllowAskDenom does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EventAllowAskDenom) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.display_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.EventAllowAskDenom.exponent":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.EventAllowAskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.EventAllowAskDenom does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EventAllowAskDenom) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.EventAllowAskDenom", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EventAllowAskDenom) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EventAllowAskDenom) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_EventAllowAskDenom) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EventAllowAskDenom) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EventAllowAskDenom)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Denom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DisplayDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Exponent != 0 {
			n += 1 + runtime.Sov(uint64(x.Exponent))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*EventAllowAskDenom)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Exponent != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Exponent))
			i--
			dAtA[i] = 0x18
		}
		if len(x.DisplayDenom) > 0 {
			i -= len(x.DisplayDenom)
			copy(dAtA[i:], x.DisplayDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DisplayDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Denom) > 0 {
			i -= len(x.Denom)
			copy(dAtA[i:], x.Denom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Denom)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*EventAllowAskDenom)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventAllowAskDenom: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EventAllowAskDenom: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Denom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisplayDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DisplayDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
				}
				x.Exponent = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Exponent |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: regen/ecocredit/v1alpha2/events.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EventCreateClass is an event emitted when a credit class is created.
type EventCreateClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// class_id is the unique ID of credit class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// admin is the admin of the credit class.
	Admin string `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *EventCreateClass) Reset() {
	*x = EventCreateClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventCreateClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventCreateClass) ProtoMessage() {}

// Deprecated: Use EventCreateClass.ProtoReflect.Descriptor instead.
func (*EventCreateClass) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP(), []int{0}
}

func (x *EventCreateClass) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *EventCreateClass) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

// EventCreateProject is an event emitted when a project is created.
type EventCreateProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project_id is the unique ID of the project.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// class_id is the unique ID of credit class for this project.
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// issuer is the issuer of the credit batches for this project.
	Issuer string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// project_location is the location of the project.
	// Full documentation can be found in MsgCreateProject.project_location.
	ProjectLocation string `protobuf:"bytes,4,opt,name=project_location,json=projectLocation,proto3" json:"project_location,omitempty"`
}

func (x *EventCreateProject) Reset() {
	*x = EventCreateProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventCreateProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventCreateProject) ProtoMessage() {}

// Deprecated: Use EventCreateProject.ProtoReflect.Descriptor instead.
func (*EventCreateProject) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP(), []int{1}
}

func (x *EventCreateProject) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *EventCreateProject) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *EventCreateProject) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *EventCreateProject) GetProjectLocation() string {
	if x != nil {
		return x.ProjectLocation
	}
	return ""
}

// EventCreateBatch is an event emitted when a credit batch is created.
type EventCreateBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// class_id is the unique ID of credit class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// batch_denom is the unique ID of credit batch.
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// issuer is the account address of the issuer of the credit batch.
	Issuer string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// total_amount is the total number of credits in the credit batch.
	TotalAmount string `protobuf:"bytes,4,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	// start_date is the beginning of the period during which this credit batch
	// was quantified and verified.
	StartDate string `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// end_date is the end of the period during which this credit batch was
	// quantified and verified.
	EndDate string `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// project_location is the location of the project.
	// Full documentation can be found in MsgCreateProject.project_location.
	ProjectLocation string `protobuf:"bytes,7,opt,name=project_location,json=projectLocation,proto3" json:"project_location,omitempty"`
	// project_id is the unique ID of the project this batch belongs to.
	ProjectId string `protobuf:"bytes,8,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *EventCreateBatch) Reset() {
	*x = EventCreateBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventCreateBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventCreateBatch) ProtoMessage() {}

// Deprecated: Use EventCreateBatch.ProtoReflect.Descriptor instead.
func (*EventCreateBatch) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP(), []int{2}
}

func (x *EventCreateBatch) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *EventCreateBatch) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *EventCreateBatch) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *EventCreateBatch) GetTotalAmount() string {
	if x != nil {
		return x.TotalAmount
	}
	return ""
}

func (x *EventCreateBatch) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *EventCreateBatch) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *EventCreateBatch) GetProjectLocation() string {
	if x != nil {
		return x.ProjectLocation
	}
	return ""
}

func (x *EventCreateBatch) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

// EventReceive is an event emitted when credits are received either upon
// creation of a new batch or upon transfer. Each batch_denom created or
// transferred will result in a separate EventReceive for easy indexing.
type EventReceive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sender is the sender of the credits in the case that this event is the
	// result of a transfer. It will not be set when credits are received at
	// initial issuance.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// recipient is the recipient of the credits
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// batch_denom is the unique ID of credit batch.
	BatchDenom string `protobuf:"bytes,3,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// tradable_amount is the decimal number of tradable credits received.
	TradableAmount string `protobuf:"bytes,4,opt,name=tradable_amount,json=tradableAmount,proto3" json:"tradable_amount,omitempty"`
	// retired_amount is the decimal number of retired credits received.
	RetiredAmount string `protobuf:"bytes,5,opt,name=retired_amount,json=retiredAmount,proto3" json:"retired_amount,omitempty"`
}

func (x *EventReceive) Reset() {
	*x = EventReceive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventReceive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventReceive) ProtoMessage() {}

// Deprecated: Use EventReceive.ProtoReflect.Descriptor instead.
func (*EventReceive) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP(), []int{3}
}

func (x *EventReceive) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *EventReceive) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *EventReceive) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *EventReceive) GetTradableAmount() string {
	if x != nil {
		return x.TradableAmount
	}
	return ""
}

func (x *EventReceive) GetRetiredAmount() string {
	if x != nil {
		return x.RetiredAmount
	}
	return ""
}

// EventRetire is an event emitted when credits are retired. When credits are
// retired from multiple batches in the same transaction, a separate event is
// emitted for each batch_denom. This allows for easier indexing.
type EventRetire struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// retirer is the account which has done the "retiring". This will be the
	// account receiving credits in the case that credits were retired upon
	// issuance using Msg/CreateBatch or retired upon transfer using Msg/Send.
	Retirer string `protobuf:"bytes,1,opt,name=retirer,proto3" json:"retirer,omitempty"`
	// batch_denom is the unique ID of credit batch.
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// amount is the decimal number of credits that have been retired.
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// location is the location of the beneficiary or buyer of the retired
	// credits. It is a string of the form
	// <country-code>[-<sub-national-code>[ <postal-code>]], with the first two
	// fields conforming to ISO 3166-2, and postal-code being up to 64
	// alphanumeric characters.
	Location string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *EventRetire) Reset() {
	*x = EventRetire{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRetire) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRetire) ProtoMessage() {}

// Deprecated: Use EventRetire.ProtoReflect.Descriptor instead.
func (*EventRetire) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP(), []int{4}
}

func (x *EventRetire) GetRetirer() string {
	if x != nil {
		return x.Retirer
	}
	return ""
}

func (x *EventRetire) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *EventRetire) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *EventRetire) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

// EventCancel is an event emitted when credits are cancelled. When credits are
// cancelled from multiple batches in the same transaction, a separate event is
// emitted for each batch_denom. This allows for easier indexing.
type EventCancel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// canceller is the account which has cancelled the credits, which should be
	// the holder of the credits.
	Canceller string `protobuf:"bytes,1,opt,name=canceller,proto3" json:"canceller,omitempty"`
	// batch_denom is the unique ID of credit batch.
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// amount is the decimal number of credits that have been cancelled.
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *EventCancel) Reset() {
	*x = EventCancel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventCancel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventCancel) ProtoMessage() {}

// Deprecated: Use EventCancel.ProtoReflect.Descriptor instead.
func (*EventCancel) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP(), []int{5}
}

func (x *EventCancel) GetCanceller() string {
	if x != nil {
		return x.Canceller
	}
	return ""
}

func (x *EventCancel) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *EventCancel) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

// EventSell is an event emitted when a sell order is created.
type EventSell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// order_id is the unique ID of sell order.
	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// batch_denom is the credit batch being sold.
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// quantity is the quantity of credits being sold.
	Quantity string `protobuf:"bytes,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// ask_price is the price the seller is asking for each unit of the
	// batch_denom. Each credit unit of the batch will be sold for at least the
	// ask_price or more.
	AskPrice *v1beta1.Coin `protobuf:"bytes,4,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	// disable_auto_retire disables auto-retirement of credits which allows a
	// buyer to disable auto-retirement in their buy order enabling them to
	// resell the credits to another buyer.
	DisableAutoRetire bool `protobuf:"varint,5,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// expiration is an optional timestamp when the sell order expires. When the
	// expiration time is reached, the sell order is removed from state.
	Expiration *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *EventSell) Reset() {
	*x = EventSell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSell) ProtoMessage() {}

// Deprecated: Use EventSell.ProtoReflect.Descriptor instead.
func (*EventSell) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP(), []int{6}
}

func (x *EventSell) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *EventSell) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *EventSell) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *EventSell) GetAskPrice() *v1beta1.Coin {
	if x != nil {
		return x.AskPrice
	}
	return nil
}

func (x *EventSell) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *EventSell) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

// EventUpdateSellOrder is an event emitted when a sell order is updated.
type EventUpdateSellOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// owner is the owner of the sell orders.
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	//  sell_order_id is the ID of an existing sell order.
	SellOrderId uint64 `protobuf:"varint,2,opt,name=sell_order_id,json=sellOrderId,proto3" json:"sell_order_id,omitempty"`
	// batch_denom is the credit batch being sold.
	BatchDenom string `protobuf:"bytes,3,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// new_quantity is the updated quantity of credits available to sell, if it
	// is set to zero then the order is cancelled.
	NewQuantity string `protobuf:"bytes,4,opt,name=new_quantity,json=newQuantity,proto3" json:"new_quantity,omitempty"`
	// new_ask_price is the new ask price for this sell order
	NewAskPrice *v1beta1.Coin `protobuf:"bytes,5,opt,name=new_ask_price,json=newAskPrice,proto3" json:"new_ask_price,omitempty"`
	// disable_auto_retire updates the disable_auto_retire field in the sell order.
	DisableAutoRetire bool `protobuf:"varint,6,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// new_expiration is an optional timestamp when the sell order expires. When the
	// expiration time is reached, the sell order is removed from state.
	NewExpiration *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=new_expiration,json=newExpiration,proto3" json:"new_expiration,omitempty"`
}

func (x *EventUpdateSellOrder) Reset() {
	*x = EventUpdateSellOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventUpdateSellOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventUpdateSellOrder) ProtoMessage() {}

// Deprecated: Use EventUpdateSellOrder.ProtoReflect.Descriptor instead.
func (*EventUpdateSellOrder) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP(), []int{7}
}

func (x *EventUpdateSellOrder) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *EventUpdateSellOrder) GetSellOrderId() uint64 {
	if x != nil {
		return x.SellOrderId
	}
	return 0
}

func (x *EventUpdateSellOrder) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *EventUpdateSellOrder) GetNewQuantity() string {
	if x != nil {
		return x.NewQuantity
	}
	return ""
}

func (x *EventUpdateSellOrder) GetNewAskPrice() *v1beta1.Coin {
	if x != nil {
		return x.NewAskPrice
	}
	return nil
}

func (x *EventUpdateSellOrder) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *EventUpdateSellOrder) GetNewExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.NewExpiration
	}
	return nil
}

// EventBuyOrderCreated is an event emitted when a buy order is created.
type EventBuyOrderCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// buy_order_id is the unique ID of buy order.
	BuyOrderId uint64 `protobuf:"varint,1,opt,name=buy_order_id,json=buyOrderId,proto3" json:"buy_order_id,omitempty"`
	// sell_order_id is the sell order ID against which the buyer is trying to buy.
	SellOrderId uint64 `protobuf:"varint,2,opt,name=sell_order_id,json=sellOrderId,proto3" json:"sell_order_id,omitempty"`
	// quantity is the quantity of credits to buy. If the quantity of credits
	// available is less than this amount the order will be partially filled
	// unless disable_partial_fill is true.
	Quantity string `protobuf:"bytes,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// bid price is the bid price for this buy order. A credit unit will be
	// settled at a purchase price that is no more than the bid price. The
	// buy order will fail if the buyer does not have enough funds available
	// to complete the purchase.
	BidPrice *v1beta1.Coin `protobuf:"bytes,4,opt,name=bid_price,json=bidPrice,proto3" json:"bid_price,omitempty"`
	// disable_auto_retire allows auto-retirement to be disabled. If it is set to true
	// the credits will not auto-retire and can be resold assuming that the
	// corresponding sell order has auto-retirement disabled. If the sell order
	// hasn't disabled auto-retirement and the buy order tries to disable it,
	// that buy order will fail.
	DisableAutoRetire bool `protobuf:"varint,5,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// disable_partial_fill disables the default behavior of partially filling
	// buy orders if the requested quantity is not available.
	DisablePartialFill bool `protobuf:"varint,6,opt,name=disable_partial_fill,json=disablePartialFill,proto3" json:"disable_partial_fill,omitempty"`
	// retirement_location is the optional retirement location for the credits
	// which will be used only if disable_auto_retire is false.
	RetirementLocation string `protobuf:"bytes,7,opt,name=retirement_location,json=retirementLocation,proto3" json:"retirement_location,omitempty"`
	// expiration is the optional timestamp when the buy order expires. When the
	// expiration time is reached, the buy order is removed from state.
	Expiration *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *EventBuyOrderCreated) Reset() {
	*x = EventBuyOrderCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventBuyOrderCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventBuyOrderCreated) ProtoMessage() {}

// Deprecated: Use EventBuyOrderCreated.ProtoReflect.Descriptor instead.
func (*EventBuyOrderCreated) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP(), []int{8}
}

func (x *EventBuyOrderCreated) GetBuyOrderId() uint64 {
	if x != nil {
		return x.BuyOrderId
	}
	return 0
}

func (x *EventBuyOrderCreated) GetSellOrderId() uint64 {
	if x != nil {
		return x.SellOrderId
	}
	return 0
}

func (x *EventBuyOrderCreated) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *EventBuyOrderCreated) GetBidPrice() *v1beta1.Coin {
	if x != nil {
		return x.BidPrice
	}
	return nil
}

func (x *EventBuyOrderCreated) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *EventBuyOrderCreated) GetDisablePartialFill() bool {
	if x != nil {
		return x.DisablePartialFill
	}
	return false
}

func (x *EventBuyOrderCreated) GetRetirementLocation() string {
	if x != nil {
		return x.RetirementLocation
	}
	return ""
}

func (x *EventBuyOrderCreated) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

// EventBuyOrderFilled is an event emitted when a buy order is filled.
type EventBuyOrderFilled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// buy_order_id is the unique ID of the buy order.
	BuyOrderId uint64 `protobuf:"varint,1,opt,name=buy_order_id,json=buyOrderId,proto3" json:"buy_order_id,omitempty"`
	// sell_order_id is the unique ID of the sell order.
	SellOrderId uint64 `protobuf:"varint,2,opt,name=sell_order_id,json=sellOrderId,proto3" json:"sell_order_id,omitempty"`
	// batch_denom is the credit batch ID of the purchased credits.
	BatchDenom string `protobuf:"bytes,3,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// quantity is the quantity of the purchased credits.
	Quantity string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// total_price is the total price for the purchased credits.
	TotalPrice *v1beta1.Coin `protobuf:"bytes,5,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *EventBuyOrderFilled) Reset() {
	*x = EventBuyOrderFilled{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventBuyOrderFilled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventBuyOrderFilled) ProtoMessage() {}

// Deprecated: Use EventBuyOrderFilled.ProtoReflect.Descriptor instead.
func (*EventBuyOrderFilled) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP(), []int{9}
}

func (x *EventBuyOrderFilled) GetBuyOrderId() uint64 {
	if x != nil {
		return x.BuyOrderId
	}
	return 0
}

func (x *EventBuyOrderFilled) GetSellOrderId() uint64 {
	if x != nil {
		return x.SellOrderId
	}
	return 0
}

func (x *EventBuyOrderFilled) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *EventBuyOrderFilled) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *EventBuyOrderFilled) GetTotalPrice() *v1beta1.Coin {
	if x != nil {
		return x.TotalPrice
	}
	return nil
}

// EventAllowAskDenom is an event emitted when an ask denom is added.
type EventAllowAskDenom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// denom is the denom to allow (ex. ibc/GLKHDSG423SGS)
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// display_denom is the denom to display to the user and is informational
	DisplayDenom string `protobuf:"bytes,2,opt,name=display_denom,json=displayDenom,proto3" json:"display_denom,omitempty"`
	// exponent is the exponent that relates the denom to the display_denom and is
	// informational
	Exponent uint32 `protobuf:"varint,3,opt,name=exponent,proto3" json:"exponent,omitempty"`
}

func (x *EventAllowAskDenom) Reset() {
	*x = EventAllowAskDenom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_events_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventAllowAskDenom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventAllowAskDenom) ProtoMessage() {}

// Deprecated: Use EventAllowAskDenom.ProtoReflect.Descriptor instead.
func (*EventAllowAskDenom) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP(), []int{10}
}

func (x *EventAllowAskDenom) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *EventAllowAskDenom) GetDisplayDenom() string {
	if x != nil {
		return x.DisplayDenom
	}
	return ""
}

func (x *EventAllowAskDenom) GetExponent() uint32 {
	if x != nil {
		return x.Exponent
	}
	return 0
}

var File_regen_ecocredit_v1alpha2_events_proto protoreflect.FileDescriptor

var file_regen_ecocredit_v1alpha2_events_proto_rawDesc = []byte{
	0x0a, 0x25, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65,
	0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x91, 0x01,
	0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x8d, 0x02, 0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e,
	0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x22, 0xb5, 0x01, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x64, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x69,
	0x72, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7c, 0x0a, 0x0b, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x69,
	0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x74, 0x69, 0x72,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65,
	0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x64, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8d, 0x02,
	0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x69,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f,
	0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xcc, 0x02,
	0x0a, 0x14, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c,
	0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d,
	0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f,
	0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x73, 0x6b, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x41, 0x73, 0x6b, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74,
	0x69, 0x72, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0d, 0x6e,
	0x65, 0x77, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x85, 0x03, 0x0a,
	0x14, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x75, 0x79,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x08, 0x62, 0x69, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f,
	0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12,
	0x30, 0x0a, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x69, 0x6c,
	0x6c, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd4, 0x01, 0x0a, 0x13, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x75,
	0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0c,
	0x62, 0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x62, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65,
	0x6e, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x3a, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x6b, 0x0a, 0x12, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x73, 0x6b, 0x44, 0x65, 0x6e, 0x6f,
	0x6d, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x42, 0x83, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x3b, 0x65, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2, 0x02,
	0x03, 0x52, 0x45, 0x58, 0xaa, 0x02, 0x18, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x45, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xca,
	0x02, 0x18, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xe2, 0x02, 0x24, 0x52, 0x65, 0x67,
	0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1a, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_v1alpha2_events_proto_rawDescOnce sync.Once
	file_regen_ecocredit_v1alpha2_events_proto_rawDescData = file_regen_ecocredit_v1alpha2_events_proto_rawDesc
)

func file_regen_ecocredit_v1alpha2_events_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_v1alpha2_events_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_v1alpha2_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_v1alpha2_events_proto_rawDescData)
	})
	return file_regen_ecocredit_v1alpha2_events_proto_rawDescData
}

var file_regen_ecocredit_v1alpha2_events_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_regen_ecocredit_v1alpha2_events_proto_goTypes = []interface{}{
	(*EventCreateClass)(nil),      // 0: regen.ecocredit.v1alpha2.EventCreateClass
	(*EventCreateProject)(nil),    // 1: regen.ecocredit.v1alpha2.EventCreateProject
	(*EventCreateBatch)(nil),      // 2: regen.ecocredit.v1alpha2.EventCreateBatch
	(*EventReceive)(nil),          // 3: regen.ecocredit.v1alpha2.EventReceive
	(*EventRetire)(nil),           // 4: regen.ecocredit.v1alpha2.EventRetire
	(*EventCancel)(nil),           // 5: regen.ecocredit.v1alpha2.EventCancel
	(*EventSell)(nil),             // 6: regen.ecocredit.v1alpha2.EventSell
	(*EventUpdateSellOrder)(nil),  // 7: regen.ecocredit.v1alpha2.EventUpdateSellOrder
	(*EventBuyOrderCreated)(nil),  // 8: regen.ecocredit.v1alpha2.EventBuyOrderCreated
	(*EventBuyOrderFilled)(nil),   // 9: regen.ecocredit.v1alpha2.EventBuyOrderFilled
	(*EventAllowAskDenom)(nil),    // 10: regen.ecocredit.v1alpha2.EventAllowAskDenom
	(*v1beta1.Coin)(nil),          // 11: cosmos.base.v1beta1.Coin
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_regen_ecocredit_v1alpha2_events_proto_depIdxs = []int32{
	11, // 0: regen.ecocredit.v1alpha2.EventSell.ask_price:type_name -> cosmos.base.v1beta1.Coin
	12, // 1: regen.ecocredit.v1alpha2.EventSell.expiration:type_name -> google.protobuf.Timestamp
	11, // 2: regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_ask_price:type_name -> cosmos.base.v1beta1.Coin
	12, // 3: regen.ecocredit.v1alpha2.EventUpdateSellOrder.new_expiration:type_name -> google.protobuf.Timestamp
	11, // 4: regen.ecocredit.v1alpha2.EventBuyOrderCreated.bid_price:type_name -> cosmos.base.v1beta1.Coin
	12, // 5: regen.ecocredit.v1alpha2.EventBuyOrderCreated.expiration:type_name -> google.protobuf.Timestamp
	11, // 6: regen.ecocredit.v1alpha2.EventBuyOrderFilled.total_price:type_name -> cosmos.base.v1beta1.Coin
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_v1alpha2_events_proto_init() }
func file_regen_ecocredit_v1alpha2_events_proto_init() {
	if File_regen_ecocredit_v1alpha2_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_v1alpha2_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventCreateClass); i {
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
		file_regen_ecocredit_v1alpha2_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventCreateProject); i {
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
		file_regen_ecocredit_v1alpha2_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventCreateBatch); i {
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
		file_regen_ecocredit_v1alpha2_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventReceive); i {
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
		file_regen_ecocredit_v1alpha2_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRetire); i {
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
		file_regen_ecocredit_v1alpha2_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventCancel); i {
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
		file_regen_ecocredit_v1alpha2_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventSell); i {
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
		file_regen_ecocredit_v1alpha2_events_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventUpdateSellOrder); i {
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
		file_regen_ecocredit_v1alpha2_events_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventBuyOrderCreated); i {
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
		file_regen_ecocredit_v1alpha2_events_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventBuyOrderFilled); i {
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
		file_regen_ecocredit_v1alpha2_events_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventAllowAskDenom); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_v1alpha2_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_ecocredit_v1alpha2_events_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_v1alpha2_events_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_v1alpha2_events_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_v1alpha2_events_proto = out.File
	file_regen_ecocredit_v1alpha2_events_proto_rawDesc = nil
	file_regen_ecocredit_v1alpha2_events_proto_goTypes = nil
	file_regen_ecocredit_v1alpha2_events_proto_depIdxs = nil
}