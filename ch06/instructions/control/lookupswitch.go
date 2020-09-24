package control

import (
	"JVM-GO/ch06/instructions/base"
	"JVM-GO/ch06/rtda"
)

/*
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/
// Access jump table by key match and jump
type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// 先跳过padding
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	// matchOffsets有点像Map，它的key是case值，value是跳转偏移量
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	// 先从操作数栈中弹出一个int变量
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 {
		// 然后用它查找matchOffsets，看是否能找到匹配的key
		if self.matchOffsets[i] == key {
			// 如果能，则按照value给出的偏移量跳转
			offset := self.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	// 否则按照defaultOffset跳转
	base.Branch(frame, int(self.defaultOffset))
}
