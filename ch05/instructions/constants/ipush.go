package constants

import (
	"JVM-GO/ch05/instructions/base"
	"JVM-GO/ch05/rtda"
)

/**
bipush指令从操作数中获取一个byte型整数，扩展成int型，然后推入栈顶
 */
type BIPUSH struct { val int8 } // Push byte
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

/**
sipush指令从操作数中获取一个short型整数，扩展成int型，然后推入栈顶
 */
type SIPUSH struct { val int16 } // Push short
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}