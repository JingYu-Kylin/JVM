package classfile

/**
类或者接口的符号引用
CONSTANT_Class_info {
	u1 tag;
	u2 name_index;
}
类和超类索引，以及接口表中的接口索引指向的都是CONSTANT_Class_info常量
 */

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.ReadUint16()
}
func (self *ConstantClassInfo) Name() string {
	return self.cp.GetUtf8(self.nameIndex)
}
