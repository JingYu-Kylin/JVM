package references

import (
	"JVM-GO/ch06/instructions/base"
	"JVM-GO/ch06/rtda"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
