package classfile

/*
	之前整体构建了class文件和常量池，现在的问题是，方法的实体在哪里？
	或者说方法的字节码存在什么地方？属性表。
	常量类型是由虚拟机严格规定的，但属性不是，所以虚拟机实现者可以自定义自己的属性类型。
	所以，Java虚拟机不依据tag区分属性，而是使用属性名来区分不同的属性。属性数据放在属性名之后的 u1 表中。
	这样Java虚拟机就可以跳过自己无法识别的属性。
*/

// AttributeInfo
/*
	attribute_info {
		u2 attribute_name_index;
		u4 attribute_length;
		u1 info[attribute_length];
	}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// readAttributes 函数读取属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// readAttribute 读取单个属性
// 刚刚注意到这里有两种传递参数的方式，分别是 ClassReader 的指针传递和 ConstantPool 的对象传递
// 区别在于，指针指向原对象，任何对指针的修改都是在修改原对象，而对象传递的是原对象的副本，彼此独立
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

// readAttribute() 先读取属性名索引，根据它从常量池中找到属性名，然后读取属性长度，接着调用newAttributeInfo()函数创建具体的属性实例。
// newAttributeInfo Java虚拟机内置了 23 种属性，先解析其中的 8 种
// 23 种预定义属性可以分为三组：第一组是实现Java虚拟机必须的，共有 5 种；第二组属性是Java类库所必须的，共有 12 种；第三组属性主要是提供给工具使用的，共有 6 种。
// 第三组属于可选属性，即可以不出现在class文件中。如果 class 文件中存在第三组属性，则Java虚拟机和Java类库也可以使用，比如使用LineNumberTable属性在异常堆栈中显示行号（难道这是异常信息有行号的原因？）
// TODO 后续可能会自己写完剩下 15 个属性
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}

	}
}
