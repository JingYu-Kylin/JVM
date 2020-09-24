package references

import (
	"JVM-GO/ch06/instructions/base"
	"JVM-GO/ch06/rtda"
	"JVM-GO/ch06/rtda/heap"
)

// Determine if object is of given type
type INSTANCE_OF struct{ base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	// 先弹出对象引用，如果是null，则把0推入操作数栈
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	// 如果对象引用不是null，则解析类符号引用，判断对象是否是类的实例，然后把判断结果推入操作数栈
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}