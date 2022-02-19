package heap

import "jvmgo/jvmgo/ch06/classfile"

/*
	方法信息
*/

// Method maxStack和maxLocals字段分别存放操作数栈和局部变量表大小。这两个值是由编译器计算好的。code字段存放方法字节码
type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func (e *Method) copyAttributes(method *classfile.MemberInfo) {
	if codeAttr := method.CodeAttribute(); codeAttr != nil {
		e.maxStack = codeAttr.MaxStack()
		e.maxLocals = codeAttr.MaxLocals()
		e.code = codeAttr.Code()
	}
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}
