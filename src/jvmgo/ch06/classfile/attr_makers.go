package classfile

/*
	Deprecated 和 Synthetic 是最简单的两个属性，仅起到标记作用。
	可以出现在ClassFile、field_info和method_info结构中。
*/

// DeprecatedAttribute 此属性用于标记该类、接口、方法、字段已经不建议使用
/*
	Deprecated_attribute{
		u2 attribute_name_index;
		u4 attribute_length;
	}
*/
// 因为不包含任何数据，所以attribute_length的值必须是 0
type DeprecatedAttribute struct {
	MarkerAttribute
}

// SyntheticAttribute 此属性用于标记源文件中不存在，由编译器生成的类成员，
// 引入Synthetic属性的主要是为了支持嵌套类和嵌套接口。
/*
	Synthetic_attribute{
		u2 attribute_name_index;
		u4 attribute_length;
	}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (e *MarkerAttribute) readInfo(reader *ClassReader) {
	// 因为两个属性没有内容，所以readInfo方法为空
}
