// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.6
// source: proto/review/review_service.proto

package review

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_review_review_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[0]
	if x != nil {
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
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{0}
}

type Review struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RecipeId      string                 `protobuf:"bytes,2,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Rating        float32                `protobuf:"fixed32,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Comment       string                 `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Review) Reset() {
	*x = Review{}
	mi := &file_proto_review_review_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{1}
}

func (x *Review) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Review) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *Review) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Review) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Review) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type ReviewNats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AuthorId      string                 `protobuf:"bytes,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	RecipeId      string                 `protobuf:"bytes,2,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	Rating        float32                `protobuf:"fixed32,3,opt,name=rating,proto3" json:"rating,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewNats) Reset() {
	*x = ReviewNats{}
	mi := &file_proto_review_review_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewNats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewNats) ProtoMessage() {}

func (x *ReviewNats) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewNats.ProtoReflect.Descriptor instead.
func (*ReviewNats) Descriptor() ([]byte, []int) {
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReviewNats) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *ReviewNats) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *ReviewNats) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type ReviewCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecipeId      string                 `protobuf:"bytes,1,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Rating        float32                `protobuf:"fixed32,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Comment       string                 `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewCreateRequest) Reset() {
	*x = ReviewCreateRequest{}
	mi := &file_proto_review_review_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewCreateRequest) ProtoMessage() {}

func (x *ReviewCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewCreateRequest.ProtoReflect.Descriptor instead.
func (*ReviewCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReviewCreateRequest) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *ReviewCreateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReviewCreateRequest) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *ReviewCreateRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type ReviewCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewCreateResponse) Reset() {
	*x = ReviewCreateResponse{}
	mi := &file_proto_review_review_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewCreateResponse) ProtoMessage() {}

func (x *ReviewCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewCreateResponse.ProtoReflect.Descriptor instead.
func (*ReviewCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{4}
}

func (x *ReviewCreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReviewGetListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reviews       []*Review              `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewGetListResponse) Reset() {
	*x = ReviewGetListResponse{}
	mi := &file_proto_review_review_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewGetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewGetListResponse) ProtoMessage() {}

func (x *ReviewGetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewGetListResponse.ProtoReflect.Descriptor instead.
func (*ReviewGetListResponse) Descriptor() ([]byte, []int) {
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{5}
}

func (x *ReviewGetListResponse) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type ReviewGetByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewGetByIdRequest) Reset() {
	*x = ReviewGetByIdRequest{}
	mi := &file_proto_review_review_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewGetByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewGetByIdRequest) ProtoMessage() {}

func (x *ReviewGetByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewGetByIdRequest.ProtoReflect.Descriptor instead.
func (*ReviewGetByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{6}
}

func (x *ReviewGetByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReviewGetByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Review        *Review                `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewGetByIdResponse) Reset() {
	*x = ReviewGetByIdResponse{}
	mi := &file_proto_review_review_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewGetByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewGetByIdResponse) ProtoMessage() {}

func (x *ReviewGetByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewGetByIdResponse.ProtoReflect.Descriptor instead.
func (*ReviewGetByIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{7}
}

func (x *ReviewGetByIdResponse) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

type ReviewUpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RecipeId      string                 `protobuf:"bytes,2,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Rating        float32                `protobuf:"fixed32,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Comment       string                 `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewUpdateRequest) Reset() {
	*x = ReviewUpdateRequest{}
	mi := &file_proto_review_review_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewUpdateRequest) ProtoMessage() {}

func (x *ReviewUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewUpdateRequest.ProtoReflect.Descriptor instead.
func (*ReviewUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{8}
}

func (x *ReviewUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReviewUpdateRequest) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *ReviewUpdateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReviewUpdateRequest) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *ReviewUpdateRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type ReviewUpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Review        *Review                `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewUpdateResponse) Reset() {
	*x = ReviewUpdateResponse{}
	mi := &file_proto_review_review_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewUpdateResponse) ProtoMessage() {}

func (x *ReviewUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewUpdateResponse.ProtoReflect.Descriptor instead.
func (*ReviewUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{9}
}

func (x *ReviewUpdateResponse) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

type ReviewDeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewDeleteRequest) Reset() {
	*x = ReviewDeleteRequest{}
	mi := &file_proto_review_review_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewDeleteRequest) ProtoMessage() {}

func (x *ReviewDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewDeleteRequest.ProtoReflect.Descriptor instead.
func (*ReviewDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{10}
}

func (x *ReviewDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReviewDeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewDeleteResponse) Reset() {
	*x = ReviewDeleteResponse{}
	mi := &file_proto_review_review_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewDeleteResponse) ProtoMessage() {}

func (x *ReviewDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewDeleteResponse.ProtoReflect.Descriptor instead.
func (*ReviewDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_review_review_service_proto_rawDescGZIP(), []int{11}
}

func (x *ReviewDeleteResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_review_review_service_proto protoreflect.FileDescriptor

const file_proto_review_review_service_proto_rawDesc = "" +
	"\n" +
	"!proto/review/review_service.proto\x12\x06review\"\a\n" +
	"\x05Empty\"\x80\x01\n" +
	"\x06Review\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\trecipe_id\x18\x02 \x01(\tR\brecipeId\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\tR\x06userId\x12\x16\n" +
	"\x06rating\x18\x04 \x01(\x02R\x06rating\x12\x18\n" +
	"\acomment\x18\x05 \x01(\tR\acomment\"^\n" +
	"\n" +
	"ReviewNats\x12\x1b\n" +
	"\tauthor_id\x18\x01 \x01(\tR\bauthorId\x12\x1b\n" +
	"\trecipe_id\x18\x02 \x01(\tR\brecipeId\x12\x16\n" +
	"\x06rating\x18\x03 \x01(\x02R\x06rating\"}\n" +
	"\x13ReviewCreateRequest\x12\x1b\n" +
	"\trecipe_id\x18\x01 \x01(\tR\brecipeId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x16\n" +
	"\x06rating\x18\x03 \x01(\x02R\x06rating\x12\x18\n" +
	"\acomment\x18\x04 \x01(\tR\acomment\"&\n" +
	"\x14ReviewCreateResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"A\n" +
	"\x15ReviewGetListResponse\x12(\n" +
	"\areviews\x18\x01 \x03(\v2\x0e.review.ReviewR\areviews\"&\n" +
	"\x14ReviewGetByIdRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"?\n" +
	"\x15ReviewGetByIdResponse\x12&\n" +
	"\x06review\x18\x01 \x01(\v2\x0e.review.ReviewR\x06review\"\x8d\x01\n" +
	"\x13ReviewUpdateRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\trecipe_id\x18\x02 \x01(\tR\brecipeId\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\tR\x06userId\x12\x16\n" +
	"\x06rating\x18\x04 \x01(\x02R\x06rating\x12\x18\n" +
	"\acomment\x18\x05 \x01(\tR\acomment\">\n" +
	"\x14ReviewUpdateResponse\x12&\n" +
	"\x06review\x18\x01 \x01(\v2\x0e.review.ReviewR\x06review\"%\n" +
	"\x13ReviewDeleteRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\".\n" +
	"\x14ReviewDeleteResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2\xfd\x02\n" +
	"\rReviewService\x12I\n" +
	"\fReviewCreate\x12\x1b.review.ReviewCreateRequest\x1a\x1c.review.ReviewCreateResponse\x12=\n" +
	"\rReviewGetList\x12\r.review.Empty\x1a\x1d.review.ReviewGetListResponse\x12L\n" +
	"\rReviewGetById\x12\x1c.review.ReviewGetByIdRequest\x1a\x1d.review.ReviewGetByIdResponse\x12I\n" +
	"\fReviewUpdate\x12\x1b.review.ReviewUpdateRequest\x1a\x1c.review.ReviewUpdateResponse\x12I\n" +
	"\fReviewDelete\x12\x1b.review.ReviewDeleteRequest\x1a\x1c.review.ReviewDeleteResponseB\n" +
	"Z\b./reviewb\x06proto3"

var (
	file_proto_review_review_service_proto_rawDescOnce sync.Once
	file_proto_review_review_service_proto_rawDescData []byte
)

func file_proto_review_review_service_proto_rawDescGZIP() []byte {
	file_proto_review_review_service_proto_rawDescOnce.Do(func() {
		file_proto_review_review_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_review_review_service_proto_rawDesc), len(file_proto_review_review_service_proto_rawDesc)))
	})
	return file_proto_review_review_service_proto_rawDescData
}

var file_proto_review_review_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_review_review_service_proto_goTypes = []any{
	(*Empty)(nil),                 // 0: review.Empty
	(*Review)(nil),                // 1: review.Review
	(*ReviewNats)(nil),            // 2: review.ReviewNats
	(*ReviewCreateRequest)(nil),   // 3: review.ReviewCreateRequest
	(*ReviewCreateResponse)(nil),  // 4: review.ReviewCreateResponse
	(*ReviewGetListResponse)(nil), // 5: review.ReviewGetListResponse
	(*ReviewGetByIdRequest)(nil),  // 6: review.ReviewGetByIdRequest
	(*ReviewGetByIdResponse)(nil), // 7: review.ReviewGetByIdResponse
	(*ReviewUpdateRequest)(nil),   // 8: review.ReviewUpdateRequest
	(*ReviewUpdateResponse)(nil),  // 9: review.ReviewUpdateResponse
	(*ReviewDeleteRequest)(nil),   // 10: review.ReviewDeleteRequest
	(*ReviewDeleteResponse)(nil),  // 11: review.ReviewDeleteResponse
}
var file_proto_review_review_service_proto_depIdxs = []int32{
	1,  // 0: review.ReviewGetListResponse.reviews:type_name -> review.Review
	1,  // 1: review.ReviewGetByIdResponse.review:type_name -> review.Review
	1,  // 2: review.ReviewUpdateResponse.review:type_name -> review.Review
	3,  // 3: review.ReviewService.ReviewCreate:input_type -> review.ReviewCreateRequest
	0,  // 4: review.ReviewService.ReviewGetList:input_type -> review.Empty
	6,  // 5: review.ReviewService.ReviewGetById:input_type -> review.ReviewGetByIdRequest
	8,  // 6: review.ReviewService.ReviewUpdate:input_type -> review.ReviewUpdateRequest
	10, // 7: review.ReviewService.ReviewDelete:input_type -> review.ReviewDeleteRequest
	4,  // 8: review.ReviewService.ReviewCreate:output_type -> review.ReviewCreateResponse
	5,  // 9: review.ReviewService.ReviewGetList:output_type -> review.ReviewGetListResponse
	7,  // 10: review.ReviewService.ReviewGetById:output_type -> review.ReviewGetByIdResponse
	9,  // 11: review.ReviewService.ReviewUpdate:output_type -> review.ReviewUpdateResponse
	11, // 12: review.ReviewService.ReviewDelete:output_type -> review.ReviewDeleteResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_review_review_service_proto_init() }
func file_proto_review_review_service_proto_init() {
	if File_proto_review_review_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_review_review_service_proto_rawDesc), len(file_proto_review_review_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_review_review_service_proto_goTypes,
		DependencyIndexes: file_proto_review_review_service_proto_depIdxs,
		MessageInfos:      file_proto_review_review_service_proto_msgTypes,
	}.Build()
	File_proto_review_review_service_proto = out.File
	file_proto_review_review_service_proto_goTypes = nil
	file_proto_review_review_service_proto_depIdxs = nil
}
