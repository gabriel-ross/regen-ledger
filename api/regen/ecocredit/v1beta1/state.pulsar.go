package ecocreditv1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_CreditType              protoreflect.MessageDescriptor
	fd_CreditType_abbreviation protoreflect.FieldDescriptor
	fd_CreditType_name         protoreflect.FieldDescriptor
	fd_CreditType_unit         protoreflect.FieldDescriptor
	fd_CreditType_precision    protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_state_proto_init()
	md_CreditType = File_regen_ecocredit_v1beta1_state_proto.Messages().ByName("CreditType")
	fd_CreditType_abbreviation = md_CreditType.Fields().ByName("abbreviation")
	fd_CreditType_name = md_CreditType.Fields().ByName("name")
	fd_CreditType_unit = md_CreditType.Fields().ByName("unit")
	fd_CreditType_precision = md_CreditType.Fields().ByName("precision")
}

var _ protoreflect.Message = (*fastReflection_CreditType)(nil)

type fastReflection_CreditType CreditType

func (x *CreditType) ProtoReflect() protoreflect.Message {
	return (*fastReflection_CreditType)(x)
}

func (x *CreditType) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_CreditType_messageType fastReflection_CreditType_messageType
var _ protoreflect.MessageType = fastReflection_CreditType_messageType{}

type fastReflection_CreditType_messageType struct{}

func (x fastReflection_CreditType_messageType) Zero() protoreflect.Message {
	return (*fastReflection_CreditType)(nil)
}
func (x fastReflection_CreditType_messageType) New() protoreflect.Message {
	return new(fastReflection_CreditType)
}
func (x fastReflection_CreditType_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_CreditType
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_CreditType) Descriptor() protoreflect.MessageDescriptor {
	return md_CreditType
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_CreditType) Type() protoreflect.MessageType {
	return _fastReflection_CreditType_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_CreditType) New() protoreflect.Message {
	return new(fastReflection_CreditType)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_CreditType) Interface() protoreflect.ProtoMessage {
	return (*CreditType)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_CreditType) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Abbreviation != "" {
		value := protoreflect.ValueOfString(x.Abbreviation)
		if !f(fd_CreditType_abbreviation, value) {
			return
		}
	}
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_CreditType_name, value) {
			return
		}
	}
	if x.Unit != "" {
		value := protoreflect.ValueOfString(x.Unit)
		if !f(fd_CreditType_unit, value) {
			return
		}
	}
	if x.Precision != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Precision)
		if !f(fd_CreditType_precision, value) {
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
func (x *fastReflection_CreditType) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.CreditType.abbreviation":
		return x.Abbreviation != ""
	case "regen.ecocredit.v1beta1.CreditType.name":
		return x.Name != ""
	case "regen.ecocredit.v1beta1.CreditType.unit":
		return x.Unit != ""
	case "regen.ecocredit.v1beta1.CreditType.precision":
		return x.Precision != uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.CreditType does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CreditType) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.CreditType.abbreviation":
		x.Abbreviation = ""
	case "regen.ecocredit.v1beta1.CreditType.name":
		x.Name = ""
	case "regen.ecocredit.v1beta1.CreditType.unit":
		x.Unit = ""
	case "regen.ecocredit.v1beta1.CreditType.precision":
		x.Precision = uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.CreditType does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_CreditType) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.CreditType.abbreviation":
		value := x.Abbreviation
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.CreditType.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.CreditType.unit":
		value := x.Unit
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.CreditType.precision":
		value := x.Precision
		return protoreflect.ValueOfUint32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.CreditType does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_CreditType) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.CreditType.abbreviation":
		x.Abbreviation = value.Interface().(string)
	case "regen.ecocredit.v1beta1.CreditType.name":
		x.Name = value.Interface().(string)
	case "regen.ecocredit.v1beta1.CreditType.unit":
		x.Unit = value.Interface().(string)
	case "regen.ecocredit.v1beta1.CreditType.precision":
		x.Precision = uint32(value.Uint())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.CreditType does not contain field %s", fd.FullName()))
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
func (x *fastReflection_CreditType) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.CreditType.abbreviation":
		panic(fmt.Errorf("field abbreviation of message regen.ecocredit.v1beta1.CreditType is not mutable"))
	case "regen.ecocredit.v1beta1.CreditType.name":
		panic(fmt.Errorf("field name of message regen.ecocredit.v1beta1.CreditType is not mutable"))
	case "regen.ecocredit.v1beta1.CreditType.unit":
		panic(fmt.Errorf("field unit of message regen.ecocredit.v1beta1.CreditType is not mutable"))
	case "regen.ecocredit.v1beta1.CreditType.precision":
		panic(fmt.Errorf("field precision of message regen.ecocredit.v1beta1.CreditType is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.CreditType does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_CreditType) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.CreditType.abbreviation":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.CreditType.name":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.CreditType.unit":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.CreditType.precision":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.CreditType does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_CreditType) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.CreditType", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_CreditType) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CreditType) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_CreditType) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_CreditType) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*CreditType)
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
		l = len(x.Abbreviation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Unit)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Precision != 0 {
			n += 1 + runtime.Sov(uint64(x.Precision))
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
		x := input.Message.Interface().(*CreditType)
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
		if x.Precision != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Precision))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Unit) > 0 {
			i -= len(x.Unit)
			copy(dAtA[i:], x.Unit)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Unit)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Abbreviation) > 0 {
			i -= len(x.Abbreviation)
			copy(dAtA[i:], x.Abbreviation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Abbreviation)))
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
		x := input.Message.Interface().(*CreditType)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CreditType: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CreditType: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Abbreviation", wireType)
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
				x.Abbreviation = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
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
				x.Unit = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
				}
				x.Precision = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Precision |= uint32(b&0x7F) << shift
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

var (
	md_ClassInfo             protoreflect.MessageDescriptor
	fd_ClassInfo_id          protoreflect.FieldDescriptor
	fd_ClassInfo_name        protoreflect.FieldDescriptor
	fd_ClassInfo_admin       protoreflect.FieldDescriptor
	fd_ClassInfo_metadata    protoreflect.FieldDescriptor
	fd_ClassInfo_credit_type protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_state_proto_init()
	md_ClassInfo = File_regen_ecocredit_v1beta1_state_proto.Messages().ByName("ClassInfo")
	fd_ClassInfo_id = md_ClassInfo.Fields().ByName("id")
	fd_ClassInfo_name = md_ClassInfo.Fields().ByName("name")
	fd_ClassInfo_admin = md_ClassInfo.Fields().ByName("admin")
	fd_ClassInfo_metadata = md_ClassInfo.Fields().ByName("metadata")
	fd_ClassInfo_credit_type = md_ClassInfo.Fields().ByName("credit_type")
}

var _ protoreflect.Message = (*fastReflection_ClassInfo)(nil)

type fastReflection_ClassInfo ClassInfo

func (x *ClassInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ClassInfo)(x)
}

func (x *ClassInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ClassInfo_messageType fastReflection_ClassInfo_messageType
var _ protoreflect.MessageType = fastReflection_ClassInfo_messageType{}

type fastReflection_ClassInfo_messageType struct{}

func (x fastReflection_ClassInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ClassInfo)(nil)
}
func (x fastReflection_ClassInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_ClassInfo)
}
func (x fastReflection_ClassInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ClassInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ClassInfo) Type() protoreflect.MessageType {
	return _fastReflection_ClassInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ClassInfo) New() protoreflect.Message {
	return new(fastReflection_ClassInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ClassInfo) Interface() protoreflect.ProtoMessage {
	return (*ClassInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ClassInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_ClassInfo_id, value) {
			return
		}
	}
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_ClassInfo_name, value) {
			return
		}
	}
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_ClassInfo_admin, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_ClassInfo_metadata, value) {
			return
		}
	}
	if x.CreditType != "" {
		value := protoreflect.ValueOfString(x.CreditType)
		if !f(fd_ClassInfo_credit_type, value) {
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
func (x *fastReflection_ClassInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassInfo.id":
		return x.Id != uint64(0)
	case "regen.ecocredit.v1beta1.ClassInfo.name":
		return x.Name != ""
	case "regen.ecocredit.v1beta1.ClassInfo.admin":
		return x.Admin != ""
	case "regen.ecocredit.v1beta1.ClassInfo.metadata":
		return len(x.Metadata) != 0
	case "regen.ecocredit.v1beta1.ClassInfo.credit_type":
		return x.CreditType != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassInfo.id":
		x.Id = uint64(0)
	case "regen.ecocredit.v1beta1.ClassInfo.name":
		x.Name = ""
	case "regen.ecocredit.v1beta1.ClassInfo.admin":
		x.Admin = ""
	case "regen.ecocredit.v1beta1.ClassInfo.metadata":
		x.Metadata = nil
	case "regen.ecocredit.v1beta1.ClassInfo.credit_type":
		x.CreditType = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ClassInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.ClassInfo.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1beta1.ClassInfo.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.ClassInfo.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.ClassInfo.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.v1beta1.ClassInfo.credit_type":
		value := x.CreditType
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassInfo does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ClassInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassInfo.id":
		x.Id = value.Uint()
	case "regen.ecocredit.v1beta1.ClassInfo.name":
		x.Name = value.Interface().(string)
	case "regen.ecocredit.v1beta1.ClassInfo.admin":
		x.Admin = value.Interface().(string)
	case "regen.ecocredit.v1beta1.ClassInfo.metadata":
		x.Metadata = value.Bytes()
	case "regen.ecocredit.v1beta1.ClassInfo.credit_type":
		x.CreditType = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassInfo does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ClassInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassInfo.id":
		panic(fmt.Errorf("field id of message regen.ecocredit.v1beta1.ClassInfo is not mutable"))
	case "regen.ecocredit.v1beta1.ClassInfo.name":
		panic(fmt.Errorf("field name of message regen.ecocredit.v1beta1.ClassInfo is not mutable"))
	case "regen.ecocredit.v1beta1.ClassInfo.admin":
		panic(fmt.Errorf("field admin of message regen.ecocredit.v1beta1.ClassInfo is not mutable"))
	case "regen.ecocredit.v1beta1.ClassInfo.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1beta1.ClassInfo is not mutable"))
	case "regen.ecocredit.v1beta1.ClassInfo.credit_type":
		panic(fmt.Errorf("field credit_type of message regen.ecocredit.v1beta1.ClassInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ClassInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassInfo.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1beta1.ClassInfo.name":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.ClassInfo.admin":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.ClassInfo.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.v1beta1.ClassInfo.credit_type":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ClassInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.ClassInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ClassInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassInfo) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ClassInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ClassInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ClassInfo)
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
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.CreditType)
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
		x := input.Message.Interface().(*ClassInfo)
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
		if len(x.CreditType) > 0 {
			i -= len(x.CreditType)
			copy(dAtA[i:], x.CreditType)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CreditType)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
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
		x := input.Message.Interface().(*ClassInfo)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
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
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreditType", wireType)
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
				x.CreditType = string(dAtA[iNdEx:postIndex])
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
	md_ClassIssuer          protoreflect.MessageDescriptor
	fd_ClassIssuer_class_id protoreflect.FieldDescriptor
	fd_ClassIssuer_issuer   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_state_proto_init()
	md_ClassIssuer = File_regen_ecocredit_v1beta1_state_proto.Messages().ByName("ClassIssuer")
	fd_ClassIssuer_class_id = md_ClassIssuer.Fields().ByName("class_id")
	fd_ClassIssuer_issuer = md_ClassIssuer.Fields().ByName("issuer")
}

var _ protoreflect.Message = (*fastReflection_ClassIssuer)(nil)

type fastReflection_ClassIssuer ClassIssuer

func (x *ClassIssuer) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ClassIssuer)(x)
}

func (x *ClassIssuer) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ClassIssuer_messageType fastReflection_ClassIssuer_messageType
var _ protoreflect.MessageType = fastReflection_ClassIssuer_messageType{}

type fastReflection_ClassIssuer_messageType struct{}

func (x fastReflection_ClassIssuer_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ClassIssuer)(nil)
}
func (x fastReflection_ClassIssuer_messageType) New() protoreflect.Message {
	return new(fastReflection_ClassIssuer)
}
func (x fastReflection_ClassIssuer_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassIssuer
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ClassIssuer) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassIssuer
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ClassIssuer) Type() protoreflect.MessageType {
	return _fastReflection_ClassIssuer_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ClassIssuer) New() protoreflect.Message {
	return new(fastReflection_ClassIssuer)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ClassIssuer) Interface() protoreflect.ProtoMessage {
	return (*ClassIssuer)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ClassIssuer) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_ClassIssuer_class_id, value) {
			return
		}
	}
	if x.Issuer != "" {
		value := protoreflect.ValueOfString(x.Issuer)
		if !f(fd_ClassIssuer_issuer, value) {
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
func (x *fastReflection_ClassIssuer) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassIssuer.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1beta1.ClassIssuer.issuer":
		return x.Issuer != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassIssuer"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassIssuer does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassIssuer) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassIssuer.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1beta1.ClassIssuer.issuer":
		x.Issuer = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassIssuer"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassIssuer does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ClassIssuer) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.ClassIssuer.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.ClassIssuer.issuer":
		value := x.Issuer
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassIssuer"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassIssuer does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ClassIssuer) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassIssuer.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1beta1.ClassIssuer.issuer":
		x.Issuer = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassIssuer"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassIssuer does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ClassIssuer) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassIssuer.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1beta1.ClassIssuer is not mutable"))
	case "regen.ecocredit.v1beta1.ClassIssuer.issuer":
		panic(fmt.Errorf("field issuer of message regen.ecocredit.v1beta1.ClassIssuer is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassIssuer"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassIssuer does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ClassIssuer) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassIssuer.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.ClassIssuer.issuer":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassIssuer"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassIssuer does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ClassIssuer) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.ClassIssuer", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ClassIssuer) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassIssuer) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ClassIssuer) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ClassIssuer) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ClassIssuer)
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
		l = len(x.Issuer)
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
		x := input.Message.Interface().(*ClassIssuer)
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
		if len(x.Issuer) > 0 {
			i -= len(x.Issuer)
			copy(dAtA[i:], x.Issuer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuer)))
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
		x := input.Message.Interface().(*ClassIssuer)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassIssuer: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassIssuer: illegal tag %d (wire type %d)", fieldNum, wire)
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
	md_ProjectInfo                  protoreflect.MessageDescriptor
	fd_ProjectInfo_id               protoreflect.FieldDescriptor
	fd_ProjectInfo_name             protoreflect.FieldDescriptor
	fd_ProjectInfo_class_id         protoreflect.FieldDescriptor
	fd_ProjectInfo_project_location protoreflect.FieldDescriptor
	fd_ProjectInfo_metadata         protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_state_proto_init()
	md_ProjectInfo = File_regen_ecocredit_v1beta1_state_proto.Messages().ByName("ProjectInfo")
	fd_ProjectInfo_id = md_ProjectInfo.Fields().ByName("id")
	fd_ProjectInfo_name = md_ProjectInfo.Fields().ByName("name")
	fd_ProjectInfo_class_id = md_ProjectInfo.Fields().ByName("class_id")
	fd_ProjectInfo_project_location = md_ProjectInfo.Fields().ByName("project_location")
	fd_ProjectInfo_metadata = md_ProjectInfo.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_ProjectInfo)(nil)

type fastReflection_ProjectInfo ProjectInfo

func (x *ProjectInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ProjectInfo)(x)
}

func (x *ProjectInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ProjectInfo_messageType fastReflection_ProjectInfo_messageType
var _ protoreflect.MessageType = fastReflection_ProjectInfo_messageType{}

type fastReflection_ProjectInfo_messageType struct{}

func (x fastReflection_ProjectInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ProjectInfo)(nil)
}
func (x fastReflection_ProjectInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_ProjectInfo)
}
func (x fastReflection_ProjectInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ProjectInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ProjectInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_ProjectInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ProjectInfo) Type() protoreflect.MessageType {
	return _fastReflection_ProjectInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ProjectInfo) New() protoreflect.Message {
	return new(fastReflection_ProjectInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ProjectInfo) Interface() protoreflect.ProtoMessage {
	return (*ProjectInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ProjectInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_ProjectInfo_id, value) {
			return
		}
	}
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_ProjectInfo_name, value) {
			return
		}
	}
	if x.ClassId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ClassId)
		if !f(fd_ProjectInfo_class_id, value) {
			return
		}
	}
	if x.ProjectLocation != "" {
		value := protoreflect.ValueOfString(x.ProjectLocation)
		if !f(fd_ProjectInfo_project_location, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_ProjectInfo_metadata, value) {
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
func (x *fastReflection_ProjectInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ProjectInfo.id":
		return x.Id != uint64(0)
	case "regen.ecocredit.v1beta1.ProjectInfo.name":
		return x.Name != ""
	case "regen.ecocredit.v1beta1.ProjectInfo.class_id":
		return x.ClassId != uint64(0)
	case "regen.ecocredit.v1beta1.ProjectInfo.project_location":
		return x.ProjectLocation != ""
	case "regen.ecocredit.v1beta1.ProjectInfo.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ProjectInfo.id":
		x.Id = uint64(0)
	case "regen.ecocredit.v1beta1.ProjectInfo.name":
		x.Name = ""
	case "regen.ecocredit.v1beta1.ProjectInfo.class_id":
		x.ClassId = uint64(0)
	case "regen.ecocredit.v1beta1.ProjectInfo.project_location":
		x.ProjectLocation = ""
	case "regen.ecocredit.v1beta1.ProjectInfo.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ProjectInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.ProjectInfo.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1beta1.ProjectInfo.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.ProjectInfo.class_id":
		value := x.ClassId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1beta1.ProjectInfo.project_location":
		value := x.ProjectLocation
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.ProjectInfo.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectInfo does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ProjectInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ProjectInfo.id":
		x.Id = value.Uint()
	case "regen.ecocredit.v1beta1.ProjectInfo.name":
		x.Name = value.Interface().(string)
	case "regen.ecocredit.v1beta1.ProjectInfo.class_id":
		x.ClassId = value.Uint()
	case "regen.ecocredit.v1beta1.ProjectInfo.project_location":
		x.ProjectLocation = value.Interface().(string)
	case "regen.ecocredit.v1beta1.ProjectInfo.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectInfo does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ProjectInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ProjectInfo.id":
		panic(fmt.Errorf("field id of message regen.ecocredit.v1beta1.ProjectInfo is not mutable"))
	case "regen.ecocredit.v1beta1.ProjectInfo.name":
		panic(fmt.Errorf("field name of message regen.ecocredit.v1beta1.ProjectInfo is not mutable"))
	case "regen.ecocredit.v1beta1.ProjectInfo.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1beta1.ProjectInfo is not mutable"))
	case "regen.ecocredit.v1beta1.ProjectInfo.project_location":
		panic(fmt.Errorf("field project_location of message regen.ecocredit.v1beta1.ProjectInfo is not mutable"))
	case "regen.ecocredit.v1beta1.ProjectInfo.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1beta1.ProjectInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ProjectInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ProjectInfo.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1beta1.ProjectInfo.name":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.ProjectInfo.class_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1beta1.ProjectInfo.project_location":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.ProjectInfo.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ProjectInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.ProjectInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ProjectInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectInfo) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ProjectInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ProjectInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ProjectInfo)
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
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ClassId != 0 {
			n += 1 + runtime.Sov(uint64(x.ClassId))
		}
		l = len(x.ProjectLocation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
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
		x := input.Message.Interface().(*ProjectInfo)
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
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.ProjectLocation) > 0 {
			i -= len(x.ProjectLocation)
			copy(dAtA[i:], x.ProjectLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectLocation)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0x1a
		}
		if x.ClassId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ClassId))
			i--
			dAtA[i] = 0x10
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
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
		x := input.Message.Interface().(*ProjectInfo)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ProjectInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ProjectInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
				}
				x.ClassId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ClassId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
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
	md_BatchInfo             protoreflect.MessageDescriptor
	fd_BatchInfo_id          protoreflect.FieldDescriptor
	fd_BatchInfo_project_id  protoreflect.FieldDescriptor
	fd_BatchInfo_batch_denom protoreflect.FieldDescriptor
	fd_BatchInfo_metadata    protoreflect.FieldDescriptor
	fd_BatchInfo_start_date  protoreflect.FieldDescriptor
	fd_BatchInfo_end_date    protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_state_proto_init()
	md_BatchInfo = File_regen_ecocredit_v1beta1_state_proto.Messages().ByName("BatchInfo")
	fd_BatchInfo_id = md_BatchInfo.Fields().ByName("id")
	fd_BatchInfo_project_id = md_BatchInfo.Fields().ByName("project_id")
	fd_BatchInfo_batch_denom = md_BatchInfo.Fields().ByName("batch_denom")
	fd_BatchInfo_metadata = md_BatchInfo.Fields().ByName("metadata")
	fd_BatchInfo_start_date = md_BatchInfo.Fields().ByName("start_date")
	fd_BatchInfo_end_date = md_BatchInfo.Fields().ByName("end_date")
}

var _ protoreflect.Message = (*fastReflection_BatchInfo)(nil)

type fastReflection_BatchInfo BatchInfo

func (x *BatchInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BatchInfo)(x)
}

func (x *BatchInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BatchInfo_messageType fastReflection_BatchInfo_messageType
var _ protoreflect.MessageType = fastReflection_BatchInfo_messageType{}

type fastReflection_BatchInfo_messageType struct{}

func (x fastReflection_BatchInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BatchInfo)(nil)
}
func (x fastReflection_BatchInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_BatchInfo)
}
func (x fastReflection_BatchInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BatchInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BatchInfo) Type() protoreflect.MessageType {
	return _fastReflection_BatchInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BatchInfo) New() protoreflect.Message {
	return new(fastReflection_BatchInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BatchInfo) Interface() protoreflect.ProtoMessage {
	return (*BatchInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BatchInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_BatchInfo_id, value) {
			return
		}
	}
	if x.ProjectId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProjectId)
		if !f(fd_BatchInfo_project_id, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_BatchInfo_batch_denom, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_BatchInfo_metadata, value) {
			return
		}
	}
	if x.StartDate != nil {
		value := protoreflect.ValueOfMessage(x.StartDate.ProtoReflect())
		if !f(fd_BatchInfo_start_date, value) {
			return
		}
	}
	if x.EndDate != nil {
		value := protoreflect.ValueOfMessage(x.EndDate.ProtoReflect())
		if !f(fd_BatchInfo_end_date, value) {
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
func (x *fastReflection_BatchInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchInfo.id":
		return x.Id != uint64(0)
	case "regen.ecocredit.v1beta1.BatchInfo.project_id":
		return x.ProjectId != uint64(0)
	case "regen.ecocredit.v1beta1.BatchInfo.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1beta1.BatchInfo.metadata":
		return len(x.Metadata) != 0
	case "regen.ecocredit.v1beta1.BatchInfo.start_date":
		return x.StartDate != nil
	case "regen.ecocredit.v1beta1.BatchInfo.end_date":
		return x.EndDate != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchInfo.id":
		x.Id = uint64(0)
	case "regen.ecocredit.v1beta1.BatchInfo.project_id":
		x.ProjectId = uint64(0)
	case "regen.ecocredit.v1beta1.BatchInfo.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1beta1.BatchInfo.metadata":
		x.Metadata = nil
	case "regen.ecocredit.v1beta1.BatchInfo.start_date":
		x.StartDate = nil
	case "regen.ecocredit.v1beta1.BatchInfo.end_date":
		x.EndDate = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BatchInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.BatchInfo.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1beta1.BatchInfo.project_id":
		value := x.ProjectId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1beta1.BatchInfo.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.BatchInfo.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.v1beta1.BatchInfo.start_date":
		value := x.StartDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1beta1.BatchInfo.end_date":
		value := x.EndDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchInfo does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BatchInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchInfo.id":
		x.Id = value.Uint()
	case "regen.ecocredit.v1beta1.BatchInfo.project_id":
		x.ProjectId = value.Uint()
	case "regen.ecocredit.v1beta1.BatchInfo.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1beta1.BatchInfo.metadata":
		x.Metadata = value.Bytes()
	case "regen.ecocredit.v1beta1.BatchInfo.start_date":
		x.StartDate = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.ecocredit.v1beta1.BatchInfo.end_date":
		x.EndDate = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchInfo does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BatchInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchInfo.start_date":
		if x.StartDate == nil {
			x.StartDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.StartDate.ProtoReflect())
	case "regen.ecocredit.v1beta1.BatchInfo.end_date":
		if x.EndDate == nil {
			x.EndDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.EndDate.ProtoReflect())
	case "regen.ecocredit.v1beta1.BatchInfo.id":
		panic(fmt.Errorf("field id of message regen.ecocredit.v1beta1.BatchInfo is not mutable"))
	case "regen.ecocredit.v1beta1.BatchInfo.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.v1beta1.BatchInfo is not mutable"))
	case "regen.ecocredit.v1beta1.BatchInfo.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1beta1.BatchInfo is not mutable"))
	case "regen.ecocredit.v1beta1.BatchInfo.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1beta1.BatchInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BatchInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchInfo.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1beta1.BatchInfo.project_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1beta1.BatchInfo.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.BatchInfo.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.v1beta1.BatchInfo.start_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1beta1.BatchInfo.end_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BatchInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.BatchInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BatchInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchInfo) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BatchInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BatchInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BatchInfo)
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
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		if x.ProjectId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProjectId))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.StartDate != nil {
			l = options.Size(x.StartDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.EndDate != nil {
			l = options.Size(x.EndDate)
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
		x := input.Message.Interface().(*BatchInfo)
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
		if x.EndDate != nil {
			encoded, err := options.Marshal(x.EndDate)
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
		if x.StartDate != nil {
			encoded, err := options.Marshal(x.StartDate)
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
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
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
		if x.ProjectId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProjectId))
			i--
			dAtA[i] = 0x10
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
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
		x := input.Message.Interface().(*BatchInfo)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
				}
				x.ProjectId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProjectId |= uint64(b&0x7F) << shift
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
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
				if x.StartDate == nil {
					x.StartDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.StartDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
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
				if x.EndDate == nil {
					x.EndDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.EndDate); err != nil {
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
	md_ClassSequence               protoreflect.MessageDescriptor
	fd_ClassSequence_credit_type   protoreflect.FieldDescriptor
	fd_ClassSequence_next_class_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_state_proto_init()
	md_ClassSequence = File_regen_ecocredit_v1beta1_state_proto.Messages().ByName("ClassSequence")
	fd_ClassSequence_credit_type = md_ClassSequence.Fields().ByName("credit_type")
	fd_ClassSequence_next_class_id = md_ClassSequence.Fields().ByName("next_class_id")
}

var _ protoreflect.Message = (*fastReflection_ClassSequence)(nil)

type fastReflection_ClassSequence ClassSequence

func (x *ClassSequence) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ClassSequence)(x)
}

func (x *ClassSequence) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ClassSequence_messageType fastReflection_ClassSequence_messageType
var _ protoreflect.MessageType = fastReflection_ClassSequence_messageType{}

type fastReflection_ClassSequence_messageType struct{}

func (x fastReflection_ClassSequence_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ClassSequence)(nil)
}
func (x fastReflection_ClassSequence_messageType) New() protoreflect.Message {
	return new(fastReflection_ClassSequence)
}
func (x fastReflection_ClassSequence_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassSequence
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ClassSequence) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassSequence
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ClassSequence) Type() protoreflect.MessageType {
	return _fastReflection_ClassSequence_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ClassSequence) New() protoreflect.Message {
	return new(fastReflection_ClassSequence)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ClassSequence) Interface() protoreflect.ProtoMessage {
	return (*ClassSequence)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ClassSequence) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.CreditType != "" {
		value := protoreflect.ValueOfString(x.CreditType)
		if !f(fd_ClassSequence_credit_type, value) {
			return
		}
	}
	if x.NextClassId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.NextClassId)
		if !f(fd_ClassSequence_next_class_id, value) {
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
func (x *fastReflection_ClassSequence) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassSequence.credit_type":
		return x.CreditType != ""
	case "regen.ecocredit.v1beta1.ClassSequence.next_class_id":
		return x.NextClassId != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassSequence does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassSequence) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassSequence.credit_type":
		x.CreditType = ""
	case "regen.ecocredit.v1beta1.ClassSequence.next_class_id":
		x.NextClassId = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassSequence does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ClassSequence) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.ClassSequence.credit_type":
		value := x.CreditType
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.ClassSequence.next_class_id":
		value := x.NextClassId
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassSequence does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ClassSequence) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassSequence.credit_type":
		x.CreditType = value.Interface().(string)
	case "regen.ecocredit.v1beta1.ClassSequence.next_class_id":
		x.NextClassId = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassSequence does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ClassSequence) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassSequence.credit_type":
		panic(fmt.Errorf("field credit_type of message regen.ecocredit.v1beta1.ClassSequence is not mutable"))
	case "regen.ecocredit.v1beta1.ClassSequence.next_class_id":
		panic(fmt.Errorf("field next_class_id of message regen.ecocredit.v1beta1.ClassSequence is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassSequence does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ClassSequence) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ClassSequence.credit_type":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.ClassSequence.next_class_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ClassSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ClassSequence does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ClassSequence) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.ClassSequence", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ClassSequence) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassSequence) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ClassSequence) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ClassSequence) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ClassSequence)
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
		l = len(x.CreditType)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.NextClassId != 0 {
			n += 1 + runtime.Sov(uint64(x.NextClassId))
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
		x := input.Message.Interface().(*ClassSequence)
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
		if x.NextClassId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.NextClassId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.CreditType) > 0 {
			i -= len(x.CreditType)
			copy(dAtA[i:], x.CreditType)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CreditType)))
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
		x := input.Message.Interface().(*ClassSequence)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassSequence: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassSequence: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreditType", wireType)
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
				x.CreditType = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NextClassId", wireType)
				}
				x.NextClassId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.NextClassId |= uint64(b&0x7F) << shift
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

var (
	md_ProjectSequence                 protoreflect.MessageDescriptor
	fd_ProjectSequence_class_id        protoreflect.FieldDescriptor
	fd_ProjectSequence_next_project_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_state_proto_init()
	md_ProjectSequence = File_regen_ecocredit_v1beta1_state_proto.Messages().ByName("ProjectSequence")
	fd_ProjectSequence_class_id = md_ProjectSequence.Fields().ByName("class_id")
	fd_ProjectSequence_next_project_id = md_ProjectSequence.Fields().ByName("next_project_id")
}

var _ protoreflect.Message = (*fastReflection_ProjectSequence)(nil)

type fastReflection_ProjectSequence ProjectSequence

func (x *ProjectSequence) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ProjectSequence)(x)
}

func (x *ProjectSequence) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ProjectSequence_messageType fastReflection_ProjectSequence_messageType
var _ protoreflect.MessageType = fastReflection_ProjectSequence_messageType{}

type fastReflection_ProjectSequence_messageType struct{}

func (x fastReflection_ProjectSequence_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ProjectSequence)(nil)
}
func (x fastReflection_ProjectSequence_messageType) New() protoreflect.Message {
	return new(fastReflection_ProjectSequence)
}
func (x fastReflection_ProjectSequence_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ProjectSequence
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ProjectSequence) Descriptor() protoreflect.MessageDescriptor {
	return md_ProjectSequence
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ProjectSequence) Type() protoreflect.MessageType {
	return _fastReflection_ProjectSequence_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ProjectSequence) New() protoreflect.Message {
	return new(fastReflection_ProjectSequence)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ProjectSequence) Interface() protoreflect.ProtoMessage {
	return (*ProjectSequence)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ProjectSequence) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ClassId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ClassId)
		if !f(fd_ProjectSequence_class_id, value) {
			return
		}
	}
	if x.NextProjectId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.NextProjectId)
		if !f(fd_ProjectSequence_next_project_id, value) {
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
func (x *fastReflection_ProjectSequence) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ProjectSequence.class_id":
		return x.ClassId != uint64(0)
	case "regen.ecocredit.v1beta1.ProjectSequence.next_project_id":
		return x.NextProjectId != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectSequence does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectSequence) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ProjectSequence.class_id":
		x.ClassId = uint64(0)
	case "regen.ecocredit.v1beta1.ProjectSequence.next_project_id":
		x.NextProjectId = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectSequence does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ProjectSequence) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.ProjectSequence.class_id":
		value := x.ClassId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1beta1.ProjectSequence.next_project_id":
		value := x.NextProjectId
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectSequence does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ProjectSequence) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ProjectSequence.class_id":
		x.ClassId = value.Uint()
	case "regen.ecocredit.v1beta1.ProjectSequence.next_project_id":
		x.NextProjectId = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectSequence does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ProjectSequence) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ProjectSequence.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1beta1.ProjectSequence is not mutable"))
	case "regen.ecocredit.v1beta1.ProjectSequence.next_project_id":
		panic(fmt.Errorf("field next_project_id of message regen.ecocredit.v1beta1.ProjectSequence is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectSequence does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ProjectSequence) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.ProjectSequence.class_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1beta1.ProjectSequence.next_project_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.ProjectSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.ProjectSequence does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ProjectSequence) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.ProjectSequence", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ProjectSequence) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectSequence) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ProjectSequence) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ProjectSequence) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ProjectSequence)
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
		if x.ClassId != 0 {
			n += 1 + runtime.Sov(uint64(x.ClassId))
		}
		if x.NextProjectId != 0 {
			n += 1 + runtime.Sov(uint64(x.NextProjectId))
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
		x := input.Message.Interface().(*ProjectSequence)
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
		if x.NextProjectId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.NextProjectId))
			i--
			dAtA[i] = 0x10
		}
		if x.ClassId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ClassId))
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
		x := input.Message.Interface().(*ProjectSequence)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ProjectSequence: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ProjectSequence: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
				}
				x.ClassId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ClassId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NextProjectId", wireType)
				}
				x.NextProjectId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.NextProjectId |= uint64(b&0x7F) << shift
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

var (
	md_BatchSequence               protoreflect.MessageDescriptor
	fd_BatchSequence_project_id    protoreflect.FieldDescriptor
	fd_BatchSequence_next_batch_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_state_proto_init()
	md_BatchSequence = File_regen_ecocredit_v1beta1_state_proto.Messages().ByName("BatchSequence")
	fd_BatchSequence_project_id = md_BatchSequence.Fields().ByName("project_id")
	fd_BatchSequence_next_batch_id = md_BatchSequence.Fields().ByName("next_batch_id")
}

var _ protoreflect.Message = (*fastReflection_BatchSequence)(nil)

type fastReflection_BatchSequence BatchSequence

func (x *BatchSequence) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BatchSequence)(x)
}

func (x *BatchSequence) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BatchSequence_messageType fastReflection_BatchSequence_messageType
var _ protoreflect.MessageType = fastReflection_BatchSequence_messageType{}

type fastReflection_BatchSequence_messageType struct{}

func (x fastReflection_BatchSequence_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BatchSequence)(nil)
}
func (x fastReflection_BatchSequence_messageType) New() protoreflect.Message {
	return new(fastReflection_BatchSequence)
}
func (x fastReflection_BatchSequence_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchSequence
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BatchSequence) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchSequence
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BatchSequence) Type() protoreflect.MessageType {
	return _fastReflection_BatchSequence_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BatchSequence) New() protoreflect.Message {
	return new(fastReflection_BatchSequence)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BatchSequence) Interface() protoreflect.ProtoMessage {
	return (*BatchSequence)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BatchSequence) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProjectId != "" {
		value := protoreflect.ValueOfString(x.ProjectId)
		if !f(fd_BatchSequence_project_id, value) {
			return
		}
	}
	if x.NextBatchId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.NextBatchId)
		if !f(fd_BatchSequence_next_batch_id, value) {
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
func (x *fastReflection_BatchSequence) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchSequence.project_id":
		return x.ProjectId != ""
	case "regen.ecocredit.v1beta1.BatchSequence.next_batch_id":
		return x.NextBatchId != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSequence does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchSequence) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchSequence.project_id":
		x.ProjectId = ""
	case "regen.ecocredit.v1beta1.BatchSequence.next_batch_id":
		x.NextBatchId = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSequence does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BatchSequence) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.BatchSequence.project_id":
		value := x.ProjectId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.BatchSequence.next_batch_id":
		value := x.NextBatchId
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSequence does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BatchSequence) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchSequence.project_id":
		x.ProjectId = value.Interface().(string)
	case "regen.ecocredit.v1beta1.BatchSequence.next_batch_id":
		x.NextBatchId = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSequence does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BatchSequence) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchSequence.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.v1beta1.BatchSequence is not mutable"))
	case "regen.ecocredit.v1beta1.BatchSequence.next_batch_id":
		panic(fmt.Errorf("field next_batch_id of message regen.ecocredit.v1beta1.BatchSequence is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSequence does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BatchSequence) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchSequence.project_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.BatchSequence.next_batch_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSequence"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSequence does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BatchSequence) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.BatchSequence", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BatchSequence) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchSequence) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BatchSequence) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BatchSequence) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BatchSequence)
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
		if x.NextBatchId != 0 {
			n += 1 + runtime.Sov(uint64(x.NextBatchId))
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
		x := input.Message.Interface().(*BatchSequence)
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
		if x.NextBatchId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.NextBatchId))
			i--
			dAtA[i] = 0x10
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
		x := input.Message.Interface().(*BatchSequence)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchSequence: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchSequence: illegal tag %d (wire type %d)", fieldNum, wire)
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
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NextBatchId", wireType)
				}
				x.NextBatchId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.NextBatchId |= uint64(b&0x7F) << shift
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

var (
	md_BatchBalance          protoreflect.MessageDescriptor
	fd_BatchBalance_address  protoreflect.FieldDescriptor
	fd_BatchBalance_batch_id protoreflect.FieldDescriptor
	fd_BatchBalance_tradable protoreflect.FieldDescriptor
	fd_BatchBalance_retired  protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_state_proto_init()
	md_BatchBalance = File_regen_ecocredit_v1beta1_state_proto.Messages().ByName("BatchBalance")
	fd_BatchBalance_address = md_BatchBalance.Fields().ByName("address")
	fd_BatchBalance_batch_id = md_BatchBalance.Fields().ByName("batch_id")
	fd_BatchBalance_tradable = md_BatchBalance.Fields().ByName("tradable")
	fd_BatchBalance_retired = md_BatchBalance.Fields().ByName("retired")
}

var _ protoreflect.Message = (*fastReflection_BatchBalance)(nil)

type fastReflection_BatchBalance BatchBalance

func (x *BatchBalance) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BatchBalance)(x)
}

func (x *BatchBalance) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BatchBalance_messageType fastReflection_BatchBalance_messageType
var _ protoreflect.MessageType = fastReflection_BatchBalance_messageType{}

type fastReflection_BatchBalance_messageType struct{}

func (x fastReflection_BatchBalance_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BatchBalance)(nil)
}
func (x fastReflection_BatchBalance_messageType) New() protoreflect.Message {
	return new(fastReflection_BatchBalance)
}
func (x fastReflection_BatchBalance_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchBalance
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BatchBalance) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchBalance
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BatchBalance) Type() protoreflect.MessageType {
	return _fastReflection_BatchBalance_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BatchBalance) New() protoreflect.Message {
	return new(fastReflection_BatchBalance)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BatchBalance) Interface() protoreflect.ProtoMessage {
	return (*BatchBalance)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BatchBalance) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Address) != 0 {
		value := protoreflect.ValueOfBytes(x.Address)
		if !f(fd_BatchBalance_address, value) {
			return
		}
	}
	if x.BatchId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BatchId)
		if !f(fd_BatchBalance_batch_id, value) {
			return
		}
	}
	if x.Tradable != "" {
		value := protoreflect.ValueOfString(x.Tradable)
		if !f(fd_BatchBalance_tradable, value) {
			return
		}
	}
	if x.Retired != "" {
		value := protoreflect.ValueOfString(x.Retired)
		if !f(fd_BatchBalance_retired, value) {
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
func (x *fastReflection_BatchBalance) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchBalance.address":
		return len(x.Address) != 0
	case "regen.ecocredit.v1beta1.BatchBalance.batch_id":
		return x.BatchId != uint64(0)
	case "regen.ecocredit.v1beta1.BatchBalance.tradable":
		return x.Tradable != ""
	case "regen.ecocredit.v1beta1.BatchBalance.retired":
		return x.Retired != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchBalance does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchBalance) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchBalance.address":
		x.Address = nil
	case "regen.ecocredit.v1beta1.BatchBalance.batch_id":
		x.BatchId = uint64(0)
	case "regen.ecocredit.v1beta1.BatchBalance.tradable":
		x.Tradable = ""
	case "regen.ecocredit.v1beta1.BatchBalance.retired":
		x.Retired = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchBalance does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BatchBalance) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.BatchBalance.address":
		value := x.Address
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.v1beta1.BatchBalance.batch_id":
		value := x.BatchId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1beta1.BatchBalance.tradable":
		value := x.Tradable
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.BatchBalance.retired":
		value := x.Retired
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchBalance does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BatchBalance) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchBalance.address":
		x.Address = value.Bytes()
	case "regen.ecocredit.v1beta1.BatchBalance.batch_id":
		x.BatchId = value.Uint()
	case "regen.ecocredit.v1beta1.BatchBalance.tradable":
		x.Tradable = value.Interface().(string)
	case "regen.ecocredit.v1beta1.BatchBalance.retired":
		x.Retired = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchBalance does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BatchBalance) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchBalance.address":
		panic(fmt.Errorf("field address of message regen.ecocredit.v1beta1.BatchBalance is not mutable"))
	case "regen.ecocredit.v1beta1.BatchBalance.batch_id":
		panic(fmt.Errorf("field batch_id of message regen.ecocredit.v1beta1.BatchBalance is not mutable"))
	case "regen.ecocredit.v1beta1.BatchBalance.tradable":
		panic(fmt.Errorf("field tradable of message regen.ecocredit.v1beta1.BatchBalance is not mutable"))
	case "regen.ecocredit.v1beta1.BatchBalance.retired":
		panic(fmt.Errorf("field retired of message regen.ecocredit.v1beta1.BatchBalance is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchBalance does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BatchBalance) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchBalance.address":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.v1beta1.BatchBalance.batch_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1beta1.BatchBalance.tradable":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.BatchBalance.retired":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchBalance does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BatchBalance) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.BatchBalance", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BatchBalance) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchBalance) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BatchBalance) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BatchBalance) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BatchBalance)
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
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.BatchId != 0 {
			n += 1 + runtime.Sov(uint64(x.BatchId))
		}
		l = len(x.Tradable)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Retired)
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
		x := input.Message.Interface().(*BatchBalance)
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
		if len(x.Retired) > 0 {
			i -= len(x.Retired)
			copy(dAtA[i:], x.Retired)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Retired)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Tradable) > 0 {
			i -= len(x.Tradable)
			copy(dAtA[i:], x.Tradable)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Tradable)))
			i--
			dAtA[i] = 0x1a
		}
		if x.BatchId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BatchId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
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
		x := input.Message.Interface().(*BatchBalance)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchBalance: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchBalance: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = append(x.Address[:0], dAtA[iNdEx:postIndex]...)
				if x.Address == nil {
					x.Address = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchId", wireType)
				}
				x.BatchId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BatchId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tradable", wireType)
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
				x.Tradable = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Retired", wireType)
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
				x.Retired = string(dAtA[iNdEx:postIndex])
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
	md_BatchSupply                  protoreflect.MessageDescriptor
	fd_BatchSupply_batch_id         protoreflect.FieldDescriptor
	fd_BatchSupply_tradable_amount  protoreflect.FieldDescriptor
	fd_BatchSupply_retired_amount   protoreflect.FieldDescriptor
	fd_BatchSupply_cancelled_amount protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1beta1_state_proto_init()
	md_BatchSupply = File_regen_ecocredit_v1beta1_state_proto.Messages().ByName("BatchSupply")
	fd_BatchSupply_batch_id = md_BatchSupply.Fields().ByName("batch_id")
	fd_BatchSupply_tradable_amount = md_BatchSupply.Fields().ByName("tradable_amount")
	fd_BatchSupply_retired_amount = md_BatchSupply.Fields().ByName("retired_amount")
	fd_BatchSupply_cancelled_amount = md_BatchSupply.Fields().ByName("cancelled_amount")
}

var _ protoreflect.Message = (*fastReflection_BatchSupply)(nil)

type fastReflection_BatchSupply BatchSupply

func (x *BatchSupply) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BatchSupply)(x)
}

func (x *BatchSupply) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BatchSupply_messageType fastReflection_BatchSupply_messageType
var _ protoreflect.MessageType = fastReflection_BatchSupply_messageType{}

type fastReflection_BatchSupply_messageType struct{}

func (x fastReflection_BatchSupply_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BatchSupply)(nil)
}
func (x fastReflection_BatchSupply_messageType) New() protoreflect.Message {
	return new(fastReflection_BatchSupply)
}
func (x fastReflection_BatchSupply_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchSupply
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BatchSupply) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchSupply
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BatchSupply) Type() protoreflect.MessageType {
	return _fastReflection_BatchSupply_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BatchSupply) New() protoreflect.Message {
	return new(fastReflection_BatchSupply)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BatchSupply) Interface() protoreflect.ProtoMessage {
	return (*BatchSupply)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BatchSupply) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BatchId)
		if !f(fd_BatchSupply_batch_id, value) {
			return
		}
	}
	if x.TradableAmount != "" {
		value := protoreflect.ValueOfString(x.TradableAmount)
		if !f(fd_BatchSupply_tradable_amount, value) {
			return
		}
	}
	if x.RetiredAmount != "" {
		value := protoreflect.ValueOfString(x.RetiredAmount)
		if !f(fd_BatchSupply_retired_amount, value) {
			return
		}
	}
	if x.CancelledAmount != "" {
		value := protoreflect.ValueOfString(x.CancelledAmount)
		if !f(fd_BatchSupply_cancelled_amount, value) {
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
func (x *fastReflection_BatchSupply) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchSupply.batch_id":
		return x.BatchId != uint64(0)
	case "regen.ecocredit.v1beta1.BatchSupply.tradable_amount":
		return x.TradableAmount != ""
	case "regen.ecocredit.v1beta1.BatchSupply.retired_amount":
		return x.RetiredAmount != ""
	case "regen.ecocredit.v1beta1.BatchSupply.cancelled_amount":
		return x.CancelledAmount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSupply"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSupply does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchSupply) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchSupply.batch_id":
		x.BatchId = uint64(0)
	case "regen.ecocredit.v1beta1.BatchSupply.tradable_amount":
		x.TradableAmount = ""
	case "regen.ecocredit.v1beta1.BatchSupply.retired_amount":
		x.RetiredAmount = ""
	case "regen.ecocredit.v1beta1.BatchSupply.cancelled_amount":
		x.CancelledAmount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSupply"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSupply does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BatchSupply) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1beta1.BatchSupply.batch_id":
		value := x.BatchId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1beta1.BatchSupply.tradable_amount":
		value := x.TradableAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.BatchSupply.retired_amount":
		value := x.RetiredAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1beta1.BatchSupply.cancelled_amount":
		value := x.CancelledAmount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSupply"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSupply does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BatchSupply) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchSupply.batch_id":
		x.BatchId = value.Uint()
	case "regen.ecocredit.v1beta1.BatchSupply.tradable_amount":
		x.TradableAmount = value.Interface().(string)
	case "regen.ecocredit.v1beta1.BatchSupply.retired_amount":
		x.RetiredAmount = value.Interface().(string)
	case "regen.ecocredit.v1beta1.BatchSupply.cancelled_amount":
		x.CancelledAmount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSupply"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSupply does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BatchSupply) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchSupply.batch_id":
		panic(fmt.Errorf("field batch_id of message regen.ecocredit.v1beta1.BatchSupply is not mutable"))
	case "regen.ecocredit.v1beta1.BatchSupply.tradable_amount":
		panic(fmt.Errorf("field tradable_amount of message regen.ecocredit.v1beta1.BatchSupply is not mutable"))
	case "regen.ecocredit.v1beta1.BatchSupply.retired_amount":
		panic(fmt.Errorf("field retired_amount of message regen.ecocredit.v1beta1.BatchSupply is not mutable"))
	case "regen.ecocredit.v1beta1.BatchSupply.cancelled_amount":
		panic(fmt.Errorf("field cancelled_amount of message regen.ecocredit.v1beta1.BatchSupply is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSupply"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSupply does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BatchSupply) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1beta1.BatchSupply.batch_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1beta1.BatchSupply.tradable_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.BatchSupply.retired_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1beta1.BatchSupply.cancelled_amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1beta1.BatchSupply"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1beta1.BatchSupply does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BatchSupply) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1beta1.BatchSupply", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BatchSupply) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchSupply) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BatchSupply) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BatchSupply) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BatchSupply)
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
		if x.BatchId != 0 {
			n += 1 + runtime.Sov(uint64(x.BatchId))
		}
		l = len(x.TradableAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RetiredAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.CancelledAmount)
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
		x := input.Message.Interface().(*BatchSupply)
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
		if len(x.CancelledAmount) > 0 {
			i -= len(x.CancelledAmount)
			copy(dAtA[i:], x.CancelledAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CancelledAmount)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.RetiredAmount) > 0 {
			i -= len(x.RetiredAmount)
			copy(dAtA[i:], x.RetiredAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetiredAmount)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.TradableAmount) > 0 {
			i -= len(x.TradableAmount)
			copy(dAtA[i:], x.TradableAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TradableAmount)))
			i--
			dAtA[i] = 0x12
		}
		if x.BatchId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BatchId))
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
		x := input.Message.Interface().(*BatchSupply)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchSupply: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchSupply: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchId", wireType)
				}
				x.BatchId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BatchId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
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
			case 3:
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
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CancelledAmount", wireType)
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
				x.CancelledAmount = string(dAtA[iNdEx:postIndex])
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: regen/ecocredit/v1beta1/state.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CreditType defines the measurement unit/precision of a certain credit type
// (e.g. carbon, biodiversity...)
type CreditType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// abbreviation is a 1-3 character uppercase abbreviation of the CreditType
	// name, used in batch denominations within the CreditType. It must be unique.
	Abbreviation string `protobuf:"bytes,1,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	// the type of credit (e.g. carbon, biodiversity, etc)
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// the measurement unit (e.g. kg, ton, etc)
	Unit string `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	// the decimal precision
	Precision uint32 `protobuf:"varint,4,opt,name=precision,proto3" json:"precision,omitempty"`
}

func (x *CreditType) Reset() {
	*x = CreditType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditType) ProtoMessage() {}

// Deprecated: Use CreditType.ProtoReflect.Descriptor instead.
func (*CreditType) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_state_proto_rawDescGZIP(), []int{0}
}

func (x *CreditType) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *CreditType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreditType) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *CreditType) GetPrecision() uint32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

// ClassInfo represents the high-level on-chain information for a credit class.
type ClassInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the unique ID of credit class.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// abbrev is the unique string name for this credit class formed from its credit type
	// and an auto-generated integer.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// admin is the admin of the credit class.
	Admin string `protobuf:"bytes,3,opt,name=admin,proto3" json:"admin,omitempty"`
	// metadata is any arbitrary metadata to attached to the credit class.
	Metadata []byte `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// credit_type is the abbreviation of the credit type.
	CreditType string `protobuf:"bytes,5,opt,name=credit_type,json=creditType,proto3" json:"credit_type,omitempty"`
}

func (x *ClassInfo) Reset() {
	*x = ClassInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassInfo) ProtoMessage() {}

// Deprecated: Use ClassInfo.ProtoReflect.Descriptor instead.
func (*ClassInfo) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_state_proto_rawDescGZIP(), []int{1}
}

func (x *ClassInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ClassInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClassInfo) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *ClassInfo) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ClassInfo) GetCreditType() string {
	if x != nil {
		return x.CreditType
	}
	return ""
}

// ClassIssuers is a JOIN table for Class Info that stores the credit class issuers
type ClassIssuer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// class_id is the unique ID of credit class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// issuer is the approved issuer of the credit class.
	Issuer string `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
}

func (x *ClassIssuer) Reset() {
	*x = ClassIssuer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassIssuer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassIssuer) ProtoMessage() {}

// Deprecated: Use ClassIssuer.ProtoReflect.Descriptor instead.
func (*ClassIssuer) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_state_proto_rawDescGZIP(), []int{2}
}

func (x *ClassIssuer) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *ClassIssuer) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

// ProjectInfo represents the high-level on-chain information for a project.
type ProjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the unique ID of the project
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// name is an unique name of this project formed from its credit class
	// name and either an auto-generated number or user provided string.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// class_id is the ID of credit class for this project.
	ClassId uint64 `protobuf:"varint,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// project_location is the location of the project.
	// Full documentation can be found in MsgCreateProject.project_location.
	ProjectLocation string `protobuf:"bytes,4,opt,name=project_location,json=projectLocation,proto3" json:"project_location,omitempty"`
	// metadata is any arbitrary metadata attached to the project.
	Metadata []byte `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ProjectInfo) Reset() {
	*x = ProjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInfo) ProtoMessage() {}

// Deprecated: Use ProjectInfo.ProtoReflect.Descriptor instead.
func (*ProjectInfo) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_state_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectInfo) GetClassId() uint64 {
	if x != nil {
		return x.ClassId
	}
	return 0
}

func (x *ProjectInfo) GetProjectLocation() string {
	if x != nil {
		return x.ProjectLocation
	}
	return ""
}

func (x *ProjectInfo) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// BatchInfo represents the high-level on-chain information for a credit batch.
type BatchInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is an auto-incrementing integer to succinctly identify the batch
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// project_id is the unique ID of the project this batch belongs to.
	ProjectId uint64 `protobuf:"varint,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// batch_denom is the unique string identifier of the credit batch formed
	// from the project name, batch sequence number and dates.
	BatchDenom string `protobuf:"bytes,3,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// metadata is any arbitrary metadata attached to the credit batch.
	Metadata []byte `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// start_date is the beginning of the period during which this credit batch
	// was quantified and verified.
	StartDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// end_date is the end of the period during which this credit batch was
	// quantified and verified.
	EndDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *BatchInfo) Reset() {
	*x = BatchInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchInfo) ProtoMessage() {}

// Deprecated: Use BatchInfo.ProtoReflect.Descriptor instead.
func (*BatchInfo) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_state_proto_rawDescGZIP(), []int{4}
}

func (x *BatchInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BatchInfo) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *BatchInfo) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *BatchInfo) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *BatchInfo) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *BatchInfo) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

// ClassSequence is a sequence number for creating credit class identifiers for each
// credit type.
type ClassSequence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// credit_type is the credit type abbreviation
	CreditType string `protobuf:"bytes,1,opt,name=credit_type,json=creditType,proto3" json:"credit_type,omitempty"`
	// next_class_id is the next class ID for this credit type
	NextClassId uint64 `protobuf:"varint,2,opt,name=next_class_id,json=nextClassId,proto3" json:"next_class_id,omitempty"`
}

func (x *ClassSequence) Reset() {
	*x = ClassSequence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassSequence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassSequence) ProtoMessage() {}

// Deprecated: Use ClassSequence.ProtoReflect.Descriptor instead.
func (*ClassSequence) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_state_proto_rawDescGZIP(), []int{5}
}

func (x *ClassSequence) GetCreditType() string {
	if x != nil {
		return x.CreditType
	}
	return ""
}

func (x *ClassSequence) GetNextClassId() uint64 {
	if x != nil {
		return x.NextClassId
	}
	return 0
}

// ProjectSequence stores and increments the sequence number for projects
// within a given credit class.
type ProjectSequence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// class_id is the id of the credit class
	ClassId uint64 `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// next_project_id is the sequence number for the project
	NextProjectId uint64 `protobuf:"varint,2,opt,name=next_project_id,json=nextProjectId,proto3" json:"next_project_id,omitempty"`
}

func (x *ProjectSequence) Reset() {
	*x = ProjectSequence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectSequence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectSequence) ProtoMessage() {}

// Deprecated: Use ProjectSequence.ProtoReflect.Descriptor instead.
func (*ProjectSequence) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_state_proto_rawDescGZIP(), []int{6}
}

func (x *ProjectSequence) GetClassId() uint64 {
	if x != nil {
		return x.ClassId
	}
	return 0
}

func (x *ProjectSequence) GetNextProjectId() uint64 {
	if x != nil {
		return x.NextProjectId
	}
	return 0
}

// BatchSequence tracks the sequence number for batches within a project
type BatchSequence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project_id is the id of the project for a batch
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// next_batch_id is a sequence number incrementing on each issued batch
	NextBatchId uint64 `protobuf:"varint,2,opt,name=next_batch_id,json=nextBatchId,proto3" json:"next_batch_id,omitempty"`
}

func (x *BatchSequence) Reset() {
	*x = BatchSequence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchSequence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchSequence) ProtoMessage() {}

// Deprecated: Use BatchSequence.ProtoReflect.Descriptor instead.
func (*BatchSequence) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_state_proto_rawDescGZIP(), []int{7}
}

func (x *BatchSequence) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *BatchSequence) GetNextBatchId() uint64 {
	if x != nil {
		return x.NextBatchId
	}
	return 0
}

// BatchBalance stores each users credit balance
type BatchBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the address of the credit holder
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// batch_id is the id of the credit batch
	BatchId uint64 `protobuf:"varint,2,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	// tradable is the tradable amount of credits
	Tradable string `protobuf:"bytes,3,opt,name=tradable,proto3" json:"tradable,omitempty"`
	// retired is the retired amount of credits
	Retired string `protobuf:"bytes,4,opt,name=retired,proto3" json:"retired,omitempty"`
}

func (x *BatchBalance) Reset() {
	*x = BatchBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchBalance) ProtoMessage() {}

// Deprecated: Use BatchBalance.ProtoReflect.Descriptor instead.
func (*BatchBalance) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_state_proto_rawDescGZIP(), []int{8}
}

func (x *BatchBalance) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *BatchBalance) GetBatchId() uint64 {
	if x != nil {
		return x.BatchId
	}
	return 0
}

func (x *BatchBalance) GetTradable() string {
	if x != nil {
		return x.Tradable
	}
	return ""
}

func (x *BatchBalance) GetRetired() string {
	if x != nil {
		return x.Retired
	}
	return ""
}

// BatchSupply tracks the supply of a credit batch
type BatchSupply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batch_id is the id of the batch
	BatchId uint64 `protobuf:"varint,1,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	// tradable_amount is the total number of tradable credits in the credit batch.
	// Some of the issued credits may be cancelled and will be removed from
	// tradable_amount and tracked in amount_cancelled. tradable_amount + retired_amount +
	// amount_cancelled will always sum to the original credit issuance amount.
	TradableAmount string `protobuf:"bytes,2,opt,name=tradable_amount,json=tradableAmount,proto3" json:"tradable_amount,omitempty"`
	// retired_amount is the total amount of credits that have been retired from the credit batch.
	RetiredAmount string `protobuf:"bytes,3,opt,name=retired_amount,json=retiredAmount,proto3" json:"retired_amount,omitempty"`
	// cancelled_amount is the number of credits in the batch that have been
	// cancelled, effectively undoing the issuance. The sum of total_amount and
	// amount_cancelled will always equal the original credit issuance amount.
	CancelledAmount string `protobuf:"bytes,4,opt,name=cancelled_amount,json=cancelledAmount,proto3" json:"cancelled_amount,omitempty"`
}

func (x *BatchSupply) Reset() {
	*x = BatchSupply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1beta1_state_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchSupply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchSupply) ProtoMessage() {}

// Deprecated: Use BatchSupply.ProtoReflect.Descriptor instead.
func (*BatchSupply) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1beta1_state_proto_rawDescGZIP(), []int{9}
}

func (x *BatchSupply) GetBatchId() uint64 {
	if x != nil {
		return x.BatchId
	}
	return 0
}

func (x *BatchSupply) GetTradableAmount() string {
	if x != nil {
		return x.TradableAmount
	}
	return ""
}

func (x *BatchSupply) GetRetiredAmount() string {
	if x != nil {
		return x.RetiredAmount
	}
	return ""
}

func (x *BatchSupply) GetCancelledAmount() string {
	if x != nil {
		return x.CancelledAmount
	}
	return ""
}

var File_regen_ecocredit_v1beta1_state_proto protoreflect.FileDescriptor

var file_regen_ecocredit_v1beta1_state_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1d,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x24, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x1e, 0x0a,
	0x0e, 0x0a, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x18, 0x01, 0x18, 0x01, 0x22, 0xbc, 0x01,
	0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x3a, 0x38, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x32, 0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x18, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x10, 0x03, 0x18, 0x02, 0x22, 0x5d, 0x0a, 0x0b,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x3a, 0x1b,
	0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x15, 0x0a, 0x11, 0x0a, 0x0f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x2c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x22, 0xba, 0x01, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x3a, 0x25, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x1f, 0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x2c, 0x6e, 0x61,
	0x6d, 0x65, 0x10, 0x01, 0x18, 0x01, 0x18, 0x04, 0x22, 0xae, 0x02, 0x0a, 0x09, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x3a, 0x43, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x3d, 0x0a, 0x06, 0x0a, 0x02,
	0x69, 0x64, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x10, 0x01, 0x18, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x10, 0x03, 0x18, 0x05, 0x22, 0x6d, 0x0a, 0x0d, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x3a,
	0x17, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x11, 0x0a, 0x0d, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x22, 0x6a, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x3a, 0x14,
	0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x0e, 0x0a, 0x0a, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x22, 0x6a, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6e, 0x65, 0x78,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x3a, 0x16, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x10,
	0x0a, 0x0c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x22, 0xad, 0x01, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x72, 0x61, 0x64, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x3a, 0x32, 0xf2, 0x9e,
	0xd3, 0x8e, 0x03, 0x2c, 0x0a, 0x12, 0x0a, 0x10, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2c,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x10, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x69, 0x64, 0x2c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x10, 0x01, 0x18, 0x09,
	0x22, 0xb9, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x74,
	0x72, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x74, 0x69, 0x72, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x14, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x0e, 0x0a, 0x0a,
	0x0a, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x42, 0xfb, 0x01, 0x0a,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x65, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x52, 0x45, 0x58, 0xaa, 0x02, 0x17, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x45, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02,
	0x17, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x23, 0x52, 0x65, 0x67, 0x65, 0x6e,
	0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x19, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_regen_ecocredit_v1beta1_state_proto_rawDescOnce sync.Once
	file_regen_ecocredit_v1beta1_state_proto_rawDescData = file_regen_ecocredit_v1beta1_state_proto_rawDesc
)

func file_regen_ecocredit_v1beta1_state_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_v1beta1_state_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_v1beta1_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_v1beta1_state_proto_rawDescData)
	})
	return file_regen_ecocredit_v1beta1_state_proto_rawDescData
}

var file_regen_ecocredit_v1beta1_state_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_regen_ecocredit_v1beta1_state_proto_goTypes = []interface{}{
	(*CreditType)(nil),            // 0: regen.ecocredit.v1beta1.CreditType
	(*ClassInfo)(nil),             // 1: regen.ecocredit.v1beta1.ClassInfo
	(*ClassIssuer)(nil),           // 2: regen.ecocredit.v1beta1.ClassIssuer
	(*ProjectInfo)(nil),           // 3: regen.ecocredit.v1beta1.ProjectInfo
	(*BatchInfo)(nil),             // 4: regen.ecocredit.v1beta1.BatchInfo
	(*ClassSequence)(nil),         // 5: regen.ecocredit.v1beta1.ClassSequence
	(*ProjectSequence)(nil),       // 6: regen.ecocredit.v1beta1.ProjectSequence
	(*BatchSequence)(nil),         // 7: regen.ecocredit.v1beta1.BatchSequence
	(*BatchBalance)(nil),          // 8: regen.ecocredit.v1beta1.BatchBalance
	(*BatchSupply)(nil),           // 9: regen.ecocredit.v1beta1.BatchSupply
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_regen_ecocredit_v1beta1_state_proto_depIdxs = []int32{
	10, // 0: regen.ecocredit.v1beta1.BatchInfo.start_date:type_name -> google.protobuf.Timestamp
	10, // 1: regen.ecocredit.v1beta1.BatchInfo.end_date:type_name -> google.protobuf.Timestamp
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_v1beta1_state_proto_init() }
func file_regen_ecocredit_v1beta1_state_proto_init() {
	if File_regen_ecocredit_v1beta1_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_v1beta1_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditType); i {
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
		file_regen_ecocredit_v1beta1_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassInfo); i {
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
		file_regen_ecocredit_v1beta1_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassIssuer); i {
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
		file_regen_ecocredit_v1beta1_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInfo); i {
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
		file_regen_ecocredit_v1beta1_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchInfo); i {
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
		file_regen_ecocredit_v1beta1_state_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassSequence); i {
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
		file_regen_ecocredit_v1beta1_state_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectSequence); i {
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
		file_regen_ecocredit_v1beta1_state_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchSequence); i {
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
		file_regen_ecocredit_v1beta1_state_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchBalance); i {
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
		file_regen_ecocredit_v1beta1_state_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchSupply); i {
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
			RawDescriptor: file_regen_ecocredit_v1beta1_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_ecocredit_v1beta1_state_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_v1beta1_state_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_v1beta1_state_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_v1beta1_state_proto = out.File
	file_regen_ecocredit_v1beta1_state_proto_rawDesc = nil
	file_regen_ecocredit_v1beta1_state_proto_goTypes = nil
	file_regen_ecocredit_v1beta1_state_proto_depIdxs = nil
}