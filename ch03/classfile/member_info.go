package classfile

/*
字段和方法
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

/**
 * 和类一样，字段和方法也有自己的访问标志。
 * 访问标志之后是一个常量池索引，给出字段名或方法名，然后又是一个常量池索引，给出字段或方法的描述符，最后是属性表。
 */
type MemberInfo struct {
	// cp字段保存常量池指针，后面会用到它
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}
/**
 * 读取字段表或方法表
 */
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.ReadUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

/**
 * 读取字段或方法数据
 */
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.ReadUint16(),
		nameIndex:       reader.ReadUint16(),
		descriptorIndex: reader.ReadUint16(),
		attributes:      ReadAttributes(reader, cp),
	}
}

/**
 * 获取访问标志
 */
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

/**
 * 从常量池查找字段或方法名
 */
func (self *MemberInfo) Name() string {
	return self.cp.GetUtf8(self.nameIndex)
}

/**
 * 从常量池查找字段或方法描述符
 */
func (self *MemberInfo) Descriptor() string {
	return self.cp.GetUtf8(self.descriptorIndex)
}