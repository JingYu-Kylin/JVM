package classfile

/**
属性是可以扩展的，不同的虚拟机实现可以定义自己的属性类型。
由于这个原因，Java虚拟机规范没有使用tag，而是使用属性名来区别不同的属性。
属性数据放在属性名之后的u1表中，这样Java虚拟机实现就可以跳过自己无法识别的属性
attribute_info {
	u2 attribute_name_index;
	u4 attribute_length;
	u1 info[attribute_length];
}
属性表中存放的属性名实际上并不是编码后的字符串，而是常量池索引，指向常量池中的CONSTANT_Utf8_info常量
 */
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

/**
 * 读取属性表
 */
func ReadAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.ReadUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

/**
 * 读取单个属性
 */
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	// 先读取属性名索引
	attrNameIndex := reader.ReadUint16()
	// 根据它从常量池中找到属性名
	attrName := cp.GetUtf8(attrNameIndex)
	// 然后读取属性长度
	attrLen := reader.ReadUint32()
	// 创建具体的属性实例
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}
/**
Java虚拟机规范预定义了23种属性，先解析其中的8种
23种预定义属性可以分为三组。
第一组属性是实现Java虚拟机所必需的，共有5种；
第二组属性是Java类库所必需的，共有12种；
第三组属性主要提供给工具使用，共有6种。
第三组属性是可选的，也就是说可以不出现在class文件中。如果class文件中存在第三组属性，Java虚拟机实现或者Java类库也是可以利用它们的，比如使用LineNumberTable属性在异常堆栈中显示行号
 */
func newAttributeInfo(attrName string, attrLen uint32,
	cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code": return &CodeAttribute{cp: cp}
	case "ConstantValue": return &ConstantValueAttribute{}
	case "Deprecated": return &DeprecatedAttribute{}
	case "Exceptions": return &ExceptionsAttribute{}
	case "LineNumberTable": return &LineNumberTableAttribute{}
	case "LocalVariableTable": return &LocalVariableTableAttribute{}
	case "SourceFile": return &SourceFileAttribute{cp: cp}
	case "Synthetic": return &SyntheticAttribute{}
	default: return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
