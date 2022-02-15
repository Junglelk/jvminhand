package classfile

// CONSTANT_Fieldref_info表示字段符号引用，CONSTANT_Methodref_info 表示普通（非接口）方法符号引用，
// CONSTANT_InterfaceMethodref_info 表示接口方法符号引用，这三个结构完全一致
/*
  	CONSTANT_Fieldref_info{
		u1 tag;
		u2 class_index;
		y2 name_and_type_index;
	}
*/

// ConstantMemberrefInfo 此处先定义一个统一的结构体ConstantMemberrefInfo
//  class_index 和 name_and_type_index 都是常量池索引，分别指向 CONSTANT_Class_info 和 CONSTANT_NameAndType_info 常量。
// 由于go没有继承的概念，而如果一样的代码复制三份则过于丑陋，所以使用组合去模拟继承
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (e *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	e.classIndex = reader.readUint16()
	e.nameAndTypeIndex = reader.readUint16()
}

func (e *ConstantMemberrefInfo) ClassName() string {
	return e.cp.getClassName(e.classIndex)
}

func (e ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return e.cp.getNameAndType(e.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
