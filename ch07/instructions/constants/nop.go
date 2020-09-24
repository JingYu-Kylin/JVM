package constants

import (
	"JVM-GO/ch07/instructions/base"
	"JVM-GO/ch07/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }
func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么也不用做
}