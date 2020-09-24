package references

import (
	"JVM-GO/ch06/instructions/base"
	"JVM-GO/ch06/rtda"
	"JVM-GO/ch06/rtda/heap"
)

// Fetch field from object
type GET_FIELD struct{ base.Index16Instruction }

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	// 字段符号引用解析
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 弹出对象引用，如果是null，则抛出NullPointerException
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	// 根据字段类型，获取相应的实例变量值，然后推入操作数栈
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()

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