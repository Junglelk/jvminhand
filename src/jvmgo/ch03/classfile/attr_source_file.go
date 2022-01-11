package classfile

// SourceFileAttribute SourceFile 是可选定长属性，只会出现在ClassFile结构中，
// 用于指出源文件名
/*
	SourceFile_attribute{
		u2 attribute_name_index;
		u4 attribute_length;
		u2 sourcefile_index;
	}
*/
// attribute_length 的值必须是 2 (Java虚拟机规范规定).sourcefile_index 是常量池索引，指向 CONSTANT_Utf8_info 常量
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (e *SourceFileAttribute) readInfo(reader *ClassReader) {
	e.sourceFileIndex = reader.readUint16()
}

func (e *SourceFileAttribute) FileName() string {
	return e.cp.getUtf8(e.sourceFileIndex)
}
