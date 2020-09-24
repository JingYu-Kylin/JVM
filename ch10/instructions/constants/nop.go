package constants

import "JVM-GO/ch10/instructions/base"
import "JVM-GO/ch10/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
