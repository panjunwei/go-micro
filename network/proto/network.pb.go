// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package go_micro_network

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/micro/go-micro/router/proto"
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

// Empty request
type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

// ListResponse is returned by ListNodes and ListNeighbours
type ListResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{1}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// NeighbourhoodRequest is sent to query node neighbourhood
type NeighbourhoodRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NeighbourhoodRequest) Reset()         { *m = NeighbourhoodRequest{} }
func (m *NeighbourhoodRequest) String() string { return proto.CompactTextString(m) }
func (*NeighbourhoodRequest) ProtoMessage()    {}
func (*NeighbourhoodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{2}
}

func (m *NeighbourhoodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NeighbourhoodRequest.Unmarshal(m, b)
}
func (m *NeighbourhoodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NeighbourhoodRequest.Marshal(b, m, deterministic)
}
func (m *NeighbourhoodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NeighbourhoodRequest.Merge(m, src)
}
func (m *NeighbourhoodRequest) XXX_Size() int {
	return xxx_messageInfo_NeighbourhoodRequest.Size(m)
}
func (m *NeighbourhoodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NeighbourhoodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NeighbourhoodRequest proto.InternalMessageInfo

func (m *NeighbourhoodRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// NeighbourhoodResponse contains node neighbourhood hierarchy
type NeighbourhoodResponse struct {
	Neighbourhood        *Neighbour `protobuf:"bytes,1,opt,name=neighbourhood,proto3" json:"neighbourhood,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NeighbourhoodResponse) Reset()         { *m = NeighbourhoodResponse{} }
func (m *NeighbourhoodResponse) String() string { return proto.CompactTextString(m) }
func (*NeighbourhoodResponse) ProtoMessage()    {}
func (*NeighbourhoodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{3}
}

func (m *NeighbourhoodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NeighbourhoodResponse.Unmarshal(m, b)
}
func (m *NeighbourhoodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NeighbourhoodResponse.Marshal(b, m, deterministic)
}
func (m *NeighbourhoodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NeighbourhoodResponse.Merge(m, src)
}
func (m *NeighbourhoodResponse) XXX_Size() int {
	return xxx_messageInfo_NeighbourhoodResponse.Size(m)
}
func (m *NeighbourhoodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NeighbourhoodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NeighbourhoodResponse proto.InternalMessageInfo

func (m *NeighbourhoodResponse) GetNeighbourhood() *Neighbour {
	if m != nil {
		return m.Neighbourhood
	}
	return nil
}

// Node is network node
type Node struct {
	// node ide
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// node address
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{4}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// Connect is sent when the node connects to the network
type Connect struct {
	// network mode
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{5}
}

func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (m *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(m, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Close is sent when the node disconnects from the network
type Close struct {
	// network node
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Close) Reset()         { *m = Close{} }
func (m *Close) String() string { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()    {}
func (*Close) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{6}
}

func (m *Close) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Close.Unmarshal(m, b)
}
func (m *Close) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Close.Marshal(b, m, deterministic)
}
func (m *Close) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Close.Merge(m, src)
}
func (m *Close) XXX_Size() int {
	return xxx_messageInfo_Close.Size(m)
}
func (m *Close) XXX_DiscardUnknown() {
	xxx_messageInfo_Close.DiscardUnknown(m)
}

var xxx_messageInfo_Close proto.InternalMessageInfo

func (m *Close) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Solicit is sent when requesting route advertisement from the network nodes
type Solicit struct {
	// network node
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Solicit) Reset()         { *m = Solicit{} }
func (m *Solicit) String() string { return proto.CompactTextString(m) }
func (*Solicit) ProtoMessage()    {}
func (*Solicit) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{7}
}

func (m *Solicit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Solicit.Unmarshal(m, b)
}
func (m *Solicit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Solicit.Marshal(b, m, deterministic)
}
func (m *Solicit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Solicit.Merge(m, src)
}
func (m *Solicit) XXX_Size() int {
	return xxx_messageInfo_Solicit.Size(m)
}
func (m *Solicit) XXX_DiscardUnknown() {
	xxx_messageInfo_Solicit.DiscardUnknown(m)
}

var xxx_messageInfo_Solicit proto.InternalMessageInfo

func (m *Solicit) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Neighbour is used to nnounce node neighbourhood
type Neighbour struct {
	// network node
	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// neighbours
	Neighbours           []*Node  `protobuf:"bytes,3,rep,name=neighbours,proto3" json:"neighbours,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Neighbour) Reset()         { *m = Neighbour{} }
func (m *Neighbour) String() string { return proto.CompactTextString(m) }
func (*Neighbour) ProtoMessage()    {}
func (*Neighbour) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{8}
}

func (m *Neighbour) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Neighbour.Unmarshal(m, b)
}
func (m *Neighbour) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Neighbour.Marshal(b, m, deterministic)
}
func (m *Neighbour) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Neighbour.Merge(m, src)
}
func (m *Neighbour) XXX_Size() int {
	return xxx_messageInfo_Neighbour.Size(m)
}
func (m *Neighbour) XXX_DiscardUnknown() {
	xxx_messageInfo_Neighbour.DiscardUnknown(m)
}

var xxx_messageInfo_Neighbour proto.InternalMessageInfo

func (m *Neighbour) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Neighbour) GetNeighbours() []*Node {
	if m != nil {
		return m.Neighbours
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "go.micro.network.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.network.ListResponse")
	proto.RegisterType((*NeighbourhoodRequest)(nil), "go.micro.network.NeighbourhoodRequest")
	proto.RegisterType((*NeighbourhoodResponse)(nil), "go.micro.network.NeighbourhoodResponse")
	proto.RegisterType((*Node)(nil), "go.micro.network.Node")
	proto.RegisterType((*Connect)(nil), "go.micro.network.Connect")
	proto.RegisterType((*Close)(nil), "go.micro.network.Close")
	proto.RegisterType((*Solicit)(nil), "go.micro.network.Solicit")
	proto.RegisterType((*Neighbour)(nil), "go.micro.network.Neighbour")
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xfd, 0x28, 0xf0, 0x35, 0x0c, 0x1f, 0x5f, 0xcc, 0x46, 0x4d, 0x53, 0x83, 0x21, 0x7b, 0x40,
	0x62, 0xb4, 0x18, 0x08, 0x9e, 0xbc, 0x18, 0x0e, 0x5e, 0x08, 0x87, 0x7a, 0xf3, 0x66, 0xbb, 0x9b,
	0xb2, 0x11, 0x3a, 0xb8, 0xbb, 0x8d, 0x7f, 0xc0, 0x1f, 0x6e, 0xba, 0x5d, 0xb0, 0x80, 0x60, 0xb8,
	0x75, 0xe6, 0xbd, 0x37, 0x6f, 0xa7, 0xfb, 0x16, 0x5a, 0x29, 0xd7, 0x1f, 0x28, 0xdf, 0x82, 0xa5,
	0x44, 0x8d, 0xe4, 0x24, 0xc1, 0x60, 0x21, 0x62, 0x89, 0x81, 0xed, 0xfb, 0xc3, 0x44, 0xe8, 0x59,
	0x16, 0x05, 0x31, 0x2e, 0xfa, 0x06, 0xe9, 0x27, 0x78, 0x5b, 0x7c, 0x48, 0xcc, 0x34, 0x97, 0x7d,
	0xa3, 0xb4, 0x45, 0x31, 0x86, 0xb6, 0xa0, 0x39, 0x11, 0x4a, 0x87, 0xfc, 0x3d, 0xe3, 0x4a, 0xd3,
	0x07, 0xf8, 0x57, 0x94, 0x6a, 0x89, 0xa9, 0xe2, 0xe4, 0x06, 0xea, 0x29, 0x32, 0xae, 0xbc, 0x4a,
	0xa7, 0xda, 0x6b, 0x0e, 0xce, 0x83, 0x6d, 0xd7, 0x60, 0x8a, 0x8c, 0x87, 0x05, 0x89, 0x76, 0xe1,
	0x74, 0xca, 0x45, 0x32, 0x8b, 0x30, 0x93, 0x33, 0x44, 0x66, 0xa7, 0x92, 0xff, 0xe0, 0x08, 0xe6,
	0x55, 0x3a, 0x95, 0x5e, 0x23, 0x74, 0x04, 0xa3, 0x2f, 0x70, 0xb6, 0xc5, 0xb3, 0x76, 0x8f, 0xf9,
	0x96, 0x25, 0xc0, 0x68, 0x9a, 0x83, 0x8b, 0x1f, 0x6c, 0x57, 0xb4, 0x70, 0x53, 0x41, 0xef, 0xa0,
	0x96, 0x1f, 0x69, 0xdb, 0x93, 0x78, 0xe0, 0xbe, 0x32, 0x26, 0xb9, 0x52, 0x9e, 0x63, 0x9a, 0xab,
	0x92, 0x8e, 0xc0, 0x1d, 0x63, 0x9a, 0xf2, 0x58, 0x93, 0x6b, 0xa8, 0xe5, 0x9b, 0x58, 0xdb, 0x7d,
	0xdb, 0x1a, 0x0e, 0x1d, 0x42, 0x7d, 0x3c, 0x47, 0xc5, 0x8f, 0x12, 0x8d, 0xc0, 0x7d, 0xc6, 0xb9,
	0x88, 0xc5, 0x71, 0x5e, 0x08, 0x8d, 0xf5, 0xc2, 0xc7, 0x08, 0xc9, 0x3d, 0xc0, 0xfa, 0xf7, 0x28,
	0xaf, 0x7a, 0xf0, 0x12, 0x4b, 0xcc, 0xc1, 0xa7, 0x03, 0xee, 0xb4, 0x00, 0xc9, 0x13, 0x80, 0xc9,
	0x44, 0x1e, 0x1b, 0x45, 0xbc, 0x6f, 0xb5, 0x0d, 0x92, 0xbd, 0x65, 0xbf, 0xbd, 0x83, 0x94, 0xa3,
	0x44, 0xff, 0x90, 0x09, 0x34, 0xf2, 0x4e, 0x6e, 0xa6, 0x48, 0x7b, 0xf7, 0x14, 0xa5, 0x20, 0xfa,
	0x97, 0xfb, 0xe0, 0xf5, 0xb4, 0x08, 0x5a, 0x1b, 0x21, 0x22, 0xdd, 0x03, 0x29, 0x29, 0xa5, 0xd1,
	0xbf, 0xfa, 0x95, 0xb7, 0xf2, 0x88, 0xfe, 0x9a, 0x47, 0x32, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x59, 0xcf, 0xab, 0xb5, 0x7c, 0x03, 0x00, 0x00,
}
