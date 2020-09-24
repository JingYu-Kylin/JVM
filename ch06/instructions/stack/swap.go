package stack

import (
	"JVM-GO/ch06/instructions/base"
	"JVM-GO/ch06/rtda"
)

// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

/**
swap指令交换栈顶的两个变量
bottom -> top
[...][c][b][a]
          \/
          /\
         V  V
[...][c][a][b]
 */
func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
