package classfile

/*
	常量池小结：
		可以把常量池中的常量分为两类，字面量和符号引用。字面量包括数字常量和字符串常量，符号引用包括类和接口名、地段和方法信息等。
		除了字面量其他常量都是通过索引直接或间接指向 CONSTANT_Utf8_info 常量
*/

// ConstantPool 使用type关键字来定义一个类型，用Java的话讲，就是创建一个类，这个类可以是结构体（以struct开头，后跟大括号），可以是某种类型
type ConstantPool []ConstantInfo

// 常量池实际上也是一个表，但有以下三点需要注意...
// 1. 表头给出的常量池大小比实际上大 1
// 2. 有效的常量池索引是 1 ~ n-1 ， 0 是无效的索引，表示不指向任何常量
// 3. CONSTANT_LONG_info 和 CONSTANT_Double_info 各占两个位置。即如果常量池中存在这两种变量，实际的常量数量要比 n-1 要小，且 1 ~ n-1 中的某些数也会变成无效索引
// 这是一个函数（废话）
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	// 索引从 1 开始
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			// 这两种类型占了两个位置
			i++
		}
	}
	return cp
}

// 这是一个方法。方法本质上就是函数，但方法有接收者，接收者类型就是 func 关键字和函数名中间的小括号内容
// 按照索引查找常量
func (e ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	// 书上这里是 if cpInfo := e[index]; cpInfo != nil
	// 但这样写编辑器会报错，搜了一下，nil是针对空指针的类型，若是判断是否为空结构体应写为 cpInfo != struct {}{}
	// struct{} 是空结构体，而struct{}{} 是空结构体的一个字面量实现
	// 此处就先写成判空指针的形式
	if cpInfo := e[index]; &cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool")
}

// 从常量池查找字段或方法的名字和描述符
func (e ConstantPool) getNameAndType(index uint16) (string, string) {
	// 这里是一个断言
	ntInfo := e.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := e.getUtf8(ntInfo.nameIndex)
	_type := e.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 从常量池中查找类名
func (e ConstantPool) getClassName(index uint16) string {
	classInfo := e.getConstantInfo(index).(*ConstantClassInfo)
	return e.getUtf8(classInfo.nameIndex)
}

// 从常量池查找UTF-8字符串
func (e ConstantPool) getUtf8(index uint16) string {
	utf8Info := e.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
