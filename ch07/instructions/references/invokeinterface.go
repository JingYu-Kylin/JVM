package references

import (
	"JVM-GO/ch07/instructions/base"
	"JVM-GO/ch07/rtda"
	"JVM-GO/ch07/rtda/heap"
)

/**
注意，和其他三条方法调用指令略有不同，在字节码中，invokeinterface指令的操作码后面跟着4字节而非2字节。
前两字节的含义和其他指令相同，是个uint16运行时常量池索引。
第3字节的值是给方法传递参数需要的slot数，其含义和给Method结构体定义的argSlotCount字段相同。
正如我们所知，这个数是可以根据方法描述符计算出来的，它的存在仅仅是因为历史原因。
第4字节是留给Oracle的某些Java虚拟机实现用的，它的值必须是0。
该字节的存在是为了保证Java虚拟机可以向后兼容。
 */
// Invoke interface method
type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	// 先从运行时常量池中拿到并解析接口方法符号引用，如果解析后的方法是静态方法或私有方法，则抛出IncompatibleClassChangeError异常。
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 从操作数栈中弹出this引用，如果引用是null，则抛出NullPointerException异常。
	// 如果引用所指对象的类没有实现解析出来的接口，则抛出IncompatibleClassChangeError异常。
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException") // todo
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 查找最终要调用的方法。如果找不到，或者找到的方法是抽象的，则抛出Abstract-MethodError异常。
	// 如果找到的方法不是public，则抛出IllegalAccessError异常，
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	// 否则，一切正常，调用方法
	base.InvokeMethod(frame, methodToBeInvoked)
}
