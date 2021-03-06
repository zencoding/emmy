// Code generated by protoc-gen-go.
// source: enums.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Valid schema types
type SchemaType int32

const (
	SchemaType_PEDERSEN    SchemaType = 0
	SchemaType_PEDERSEN_EC SchemaType = 1
	SchemaType_SCHNORR     SchemaType = 2
	SchemaType_SCHNORR_EC  SchemaType = 3
	SchemaType_CSPAILLIER  SchemaType = 4
	SchemaType_QR          SchemaType = 5
	SchemaType_QNR         SchemaType = 6
)

var SchemaType_name = map[int32]string{
	0: "PEDERSEN",
	1: "PEDERSEN_EC",
	2: "SCHNORR",
	3: "SCHNORR_EC",
	4: "CSPAILLIER",
	5: "QR",
	6: "QNR",
}
var SchemaType_value = map[string]int32{
	"PEDERSEN":    0,
	"PEDERSEN_EC": 1,
	"SCHNORR":     2,
	"SCHNORR_EC":  3,
	"CSPAILLIER":  4,
	"QR":          5,
	"QNR":         6,
}

func (x SchemaType) String() string {
	return proto.EnumName(SchemaType_name, int32(x))
}
func (SchemaType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// Valid schema variants
type SchemaVariant int32

const (
	SchemaVariant_SIGMA SchemaVariant = 0
	SchemaVariant_ZKP   SchemaVariant = 1
	SchemaVariant_ZKPOK SchemaVariant = 2
)

var SchemaVariant_name = map[int32]string{
	0: "SIGMA",
	1: "ZKP",
	2: "ZKPOK",
}
var SchemaVariant_value = map[string]int32{
	"SIGMA": 0,
	"ZKP":   1,
	"ZKPOK": 2,
}

func (x SchemaVariant) String() string {
	return proto.EnumName(SchemaVariant_name, int32(x))
}
func (SchemaVariant) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterEnum("protobuf.SchemaType", SchemaType_name, SchemaType_value)
	proto.RegisterEnum("protobuf.SchemaVariant", SchemaVariant_name, SchemaVariant_value)
}

func init() { proto.RegisterFile("enums.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcd, 0x2b, 0xcd,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x5a, 0x99,
	0x5c, 0x5c, 0xc1, 0xc9, 0x19, 0xa9, 0xb9, 0x89, 0x21, 0x95, 0x05, 0xa9, 0x42, 0x3c, 0x5c, 0x1c,
	0x01, 0xae, 0x2e, 0xae, 0x41, 0xc1, 0xae, 0x7e, 0x02, 0x0c, 0x42, 0xfc, 0x5c, 0xdc, 0x30, 0x5e,
	0xbc, 0xab, 0xb3, 0x00, 0xa3, 0x10, 0x37, 0x17, 0x7b, 0xb0, 0xb3, 0x87, 0x9f, 0x7f, 0x50, 0x90,
	0x00, 0x93, 0x10, 0x1f, 0x50, 0x27, 0x84, 0x03, 0x92, 0x64, 0x06, 0xf1, 0x9d, 0x83, 0x03, 0x1c,
	0x3d, 0x7d, 0x7c, 0x3c, 0x5d, 0x83, 0x04, 0x58, 0x84, 0xd8, 0xb8, 0x98, 0x02, 0x83, 0x04, 0x58,
	0x85, 0xd8, 0xb9, 0x98, 0x03, 0xfd, 0x82, 0x04, 0xd8, 0xb4, 0xf4, 0xb8, 0x78, 0x21, 0x56, 0x85,
	0x25, 0x16, 0x65, 0x26, 0xe6, 0x95, 0x08, 0x71, 0x72, 0xb1, 0x06, 0x7b, 0xba, 0xfb, 0x3a, 0x02,
	0xad, 0x02, 0x2a, 0x8a, 0xf2, 0x0e, 0x00, 0x5a, 0x01, 0x14, 0x03, 0x32, 0xfc, 0xbd, 0x05, 0x98,
	0x92, 0xd8, 0xc0, 0x8e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x97, 0x15, 0x3c, 0x7b, 0xba,
	0x00, 0x00, 0x00,
}
