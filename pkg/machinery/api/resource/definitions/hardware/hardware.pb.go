// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: resource/definitions/hardware/hardware.proto

package hardware

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MemoryModuleSpec represents a single Memory.
type MemoryModuleSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size          uint32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	DeviceLocator string `protobuf:"bytes,2,opt,name=device_locator,json=deviceLocator,proto3" json:"device_locator,omitempty"`
	BankLocator   string `protobuf:"bytes,3,opt,name=bank_locator,json=bankLocator,proto3" json:"bank_locator,omitempty"`
	Speed         uint32 `protobuf:"varint,4,opt,name=speed,proto3" json:"speed,omitempty"`
	Manufacturer  string `protobuf:"bytes,5,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	SerialNumber  string `protobuf:"bytes,6,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	AssetTag      string `protobuf:"bytes,7,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	ProductName   string `protobuf:"bytes,8,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
}

func (x *MemoryModuleSpec) Reset() {
	*x = MemoryModuleSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryModuleSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryModuleSpec) ProtoMessage() {}

func (x *MemoryModuleSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryModuleSpec.ProtoReflect.Descriptor instead.
func (*MemoryModuleSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_hardware_hardware_proto_rawDescGZIP(), []int{0}
}

func (x *MemoryModuleSpec) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MemoryModuleSpec) GetDeviceLocator() string {
	if x != nil {
		return x.DeviceLocator
	}
	return ""
}

func (x *MemoryModuleSpec) GetBankLocator() string {
	if x != nil {
		return x.BankLocator
	}
	return ""
}

func (x *MemoryModuleSpec) GetSpeed() uint32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *MemoryModuleSpec) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *MemoryModuleSpec) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *MemoryModuleSpec) GetAssetTag() string {
	if x != nil {
		return x.AssetTag
	}
	return ""
}

func (x *MemoryModuleSpec) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

// PCIDeviceSpec represents a single processor.
type PCIDeviceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class      string `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	Subclass   string `protobuf:"bytes,2,opt,name=subclass,proto3" json:"subclass,omitempty"`
	Vendor     string `protobuf:"bytes,3,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Product    string `protobuf:"bytes,4,opt,name=product,proto3" json:"product,omitempty"`
	ClassId    string `protobuf:"bytes,5,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	SubclassId string `protobuf:"bytes,6,opt,name=subclass_id,json=subclassId,proto3" json:"subclass_id,omitempty"`
	VendorId   string `protobuf:"bytes,7,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	ProductId  string `protobuf:"bytes,8,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *PCIDeviceSpec) Reset() {
	*x = PCIDeviceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PCIDeviceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCIDeviceSpec) ProtoMessage() {}

func (x *PCIDeviceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCIDeviceSpec.ProtoReflect.Descriptor instead.
func (*PCIDeviceSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_hardware_hardware_proto_rawDescGZIP(), []int{1}
}

func (x *PCIDeviceSpec) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *PCIDeviceSpec) GetSubclass() string {
	if x != nil {
		return x.Subclass
	}
	return ""
}

func (x *PCIDeviceSpec) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *PCIDeviceSpec) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *PCIDeviceSpec) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *PCIDeviceSpec) GetSubclassId() string {
	if x != nil {
		return x.SubclassId
	}
	return ""
}

func (x *PCIDeviceSpec) GetVendorId() string {
	if x != nil {
		return x.VendorId
	}
	return ""
}

func (x *PCIDeviceSpec) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

// ProcessorSpec represents a single processor.
type ProcessorSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Socket       string `protobuf:"bytes,1,opt,name=socket,proto3" json:"socket,omitempty"`
	Manufacturer string `protobuf:"bytes,2,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	ProductName  string `protobuf:"bytes,3,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	MaxSpeed     uint32 `protobuf:"varint,4,opt,name=max_speed,json=maxSpeed,proto3" json:"max_speed,omitempty"`
	BootSpeed    uint32 `protobuf:"varint,5,opt,name=boot_speed,json=bootSpeed,proto3" json:"boot_speed,omitempty"`
	Status       uint32 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	SerialNumber string `protobuf:"bytes,7,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	AssetTag     string `protobuf:"bytes,8,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	PartNumber   string `protobuf:"bytes,9,opt,name=part_number,json=partNumber,proto3" json:"part_number,omitempty"`
	CoreCount    uint32 `protobuf:"varint,10,opt,name=core_count,json=coreCount,proto3" json:"core_count,omitempty"`
	CoreEnabled  uint32 `protobuf:"varint,11,opt,name=core_enabled,json=coreEnabled,proto3" json:"core_enabled,omitempty"`
	ThreadCount  uint32 `protobuf:"varint,12,opt,name=thread_count,json=threadCount,proto3" json:"thread_count,omitempty"`
}

func (x *ProcessorSpec) Reset() {
	*x = ProcessorSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessorSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessorSpec) ProtoMessage() {}

func (x *ProcessorSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessorSpec.ProtoReflect.Descriptor instead.
func (*ProcessorSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_hardware_hardware_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessorSpec) GetSocket() string {
	if x != nil {
		return x.Socket
	}
	return ""
}

func (x *ProcessorSpec) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *ProcessorSpec) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *ProcessorSpec) GetMaxSpeed() uint32 {
	if x != nil {
		return x.MaxSpeed
	}
	return 0
}

func (x *ProcessorSpec) GetBootSpeed() uint32 {
	if x != nil {
		return x.BootSpeed
	}
	return 0
}

func (x *ProcessorSpec) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ProcessorSpec) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *ProcessorSpec) GetAssetTag() string {
	if x != nil {
		return x.AssetTag
	}
	return ""
}

func (x *ProcessorSpec) GetPartNumber() string {
	if x != nil {
		return x.PartNumber
	}
	return ""
}

func (x *ProcessorSpec) GetCoreCount() uint32 {
	if x != nil {
		return x.CoreCount
	}
	return 0
}

func (x *ProcessorSpec) GetCoreEnabled() uint32 {
	if x != nil {
		return x.CoreEnabled
	}
	return 0
}

func (x *ProcessorSpec) GetThreadCount() uint32 {
	if x != nil {
		return x.ThreadCount
	}
	return 0
}

// SystemInformationSpec represents the system information obtained from smbios.
type SystemInformationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manufacturer string `protobuf:"bytes,1,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	ProductName  string `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Version      string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	SerialNumber string `protobuf:"bytes,4,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Uuid         string `protobuf:"bytes,5,opt,name=uuid,proto3" json:"uuid,omitempty"`
	WakeUpType   string `protobuf:"bytes,6,opt,name=wake_up_type,json=wakeUpType,proto3" json:"wake_up_type,omitempty"`
	SkuNumber    string `protobuf:"bytes,7,opt,name=sku_number,json=skuNumber,proto3" json:"sku_number,omitempty"`
}

func (x *SystemInformationSpec) Reset() {
	*x = SystemInformationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemInformationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemInformationSpec) ProtoMessage() {}

func (x *SystemInformationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemInformationSpec.ProtoReflect.Descriptor instead.
func (*SystemInformationSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_hardware_hardware_proto_rawDescGZIP(), []int{3}
}

func (x *SystemInformationSpec) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *SystemInformationSpec) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *SystemInformationSpec) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SystemInformationSpec) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *SystemInformationSpec) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SystemInformationSpec) GetWakeUpType() string {
	if x != nil {
		return x.WakeUpType
	}
	return ""
}

func (x *SystemInformationSpec) GetSkuNumber() string {
	if x != nil {
		return x.SkuNumber
	}
	return ""
}

var File_resource_definitions_hardware_hardware_proto protoreflect.FileDescriptor

var file_resource_definitions_hardware_hardware_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x2f,
	0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23,
	0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x64, 0x77,
	0x61, 0x72, 0x65, 0x22, 0x8f, 0x02, 0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x6e, 0x6b, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74,
	0x61, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54,
	0x61, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xeb, 0x01, 0x0a, 0x0d, 0x50, 0x43, 0x49, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x75, 0x62, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x75, 0x62, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x62,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x22, 0x8a, 0x03, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x74, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x72, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x61, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xf2, 0x01, 0x0a, 0x15, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x77, 0x61, 0x6b, 0x65, 0x5f, 0x75, 0x70, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x61, 0x6b, 0x65,
	0x55, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6b, 0x75, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6b, 0x75, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x7a, 0x0a, 0x2b, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x61, 0x6c,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x64,
	0x77, 0x61, 0x72, 0x65, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f,
	0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_definitions_hardware_hardware_proto_rawDescOnce sync.Once
	file_resource_definitions_hardware_hardware_proto_rawDescData = file_resource_definitions_hardware_hardware_proto_rawDesc
)

func file_resource_definitions_hardware_hardware_proto_rawDescGZIP() []byte {
	file_resource_definitions_hardware_hardware_proto_rawDescOnce.Do(func() {
		file_resource_definitions_hardware_hardware_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_definitions_hardware_hardware_proto_rawDescData)
	})
	return file_resource_definitions_hardware_hardware_proto_rawDescData
}

var file_resource_definitions_hardware_hardware_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_resource_definitions_hardware_hardware_proto_goTypes = []any{
	(*MemoryModuleSpec)(nil),      // 0: talos.resource.definitions.hardware.MemoryModuleSpec
	(*PCIDeviceSpec)(nil),         // 1: talos.resource.definitions.hardware.PCIDeviceSpec
	(*ProcessorSpec)(nil),         // 2: talos.resource.definitions.hardware.ProcessorSpec
	(*SystemInformationSpec)(nil), // 3: talos.resource.definitions.hardware.SystemInformationSpec
}
var file_resource_definitions_hardware_hardware_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resource_definitions_hardware_hardware_proto_init() }
func file_resource_definitions_hardware_hardware_proto_init() {
	if File_resource_definitions_hardware_hardware_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_definitions_hardware_hardware_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MemoryModuleSpec); i {
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
		file_resource_definitions_hardware_hardware_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PCIDeviceSpec); i {
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
		file_resource_definitions_hardware_hardware_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ProcessorSpec); i {
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
		file_resource_definitions_hardware_hardware_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SystemInformationSpec); i {
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
			RawDescriptor: file_resource_definitions_hardware_hardware_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_definitions_hardware_hardware_proto_goTypes,
		DependencyIndexes: file_resource_definitions_hardware_hardware_proto_depIdxs,
		MessageInfos:      file_resource_definitions_hardware_hardware_proto_msgTypes,
	}.Build()
	File_resource_definitions_hardware_hardware_proto = out.File
	file_resource_definitions_hardware_hardware_proto_rawDesc = nil
	file_resource_definitions_hardware_hardware_proto_goTypes = nil
	file_resource_definitions_hardware_hardware_proto_depIdxs = nil
}
