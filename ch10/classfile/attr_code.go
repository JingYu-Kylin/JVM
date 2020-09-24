package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}
func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}
func (self *CodeAttribute) Code() []byte {
	return self.code
}
func (self *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return self.exceptionTable
}

func (self *CodeAttribute) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}

type ExceptionTableEntry struct {
	// start_pc和end_pc可以锁定一部分字节码，这部分字节码对应某个可能抛出异常的try{}代码块
	startPc   uint16
	endPc     uint16
	// handler_pc就指出负责异常处理的catch{}块在哪里。
	handlerPc uint16
	// catch_type是个索引，通过它可以从运行时常量池中查到一个类符号引用，解析后的类是个异常类
	catchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}
