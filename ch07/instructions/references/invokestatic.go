package references

import (
	"JVM-GO/ch07/instructions/base"
	"JVM-GO/ch07/rtda"
	"JVM-GO/ch07/rtda/heap"
)

// Invoke a class (static) method
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	// 假定解析符号引用后得到方法M
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	// M必须是静态方法，否则抛出Incompatible-ClassChangeError异常
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := resolvedMethod.Class()
	// M不能是类初始化方法
	// 类初始化方法只能由Java虚拟机调用，不能使用invokestatic指令调用。
	if !class.InitStarted() {
		frame.RevertNextPC()
		// 如果声明M的类还没有被初始化，则要先初始化该类
		base.InitClass(frame.Thread(), class)
		return
	}
	// 对于invokestatic指令，M就是最终要执行的方法，调用InvokeMethod（）函数执行该方法。
	base.InvokeMethod(frame, resolvedMethod)
}