package classfile

/**
SourceFile是可选定长属性
只会出现在ClassFile结构中，用于指出源文件名
SourceFile_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	u2 sourcefile_index;
}
attribute_length的值必须是2
sourcefile_index是常量池索引，指向CONSTANT_Utf8_info常量
 */
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}
func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.ReadUint16()
}
func (self *SourceFileAttribute) FileName() string {
	return self.cp.GetUtf8(self.sourceFileIndex)
}
