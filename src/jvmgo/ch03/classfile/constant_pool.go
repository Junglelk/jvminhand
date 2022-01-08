package classfile

// ConstantPool 使用type关键字来定义一个类型，用Java的话讲，就是创建一个类，这个类可以是结构体（以struct开头，后跟大括号），可以是某种类型
type ConstantPool []ConstantInfo

// 这是一个函数（废话）
func readConstantPool(reader *ClassReader) ConstantPool {

}

// 这是一个方法。方法本质上就是函数，但方法有接收者，接收者类型就是 func 关键字和函数名中间的小括号内容
func (e ConstantPool) getConstantInfo(index uint16) ContantInfo {

}

func (e ConstantPool) getNameAndType(index uint16) (string, string) {

}
