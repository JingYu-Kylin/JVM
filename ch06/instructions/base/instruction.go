package base

import "JVM-GO/ch06/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) // 从字节码中提取操作数
	Execute(frame *rtda.Frame) // 执行指令逻辑
}

/**
表示没有操作数的指令，所以没有定义任何字段
 */
type NoOperandsInstruction struct {
}
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// 什么也不用读
}

/**
跳转指令
 */
type BranchInstruction struct {
	Offset int // 跳转偏移量
}
/**
从字节码中读取一个uint16整数，转成int后赋给Offset字段
 */
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/**
存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出。
把这类指令抽象成Index8Instruction结构体
 */
type Index8Instruction struct {
	Index uint // 局部变量表索引
}
/**
从字节码中读取一个int8整数，转成uint后赋给Index字段
 */
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

/**
有一些指令需要访问运行时常量池，常量池索引由两字节操作数给出。
把这类指令抽象成Index16Instruction结构体
 */
type Index16Instruction struct {
	Index uint // 常量池索引
}
/**
从字节码中读取一个uint16整数，转成uint后赋给Index字段
 */
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
