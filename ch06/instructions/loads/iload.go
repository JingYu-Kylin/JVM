package loads

import (
	"JVM-GO/ch06/instructions/base"
	"JVM-GO/ch06/rtda"
)

// Load int from local variable
/**
iload指令的索引来自操作数
 */
type ILOAD struct{ base.Index8Instruction }
func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, self.Index)
}

/**
其余4条指令的索引隐含在操作码中
 */
type ILOAD_0 struct{ base.NoOperandsInstruction }
func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct{ base.NoOperandsInstruction }
func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }
func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }
func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

/**
为了避免重复代码，定义一个函数供iload系列指令使用
 */
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}