package references

import (
	"JVM-GO/ch07/instructions/base"
	"JVM-GO/ch07/rtda"
	"JVM-GO/ch07/rtda/heap"
)

// Create new object
type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	// todo: init class
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	// 因为接口和抽象类都不能实例化，所以如果解析后的类是接口或抽象类，
	// 按照Java虚拟机规范规定，需要抛出InstantiationError异常。
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	// 如果解析后的类还没有初始化，则需要先初始化类。
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}

