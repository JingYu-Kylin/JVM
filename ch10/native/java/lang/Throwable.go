package lang

import "fmt"
import "JVM-GO/ch10/native"
import "JVM-GO/ch10/rtda"
import "JVM-GO/ch10/rtda/heap"

const jlThrowable = "java/lang/Throwable"

/**
StackTraceElement结构体用来记录Java虚拟机栈帧信息：
lineNumber字段给出帧正在执行哪行代码；
methodName字段给出方法名；
className字段给出声明方法的类名；
fileName字段给出类所在的文件名。
 */
type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}

func init() {
	native.Register(jlThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
// (I)Ljava/lang/Throwable;
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	// 由于栈顶两帧正在执行fillInStackTrace（int）和fillInStackTrace（）方法，所以需要跳过这两帧。
	// 这两帧下面的几帧正在执行异常类的构造函数，所以也要跳过，具体要跳过多少帧数则要看异常类的继承层次。
	skip := distanceToObject(tObj.Class()) + 2
	// 计算好需要跳过的帧之后，调用Thread结构体的GetFrames（）方法拿到完整的Java虚拟机栈，然后reslice一下就是真正需要的帧。
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		// createStackTraceElement（）函数根据帧创建StackTraceElement实例
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

// distanceToObject（）函数计算所需跳过的帧数
func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}
