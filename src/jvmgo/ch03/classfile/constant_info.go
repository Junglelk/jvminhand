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
