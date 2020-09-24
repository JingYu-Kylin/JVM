package classfile

/**
字段或方法的名称和描述符
CONSTANT_NameAndType_info {
	u1 tag;
	u2 name_index;
	u2 descriptor_index;
}
字段或方法名由name_index给出
字段或方法的描述符由descriptor_index给出
name_index和descriptor_index都是常量池索引，指向CONSTANT_Utf8_info常量
字段和方法名就是代码中出现的（或者编译器生成的）字段或方法的名字。
 */

/**
Java虚拟机规范定义了一种简单的语法来描述字段和方法
1）类型描述符
	①基本类型byte、short、char、int、long、float和double的描述符是单个字母，分别对应B、S、C、I、J、F和D。注意，long的描述符是J而不是L。
	②引用类型的描述符是L＋类的完全限定名＋分号。
	③数组类型的描述符是[＋数组元素类型描述符。
2）字段描述符就是字段类型的描述符。
3）方法描述符是（分号分隔的参数类型描述符）+返回值类型描述符，其中void返回值由单个字母V表示
更详细的介绍可以参考Java虚拟机规范
 */

type ConstantNameAndTypeInfo struct {
	nameIndex uint16
	descriptorIndex uint16
}
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.ReadUint16()
	self.descriptorIndex = reader.ReadUint16()
}
