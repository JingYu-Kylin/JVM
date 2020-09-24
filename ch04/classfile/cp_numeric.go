package classfile

import (
	"math"
)

/**
int型常量
CONSTANT_Integer_info {
	u1 tag;
	u4 bytes;
}
CONSTANT_Integer_info正好可以容纳一个Java的int型常量，
但实际上比int更小的boolean、byte、short和char类型的常量也放在CONSTANT_Integer_info中
 */
type ConstantIntegerInfo struct {
	val int32
}
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	// 先读取一个uint32数据
	bytes := reader.ReadUint32()
	// 然后把它转型成int32类型
	self.val = int32(bytes)
}

/**
IEEE754单精度浮点数常量
CONSTANT_Float_info {
	u1 tag;
	u4 bytes;
}
 */
type ConstantFloatInfo struct {
	val float32
}
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	// 读取一个uint32数据
	bytes := reader.ReadUint32()
	// 转换成float32类型
	self.val = math.Float32frombits(bytes)
}

/**
Long型常量
CONSTANT_Long_info {
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}
 */
type ConstantLongInfo struct {
	val int64
}
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.ReadUint64()
	self.val = int64(bytes)
}

/**
IEEE754双精度浮点数
CONSTANT_Double_info {
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}
 */
type ConstantDoubleInfo struct {
	val float64
}
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.ReadUint64()
	self.val = math.Float64frombits(bytes)
}


