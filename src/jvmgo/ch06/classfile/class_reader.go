package classfile

import "encoding/binary"

type ClassReader struct {
	// 在go中，byte为8比特无符号整数
	data []byte
}

// Java虚拟机规范定义了u1、u2、u4三种数据类型来表示1、2和4字节无符号整数
// 相同类型的多条数据一般按照表（table）的形式存储在class文件中。表由表头和表项（item）构成，表头是u2或u4的整数，假设表头是 n ，后面就紧跟 n 个表项数据
func (e *ClassReader) readUint8() uint8 {
	// 取出并返回一个字节
	val := e.data[0]
	e.data = e.data[1:]
	return val
}

// 使用uint16读取u2类型的数据
// BigEndian 是ByteOrder的大端实现
func (e *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(e.data)
	e.data = e.data[2:]
	return val
}

func (e *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(e.data)
	e.data = e.data[4:]
	return val
}

// 读取uint64 虽然Java虚拟机规范并没有定义u8
func (e *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(e.data)
	e.data = e.data[8:]
	return val
}

func (e *ClassReader) readUint16s() []uint16 {
	n := e.readUint16()
	// make 用于初始化slice, map 或 chan
	// 这个的意思是，新建一个uint16的切片，长度和容量均为 n
	s := make([]uint16, n)
	for i := range s {
		s[i] = e.readUint16()
	}
	return s
}

// 用于读取指定数量的字节
func (e *ClassReader) readBytes(n uint32) []byte {
	// 从下标 0 到下标 n
	bytes := e.data[:n]
	// 从第 n 个到最后
	e.data = e.data[n:]
	return bytes
}
