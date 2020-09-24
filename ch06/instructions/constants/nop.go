package constants

import (
	"JVM-GO/ch06/instructions/base"
	"JVM-GO/ch06/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }
func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么也不用做
}