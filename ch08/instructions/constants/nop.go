package constants

import (
	"JVM-GO/ch08/instructions/base"
	"JVM-GO/ch08/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
