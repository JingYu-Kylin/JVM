package heap

import "JVM-GO/ch10/classfile"

type ExceptionTable []*ExceptionHandler // ExceptionTable只是[]*ExceptionHandler的别名

type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlerPc int
	catchType *ClassRef
}

// 把class文件中的异常处理表转换成ExceptionTable类型
func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp), //异常处理项的catchType有可能是0
		}
	}

	return table
}

// 从运行时常量池中查找类符号引用
func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		return nil // catch all
	}
	return cp.GetConstant(index).(*ClassRef)
}

// 搜索异常处理表
func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self {
		// jvms: The start_pc is inclusive and end_pc is exclusive
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}

	// 第一，startPc给出的是try{}语句块的第一条指令，endPc给出的则是try{}语句块的下一条指令。
	// 第二，如果catchType是nil（在class文件中是0），表示可以处理所有异常，这是用来实现finally子句的。
	return nil
}
