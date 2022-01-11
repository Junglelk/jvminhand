package classfile

// ConstantValueAttribute 是定长属性，只会出现在 field_info 中，用于表示常量表达式的值。
/*
	ConstantValue_attribute{
		u2 attribute_name_index;
		u4 attribute_length;
		u2 constantvalue_index;
	}
*/
// attribute_length 的值必须为 2 (Java虚拟机规范规定)；
// constantvalue_index 是常量池索引，但具体指向哪种常量因字段类型而异
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (e *ConstantValueAttribute) readInfo(reader *ClassReader) {
	e.constantValueIndex = reader.readUint16()
}

func (e *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return e.constantValueIndex
}
