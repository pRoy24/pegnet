// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opr.proto

package oprencoding

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

type ProtoOPRMin struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Winners              []string `protobuf:"bytes,4,rep,name=winners,proto3" json:"winners,omitempty"`
	PEG                  uint64   `protobuf:"varint,5,opt,name=PEG,proto3" json:"PEG,omitempty"`
	PUSD                 uint64   `protobuf:"varint,6,opt,name=pUSD,proto3" json:"pUSD,omitempty"`
	PEUR                 uint64   `protobuf:"varint,7,opt,name=pEUR,proto3" json:"pEUR,omitempty"`
	PJPY                 uint64   `protobuf:"varint,8,opt,name=pJPY,proto3" json:"pJPY,omitempty"`
	PGBP                 uint64   `protobuf:"varint,9,opt,name=pGBP,proto3" json:"pGBP,omitempty"`
	PCAD                 uint64   `protobuf:"varint,10,opt,name=pCAD,proto3" json:"pCAD,omitempty"`
	PCHF                 uint64   `protobuf:"varint,11,opt,name=pCHF,proto3" json:"pCHF,omitempty"`
	PINR                 uint64   `protobuf:"varint,12,opt,name=pINR,proto3" json:"pINR,omitempty"`
	PSGD                 uint64   `protobuf:"varint,13,opt,name=pSGD,proto3" json:"pSGD,omitempty"`
	PCNY                 uint64   `protobuf:"varint,14,opt,name=pCNY,proto3" json:"pCNY,omitempty"`
	PHKD                 uint64   `protobuf:"varint,15,opt,name=pHKD,proto3" json:"pHKD,omitempty"`
	PKRW                 uint64   `protobuf:"varint,16,opt,name=pKRW,proto3" json:"pKRW,omitempty"`
	PBRL                 uint64   `protobuf:"varint,17,opt,name=pBRL,proto3" json:"pBRL,omitempty"`
	PPHP                 uint64   `protobuf:"varint,18,opt,name=pPHP,proto3" json:"pPHP,omitempty"`
	PMXN                 uint64   `protobuf:"varint,19,opt,name=pMXN,proto3" json:"pMXN,omitempty"`
	PXAU                 uint64   `protobuf:"varint,20,opt,name=pXAU,proto3" json:"pXAU,omitempty"`
	PXAG                 uint64   `protobuf:"varint,21,opt,name=pXAG,proto3" json:"pXAG,omitempty"`
	PXBT                 uint64   `protobuf:"varint,22,opt,name=pXBT,proto3" json:"pXBT,omitempty"`
	PETH                 uint64   `protobuf:"varint,23,opt,name=pETH,proto3" json:"pETH,omitempty"`
	PLTC                 uint64   `protobuf:"varint,24,opt,name=pLTC,proto3" json:"pLTC,omitempty"`
	PRVN                 uint64   `protobuf:"varint,25,opt,name=pRVN,proto3" json:"pRVN,omitempty"`
	PXBC                 uint64   `protobuf:"varint,26,opt,name=pXBC,proto3" json:"pXBC,omitempty"`
	PFCT                 uint64   `protobuf:"varint,27,opt,name=pFCT,proto3" json:"pFCT,omitempty"`
	PBNB                 uint64   `protobuf:"varint,28,opt,name=pBNB,proto3" json:"pBNB,omitempty"`
	PXLM                 uint64   `protobuf:"varint,29,opt,name=pXLM,proto3" json:"pXLM,omitempty"`
	PADA                 uint64   `protobuf:"varint,30,opt,name=pADA,proto3" json:"pADA,omitempty"`
	PXMR                 uint64   `protobuf:"varint,31,opt,name=pXMR,proto3" json:"pXMR,omitempty"`
	PDASH                uint64   `protobuf:"varint,32,opt,name=pDASH,proto3" json:"pDASH,omitempty"`
	PZEC                 uint64   `protobuf:"varint,33,opt,name=pZEC,proto3" json:"pZEC,omitempty"`
	PDCR                 uint64   `protobuf:"varint,34,opt,name=pDCR,proto3" json:"pDCR,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoOPRMin) Reset()         { *m = ProtoOPRMin{} }
func (m *ProtoOPRMin) String() string { return proto.CompactTextString(m) }
func (*ProtoOPRMin) ProtoMessage()    {}
func (*ProtoOPRMin) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c939583bc8284b5, []int{0}
}

func (m *ProtoOPRMin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoOPRMin.Unmarshal(m, b)
}
func (m *ProtoOPRMin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoOPRMin.Marshal(b, m, deterministic)
}
func (m *ProtoOPRMin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoOPRMin.Merge(m, src)
}
func (m *ProtoOPRMin) XXX_Size() int {
	return xxx_messageInfo_ProtoOPRMin.Size(m)
}
func (m *ProtoOPRMin) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoOPRMin.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoOPRMin proto.InternalMessageInfo

func (m *ProtoOPRMin) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ProtoOPRMin) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProtoOPRMin) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ProtoOPRMin) GetWinners() []string {
	if m != nil {
		return m.Winners
	}
	return nil
}

func (m *ProtoOPRMin) GetPEG() uint64 {
	if m != nil {
		return m.PEG
	}
	return 0
}

func (m *ProtoOPRMin) GetPUSD() uint64 {
	if m != nil {
		return m.PUSD
	}
	return 0
}

func (m *ProtoOPRMin) GetPEUR() uint64 {
	if m != nil {
		return m.PEUR
	}
	return 0
}

func (m *ProtoOPRMin) GetPJPY() uint64 {
	if m != nil {
		return m.PJPY
	}
	return 0
}

func (m *ProtoOPRMin) GetPGBP() uint64 {
	if m != nil {
		return m.PGBP
	}
	return 0
}

func (m *ProtoOPRMin) GetPCAD() uint64 {
	if m != nil {
		return m.PCAD
	}
	return 0
}

func (m *ProtoOPRMin) GetPCHF() uint64 {
	if m != nil {
		return m.PCHF
	}
	return 0
}

func (m *ProtoOPRMin) GetPINR() uint64 {
	if m != nil {
		return m.PINR
	}
	return 0
}

func (m *ProtoOPRMin) GetPSGD() uint64 {
	if m != nil {
		return m.PSGD
	}
	return 0
}

func (m *ProtoOPRMin) GetPCNY() uint64 {
	if m != nil {
		return m.PCNY
	}
	return 0
}

func (m *ProtoOPRMin) GetPHKD() uint64 {
	if m != nil {
		return m.PHKD
	}
	return 0
}

func (m *ProtoOPRMin) GetPKRW() uint64 {
	if m != nil {
		return m.PKRW
	}
	return 0
}

func (m *ProtoOPRMin) GetPBRL() uint64 {
	if m != nil {
		return m.PBRL
	}
	return 0
}

func (m *ProtoOPRMin) GetPPHP() uint64 {
	if m != nil {
		return m.PPHP
	}
	return 0
}

func (m *ProtoOPRMin) GetPMXN() uint64 {
	if m != nil {
		return m.PMXN
	}
	return 0
}

func (m *ProtoOPRMin) GetPXAU() uint64 {
	if m != nil {
		return m.PXAU
	}
	return 0
}

func (m *ProtoOPRMin) GetPXAG() uint64 {
	if m != nil {
		return m.PXAG
	}
	return 0
}

func (m *ProtoOPRMin) GetPXBT() uint64 {
	if m != nil {
		return m.PXBT
	}
	return 0
}

func (m *ProtoOPRMin) GetPETH() uint64 {
	if m != nil {
		return m.PETH
	}
	return 0
}

func (m *ProtoOPRMin) GetPLTC() uint64 {
	if m != nil {
		return m.PLTC
	}
	return 0
}

func (m *ProtoOPRMin) GetPRVN() uint64 {
	if m != nil {
		return m.PRVN
	}
	return 0
}

func (m *ProtoOPRMin) GetPXBC() uint64 {
	if m != nil {
		return m.PXBC
	}
	return 0
}

func (m *ProtoOPRMin) GetPFCT() uint64 {
	if m != nil {
		return m.PFCT
	}
	return 0
}

func (m *ProtoOPRMin) GetPBNB() uint64 {
	if m != nil {
		return m.PBNB
	}
	return 0
}

func (m *ProtoOPRMin) GetPXLM() uint64 {
	if m != nil {
		return m.PXLM
	}
	return 0
}

func (m *ProtoOPRMin) GetPADA() uint64 {
	if m != nil {
		return m.PADA
	}
	return 0
}

func (m *ProtoOPRMin) GetPXMR() uint64 {
	if m != nil {
		return m.PXMR
	}
	return 0
}

func (m *ProtoOPRMin) GetPDASH() uint64 {
	if m != nil {
		return m.PDASH
	}
	return 0
}

func (m *ProtoOPRMin) GetPZEC() uint64 {
	if m != nil {
		return m.PZEC
	}
	return 0
}

func (m *ProtoOPRMin) GetPDCR() uint64 {
	if m != nil {
		return m.PDCR
	}
	return 0
}

func init() {
	proto.RegisterType((*ProtoOPRMin)(nil), "oprencoding.ProtoOPRMin")
}

func init() { proto.RegisterFile("opr.proto", fileDescriptor_9c939583bc8284b5) }

var fileDescriptor_9c939583bc8284b5 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xd2, 0xcb, 0x72, 0x9b, 0x30,
	0x14, 0x06, 0xe0, 0xf1, 0x35, 0x35, 0x6e, 0xd3, 0x94, 0xa6, 0xe9, 0xdf, 0x3b, 0xcd, 0xca, 0xab,
	0x6e, 0xfa, 0x04, 0xba, 0x60, 0x94, 0x18, 0x54, 0x8d, 0x0c, 0x09, 0xce, 0xae, 0x2d, 0x9e, 0x84,
	0x0d, 0x30, 0x38, 0x33, 0x7d, 0xbe, 0xbe, 0x59, 0xc6, 0xd2, 0x21, 0x3b, 0x9d, 0x6f, 0xfe, 0x73,
	0x8e, 0xd0, 0x10, 0x2c, 0xda, 0xae, 0xff, 0xd1, 0xf5, 0xed, 0x63, 0x1b, 0x2e, 0xdb, 0xae, 0xdf,
	0x37, 0x7f, 0xdb, 0xaa, 0x6e, 0xee, 0x2f, 0xff, 0xcf, 0x82, 0xa5, 0x39, 0xf2, 0x2f, 0x63, 0xb3,
	0xba, 0x09, 0x11, 0x9c, 0xfc, 0xae, 0xaa, 0x7e, 0x7f, 0x38, 0x60, 0x14, 0x8d, 0x56, 0x0b, 0x3b,
	0x94, 0xe1, 0x69, 0x30, 0xae, 0x2b, 0x8c, 0x1d, 0x8e, 0xeb, 0x2a, 0xbc, 0x08, 0xe6, 0x0f, 0xfb,
	0xfa, 0xfe, 0xe1, 0x11, 0x93, 0x68, 0xb4, 0x9a, 0x59, 0xaa, 0x8e, 0x13, 0xfe, 0xd5, 0x4d, 0xb3,
	0xef, 0x0f, 0x98, 0x46, 0x93, 0xe3, 0x04, 0x2a, 0xc3, 0xb3, 0x60, 0x62, 0xe2, 0x04, 0xb3, 0x68,
	0xb4, 0x9a, 0xda, 0xe3, 0x31, 0x0c, 0x83, 0x69, 0x57, 0x6c, 0x25, 0xe6, 0x8e, 0xdc, 0xd9, 0x59,
	0x5c, 0x58, 0x9c, 0x90, 0xc5, 0x85, 0x75, 0x76, 0x6d, 0x76, 0x78, 0x41, 0x76, 0x6d, 0x76, 0xce,
	0x12, 0x6e, 0xb0, 0x20, 0x4b, 0xb8, 0x71, 0x26, 0x98, 0x44, 0x40, 0x26, 0x98, 0x9f, 0x27, 0xd4,
	0x1a, 0xcb, 0xc1, 0xd4, 0xda, 0xd9, 0x95, 0xb6, 0x78, 0x49, 0x76, 0xa5, 0xfd, 0x8e, 0x6d, 0x22,
	0xf1, 0x8a, 0x6c, 0x9b, 0x50, 0xaf, 0xde, 0xe1, 0x74, 0xe8, 0xd5, 0x7e, 0xaf, 0xda, 0x48, 0xbc,
	0x26, 0x53, 0x1b, 0x9f, 0xdb, 0xd8, 0x5b, 0x9c, 0x91, 0x6d, 0xec, 0xad, 0x33, 0x6e, 0x53, 0xbc,
	0x21, 0xe3, 0x36, 0x75, 0x66, 0x94, 0x41, 0x48, 0x66, 0x94, 0xbf, 0x73, 0x56, 0x6a, 0xbc, 0x25,
	0xcb, 0x4a, 0xed, 0xac, 0x64, 0x05, 0xce, 0xc9, 0x4a, 0x56, 0x90, 0x25, 0x78, 0xf7, 0x6c, 0xfe,
	0xfd, 0x4a, 0x9e, 0xe3, 0x62, 0x30, 0x9e, 0xfb, 0xf7, 0xcb, 0x15, 0xde, 0x0f, 0xef, 0x97, 0x2b,
	0x67, 0x69, 0x2e, 0x00, 0xb2, 0x34, 0x17, 0xce, 0xec, 0x8d, 0xc6, 0x07, 0x32, 0x7b, 0x43, 0x7b,
	0xb9, 0xc0, 0xc7, 0xe7, 0x79, 0x3e, 0xb7, 0x16, 0x39, 0x3e, 0x91, 0xad, 0x85, 0xdf, 0xc1, 0x35,
	0xc7, 0xe7, 0xe1, 0xdb, 0x34, 0xf7, 0xbd, 0x69, 0x86, 0x2f, 0x43, 0x6f, 0x9a, 0x39, 0x63, 0x92,
	0xe1, 0x2b, 0x19, 0x93, 0xcc, 0xe7, 0x32, 0x8b, 0x6f, 0x43, 0x2e, 0xb3, 0xe1, 0x79, 0x30, 0xeb,
	0x24, 0xdb, 0x2a, 0x44, 0x0e, 0x7d, 0xe1, 0x92, 0x77, 0xb1, 0xc0, 0x77, 0x4a, 0xde, 0xc5, 0xfe,
	0x36, 0x52, 0x58, 0x5c, 0x92, 0x49, 0x61, 0xff, 0xcc, 0xdd, 0x7f, 0xfd, 0xf3, 0x29, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0x65, 0x21, 0x3e, 0xe4, 0x02, 0x00, 0x00,
}
