package heap

import (
	"JVM-GO/ch06/classfile"
	"fmt"
)

type Constant interface{}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

/**
把class文件中的常量池转换成运行时常量池
 */
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		// int或float型常量，直接取出常量值，放进consts中即可
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value() // int32
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()  // float32
		// long或double型常量，也是直接提取常量值放进consts中。
		// 但是要注意，这两种类型的常量在常量池中都是占据两个位
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()  // int64
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value() // float64
			i++
		// 字符串常量，直接取出Go语言字符串，放进consts中
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()
		// 剩下4种类型的常量需要处理，分别是类、字段、方法和接口方法的符号引用
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *classfile.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		default:
			// todo
		}
	}

	return rtCp
}

/**
根据索引返回常量
 */
func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
