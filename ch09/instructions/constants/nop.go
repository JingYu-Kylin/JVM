package constants

import "JVM-GO/ch09/instructions/base"
import "JVM-GO/ch09/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
