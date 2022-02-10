package base

// BytecodeReader code字段存放字节码，pc字段记录读取到了哪个字节。
type BytecodeReader struct {
	code []byte
	pc   int
}

func (e *BytecodeReader) Reset(code []byte, pc int) {
	e.code = code
	e.pc = pc
}

func (e *BytecodeReader) ReadUint8() uint8 {
	i := e.code[e.pc]
	e.pc++
	return i
}

func (e *BytecodeReader) ReadInt8() int8 {
	return int8(e.ReadUint8())
}

func (e *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(e.code[e.ReadUint8()])
	byte2 := uint16(e.code[e.ReadUint8()])
	return (byte1 << 8) | byte2
}

func (e *BytecodeReader) ReadInt16() int16 {
	return int16(e.ReadUint16())
}

func (e *BytecodeReader) ReadUint32() uint32 {
	byte1 := uint32(e.code[e.ReadUint8()])
	byte2 := uint32(e.code[e.ReadUint8()])
	byte3 := uint32(e.code[e.ReadUint8()])
	byte4 := uint32(e.code[e.ReadUint8()])
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (e *BytecodeReader) ReadInt32() int32 {
	return int32(e.ReadUint32())
}

func (e *BytecodeReader) SkipPadding() {
	for e.pc%4 != 0 {
		e.ReadInt8()
	}
}

func (e *BytecodeReader) ReadUint32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = e.ReadInt32()
	}
	return ints
}
