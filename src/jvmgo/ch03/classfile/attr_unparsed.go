package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (e *UnparsedAttribute) readInfo(reader *ClassReader) {
	e.info = reader.readBytes(e.length)
}
