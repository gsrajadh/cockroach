// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/ccl/sqlccl/backup.proto
// DO NOT EDIT!

/*
	Package sqlccl is a generated protocol buffer package.

	It is generated from these files:
		cockroach/pkg/ccl/sqlccl/backup.proto

	It has these top-level messages:
		BackupDescriptor
*/
package sqlccl

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_build "github.com/cockroachdb/cockroach/pkg/build"
import cockroach_roachpb3 "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_roachpb1 "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_sql_sqlbase1 "github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import github_com_cockroachdb_cockroach_pkg_util_uuid "github.com/cockroachdb/cockroach/pkg/util/uuid"
import github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// BackupDescriptor represents a consistent snapshot of ranges.
//
// Each range snapshot includes a path to data that is a diff of the data in
// that key range between a start and end timestamp. The end timestamp of all
// ranges in a backup is the same, but the start may vary (to allow individual
// tables to be backed up on different schedules).
type BackupDescriptor struct {
	StartTime cockroach_util_hlc.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime" json:"start_time"`
	EndTime   cockroach_util_hlc.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime" json:"end_time"`
	// Spans contains the spans requested for backup. The keyranges covered by
	// `files` may be a subset of this if there were ranges with no changes since
	// the last backup.
	Spans         []cockroach_roachpb1.Span                           `protobuf:"bytes,3,rep,name=spans" json:"spans"`
	Files         []BackupDescriptor_File                             `protobuf:"bytes,4,rep,name=files" json:"files"`
	Descriptors   []cockroach_sql_sqlbase1.Descriptor                 `protobuf:"bytes,5,rep,name=descriptors" json:"descriptors"`
	EntryCounts   cockroach_roachpb3.BulkOpSummary                    `protobuf:"bytes,12,opt,name=entry_counts,json=entryCounts" json:"entry_counts"`
	Dir           cockroach_roachpb3.ExportStorage                    `protobuf:"bytes,7,opt,name=dir" json:"dir"`
	FormatVersion uint32                                              `protobuf:"varint,8,opt,name=format_version,json=formatVersion,proto3" json:"format_version,omitempty"`
	ClusterID     github_com_cockroachdb_cockroach_pkg_util_uuid.UUID `protobuf:"bytes,9,opt,name=cluster_id,json=clusterId,proto3,customtype=github.com/cockroachdb/cockroach/pkg/util/uuid.UUID" json:"cluster_id"`
	// node_id and build_info of the gateway node (which writes the descriptor).
	NodeID    github_com_cockroachdb_cockroach_pkg_roachpb.NodeID `protobuf:"varint,10,opt,name=node_id,json=nodeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.NodeID" json:"node_id,omitempty"`
	BuildInfo cockroach_build.Info                                `protobuf:"bytes,11,opt,name=build_info,json=buildInfo" json:"build_info"`
}

func (m *BackupDescriptor) Reset()                    { *m = BackupDescriptor{} }
func (m *BackupDescriptor) String() string            { return proto.CompactTextString(m) }
func (*BackupDescriptor) ProtoMessage()               {}
func (*BackupDescriptor) Descriptor() ([]byte, []int) { return fileDescriptorBackup, []int{0} }

// BackupDescriptor_File represents a file that contains the diff for a key
// range between two timestamps.
type BackupDescriptor_File struct {
	Span     cockroach_roachpb1.Span `protobuf:"bytes,1,opt,name=span" json:"span"`
	Path     string                  `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Sha512   []byte                  `protobuf:"bytes,4,opt,name=sha512,proto3" json:"sha512,omitempty"`
	DataSize uint64                  `protobuf:"varint,5,opt,name=data_size,json=dataSize,proto3" json:"data_size,omitempty"`
}

func (m *BackupDescriptor_File) Reset()                    { *m = BackupDescriptor_File{} }
func (m *BackupDescriptor_File) String() string            { return proto.CompactTextString(m) }
func (*BackupDescriptor_File) ProtoMessage()               {}
func (*BackupDescriptor_File) Descriptor() ([]byte, []int) { return fileDescriptorBackup, []int{0, 0} }

func init() {
	proto.RegisterType((*BackupDescriptor)(nil), "cockroach.ccl.sqlccl.BackupDescriptor")
	proto.RegisterType((*BackupDescriptor_File)(nil), "cockroach.ccl.sqlccl.BackupDescriptor.File")
}
func (m *BackupDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupDescriptor) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintBackup(dAtA, i, uint64(m.StartTime.Size()))
	n1, err := m.StartTime.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintBackup(dAtA, i, uint64(m.EndTime.Size()))
	n2, err := m.EndTime.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Spans) > 0 {
		for _, msg := range m.Spans {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintBackup(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Files) > 0 {
		for _, msg := range m.Files {
			dAtA[i] = 0x22
			i++
			i = encodeVarintBackup(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Descriptors) > 0 {
		for _, msg := range m.Descriptors {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintBackup(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x3a
	i++
	i = encodeVarintBackup(dAtA, i, uint64(m.Dir.Size()))
	n3, err := m.Dir.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if m.FormatVersion != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.FormatVersion))
	}
	dAtA[i] = 0x4a
	i++
	i = encodeVarintBackup(dAtA, i, uint64(m.ClusterID.Size()))
	n4, err := m.ClusterID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.NodeID != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.NodeID))
	}
	dAtA[i] = 0x5a
	i++
	i = encodeVarintBackup(dAtA, i, uint64(m.BuildInfo.Size()))
	n5, err := m.BuildInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x62
	i++
	i = encodeVarintBackup(dAtA, i, uint64(m.EntryCounts.Size()))
	n6, err := m.EntryCounts.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func (m *BackupDescriptor_File) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupDescriptor_File) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintBackup(dAtA, i, uint64(m.Span.Size()))
	n7, err := m.Span.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	if len(m.Path) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBackup(dAtA, i, uint64(len(m.Path)))
		i += copy(dAtA[i:], m.Path)
	}
	if len(m.Sha512) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintBackup(dAtA, i, uint64(len(m.Sha512)))
		i += copy(dAtA[i:], m.Sha512)
	}
	if m.DataSize != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.DataSize))
	}
	return i, nil
}

func encodeFixed64Backup(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Backup(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintBackup(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BackupDescriptor) Size() (n int) {
	var l int
	_ = l
	l = m.StartTime.Size()
	n += 1 + l + sovBackup(uint64(l))
	l = m.EndTime.Size()
	n += 1 + l + sovBackup(uint64(l))
	if len(m.Spans) > 0 {
		for _, e := range m.Spans {
			l = e.Size()
			n += 1 + l + sovBackup(uint64(l))
		}
	}
	if len(m.Files) > 0 {
		for _, e := range m.Files {
			l = e.Size()
			n += 1 + l + sovBackup(uint64(l))
		}
	}
	if len(m.Descriptors) > 0 {
		for _, e := range m.Descriptors {
			l = e.Size()
			n += 1 + l + sovBackup(uint64(l))
		}
	}
	l = m.Dir.Size()
	n += 1 + l + sovBackup(uint64(l))
	if m.FormatVersion != 0 {
		n += 1 + sovBackup(uint64(m.FormatVersion))
	}
	l = m.ClusterID.Size()
	n += 1 + l + sovBackup(uint64(l))
	if m.NodeID != 0 {
		n += 1 + sovBackup(uint64(m.NodeID))
	}
	l = m.BuildInfo.Size()
	n += 1 + l + sovBackup(uint64(l))
	l = m.EntryCounts.Size()
	n += 1 + l + sovBackup(uint64(l))
	return n
}

func (m *BackupDescriptor_File) Size() (n int) {
	var l int
	_ = l
	l = m.Span.Size()
	n += 1 + l + sovBackup(uint64(l))
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovBackup(uint64(l))
	}
	l = len(m.Sha512)
	if l > 0 {
		n += 1 + l + sovBackup(uint64(l))
	}
	if m.DataSize != 0 {
		n += 1 + sovBackup(uint64(m.DataSize))
	}
	return n
}

func sovBackup(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBackup(x uint64) (n int) {
	return sovBackup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BackupDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackupDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StartTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EndTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spans = append(m.Spans, cockroach_roachpb1.Span{})
			if err := m.Spans[len(m.Spans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Files", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Files = append(m.Files, BackupDescriptor_File{})
			if err := m.Files[len(m.Files)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Descriptors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Descriptors = append(m.Descriptors, cockroach_sql_sqlbase1.Descriptor{})
			if err := m.Descriptors[len(m.Descriptors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dir", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Dir.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FormatVersion", wireType)
			}
			m.FormatVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FormatVersion |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClusterID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= (github_com_cockroachdb_cockroach_pkg_roachpb.NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuildInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BuildInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntryCounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EntryCounts.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
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
func (m *BackupDescriptor_File) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: File: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: File: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Span", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Span.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sha512", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sha512 = append(m.Sha512[:0], dAtA[iNdEx:postIndex]...)
			if m.Sha512 == nil {
				m.Sha512 = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataSize", wireType)
			}
			m.DataSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DataSize |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
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
func skipBackup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBackup
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
					return 0, ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBackup
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBackup
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBackup
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBackup(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBackup = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBackup   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/ccl/sqlccl/backup.proto", fileDescriptorBackup) }

var fileDescriptorBackup = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xed, 0x7c, 0x99, 0xfc, 0x39, 0xed, 0xa7, 0xca, 0x2a, 0x30, 0x0a, 0x22, 0x19, 0x90, 0x8a,
	0x22, 0x90, 0x3c, 0x6a, 0x2b, 0x24, 0xc4, 0x02, 0x89, 0x34, 0xfc, 0x4c, 0x17, 0x20, 0x25, 0x94,
	0x45, 0x37, 0x91, 0x63, 0x3b, 0x89, 0x95, 0xc9, 0x78, 0x6a, 0x7b, 0x10, 0xed, 0x13, 0xb0, 0xe4,
	0x05, 0x78, 0x9f, 0x2e, 0x59, 0x22, 0x16, 0x11, 0x84, 0xb7, 0x60, 0x85, 0xec, 0x99, 0xb4, 0x09,
	0x0a, 0xa2, 0xab, 0xdc, 0x5c, 0x9f, 0x73, 0xee, 0xdc, 0x7b, 0xcf, 0x05, 0xbb, 0x44, 0x90, 0x89,
	0x14, 0x98, 0x8c, 0x83, 0x64, 0x32, 0x0a, 0x08, 0x89, 0x02, 0x75, 0x1a, 0x99, 0x9f, 0x01, 0x26,
	0x93, 0x34, 0x41, 0x89, 0x14, 0x5a, 0xc0, 0x9d, 0x4b, 0x18, 0x22, 0x24, 0x42, 0x19, 0xa4, 0xde,
	0x58, 0x25, 0x0f, 0x52, 0x1e, 0xd1, 0x80, 0xc7, 0x43, 0x91, 0xb1, 0xea, 0xcd, 0xd5, 0x77, 0x1b,
	0x25, 0x83, 0x00, 0x27, 0x3c, 0x07, 0xf8, 0xeb, 0x01, 0x14, 0x6b, 0x9c, 0x23, 0x1e, 0xac, 0x22,
	0xd4, 0xa9, 0xfd, 0xbe, 0x01, 0x56, 0x2c, 0x50, 0x5a, 0xa6, 0x44, 0xa7, 0x92, 0xd1, 0x1c, 0x7b,
	0x7f, 0x15, 0x9b, 0x6a, 0x1e, 0x05, 0xe3, 0x88, 0x04, 0x9a, 0x4f, 0x99, 0xd2, 0x78, 0x9a, 0x37,
	0x53, 0xdf, 0x19, 0x89, 0x91, 0xb0, 0x61, 0x60, 0xa2, 0x2c, 0x7b, 0xef, 0x73, 0x19, 0x6c, 0xb7,
	0x6d, 0xcf, 0x1d, 0xa6, 0x88, 0xe4, 0x89, 0x16, 0x12, 0xb6, 0x01, 0x50, 0x1a, 0x4b, 0xdd, 0x37,
	0x1a, 0x9e, 0xe3, 0x3b, 0xad, 0xda, 0xfe, 0x1d, 0x74, 0x35, 0x0c, 0x53, 0x03, 0x8d, 0x23, 0x82,
	0xde, 0x2e, 0x6a, 0xb4, 0xdd, 0x8b, 0x59, 0x73, 0xa3, 0x5b, 0xb5, 0x34, 0x93, 0x85, 0x4f, 0x41,
	0x85, 0xc5, 0x34, 0x53, 0xf8, 0xef, 0xfa, 0x0a, 0x65, 0x16, 0x53, 0xcb, 0x3f, 0x00, 0x45, 0x95,
	0xe0, 0x58, 0x79, 0x05, 0xbf, 0xd0, 0xaa, 0xed, 0xdf, 0x5a, 0x22, 0xe7, 0x03, 0x43, 0xbd, 0x04,
	0xc7, 0x39, 0x2d, 0xc3, 0xc2, 0x97, 0xa0, 0x38, 0xe4, 0x11, 0x53, 0x9e, 0x6b, 0x49, 0x0f, 0xd1,
	0xba, 0x05, 0xa2, 0x3f, 0xfb, 0x45, 0x2f, 0x78, 0xc4, 0x16, 0x42, 0x96, 0x0f, 0x43, 0x50, 0xa3,
	0x97, 0xef, 0xca, 0x2b, 0x5a, 0xb9, 0xbb, 0x4b, 0x72, 0xea, 0xd4, 0xca, 0x99, 0x95, 0xa0, 0x2b,
	0xa5, 0x5c, 0x64, 0x99, 0x0b, 0x1f, 0x83, 0x02, 0xe5, 0xd2, 0x2b, 0xdb, 0x19, 0xf8, 0x6b, 0xda,
	0x78, 0xfe, 0x21, 0x11, 0x52, 0xf7, 0xb4, 0x90, 0x78, 0xb4, 0xf8, 0x0c, 0x43, 0x81, 0xbb, 0xe0,
	0xff, 0xa1, 0x90, 0x53, 0xac, 0xfb, 0xef, 0x99, 0x54, 0x5c, 0xc4, 0x5e, 0xc5, 0x77, 0x5a, 0x5b,
	0xdd, 0xad, 0x2c, 0xfb, 0x2e, 0x4b, 0xc2, 0x11, 0x00, 0x24, 0x4a, 0x95, 0x66, 0xb2, 0xcf, 0xa9,
	0x57, 0xf5, 0x9d, 0xd6, 0x66, 0xfb, 0x95, 0x51, 0xf9, 0x36, 0x6b, 0x1e, 0x8c, 0xb8, 0x1e, 0xa7,
	0x03, 0x44, 0xc4, 0x34, 0xb8, 0xac, 0x4c, 0x07, 0xc1, 0x1a, 0xcf, 0xa4, 0x29, 0xa7, 0xe8, 0xf8,
	0x38, 0xec, 0xcc, 0x67, 0xcd, 0xea, 0x61, 0x26, 0x18, 0x76, 0xba, 0xd5, 0x5c, 0x3b, 0xa4, 0xf0,
	0x04, 0x94, 0x63, 0x41, 0x99, 0xa9, 0x02, 0x7c, 0xa7, 0x55, 0x6c, 0x3f, 0x9b, 0xcf, 0x9a, 0xa5,
	0xd7, 0x82, 0xb2, 0xb0, 0xf3, 0xeb, 0xba, 0xb5, 0x16, 0x5d, 0x67, 0xb4, 0x6e, 0xc9, 0x28, 0x86,
	0x14, 0x3e, 0x01, 0xc0, 0x1e, 0x52, 0xdf, 0x1c, 0x92, 0x57, 0xb3, 0xc3, 0xba, 0xb1, 0x34, 0x2c,
	0xfb, 0x88, 0xc2, 0x78, 0x28, 0x16, 0x56, 0xb3, 0x19, 0x93, 0x80, 0x21, 0xd8, 0x64, 0xb1, 0x96,
	0x67, 0x7d, 0x22, 0xd2, 0x58, 0x2b, 0x6f, 0xf3, 0xaf, 0xa3, 0x6e, 0xa7, 0xd1, 0xe4, 0x4d, 0xd2,
	0x4b, 0xa7, 0x53, 0x2c, 0xcf, 0x16, 0xcb, 0xb2, 0xdc, 0x43, 0x4b, 0xad, 0x7f, 0x74, 0x80, 0x6b,
	0xdc, 0x00, 0xf7, 0x80, 0x6b, 0x2c, 0x95, 0x9b, 0xff, 0x1f, 0xee, 0xb3, 0x50, 0x08, 0x81, 0x9b,
	0x60, 0x3d, 0xb6, 0x6e, 0xaf, 0x76, 0x6d, 0x0c, 0x6f, 0x82, 0x92, 0x1a, 0xe3, 0x47, 0x7b, 0xfb,
	0x9e, 0x6b, 0xf6, 0xd2, 0xcd, 0xff, 0xc1, 0xdb, 0xa0, 0x6a, 0xce, 0xbd, 0xaf, 0xf8, 0x39, 0xf3,
	0x8a, 0xbe, 0xd3, 0x72, 0xbb, 0x15, 0x93, 0xe8, 0xf1, 0x73, 0x76, 0xe4, 0x56, 0x0a, 0xdb, 0xee,
	0x91, 0x5b, 0x29, 0x6d, 0x97, 0xdb, 0xfe, 0xc5, 0x8f, 0xc6, 0xc6, 0xc5, 0xbc, 0xe1, 0x7c, 0x99,
	0x37, 0x9c, 0xaf, 0xf3, 0x86, 0xf3, 0x7d, 0xde, 0x70, 0x3e, 0xfd, 0x6c, 0x6c, 0x9c, 0x94, 0x32,
	0x37, 0x0f, 0x4a, 0xf6, 0x90, 0x0f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xef, 0x12, 0x8f, 0x84,
	0xd4, 0x04, 0x00, 0x00,
}
