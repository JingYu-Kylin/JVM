package comparisons

import (
	"JVM-GO/ch07/instructions/base"
	"JVM-GO/ch07/rtda"
)

// Compare long
type LCMP struct{ base.NoOperandsInstruction }

func (self *LCMP) Execute(frame *rtda.Frame) {
	// 把栈顶的两个long变量弹出
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	//进行比较, 然后把比较结果（int型0、1或-1）推入栈顶，
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
