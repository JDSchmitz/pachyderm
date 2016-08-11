// Code generated by protoc-gen-go.
// source: server/pfs/db/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	server/pfs/db/persist/persist.proto

It has these top-level messages:
	Clock
	ClockID
	BranchClock
	Repo
	BlockRef
	Diff
	Commit
	ProvenanceCommit
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import _ "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileType int32

const (
	FileType_NONE FileType = 0
	FileType_FILE FileType = 1
	FileType_DIR  FileType = 2
)

var FileType_name = map[int32]string{
	0: "NONE",
	1: "FILE",
	2: "DIR",
}
var FileType_value = map[string]int32{
	"NONE": 0,
	"FILE": 1,
	"DIR":  2,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}
func (FileType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Clock struct {
	// a document either has these two fields
	Branch string `protobuf:"bytes,1,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,2,opt,name=clock" json:"clock,omitempty"`
}

func (m *Clock) Reset()                    { *m = Clock{} }
func (m *Clock) String() string            { return proto.CompactTextString(m) }
func (*Clock) ProtoMessage()               {}
func (*Clock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClockID struct {
	ID     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo   string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Branch string `protobuf:"bytes,3,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,4,opt,name=clock" json:"clock,omitempty"`
}

func (m *ClockID) Reset()                    { *m = ClockID{} }
func (m *ClockID) String() string            { return proto.CompactTextString(m) }
func (*ClockID) ProtoMessage()               {}
func (*ClockID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BranchClock struct {
	Clocks []*Clock `protobuf:"bytes,1,rep,name=clocks" json:"clocks,omitempty"`
}

func (m *BranchClock) Reset()                    { *m = BranchClock{} }
func (m *BranchClock) String() string            { return proto.CompactTextString(m) }
func (*BranchClock) ProtoMessage()               {}
func (*BranchClock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BranchClock) GetClocks() []*Clock {
	if m != nil {
		return m.Clocks
	}
	return nil
}

type Repo struct {
	Name       string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Created    *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	Size       uint64                     `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Provenance []string                   `protobuf:"bytes,4,rep,name=provenance" json:"provenance,omitempty"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Repo) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type BlockRef struct {
	Hash  string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Lower uint64 `protobuf:"varint,2,opt,name=lower" json:"lower,omitempty"`
	Upper uint64 `protobuf:"varint,3,opt,name=upper" json:"upper,omitempty"`
}

func (m *BlockRef) Reset()                    { *m = BlockRef{} }
func (m *BlockRef) String() string            { return proto.CompactTextString(m) }
func (*BlockRef) ProtoMessage()               {}
func (*BlockRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Diff struct {
	ID       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo     string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	CommitID string `protobuf:"bytes,3,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	Path     string `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
	// block_refs and delete cannot both be set
	BlockRefs []*BlockRef                `protobuf:"bytes,5,rep,name=block_refs,json=blockRefs" json:"block_refs,omitempty"`
	Delete    bool                       `protobuf:"varint,6,opt,name=delete" json:"delete,omitempty"`
	Size      uint64                     `protobuf:"varint,7,opt,name=size" json:"size,omitempty"`
	Clocks    []*Clock                   `protobuf:"bytes,8,rep,name=clocks" json:"clocks,omitempty"`
	FileType  FileType                   `protobuf:"varint,9,opt,name=file_type,json=fileType,enum=FileType" json:"file_type,omitempty"`
	Modified  *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=modified" json:"modified,omitempty"`
}

func (m *Diff) Reset()                    { *m = Diff{} }
func (m *Diff) String() string            { return proto.CompactTextString(m) }
func (*Diff) ProtoMessage()               {}
func (*Diff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Diff) GetBlockRefs() []*BlockRef {
	if m != nil {
		return m.BlockRefs
	}
	return nil
}

func (m *Diff) GetClocks() []*Clock {
	if m != nil {
		return m.Clocks
	}
	return nil
}

func (m *Diff) GetModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

type Commit struct {
	ID           string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo         string                     `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	BranchClocks []*BranchClock             `protobuf:"bytes,3,rep,name=branch_clocks,json=branchClocks" json:"branch_clocks,omitempty"`
	Started      *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=started" json:"started,omitempty"`
	Finished     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=finished" json:"finished,omitempty"`
	Cancelled    bool                       `protobuf:"varint,6,opt,name=cancelled" json:"cancelled,omitempty"`
	Provenance   []*ProvenanceCommit        `protobuf:"bytes,7,rep,name=provenance" json:"provenance,omitempty"`
	Size         uint64                     `protobuf:"varint,8,opt,name=size" json:"size,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Commit) GetBranchClocks() []*BranchClock {
	if m != nil {
		return m.BranchClocks
	}
	return nil
}

func (m *Commit) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *Commit) GetFinished() *google_protobuf.Timestamp {
	if m != nil {
		return m.Finished
	}
	return nil
}

func (m *Commit) GetProvenance() []*ProvenanceCommit {
	if m != nil {
		return m.Provenance
	}
	return nil
}

type ProvenanceCommit struct {
	ID   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
}

func (m *ProvenanceCommit) Reset()                    { *m = ProvenanceCommit{} }
func (m *ProvenanceCommit) String() string            { return proto.CompactTextString(m) }
func (*ProvenanceCommit) ProtoMessage()               {}
func (*ProvenanceCommit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Clock)(nil), "Clock")
	proto.RegisterType((*ClockID)(nil), "ClockID")
	proto.RegisterType((*BranchClock)(nil), "BranchClock")
	proto.RegisterType((*Repo)(nil), "Repo")
	proto.RegisterType((*BlockRef)(nil), "BlockRef")
	proto.RegisterType((*Diff)(nil), "Diff")
	proto.RegisterType((*Commit)(nil), "Commit")
	proto.RegisterType((*ProvenanceCommit)(nil), "ProvenanceCommit")
	proto.RegisterEnum("FileType", FileType_name, FileType_value)
}

func init() { proto.RegisterFile("server/pfs/db/persist/persist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x25, 0x89, 0xe3, 0xd8, 0x93, 0x52, 0x85, 0x15, 0x42, 0x51, 0x41, 0x05, 0x19, 0x09, 0x22,
	0x24, 0x1c, 0x51, 0xa0, 0x1f, 0xd0, 0xa6, 0x95, 0x82, 0x50, 0x41, 0xab, 0xde, 0x38, 0x44, 0x76,
	0xb2, 0x6e, 0x56, 0xd8, 0xb1, 0xb5, 0xeb, 0xb6, 0x82, 0x33, 0x3f, 0x04, 0x5f, 0xc8, 0xec, 0x78,
	0x9d, 0xb8, 0x55, 0xa5, 0xe4, 0xe4, 0x99, 0xe7, 0x79, 0x3b, 0x6f, 0xde, 0x0c, 0xbc, 0xd6, 0x42,
	0xdd, 0x08, 0x35, 0x2e, 0x12, 0x3d, 0x5e, 0xc4, 0xe3, 0x42, 0x28, 0x2d, 0x75, 0x59, 0x7f, 0xc3,
	0x42, 0xe5, 0x65, 0x7e, 0xf0, 0xf2, 0x2a, 0xcf, 0xaf, 0x52, 0x31, 0xa6, 0x2c, 0xbe, 0x4e, 0xc6,
	0xa5, 0xcc, 0x84, 0x2e, 0xa3, 0xac, 0xb0, 0x05, 0x87, 0xf7, 0x0b, 0x6e, 0x55, 0x54, 0x98, 0x37,
	0xaa, 0xff, 0xc1, 0x67, 0xe8, 0x9e, 0xa6, 0xf9, 0xfc, 0x27, 0x7b, 0x06, 0x6e, 0xac, 0xa2, 0xd5,
	0x7c, 0x39, 0x6c, 0xbd, 0x6a, 0x8d, 0x7c, 0x6e, 0x33, 0xf6, 0x14, 0xba, 0x73, 0x53, 0x30, 0x6c,
	0x23, 0xec, 0xf0, 0x2a, 0x09, 0x7e, 0x40, 0x8f, 0x68, 0xd3, 0x09, 0xdb, 0x87, 0xb6, 0x5c, 0x58,
	0x12, 0x46, 0x8c, 0x81, 0xa3, 0x44, 0x91, 0x53, 0xbd, 0xcf, 0x29, 0x6e, 0x3c, 0xde, 0x79, 0xf8,
	0x71, 0xa7, 0xf9, 0xf8, 0x7b, 0xe8, 0x9f, 0xd0, 0xff, 0x4a, 0xd9, 0x21, 0xb8, 0x84, 0x6b, 0x6c,
	0xd2, 0x19, 0xf5, 0x8f, 0xdc, 0x90, 0x70, 0x6e, 0xd1, 0xe0, 0x4f, 0x0b, 0x1c, 0x6e, 0xba, 0x60,
	0xe7, 0x55, 0x94, 0x09, 0xab, 0x85, 0x62, 0xf6, 0x09, 0x7a, 0x73, 0x25, 0xa2, 0x52, 0x2c, 0x48,
	0x50, 0xff, 0xe8, 0x20, 0xac, 0x1c, 0x09, 0x6b, 0x47, 0xc2, 0xcb, 0xda, 0x32, 0x5e, 0x97, 0x9a,
	0x97, 0xb4, 0xfc, 0x2d, 0x48, 0xad, 0xc3, 0x29, 0x46, 0x19, 0x80, 0x94, 0x1b, 0xb1, 0x42, 0x65,
	0x02, 0x05, 0x77, 0xb0, 0x47, 0x03, 0x09, 0xbe, 0x80, 0x77, 0x42, 0xba, 0x44, 0x62, 0xf8, 0xcb,
	0x48, 0xd7, 0x56, 0x52, 0x6c, 0x66, 0x4d, 0xf3, 0x5b, 0xa1, 0x6a, 0x23, 0x29, 0x31, 0xe8, 0xb5,
	0xd9, 0x87, 0x6d, 0x55, 0x25, 0xc1, 0xdf, 0x36, 0x38, 0x13, 0x99, 0x24, 0x3b, 0x99, 0xfb, 0x1c,
	0xfc, 0x79, 0x9e, 0x65, 0xb2, 0x9c, 0x61, 0x69, 0xe5, 0xaf, 0x57, 0x01, 0x53, 0x22, 0x14, 0x51,
	0xb9, 0x24, 0x83, 0x91, 0x60, 0x62, 0x36, 0x02, 0x88, 0x8d, 0xd2, 0x99, 0x12, 0x89, 0x1e, 0x76,
	0xc9, 0x54, 0x3f, 0xac, 0xc5, 0x73, 0x3f, 0xb6, 0x91, 0x36, 0x7b, 0x5b, 0x88, 0x54, 0x94, 0x62,
	0xe8, 0x22, 0xdf, 0xe3, 0x36, 0x5b, 0xfb, 0xd3, 0xbb, 0xe3, 0x4f, 0xbd, 0x26, 0xef, 0xa1, 0x35,
	0xb1, 0x37, 0xe0, 0x27, 0x32, 0x15, 0xb3, 0xf2, 0x57, 0x21, 0x86, 0x3e, 0x12, 0xf7, 0xb1, 0xe9,
	0x39, 0x22, 0x97, 0x08, 0x70, 0x2f, 0xb1, 0x11, 0x3b, 0x06, 0x2f, 0xcb, 0x17, 0x32, 0x91, 0xb8,
	0x32, 0xd8, 0xba, 0xb2, 0x75, 0x6d, 0xf0, 0xaf, 0x0d, 0xee, 0x29, 0x8d, 0xbd, 0x93, 0x6b, 0x1f,
	0xe0, 0x71, 0x75, 0x84, 0x33, 0xab, 0xba, 0x43, 0xaa, 0xf7, 0xc2, 0xc6, 0xe9, 0xf1, 0xbd, 0x78,
	0x93, 0x68, 0x73, 0x4b, 0xd8, 0x54, 0x99, 0x5b, 0x72, 0xb6, 0xdf, 0x92, 0x2d, 0x35, 0xf3, 0x24,
	0x72, 0x25, 0xf5, 0x12, 0x69, 0xdd, 0xed, 0xf3, 0xd4, 0xb5, 0xec, 0x05, 0xae, 0xd5, 0x1c, 0x56,
	0x9a, 0x22, 0xb1, 0xb2, 0x7f, 0x03, 0xa0, 0xfc, 0xe6, 0x35, 0xf6, 0x48, 0xfb, 0x93, 0xf0, 0xfb,
	0x1a, 0xaa, 0x9c, 0x68, 0x1e, 0xe8, 0x7a, 0x69, 0xde, 0x66, 0x69, 0xc1, 0x31, 0x0c, 0xee, 0x73,
	0x76, 0x71, 0xef, 0xdd, 0x5b, 0xf0, 0xea, 0xd5, 0x31, 0x0f, 0x9c, 0x8b, 0x6f, 0x17, 0x67, 0x83,
	0x47, 0x26, 0x3a, 0x9f, 0x7e, 0x3d, 0x1b, 0xb4, 0x58, 0x0f, 0x3a, 0x93, 0x29, 0x1f, 0xb4, 0x63,
	0x97, 0x66, 0xfc, 0xf8, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x7e, 0x51, 0xa2, 0xce, 0x04, 0x00,
	0x00,
}