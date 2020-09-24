package math

import (
	"JVM-GO/ch07/instructions/base"
	"JVM-GO/ch07/rtda"
)

// Increment local variable by constant
type IINC struct {
	Index uint
	Const int32
}

/**
从字节码里读取操作数
 */
func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

/**
从局部变量表中读取变量，给它加上常量值，再把结果写回局部变量表
 */
func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}