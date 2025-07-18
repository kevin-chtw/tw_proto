// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: match2game.proto

package sproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Match2GameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartGameReq   *StartGameReq   `protobuf:"bytes,1,opt,name=start_game_req,json=startGameReq,proto3" json:"start_game_req,omitempty"`
	CancelMatchReq *CancelMatchReq `protobuf:"bytes,2,opt,name=cancel_match_req,json=cancelMatchReq,proto3" json:"cancel_match_req,omitempty"`
}

func (x *Match2GameReq) Reset() {
	*x = Match2GameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match2game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match2GameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match2GameReq) ProtoMessage() {}

func (x *Match2GameReq) ProtoReflect() protoreflect.Message {
	mi := &file_match2game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match2GameReq.ProtoReflect.Descriptor instead.
func (*Match2GameReq) Descriptor() ([]byte, []int) {
	return file_match2game_proto_rawDescGZIP(), []int{0}
}

func (x *Match2GameReq) GetStartGameReq() *StartGameReq {
	if x != nil {
		return x.StartGameReq
	}
	return nil
}

func (x *Match2GameReq) GetCancelMatchReq() *CancelMatchReq {
	if x != nil {
		return x.CancelMatchReq
	}
	return nil
}

type StartGameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchServerId string   `protobuf:"bytes,1,opt,name=match_server_id,json=matchServerId,proto3" json:"match_server_id,omitempty"` // 匹配服务器ID
	Matchid       string   `protobuf:"bytes,2,opt,name=matchid,proto3" json:"matchid,omitempty"`
	Tableid       string   `protobuf:"bytes,3,opt,name=tableid,proto3" json:"tableid,omitempty"`
	Players       []string `protobuf:"bytes,4,rep,name=players,proto3" json:"players,omitempty"`
	GameConfig    []string `protobuf:"bytes,5,rep,name=game_config,json=gameConfig,proto3" json:"game_config,omitempty"`
	InitialChips  int32    `protobuf:"varint,6,opt,name=initial_chips,json=initialChips,proto3" json:"initial_chips,omitempty"`
	ScoreBase     int32    `protobuf:"varint,7,opt,name=score_base,json=scoreBase,proto3" json:"score_base,omitempty"`
}

func (x *StartGameReq) Reset() {
	*x = StartGameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match2game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameReq) ProtoMessage() {}

func (x *StartGameReq) ProtoReflect() protoreflect.Message {
	mi := &file_match2game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameReq.ProtoReflect.Descriptor instead.
func (*StartGameReq) Descriptor() ([]byte, []int) {
	return file_match2game_proto_rawDescGZIP(), []int{1}
}

func (x *StartGameReq) GetMatchServerId() string {
	if x != nil {
		return x.MatchServerId
	}
	return ""
}

func (x *StartGameReq) GetMatchid() string {
	if x != nil {
		return x.Matchid
	}
	return ""
}

func (x *StartGameReq) GetTableid() string {
	if x != nil {
		return x.Tableid
	}
	return ""
}

func (x *StartGameReq) GetPlayers() []string {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *StartGameReq) GetGameConfig() []string {
	if x != nil {
		return x.GameConfig
	}
	return nil
}

func (x *StartGameReq) GetInitialChips() int32 {
	if x != nil {
		return x.InitialChips
	}
	return 0
}

func (x *StartGameReq) GetScoreBase() int32 {
	if x != nil {
		return x.ScoreBase
	}
	return 0
}

type CancelMatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matchid string `protobuf:"bytes,1,opt,name=matchid,proto3" json:"matchid,omitempty"`
	Tableid string `protobuf:"bytes,2,opt,name=tableid,proto3" json:"tableid,omitempty"`
}

func (x *CancelMatchReq) Reset() {
	*x = CancelMatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match2game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelMatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelMatchReq) ProtoMessage() {}

func (x *CancelMatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_match2game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelMatchReq.ProtoReflect.Descriptor instead.
func (*CancelMatchReq) Descriptor() ([]byte, []int) {
	return file_match2game_proto_rawDescGZIP(), []int{2}
}

func (x *CancelMatchReq) GetMatchid() string {
	if x != nil {
		return x.Matchid
	}
	return ""
}

func (x *CancelMatchReq) GetTableid() string {
	if x != nil {
		return x.Tableid
	}
	return ""
}

type Match2GameAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameResultAck *GameResultAck `protobuf:"bytes,1,opt,name=game_result_ack,json=gameResultAck,proto3" json:"game_result_ack,omitempty"`
}

func (x *Match2GameAck) Reset() {
	*x = Match2GameAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match2game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match2GameAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match2GameAck) ProtoMessage() {}

func (x *Match2GameAck) ProtoReflect() protoreflect.Message {
	mi := &file_match2game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match2GameAck.ProtoReflect.Descriptor instead.
func (*Match2GameAck) Descriptor() ([]byte, []int) {
	return file_match2game_proto_rawDescGZIP(), []int{3}
}

func (x *Match2GameAck) GetGameResultAck() *GameResultAck {
	if x != nil {
		return x.GameResultAck
	}
	return nil
}

type GameResultAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matchid   string  `protobuf:"bytes,1,opt,name=matchid,proto3" json:"matchid,omitempty"`
	Tableid   string  `protobuf:"bytes,2,opt,name=tableid,proto3" json:"tableid,omitempty"`
	EndReason int32   `protobuf:"varint,3,opt,name=end_reason,json=endReason,proto3" json:"end_reason,omitempty"`
	Scores    []int32 `protobuf:"varint,4,rep,packed,name=scores,proto3" json:"scores,omitempty"`
}

func (x *GameResultAck) Reset() {
	*x = GameResultAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match2game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResultAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResultAck) ProtoMessage() {}

func (x *GameResultAck) ProtoReflect() protoreflect.Message {
	mi := &file_match2game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameResultAck.ProtoReflect.Descriptor instead.
func (*GameResultAck) Descriptor() ([]byte, []int) {
	return file_match2game_proto_rawDescGZIP(), []int{4}
}

func (x *GameResultAck) GetMatchid() string {
	if x != nil {
		return x.Matchid
	}
	return ""
}

func (x *GameResultAck) GetTableid() string {
	if x != nil {
		return x.Tableid
	}
	return ""
}

func (x *GameResultAck) GetEndReason() int32 {
	if x != nil {
		return x.EndReason
	}
	return 0
}

func (x *GameResultAck) GetScores() []int32 {
	if x != nil {
		return x.Scores
	}
	return nil
}

var File_match2game_proto protoreflect.FileDescriptor

var file_match2game_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x32, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x3a, 0x0a, 0x0e,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x40, 0x0a, 0x10, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x22, 0xe9, 0x01, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x0f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x68,
	0x69, 0x70, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x43, 0x68, 0x69, 0x70, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x42, 0x61, 0x73, 0x65, 0x22, 0x44, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x0d,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x32, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x3d, 0x0a,
	0x0f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x6b, 0x52, 0x0d, 0x67,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x6b, 0x22, 0x7a, 0x0a, 0x0d,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x73,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match2game_proto_rawDescOnce sync.Once
	file_match2game_proto_rawDescData = file_match2game_proto_rawDesc
)

func file_match2game_proto_rawDescGZIP() []byte {
	file_match2game_proto_rawDescOnce.Do(func() {
		file_match2game_proto_rawDescData = protoimpl.X.CompressGZIP(file_match2game_proto_rawDescData)
	})
	return file_match2game_proto_rawDescData
}

var file_match2game_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_match2game_proto_goTypes = []interface{}{
	(*Match2GameReq)(nil),  // 0: sproto.Match2GameReq
	(*StartGameReq)(nil),   // 1: sproto.StartGameReq
	(*CancelMatchReq)(nil), // 2: sproto.CancelMatchReq
	(*Match2GameAck)(nil),  // 3: sproto.Match2GameAck
	(*GameResultAck)(nil),  // 4: sproto.GameResultAck
}
var file_match2game_proto_depIdxs = []int32{
	1, // 0: sproto.Match2GameReq.start_game_req:type_name -> sproto.StartGameReq
	2, // 1: sproto.Match2GameReq.cancel_match_req:type_name -> sproto.CancelMatchReq
	4, // 2: sproto.Match2GameAck.game_result_ack:type_name -> sproto.GameResultAck
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_match2game_proto_init() }
func file_match2game_proto_init() {
	if File_match2game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match2game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Match2GameReq); i {
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
		file_match2game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameReq); i {
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
		file_match2game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelMatchReq); i {
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
		file_match2game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Match2GameAck); i {
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
		file_match2game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameResultAck); i {
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
			RawDescriptor: file_match2game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_match2game_proto_goTypes,
		DependencyIndexes: file_match2game_proto_depIdxs,
		MessageInfos:      file_match2game_proto_msgTypes,
	}.Build()
	File_match2game_proto = out.File
	file_match2game_proto_rawDesc = nil
	file_match2game_proto_goTypes = nil
	file_match2game_proto_depIdxs = nil
}
