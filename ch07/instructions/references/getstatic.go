package references

import (
	"JVM-GO/ch07/instructions/base"
	"JVM-GO/ch07/rtda"
	"JVM-GO/ch07/rtda/heap"
)

// Get static field from class
type GET_STATIC struct{ base.Index16Instruction }

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	// todo: init class 如果声明字段的类还没有初始化好，也需要先初始化
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 如果解析后的字段不是静态字段，也要抛出IncompatibleClassChangeError异常
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// getstatic只是读取静态变量的值，自然也就不用管它是否是final了

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	// 根据字段类型，从静态变量中取出相应的值，然后推入操作数栈顶
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		// todo
	}
}
