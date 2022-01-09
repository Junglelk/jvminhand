package classfile

// ConstantClassInfo 类和超类索引，以及接口表中的接口索引都是指向 CONSTANT_Class_info 常量
type ConstantClassInfo struct {
	cp        *ConstantPool
	nameIndex uint16
}

func (e *ConstantClassInfo) readInfo(reader *ClassReader) {
	e.nameIndex = reader.readUint16()
}

func (e *ConstantClassInfo) Name() string {
	return e.cp.getUtf8(e.nameIndex)
}
