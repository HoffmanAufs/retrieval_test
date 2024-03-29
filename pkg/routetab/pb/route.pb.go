// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: route.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// define FindRouteReq message struct here
type FindRouteReq struct {
	Dest []byte   `protobuf:"bytes,1,opt,name=Dest,proto3" json:"Dest,omitempty"`
	Path [][]byte `protobuf:"bytes,2,rep,name=Path,proto3" json:"Path,omitempty"`
}

func (m *FindRouteReq) Reset()         { *m = FindRouteReq{} }
func (m *FindRouteReq) String() string { return proto.CompactTextString(m) }
func (*FindRouteReq) ProtoMessage()    {}
func (*FindRouteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{0}
}
func (m *FindRouteReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FindRouteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FindRouteReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FindRouteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRouteReq.Merge(m, src)
}
func (m *FindRouteReq) XXX_Size() int {
	return m.Size()
}
func (m *FindRouteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRouteReq.DiscardUnknown(m)
}

var xxx_messageInfo_FindRouteReq proto.InternalMessageInfo

func (m *FindRouteReq) GetDest() []byte {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *FindRouteReq) GetPath() [][]byte {
	if m != nil {
		return m.Path
	}
	return nil
}

// define FindRouteResp message struct here
type FindRouteResp struct {
	Dest       []byte       `protobuf:"bytes,1,opt,name=Dest,proto3" json:"Dest,omitempty"`
	RouteItems []*RouteItem `protobuf:"bytes,2,rep,name=RouteItems,proto3" json:"RouteItems,omitempty"`
}

func (m *FindRouteResp) Reset()         { *m = FindRouteResp{} }
func (m *FindRouteResp) String() string { return proto.CompactTextString(m) }
func (*FindRouteResp) ProtoMessage()    {}
func (*FindRouteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{1}
}
func (m *FindRouteResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FindRouteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FindRouteResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FindRouteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRouteResp.Merge(m, src)
}
func (m *FindRouteResp) XXX_Size() int {
	return m.Size()
}
func (m *FindRouteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRouteResp.DiscardUnknown(m)
}

var xxx_messageInfo_FindRouteResp proto.InternalMessageInfo

func (m *FindRouteResp) GetDest() []byte {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *FindRouteResp) GetRouteItems() []*RouteItem {
	if m != nil {
		return m.RouteItems
	}
	return nil
}

type RouteItem struct {
	CreateTime int64        `protobuf:"varint,1,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Ttl        uint32       `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Neighbor   []byte       `protobuf:"bytes,3,opt,name=neighbor,proto3" json:"neighbor,omitempty"`
	NextHop    []*RouteItem `protobuf:"bytes,4,rep,name=nextHop,proto3" json:"nextHop,omitempty"`
}

func (m *RouteItem) Reset()         { *m = RouteItem{} }
func (m *RouteItem) String() string { return proto.CompactTextString(m) }
func (*RouteItem) ProtoMessage()    {}
func (*RouteItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{2}
}
func (m *RouteItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RouteItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RouteItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RouteItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteItem.Merge(m, src)
}
func (m *RouteItem) XXX_Size() int {
	return m.Size()
}
func (m *RouteItem) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteItem.DiscardUnknown(m)
}

var xxx_messageInfo_RouteItem proto.InternalMessageInfo

func (m *RouteItem) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *RouteItem) GetTtl() uint32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *RouteItem) GetNeighbor() []byte {
	if m != nil {
		return m.Neighbor
	}
	return nil
}

func (m *RouteItem) GetNextHop() []*RouteItem {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func init() {
	proto.RegisterType((*FindRouteReq)(nil), "routetab.FindRouteReq")
	proto.RegisterType((*FindRouteResp)(nil), "routetab.FindRouteResp")
	proto.RegisterType((*RouteItem)(nil), "routetab.RouteItem")
}

func init() { proto.RegisterFile("route.proto", fileDescriptor_0984d49a362b6b9f) }

var fileDescriptor_0984d49a362b6b9f = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xca, 0x2f, 0x2d,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x73, 0x4a, 0x12, 0x93, 0x94, 0xcc,
	0xb8, 0x78, 0xdc, 0x32, 0xf3, 0x52, 0x82, 0x40, 0xfc, 0xa0, 0xd4, 0x42, 0x21, 0x21, 0x2e, 0x16,
	0x97, 0xd4, 0xe2, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x30, 0x1b, 0x24, 0x16, 0x90,
	0x58, 0x92, 0x21, 0xc1, 0xa4, 0xc0, 0x0c, 0x12, 0x03, 0xb1, 0x95, 0x22, 0xb8, 0x78, 0x91, 0xf4,
	0x15, 0x17, 0x60, 0xd5, 0x68, 0xcc, 0xc5, 0x05, 0x56, 0xe0, 0x59, 0x92, 0x9a, 0x5b, 0x0c, 0xd6,
	0xce, 0x6d, 0x24, 0xac, 0x07, 0xb3, 0x5b, 0x0f, 0x2e, 0x17, 0x84, 0xa4, 0x4c, 0xa9, 0x83, 0x91,
	0x8b, 0x13, 0xce, 0x15, 0x92, 0xe3, 0xe2, 0x4a, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x0d, 0xc9, 0xcc,
	0x4d, 0x05, 0x1b, 0xce, 0x1c, 0x84, 0x24, 0x22, 0x24, 0xc0, 0xc5, 0x5c, 0x52, 0x92, 0x23, 0xc1,
	0xa4, 0xc0, 0xa8, 0xc1, 0x1b, 0x04, 0x62, 0x0a, 0x49, 0x71, 0x71, 0xe4, 0xa5, 0x66, 0xa6, 0x67,
	0x24, 0xe5, 0x17, 0x49, 0x30, 0x83, 0x1d, 0x03, 0xe7, 0x0b, 0xe9, 0x72, 0xb1, 0xe7, 0xa5, 0x56,
	0x94, 0x78, 0xe4, 0x17, 0x48, 0xb0, 0xe0, 0x76, 0x0d, 0x4c, 0x8d, 0x93, 0xcc, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x31, 0x15, 0x24, 0x25, 0xb1, 0x81, 0xc3, 0xd2, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xb2, 0xa0, 0x27, 0x5a, 0x01, 0x00, 0x00,
}

func (m *FindRouteReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FindRouteReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FindRouteReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		for iNdEx := len(m.Path) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Path[iNdEx])
			copy(dAtA[i:], m.Path[iNdEx])
			i = encodeVarintRoute(dAtA, i, uint64(len(m.Path[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Dest) > 0 {
		i -= len(m.Dest)
		copy(dAtA[i:], m.Dest)
		i = encodeVarintRoute(dAtA, i, uint64(len(m.Dest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FindRouteResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FindRouteResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FindRouteResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RouteItems) > 0 {
		for iNdEx := len(m.RouteItems) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RouteItems[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRoute(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Dest) > 0 {
		i -= len(m.Dest)
		copy(dAtA[i:], m.Dest)
		i = encodeVarintRoute(dAtA, i, uint64(len(m.Dest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RouteItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RouteItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RouteItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NextHop) > 0 {
		for iNdEx := len(m.NextHop) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NextHop[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRoute(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Neighbor) > 0 {
		i -= len(m.Neighbor)
		copy(dAtA[i:], m.Neighbor)
		i = encodeVarintRoute(dAtA, i, uint64(len(m.Neighbor)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Ttl != 0 {
		i = encodeVarintRoute(dAtA, i, uint64(m.Ttl))
		i--
		dAtA[i] = 0x10
	}
	if m.CreateTime != 0 {
		i = encodeVarintRoute(dAtA, i, uint64(m.CreateTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRoute(dAtA []byte, offset int, v uint64) int {
	offset -= sovRoute(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FindRouteReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Dest)
	if l > 0 {
		n += 1 + l + sovRoute(uint64(l))
	}
	if len(m.Path) > 0 {
		for _, b := range m.Path {
			l = len(b)
			n += 1 + l + sovRoute(uint64(l))
		}
	}
	return n
}

func (m *FindRouteResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Dest)
	if l > 0 {
		n += 1 + l + sovRoute(uint64(l))
	}
	if len(m.RouteItems) > 0 {
		for _, e := range m.RouteItems {
			l = e.Size()
			n += 1 + l + sovRoute(uint64(l))
		}
	}
	return n
}

func (m *RouteItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreateTime != 0 {
		n += 1 + sovRoute(uint64(m.CreateTime))
	}
	if m.Ttl != 0 {
		n += 1 + sovRoute(uint64(m.Ttl))
	}
	l = len(m.Neighbor)
	if l > 0 {
		n += 1 + l + sovRoute(uint64(l))
	}
	if len(m.NextHop) > 0 {
		for _, e := range m.NextHop {
			l = e.Size()
			n += 1 + l + sovRoute(uint64(l))
		}
	}
	return n
}

func sovRoute(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRoute(x uint64) (n int) {
	return sovRoute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FindRouteReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
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
			return fmt.Errorf("proto: FindRouteReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FindRouteReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dest = append(m.Dest[:0], dAtA[iNdEx:postIndex]...)
			if m.Dest == nil {
				m.Dest = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = append(m.Path, make([]byte, postIndex-iNdEx))
			copy(m.Path[len(m.Path)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoute
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FindRouteResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
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
			return fmt.Errorf("proto: FindRouteResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FindRouteResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dest = append(m.Dest[:0], dAtA[iNdEx:postIndex]...)
			if m.Dest == nil {
				m.Dest = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteItems", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RouteItems = append(m.RouteItems, &RouteItem{})
			if err := m.RouteItems[len(m.RouteItems)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoute
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RouteItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
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
			return fmt.Errorf("proto: RouteItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RouteItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			m.CreateTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			m.Ttl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ttl |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Neighbor", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Neighbor = append(m.Neighbor[:0], dAtA[iNdEx:postIndex]...)
			if m.Neighbor == nil {
				m.Neighbor = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextHop", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextHop = append(m.NextHop, &RouteItem{})
			if err := m.NextHop[len(m.NextHop)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoute
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRoute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoute
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRoute
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRoute
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRoute
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRoute        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoute          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRoute = fmt.Errorf("proto: unexpected end of group")
)
