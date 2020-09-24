package native

import "JVM-GO/ch09/rtda"

// 本地方法定义成一个函数，参数是Frame结构体指针，没有返回值
// frame参数就是本地方法的工作空间，也就是连接Java虚拟机和Java类库的桥梁
type NativeMethod func(frame *rtda.Frame)
// registry变量是个哈希表，值是具体的本地方法实现
var registry = map[string]NativeMethod{}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}

// 类名、方法名和方法描述符加在一起才能唯一确定一个方法，所以把它们的组合作为本地方法注册表的键
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

// 根据类名、方法名和方法描述符查找本地方法实现
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	// 如果找不到，则返回nil
	return nil
}
