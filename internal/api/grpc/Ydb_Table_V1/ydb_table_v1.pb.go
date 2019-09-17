// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_table_v1.proto

package Ydb_Table_V1

import (
	_ "github.com/yandex-cloud/ydb-go-sdk/internal/api/protos/Ydb_Table"
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

func init() { proto.RegisterFile("ydb_table_v1.proto", fileDescriptor_d67561c16f71e10d) }

var fileDescriptor_d67561c16f71e10d = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0x6f, 0x8b, 0xd3, 0x40,
	0x10, 0xc6, 0xf1, 0x8d, 0xe0, 0x5a, 0x51, 0x57, 0x41, 0x3c, 0x3d, 0x15, 0xef, 0x3c, 0x7d, 0x95,
	0x5a, 0xfd, 0x04, 0xd7, 0x56, 0x41, 0x14, 0xff, 0xb4, 0xa7, 0x20, 0x82, 0xc7, 0x26, 0x19, 0xce,
	0xa5, 0x9b, 0xec, 0xba, 0xbb, 0x29, 0x97, 0x4f, 0xec, 0xd7, 0x90, 0x26, 0x9b, 0x64, 0x36, 0xd9,
	0x5c, 0xdf, 0x95, 0x79, 0x9e, 0xf9, 0xcd, 0x93, 0x32, 0xc3, 0x12, 0x5a, 0xa6, 0xf1, 0xb9, 0x65,
	0xb1, 0x80, 0xf3, 0xed, 0x2c, 0x52, 0x5a, 0x5a, 0x49, 0x27, 0x3f, 0xd3, 0x38, 0x3a, 0xdb, 0xd5,
	0xa2, 0x1f, 0xb3, 0x83, 0x57, 0x1b, 0xbe, 0xe1, 0x99, 0x9e, 0xaa, 0x22, 0x16, 0x3c, 0x99, 0x32,
	0xc5, 0xa7, 0x95, 0xcd, 0x4c, 0xdb, 0xd6, 0xba, 0xef, 0xcd, 0xbf, 0x09, 0x99, 0x54, 0x6d, 0x6b,
	0xd0, 0x5b, 0x9e, 0x00, 0x5d, 0x91, 0x5b, 0x0b, 0x0d, 0xcc, 0xc2, 0x1a, 0x8c, 0xe1, 0x32, 0xa7,
	0x4f, 0xa3, 0x0e, 0xed, 0x29, 0x2b, 0xf8, 0x5b, 0x80, 0xb1, 0x07, 0xcf, 0xc6, 0x0d, 0x46, 0xc9,
	0xdc, 0x54, 0xcc, 0x25, 0x08, 0x08, 0x33, 0x3d, 0x25, 0xc4, 0xec, 0x19, 0x1c, 0xf3, 0x3d, 0xb9,
	0xf1, 0x11, 0x40, 0x9d, 0x0a, 0xbe, 0x05, 0xfa, 0x08, 0xd9, 0xdb, 0x6a, 0xc3, 0x7a, 0x1c, 0x16,
	0x1d, 0xe7, 0x13, 0xb9, 0x59, 0x87, 0xae, 0x0c, 0xf4, 0x70, 0xf0, 0x31, 0xd5, 0xef, 0x86, 0xf5,
	0x64, 0x4c, 0xee, 0x52, 0x2d, 0xb5, 0x54, 0x35, 0x0b, 0xa7, 0x6a, 0xab, 0xa1, 0x54, 0x48, 0x74,
	0x9c, 0x0f, 0x84, 0x9c, 0x0a, 0x0b, 0xba, 0x06, 0x61, 0x6f, 0x57, 0x6e, 0x48, 0x87, 0x23, 0x6a,
	0x17, 0x69, 0x21, 0x55, 0x39, 0x8c, 0xd4, 0x56, 0x43, 0x91, 0x90, 0xd8, 0x45, 0x6a, 0x8b, 0x86,
	0x06, 0xbd, 0x26, 0x14, 0x09, 0xab, 0x78, 0x1f, 0x4c, 0xa2, 0x79, 0xec, 0xfe, 0x75, 0x7f, 0x1f,
	0x90, 0x12, 0xde, 0x07, 0xcf, 0xe0, 0x98, 0xbf, 0xc8, 0x9d, 0x77, 0x97, 0x4a, 0x30, 0x9e, 0x2f,
	0x99, 0x65, 0xdf, 0x0a, 0xd0, 0x25, 0x7d, 0x8e, 0xba, 0xfa, 0x62, 0x43, 0x3e, 0xba, 0xd2, 0xd3,
	0xc1, 0xbf, 0x6a, 0x50, 0x4c, 0x43, 0x18, 0xde, 0x17, 0x43, 0xf0, 0xa1, 0x07, 0x27, 0x87, 0xa4,
	0xb0, 0x30, 0x96, 0xdc, 0x17, 0xc3, 0xc9, 0xfb, 0x1e, 0x07, 0x67, 0x84, 0x3a, 0x6d, 0x9d, 0xfc,
	0x81, 0x0c, 0x6a, 0xfc, 0xf1, 0xb0, 0x15, 0xc9, 0xcd, 0x80, 0x17, 0x7b, 0x5c, 0x5d, 0xfe, 0x39,
	0x5c, 0xf0, 0xfc, 0x4c, 0xb3, 0xdc, 0xb0, 0xc4, 0xee, 0x0e, 0x1c, 0xe7, 0xef, 0x8b, 0xa1, 0xfc,
	0x43, 0x8f, 0x83, 0xff, 0x26, 0x77, 0x17, 0x32, 0xcb, 0xb8, 0xc5, 0xf4, 0x23, 0x6f, 0xbd, 0x7a,
	0x6a, 0x83, 0x3f, 0xbe, 0xda, 0xe4, 0xf8, 0x29, 0xb9, 0xb7, 0x92, 0x42, 0xc4, 0x2c, 0xd9, 0xe0,
	0x09, 0xf8, 0xd3, 0x03, 0x7a, 0x33, 0xe3, 0x64, 0x9f, 0xcd, 0x4d, 0xb9, 0x20, 0xf7, 0xbd, 0xad,
	0xfd, 0xa2, 0x76, 0xb2, 0xa1, 0x27, 0x63, 0x6b, 0xed, 0x0c, 0xcd, 0x9c, 0x97, 0x7b, 0x7d, 0x6e,
	0xd0, 0x67, 0x72, 0x7b, 0x6d, 0x35, 0xb0, 0x6c, 0x05, 0x2c, 0x1d, 0x9e, 0x7c, 0x5b, 0x0d, 0x9d,
	0x3c, 0x12, 0x6b, 0xda, 0xeb, 0x6b, 0xbb, 0xa3, 0x9f, 0x17, 0x62, 0xf3, 0x5d, 0x19, 0xd0, 0xd6,
	0x3b, 0xfa, 0xae, 0x1c, 0x3a, 0x7a, 0xac, 0xd6, 0xb0, 0xf9, 0x43, 0xf2, 0x20, 0x91, 0x59, 0x54,
	0xb2, 0x3c, 0x85, 0xcb, 0xa8, 0x4c, 0xe3, 0xa8, 0x7e, 0x87, 0xb6, 0xb3, 0xf8, 0x7a, 0xf5, 0x16,
	0xbd, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x3f, 0xf0, 0xdb, 0xd9, 0x06, 0x00, 0x00,
}

const (
	CreateSession        = "/Ydb.Table.V1.TableService/CreateSession"
	DeleteSession        = "/Ydb.Table.V1.TableService/DeleteSession"
	KeepAlive            = "/Ydb.Table.V1.TableService/KeepAlive"
	CreateTable          = "/Ydb.Table.V1.TableService/CreateTable"
	DropTable            = "/Ydb.Table.V1.TableService/DropTable"
	AlterTable           = "/Ydb.Table.V1.TableService/AlterTable"
	CopyTable            = "/Ydb.Table.V1.TableService/CopyTable"
	CopyTables           = "/Ydb.Table.V1.TableService/CopyTables"
	DescribeTable        = "/Ydb.Table.V1.TableService/DescribeTable"
	ExplainDataQuery     = "/Ydb.Table.V1.TableService/ExplainDataQuery"
	PrepareDataQuery     = "/Ydb.Table.V1.TableService/PrepareDataQuery"
	ExecuteDataQuery     = "/Ydb.Table.V1.TableService/ExecuteDataQuery"
	ExecuteSchemeQuery   = "/Ydb.Table.V1.TableService/ExecuteSchemeQuery"
	BeginTransaction     = "/Ydb.Table.V1.TableService/BeginTransaction"
	CommitTransaction    = "/Ydb.Table.V1.TableService/CommitTransaction"
	RollbackTransaction  = "/Ydb.Table.V1.TableService/RollbackTransaction"
	DescribeTableOptions = "/Ydb.Table.V1.TableService/DescribeTableOptions"
	StreamReadTable      = "/Ydb.Table.V1.TableService/StreamReadTable"
	BulkUpsert           = "/Ydb.Table.V1.TableService/BulkUpsert"
)
