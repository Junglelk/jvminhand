package classfile

// LocalVariableTableAttribute 属性表里存的方法的局部变量信息。LineNumberTable、LocalVariableTable和前面的SourceFile属性都是调试信息，都不是运行时必须的。
/*
	LocalVariableTable_attribute {
    	u2 attribute_name_index;
    	u4 attribute_length;
    	u2 local_variable_table_length;
    	{   u2 start_pc;
    	    u2 length;
    	    u2 name_index;
    	    u2 descriptor_index;
    	    u2 index;
    	} local_variable_table[local_variable_table_length];
	}
*/
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (e *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	e.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range e.localVariableTable {
		e.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
