// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lspd.proto

package lspd

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PaymentInformation struct {
	PaymentHash          []byte   `protobuf:"bytes,1,opt,name=payment_hash,json=paymentHash,proto3" json:"payment_hash,omitempty"`
	PaymentSecret        []byte   `protobuf:"bytes,2,opt,name=payment_secret,json=paymentSecret,proto3" json:"payment_secret,omitempty"`
	Destination          []byte   `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	IncomingAmountMsat   int64    `protobuf:"varint,4,opt,name=incoming_amount_msat,json=incomingAmountMsat,proto3" json:"incoming_amount_msat,omitempty"`
	OutgoingAmountMsat   int64    `protobuf:"varint,5,opt,name=outgoing_amount_msat,json=outgoingAmountMsat,proto3" json:"outgoing_amount_msat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentInformation) Reset()         { *m = PaymentInformation{} }
func (m *PaymentInformation) String() string { return proto.CompactTextString(m) }
func (*PaymentInformation) ProtoMessage()    {}
func (*PaymentInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69a0f5a734bca26, []int{0}
}

func (m *PaymentInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentInformation.Unmarshal(m, b)
}
func (m *PaymentInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentInformation.Marshal(b, m, deterministic)
}
func (m *PaymentInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentInformation.Merge(m, src)
}
func (m *PaymentInformation) XXX_Size() int {
	return xxx_messageInfo_PaymentInformation.Size(m)
}
func (m *PaymentInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentInformation.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentInformation proto.InternalMessageInfo

func (m *PaymentInformation) GetPaymentHash() []byte {
	if m != nil {
		return m.PaymentHash
	}
	return nil
}

func (m *PaymentInformation) GetPaymentSecret() []byte {
	if m != nil {
		return m.PaymentSecret
	}
	return nil
}

func (m *PaymentInformation) GetDestination() []byte {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *PaymentInformation) GetIncomingAmountMsat() int64 {
	if m != nil {
		return m.IncomingAmountMsat
	}
	return 0
}

func (m *PaymentInformation) GetOutgoingAmountMsat() int64 {
	if m != nil {
		return m.OutgoingAmountMsat
	}
	return 0
}

func init() {
	proto.RegisterType((*PaymentInformation)(nil), "lspd.PaymentInformation")
}

func init() {
	proto.RegisterFile("lspd.proto", fileDescriptor_c69a0f5a734bca26)
}

var fileDescriptor_c69a0f5a734bca26 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x29, 0x2e, 0x48,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x9e, 0x30, 0x72, 0x09, 0x05,
	0x24, 0x56, 0xe6, 0xa6, 0xe6, 0x95, 0x78, 0xe6, 0xa5, 0xe5, 0x17, 0xe5, 0x26, 0x96, 0x64, 0xe6,
	0xe7, 0x09, 0x29, 0x72, 0xf1, 0x14, 0x40, 0x44, 0xe3, 0x33, 0x12, 0x8b, 0x33, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x78, 0x82, 0xb8, 0xa1, 0x62, 0x1e, 0x89, 0xc5, 0x19, 0x42, 0xaa, 0x5c, 0x7c, 0x30,
	0x25, 0xc5, 0xa9, 0xc9, 0x45, 0xa9, 0x25, 0x12, 0x4c, 0x60, 0x45, 0xbc, 0x50, 0xd1, 0x60, 0xb0,
	0xa0, 0x90, 0x02, 0x17, 0x77, 0x4a, 0x6a, 0x71, 0x49, 0x66, 0x1e, 0xd8, 0x60, 0x09, 0x66, 0x88,
	0x41, 0x48, 0x42, 0x42, 0x06, 0x5c, 0x22, 0x99, 0x79, 0xc9, 0xf9, 0xb9, 0x99, 0x79, 0xe9, 0xf1,
	0x89, 0xb9, 0xf9, 0xa5, 0x79, 0x25, 0xf1, 0xb9, 0xc5, 0x89, 0x25, 0x12, 0x2c, 0x0a, 0x8c, 0x1a,
	0xcc, 0x41, 0x42, 0x30, 0x39, 0x47, 0xb0, 0x94, 0x6f, 0x71, 0x62, 0x09, 0x48, 0x47, 0x7e, 0x69,
	0x49, 0x7a, 0x3e, 0xba, 0x0e, 0x56, 0x88, 0x0e, 0x98, 0x1c, 0x42, 0x47, 0x12, 0x1b, 0xd8, 0xcf,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xa9, 0xb9, 0x9b, 0x01, 0x01, 0x00, 0x00,
}