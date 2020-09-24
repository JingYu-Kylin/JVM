package rtda

import "JVM-GO/ch07/rtda/heap"

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
	// 为了通过frame变量拿到当前类的运行时常量池
	method       *heap.Method
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
		// 执行方法所需的局部变量表大小和操作数栈深度是由编译器预先计算好的，存储在class文件method_info结构的Code属性中，
	}
}

// getters & setters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) Method() *heap.Method {
	return self.method
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}