package references

import (
	"JVM-GO/ch08/instructions/base"
	"JVM-GO/ch08/rtda"
	"JVM-GO/ch08/rtda/heap"
)

const (
	//Array Type  atype常量
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

// Create new array
// newarray指令需要两个操作数。第一个操作数是一个uint8整数，在字节码中紧跟在指令操作码后面，表示要创建哪种类型的数组。
// Java虚拟机规范把这个操作数叫作atype，并且规定了它的有效值。
type NEW_ARRAY struct {
	atype uint8
}

// 读取atype的值
func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype = reader.ReadUint8()
}
// 根据atype和count创建基本类型数组
func (self *NEW_ARRAY) Execute(frame *rtda.Frame) {
	// newarray指令的第二个操作数是count，从操作数栈中弹出，表示数组长度
	stack := frame.OperandStack()
	count := stack.PopInt()
	// 如果count小于0，则抛出NegativeArraySizeException异常
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	// 否则根据atype值使用当前类的类加载器加载数组类，然后创建数组对象并推入操作数栈
	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, self.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}
