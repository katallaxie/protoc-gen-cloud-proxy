// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: example/example.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	api "github.com/katallaxie/protoc-gen-cloud-proxy/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Song ...
type Song struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Artist ...
	Artist string `protobuf:"bytes,1,opt,name=artist,proto3" json:"artist,omitempty"`
	// SongTitle ...
	SongTitle string `protobuf:"bytes,2,opt,name=song_title,json=songTitle,proto3" json:"song_title,omitempty"`
	// AlbumTitle ...
	AlbumTitle string `protobuf:"bytes,3,opt,name=album_title,json=albumTitle,proto3" json:"album_title,omitempty"`
	// Year ...
	Year string `protobuf:"bytes,4,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *Song) Reset() {
	*x = Song{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Song) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Song) ProtoMessage() {}

func (x *Song) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Song.ProtoReflect.Descriptor instead.
func (*Song) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{0}
}

func (x *Song) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *Song) GetSongTitle() string {
	if x != nil {
		return x.SongTitle
	}
	return ""
}

func (x *Song) GetAlbumTitle() string {
	if x != nil {
		return x.AlbumTitle
	}
	return ""
}

func (x *Song) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

// Insert ...
type Insert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Insert) Reset() {
	*x = Insert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insert) ProtoMessage() {}

func (x *Insert) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insert.ProtoReflect.Descriptor instead.
func (*Insert) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{1}
}

// Update ...
type Update struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Update) Reset() {
	*x = Update{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update) ProtoMessage() {}

func (x *Update) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Update.ProtoReflect.Descriptor instead.
func (*Update) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{2}
}

// ListSongs ...
type ListSongs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSongs) Reset() {
	*x = ListSongs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSongs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSongs) ProtoMessage() {}

func (x *ListSongs) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSongs.ProtoReflect.Descriptor instead.
func (*ListSongs) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{3}
}

// ReceiveInserts ...
type ReceiveInserts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceiveInserts) Reset() {
	*x = ReceiveInserts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveInserts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveInserts) ProtoMessage() {}

func (x *ReceiveInserts) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveInserts.ProtoReflect.Descriptor instead.
func (*ReceiveInserts) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{4}
}

// Empty ...
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{5}
}

// Request ...
type Insert_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Song *Song `protobuf:"bytes,1,opt,name=Song,proto3" json:"Song,omitempty"`
}

func (x *Insert_Request) Reset() {
	*x = Insert_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insert_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insert_Request) ProtoMessage() {}

func (x *Insert_Request) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insert_Request.ProtoReflect.Descriptor instead.
func (*Insert_Request) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Insert_Request) GetSong() *Song {
	if x != nil {
		return x.Song
	}
	return nil
}

// Response ...
type Insert_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Insert_Response) Reset() {
	*x = Insert_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insert_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insert_Response) ProtoMessage() {}

func (x *Insert_Response) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insert_Response.ProtoReflect.Descriptor instead.
func (*Insert_Response) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Insert_Response) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// Request ...
type Update_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Song *Song `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
}

func (x *Update_Request) Reset() {
	*x = Update_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Update_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update_Request) ProtoMessage() {}

func (x *Update_Request) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Update_Request.ProtoReflect.Descriptor instead.
func (*Update_Request) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Update_Request) GetSong() *Song {
	if x != nil {
		return x.Song
	}
	return nil
}

// Response ...
type Update_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Update_Response) Reset() {
	*x = Update_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Update_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update_Response) ProtoMessage() {}

func (x *Update_Response) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Update_Response.ProtoReflect.Descriptor instead.
func (*Update_Response) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Update_Response) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// Request ...
type ListSongs_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSongs_Request) Reset() {
	*x = ListSongs_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSongs_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSongs_Request) ProtoMessage() {}

func (x *ListSongs_Request) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSongs_Request.ProtoReflect.Descriptor instead.
func (*ListSongs_Request) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{3, 0}
}

// Response ...
type ListSongs_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Songs []*Song `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"`
}

func (x *ListSongs_Response) Reset() {
	*x = ListSongs_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSongs_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSongs_Response) ProtoMessage() {}

func (x *ListSongs_Response) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSongs_Response.ProtoReflect.Descriptor instead.
func (*ListSongs_Response) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{3, 1}
}

func (x *ListSongs_Response) GetSongs() []*Song {
	if x != nil {
		return x.Songs
	}
	return nil
}

// Request ...
type ReceiveInserts_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceiveInserts_Request) Reset() {
	*x = ReceiveInserts_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveInserts_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveInserts_Request) ProtoMessage() {}

func (x *ReceiveInserts_Request) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveInserts_Request.ProtoReflect.Descriptor instead.
func (*ReceiveInserts_Request) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{4, 0}
}

var File_example_example_proto protoreflect.FileDescriptor

var file_example_example_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x09, 0x73, 0x71, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x04, 0x53,
	0x6f, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0f, 0xf2, 0x94, 0x8c, 0x2f, 0x0a, 0x52, 0x08, 0x12, 0x06, 0x41, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0a,
	0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x12, 0xf2, 0x94, 0x8c, 0x2f, 0x0d, 0x52, 0x0b, 0x12, 0x09, 0x53, 0x6f, 0x6e, 0x67, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x6f, 0x6e, 0x67, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x2c, 0x0a, 0x0b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xf2, 0x94, 0x8c, 0x2f, 0x06, 0x52, 0x04, 0x0a, 0x02, 0x3a,
	0x74, 0x52, 0x0a, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xf2, 0x94, 0x8c,
	0x2f, 0x06, 0x52, 0x04, 0x0a, 0x02, 0x3a, 0x79, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x54,
	0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x1a, 0x2a, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x53, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x04,
	0x53, 0x6f, 0x6e, 0x67, 0x1a, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x2a,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x73, 0x6f, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x6f, 0x6e, 0x67, 0x52, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x1a, 0x1e, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x05, 0x73, 0x6f, 0x6e, 0x67,
	0x73, 0x22, 0x1b, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x73, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xc2, 0x03, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48, 0xf2,
	0x94, 0x8c, 0x2f, 0x43, 0x12, 0x41, 0x0a, 0x36, 0x61, 0x72, 0x6e, 0x3a, 0x61, 0x77, 0x73, 0x3a,
	0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x3a, 0x65, 0x75, 0x2d, 0x77, 0x65, 0x73, 0x74, 0x2d, 0x31,
	0x3a, 0x32, 0x39, 0x31, 0x33, 0x33, 0x39, 0x30, 0x38, 0x38, 0x39, 0x33, 0x35, 0x3a, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x6d, 0x79, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x12, 0x07,
	0x24, 0x4c, 0x41, 0x54, 0x45, 0x53, 0x54, 0x12, 0x88, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4f, 0xf2, 0x94, 0x8c, 0x2f, 0x4a, 0x1a, 0x48, 0x62, 0x46, 0x0a, 0x05, 0x4d, 0x75,
	0x73, 0x69, 0x63, 0x12, 0x07, 0x41, 0x4c, 0x4c, 0x5f, 0x4e, 0x45, 0x57, 0x3a, 0x11, 0x0a, 0x03,
	0x23, 0x41, 0x54, 0x12, 0x0a, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x3a,
	0x0a, 0x0a, 0x02, 0x23, 0x59, 0x12, 0x04, 0x59, 0x65, 0x61, 0x72, 0x4a, 0x15, 0x53, 0x45, 0x54,
	0x20, 0x23, 0x59, 0x20, 0x3d, 0x20, 0x3a, 0x79, 0x2c, 0x20, 0x23, 0x41, 0x54, 0x20, 0x3d, 0x20,
	0x3a, 0x74, 0x12, 0x59, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x61, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x51, 0x53, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x6d, 0x61, 0x7a,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x51, 0x53, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x07, 0xf2, 0x94, 0x8c, 0x2f, 0x02, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4d, 0x0a,
	0x0b, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x61,
	0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x51, 0x53, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11,
	0x2e, 0x61, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x07, 0xf2, 0x94, 0x8c, 0x2f, 0x02, 0x22, 0x00, 0x28, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_example_proto_rawDescOnce sync.Once
	file_example_example_proto_rawDescData = file_example_example_proto_rawDesc
)

func file_example_example_proto_rawDescGZIP() []byte {
	file_example_example_proto_rawDescOnce.Do(func() {
		file_example_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_example_proto_rawDescData)
	})
	return file_example_example_proto_rawDescData
}

var file_example_example_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_example_example_proto_goTypes = []interface{}{
	(*Song)(nil),                        // 0: proto.Song
	(*Insert)(nil),                      // 1: proto.Insert
	(*Update)(nil),                      // 2: proto.Update
	(*ListSongs)(nil),                   // 3: proto.ListSongs
	(*ReceiveInserts)(nil),              // 4: proto.ReceiveInserts
	(*Empty)(nil),                       // 5: proto.Empty
	(*Insert_Request)(nil),              // 6: proto.Insert.Request
	(*Insert_Response)(nil),             // 7: proto.Insert.Response
	(*Update_Request)(nil),              // 8: proto.Update.Request
	(*Update_Response)(nil),             // 9: proto.Update.Response
	(*ListSongs_Request)(nil),           // 10: proto.ListSongs.Request
	(*ListSongs_Response)(nil),          // 11: proto.ListSongs.Response
	(*ReceiveInserts_Request)(nil),      // 12: proto.ReceiveInserts.Request
	(*api.SQS_ReceiveMessageInput)(nil), // 13: amazon.api.SQS.ReceiveMessageInput
	(*api.SQS_SendMessageInput)(nil),    // 14: amazon.api.SQS.SendMessageInput
	(*api.SQS_Message)(nil),             // 15: amazon.api.SQS.Message
	(*api.Empty)(nil),                   // 16: amazon.api.Empty
}
var file_example_example_proto_depIdxs = []int32{
	0,  // 0: proto.Insert.Request.Song:type_name -> proto.Song
	0,  // 1: proto.Update.Request.song:type_name -> proto.Song
	0,  // 2: proto.ListSongs.Response.songs:type_name -> proto.Song
	6,  // 3: proto.Example.Insert:input_type -> proto.Insert.Request
	8,  // 4: proto.Example.Update:input_type -> proto.Update.Request
	13, // 5: proto.Example.ReceiveInserts:input_type -> amazon.api.SQS.ReceiveMessageInput
	14, // 6: proto.Example.SendInserts:input_type -> amazon.api.SQS.SendMessageInput
	7,  // 7: proto.Example.Insert:output_type -> proto.Insert.Response
	9,  // 8: proto.Example.Update:output_type -> proto.Update.Response
	15, // 9: proto.Example.ReceiveInserts:output_type -> amazon.api.SQS.Message
	16, // 10: proto.Example.SendInserts:output_type -> amazon.api.Empty
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_example_example_proto_init() }
func file_example_example_proto_init() {
	if File_example_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Song); i {
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
		file_example_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insert); i {
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
		file_example_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Update); i {
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
		file_example_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSongs); i {
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
		file_example_example_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveInserts); i {
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
		file_example_example_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_example_example_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insert_Request); i {
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
		file_example_example_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insert_Response); i {
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
		file_example_example_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Update_Request); i {
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
		file_example_example_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Update_Response); i {
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
		file_example_example_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSongs_Request); i {
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
		file_example_example_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSongs_Response); i {
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
		file_example_example_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveInserts_Request); i {
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
			RawDescriptor: file_example_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_example_proto_goTypes,
		DependencyIndexes: file_example_example_proto_depIdxs,
		MessageInfos:      file_example_example_proto_msgTypes,
	}.Build()
	File_example_example_proto = out.File
	file_example_example_proto_rawDesc = nil
	file_example_example_proto_goTypes = nil
	file_example_example_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleClient interface {
	// Insert ...
	Insert(ctx context.Context, in *Insert_Request, opts ...grpc.CallOption) (*Insert_Response, error)
	// Update ...
	Update(ctx context.Context, in *Update_Request, opts ...grpc.CallOption) (*Update_Response, error)
	// ReceiveInserts ...
	ReceiveInserts(ctx context.Context, in *api.SQS_ReceiveMessageInput, opts ...grpc.CallOption) (Example_ReceiveInsertsClient, error)
	// SendInserts ...
	SendInserts(ctx context.Context, opts ...grpc.CallOption) (Example_SendInsertsClient, error)
}

type exampleClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleClient(cc grpc.ClientConnInterface) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Insert(ctx context.Context, in *Insert_Request, opts ...grpc.CallOption) (*Insert_Response, error) {
	out := new(Insert_Response)
	err := c.cc.Invoke(ctx, "/proto.Example/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleClient) Update(ctx context.Context, in *Update_Request, opts ...grpc.CallOption) (*Update_Response, error) {
	out := new(Update_Response)
	err := c.cc.Invoke(ctx, "/proto.Example/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleClient) ReceiveInserts(ctx context.Context, in *api.SQS_ReceiveMessageInput, opts ...grpc.CallOption) (Example_ReceiveInsertsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Example_serviceDesc.Streams[0], "/proto.Example/ReceiveInserts", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleReceiveInsertsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Example_ReceiveInsertsClient interface {
	Recv() (*api.SQS_Message, error)
	grpc.ClientStream
}

type exampleReceiveInsertsClient struct {
	grpc.ClientStream
}

func (x *exampleReceiveInsertsClient) Recv() (*api.SQS_Message, error) {
	m := new(api.SQS_Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleClient) SendInserts(ctx context.Context, opts ...grpc.CallOption) (Example_SendInsertsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Example_serviceDesc.Streams[1], "/proto.Example/SendInserts", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleSendInsertsClient{stream}
	return x, nil
}

type Example_SendInsertsClient interface {
	Send(*api.SQS_SendMessageInput) error
	CloseAndRecv() (*api.Empty, error)
	grpc.ClientStream
}

type exampleSendInsertsClient struct {
	grpc.ClientStream
}

func (x *exampleSendInsertsClient) Send(m *api.SQS_SendMessageInput) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exampleSendInsertsClient) CloseAndRecv() (*api.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(api.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExampleServer is the server API for Example service.
type ExampleServer interface {
	// Insert ...
	Insert(context.Context, *Insert_Request) (*Insert_Response, error)
	// Update ...
	Update(context.Context, *Update_Request) (*Update_Response, error)
	// ReceiveInserts ...
	ReceiveInserts(*api.SQS_ReceiveMessageInput, Example_ReceiveInsertsServer) error
	// SendInserts ...
	SendInserts(Example_SendInsertsServer) error
}

// UnimplementedExampleServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (*UnimplementedExampleServer) Insert(context.Context, *Insert_Request) (*Insert_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (*UnimplementedExampleServer) Update(context.Context, *Update_Request) (*Update_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedExampleServer) ReceiveInserts(*api.SQS_ReceiveMessageInput, Example_ReceiveInsertsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveInserts not implemented")
}
func (*UnimplementedExampleServer) SendInserts(Example_SendInsertsServer) error {
	return status.Errorf(codes.Unimplemented, "method SendInserts not implemented")
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Insert_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Example/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Insert(ctx, req.(*Insert_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Example_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Update_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Example/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Update(ctx, req.(*Update_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Example_ReceiveInserts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(api.SQS_ReceiveMessageInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServer).ReceiveInserts(m, &exampleReceiveInsertsServer{stream})
}

type Example_ReceiveInsertsServer interface {
	Send(*api.SQS_Message) error
	grpc.ServerStream
}

type exampleReceiveInsertsServer struct {
	grpc.ServerStream
}

func (x *exampleReceiveInsertsServer) Send(m *api.SQS_Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Example_SendInserts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServer).SendInserts(&exampleSendInsertsServer{stream})
}

type Example_SendInsertsServer interface {
	SendAndClose(*api.Empty) error
	Recv() (*api.SQS_SendMessageInput, error)
	grpc.ServerStream
}

type exampleSendInsertsServer struct {
	grpc.ServerStream
}

func (x *exampleSendInsertsServer) SendAndClose(m *api.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exampleSendInsertsServer) Recv() (*api.SQS_SendMessageInput, error) {
	m := new(api.SQS_SendMessageInput)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _Example_Insert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Example_Update_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveInserts",
			Handler:       _Example_ReceiveInserts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendInserts",
			Handler:       _Example_SendInserts_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "example/example.proto",
}
