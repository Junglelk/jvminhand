package heap

import "jvmgo/jvmgo/ch06/classfile"

/*
	字段信息
*/

// Field 目前信息完全由类成员信息构成
type Field struct {
	ClassMember
	slotId uint
}

// newFields 创建字段表 MemberInfo 字段和方法表。字段和方法的基本结构大致相同，不同的仅有属性表
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}
