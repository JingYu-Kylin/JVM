package classfile

/**
CONSTANT_Fieldref_info表示字段符号引用，
CONSTANT_Methodref_info表示普通（非接口）方法符号引用，
CONSTANT_InterfaceMethodref_info表示接口方法符号引用

这三种常量结构一模一样
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */

// 三个结构体“继承”ConstantMemberrefInfo
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}
func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.ReadUint16()
	self.nameAndTypeIndex = reader.ReadUint16()
}
func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.GetClassName(self.classIndex)
}
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}