package stores

import (
	"JVM-GO/ch05/instructions/base"
	"JVM-GO/ch05/rtda"
)

// Store long into local variable
/**
lstore指令的索引来自操作数
 */
type LSTORE struct{ base.Index8Instruction }
func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(self.Index))
}

/**
其余4条指令的索引隐含在操作码中，
 */
type LSTORE_0 struct{ base.NoOperandsInstruction }
func (self *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct{ base.NoOperandsInstruction }
func (self *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct{ base.NoOperandsInstruction }
func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct{ base.NoOperandsInstruction }
func (self *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}

/**
同样定义一个函数供5条指令使用
 */
func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}