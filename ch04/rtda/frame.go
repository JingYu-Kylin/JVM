package rtda

/**
 * 帧
 */
type Frame struct {
	lower *Frame //用来实现链表数据结构
	localVars LocalVars //局部变量表指针
	operandStack *OperandStack //操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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
