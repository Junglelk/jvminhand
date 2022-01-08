package classfile

// MemberInfo 字段和方法表。字段和方法的基本结构大致相同，不同的仅有属性表
type MemberInfo struct {
	cp ConstantPool
	// 访问标志
	accessFlags uint16
	// 常量池索引，给出字段名或方法名
	nameIndex uint16
	// 常量池索引，给出字段或方法的描述符
	descriptorIndex uint16
	// 属性表
	attribute []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	// 新建一个结构类型，使用结构字面量来初始化，& 符号表示取指针，整个语句表示返回一个MemberInfo指针
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attribute:       readAttributes(reader, cp),
	}
}

func (e *MemberInfo) AccessFlags() uint16 {
	return e.accessFlags
}

// Name 从常量池中寻找字段或方法名
func (e *MemberInfo) Name() string {
	return e.cp.getUtf8(e.nameIndex)
}

// Descriptor 从常量池中查找字段或方法描述符
func (e *MemberInfo) Descriptor() string {
	return e.cp.getUtf8(e.descriptorIndex)

}
