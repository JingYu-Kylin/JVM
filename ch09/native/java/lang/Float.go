package lang

import "math"
import "JVM-GO/ch09/native"
import "JVM-GO/ch09/rtda"

const jlFloat = "java/lang/Float"

func init() {
	native.Register(jlFloat, "floatToRawIntBits", "(F)I", floatToRawIntBits)
	native.Register(jlFloat, "intBitsToFloat", "(I)F", intBitsToFloat)
}

// public static native int floatToRawIntBits(float value);
// (F)I
// Go语言的math包提供了一个类似函数：Float32bits（），正好可以用来实现floatToRaw-IntBits（）方法
func floatToRawIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value) // todo
	frame.OperandStack().PushInt(int32(bits))
}

// public static native float intBitsToFloat(int bits);
// (I)F
func intBitsToFloat(frame *rtda.Frame) {
	bits := frame.LocalVars().GetInt(0)
	value := math.Float32frombits(uint32(bits)) // todo
	frame.OperandStack().PushFloat(value)
}
