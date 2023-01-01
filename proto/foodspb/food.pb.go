// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: food.proto

package foodspb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ingredients struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IDfood int64  `protobuf:"varint,2,opt,name=IDfood,proto3" json:"IDfood,omitempty"`
	Status bool   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Ingredients) Reset() {
	*x = Ingredients{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ingredients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ingredients) ProtoMessage() {}

func (x *Ingredients) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ingredients.ProtoReflect.Descriptor instead.
func (*Ingredients) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{0}
}

func (x *Ingredients) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ingredients) GetIDfood() int64 {
	if x != nil {
		return x.IDfood
	}
	return 0
}

func (x *Ingredients) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type IngredientsOnlyName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *IngredientsOnlyName) Reset() {
	*x = IngredientsOnlyName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngredientsOnlyName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngredientsOnlyName) ProtoMessage() {}

func (x *IngredientsOnlyName) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngredientsOnlyName.ProtoReflect.Descriptor instead.
func (*IngredientsOnlyName) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{1}
}

func (x *IngredientsOnlyName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreatedFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price       float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Ingredients []string `protobuf:"bytes,3,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
}

func (x *CreatedFoodRequest) Reset() {
	*x = CreatedFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedFoodRequest) ProtoMessage() {}

func (x *CreatedFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedFoodRequest.ProtoReflect.Descriptor instead.
func (*CreatedFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{2}
}

func (x *CreatedFoodRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatedFoodRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreatedFoodRequest) GetIngredients() []string {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

type Food struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price       float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Ingredients []string               `protobuf:"bytes,4,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdateAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
	Status      bool                   `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Food) Reset() {
	*x = Food{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Food) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Food) ProtoMessage() {}

func (x *Food) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Food.ProtoReflect.Descriptor instead.
func (*Food) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{3}
}

func (x *Food) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Food) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Food) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Food) GetIngredients() []string {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

func (x *Food) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Food) GetUpdateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

func (x *Food) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type FoodUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *FoodUpdateRequest) Reset() {
	*x = FoodUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodUpdateRequest) ProtoMessage() {}

func (x *FoodUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodUpdateRequest.ProtoReflect.Descriptor instead.
func (*FoodUpdateRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{4}
}

func (x *FoodUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FoodUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FoodUpdateRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type FoodWithoutIngredients struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price     float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdateAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
	Status    bool                   `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FoodWithoutIngredients) Reset() {
	*x = FoodWithoutIngredients{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodWithoutIngredients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodWithoutIngredients) ProtoMessage() {}

func (x *FoodWithoutIngredients) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodWithoutIngredients.ProtoReflect.Descriptor instead.
func (*FoodWithoutIngredients) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{5}
}

func (x *FoodWithoutIngredients) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FoodWithoutIngredients) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FoodWithoutIngredients) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *FoodWithoutIngredients) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *FoodWithoutIngredients) GetUpdateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

func (x *FoodWithoutIngredients) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type FoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FoodRequest) Reset() {
	*x = FoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodRequest) ProtoMessage() {}

func (x *FoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodRequest.ProtoReflect.Descriptor instead.
func (*FoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{6}
}

func (x *FoodRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetFoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetFoodsRequest) Reset() {
	*x = GetFoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFoodsRequest) ProtoMessage() {}

func (x *GetFoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFoodsRequest.ProtoReflect.Descriptor instead.
func (*GetFoodsRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{7}
}

func (x *GetFoodsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type FoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Food *Food `protobuf:"bytes,1,opt,name=food,proto3" json:"food,omitempty"`
}

func (x *FoodResponse) Reset() {
	*x = FoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodResponse) ProtoMessage() {}

func (x *FoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodResponse.ProtoReflect.Descriptor instead.
func (*FoodResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{8}
}

func (x *FoodResponse) GetFood() *Food {
	if x != nil {
		return x.Food
	}
	return nil
}

type FoodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foods []*Food `protobuf:"bytes,1,rep,name=foods,proto3" json:"foods,omitempty"`
}

func (x *FoodsResponse) Reset() {
	*x = FoodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodsResponse) ProtoMessage() {}

func (x *FoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodsResponse.ProtoReflect.Descriptor instead.
func (*FoodsResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{9}
}

func (x *FoodsResponse) GetFoods() []*Food {
	if x != nil {
		return x.Foods
	}
	return nil
}

type FoodDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *FoodDeleteResponse) Reset() {
	*x = FoodDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodDeleteResponse) ProtoMessage() {}

func (x *FoodDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodDeleteResponse.ProtoReflect.Descriptor instead.
func (*FoodDeleteResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{10}
}

func (x *FoodDeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_food_proto protoreflect.FileDescriptor

var file_food_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6f,
	0x6f, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x44, 0x66, 0x6f, 0x6f, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x49, 0x44, 0x66, 0x6f, 0x6f, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x29, 0x0a, 0x13, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x60, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0xec, 0x01, 0x0a, 0x04, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x4d, 0x0a, 0x11, 0x46, 0x6f, 0x6f, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0xdc, 0x01, 0x0a, 0x16, 0x46, 0x6f, 0x6f, 0x64, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75,
	0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x36, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x1d, 0x0a, 0x0b, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x0c, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64,
	0x52, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x22, 0x31, 0x0a, 0x0d, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x66, 0x6f, 0x6f, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f,
	0x6f, 0x64, 0x52, 0x05, 0x66, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x46, 0x6f, 0x6f,
	0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xaa, 0x02, 0x0a, 0x0b, 0x46, 0x6f,
	0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x18, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f,
	0x64, 0x12, 0x11, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6f, 0x64, 0x73, 0x12, 0x15, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x66, 0x6f,
	0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x17,
	0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46,
	0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x11, 0x2e, 0x66, 0x6f, 0x6f, 0x64,
	0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66,
	0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x66, 0x6f, 0x6f, 0x64, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_food_proto_rawDescOnce sync.Once
	file_food_proto_rawDescData = file_food_proto_rawDesc
)

func file_food_proto_rawDescGZIP() []byte {
	file_food_proto_rawDescOnce.Do(func() {
		file_food_proto_rawDescData = protoimpl.X.CompressGZIP(file_food_proto_rawDescData)
	})
	return file_food_proto_rawDescData
}

var file_food_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_food_proto_goTypes = []interface{}{
	(*Ingredients)(nil),            // 0: food.Ingredients
	(*IngredientsOnlyName)(nil),    // 1: food.IngredientsOnlyName
	(*CreatedFoodRequest)(nil),     // 2: food.CreatedFoodRequest
	(*Food)(nil),                   // 3: food.Food
	(*FoodUpdateRequest)(nil),      // 4: food.FoodUpdateRequest
	(*FoodWithoutIngredients)(nil), // 5: food.FoodWithoutIngredients
	(*FoodRequest)(nil),            // 6: food.FoodRequest
	(*GetFoodsRequest)(nil),        // 7: food.GetFoodsRequest
	(*FoodResponse)(nil),           // 8: food.FoodResponse
	(*FoodsResponse)(nil),          // 9: food.FoodsResponse
	(*FoodDeleteResponse)(nil),     // 10: food.FoodDeleteResponse
	(*timestamppb.Timestamp)(nil),  // 11: google.protobuf.Timestamp
}
var file_food_proto_depIdxs = []int32{
	11, // 0: food.Food.createdAt:type_name -> google.protobuf.Timestamp
	11, // 1: food.Food.updateAt:type_name -> google.protobuf.Timestamp
	11, // 2: food.FoodWithoutIngredients.createdAt:type_name -> google.protobuf.Timestamp
	11, // 3: food.FoodWithoutIngredients.updateAt:type_name -> google.protobuf.Timestamp
	3,  // 4: food.FoodResponse.food:type_name -> food.Food
	3,  // 5: food.FoodsResponse.foods:type_name -> food.Food
	2,  // 6: food.FoodService.CreatedFood:input_type -> food.CreatedFoodRequest
	6,  // 7: food.FoodService.GetFood:input_type -> food.FoodRequest
	7,  // 8: food.FoodService.GetFoods:input_type -> food.GetFoodsRequest
	4,  // 9: food.FoodService.UpdateFood:input_type -> food.FoodUpdateRequest
	6,  // 10: food.FoodService.DeleteFood:input_type -> food.FoodRequest
	8,  // 11: food.FoodService.CreatedFood:output_type -> food.FoodResponse
	8,  // 12: food.FoodService.GetFood:output_type -> food.FoodResponse
	9,  // 13: food.FoodService.GetFoods:output_type -> food.FoodsResponse
	8,  // 14: food.FoodService.UpdateFood:output_type -> food.FoodResponse
	10, // 15: food.FoodService.DeleteFood:output_type -> food.FoodDeleteResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_food_proto_init() }
func file_food_proto_init() {
	if File_food_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_food_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ingredients); i {
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
		file_food_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngredientsOnlyName); i {
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
		file_food_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedFoodRequest); i {
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
		file_food_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Food); i {
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
		file_food_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodUpdateRequest); i {
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
		file_food_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodWithoutIngredients); i {
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
		file_food_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodRequest); i {
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
		file_food_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFoodsRequest); i {
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
		file_food_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodResponse); i {
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
		file_food_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodsResponse); i {
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
		file_food_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodDeleteResponse); i {
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
			RawDescriptor: file_food_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_food_proto_goTypes,
		DependencyIndexes: file_food_proto_depIdxs,
		MessageInfos:      file_food_proto_msgTypes,
	}.Build()
	File_food_proto = out.File
	file_food_proto_rawDesc = nil
	file_food_proto_goTypes = nil
	file_food_proto_depIdxs = nil
}
