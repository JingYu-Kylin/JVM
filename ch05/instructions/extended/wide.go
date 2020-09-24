package extended

import (
	"JVM-GO/ch05/instructions/base"
	"JVM-GO/ch05/instructions/loads"
	"JVM-GO/ch05/instructions/math"
	"JVM-GO/ch05/instructions/stores"
	"JVM-GO/ch05/rtda"
)

// Extend local variable index by additional bytes
type WIDE struct {
	// wide指令改变其他指令的行为，modifiedInstruction字段存放被改变的指令。
	modifiedInstruction base.Instruction
}
/**
wide指令需要自己解码出modifiedInstruction
 */
func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	// 从字节码中读取一字节的操作码，然后创建子指令实例，最后读取子指令的操作数。
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		// 加载指令和存储指令都只有一个操作数，需要扩展成2字节
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x16:
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x17:
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x18:
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x19:
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x36:
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x37:
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x38:
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x39:
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x3a:
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x84:
		// iinc指令有两个操作数，都需要扩展成2字节
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0xa9: // 因为没有实现ret指令，所以暂时调用panic（）函数终止程序执行
		panic("Unsupported opcode: 0xa9!")
	}
}

/**
wide指令只是增加了索引宽度，并不改变子指令操作，所以其Execute（）方法只要调用子指令的Execute（）方法即可
 */
func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}
