package classfile

import "fmt"

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
