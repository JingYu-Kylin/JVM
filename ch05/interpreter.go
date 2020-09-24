package main

import (
	"JVM-GO/ch05/classfile"
	"JVM-GO/ch05/instructions"
	"JVM-GO/ch05/instructions/base"
	"JVM-GO/ch05/rtda"
	"fmt"
)


func interpret(methodInfo *classfile.MemberInfo) {
	// interpret（）方法的参数是MemberInfo指针，调用MemberInfo结构体的CodeAttribute（）方法可以获取它的Code属性
	codeAttr := methodInfo.CodeAttribute()
	// 进一步获得执行方法所需的局部变量表和操作数栈空间，以及方法的字节码
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	// 先创建一个Thread实例，然后创建一个帧并把它推入Java虚拟机栈顶
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	// 最后执行方法
	defer catchErr(frame)
	loop(thread, bytecode)
}

/**
因为每个方法的最后一条指令都是某个return指令，而还没有实现return指令，所以方法在执行过程中必定会出现错误，此时解释器逻辑会转到catchErr（）函数
 */
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

/**
把局部变量表和操作数栈的内容打印出来，以此来观察方法的执行结果
 */
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
