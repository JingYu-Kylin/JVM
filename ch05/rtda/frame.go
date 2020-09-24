package rtda

/**
 * 帧
 */
type Frame struct {
	lower *Frame //用来实现链表数据结构
	localVars LocalVars //局部变量表指针
	operandStack *OperandStack //操作数栈指针

	// 这两个字段主要是为了实现跳转指令而添加的
	thread       *Thread
	nextPC       int // the next instruction after the call
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
		// 执行方法所需的局部变量表大小和操作数栈深度是由编译器预先计算好的，存储在class文件method_info结构的Code属性中，
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
