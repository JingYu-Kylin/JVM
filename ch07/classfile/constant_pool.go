package classfile

/**
常量池
表头给出的常量池大小比实际大1。假设表头给出的值是n，那么常量池的实际大小是n–1。
有效的常量池索引是1~n–1。0是无效索引，表示不指向任何常量。
CONSTANT_Long_info和CONSTANT_Double_info各占两个位置。也就是说，如果常量池中存在这两种常量，实际的常量数量比n–1还要少，而且1~n–1的某些数也会变成无效索引
 */

type ConstantPool []ConstantInfo

/**
 * 读取常量池
 */
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.ReadUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ { // 注意索引从1开始
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 占两个位置
		}
	}
	return cp
}

/**
 * 按索引查找常量
 */
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

/**
 * 从常量池查找字段或方法的名字和描述符
 */
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.GetUtf8(ntInfo.nameIndex)
	_type := self.GetUtf8(ntInfo.descriptorIndex)
	return name, _type
}

/**
 * 从常量池查找类名
 */
func (self ConstantPool) GetClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.GetUtf8(classInfo.nameIndex)
}

/**
 * 从常量池查找UTF-8字符串
 */
func (self ConstantPool) GetUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
