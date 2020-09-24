package classfile

/**
Code是变长属性
只存在于method_info结构中。
Code属性中存放字节码等方法相关信息
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;// 操作数栈的最大深度
    u2 max_locals;// 局部变量表大小
	// 字节码，存在u1表中
    u4 code_length;
    u1 code[code_length];
	// 异常处理表
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
	// 属性表
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
	self.maxStack = reader.ReadUint16()
	self.maxLocals = reader.ReadUint16()
	codeLength := reader.ReadUint32()
	self.code = reader.ReadBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = ReadAttributes(reader, self.cp)
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


type ExceptionTableEntry struct {
	startPc uint16
	endPc uint16
	handlerPc uint16
	catchType uint16
}
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.ReadUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc: reader.ReadUint16(),
			endPc: reader.ReadUint16(),
			handlerPc: reader.ReadUint16(),
			catchType: reader.ReadUint16(),
		}
	}
	return exceptionTable
}
