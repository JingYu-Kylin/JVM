package references

import (
	"JVM-GO/ch06/instructions/base"
	"JVM-GO/ch06/rtda"
	"JVM-GO/ch06/rtda/heap"
)

// Check whether object is of given type
type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	// 先从操作数栈中弹出对象引用，再推回去，这样就不会改变操作数栈的状态。
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	// 如果引用是null，则指令执行结束。也就是说，null引用可以转换成任何类型，
	if ref == nil {
		return
	}
	// 否则解析类符号引用
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	// 判断对象是否是类的实例。如果是的话，指令执行结束，否则抛出ClassCastException。
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}