package base

import (
	"JVM-GO/ch07/rtda"
	"JVM-GO/ch07/rtda/heap"
)

// jvms 5.5
func InitClass(thread *rtda.Thread, class *heap.Class) {
	// 先调用StartInit（）方法把类的initStarted状态设置成true以免进入死循环
	class.StartInit()
	// 然后调用scheduleClinit（）函数准备执行类的初始化方法
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
