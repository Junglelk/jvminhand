package classfile

// ConstantStringInfo 结构体本身不存储字符串数据，只存了常量池索引
// 这正好也对应java语言中 String 类是常量的设计方式
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

// 获取常量池索引
func (e ConstantStringInfo) readInfo(reader *ClassReader) {
	e.stringIndex = reader.readUint16()
}

func (e *ConstantStringInfo) String() string {
	return e.cp.getUtf8(e.stringIndex)
}
