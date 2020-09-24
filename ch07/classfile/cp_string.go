package classfile

/**
java.lang.String字面量
CONSTANT_String_info {
	u1 tag;
	u2 string_index;
}
 */
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

/**
 * 读取常量池索引
 */
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.ReadUint16()
}

/**
 * 按索引从常量池中查找字符串
 */
func (self *ConstantStringInfo) String() string {
	return self.cp.GetUtf8(self.stringIndex)
}
