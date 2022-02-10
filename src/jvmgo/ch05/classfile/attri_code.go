package classfile

// CodeAttribute  是变长字段，只存在于 method_info 结构中。Code属性存放字节码等方法相关信息。
/*
	Code_attribute{
		u2 attribute_name_index;
		u4 attribute_length;
		u2 max_stack;
		u2 max_locals;
		u4 code_length;
		u1 code[code_length];
		u2 exception_table_length;
		{
			u2 start_pc;
			u2 end_pc;
			u2 handler_pc;
			u2 catch_type;
		}exception_table[exception_table_length]
		u2 attributes_count
		attribute_info attributes[attributes_count]
}
*/
// max_stack 给出操作数栈的最大深度，max_locals 给出局部变量表大小。字节码存在 u1 表中。最后是异常处理表和属性表。
// 在运行时数据区，实现操作数栈和局部变量表时，max_stack 和 max_locals 就会派上用场。
// 在指令集和解释器处，会用到字节码
// 异常处理时，会用到异常处理表
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (e *CodeAttribute) MaxStack() uint {
	return uint(e.maxStack)
}
func (e *CodeAttribute) MaxLocals() uint {
	return uint(e.maxLocals)
}
func (e *CodeAttribute) Code() []byte {
	return e.code
}

func (e *CodeAttribute) readInfo(reader *ClassReader) {
	e.maxStack = reader.readUint16()
	e.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	e.code = reader.readBytes(codeLength)
	e.exceptionTable = readExceptionTable(reader)
	e.attributes = readAttributes(reader, e.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
