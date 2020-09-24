package classfile

import "encoding/binary"

/**
 * 数据在class文件中以大端（big-endian）方式存储
 * Java虚拟机规范定义了u1、u2和u4三种数据类型来表示1、2和4字节无符号整数，
 * 分别对应Go语言的uint8、uint16和uint32类型
 * 相同类型的多条数据一般按表（table）的形式存储在class文件中。
 * 表由表头和表项（item）构成，表头是u2或u4整数。
 * 假设表头是n，后面就紧跟着n个表项数据。
 */

type ClassReader struct {
	data []byte
}

/**
 * 读取u1类型数据
 */
func (self *ClassReader) ReadUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

/**
 * 读取u2类型数据
 */
func (self *ClassReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

/**
 * 读取u4类型数据
 */
func (self *ClassReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

/**
 * 读取uint64（Java虚拟机规范并没有定义u8）类型数据
 */
func (self *ClassReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

/**
 * 读取uint16表，表的大小由开头的uint16数据指出
 */
func (self *ClassReader) ReadUint16s() []uint16 {
	n := self.ReadUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.ReadUint16()
	}
	return s
}

/**
 * 用于读取指定数量的字节
 */
func (self *ClassReader) ReadBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
