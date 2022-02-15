package classfile

import (
	"fmt"
	"unicode/utf16"
)

// ConstantUtf8Info 字符串在class文件中是以 MUTF-8 方式编码的，虽然编码和 UTF-8 大致相同，但并不兼容
// 区别在于：
// 1. null字符会被编码成 2 字节
// 2. 补充字符按照 UTF-16 拆分为代理对分别编码
type ConstantUtf8Info struct {
	str string
}

// readInfo 方法先读取出 []byte ，然后调用 decodeMUTF8() 函数把它解码成 Go 字符串。
func (e *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	e.str = decodeMUTF8(bytes)
}

// 这个是完整版的 decodeMUTF8 函数，反正我看不懂，简化版的只有一行 return string(bytes)
func decodeMUTF8(bytes []byte) string {
	length := len(bytes)
	charArr := make([]uint16, length)

	var c, char2, char3 uint16
	count := 0
	charArrCount := 0

	for count < length {
		c = uint16(bytes[count])
		if c > 127 {
			break
		}
		count++
		charArr[charArrCount] = c
		charArrCount++
	}

	for count < length {
		c = uint16(bytes[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			charArr[charArrCount] = c
			charArrCount++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > length {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			charArr[charArrCount] = c&0x1F<<6 | char2&0x3F
			charArrCount++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > length {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-2])
			char3 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count-1))
			}
			charArr[charArrCount] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			charArrCount++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than length
	charArr = charArr[0:charArrCount]
	runes := utf16.Decode(charArr)
	return string(runes)
}
