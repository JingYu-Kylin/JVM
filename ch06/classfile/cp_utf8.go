package classfile

import (
	"fmt"
	"unicode/utf16"
)

/**
MUTF-8编码的字符串
CONSTANT_Utf8_info {
	u1 tag;
	u2 length;
	u1 bytes[length];
}
字符串在class文件中是以MUTF-8（Modified UTF-8）方式编码的
Java序列化机制也使用了MUTF-8编码。
java.io.DataInput和java.io.DataOutput接口分别定义了readUTF（）和writeUTF（）方法，可以读写MUTF-8编码的字符串
 */

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	// 先读取出[]byte
	length := uint32(reader.ReadUint16())
	bytes := reader.ReadBytes(length)
	// 解码成Go字符串
	self.str = decodeMUTF8(bytes)
}

/**
 * 简化版的readMUTF8（）
 * 假设字符串中不包含null字符或补充字符
 */
//func decodeMUTF8(bytes []byte) string {
//	return string(bytes)
//}

/**
 * mutf8 -> utf16 -> utf32 -> string
 * 根据java.io.DataInputStream.readUTF（）方法改写
 * 见 java.io.DataInputStream.readUTF(DataInput)
 */
func decodeMUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}

