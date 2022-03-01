package heap

import (
	"jvmgo/jvmgo/ch06/classfile"
	"strings"
)

/*
	方法区是运行时数据区的一块逻辑区域，由多个线程共享。
	方法区主要存放从class文件获取的类信息。类变量也存放在方法区内。
	虚拟机第一次使用到某一个类时会搜索类路径，找到相应的class文件，然后读取并解析class文件，
	把相关信息存放在方法区。虚拟机规范并没有规定方法区的具体位置，也没有规定大小信息、是否参与垃圾回收、方法区如何存放类数据。
*/

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (e *Class) IsPublic() bool {
	return 0 != e.accessFlags&ACC_PUBLIC
}
func (e *Class) IsFinal() bool {
	return 0 != e.accessFlags&ACC_FINAL
}

func (e *Class) IsSuper() bool {
	return 0 != e.accessFlags&ACC_SUPER
}

func (e *Class) IsInterface() bool {
	return 0 != e.accessFlags&ACC_INTERFACE
}

func (e *Class) IsAbstract() bool {
	return 0 != e.accessFlags&ACC_ABSTRACT
}
func (e *Class) IsSynthetic() bool {
	return 0 != e.accessFlags&ACC_SYNTHETIC
}
func (e *Class) IsAnnotation() bool {
	return 0 != e.accessFlags&ACC_ANNOTATION
}
func (e *Class) IsEnum() bool {
	return 0 != e.accessFlags&ACC_ENUM
}

func (e *Class) IsAccessibleTo(d *Class) bool {
	// 类d如果想访问类e，需要满足以下两个条件之一
	// 1. e是public的
	// 2. 两个类在同一个包里
	return e.IsPublic() || e.getPackageName() == d.getPackageName()
}

func (e *Class) getPackageName() string {
	if index := strings.LastIndex(e.name, "/"); index > 0 {
		return e.name[:index]
	}
	return ""
}
