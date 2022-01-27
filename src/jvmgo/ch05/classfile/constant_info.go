package classfile

// 常量池中存放的数据各不相同，所以每种常量的格式也不一致。
// 常量数据的第一个字节为 tag ，用来区分常量类型。
// java虚拟机中给出的常量结构为
// cp_info{
// 		u1 tag;
// 		u1 info[];
// }
// 下面是java中定义的常量的 tag 值
const (
	ConstantClass              = 7
	ConstantFieldRef           = 9
	ConstantMethodRef          = 10
	ConstantInterfaceMethodRef = 11
	ConstantString             = 8
	ConstantInteger            = 3
	ConstantFloat              = 4
	ConstantLong               = 5
	ConstantDouble             = 6
	ConstantNameAndType        = 12
	ConstantUtf8               = 1
	ConstantMethodHandle       = 15
	ConstantMethodtype         = 16
	ConstantInvokeDynamic      = 18
)

type ConstantInfo interface {
	// 读取常量信息，需要由集体的常量结构体来实现
	readInfo(reader *ClassReader)
}

// readConstantInfo 函数先读取出 tag 值，然后调用newConstantInfo() 函数创建具体的常量
// 然后调用常量的 readInfo() 方法读取常量信息
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	// 创建常量后，读取该常量的值
	c.readInfo(reader)
	return c
}

// newConstantInfo 根据tag值创建具体的常量
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {

	case ConstantInteger:
		return &ConstantIntegerInfo{}
	case ConstantFloat:
		return &ConstantFloatInfo{}
	case ConstantLong:
		return &ConstantLongInfo{}
	case ConstantDouble:
		return &ConstantDoubleInfo{}
	case ConstantUtf8:
		return &ConstantUtf8Info{}
	case ConstantString:
		return &ConstantStringInfo{}
	case ConstantClass:
		return &ConstantClassInfo{}
	case ConstantFieldRef:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case ConstantMethodRef:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case ConstantInterfaceMethodRef:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case ConstantNameAndType:
		return &ConstantNameAndTypeInfo{}
	case ConstantMethodtype:
		return &ConstantMethodTypeInfo{}
	case ConstantMethodHandle:
		return &ConstantMethodHandleInfo{}
	case ConstantInvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag !")
	}
}
