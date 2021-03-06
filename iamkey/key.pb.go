// Code generated by protoc-gen-go. DO NOT EDIT.
// source: key.proto

package iamkey // import "github.com/fedorpatlin/go-sdk/iamkey"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Key is resource managed by IAM Key Service.
// Can be issued for User or Service Account, but key authorization is supported only for Service Accounts.
// Issued key contains private part that is not saved on server side, and should be saved by client.
type Key struct {
	// ID of the Key resource.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Subject:
	//	*Key_UserAccountId
	//	*Key_ServiceAccountId
	Subject isKey_Subject `protobuf_oneof:"subject"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Description of the Key resource. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// An algorithm used to generate a key pair of the Key resource.
	KeyAlgorithm v1.Key_Algorithm `protobuf:"varint,6,opt,name=key_algorithm,json=keyAlgorithm,proto3,enum=yandex.cloud.iam.v1.Key_Algorithm" json:"key_algorithm,omitempty"`
	// A public key of the Key resource.
	PublicKey string `protobuf:"bytes,7,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// A public key of the Key resource.
	PrivateKey           string   `protobuf:"bytes,8,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_key_fecd1f348d833dbd, []int{0}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (dst *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(dst, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

type isKey_Subject interface {
	isKey_Subject()
}

type Key_UserAccountId struct {
	UserAccountId string `protobuf:"bytes,2,opt,name=user_account_id,json=userAccountId,proto3,oneof"`
}
type Key_ServiceAccountId struct {
	ServiceAccountId string `protobuf:"bytes,3,opt,name=service_account_id,json=serviceAccountId,proto3,oneof"`
}

func (*Key_UserAccountId) isKey_Subject()    {}
func (*Key_ServiceAccountId) isKey_Subject() {}

func (m *Key) GetSubject() isKey_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Key) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Key) GetUserAccountId() string {
	if x, ok := m.GetSubject().(*Key_UserAccountId); ok {
		return x.UserAccountId
	}
	return ""
}

func (m *Key) GetServiceAccountId() string {
	if x, ok := m.GetSubject().(*Key_ServiceAccountId); ok {
		return x.ServiceAccountId
	}
	return ""
}

func (m *Key) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Key) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Key) GetKeyAlgorithm() v1.Key_Algorithm {
	if m != nil {
		return m.KeyAlgorithm
	}
	return v1.Key_ALGORITHM_UNSPECIFIED
}

func (m *Key) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *Key) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Key) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Key_OneofMarshaler, _Key_OneofUnmarshaler, _Key_OneofSizer, []interface{}{
		(*Key_UserAccountId)(nil),
		(*Key_ServiceAccountId)(nil),
	}
}

func _Key_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Key)
	// subject
	switch x := m.Subject.(type) {
	case *Key_UserAccountId:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.UserAccountId)
	case *Key_ServiceAccountId:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ServiceAccountId)
	case nil:
	default:
		return fmt.Errorf("Key.Subject has unexpected type %T", x)
	}
	return nil
}

func _Key_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Key)
	switch tag {
	case 2: // subject.user_account_id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Subject = &Key_UserAccountId{x}
		return true, err
	case 3: // subject.service_account_id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Subject = &Key_ServiceAccountId{x}
		return true, err
	default:
		return false, nil
	}
}

func _Key_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Key)
	// subject
	switch x := m.Subject.(type) {
	case *Key_UserAccountId:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.UserAccountId)))
		n += len(x.UserAccountId)
	case *Key_ServiceAccountId:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ServiceAccountId)))
		n += len(x.ServiceAccountId)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Key)(nil), "yandex.cloud.sdk.v1.Key")
}

func init() { proto.RegisterFile("key.proto", fileDescriptor_key_fecd1f348d833dbd) }

var fileDescriptor_key_fecd1f348d833dbd = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x41, 0x6f, 0xa3, 0x30,
	0x10, 0x85, 0x17, 0xb2, 0x9b, 0x2c, 0xce, 0x26, 0xbb, 0xf2, 0x5e, 0x50, 0xa4, 0x28, 0x28, 0x27,
	0xa4, 0x2a, 0x46, 0x49, 0x4f, 0x3d, 0x26, 0x97, 0xb6, 0xca, 0x0d, 0xf5, 0xd4, 0x0b, 0x32, 0xf6,
	0x94, 0xba, 0x40, 0x8c, 0x8c, 0x41, 0xf5, 0xbf, 0xed, 0x4f, 0xa9, 0x30, 0x24, 0x4d, 0x2f, 0x96,
	0xfc, 0xde, 0x37, 0xf6, 0xcc, 0x1b, 0xe4, 0xe5, 0x60, 0x48, 0xa5, 0xa4, 0x96, 0xf8, 0xbf, 0xa1,
	0x27, 0x0e, 0xef, 0x84, 0x15, 0xb2, 0xe1, 0xa4, 0xe6, 0x39, 0x69, 0xb7, 0x8b, 0x55, 0x26, 0x65,
	0x56, 0x40, 0x64, 0x91, 0xb4, 0x79, 0x89, 0xb4, 0x28, 0xa1, 0xd6, 0xb4, 0xac, 0xfa, 0xaa, 0xc5,
	0xb2, 0xaf, 0x8a, 0x6c, 0x55, 0x24, 0x68, 0x19, 0xb5, 0xdb, 0xe8, 0xf2, 0xe8, 0xfa, 0xc3, 0x45,
	0xa3, 0x23, 0x18, 0x3c, 0x47, 0xae, 0xe0, 0xbe, 0x13, 0x38, 0xa1, 0x17, 0xbb, 0x82, 0xe3, 0x10,
	0xfd, 0x6d, 0x6a, 0x50, 0x09, 0x65, 0x4c, 0x36, 0x27, 0x9d, 0x08, 0xee, 0xbb, 0x9d, 0xf9, 0xf0,
	0x23, 0x9e, 0x75, 0xc6, 0xbe, 0xd7, 0x1f, 0x39, 0x26, 0x08, 0xd7, 0xa0, 0x5a, 0xc1, 0xe0, 0x1a,
	0x1e, 0x0d, 0xf0, 0xbf, 0xc1, 0xfb, 0xe2, 0xef, 0x10, 0x62, 0x0a, 0xa8, 0x06, 0x9e, 0x50, 0xed,
	0xff, 0x0c, 0x9c, 0x70, 0xba, 0x5b, 0x90, 0x7e, 0x0c, 0x72, 0x1e, 0x83, 0x3c, 0x9d, 0xc7, 0x88,
	0xbd, 0x81, 0xde, 0x6b, 0x1c, 0xa0, 0x29, 0x87, 0x9a, 0x29, 0x51, 0x69, 0x21, 0x4f, 0xfe, 0x2f,
	0xdb, 0xed, 0xb5, 0x84, 0xef, 0xd1, 0x2c, 0x07, 0x93, 0xd0, 0x22, 0x93, 0x4a, 0xe8, 0xd7, 0xd2,
	0x1f, 0x07, 0x4e, 0x38, 0xdf, 0xad, 0xc9, 0xb7, 0xec, 0x04, 0x2d, 0x49, 0xbb, 0x25, 0x47, 0x30,
	0x64, 0x7f, 0x26, 0xe3, 0x3f, 0x39, 0x98, 0xcb, 0x0d, 0x2f, 0x11, 0xaa, 0x9a, 0xb4, 0x10, 0x2c,
	0xc9, 0xc1, 0xf8, 0x13, 0xfb, 0x93, 0xd7, 0x2b, 0x5d, 0x5c, 0x2b, 0x34, 0xad, 0x94, 0x68, 0xa9,
	0x06, 0xeb, 0xff, 0xb6, 0x3e, 0x1a, 0xa4, 0x23, 0x98, 0x83, 0x87, 0x26, 0x75, 0x93, 0xbe, 0x01,
	0xd3, 0x87, 0xcd, 0xf3, 0x4d, 0x9a, 0x0e, 0x0d, 0x6c, 0x34, 0xd0, 0x92, 0xa8, 0x66, 0x58, 0x87,
	0x3d, 0x37, 0x99, 0x8c, 0x6a, 0x9e, 0x77, 0xbb, 0xc9, 0xc1, 0xa4, 0x63, 0x9b, 0xc1, 0xed, 0x67,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0xc2, 0x3d, 0x22, 0xfa, 0x01, 0x00, 0x00,
}
