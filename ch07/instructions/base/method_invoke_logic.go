package base

import (
	"JVM-GO/ch07/rtda"
	"JVM-GO/ch07/rtda/heap"
	"fmt"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	// 创建新的帧并推入Java虚拟机栈
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	// 传递参数
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	// hack!
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
