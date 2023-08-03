// Code generated by go-vk from vk.xml at 2023-08-03 23:04:04.862517702 +0200 CEST m=+2.841185509. DO NOT EDIT.

package vk

//go:generate stringer -output=enum_metal_string_0.go -type=ExportMetalObjectTypeFlagBitsEXT
type ExportMetalObjectTypeFlagBitsEXT = ExportMetalObjectTypeFlagsEXT

const (
	EXPORT_METAL_OBJECT_TYPE_METAL_DEVICE_BIT_EXT        ExportMetalObjectTypeFlagBitsEXT = 1 << 0
	EXPORT_METAL_OBJECT_TYPE_METAL_COMMAND_QUEUE_BIT_EXT ExportMetalObjectTypeFlagBitsEXT = 1 << 1
	EXPORT_METAL_OBJECT_TYPE_METAL_BUFFER_BIT_EXT        ExportMetalObjectTypeFlagBitsEXT = 1 << 2
	EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT       ExportMetalObjectTypeFlagBitsEXT = 1 << 3
	EXPORT_METAL_OBJECT_TYPE_METAL_IOSURFACE_BIT_EXT     ExportMetalObjectTypeFlagBitsEXT = 1 << 4
	EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT  ExportMetalObjectTypeFlagBitsEXT = 1 << 5
)

// Platform-specific values for VkStructureType
const (
	STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT       StructureType = 1000217000
	STRUCTURE_TYPE_EXPORT_METAL_OBJECT_CREATE_INFO_EXT StructureType = 1000311000
	STRUCTURE_TYPE_EXPORT_METAL_OBJECTS_INFO_EXT       StructureType = 1000311001
	STRUCTURE_TYPE_EXPORT_METAL_DEVICE_INFO_EXT        StructureType = 1000311002
	STRUCTURE_TYPE_EXPORT_METAL_COMMAND_QUEUE_INFO_EXT StructureType = 1000311003
	STRUCTURE_TYPE_EXPORT_METAL_BUFFER_INFO_EXT        StructureType = 1000311004
	STRUCTURE_TYPE_IMPORT_METAL_BUFFER_INFO_EXT        StructureType = 1000311005
	STRUCTURE_TYPE_EXPORT_METAL_TEXTURE_INFO_EXT       StructureType = 1000311006
	STRUCTURE_TYPE_IMPORT_METAL_TEXTURE_INFO_EXT       StructureType = 1000311007
	STRUCTURE_TYPE_EXPORT_METAL_IO_SURFACE_INFO_EXT    StructureType = 1000311008
	STRUCTURE_TYPE_IMPORT_METAL_IO_SURFACE_INFO_EXT    StructureType = 1000311009
	STRUCTURE_TYPE_EXPORT_METAL_SHARED_EVENT_INFO_EXT  StructureType = 1000311010
	STRUCTURE_TYPE_IMPORT_METAL_SHARED_EVENT_INFO_EXT  StructureType = 1000311011
)
