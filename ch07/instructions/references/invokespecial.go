package references

import (
	"JVM-GO/ch07/instructions/base"
	"JVM-GO/ch07/rtda"
	"JVM-GO/ch07/rtda/heap"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	// 先拿到当前类、当前常量池、方法符号引用，然后解析符号引用，拿到解析后的类和方法
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	// 假定从方法符号引用中解析出来的类是C，方法是M。
	// 如果M是构造函数，则声明M的类必须是C，否则抛出NoSuchMethodError异常。
	// 如果M是静态方法，则抛出IncompatibleClassChangeError异常。
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 从操作数栈中弹出this引用，如果该引用是null，抛出NullPointerException异常。
	// 注意，在传递参数之前，不能破坏操作数栈的状态。
	// 给OperandStack结构体添加一个GetRefFromTop（）方法，该方法返回距离操作数栈顶n个单元格的引用变量。
	// 比如GetRefFromTop（0）返回操作数栈顶引用，GetRefFromTop（1）返回从栈顶开始的倒数第二个引用，等等
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	// 上面的判断确保protected方法只能被声明该方法的类或子类调用。
	// 如果违反这一规定，则抛出IllegalAccessError异常
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}
	// 如果调用的中超类中的函数，但不是构造函数，且当前类的ACC_SUPER标志被设置，需要一个额外的过程查找最终要调用的方法；
	// 否则前面从方法符号引用中解析出来的方法就是要调用的方法。
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {

		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			methodRef.Name(), methodRef.Descriptor())
	}
	// 如果查找过程失败，或者找到的方法是抽象的，抛出AbstractMethodError异常。
	// 最后，如果一切正常，就调用方法。
	// 这里之所以这么复杂，是因为调用超类的（非构造函数）方法需要特别处理。
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
