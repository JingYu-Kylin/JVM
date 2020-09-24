package control

import (
	"JVM-GO/ch06/instructions/base"
	"JVM-GO/ch06/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
