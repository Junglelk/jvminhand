package heap

import "jvmgo/jvmgo/ch06/classfile"

// MemberRef 成员引用，显然，除了通用的之外还添加了成员名和描述符
type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (e *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	e.className = refInfo.ClassName()
	e.name, e.descriptor = refInfo.NameAndDescriptor()
}
