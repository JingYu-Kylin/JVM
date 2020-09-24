package stack

import (
	"JVM-GO/ch07/instructions/base"
	"JVM-GO/ch07/rtda"
)
/**
pop指令只能用于弹出int、float等占用一个操作数栈位置的变量。
double和long变量在操作数栈中占据两个位置，需要使用pop2指令弹出
 */
type POP struct{ base.NoOperandsInstruction }
type POP2 struct{ base.NoOperandsInstruction }

/**
pop指令把栈顶变量弹出
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
 */
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

/**
pop2指令把栈顶变量弹出
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
 */
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}