package classfile

import "fmt"

// ClassFile Java虚拟机规范中接口表、字段表、方法表、属性表都有各自的计数器
// All problems in computer science can be solved by another level of indirection.
// -- David Wheeler
// ClassFile 就是为了实现类加载而增加的中间层
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
	methods []*MemberInfo
	// 属性表
	attributes []AttributeInfo
}

// Parse 函数把[]byte解析成ClassFile结构体
// 函数的返回类型可以有变量名，这两个变量名可以作为已经声明的指针在函数体内使用
func Parse(classData []byte) (cf *ClassFile, err error) {
	// 前面说过，defer修饰的函数是在函数结束后执行；recover()函数用于处理panic，勉强相当于try-catch
	// 这段代码用于处理解析过程中的异常
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			// type assertion ，用于判断接口的实现类型是否为 T
			// x.(T) : asserts that x is not nil and that the value stored in x is of type T
			// 可用一个参数或两个参数接收，一个参数接受时，断言失败会panic，两个参数接受时，断言失败第二个会被赋值为false
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	// return 语句后面为空，一种情况是结束函数运行，函数无返回值；
	// 另一种情况是函数有返回值，返回值有变量名，此变量已经在函数体内使用，所以return结束后，该变量就是返回值，可省略
	return
}

// 读取并验证一个类文件
func (e *ClassFile) read(reader *ClassReader) {
	// 读取并验证魔数
	e.readAndCheckMagic(reader)
	// 读取并验证版本
	e.readAndCheckVersion(reader)
	// 读取常量池
	e.constantPool = readConstantPool(reader)
	// 获取访问修饰符
	e.accessFlags = reader.readUint16()
	// 当前类索引
	e.thisClass = reader.readUint16()
	// 父类索引
	e.superClass = reader.readUint16()
	// 接口表
	e.interfaces = reader.readUint16s()
	// 字段表
	e.fields = readMembers(reader, e.constantPool)
	// 方法表
	e.methods = readMembers(reader, e.constantPool)
	// 属性表
	e.attributes = readAttributes(reader, e.constantPool)

}

// class文件以16进制数字 CAFEBABE 开头，一共 8 个字节，所以使用 readUint32 函数读取
func (e *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	// 校验魔数是否正确
	if magic != 0xCAFEBABE {
		// java虚拟机规范中规定如果魔数不符合，则抛出“java.lang.ClassFormatError”异常
		// 但异常本身需要java虚拟机支持，目前未编写到这一步，先使用panic来终止程序运行
		panic("java.lang.ClassFormatError:magic!")
	}
}

// 读取并验证版本号，java的版本号分为次版本号和主版本号，都是u2类型，如果主版本号为M，次版本号为m，则完整版本号可以写为 M.m
// 次版本号JDK 1.2 （含）之后就没有再使用过了，版本号在 JDK 1.0.2 时为 45.0~45.3 , JDK 1.1 为45.0 ~ 45.65535，从 JDK 1.2 开始，版本号从 46.0 递增，Java8 为 52.0
func (e *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// 目前仅解析到Java8
	e.minorVersion = reader.readUint16()
	e.majorVersion = reader.readUint16()
	switch e.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if e.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersion")
}

// MinorVersion getter
func (e *ClassFile) MinorVersion() uint16 {
	return e.minorVersion
}

// MajorVersion getter
func (e *ClassFile) MajorVersion() uint16 {
	return e.majorVersion
}

// ConstantPool getter
// 版本号之后是常量池
func (e *ClassFile) ConstantPool() ConstantPool {
	return e.constantPool
}

// AccessFlags getter
// 常量池之后是类访问标记，访问标识符是一个 16 位的 bitmask，指出class文件定义的是类还是接口，访问级别是 public 还是 private 等。
func (e *ClassFile) AccessFlags() uint16 {
	return e.accessFlags
}

// Fields getter
func (e *ClassFile) Fields() []*MemberInfo {
	return e.fields
}

// Methods getter
func (e *ClassFile) Methods() []*MemberInfo {
	return e.methods
}

// ClassName 函数从常量池中查找类名
func (e *ClassFile) ClassName() string {
	return e.constantPool.getClassName(e.thisClass)
}

// SuperClassName 函数从常量池中查找超类名
func (e *ClassFile) SuperClassName() string {
	if e.superClass > 0 {
		return e.constantPool.getClassName(e.superClass)
	}
	// 没有超类的情况，Java中只有java.lang.Object无超类
	return ""
}

func (e *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(e.interfaces))

	for i, cpIndex := range e.interfaces {
		interfaceNames[i] = e.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
