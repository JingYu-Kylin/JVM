package lang

import "JVM-GO/ch09/native"
import "JVM-GO/ch09/rtda"
import "JVM-GO/ch09/rtda/heap"

const jlSystem = "java/lang/System"

func init() {
	native.Register(jlSystem, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
func arraycopy(frame *rtda.Frame) {
	// 先从局部变量表中拿到5个参数
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)
	// 源数组和目标数组都不能是null，否则按照System类的Javadoc应该抛出NullPointerException异常
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}
	// 源数组和目标数组必须兼容才能拷贝，否则应该抛出ArrayStoreExceptio异常
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}
	// 接下来检查srcPos、destPos和length参数，如果有问题则抛出IndexOutOfBoundsException异常
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}
	// 最后，参数合法，调用ArrayCopy（）函数进行数组拷贝
	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()
	// 首先确保src和dest都是数组
	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	// 然后检查数组类型。如果两者都是引用数组，则可以拷贝，否则两者必须是相同类型的基本类型数组
	if srcClass.ComponentClass().IsPrimitive() ||
		destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}
