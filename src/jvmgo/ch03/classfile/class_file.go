package classfile

// ClassFile Java虚拟机规范中接口表、字段表、方法表、属性表都有各自的计数器
// Go语言中的访问控制非常简单，只有公开和私有两种。所有首字母大写的类型、结构体、字段、变量、函数、方法等等都是公开的，
// 首字母小写的都是私有的。
type ClassFile struct {
	// 魔数
	magic uint32
	// 副版本号
	minorVersion uint16
	// 主版本号
	majorVersion uint16
	// 常量池
	constantPool ConstantPool
	// 访问标记
	accessFlags uint16
	// 类索引
	thisClass uint16
	// 父类索引
	superClass uint16
	// 接口表
	interfaces []uint16
	// 字段表
	fields []*MemberInfo
	// 方法表
	method []*MemberInfo
	// 属性表
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {

}

func (e *ClassFile) read(reader *ClassReader) {

}

func (e *ClassFile) readAndCheckMagic(reader *ClassReader) {

}

func (e *ClassFile) readAndCheckVersion(reader *ClassReader) {

}

// MinorVersion getter
func (e *ClassFile) MinorVersion() uint16 {

}

// MajorVersion getter
func (e *ClassFile) MajorVersion() uint16 {

}

// ConstantPool getter
func (e *ClassFile) ConstantPool() ConstantPool {

}

// AccessFlags getter
func (e *ClassFile) AccessFlags() uint16 {

}

// Fields getter
func (e *ClassFile) Fields() []MemberInfo {

}

// Methods getter
func (e *ClassFile) Methods() []MemberInfo {

}

func (e *ClassFile) ClassName() string {

}

func (e *ClassFile) SuperClassName() string {

}

func (e *ClassFile) InterfaceName() []string {

}
