package rtda

import "JVM-GO/ch07/rtda/heap"

/**
 * 线程
 */
type Thread struct {
	pc int //pc寄存器
	stack *Stack //Stack结构体（Java虚拟机栈）指针
}
func NewThread() *Thread {
	return &Thread{
		//newStack（）函数创建Stack结构体实例，它的参数表示要创建的Stack最多可以容纳多少帧
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int { return self.pc } // getter
func (self *Thread) SetPC(pc int) { self.pc = pc } // setter

/**
 * 调用Stack结构体
 */
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

/**
 * 返回当前帧
 */
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

