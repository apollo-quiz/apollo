// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: questions.proto

package questions

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

type QuestionType int32

const (
	QuestionType_NONE         QuestionType = 0
	QuestionType_MULTI_CHOICE QuestionType = 2
	QuestionType_OPEN_ANSWER  QuestionType = 1
)

// Enum value maps for QuestionType.
var (
	QuestionType_name = map[int32]string{
		0: "NONE",
		2: "MULTI_CHOICE",
		1: "OPEN_ANSWER",
	}
	QuestionType_value = map[string]int32{
		"NONE":         0,
		"MULTI_CHOICE": 2,
		"OPEN_ANSWER":  1,
	}
)

func (x QuestionType) Enum() *QuestionType {
	p := new(QuestionType)
	*p = x
	return p
}

func (x QuestionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuestionType) Descriptor() protoreflect.EnumDescriptor {
	return file_questions_proto_enumTypes[0].Descriptor()
}

func (QuestionType) Type() protoreflect.EnumType {
	return &file_questions_proto_enumTypes[0]
}

func (x QuestionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuestionType.Descriptor instead.
func (QuestionType) EnumDescriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{0}
}

type MultiChoiceQuestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points      int32    `protobuf:"varint,1,opt,name=points,proto3" json:"points,omitempty"`
	Option      []string `protobuf:"bytes,2,rep,name=option,proto3" json:"option,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *MultiChoiceQuestion) Reset() {
	*x = MultiChoiceQuestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiChoiceQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiChoiceQuestion) ProtoMessage() {}

func (x *MultiChoiceQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiChoiceQuestion.ProtoReflect.Descriptor instead.
func (*MultiChoiceQuestion) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{0}
}

func (x *MultiChoiceQuestion) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *MultiChoiceQuestion) GetOption() []string {
	if x != nil {
		return x.Option
	}
	return nil
}

func (x *MultiChoiceQuestion) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type OpenQuestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points      int32  `protobuf:"varint,1,opt,name=points,proto3" json:"points,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *OpenQuestion) Reset() {
	*x = OpenQuestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenQuestion) ProtoMessage() {}

func (x *OpenQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenQuestion.ProtoReflect.Descriptor instead.
func (*OpenQuestion) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{1}
}

func (x *OpenQuestion) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *OpenQuestion) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type QuestionType `protobuf:"varint,2,opt,name=type,proto3,enum=common.questions.QuestionType" json:"type,omitempty"`
	// Types that are assignable to Question:
	//
	//	*Question_MultiChoiceQuestion
	//	*Question_OpenQuestion
	Question isQuestion_Question `protobuf_oneof:"question"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{2}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetType() QuestionType {
	if x != nil {
		return x.Type
	}
	return QuestionType_NONE
}

func (m *Question) GetQuestion() isQuestion_Question {
	if m != nil {
		return m.Question
	}
	return nil
}

func (x *Question) GetMultiChoiceQuestion() *MultiChoiceQuestion {
	if x, ok := x.GetQuestion().(*Question_MultiChoiceQuestion); ok {
		return x.MultiChoiceQuestion
	}
	return nil
}

func (x *Question) GetOpenQuestion() *OpenQuestion {
	if x, ok := x.GetQuestion().(*Question_OpenQuestion); ok {
		return x.OpenQuestion
	}
	return nil
}

type isQuestion_Question interface {
	isQuestion_Question()
}

type Question_MultiChoiceQuestion struct {
	MultiChoiceQuestion *MultiChoiceQuestion `protobuf:"bytes,3,opt,name=multiChoiceQuestion,proto3,oneof"`
}

type Question_OpenQuestion struct {
	OpenQuestion *OpenQuestion `protobuf:"bytes,4,opt,name=openQuestion,proto3,oneof"`
}

func (*Question_MultiChoiceQuestion) isQuestion_Question() {}

func (*Question_OpenQuestion) isQuestion_Question() {}

var File_questions_proto protoreflect.FileDescriptor

var file_questions_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x67, 0x0a, 0x13, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x0c,
	0x4f, 0x70, 0x65, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xfb, 0x01, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x59, 0x0a, 0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x13, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x6e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x3b, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x10,
	0x01, 0x42, 0x12, 0x5a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_questions_proto_rawDescOnce sync.Once
	file_questions_proto_rawDescData = file_questions_proto_rawDesc
)

func file_questions_proto_rawDescGZIP() []byte {
	file_questions_proto_rawDescOnce.Do(func() {
		file_questions_proto_rawDescData = protoimpl.X.CompressGZIP(file_questions_proto_rawDescData)
	})
	return file_questions_proto_rawDescData
}

var file_questions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_questions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_questions_proto_goTypes = []interface{}{
	(QuestionType)(0),           // 0: common.questions.QuestionType
	(*MultiChoiceQuestion)(nil), // 1: common.questions.MultiChoiceQuestion
	(*OpenQuestion)(nil),        // 2: common.questions.OpenQuestion
	(*Question)(nil),            // 3: common.questions.Question
}
var file_questions_proto_depIdxs = []int32{
	0, // 0: common.questions.Question.type:type_name -> common.questions.QuestionType
	1, // 1: common.questions.Question.multiChoiceQuestion:type_name -> common.questions.MultiChoiceQuestion
	2, // 2: common.questions.Question.openQuestion:type_name -> common.questions.OpenQuestion
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_questions_proto_init() }
func file_questions_proto_init() {
	if File_questions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_questions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiChoiceQuestion); i {
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
		file_questions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenQuestion); i {
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
		file_questions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
	file_questions_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Question_MultiChoiceQuestion)(nil),
		(*Question_OpenQuestion)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_questions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_questions_proto_goTypes,
		DependencyIndexes: file_questions_proto_depIdxs,
		EnumInfos:         file_questions_proto_enumTypes,
		MessageInfos:      file_questions_proto_msgTypes,
	}.Build()
	File_questions_proto = out.File
	file_questions_proto_rawDesc = nil
	file_questions_proto_goTypes = nil
	file_questions_proto_depIdxs = nil
}
