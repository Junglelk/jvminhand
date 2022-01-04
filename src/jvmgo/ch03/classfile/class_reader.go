package classfile

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

func (e *ClassReader) readUint16() uint16 {

}

func (e *ClassReader) readUint32() uint32 {

}

func (e *ClassReader) readUint64() uint64 {

}

func (e *ClassReader) readUint16s() []uint16 {

}

func (e *ClassReader) readBytes(length uint32) []byte {

}
