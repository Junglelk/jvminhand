package heap

import "jvmgo/jvmgo/ch06/classfile"

/*
	字段和方法都属于类成员，它们的访问标识符、名字、描述符是一致的。
	下面结构体中class字段决定了可以通过方法或字段获取到它所属的类。

	方法描述符比方法签名多了一个返回值。
*/

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (e *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	e.accessFlags = memberInfo.AccessFlags()
	e.name = memberInfo.Name()
	e.descriptor = memberInfo.Descriptor()
}

// 下面是常规的getter和访问控制符判断方法

func (e *ClassMember) AccessFlags() uint16 {
	return e.accessFlags
}

func (e *ClassMember) Name() string {
	return e.name
}

func (e *ClassMember) Descriptor() string {
	return e.descriptor
}

func (e *ClassMember) IsPublic() bool {
	return 0 != e.accessFlags&ACC_PUBLIC
}
func (e *ClassMember) IsFinal() bool {
	return 0 != e.accessFlags&ACC_FINAL
}

func (e *ClassMember) IsSuper() bool {
	return 0 != e.accessFlags&ACC_SUPER
}

func (e *ClassMember) IsInterface() bool {
	return 0 != e.accessFlags&ACC_INTERFACE
}

func (e *ClassMember) IsAbstract() bool {
	return 0 != e.accessFlags&ACC_ABSTRACT
}
func (e *ClassMember) IsSynthetic() bool {
	return 0 != e.accessFlags&ACC_SYNTHETIC
}
func (e *ClassMember) IsAnnotation() bool {
	return 0 != e.accessFlags&ACC_ANNOTATION
}
func (e *ClassMember) IsEnum() bool {
	return 0 != e.accessFlags&ACC_ENUM
}
