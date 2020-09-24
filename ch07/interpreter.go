package main

import (
	"JVM-GO/ch07/instructions"
	"JVM-GO/ch07/instructions/base"
	"JVM-GO/ch07/rtda"
	"JVM-GO/ch07/rtda/heap"
	"fmt"
)


func interpret(method *heap.Method, logInst bool) {
	// 先创建一个Thread实例，然后创建一个帧并把它推入Java虚拟机栈顶
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	// 最后执行方法
	// logInst参数控制是否把指令执行信息打印到控制台
	defer catchErr(thread)
	loop(thread, logInst)
}

/**
打印出错信息
因为每个方法的最后一条指令都是某个return指令，而还没有实现return指令，所以方法在执行过程中必定会出现错误，此时解释器逻辑会转到catchErr（）函数
 */
func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

/**
把局部变量表和操作数栈的内容打印出来，以此来观察方法的执行结果
在每次循环开始，先拿到当前帧，然后根据pc从当前方法中解码出一条指令。
指令执行完毕之后，判断Java虚拟机栈中是否还有帧。
如果没有则退出循环；否则继续。
 */
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

/**
在方法执行过程中打印指令信息
 */
func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

/**
打印Java虚拟机栈信息
 */
func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}