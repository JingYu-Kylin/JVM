package classfile

import (
	"fmt"
)

/**
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

/**
 * 把[]byte解析成ClassFile结构体
 */
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

/**
 * 依次调用其他方法解析class文件
 */
func (self *ClassFile) read(reader *ClassReader) {
	// 很多文件格式都会规定满足该格式的文件必须以某几个固定字节开头，这几个字节主要起标识作用，叫作魔数（magic number）。
	self.readAndCheckMagic(reader)
	// 魔数之后是class文件的次版本号和主版本号，都是u2类型
	self.readAndCheckVersion(reader)
	// 版本号之后是常量池
	self.constantPool = ReadConstantPool(reader)
	// 常量池之后是类访问标志，这是一个16位的“bitmask”，指出class文件定义的是类还是接口，访问级别是public还是private，等等
	self.accessFlags = reader.ReadUint16()
	// 类访问标志之后是两个u2类型的常量池索引，分别给出类名和超类名
	// class文件存储的类名类似完全限定名，但是把点换成了斜线，Java语言规范把这种名字叫作二进制名（binary names）
	// 因为每个类都有名字，所以thisClass必须是有效的常量池索引。
	// 除java.lang.Object之外，其他类都有超类，所以superClass只在Object.class中是0，在其他class文件中必须是有效的常量池索引
	self.thisClass = reader.ReadUint16()
	self.superClass = reader.ReadUint16()
	// 类和超类索引后面是接口索引表，表中存放的也是常量池索引，给出该类实现的所有接口的名字
	self.interfaces = reader.ReadUint16s()
	// 接口索引表之后是字段表和方法表，分别存储字段和方法信息
	// 字段和方法的基本结构大致相同，差别仅在于属性表
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = ReadAttributes(reader, self.constantPool) //见3.4
}

/**
 * 验证魔数
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.ReadUint32()
	// class文件的魔数是“0xCAFEBABE”
	if magic != 0xCAFEBABE {
		// Java虚拟机规范规定，如果加载的class文件不符合要求的格式，Java虚拟机实现就抛出java.lang.ClassFormatError异常
		// 暂时先调用panic（）方法终止程序执行
		panic("java.lang.ClassFormatError: magic!")
	}
}

/**
 * 验证版本
 */
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.ReadUint16()
	self.majorVersion = reader.ReadUint16()
	switch self.majorVersion {
	// 特定的Java虚拟机实现只能支持版本号在某个范围内的class文件
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	// 如果版本号不在支持的范围内，Java虚拟机实现就抛出java.lang.UnsupportedClassVersionError异常
	// 如果遇到其他版本号，暂时先调用panic（）方法终止程序执行。
	panic("java.lang.UnsupportedClassVersionError!")
}

/**
 * 这6个方法是Getter方法，把结构体的字段暴露给其他包使用
 */
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

/**
 * 从常量池查找类名
 */
func (self *ClassFile) ClassName() string {
	return self.constantPool.GetClassName(self.thisClass)
}

/**
 * 从常量池查找超类名
 */
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.GetClassName(self.superClass)
	}
	return "" // 只有java.lang.Object没有超类
}

/**
 * 从常量池查找接口名
 */
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.GetClassName(cpIndex)
	}
	return interfaceNames
}

