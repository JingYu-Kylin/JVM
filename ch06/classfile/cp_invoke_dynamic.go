package classfile

/**
CONSTANT_MethodType_info、CONSTANT_MethodHandle_info和CONSTANT_InvokeDynamic_info是Java SE 7才添加到class文件中的，目的是支持新增的invokedynamic指令
可以把常量池中的常量分为两类：字面量（literal）和符号引用（symbolic reference）。
字面量包括数字常量和字符串常量，符号引用包括类和接口名、字段和方法信息等。
除了字面量，其他常量都是通过索引直接或间接指向CONSTANT_Utf8_info常量
 */

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.referenceKind = reader.ReadUint8()
	self.referenceIndex = reader.ReadUint16()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.ReadUint16()
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.ReadUint16()
	self.nameAndTypeIndex = reader.ReadUint16()
}
