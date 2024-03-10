// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: answers.proto

package answers

import (
	questions "common/questions"
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

type MultiChoiceAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer int32 `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *MultiChoiceAnswer) Reset() {
	*x = MultiChoiceAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiChoiceAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiChoiceAnswer) ProtoMessage() {}

func (x *MultiChoiceAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_answers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiChoiceAnswer.ProtoReflect.Descriptor instead.
func (*MultiChoiceAnswer) Descriptor() ([]byte, []int) {
	return file_answers_proto_rawDescGZIP(), []int{0}
}

func (x *MultiChoiceAnswer) GetAnswer() int32 {
	if x != nil {
		return x.Answer
	}
	return 0
}

type OpenAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *OpenAnswer) Reset() {
	*x = OpenAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAnswer) ProtoMessage() {}

func (x *OpenAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_answers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAnswer.ProtoReflect.Descriptor instead.
func (*OpenAnswer) Descriptor() ([]byte, []int) {
	return file_answers_proto_rawDescGZIP(), []int{1}
}

func (x *OpenAnswer) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type questions.QuestionType `protobuf:"varint,1,opt,name=type,proto3,enum=common.questions.QuestionType" json:"type,omitempty"`
	// Types that are assignable to Question:
	//
	//	*Answer_MultiChoiceQuestion
	//	*Answer_OpenQuestion
	Question isAnswer_Question `protobuf_oneof:"question"`
	// Types that are assignable to Answer:
	//
	//	*Answer_MultiChoiceAnswer
	//	*Answer_OpenAnswer
	Answer isAnswer_Answer `protobuf_oneof:"answer"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_answers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_answers_proto_rawDescGZIP(), []int{2}
}

func (x *Answer) GetType() questions.QuestionType {
	if x != nil {
		return x.Type
	}
	return questions.QuestionType(0)
}

func (m *Answer) GetQuestion() isAnswer_Question {
	if m != nil {
		return m.Question
	}
	return nil
}

func (x *Answer) GetMultiChoiceQuestion() *questions.MultiChoiceQuestion {
	if x, ok := x.GetQuestion().(*Answer_MultiChoiceQuestion); ok {
		return x.MultiChoiceQuestion
	}
	return nil
}

func (x *Answer) GetOpenQuestion() *questions.OpenQuestion {
	if x, ok := x.GetQuestion().(*Answer_OpenQuestion); ok {
		return x.OpenQuestion
	}
	return nil
}

func (m *Answer) GetAnswer() isAnswer_Answer {
	if m != nil {
		return m.Answer
	}
	return nil
}

func (x *Answer) GetMultiChoiceAnswer() *MultiChoiceAnswer {
	if x, ok := x.GetAnswer().(*Answer_MultiChoiceAnswer); ok {
		return x.MultiChoiceAnswer
	}
	return nil
}

func (x *Answer) GetOpenAnswer() *OpenAnswer {
	if x, ok := x.GetAnswer().(*Answer_OpenAnswer); ok {
		return x.OpenAnswer
	}
	return nil
}

type isAnswer_Question interface {
	isAnswer_Question()
}

type Answer_MultiChoiceQuestion struct {
	MultiChoiceQuestion *questions.MultiChoiceQuestion `protobuf:"bytes,2,opt,name=multiChoiceQuestion,proto3,oneof"`
}

type Answer_OpenQuestion struct {
	OpenQuestion *questions.OpenQuestion `protobuf:"bytes,3,opt,name=openQuestion,proto3,oneof"`
}

func (*Answer_MultiChoiceQuestion) isAnswer_Question() {}

func (*Answer_OpenQuestion) isAnswer_Question() {}

type isAnswer_Answer interface {
	isAnswer_Answer()
}

type Answer_MultiChoiceAnswer struct {
	MultiChoiceAnswer *MultiChoiceAnswer `protobuf:"bytes,4,opt,name=multiChoiceAnswer,proto3,oneof"`
}

type Answer_OpenAnswer struct {
	OpenAnswer *OpenAnswer `protobuf:"bytes,5,opt,name=openAnswer,proto3,oneof"`
}

func (*Answer_MultiChoiceAnswer) isAnswer_Answer() {}

func (*Answer_OpenAnswer) isAnswer_Answer() {}

var File_answers_proto protoreflect.FileDescriptor

var file_answers_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x1a,
	0x0f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2b, 0x0a, 0x11, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x24, 0x0a,
	0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x22, 0x84, 0x03, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x32,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x59, 0x0a, 0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x43,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a,
	0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x11, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x2e,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x48, 0x01, 0x52, 0x11, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x6e,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x48, 0x01, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x08, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x42, 0x10, 0x5a, 0x0e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_answers_proto_rawDescOnce sync.Once
	file_answers_proto_rawDescData = file_answers_proto_rawDesc
)

func file_answers_proto_rawDescGZIP() []byte {
	file_answers_proto_rawDescOnce.Do(func() {
		file_answers_proto_rawDescData = protoimpl.X.CompressGZIP(file_answers_proto_rawDescData)
	})
	return file_answers_proto_rawDescData
}

var file_answers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_answers_proto_goTypes = []interface{}{
	(*MultiChoiceAnswer)(nil),             // 0: common.answers.MultiChoiceAnswer
	(*OpenAnswer)(nil),                    // 1: common.answers.OpenAnswer
	(*Answer)(nil),                        // 2: common.answers.Answer
	(questions.QuestionType)(0),           // 3: common.questions.QuestionType
	(*questions.MultiChoiceQuestion)(nil), // 4: common.questions.MultiChoiceQuestion
	(*questions.OpenQuestion)(nil),        // 5: common.questions.OpenQuestion
}
var file_answers_proto_depIdxs = []int32{
	3, // 0: common.answers.Answer.type:type_name -> common.questions.QuestionType
	4, // 1: common.answers.Answer.multiChoiceQuestion:type_name -> common.questions.MultiChoiceQuestion
	5, // 2: common.answers.Answer.openQuestion:type_name -> common.questions.OpenQuestion
	0, // 3: common.answers.Answer.multiChoiceAnswer:type_name -> common.answers.MultiChoiceAnswer
	1, // 4: common.answers.Answer.openAnswer:type_name -> common.answers.OpenAnswer
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_answers_proto_init() }
func file_answers_proto_init() {
	if File_answers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_answers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiChoiceAnswer); i {
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
		file_answers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAnswer); i {
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
		file_answers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
	file_answers_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Answer_MultiChoiceQuestion)(nil),
		(*Answer_OpenQuestion)(nil),
		(*Answer_MultiChoiceAnswer)(nil),
		(*Answer_OpenAnswer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_answers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_answers_proto_goTypes,
		DependencyIndexes: file_answers_proto_depIdxs,
		MessageInfos:      file_answers_proto_msgTypes,
	}.Build()
	File_answers_proto = out.File
	file_answers_proto_rawDesc = nil
	file_answers_proto_goTypes = nil
	file_answers_proto_depIdxs = nil
}