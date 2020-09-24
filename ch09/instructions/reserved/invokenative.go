package reserved

import "JVM-GO/ch09/instructions/base"
import "JVM-GO/ch09/rtda"
import "JVM-GO/ch09/native"
// 如果没有任何包依赖lang包，它就不会被编译进可执行文件，上面的本地方法也就不会被注册。
// 所以需要一个地方导入lang包，把它放在invokenative.go文件中。
// 由于没有显示使用lang中的变量或函数，所以必须在包名前面加上下划线，否则无法通过编译。
// 这个技术在Go语言中叫作“import for side effect”。
import _ "JVM-GO/ch09/native/java/lang"
import _ "JVM-GO/ch09/native/sun/misc"

// Invoke native method 0xFE（invokenative）指令
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

/**
根据类名、方法名和方法描述符从本地方法注册表中查找本地方法实现，
如果找不到，则抛出UnsatisfiedLinkError异常，否则直接调用本地方法
 */
func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
